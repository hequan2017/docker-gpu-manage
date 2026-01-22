package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// K8sPermission K8s资源权限定义
type K8sPermission struct {
	global.GVA_MODEL
	// 基本信息
	Name        string `json:"name" gorm:"type:varchar(100);not null;index;comment:权限名称"`
	Description string `json:"description" gorm:"type:varchar(500);comment:权限描述"`
	Enabled     bool   `json:"enabled" gorm:"type:tinyint(1);default:1;comment:是否启用"`

	// 权限范围
	ScopeType   string `json:"scopeType" gorm:"type:varchar(20);index;comment:权限范围类型"` // global, cluster, namespace
	ClusterName string `json:"clusterName" gorm:"type:varchar(100);index;comment:集群名称"`
	Namespace   string `json:"namespace" gorm:"type:varchar(100);index;comment:命名空间"`

	// 资源类型
	ResourceTypes string `json:"resourceTypes" gorm:"type:text;comment:资源类型列表(JSON)"` // ["pod", "deployment", "service"]

	// 操作权限
	Actions string `json:"actions" gorm:"type:text;comment:允许的操作列表(JSON)"` // ["get", "list", "create", "update", "delete", "exec", "scale"]

	// 约束条件
	Constraints string `json:"constraints" gorm:"type:text;comment:额外约束条件(JSON)"` // 如标签选择器等
}

// TableName 指定表名
func (K8sPermission) TableName() string {
	return "k8s_permissions"
}

// K8sRolePermission K8s角色权限关联
type K8sRolePermission struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	PermissionID uint   `json:"permissionId" gorm:"index;comment:权限ID"`
	RoleID       uint   `json:"roleId" gorm:"index;comment:角色ID"`
	CreatedAt    int64  `json:"createdAt" gorm:"autoCreateTime;comment:创建时间"`
}

// TableName 指定表名
func (K8sRolePermission) TableName() string {
	return "k8s_role_permissions"
}

// K8sUserPermission K8s用户权限关联（直接授权）
type K8sUserPermission struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	PermissionID uint   `json:"permissionId" gorm:"index;comment:权限ID"`
	UserID       uint   `json:"userId" gorm:"index;comment:用户ID"`
	ExpiresAt    *int64 `json:"expiresAt" gorm:"index;comment:过期时间(可选)"`
	CreatedAt    int64  `json:"createdAt" gorm:"autoCreateTime;comment:创建时间"`
}

// TableName 指定表名
func (K8sUserPermission) TableName() string {
	return "k8s_user_permissions"
}

// K8sPermissionCheck 权限检查请求
type K8sPermissionCheck struct {
	UserID      uint   `json:"userId" binding:"required"`      // 用户ID
	UserRoleIDs []uint `json:"userRoleIds"`                   // 用户角色ID列表
	Action      string `json:"action" binding:"required"`     // 操作类型: get, list, create, update, delete, exec, scale
	Resource    string `json:"resource" binding:"required"`   // 资源类型: cluster, pod, deployment, service, namespace
	ClusterName string `json:"clusterName"`                  // 集群名称
	Namespace   string `json:"namespace"`                    // 命名空间
}

// PermissionConstraint 权限约束
type PermissionConstraint struct {
	LabelSelector map[string]string `json:"labelSelector,omitempty"` // 标签选择器
	FieldSelector string            `json:"fieldSelector,omitempty"` // 字段选择器
	Names        []string          `json:"names,omitempty"`         // 资源名称白名单
	ExcludeNames []string          `json:"excludeNames,omitempty"`  // 资源名称黑名单
}

// ResourceTypeSet 资源类型集合
type ResourceTypeSet struct {
	Include []string `json:"include"` // 包含的资源类型
	Exclude []string `json:"exclude"` // 排除的资源类型
}

// ActionSet 操作集合
type ActionSet struct {
	Include []string `json:"include"` // 允许的操作
	Exclude []string `json:"exclude"` // 禁止的操作
}
