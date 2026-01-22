package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sAuditRouter struct{}

// InitK8sAuditRouter 初始化审计日志路由
func (r *K8sAuditRouter) InitK8sAuditRouter(Router *gin.RouterGroup) {
	auditApi := k8sv1.ApiGroupApp.K8sAuditApi
	auditRouter := Router.Group("audit")
	{
		auditRouter.GET("list", auditApi.GetAuditLogs)           // 获取审计日志列表
		auditRouter.GET("stats", auditApi.GetAuditLogStats)      // 获取审计统计
		auditRouter.GET("client-stats", auditApi.GetClientStats) // 获取客户端统计
		auditRouter.DELETE("cleanup", auditApi.DeleteOldLogs)    // 清理旧日志
		auditRouter.GET("export", auditApi.ExportAuditLogs)      // 导出审计日志
	}
}
