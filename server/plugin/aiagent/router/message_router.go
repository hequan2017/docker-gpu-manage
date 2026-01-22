package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Message = new(message)

type message struct{}

// Init 初始化 消息 路由信息
func (r *message) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("message").Use(middleware.OperationRecord())
		group.DELETE("deleteMessage", apiMessage.DeleteMessage) // 删除消息
	}
	{
		group := private.Group("message")
		group.GET("getMessageList", apiMessage.GetMessageList) // 获取消息列表
	}
}
