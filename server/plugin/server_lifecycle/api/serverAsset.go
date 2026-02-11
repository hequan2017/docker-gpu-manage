package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var ServerAsset = new(asset)

type asset struct {}

// CreateServerAsset 创建服务器资产
// @Tags ServerAsset
// @Summary 创建服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "创建服务器资产"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /asset/createServerAsset [post]
func (a *asset) CreateServerAsset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ServerAsset
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceServerAsset.CreateServerAsset(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteServerAsset 删除服务器资产
// @Tags ServerAsset
// @Summary 删除服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "删除服务器资产"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /asset/deleteServerAsset [delete]
func (a *asset) DeleteServerAsset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceServerAsset.DeleteServerAsset(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteServerAssetByIds 批量删除服务器资产
// @Tags ServerAsset
// @Summary 批量删除服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /asset/deleteServerAssetByIds [delete]
func (a *asset) DeleteServerAssetByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceServerAsset.DeleteServerAssetByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateServerAsset 更新服务器资产
// @Tags ServerAsset
// @Summary 更新服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "更新服务器资产"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /asset/updateServerAsset [put]
func (a *asset) UpdateServerAsset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ServerAsset
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceServerAsset.UpdateServerAsset(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindServerAsset 用id查询服务器资产
// @Tags ServerAsset
// @Summary 用id查询服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询服务器资产"
// @Success 200 {object} response.Response{data=model.ServerAsset,msg=string} "查询成功"
// @Router /asset/findServerAsset [get]
func (a *asset) FindServerAsset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reasset, err := serviceServerAsset.GetServerAsset(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reasset, c)
}
// GetServerAssetList 分页获取服务器资产列表
// @Tags ServerAsset
// @Summary 分页获取服务器资产列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ServerAssetSearch true "分页获取服务器资产列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /asset/getServerAssetList [get]
func (a *asset) GetServerAssetList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.ServerAssetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceServerAsset.GetServerAssetInfoList(ctx,pageInfo)
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
// GetServerAssetPublic 不需要鉴权的服务器资产接口
// @Tags ServerAsset
// @Summary 不需要鉴权的服务器资产接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /asset/getServerAssetPublic [get]
func (a *asset) GetServerAssetPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceServerAsset.GetServerAssetPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的服务器资产接口信息"}, "获取成功", c)
}
