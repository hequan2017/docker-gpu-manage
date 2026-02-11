package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MsSpace = new(msSpace)

type msSpace struct {}

// Init 初始化 创空间 路由信息
func (r *msSpace) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("space").Use(middleware.OperationRecord())
		group.POST("createMsSpace", apiMsSpace.CreateMsSpace)   // 新建创空间
		group.DELETE("deleteMsSpace", apiMsSpace.DeleteMsSpace) // 删除创空间
		group.DELETE("deleteMsSpaceByIds", apiMsSpace.DeleteMsSpaceByIds) // 批量删除创空间
		group.PUT("updateMsSpace", apiMsSpace.UpdateMsSpace)    // 更新创空间
	}
	{
	    group := private.Group("space")
		group.GET("findMsSpace", apiMsSpace.FindMsSpace)        // 根据ID获取创空间
		group.GET("getMsSpaceList", apiMsSpace.GetMsSpaceList)  // 获取创空间列表
	}
	{
	    group := public.Group("space")
	    group.GET("getMsSpacePublic", apiMsSpace.GetMsSpacePublic)  // 创空间开放接口
	}
}
