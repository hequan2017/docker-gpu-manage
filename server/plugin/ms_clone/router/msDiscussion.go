package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MsDiscussion = new(msDiscussion)

type msDiscussion struct {}

// Init 初始化 社区讨论 路由信息
func (r *msDiscussion) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("discussion").Use(middleware.OperationRecord())
		group.POST("createMsDiscussion", apiMsDiscussion.CreateMsDiscussion)   // 新建社区讨论
		group.DELETE("deleteMsDiscussion", apiMsDiscussion.DeleteMsDiscussion) // 删除社区讨论
		group.DELETE("deleteMsDiscussionByIds", apiMsDiscussion.DeleteMsDiscussionByIds) // 批量删除社区讨论
		group.PUT("updateMsDiscussion", apiMsDiscussion.UpdateMsDiscussion)    // 更新社区讨论
		group.POST("likeMsDiscussion", apiMsDiscussion.LikeMsDiscussion)       // 点赞社区讨论
	}
	{
		group := private.Group("discussion")
		group.GET("findMsDiscussion", apiMsDiscussion.FindMsDiscussion)        // 根据ID获取社区讨论
		group.GET("getMsDiscussionList", apiMsDiscussion.GetMsDiscussionList)  // 获取社区讨论列表
	}
	{
	    group := public.Group("discussion")
	    group.GET("getMsDiscussionDataSource", apiMsDiscussion.GetMsDiscussionDataSource)  // 获取社区讨论数据源
		group.GET("getMsDiscussionPublic", apiMsDiscussion.GetMsDiscussionPublic)  // 社区讨论开放接口
		group.POST("viewMsDiscussion", apiMsDiscussion.ViewMsDiscussion)       // 增加浏览量(公开接口)
	}
}
