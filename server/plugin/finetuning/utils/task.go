package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	finetuningModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model/request"
	"go.uber.org/zap"
)

// TaskUtil 任务工具类
type TaskUtil struct{}

// BuildCommandFromRequest 从请求构建命令
func (tu *TaskUtil) BuildCommandFromRequest(req *request.CreateFinetuningTaskRequest, task *finetuningModel.FinetuningTask) error {
	// 否则使用默认模板构建命令
	baseCmd := "python train.py"
	args := []string{
		fmt.Sprintf("--base_model=%s", req.BaseModel),
		fmt.Sprintf("--data_path=%s", req.DatasetPath),
	}

	// 添加输出路径
	if req.OutputPath != "" {
		args = append(args, fmt.Sprintf("--output_dir=%s", req.OutputPath))
	} else {
		// 使用默认输出路径
		defaultOutputPath := filepath.Join(global.GVA_CONFIG.Local.StorePath, "finetuning_outputs",
			fmt.Sprintf("%s_%d", req.Name, time.Now().Unix()))
		task.OutputPath = &defaultOutputPath
		args = append(args, fmt.Sprintf("--output_dir=%s", defaultOutputPath))
	}

	// 添加训练参数
	if len(req.TrainingArgs) > 0 {
		for key, value := range req.TrainingArgs {
			args = append(args, fmt.Sprintf("--%s=%v", key, value))
		}
	}

	// 组合命令
	command := fmt.Sprintf("%s %s", baseCmd, strings.Join(args, " "))
	task.Command = &command

	return nil
}

// ValidateRequest 验证请求参数
func (tu *TaskUtil) ValidateRequest(req *request.CreateFinetuningTaskRequest) error {
	if req.Name == "" {
		return fmt.Errorf("任务名称不能为空")
	}
	if req.BaseModel == "" {
		return fmt.Errorf("基础模型不能为空")
	}
	if req.DatasetPath == "" {
		return fmt.Errorf("数据集路径不能为空")
	}

	// 验证数据集路径是否存在
	if _, err := os.Stat(req.DatasetPath); os.IsNotExist(err) {
		return fmt.Errorf("数据集路径不存在: %s", req.DatasetPath)
	}

	// 验证基础模型路径是否存在（如果是本地路径）
	if !strings.HasPrefix(req.BaseModel, "hf://") && !strings.HasPrefix(req.BaseModel, "http://") && !strings.HasPrefix(req.BaseModel, "https://") {
		if _, err := os.Stat(req.BaseModel); os.IsNotExist(err) {
			return fmt.Errorf("基础模型路径不存在: %s", req.BaseModel)
		}
	}

	return nil
}

// CreateLogDir 创建日志目录
func (tu *TaskUtil) CreateLogDir(taskName string) (string, error) {
	logDir := filepath.Join(global.GVA_CONFIG.Local.StorePath, "finetuning_logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", fmt.Errorf("创建日志目录失败: %w", err)
	}

	logFileName := fmt.Sprintf("task_%s_%s.log", taskName, time.Now().Format("20060102_150405"))
	logPath := filepath.Join(logDir, logFileName)

	return logPath, nil
}

// SerializeConfig 序列化配置
func (tu *TaskUtil) SerializeConfig(config map[string]interface{}) (*string, error) {
	if len(config) == 0 {
		return nil, nil
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("序列化配置失败: %w", err)
	}

	configStr := string(jsonData)
	return &configStr, nil
}

// DeserializeConfig 反序列化配置
func (tu *TaskUtil) DeserializeConfig(configStr *string) (map[string]interface{}, error) {
	if configStr == nil || *configStr == "" {
		return nil, nil
	}

	var config map[string]interface{}
	if err := json.Unmarshal([]byte(*configStr), &config); err != nil {
		return nil, fmt.Errorf("反序列化配置失败: %w", err)
	}

	return config, nil
}

// CalculateProgress 计算任务进度（基于时间估算）
func (tu *TaskUtil) CalculateProgress(task *finetuningModel.FinetuningTask) float64 {
	if task.Status != finetuningModel.TaskStatusRunning {
		if task.Status == finetuningModel.TaskStatusCompleted {
			return 100.0
		}
		return task.Progress != nil ? *task.Progress : 0
	}

	// 如果有开始时间，基于时间估算进度（仅作参考）
	if task.StartedAt != nil {
		elapsed := time.Now().Unix() - *task.StartedAt
		// 假设训练需要2小时，可以基于实际情况调整
		estimatedDuration := int64(2 * 60 * 60) // 2小时
		progress := float64(elapsed) / float64(estimatedDuration) * 100
		if progress > 99 {
			progress = 99
		}
		return progress
	}

	return 0
}

// FormatDuration 格式化持续时间
func (tu *TaskUtil) FormatDuration(seconds int64) string {
	if seconds <= 0 {
		return "-"
	}

	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%d小时%d分%d秒", hours, minutes, secs)
	} else if minutes > 0 {
		return fmt.Sprintf("%d分%d秒", minutes, secs)
	}
	return fmt.Sprintf("%d秒", secs)
}

// GetTaskStatusText 获取任务状态文本
func (tu *TaskUtil) GetTaskStatusText(status string) string {
	statusMap := map[string]string{
		"pending":   "待执行",
		"running":   "执行中",
		"completed": "已完成",
		"failed":    "失败",
		"stopped":   "已停止",
	}
	if text, ok := statusMap[status]; ok {
		return text
	}
	return status
}

// GetTaskStatusType 获取任务状态类型（用于前端显示）
func (tu *TaskUtil) GetTaskStatusType(status string) string {
	typeMap := map[string]string{
		"pending":   "info",
		"running":   "primary",
		"completed": "success",
		"failed":    "danger",
		"stopped":   "warning",
	}
	if t, ok := typeMap[status]; ok {
		return t
	}
	return "info"
}

// CleanupTask 清理任务资源
func (tu *TaskUtil) CleanupTask(ctx context.Context, taskID uint) error {
	// 获取任务信息
	var task finetuningModel.FinetuningTask
	if err := global.GVA_DB.WithContext(ctx).Where("id = ?", taskID).First(&task).Error; err != nil {
		return err
	}

	// 删除日志文件
	if task.LogPath != nil {
		if err := os.Remove(*task.LogPath); err != nil && !os.IsNotExist(err) {
			global.GVA_LOG.Warn("删除日志文件失败", zap.String("path", *task.LogPath), zap.Error(err))
		}
	}

	// 可选：删除输出文件
	// if task.OutputPath != nil {
	//     if err := os.RemoveAll(*task.OutputPath); err != nil {
	//         global.GVA_LOG.Warn("删除输出目录失败", zap.String("path", *task.OutputPath), zap.Error(err))
	//     }
	// }

	return nil
}
