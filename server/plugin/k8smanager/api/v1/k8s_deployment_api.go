package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
)

type K8sDeploymentApi struct{}

// GetDeploymentList 获取Deployment列表
// @Tags K8sDeployment
// @Summary 获取Deployment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.K8sDeploymentSearch true "查询条件"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/deployment/list [get]
func (a *K8sDeploymentApi) GetDeploymentList(c *gin.Context) {
	var info request.K8sDeploymentSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	deployList, err := service.ServiceGroupApp.K8sDeploymentService.GetDeploymentList(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取Deployment列表失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(deployList, "获取成功", c)
}

// GetDeployment 获取Deployment详情
// @Tags K8sDeployment
// @Summary 获取Deployment详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param name query string true "Deployment名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/deployment/get [get]
func (a *K8sDeploymentApi) GetDeployment(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	name := c.Query("name")

	if clusterName == "" || namespace == "" || name == "" {
		response.FailWithMessage("集群名称、命名空间和Deployment名称不能为空", c)
		return
	}

	deploy, err := service.ServiceGroupApp.K8sDeploymentService.GetDeployment(c.Request.Context(), clusterName, namespace, name)
	if err != nil {
		global.GVA_LOG.Error("获取Deployment详情失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(deploy, "获取成功", c)
}

// ScaleDeployment 扩缩容Deployment
// @Tags K8sDeployment
// @Summary 扩缩容Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ScaleDeploymentRequest true "扩缩容参数"
// @Success 200 {object} response.Response{data=interface{},msg=string} "扩缩容成功"
// @Router /k8s/deployment/scale [post]
func (a *K8sDeploymentApi) ScaleDeployment(c *gin.Context) {
	var req request.ScaleDeploymentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	deploy, err := service.ServiceGroupApp.K8sDeploymentService.ScaleDeployment(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("扩缩容Deployment失败!", err)
		response.FailWithMessage("扩缩容失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(deploy, "扩缩容成功", c)
}

// RestartDeployment 重启Deployment
// @Tags K8sDeployment
// @Summary 重启Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RestartDeploymentRequest true "重启参数"
// @Success 200 {object} response.Response{msg=string} "重启成功"
// @Router /k8s/deployment/restart [post]
func (a *K8sDeploymentApi) RestartDeployment(c *gin.Context) {
	var req request.RestartDeploymentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sDeploymentService.RestartDeployment(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("重启Deployment失败!", err)
		response.FailWithMessage("重启失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("重启成功", c)
}

// DeleteDeployment 删除Deployment
// @Tags K8sDeployment
// @Summary 删除Deployment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteK8sResourceRequest true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/deployment/delete [delete]
func (a *K8sDeploymentApi) DeleteDeployment(c *gin.Context) {
	var req request.DeleteK8sResourceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sDeploymentService.DeleteDeployment(c.Request.Context(), req.ClusterName, req.Namespace, req.Name)
	if err != nil {
		global.GVA_LOG.Error("删除Deployment失败!", err)
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetDeploymentPods 获取Deployment关联的Pods
// @Tags K8sDeployment
// @Summary 获取Deployment关联的Pods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param name query string true "Deployment名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/deployment/pods [get]
func (a *K8sDeploymentApi) GetDeploymentPods(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	name := c.Query("name")

	if clusterName == "" || namespace == "" || name == "" {
		response.FailWithMessage("集群名称、命名空间和Deployment名称不能为空", c)
		return
	}

	pods, err := service.ServiceGroupApp.K8sDeploymentService.GetDeploymentPods(c.Request.Context(), clusterName, namespace, name)
	if err != nil {
		global.GVA_LOG.Error("获取Deployment Pods失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(pods, "获取成功", c)
}
