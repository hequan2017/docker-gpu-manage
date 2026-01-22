package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// K8sAuditLog K8s操作审计日志
type K8sAuditLog struct {
	global.GVA_MODEL
	// 用户信息
	UserID      uint   `json:"userId" gorm:"index;comment:用户ID"`
	Username    string `json:"username" gorm:"type:varchar(100);index;comment:用户名"`
	UserIP      string `json:"userIp" gorm:"type:varchar(50);comment:用户IP"`
	UserAgent   string `json:"userAgent" gorm:"type:varchar(500);comment:用户代理"`

	// 操作信息
	Action      string `json:"action" gorm:"type:varchar(100);index;comment:操作类型"` // create, delete, update, query, exec, scale, restart等
	Resource    string `json:"resource" gorm:"type:varchar(100);index;comment:资源类型"` // cluster, pod, deployment, service, namespace等
	ResourceID  string `json:"resourceId" gorm:"type:varchar(200);comment:资源ID"`
	ClusterName string `json:"clusterName" gorm:"type:varchar(100);index;comment:集群名称"`
	Namespace   string `json:"namespace" gorm:"type:varchar(100);comment:命名空间"`

	// 操作详情
	Description string `json:"description" gorm:"type:text;comment:操作描述"`
	RequestData string `json:"requestData" gorm:"type:text;comment:请求数据"`
	Response    string `json:"response" gorm:"type:text;comment:响应数据"`

	// 结果信息
	Status    string `json:"status" gorm:"type:varchar(20);index;comment:操作状态"` // success, failure, partial
	ErrorMsg  string `json:"errorMsg" gorm:"type:text;comment:错误信息"`
	Duration  int64  `json:"duration" gorm:"type:int;comment:执行时长(毫秒)"`

	// 额外信息
	Metadata   string `json:"metadata" gorm:"type:text;comment:额外元数据(JSON)"`
}

// TableName 指定表名
func (K8sAuditLog) TableName() string {
	return "k8s_audit_logs"
}

// K8sAuditLogSearch 审计日志搜索条件
type K8sAuditLogSearch struct {
	global.PageInfo
	StartTime  time.Time `json:"startTime" form:"startTime"`  // 开始时间
	EndTime    time.Time `json:"endTime" form:"endTime"`      // 结束时间
	UserID     uint      `json:"userId" form:"userId"`        // 用户ID
	Username   string    `json:"username" form:"username"`    // 用户名
	Action     string    `json:"action" form:"action"`        // 操作类型
	Resource   string    `json:"resource" form:"resource"`    // 资源类型
	Cluster    string    `json:"cluster" form:"cluster"`      // 集群名称
	Status     string    `json:"status" form:"status"`        // 操作状态
	Keyword    string    `json:"keyword" form:"keyword"`      // 关键词搜索
}
