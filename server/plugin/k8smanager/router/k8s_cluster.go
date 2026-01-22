package router

import (
	k8sv1 "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sClusterRouter struct{}

// InitK8sClusterRouter 初始化K8s集群路由
func (r *K8sClusterRouter) InitK8sClusterRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	clusterApi := k8sv1.ApiGroupApp.K8sClusterApi
	clusterRouter := Router.Group("cluster")
	{
		clusterRouter.POST("create", clusterApi.CreateK8sCluster)      // 创建集群
		clusterRouter.DELETE("delete", clusterApi.DeleteK8sCluster)   // 删除集群
		clusterRouter.DELETE("deleteByIds", clusterApi.DeleteK8sClusterByIds) // 批量删除集群
		clusterRouter.PUT("update", clusterApi.UpdateK8sCluster)      // 更新集群
		clusterRouter.GET("get", clusterApi.GetK8sCluster)            // 获取集群详情
		clusterRouter.GET("list", clusterApi.GetK8sClusterList)       // 获取集群列表
		clusterRouter.POST("refresh", clusterApi.RefreshClusterStatus) // 刷新集群状态
		clusterRouter.GET("all", clusterApi.GetAllClusters)           // 获取所有集群
	}
}
