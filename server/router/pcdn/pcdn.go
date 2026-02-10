package pcdn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PcdnRouter struct{}

func (r *PcdnRouter) InitPcdnRouter(Router *gin.RouterGroup) {
	pcdnRouter := Router.Group("pcdn").Use(middleware.OperationRecord())
	pcdnRouterWithoutRecord := Router.Group("pcdn")

	{
		pcdnRouter.POST("node/create", pcdnApi.CreatePcdnNode)
		pcdnRouter.DELETE("node/delete", pcdnApi.DeletePcdnNode)
		pcdnRouter.PUT("node/update", pcdnApi.UpdatePcdnNode)
		pcdnRouter.PUT("node/online", pcdnApi.SwitchPcdnNodeOnline)
		pcdnRouter.PUT("node/weight", pcdnApi.UpdatePcdnNodeWeight)

		pcdnRouter.POST("policy/create", pcdnApi.CreatePcdnPolicy)
		pcdnRouter.DELETE("policy/delete", pcdnApi.DeletePcdnPolicy)
		pcdnRouter.PUT("policy/update", pcdnApi.UpdatePcdnPolicy)
		pcdnRouter.PUT("policy/grayRelease", pcdnApi.GrayReleasePcdnPolicy)
		pcdnRouter.PUT("policy/switch", pcdnApi.SwitchPcdnPolicy)
	}

	{
		pcdnRouterWithoutRecord.GET("node/find", pcdnApi.FindPcdnNode)
		pcdnRouterWithoutRecord.GET("node/list", pcdnApi.GetPcdnNodeList)
		pcdnRouterWithoutRecord.GET("policy/find", pcdnApi.FindPcdnPolicy)
		pcdnRouterWithoutRecord.GET("policy/list", pcdnApi.GetPcdnPolicyList)
		pcdnRouterWithoutRecord.GET("task/list", pcdnApi.GetPcdnDispatchTaskList)
	}
}
