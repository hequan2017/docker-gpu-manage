package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"gorm.io/gorm"
)

var (
	// ErrPermissionDenied 权限拒绝错误
	ErrPermissionDenied = errors.New("permission denied")
	// ErrPermissionNotFound 权限不存在错误
	ErrPermissionNotFound = errors.New("permission not found")
)

// K8sPermissionService K8s权限服务
type K8sPermissionService struct{}

// CheckPermission 检查用户是否有权限执行操作
func (s *K8sPermissionService) CheckPermission(check *model.K8sPermissionCheck) error {
	// 获取用户的所有权限
	permissions, err := s.getUserPermissions(check.UserID, check.UserRoleIDs)
	if err != nil {
		return err
	}

	// 检查是否有任何权限允许该操作
	for _, perm := range permissions {
		if !perm.Enabled {
			continue
		}

		if s.matchPermission(&perm, check) {
			return nil // 有权限
		}
	}

	return ErrPermissionDenied
}

// HasPermission 检查是否有权限（返回bool）
func (s *K8sPermissionService) HasPermission(check *model.K8sPermissionCheck) bool {
	return s.CheckPermission(check) == nil
}

// getUserPermissions 获取用户的所有权限（包括角色权限和直接授予的权限）
func (s *K8sPermissionService) getUserPermissions(userID uint, roleIDs []uint) ([]model.K8sPermission, error) {
	var permissions []model.K8sPermission

	// 1. 获取直接授予用户的权限
	var userPerms []model.K8sPermission
	err := global.GVA_DB.Table("k8s_permissions").
		Joins("JOIN k8s_user_permissions ON k8s_user_permissions.permission_id = k8s_permissions.id").
		Where("k8s_user_permissions.user_id = ?", userID).
		Where("(k8s_user_permissions.expires_at IS NULL OR k8s_user_permissions.expires_at > ?)", time.Now()).
		Find(&userPerms).Error
	if err != nil {
		return nil, err
	}
	permissions = append(permissions, userPerms...)

	// 2. 获取通过角色授予的权限
	if len(roleIDs) > 0 {
		var rolePerms []model.K8sPermission
		err = global.GVA_DB.Table("k8s_permissions").
			Joins("JOIN k8s_role_permissions ON k8s_role_permissions.permission_id = k8s_permissions.id").
			Where("k8s_role_permissions.role_id IN ?", roleIDs).
			Find(&rolePerms).Error
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, rolePerms...)
	}

	return permissions, nil
}

// matchPermission 检查权限是否匹配请求
func (s *K8sPermissionService) matchPermission(perm *model.K8sPermission, check *model.K8sPermissionCheck) bool {
	// 1. 检查权限范围
	if !s.matchScope(perm, check) {
		return false
	}

	// 2. 检查资源类型
	if !s.matchResourceType(perm, check.Resource) {
		return false
	}

	// 3. 检查操作权限
	if !s.matchAction(perm, check.Action) {
		return false
	}

	return true
}

// matchScope 检查权限范围是否匹配
func (s *K8sPermissionService) matchScope(perm *model.K8sPermission, check *model.K8sPermissionCheck) bool {
	switch perm.ScopeType {
	case "global":
		// 全局权限，允许所有集群和命名空间
		return true
	case "cluster":
		// 集群级别权限，检查集群名称
		return perm.ClusterName == check.ClusterName
	case "namespace":
		// 命名空间级别权限，检查集群和命名空间
		return perm.ClusterName == check.ClusterName && perm.Namespace == check.Namespace
	default:
		return false
	}
}

// matchResourceType 检查资源类型是否匹配
func (s *K8sPermissionService) matchResourceType(perm *model.K8sPermission, resource string) bool {
	var rts model.ResourceTypeSet
	if err := json.Unmarshal([]byte(perm.ResourceTypes), &rts); err != nil {
		return false
	}

	// 检查是否在排除列表中
	for _, exclude := range rts.Exclude {
		if exclude == resource || exclude == "*" {
			return false
		}
	}

	// 检查是否在包含列表中
	for _, include := range rts.Include {
		if include == "*" || include == resource {
			return true
		}
	}

	return false
}

// matchAction 检查操作是否匹配
func (s *K8sPermissionService) matchAction(perm *model.K8sPermission, action string) bool {
	var as model.ActionSet
	if err := json.Unmarshal([]byte(perm.Actions), &as); err != nil {
		return false
	}

	// 检查是否在排除列表中
	for _, exclude := range as.Exclude {
		if exclude == action || exclude == "*" {
			return false
		}
	}

	// 检查是否在包含列表中
	for _, include := range as.Include {
		if include == "*" || include == action {
			return true
		}
	}

	return false
}

