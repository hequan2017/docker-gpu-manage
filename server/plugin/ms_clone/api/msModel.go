package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var MsModel = new(msModel)

type msModel struct {}

// CreateMsModel 创建模型库
// @Tags MsModel
// @Summary 创建模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "创建模型库"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /model/createMsModel [post]
func (a *msModel) CreateMsModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsModel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsModel.CreateMsModel(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMsModel 删除模型库
// @Tags MsModel
// @Summary 删除模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "删除模型库"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /model/deleteMsModel [delete]
func (a *msModel) DeleteMsModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMsModel.DeleteMsModel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMsModelByIds 批量删除模型库
// @Tags MsModel
// @Summary 批量删除模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /model/deleteMsModelByIds [delete]
func (a *msModel) DeleteMsModelByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMsModel.DeleteMsModelByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMsModel 更新模型库
// @Tags MsModel
// @Summary 更新模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "更新模型库"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /model/updateMsModel [put]
func (a *msModel) UpdateMsModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsModel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsModel.UpdateMsModel(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMsModel 用id查询模型库
// @Tags MsModel
// @Summary 用id查询模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询模型库"
// @Success 200 {object} response.Response{data=model.MsModel,msg=string} "查询成功"
// @Router /model/findMsModel [get]
func (a *msModel) FindMsModel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remodel, err := serviceMsModel.GetMsModel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(remodel, c)
}
// GetMsModelList 分页获取模型库列表
// @Tags MsModel
// @Summary 分页获取模型库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MsModelSearch true "分页获取模型库列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /model/getMsModelList [get]
func (a *msModel) GetMsModelList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MsModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMsModel.GetMsModelInfoList(ctx,pageInfo)
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
// GetMsModelPublic 不需要鉴权的模型库接口
// @Tags MsModel
// @Summary 不需要鉴权的模型库接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /model/getMsModelPublic [get]
func (a *msModel) GetMsModelPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMsModel.GetMsModelPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的模型库接口信息"}, "获取成功", c)
}
