package pcdn

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	pcdnModel "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	"go.uber.org/zap"
)

type RuntimeService struct {
	dispatcher Dispatcher
}

func NewRuntimeService(dispatcher Dispatcher) *RuntimeService {
	if dispatcher == nil {
		dispatcher = &MockDispatcher{}
	}
	return &RuntimeService{dispatcher: dispatcher}
}

// ProcessPendingTasks 扫描待处理任务，支持超时、重试和幂等状态推进。
func (r *RuntimeService) ProcessPendingTasks(ctx context.Context, limit int) error {
	if limit <= 0 {
		limit = 20
	}
	now := time.Now().Unix()
	var tasks []pcdnModel.PcdnDispatchTask
	err := global.GVA_DB.WithContext(ctx).
		Where("status IN ?", []string{pcdnModel.DispatchStatusPending, pcdnModel.DispatchStatusRetrying}).
		Where("next_retry_unix = 0 OR next_retry_unix <= ?", now).
		Order("created_at ASC").
		Limit(limit).
		Find(&tasks).Error
	if err != nil {
		return err
	}
	for i := range tasks {
		if err := r.handleTask(ctx, &tasks[i]); err != nil {
			global.GVA_LOG.Warn("处理PCDN任务失败", zap.Uint("id", tasks[i].ID), zap.Error(err))
		}
	}
	return nil
}

func (r *RuntimeService) handleTask(ctx context.Context, task *pcdnModel.PcdnDispatchTask) error {
	locked, err := r.markRunning(ctx, task.ID)
	if err != nil || !locked {
		return err
	}

	result := r.dispatcher.Dispatch(ctx, DispatchRequest{
		TaskID:     task.TaskID,
		TraceID:    task.TraceID,
		ContentID:  task.ContentID,
		TargetNode: task.CurrentNodeID,
		TimeoutSec: task.TimeoutSeconds,
	})

	if result.Success {
		return global.GVA_DB.WithContext(ctx).
			Model(&pcdnModel.PcdnDispatchTask{}).
			Where("id = ?", task.ID).
			Updates(map[string]any{"status": pcdnModel.DispatchStatusSuccess, "last_error": ""}).Error
	}

	nextNode := NextFallbackNode(*task)
	retryCount := task.RetryCount + 1
	updates := map[string]any{
		"retry_count": retryCount,
		"last_error":  result.Message,
	}
	if nextNode != 0 {
		updates["current_node_id"] = nextNode
	}
	if retryCount <= task.MaxRetry {
		updates["status"] = pcdnModel.DispatchStatusRetrying
		updates["next_retry_unix"] = time.Now().Unix() + int64(retryCount*2)
	} else {
		updates["status"] = pcdnModel.DispatchStatusFailed
		if nextNode == 0 {
			updates["last_error"] = fmt.Sprintf("%s; no fallback node available", result.Message)
		}
	}
	return global.GVA_DB.WithContext(ctx).
		Model(&pcdnModel.PcdnDispatchTask{}).
		Where("id = ?", task.ID).
		Updates(updates).Error
}

func (r *RuntimeService) markRunning(ctx context.Context, id uint) (bool, error) {
	res := global.GVA_DB.WithContext(ctx).
		Model(&pcdnModel.PcdnDispatchTask{}).
		Where("id = ? AND status IN ?", id, []string{pcdnModel.DispatchStatusPending, pcdnModel.DispatchStatusRetrying}).
		Updates(map[string]any{"status": pcdnModel.DispatchStatusRunning, "next_retry_unix": 0})
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected == 1, nil
}

// SyncTimeoutTasks 将长时间 Running 的任务回收为重试态。
func (r *RuntimeService) SyncTimeoutTasks(ctx context.Context) error {
	var running []pcdnModel.PcdnDispatchTask
	if err := global.GVA_DB.WithContext(ctx).Where("status = ?", pcdnModel.DispatchStatusRunning).Find(&running).Error; err != nil {
		return err
	}
	now := time.Now().Unix()
	for _, t := range running {
		deadline := t.UpdatedAt.Unix() + int64(maxInt(t.TimeoutSeconds, 8))
		if now <= deadline {
			continue
		}
		status := pcdnModel.DispatchStatusRetrying
		if t.RetryCount >= t.MaxRetry {
			status = pcdnModel.DispatchStatusFailed
		}
		nextRetry := int64(0)
		if status == pcdnModel.DispatchStatusRetrying {
			nextRetry = now + 2
		}
		err := global.GVA_DB.WithContext(ctx).
			Model(&pcdnModel.PcdnDispatchTask{}).
			Where("id = ?", t.ID).
			Updates(map[string]any{
				"status":          status,
				"retry_count":     t.RetryCount + 1,
				"next_retry_unix": nextRetry,
				"last_error":      "dispatch timeout",
			}).Error
		if err != nil {
			global.GVA_LOG.Warn("回收PCDN超时任务失败", zap.Uint("task_id", t.ID), zap.Error(err))
		}
	}
	return nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var PcdnRuntimeService = NewRuntimeService(nil)
