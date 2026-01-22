package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sDeploymentRouter struct{}

// InitK8sDeploymentRouter 初始化Deployment路由
func (r *K8sDeploymentRouter) InitK8sDeploymentRouter(Router *gin.RouterGroup) {
	deploymentApi := k8sv1.ApiGroupApp.K8sDeploymentApi
	deploymentRouter := Router.Group("deployment")
	{
		deploymentRouter.GET("list", deploymentApi.GetDeploymentList)   // 获取Deployment列表
		deploymentRouter.GET("get", deploymentApi.GetDeployment)       // 获取Deployment详情
		deploymentRouter.POST("scale", deploymentApi.ScaleDeployment)  // 扩缩容Deployment
		deploymentRouter.POST("restart", deploymentApi.RestartDeployment) // 重启Deployment
		deploymentRouter.DELETE("delete", deploymentApi.DeleteDeployment) // 删除Deployment
		deploymentRouter.GET("pods", deploymentApi.GetDeploymentPods)  // 获取Deployment关联的Pods
	}
}
