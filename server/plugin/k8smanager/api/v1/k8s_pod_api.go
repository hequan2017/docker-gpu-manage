package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
)

type K8sPodApi struct{}

// GetPodList 获取Pod列表
// @Tags K8sPod
// @Summary 获取Pod列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.K8sPodSearch true "查询条件"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/pod/list [get]
func (a *K8sPodApi) GetPodList(c *gin.Context) {
	var info request.K8sPodSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	podList, err := service.ServiceGroupApp.K8sPodService.GetPodList(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取Pod列表失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(podList, "获取成功", c)
}

// GetPod 获取Pod详情
// @Tags K8sPod
// @Summary 获取Pod详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param podName query string true "Pod名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/pod/get [get]
func (a *K8sPodApi) GetPod(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	podName := c.Query("podName")

	if clusterName == "" || namespace == "" || podName == "" {
		response.FailWithMessage("集群名称、命名空间和Pod名称不能为空", c)
		return
	}

	pod, err := service.ServiceGroupApp.K8sPodService.GetPod(c.Request.Context(), clusterName, namespace, podName)
	if err != nil {
		global.GVA_LOG.Error("获取Pod详情失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(pod, "获取成功", c)
}

// DeletePod 删除Pod
// @Tags K8sPod
// @Summary 删除Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteK8sResourceRequest true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/pod/delete [delete]
func (a *K8sPodApi) DeletePod(c *gin.Context) {
	var req request.DeleteK8sResourceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sPodService.DeletePod(c.Request.Context(), req.ClusterName, req.Namespace, req.Name)
	if err != nil {
		global.GVA_LOG.Error("删除Pod失败!", err)
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetPodLog 获取Pod日志
// @Tags K8sPod
// @Summary 获取Pod日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetPodLogRequest true "查询参数"
// @Success 200 {object} response.Response{data=string,msg=string} "获取成功"
// @Router /k8s/pod/log [post]
func (a *K8sPodApi) GetPodLog(c *gin.Context) {
	var req request.GetPodLogRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	log, err := service.ServiceGroupApp.K8sPodService.GetPodLog(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("获取Pod日志失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(log, "获取成功", c)
}

// GetPodContainers 获取Pod容器列表
// @Tags K8sPod
// @Summary 获取Pod容器列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param podName query string true "Pod名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/pod/containers [get]
func (a *K8sPodApi) GetPodContainers(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	podName := c.Query("podName")

	if clusterName == "" || namespace == "" || podName == "" {
		response.FailWithMessage("集群名称、命名空间和Pod名称不能为空", c)
		return
	}

	containers, initContainers, err := service.ServiceGroupApp.K8sPodService.GetPodContainers(c.Request.Context(), clusterName, namespace, podName)
	if err != nil {
		global.GVA_LOG.Error("获取Pod容器列表失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	result := map[string]interface{}{
		"containers":     containers,
		"initContainers": initContainers,
	}

	response.OkWithDetailed(result, "获取成功", c)
}

// GetPodEvents 获取Pod事件
// @Tags K8sPod
// @Summary 获取Pod事件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param podName query string true "Pod名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/pod/events [get]
func (a *K8sPodApi) GetPodEvents(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	podName := c.Query("podName")

	if clusterName == "" || namespace == "" || podName == "" {
		response.FailWithMessage("集群名称、命名空间和Pod名称不能为空", c)
		return
	}

	events, err := service.ServiceGroupApp.K8sPodService.GetPodEvents(c.Request.Context(), clusterName, namespace, podName)
	if err != nil {
		global.GVA_LOG.Error("获取Pod事件失败!", err)
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(events, "获取成功", c)
}
