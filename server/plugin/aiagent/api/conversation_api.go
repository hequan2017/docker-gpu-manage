package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Conversation = new(conversation)

type conversation struct{}

// CreateConversation 创建会话
// @Tags Conversation
// @Summary 创建会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "创建会话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /conversation/createConversation [post]
func (a *conversation) CreateConversation(c *gin.Context) {
	var conversation model.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 从上下文获取用户ID（uint -> int 转换）
	userIDUint := c.GetUint("user_id")
	userID := int(userIDUint)
	conversation.UserID = &userID
	err = serviceConversation.CreateConversation(&conversation)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithDetailed(conversation, "创建成功", c)
}

// DeleteConversation 删除会话
// @Tags Conversation
// @Summary 删除会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "删除会话"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /conversation/deleteConversation [delete]
func (a *conversation) DeleteConversation(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceConversation.DeleteConversation(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateConversation 更新会话
// @Tags Conversation
// @Summary 更新会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "更新会话"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /conversation/updateConversation [put]
func (a *conversation) UpdateConversation(c *gin.Context) {
	var conversation model.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceConversation.UpdateConversation(conversation)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindConversation 用id查询会话
// @Tags Conversation
// @Summary 用id查询会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Conversation true "用id查询会话"
// @Success 200 {object} response.Response{data=model.Conversation,msg=string} "查询成功"
// @Router /conversation/findConversation [get]
func (a *conversation) FindConversation(c *gin.Context) {
	ID := c.Query("ID")
	conversation, err := serviceConversation.GetConversation(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(conversation, c)
}

// GetConversationList 分页获取会话列表
// @Tags Conversation
// @Summary 分页获取会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.ConversationSearch true "分页获取会话列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /conversation/getConversationList [get]
func (a *conversation) GetConversationList(c *gin.Context) {
	var pageInfo request.ConversationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceConversation.GetConversationList(pageInfo)
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

// SetConversationActive 设置会话激活状态
// @Tags Conversation
// @Summary 设置会话激活状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "会话ID"
// @Param isActive query bool true "是否激活"
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /conversation/setActive [post]
func (a *conversation) SetConversationActive(c *gin.Context) {
	ID := c.Query("ID")
	isActive := c.DefaultQuery("isActive", "true") == "true"
	userID := int(c.GetUint("user_id"))
	err := serviceConversation.SetConversationActive(ID, userID, isActive)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetActiveConversation 获取用户的激活会话
// @Tags Conversation
// @Summary 获取用户的激活会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=model.Conversation,msg=string} "获取成功"
// @Router /conversation/getActive [get]
func (a *conversation) GetActiveConversation(c *gin.Context) {
	userID := int(c.GetUint("user_id"))
	conversation, err := serviceConversation.GetActiveConversation(userID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(conversation, c)
}
