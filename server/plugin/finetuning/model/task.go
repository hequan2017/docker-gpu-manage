package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FinetuningTask 算法微调任务 结构体
type FinetuningTask struct {
	global.GVA_MODEL
	Name         string     `json:"name" form:"name" gorm:"column:name;comment:任务名称;type:varchar(200);not null"`                                        // 任务名称
	Description  *string    `json:"description" form:"description" gorm:"column:description;comment:任务描述;type:text"`                                     // 任务描述
	UserID       *int       `json:"userID" form:"userID" gorm:"column:user_id;comment:所属用户;index"`                                                      // 所属用户
	Status       string     `json:"status" form:"status" gorm:"column:status;comment:任务状态;type:varchar(20);default:pending;index"`                      // 任务状态: pending, running, completed, failed, stopped
	Progress     *float64   `json:"progress" form:"progress" gorm:"column:progress;comment:任务进度;default:0"`                                              // 任务进度 0-100
	BaseModel    string     `json:"baseModel" form:"baseModel" gorm:"column:base_model;comment:基础模型;type:varchar(200);not null"`                         // 基础模型路径或名称
	DatasetPath  string     `json:"datasetPath" form:"datasetPath" gorm:"column:dataset_path;comment:数据集路径;type:varchar(500);not null"`                 // 数据集路径
	OutputPath   *string    `json:"outputPath" form:"outputPath" gorm:"column:output_path;comment:输出路径;type:varchar(500)"`                                // 输出模型路径
	TrainingArgs *string    `json:"trainingArgs" form:"trainingArgs" gorm:"column:training_args;comment:训练参数;type:json"`                                  // 训练参数JSON配置
	GPUConfig    *string    `json:"gpuConfig" form:"gpuConfig" gorm:"column:gpu_config;comment:GPU配置;type:json"`                                          // GPU配置JSON配置
	Command      *string    `json:"command" form:"command" gorm:"column:command;comment:执行命令;type:text"`                                                  // 执行的完整命令
	LogPath      *string    `json:"logPath" form:"logPath" gorm:"column:log_path;comment:日志文件路径;type:varchar(500)"`                                     // 日志文件路径
	ErrorMessage *string    `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:错误信息;type:text"`                                  // 错误信息
	StartedAt    *int64     `json:"startedAt" form:"startedAt" gorm:"column:started_at;comment:开始时间"`                                                    // 开始时间戳
	FinishedAt   *int64     `json:"finishedAt" form:"finishedAt" gorm:"column:finished_at;comment:结束时间"`                                                  // 结束时间戳
	Pid          *int       `json:"pid" form:"pid" gorm:"column:pid;comment:进程ID"`                                                                        // 进程ID
	Metrics      *string    `json:"metrics" form:"metrics" gorm:"column:metrics;comment:训练指标;type:json"`                                                 // 训练指标JSON配置
}

// TableName FinetuningTask 自定义表名 gva_finetuning_tasks
func (FinetuningTask) TableName() string {
	return "gva_finetuning_tasks"
}

// TaskStatus 任务状态常量
const (
	TaskStatusPending   = "pending"   // 待执行
	TaskStatusRunning   = "running"   // 执行中
	TaskStatusCompleted = "completed" // 已完成
	TaskStatusFailed    = "failed"    // 失败
	TaskStatusStopped   = "stopped"   // 已停止
)
