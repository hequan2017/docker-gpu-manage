package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	commonRequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sClusterApi struct{}

// CreateK8sCluster 创建K8s集群
// @Tags K8sCluster
// @Summary 创建K8s集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "集群信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8s/cluster/create [post]
func (a *K8sClusterApi) CreateK8sCluster(c *gin.Context) {
	var cluster model.K8sCluster
	err := c.ShouldBindJSON(&cluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sClusterService.CreateK8sCluster(c.Request.Context(), &cluster)
	if err != nil {
		global.GVA_LOG.Error("创建K8s集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteK8sCluster 删除K8s集群
// @Tags K8sCluster
// @Summary 删除K8s集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "集群ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/cluster/delete [delete]
func (a *K8sClusterApi) DeleteK8sCluster(c *gin.Context) {
	var cluster model.K8sCluster
	err := c.ShouldBindJSON(&cluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sClusterService.DeleteK8sCluster(c.Request.Context(), cluster.ID)
	if err != nil {
		global.GVA_LOG.Error("删除K8s集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// DeleteK8sClusterByIds 批量删除K8s集群
// @Tags K8sCluster
// @Summary 批量删除K8s集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "集群ID列表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/cluster/deleteByIds [delete]
func (a *K8sClusterApi) DeleteK8sClusterByIds(c *gin.Context) {
	var req commonRequest.IdsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 转换 []int 为 []uint
	ids := make([]uint, len(req.Ids))
	for i, id := range req.Ids {
		ids[i] = uint(id)
	}
	err = service.ServiceGroupApp.K8sClusterService.DeleteK8sClusterByIds(c.Request.Context(), ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除K8s集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sCluster 更新K8s集群
// @Tags K8sCluster
// @Summary 更新K8s集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "集群信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8s/cluster/update [put]
func (a *K8sClusterApi) UpdateK8sCluster(c *gin.Context) {
	var cluster model.K8sCluster
	err := c.ShouldBindJSON(&cluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sClusterService.UpdateK8sCluster(c.Request.Context(), &cluster)
	if err != nil {
		global.GVA_LOG.Error("更新K8s集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// GetK8sCluster 根据ID获取K8s集群
// @Tags K8sCluster
// @Summary 根据ID获取K8s集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sCluster true "集群ID"
// @Success 200 {object} response.Response{data=model.K8sCluster,msg=string} "获取成功"
// @Router /k8s/cluster/get [get]
func (a *K8sClusterApi) GetK8sCluster(c *gin.Context) {
	var req commonRequest.GetById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	cluster, err := service.ServiceGroupApp.K8sClusterService.GetK8sCluster(c.Request.Context(), uint(req.ID))
	if err != nil {
		global.GVA_LOG.Error("获取K8s集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(cluster, "获取成功", c)
}

// GetK8sClusterList 获取K8s集群列表
// @Tags K8sCluster
// @Summary 获取K8s集群列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.K8sClusterSearch true "查询条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8s/cluster/list [get]
func (a *K8sClusterApi) GetK8sClusterList(c *gin.Context) {
	var info request.K8sClusterSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := service.ServiceGroupApp.K8sClusterService.GetK8sClusterInfoList(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取K8s集群列表失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

// RefreshClusterStatus 刷新集群状态
// @Tags K8sCluster
// @Summary 刷新集群状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Success 200 {object} response.Response{msg=string} "刷新成功"
// @Router /k8s/cluster/refresh [post]
func (a *K8sClusterApi) RefreshClusterStatus(c *gin.Context) {
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	err := service.ServiceGroupApp.K8sClusterService.RefreshClusterStatus(c.Request.Context(), clusterName)
	if err != nil {
		global.GVA_LOG.Error("刷新集群状态失败!", zap.String("error", err.Error()))
		response.FailWithMessage("刷新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("刷新成功", c)
}

// GetAllClusters 获取所有集群（用于下拉选择）
// @Tags K8sCluster
// @Summary 获取所有集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]model.K8sCluster,msg=string} "获取成功"
// @Router /k8s/cluster/all [get]
func (a *K8sClusterApi) GetAllClusters(c *gin.Context) {
	clusters, err := service.ServiceGroupApp.K8sClusterService.GetAllClusters(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error("获取所有集群失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(clusters, "获取成功", c)
}
