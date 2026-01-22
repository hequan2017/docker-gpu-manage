package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sMetricsApi struct{}

// GetClusterMetrics 获取集群指标
// @Tags K8sMetrics
// @Summary 获取集群指标
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Success 200 {object} response.Response{data=service.ClusterMetrics,msg=string} "获取成功"
// @Router /k8s/metrics/cluster [get]
func (a *K8sMetricsApi) GetClusterMetrics(c *gin.Context) {
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	metrics, err := service.ServiceGroupApp.K8sMetricsService.GetClusterMetrics(clusterName)
	if err != nil {
		global.GVA_LOG.Error("获取集群指标失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(metrics, "获取成功", c)
}

// RefreshClusterMetrics 刷新集群指标
// @Tags K8sMetrics
// @Summary 刷新集群指标
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Success 200 {object} response.Response{data=service.ClusterMetrics,msg=string} "刷新成功"
// @Router /k8s/metrics/cluster/refresh [post]
func (a *K8sMetricsApi) RefreshClusterMetrics(c *gin.Context) {
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	metrics, err := service.ServiceGroupApp.K8sMetricsService.CollectClusterMetrics(c.Request.Context(), clusterName)
	if err != nil {
		global.GVA_LOG.Error("刷新集群指标失败!", zap.String("error", err.Error()))
		response.FailWithMessage("刷新失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(metrics, "刷新成功", c)
}

// GetNodeMetrics 获取节点指标
// @Tags K8sMetrics
// @Summary 获取节点指标
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Success 200 {object} response.Response{data=[]service.NodeMetrics,msg=string} "获取成功"
// @Router /k8s/metrics/nodes [get]
func (a *K8sMetricsApi) GetNodeMetrics(c *gin.Context) {
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	metrics, err := service.ServiceGroupApp.K8sMetricsService.GetNodeMetrics(clusterName)
	if err != nil {
		global.GVA_LOG.Error("获取节点指标失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(metrics, "获取成功", c)
}

// GetPodMetrics 获取Pod指标
// @Tags K8sMetrics
// @Summary 获取Pod指标
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Success 200 {object} response.Response{data=[]service.PodMetrics,msg=string} "获取成功"
// @Router /k8s/metrics/pods [get]
func (a *K8sMetricsApi) GetPodMetrics(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")

	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	metrics, err := service.ServiceGroupApp.K8sMetricsService.GetPodMetrics(clusterName, namespace)
	if err != nil {
		global.GVA_LOG.Error("获取Pod指标失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(metrics, "获取成功", c)
}

// GetMetricsSummary 获取指标摘要
// @Tags K8sMetrics
// @Summary 获取指标摘要
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /k8s/metrics/summary [get]
func (a *K8sMetricsApi) GetMetricsSummary(c *gin.Context) {
	summary := service.ServiceGroupApp.K8sMetricsService.GetMetricsSummary()
	response.OkWithDetailed(summary, "获取成功", c)
}

// StartAutoCollector 启动自动收集
// @Tags K8sMetrics
// @Summary 启动自动收集
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "启动成功"
// @Router /k8s/metrics/collector/start [post]
func (a *K8sMetricsApi) StartAutoCollector(c *gin.Context) {
	service.ServiceGroupApp.K8sMetricsService.StartAutoCollector()
	response.OkWithMessage("自动收集已启动", c)
}

// StopAutoCollector 停止自动收集
// @Tags K8sMetrics
// @Summary 停止自动收集
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "停止成功"
// @Router /k8s/metrics/collector/stop [post]
func (a *K8sMetricsApi) StopAutoCollector(c *gin.Context) {
	service.ServiceGroupApp.K8sMetricsService.StopAutoCollector()
	response.OkWithMessage("自动收集已停止", c)
}
