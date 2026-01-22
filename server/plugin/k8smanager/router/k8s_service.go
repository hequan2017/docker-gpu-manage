package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sServiceRouter struct{}

// InitK8sServiceRouter 初始化Service路由
func (r *K8sServiceRouter) InitK8sServiceRouter(Router *gin.RouterGroup) {
	serviceApi := k8sv1.ApiGroupApp.K8sServiceApi
	serviceRouter := Router.Group("service")
	{
		serviceRouter.GET("list", serviceApi.GetServiceList)         // 获取Service列表
		serviceRouter.GET("get", serviceApi.GetService)             // 获取Service详情
		serviceRouter.DELETE("delete", serviceApi.DeleteService)    // 删除Service
		serviceRouter.GET("endpoints", serviceApi.GetServiceEndpoints) // 获取Service的Endpoints
	}
}
