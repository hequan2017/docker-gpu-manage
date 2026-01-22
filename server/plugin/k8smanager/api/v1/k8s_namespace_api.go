package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sNamespaceApi struct{}

// GetNamespaceList 获取Namespace列表
// @Tags K8sNamespace
// @Summary 获取Namespace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.K8sNamespaceSearch true "查询条件"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/namespace/list [get]
func (a *K8sNamespaceApi) GetNamespaceList(c *gin.Context) {
	var info request.K8sNamespaceSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	namespaceList, err := service.ServiceGroupApp.K8sNamespaceService.GetNamespaceList(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取Namespace列表失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(namespaceList, "获取成功", c)
}

// GetNamespace 获取Namespace详情
// @Tags K8sNamespace
// @Summary 获取Namespace详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param name query string true "Namespace名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/namespace/get [get]
func (a *K8sNamespaceApi) GetNamespace(c *gin.Context) {
	clusterName := c.Query("clusterName")
	name := c.Query("name")

	if clusterName == "" || name == "" {
		response.FailWithMessage("集群名称和Namespace名称不能为空", c)
		return
	}

	namespace, err := service.ServiceGroupApp.K8sNamespaceService.GetNamespace(c.Request.Context(), clusterName, name)
	if err != nil {
		global.GVA_LOG.Error("获取Namespace详情失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(namespace, "获取成功", c)
}

// CreateNamespace 创建Namespace
// @Tags K8sNamespace
// @Summary 创建Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateNamespace true "创建参数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8s/namespace/create [post]
type CreateNamespace struct {
	ClusterName string            `json:"clusterName" binding:"required"`
	Name        string            `json:"name" binding:"required"`
	Labels      map[string]string `json:"labels"`
}

func (a *K8sNamespaceApi) CreateNamespace(c *gin.Context) {
	var req CreateNamespace
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sNamespaceService.CreateNamespace(c.Request.Context(), req.ClusterName, req.Name, req.Labels)
	if err != nil {
		global.GVA_LOG.Error("创建Namespace失败!", zap.String("error", err.Error()))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteNamespace 删除Namespace
// @Tags K8sNamespace
// @Summary 删除Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteK8sResourceRequest true "删除参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/namespace/delete [delete]
func (a *K8sNamespaceApi) DeleteNamespace(c *gin.Context) {
	var req request.DeleteK8sResourceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sNamespaceService.DeleteNamespace(c.Request.Context(), req.ClusterName, req.Name)
	if err != nil {
		global.GVA_LOG.Error("删除Namespace失败!", zap.String("error", err.Error()))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
