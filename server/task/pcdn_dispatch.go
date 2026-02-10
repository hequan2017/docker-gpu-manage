package task

import (
	"context"

	pcdnService "github.com/flipped-aurora/gin-vue-admin/server/service/pcdn"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProcessPcdnDispatchTasks 执行PCDN调度任务下发。
func ProcessPcdnDispatchTasks() {
	if err := pcdnService.PcdnRuntimeService.ProcessPendingTasks(context.Background(), 50); err != nil {
		global.GVA_LOG.Error("处理PCDN调度任务失败", zap.Error(err))
	}
}

// SyncPcdnDispatchStatus 同步处理超时和重试状态。
func SyncPcdnDispatchStatus() {
	if err := pcdnService.PcdnRuntimeService.SyncTimeoutTasks(context.Background()); err != nil {
		global.GVA_LOG.Error("同步PCDN调度状态失败", zap.Error(err))
	}
}
