package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sServiceApi struct{}

// GetServiceList 获取Service列表
// @Tags K8sService
// @Summary 获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.K8sServiceSearch true "查询条件"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/service/list [get]
func (a *K8sServiceApi) GetServiceList(c *gin.Context) {
	var info request.K8sServiceSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	svcList, err := service.ServiceGroupApp.K8sServiceService.GetServiceList(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取Service列表失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(svcList, "获取成功", c)
}

// GetService 获取Service详情
// @Tags K8sService
// @Summary 获取Service详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param name query string true "Service名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/service/get [get]
func (a *K8sServiceApi) GetService(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	name := c.Query("name")

	if clusterName == "" || namespace == "" || name == "" {
		response.FailWithMessage("集群名称、命名空间和Service名称不能为空", c)
		return
	}

	svc, err := service.ServiceGroupApp.K8sServiceService.GetService(c.Request.Context(), clusterName, namespace, name)
	if err != nil {
		global.GVA_LOG.Error("获取Service详情失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(svc, "获取成功", c)
}

// DeleteService 删除Service
// @Tags K8sService
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteK8sResourceRequest true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/service/delete [delete]
func (a *K8sServiceApi) DeleteService(c *gin.Context) {
	var req request.DeleteK8sResourceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sServiceService.DeleteService(c.Request.Context(), req.ClusterName, req.Namespace, req.Name)
	if err != nil {
		global.GVA_LOG.Error("删除Service失败!", zap.String("error", err.Error()))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetServiceEndpoints 获取Service的Endpoints
// @Tags K8sService
// @Summary 获取Service的Endpoints
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param name query string true "Service名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/service/endpoints [get]
func (a *K8sServiceApi) GetServiceEndpoints(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	name := c.Query("name")

	if clusterName == "" || namespace == "" || name == "" {
		response.FailWithMessage("集群名称、命名空间和Service名称不能为空", c)
		return
	}

	endpoints, err := service.ServiceGroupApp.K8sServiceService.GetServiceEndpoints(c.Request.Context(), clusterName, namespace, name)
	if err != nil {
		global.GVA_LOG.Error("获取Service Endpoints失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(endpoints, "获取成功", c)
}
