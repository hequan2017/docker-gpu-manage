package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/api"
	"github.com/gin-gonic/gin"
)

type PcdnNodeRouter struct{}

func (s *PcdnNodeRouter) InitPcdnNodeRouter(Router *gin.RouterGroup) {
	pcdnNodeRouter := Router.Group("pcdn").Use(middleware.OperationRecord())
	pcdnNodeRouterWithoutRecord := Router.Group("pcdn")
	pcdnNodeApi := api.Api.PcdnNodeApi
	{
		pcdnNodeRouter.POST("createPcdnNode", pcdnNodeApi.CreatePcdnNode)             // 新建PCDN节点
		pcdnNodeRouter.DELETE("deletePcdnNode", pcdnNodeApi.DeletePcdnNode)           // 删除PCDN节点
		pcdnNodeRouter.DELETE("deletePcdnNodeByIds", pcdnNodeApi.DeletePcdnNodeByIds) // 批量删除PCDN节点
		pcdnNodeRouter.PUT("updatePcdnNode", pcdnNodeApi.UpdatePcdnNode)              // 更新PCDN节点
	}
	{
		pcdnNodeRouterWithoutRecord.GET("findPcdnNode", pcdnNodeApi.FindPcdnNode)       // 根据ID获取PCDN节点
		pcdnNodeRouterWithoutRecord.GET("getPcdnNodeList", pcdnNodeApi.GetPcdnNodeList) // 获取PCDN节点列表
	}
}
