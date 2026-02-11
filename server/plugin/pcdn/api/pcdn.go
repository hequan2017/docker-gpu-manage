package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PcdnNodeApi struct{}

// CreatePcdnNode 创建PCDN节点
// @Tags PcdnNode
// @Summary 创建PCDN节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcdnNode true "创建PCDN节点"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pcdn/createPcdnNode [post]
func (a *PcdnNodeApi) CreatePcdnNode(c *gin.Context) {
	var pcdnNode model.PcdnNode
	err := c.ShouldBindJSON(&pcdnNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.Service.PcdnNodeService.CreatePcdnNode(&pcdnNode)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePcdnNode 删除PCDN节点
// @Tags PcdnNode
// @Summary 删除PCDN节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcdnNode true "删除PCDN节点"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pcdn/deletePcdnNode [delete]
func (a *PcdnNodeApi) DeletePcdnNode(c *gin.Context) {
	ID := c.Query("ID")
	err := service.Service.PcdnNodeService.DeletePcdnNode(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePcdnNodeByIds 批量删除PCDN节点
// @Tags PcdnNode
// @Summary 批量删除PCDN节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PCDN节点"
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pcdn/deletePcdnNodeByIds [delete]
func (a *PcdnNodeApi) DeletePcdnNodeByIds(c *gin.Context) {
	IDs := c.QueryArray("ids[]")
	err := service.Service.PcdnNodeService.DeletePcdnNodeByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePcdnNode 更新PCDN节点
// @Tags PcdnNode
// @Summary 更新PCDN节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PcdnNode true "更新PCDN节点"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pcdn/updatePcdnNode [put]
func (a *PcdnNodeApi) UpdatePcdnNode(c *gin.Context) {
	var pcdnNode model.PcdnNode
	err := c.ShouldBindJSON(&pcdnNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.Service.PcdnNodeService.UpdatePcdnNode(pcdnNode)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPcdnNode 用id查询PCDN节点
// @Tags PcdnNode
// @Summary 用id查询PCDN节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PcdnNode true "用id查询PCDN节点"
// @Success 200 {object} response.Response{data=model.PcdnNode,msg=string} "查询成功"
// @Router /pcdn/findPcdnNode [get]
func (a *PcdnNodeApi) FindPcdnNode(c *gin.Context) {
	ID := c.Query("ID")
	pcdnNode, err := service.Service.PcdnNodeService.GetPcdnNode(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(pcdnNode, c)
}

// GetPcdnNodeList 分页获取PCDN节点列表
// @Tags PcdnNode
// @Summary 分页获取PCDN节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PcdnNodeSearch true "分页获取PCDN节点列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pcdn/getPcdnNodeList [get]
func (a *PcdnNodeApi) GetPcdnNodeList(c *gin.Context) {
	var pageInfo request.PcdnNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.Service.PcdnNodeService.GetPcdnNodeInfoList(pageInfo)
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
