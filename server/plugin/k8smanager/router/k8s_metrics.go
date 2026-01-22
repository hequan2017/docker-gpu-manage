package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sMetricsRouter struct{}

// InitK8sMetricsRouter 初始化监控指标路由
func (r *K8sMetricsRouter) InitK8sMetricsRouter(Router *gin.RouterGroup) {
	metricsApi := k8sv1.ApiGroupApp.K8sMetricsApi
	metricsRouter := Router.Group("metrics")
	{
		metricsRouter.GET("cluster", metricsApi.GetClusterMetrics)         // 获取集群指标
		metricsRouter.POST("cluster/refresh", metricsApi.RefreshClusterMetrics) // 刷新集群指标
		metricsRouter.GET("nodes", metricsApi.GetNodeMetrics)             // 获取节点指标
		metricsRouter.GET("pods", metricsApi.GetPodMetrics)               // 获取Pod指标
		metricsRouter.GET("summary", metricsApi.GetMetricsSummary)        // 获取指标摘要
		metricsRouter.POST("collector/start", metricsApi.StartAutoCollector) // 启动自动收集
		metricsRouter.POST("collector/stop", metricsApi.StopAutoCollector)   // 停止自动收集
	}
}
