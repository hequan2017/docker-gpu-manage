package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	commonRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// CreateFinetuningTaskRequest 创建微调任务请求
type CreateFinetuningTaskRequest struct {
	Name         string                 `json:"name" form:"name" binding:"required"`                    // 任务名称
	Description  string                 `json:"description" form:"description"`                         // 任务描述
	BaseModel    string                 `json:"baseModel" form:"baseModel" binding:"required"`          // 基础模型
	DatasetPath  string                 `json:"datasetPath" form:"datasetPath" binding:"required"`      // 数据集路径
	OutputPath   string                 `json:"outputPath" form:"outputPath"`                           // 输出路径
	TrainingArgs map[string]interface{} `json:"trainingArgs" form:"trainingArgs"`                       // 训练参数
	GPUConfig    map[string]interface{} `json:"gpuConfig" form:"gpuConfig"`                             // GPU配置
	Command      string                 `json:"command" form:"command"`                                 // 自定义命令
}

// UpdateFinetuningTaskRequest 更新微调任务请求
type UpdateFinetuningTaskRequest struct {
	global.GVA_MODEL
	Name        string  `json:"name" form:"name"`         // 任务名称
	Description *string `json:"description" form:"description"` // 任务描述
	Status      string  `json:"status" form:"status"`     // 任务状态
}

// FinetuningTaskSearch 微调任务搜索请求
type FinetuningTaskSearch struct {
	commonRequest.PageInfo
	Name        string `form:"name" search:"type:contains;column:name;table:gva_finetuning_tasks"`                         // 任务名称(模糊查询)
	Status      string `form:"status" search:"type:exact;column:status;table:gva_finetuning_tasks"`                        // 任务状态
	BaseModel   string `form:"baseModel" search:"type:contains;column:base_model;table:gva_finetuning_tasks"`              // 基础模型
	OrderKey    string `form:"orderKey" search:"type:order;column:created_at;table:gva_finetuning_tasks"`                  // 排序字段
	Desc        bool   `form:"desc" search:"type:desc"`                                                                   // 是否倒序
}

// GetFinetuningTaskById 获取任务详情请求
type GetFinetuningTaskById struct {
	ID uint `form:"id" binding:"required"` // 任务ID
}

// StopFinetuningTaskRequest 停止任务请求
type StopFinetuningTaskRequest struct {
	ID uint `form:"id" binding:"required"` // 任务ID
}

// DeleteFinetuningTaskRequest 删除任务请求
type DeleteFinetuningTaskRequest struct {
	ID uint `form:"id" binding:"required"` // 任务ID
}

// GetTaskLogRequest 获取任务日志请求
type GetTaskLogRequest struct {
	ID    uint   `form:"id" binding:"required"` // 任务ID
	Lines *int   `form:"lines"`                  // 获取日志行数
	Offset *int  `form:"offset"`                 // 日志偏移量
}
