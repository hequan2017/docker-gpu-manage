package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Conversation = new(conversation)

type conversation struct{}

// Init 初始化 会话 路由信息
func (r *conversation) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("conversation").Use(middleware.OperationRecord())
		group.POST("createConversation", apiConversation.CreateConversation)     // 新建会话
		group.DELETE("deleteConversation", apiConversation.DeleteConversation) // 删除会话
		group.PUT("updateConversation", apiConversation.UpdateConversation)    // 更新会话
		group.POST("setActive", apiConversation.SetConversationActive)         // 设置激活状态
	}
	{
		group := private.Group("conversation")
		group.GET("findConversation", apiConversation.FindConversation)       // 根据ID获取会话
		group.GET("getConversationList", apiConversation.GetConversationList) // 获取会话列表
		group.GET("getActive", apiConversation.GetActiveConversation)         // 获取激活的会话
	}
}
