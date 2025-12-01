package computenode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComputeNodeRouter struct {}

// InitComputeNodeRouter 初始化 算力节点 路由信息
func (s *ComputeNodeRouter) InitComputeNodeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	computeNodeRouter := Router.Group("computeNode").Use(middleware.OperationRecord())
	computeNodeRouterWithoutRecord := Router.Group("computeNode")
	computeNodeRouterWithoutAuth := PublicRouter.Group("computeNode")
	{
		computeNodeRouter.POST("createComputeNode", computeNodeApi.CreateComputeNode)   // 新建算力节点
		computeNodeRouter.DELETE("deleteComputeNode", computeNodeApi.DeleteComputeNode) // 删除算力节点
		computeNodeRouter.DELETE("deleteComputeNodeByIds", computeNodeApi.DeleteComputeNodeByIds) // 批量删除算力节点
		computeNodeRouter.PUT("updateComputeNode", computeNodeApi.UpdateComputeNode)    // 更新算力节点
	}
	{
		computeNodeRouterWithoutRecord.GET("findComputeNode", computeNodeApi.FindComputeNode)        // 根据ID获取算力节点
		computeNodeRouterWithoutRecord.GET("getComputeNodeList", computeNodeApi.GetComputeNodeList)  // 获取算力节点列表
	}
	{
	    computeNodeRouterWithoutAuth.GET("getComputeNodePublic", computeNodeApi.GetComputeNodePublic)  // 算力节点开放接口
	}
}
