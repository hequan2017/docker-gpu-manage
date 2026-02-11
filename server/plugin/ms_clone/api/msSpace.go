package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var MsSpace = new(msSpace)

type msSpace struct {}

// CreateMsSpace 创建创空间
// @Tags MsSpace
// @Summary 创建创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "创建创空间"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /space/createMsSpace [post]
func (a *msSpace) CreateMsSpace(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsSpace
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsSpace.CreateMsSpace(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMsSpace 删除创空间
// @Tags MsSpace
// @Summary 删除创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "删除创空间"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /space/deleteMsSpace [delete]
func (a *msSpace) DeleteMsSpace(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMsSpace.DeleteMsSpace(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMsSpaceByIds 批量删除创空间
// @Tags MsSpace
// @Summary 批量删除创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /space/deleteMsSpaceByIds [delete]
func (a *msSpace) DeleteMsSpaceByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMsSpace.DeleteMsSpaceByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMsSpace 更新创空间
// @Tags MsSpace
// @Summary 更新创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "更新创空间"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /space/updateMsSpace [put]
func (a *msSpace) UpdateMsSpace(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsSpace
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsSpace.UpdateMsSpace(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMsSpace 用id查询创空间
// @Tags MsSpace
// @Summary 用id查询创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询创空间"
// @Success 200 {object} response.Response{data=model.MsSpace,msg=string} "查询成功"
// @Router /space/findMsSpace [get]
func (a *msSpace) FindMsSpace(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	respace, err := serviceMsSpace.GetMsSpace(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(respace, c)
}
// GetMsSpaceList 分页获取创空间列表
// @Tags MsSpace
// @Summary 分页获取创空间列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MsSpaceSearch true "分页获取创空间列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /space/getMsSpaceList [get]
func (a *msSpace) GetMsSpaceList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MsSpaceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMsSpace.GetMsSpaceInfoList(ctx,pageInfo)
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
// GetMsSpacePublic 不需要鉴权的创空间接口
// @Tags MsSpace
// @Summary 不需要鉴权的创空间接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /space/getMsSpacePublic [get]
func (a *msSpace) GetMsSpacePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMsSpace.GetMsSpacePublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的创空间接口信息"}, "获取成功", c)
}
