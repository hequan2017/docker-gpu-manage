package middleware

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// K8sPermissionMiddleware K8s权限检查中间件
func K8sPermissionMiddleware() gin.HandlerFunc {
	permService := &service.K8sPermissionService{}

	return func(c *gin.Context) {
		// 超级管理员跳过权限检查
		if userID, exists := c.Get("userId"); exists {
			if uid, ok := userID.(uint); ok {
				// 这里需要检查用户是否是超级管理员
				// 假设通过用户服务或Casbin已经做了检查
				// 如果是超级管理员，直接通过
				if isSuperAdmin(c) {
					c.Next()
					return
				}
			}
		}

		// 解析请求以确定需要的权限
		permCheck := parsePermissionRequest(c)
		if permCheck == nil {
			c.Next() // 无法解析，跳过检查
			return
		}

		// 检查权限
		if err := permService.CheckPermission(permCheck); err != nil {
			global.GVA_LOG.Warn("K8s权限检查失败",
				zap.Uint("userId", permCheck.UserID),
				zap.String("action", permCheck.Action),
				zap.String("resource", permCheck.Resource),
				zap.String("cluster", permCheck.ClusterName),
				zap.Error(err))
			response.FailWithMessage("权限不足: "+err.Error(), c)
			c.Abort()
			return
		}

		c.Next()
	}
}

// parsePermissionRequest 从请求中解析权限检查请求
func parsePermissionRequest(c *gin.Context) *model.K8sPermissionCheck {
	// 获取用户信息
	userID, exists := c.Get("userId")
	if !exists {
		return nil
	}

	uid, ok := userID.(uint)
	if !ok {
		return nil
	}

	// 获取用户角色ID列表
	var roleIDs []uint
	if roles, exists := c.Get("authorityId"); exists {
		// 这里需要根据实际情况获取角色ID列表
		// 假设从Casbin或其他地方获取
	}

	// 从路径解析操作类型和资源类型
	path := c.Request.URL.Path
	if !strings.HasPrefix(path, "/k8s/") {
		return nil
	}

	parts := strings.Split(strings.TrimPrefix(path, "/k8s/"), "/")
	if len(parts) < 2 {
		return nil
	}

	resourceType := parts[0]     // cluster, pod, deployment等
	operation := parts[1]        // create, delete, update, list等

	// 获取集群名称和命名空间
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		clusterName = c.PostForm("clusterName")
	}

	namespace := c.Query("namespace")
	if namespace == "" {
		namespace = c.PostForm("namespace")
	}

	// 将操作类型映射到权限操作
	action := mapOperationToAction(operation, c.Request.Method)

	return &model.K8sPermissionCheck{
		UserID:      uid,
		UserRoleIDs: roleIDs,
		Action:      action,
		Resource:    resourceType,
		ClusterName: clusterName,
		Namespace:   namespace,
	}
}

// mapOperationToAction 将API操作映射到权限操作
func mapOperationToAction(operation, method string) string {
	switch operation {
	case "create":
		return "create"
	case "delete", "deleteByIds":
		return "delete"
	case "update":
		return "update"
	case "get", "list", "all":
		return "get"
	case "scale":
		return "scale"
	case "restart":
		return "restart"
	case "refresh":
		return "get"
	case "log":
		return "get"
	case "events":
		return "get"
	case "containers":
		return "get"
	case "endpoints":
		return "get"
	default:
		// 根据HTTP方法推断
		switch method {
		case "POST":
			return "create"
		case "PUT", "PATCH":
			return "update"
		case "DELETE":
			return "delete"
		case "GET":
			return "get"
		default:
			return "unknown"
		}
	}
}

// isSuperAdmin 检查用户是否是超级管理员
func isSuperAdmin(c *gin.Context) bool {
	// 这里需要实现实际的超级管理员检查逻辑
	// 可能通过检查用户的authorityId或其他标识
	if authorityId, exists := c.Get("authorityId"); exists {
		// 假设authorityId为"888"是超级管理员
		if id, ok := authorityId.(float64); ok && int(id) == 888 {
			return true
		}
	}
	return false
}

// GetFilteredQueryParams 获取应用权限约束后的查询参数
func GetFilteredQueryParams(c *gin.Context, originalParams map[string]interface{}) (map[string]interface{}, error) {
	permCheck := parsePermissionRequest(c)
	if permCheck == nil {
		return originalParams, nil
	}

	permService := &service.K8sPermissionService{}
	return permService.ApplyConstraints(permCheck, originalParams)
}
