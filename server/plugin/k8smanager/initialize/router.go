package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	k8smiddleware "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	// 公开路由（不需要认证）
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("k8s")
	// 私有路由（需要认证和权限）
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("k8s")
	private.Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(k8smiddleware.K8sAuditMiddleware()) // 添加审计中间件

	// 初始化各模块路由
	router.RouterGroupApp.K8sClusterRouter.InitK8sClusterRouter(public, private)
	router.RouterGroupApp.K8sPodRouter.InitK8sPodRouter(private)
	router.RouterGroupApp.K8sDeploymentRouter.InitK8sDeploymentRouter(private)
	router.RouterGroupApp.K8sServiceRouter.InitK8sServiceRouter(private)
	router.RouterGroupApp.K8sNamespaceRouter.InitK8sNamespaceRouter(private)
	router.RouterGroupApp.K8sEventRouter.InitK8sEventRouter(private)
	router.RouterGroupApp.K8sAuditRouter.InitK8sAuditRouter(private)     // 审计日志路由
	router.RouterGroupApp.K8sMetricsRouter.InitK8sMetricsRouter(private) // 监控指标路由
}