// ApplyConstraints 应用权限约束到查询参数
func (s *K8sPermissionService) ApplyConstraints(check *model.K8sPermissionCheck, queryParams map[string]interface{}) (map[string]interface{}, error) {
	permissions, err := s.getUserPermissions(check.UserID, check.UserRoleIDs)
	if err != nil {
		return nil, err
	}

	// 合并所有匹配的权限约束
	var allConstraints []model.PermissionConstraint
	for _, perm := range permissions {
		if !perm.Enabled || !s.matchPermission(&perm, check) {
			continue
		}

		if perm.Constraints != "" {
			var constraints model.PermissionConstraint
			if err := json.Unmarshal([]byte(perm.Constraints), &constraints); err == nil {
				allConstraints = append(allConstraints, constraints)
			}
		}
	}

	// 应用约束到查询参数
	result := make(map[string]interface{})
	for k, v := range queryParams {
		result[k] = v
	}

	// 如果有任何约束，应用它们
	if len(allConstraints) > 0 {
		// 标签选择器（取交集）
		labelSelectors := make(map[string]string)
		for _, constraints := range allConstraints {
			for k, v := range constraints.LabelSelector {
				labelSelectors[k] = v
			}
		}
		if len(labelSelectors) > 0 {
			result["labelSelector"] = labelSelectors
		}

		// 名称白名单（取并集）
		var names []string
		for _, constraints := range allConstraints {
			names = append(names, constraints.Names...)
		}
		if len(names) > 0 {
			result["names"] = names
		}

		// 名称黑名单
		var excludeNames []string
		for _, constraints := range allConstraints {
			excludeNames = append(excludeNames, constraints.ExcludeNames...)
		}
		if len(excludeNames) > 0 {
			result["excludeNames"] = excludeNames
		}
	}

	return result, nil
}

// CreatePermission 创建权限
func (s *K8sPermissionService) CreatePermission(perm *model.K8sPermission) error {
	return global.GVA_DB.Create(perm).Error
}

// UpdatePermission 更新权限
func (s *K8sPermissionService) UpdatePermission(perm *model.K8sPermission) error {
	return global.GVA_DB.Save(perm).Error
}

// DeletePermission 删除权限
func (s *K8sPermissionService) DeletePermission(id uint) error {
	// 先删除关联
	global.GVA_DB.Where("permission_id = ?", id).Delete(&model.K8sRolePermission{})
	global.GVA_DB.Where("permission_id = ?", id).Delete(&model.K8sUserPermission{})
	// 再删除权限
	return global.GVA_DB.Delete(&model.K8sPermission{}, id).Error
}

// GetPermission 获取权限详情
func (s *K8sPermissionService) GetPermission(id uint) (*model.K8sPermission, error) {
	var perm model.K8sPermission
	err := global.GVA_DB.First(&perm, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPermissionNotFound
		}
		return nil, err
	}
	return &perm, nil
}

// GrantPermissionToRole 将权限授予角色
func (s *K8sPermissionService) GrantPermissionToRole(permissionID, roleID uint) error {
	// 检查是否已存在
	var count int64
	global.GVA_DB.Model(&model.K8sRolePermission{}).
		Where("permission_id = ? AND role_id = ?", permissionID, roleID).
		Count(&count)
	if count > 0 {
		return nil // 已存在
	}

	return global.GVA_DB.Create(&model.K8sRolePermission{
		PermissionID: permissionID,
		RoleID:       roleID,
	}).Error
}

// RevokePermissionFromRole 撤销角色的权限
func (s *K8sPermissionService) RevokePermissionFromRole(permissionID, roleID uint) error {
	return global.GVA_DB.Where("permission_id = ? AND role_id = ?", permissionID, roleID).
		Delete(&model.K8sRolePermission{}).Error
}

// GrantPermissionToUser 将权限直接授予用户
func (s *K8sPermissionService) GrantPermissionToUser(permissionID, userID uint, expiresAt *int64) error {
	// 检查是否已存在
	var count int64
	global.GVA_DB.Model(&model.K8sUserPermission{}).
		Where("permission_id = ? AND user_id = ?", permissionID, userID).
		Count(&count)
	if count > 0 {
		return fmt.Errorf("用户已有该权限")
	}

	return global.GVA_DB.Create(&model.K8sUserPermission{
		PermissionID: permissionID,
		UserID:       userID,
		ExpiresAt:    expiresAt,
	}).Error
}

// RevokePermissionFromUser 撤销用户的权限
func (s *K8sPermissionService) RevokePermissionFromUser(permissionID, userID uint) error {
	return global.GVA_DB.Where("permission_id = ? AND user_id = ?", permissionID, userID).
		Delete(&model.K8sUserPermission{}).Error
}
