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

type PcdnPolicyApi struct{}

func (a *PcdnPolicyApi) CreatePcdnPolicy(c *gin.Context) {
	var req pcdn.PcdnPolicy
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnPolicyService.CreatePcdnPolicy(c.Request.Context(), &req); err != nil {
		global.GVA_LOG.Error("创建PCDN策略失败", zap.Error(err))
		response.FailWithMessage("创建PCDN策略失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建PCDN策略成功", c)
}

func (a *PcdnPolicyApi) DeletePcdnPolicy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("id参数错误", c)
		return
	}
	if err = pcdnPolicyService.DeletePcdnPolicy(c.Request.Context(), uint(id)); err != nil {
		global.GVA_LOG.Error("删除PCDN策略失败", zap.Error(err))
		response.FailWithMessage("删除PCDN策略失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除PCDN策略成功", c)
}

func (a *PcdnPolicyApi) UpdatePcdnPolicy(c *gin.Context) {
	var req pcdn.PcdnPolicy
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnPolicyService.UpdatePcdnPolicy(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("更新PCDN策略失败", zap.Error(err))
		response.FailWithMessage("更新PCDN策略失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新PCDN策略成功", c)
}

func (a *PcdnPolicyApi) FindPcdnPolicy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("id参数错误", c)
		return
	}
	ret, err := pcdnPolicyService.GetPcdnPolicy(c.Request.Context(), uint(id))
	if err != nil {
		global.GVA_LOG.Error("查询PCDN策略失败", zap.Error(err))
		response.FailWithMessage("查询PCDN策略失败:"+err.Error(), c)
		return
	}
	response.OkWithData(ret, c)
}

func (a *PcdnPolicyApi) GetPcdnPolicyList(c *gin.Context) {
	var req pcdnReq.PcdnPolicySearch
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcdnPolicyService.GetPcdnPolicyList(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("获取PCDN策略列表失败", zap.Error(err))
		response.FailWithMessage("获取PCDN策略列表失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: req.Page, PageSize: req.PageSize}, "获取PCDN策略列表成功", c)
}

func (a *PcdnPolicyApi) GrayReleasePcdnPolicy(c *gin.Context) {
	var req pcdnReq.PcdnPolicyGrayReleaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnPolicyService.GrayRelease(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("PCDN策略灰度发布失败", zap.Error(err))
		response.FailWithMessage("PCDN策略灰度发布失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("PCDN策略灰度发布成功", c)
}

func (a *PcdnPolicyApi) SwitchPcdnPolicy(c *gin.Context) {
	var req pcdnReq.PcdnPolicySwitchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pcdnPolicyService.SwitchPolicy(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("PCDN策略启停失败", zap.Error(err))
		response.FailWithMessage("PCDN策略启停失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("PCDN策略启停成功", c)
}
