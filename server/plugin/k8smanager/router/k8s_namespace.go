package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sNamespaceRouter struct{}

// InitK8sNamespaceRouter 初始化Namespace路由
func (r *K8sNamespaceRouter) InitK8sNamespaceRouter(Router *gin.RouterGroup) {
	namespaceApi := k8sv1.ApiGroupApp.K8sNamespaceApi
	namespaceRouter := Router.Group("namespace")
	{
		namespaceRouter.GET("list", namespaceApi.GetNamespaceList)   // 获取Namespace列表
		namespaceRouter.GET("get", namespaceApi.GetNamespace)       // 获取Namespace详情
		namespaceRouter.POST("create", namespaceApi.CreateNamespace) // 创建Namespace
		namespaceRouter.DELETE("delete", namespaceApi.DeleteNamespace) // 删除Namespace
	}
}
