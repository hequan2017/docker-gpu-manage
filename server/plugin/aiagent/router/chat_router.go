package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Chat = new(chat)

type chat struct{}

// Init 初始化 聊天 路由信息
func (r *chat) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("chat").Use(middleware.OperationRecord())
		group.POST("sendMessage", apiChat.SendMessage) // 发送消息
	}
}
