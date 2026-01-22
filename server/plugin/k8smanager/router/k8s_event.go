package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sEventRouter struct{}

// InitK8sEventRouter 初始化Event路由
func (r *K8sEventRouter) InitK8sEventRouter(Router *gin.RouterGroup) {
	eventApi := k8sv1.ApiGroupApp.K8sEventApi
	eventRouter := Router.Group("event")
	{
		eventRouter.POST("list", eventApi.GetEventList) // 获取Event列表
	}
}
