package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sPodRouter struct{}

// InitK8sPodRouter 初始化Pod路由
func (r *K8sPodRouter) InitK8sPodRouter(Router *gin.RouterGroup) {
	podApi := k8sv1.ApiGroupApp.K8sPodApi
	podRouter := Router.Group("pod")
	{
		podRouter.GET("list", podApi.GetPodList)           // 获取Pod列表
		podRouter.GET("get", podApi.GetPod)                // 获取Pod详情
		podRouter.DELETE("delete", podApi.DeletePod)       // 删除Pod
		podRouter.POST("log", podApi.GetPodLog)            // 获取Pod日志
		podRouter.GET("containers", podApi.GetPodContainers) // 获取Pod容器列表
		podRouter.GET("events", podApi.GetPodEvents)       // 获取Pod事件
	}
}
