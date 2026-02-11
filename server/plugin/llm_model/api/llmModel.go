package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var LlmModel = new(llm)

type llm struct {}

// CreateLlmModel 创建开源大模型
// @Tags LlmModel
// @Summary 创建开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "创建开源大模型"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /llm/createLlmModel [post]
func (a *llm) CreateLlmModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.LlmModel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceLlmModel.CreateLlmModel(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteLlmModel 删除开源大模型
// @Tags LlmModel
// @Summary 删除开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "删除开源大模型"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /llm/deleteLlmModel [delete]
func (a *llm) DeleteLlmModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceLlmModel.DeleteLlmModel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteLlmModelByIds 批量删除开源大模型
// @Tags LlmModel
// @Summary 批量删除开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /llm/deleteLlmModelByIds [delete]
func (a *llm) DeleteLlmModelByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceLlmModel.DeleteLlmModelByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateLlmModel 更新开源大模型
// @Tags LlmModel
// @Summary 更新开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "更新开源大模型"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /llm/updateLlmModel [put]
func (a *llm) UpdateLlmModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.LlmModel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceLlmModel.UpdateLlmModel(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindLlmModel 用id查询开源大模型
// @Tags LlmModel
// @Summary 用id查询开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询开源大模型"
// @Success 200 {object} response.Response{data=model.LlmModel,msg=string} "查询成功"
// @Router /llm/findLlmModel [get]
func (a *llm) FindLlmModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rellm, err := serviceLlmModel.GetLlmModel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(rellm, c)
}
// GetLlmModelList 分页获取开源大模型列表
// @Tags LlmModel
// @Summary 分页获取开源大模型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.LlmModelSearch true "分页获取开源大模型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /llm/getLlmModelList [get]
func (a *llm) GetLlmModelList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.LlmModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceLlmModel.GetLlmModelInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetLlmModelPublic 不需要鉴权的开源大模型接口
// @Tags LlmModel
// @Summary 不需要鉴权的开源大模型接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /llm/getLlmModelPublic [get]
func (a *llm) GetLlmModelPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceLlmModel.GetLlmModelPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的开源大模型接口信息"}, "获取成功", c)
}
