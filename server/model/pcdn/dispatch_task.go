package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// PcdnDispatchTask PCDN调度任务
type PcdnDispatchTask struct {
	global.GVA_MODEL
	TraceID    string `json:"traceId" form:"traceId" gorm:"column:trace_id;size:128;index;comment:追踪ID"`
	Status     string `json:"status" form:"status" gorm:"column:status;size:32;index;comment:任务状态"`
	NodeID     uint   `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:节点ID"`
	PolicyID   uint   `json:"policyId" form:"policyId" gorm:"column:policy_id;comment:策略ID"`
	RequestURI string `json:"requestUri" form:"requestUri" gorm:"column:request_uri;size:255;comment:请求路径"`
	Result     string `json:"result" form:"result" gorm:"column:result;type:text;comment:执行结果"`
}

func (PcdnDispatchTask) TableName() string {
	return "pcdn_dispatch_task"
}
