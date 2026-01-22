package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Message = new(message)

type message struct{}

// GetMessageList 分页获取消息列表
// @Tags Message
// @Summary 分页获取消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.MessageSearch true "分页获取消息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /message/getMessageList [get]
func (a *message) GetMessageList(c *gin.Context) {
	var pageInfo request.MessageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMessage.GetMessageList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// DeleteMessage 删除消息
// @Tags Message
// @Summary 删除消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "消息ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /message/deleteMessage [delete]
func (a *message) DeleteMessage(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceMessage.DeleteMessage(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
