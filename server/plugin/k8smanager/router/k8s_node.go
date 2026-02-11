package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sNodeRouter struct{}

// InitK8sNodeRouter 初始化Node路由
func (r *K8sNodeRouter) InitK8sNodeRouter(Router *gin.RouterGroup) {
	nodeApi := k8sv1.ApiGroupApp.K8sNodeApi
	nodeRouter := Router.Group("node")
	{
		nodeRouter.GET("list", nodeApi.GetNodeList)   // 获取Node列表
		nodeRouter.GET("get", nodeApi.GetNode)        // 获取Node详情
		nodeRouter.POST("cordon", nodeApi.CordonNode) // 设置Node调度状态
	}
}
