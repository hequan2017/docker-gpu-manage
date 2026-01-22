package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DellAssetApi struct{}

// CreateDellAsset 创建戴尔服务器资产
// @Tags DellAsset
// @Summary 创建戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DellAsset true "服务器资产信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dellAsset/createDellAsset [post]
func (daa *DellAssetApi) CreateDellAsset(c *gin.Context) {
	var dellAsset model.DellAsset
	err := c.ShouldBindJSON(&dellAsset)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceDellAsset.CreateDellAsset(&dellAsset)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDellAsset 删除戴尔服务器资产
// @Tags DellAsset
// @Summary 删除戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DellAsset true "ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dellAsset/deleteDellAsset [delete]
func (daa *DellAssetApi) DeleteDellAsset(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceDellAsset.DeleteDellAsset(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDellAssetByIds 批量删除戴尔服务器资产
// @Tags DellAsset
// @Summary 批量删除戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dellAsset/deleteDellAssetByIds [delete]
func (daa *DellAssetApi) DeleteDellAssetByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := serviceDellAsset.DeleteDellAssetByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDellAsset 更新戴尔服务器资产
// @Tags DellAsset
// @Summary 更新戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DellAsset true "服务器资产信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dellAsset/updateDellAsset [put]
func (daa *DellAssetApi) UpdateDellAsset(c *gin.Context) {
	var dellAsset model.DellAsset
	err := c.ShouldBindJSON(&dellAsset)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceDellAsset.UpdateDellAsset(dellAsset)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDellAsset 用id查询戴尔服务器资产
// @Tags DellAsset
// @Summary 用id查询戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DellAsset true "ID"
// @Success 200 {object} response.Response{data=model.DellAsset,msg=string} "查询成功"
// @Router /dellAsset/findDellAsset [get]
func (daa *DellAssetApi) FindDellAsset(c *gin.Context) {
	ID := c.Query("ID")
	reDellAsset, err := serviceDellAsset.GetDellAsset(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reDellAsset, c)
}

// GetDellAssetList 分页获取戴尔服务器资产列表
// @Tags DellAsset
// @Summary 分页获取戴尔服务器资产列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DellAssetSearch true "分页获取戴尔服务器资产列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dellAsset/getDellAssetList [get]
func (daa *DellAssetApi) GetDellAssetList(c *gin.Context) {
	var pageInfo request.DellAssetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceDellAsset.GetDellAssetList(pageInfo)
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

// GetDellAssetStatistics 获取资产统计信息
// @Tags DellAsset
// @Summary 获取资产统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /dellAsset/getStatistics [get]
func (daa *DellAssetApi) GetDellAssetStatistics(c *gin.Context) {
	stats, err := serviceDellAsset.GetDellAssetStatistics()
	if err != nil {
		global.GVA_LOG.Error("获取统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取统计信息失败", c)
		return
	}
	response.OkWithDetailed(stats, "获取成功", c)
}
