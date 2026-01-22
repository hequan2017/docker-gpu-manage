package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Chat = new(chat)

type chat struct{}

// SendMessage 发送消息并获取AI回复
// @Tags Chat
// @Summary 发送消息并获取AI回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ChatRequest true "对话请求"
// @Success 200 {object} response.Response{data=request.ChatResponse,msg=string} "发送成功"
// @Router /chat/sendMessage [post]
func (a *chat) SendMessage(c *gin.Context) {
	var chatReq request.ChatRequest
	err := c.ShouldBindJSON(&chatReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := int(c.GetUint("user_id"))
	resp, err := serviceChat.SendMessage(chatReq, userID)
	if err != nil {
		global.GVA_LOG.Error("发送消息失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(resp, c)
}
