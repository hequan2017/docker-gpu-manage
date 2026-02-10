package pcdn

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PcdnNodeApi struct{}

func (a *PcdnNodeApi) CreatePcdnNode(c *gin.Context) {
	var req pcdn.PcdnNode
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnNodeService.CreatePcdnNode(c.Request.Context(), &req); err != nil {
		global.GVA_LOG.Error("创建PCDN节点失败", zap.Error(err))
		response.FailWithMessage("创建PCDN节点失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建PCDN节点成功", c)
}

func (a *PcdnNodeApi) DeletePcdnNode(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("id参数错误", c)
		return
	}
	if err = pcdnNodeService.DeletePcdnNode(c.Request.Context(), uint(id)); err != nil {
		global.GVA_LOG.Error("删除PCDN节点失败", zap.Error(err))
		response.FailWithMessage("删除PCDN节点失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除PCDN节点成功", c)
}

func (a *PcdnNodeApi) UpdatePcdnNode(c *gin.Context) {
	var req pcdn.PcdnNode
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnNodeService.UpdatePcdnNode(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("更新PCDN节点失败", zap.Error(err))
		response.FailWithMessage("更新PCDN节点失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新PCDN节点成功", c)
}

func (a *PcdnNodeApi) FindPcdnNode(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("id参数错误", c)
		return
	}
	ret, err := pcdnNodeService.GetPcdnNode(c.Request.Context(), uint(id))
	if err != nil {
		global.GVA_LOG.Error("查询PCDN节点失败", zap.Error(err))
		response.FailWithMessage("查询PCDN节点失败:"+err.Error(), c)
		return
	}
	response.OkWithData(ret, c)
}

func (a *PcdnNodeApi) GetPcdnNodeList(c *gin.Context) {
	var req pcdnReq.PcdnNodeSearch
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcdnNodeService.GetPcdnNodeList(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("获取PCDN节点列表失败", zap.Error(err))
		response.FailWithMessage("获取PCDN节点列表失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: req.Page, PageSize: req.PageSize}, "获取PCDN节点列表成功", c)
}

func (a *PcdnNodeApi) SwitchPcdnNodeOnline(c *gin.Context) {
	var req pcdnReq.PcdnNodeOnlineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnNodeService.SwitchOnline(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("切换PCDN节点上下线失败", zap.Error(err))
		response.FailWithMessage("切换PCDN节点上下线失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("切换PCDN节点上下线成功", c)
}

func (a *PcdnNodeApi) UpdatePcdnNodeWeight(c *gin.Context) {
	var req pcdnReq.PcdnNodeWeightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnNodeService.UpdateWeight(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("调整PCDN节点权重失败", zap.Error(err))
		response.FailWithMessage("调整PCDN节点权重失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("调整PCDN节点权重成功", c)
}
