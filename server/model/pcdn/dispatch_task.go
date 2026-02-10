package pcdn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

const (
	DispatchStatusPending  = "pending"
	DispatchStatusRunning  = "running"
	DispatchStatusSuccess  = "success"
	DispatchStatusFailed   = "failed"
	DispatchStatusRetrying = "retrying"
)

// PcdnDispatchTask 记录调度计划及下发执行结果，支持审计与重放。
type PcdnDispatchTask struct {
	global.GVA_MODEL
	TaskID           string         `json:"taskId" gorm:"column:task_id;size:64;not null;uniqueIndex:uk_task_trace,priority:1;index:idx_task_trace,priority:1"`
	TraceID          string         `json:"traceId" gorm:"column:trace_id;size:64;not null;uniqueIndex:uk_task_trace,priority:2;index:idx_task_trace,priority:2"`
	ContentID        string         `json:"contentId" gorm:"column:content_id;size:128;not null;index"`
	UserRegion       string         `json:"userRegion" gorm:"column:user_region;size:64;not null;index"`
	UserISP          string         `json:"userIsp" gorm:"column:user_isp;size:64;default:''"`
	TopN             int            `json:"topN" gorm:"column:top_n;default:3"`
	Status           string         `json:"status" gorm:"column:status;size:32;not null;index"`
	RetryCount       int            `json:"retryCount" gorm:"column:retry_count;default:0"`
	MaxRetry         int            `json:"maxRetry" gorm:"column:max_retry;default:3"`
	TimeoutSeconds   int            `json:"timeoutSeconds" gorm:"column:timeout_seconds;default:8"`
	NextRetryUnix    int64          `json:"nextRetryUnix" gorm:"column:next_retry_unix;default:0;index"`
	LastError        string         `json:"lastError" gorm:"column:last_error;size:1024;default:''"`
	Candidates       common.JSONMap `json:"candidates" gorm:"column:candidates;type:json"`
	MetricsSnapshot  common.JSONMap `json:"metricsSnapshot" gorm:"column:metrics_snapshot;type:json"`
	DispatchProtocol string         `json:"dispatchProtocol" gorm:"column:dispatch_protocol;size:64;default:'mock'"`
	PrimaryNodeID    uint           `json:"primaryNodeId" gorm:"column:primary_node_id;default:0"`
	CurrentNodeID    uint           `json:"currentNodeId" gorm:"column:current_node_id;default:0"`
}

func (PcdnDispatchTask) TableName() string {
	return "pcdn_dispatch_task"
}
