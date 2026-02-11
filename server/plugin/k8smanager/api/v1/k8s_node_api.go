package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sNodeApi struct{}

// GetNodeList 获取Node列表
// @Tags K8sNode
// @Summary 获取Node列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/node/list [get]
func (a *K8sNodeApi) GetNodeList(c *gin.Context) {
	clusterName := c.Query("clusterName")
	if clusterName == "" {
		response.FailWithMessage("集群名称不能为空", c)
		return
	}

	nodeList, err := service.ServiceGroupApp.K8sNodeService.GetNodeList(c.Request.Context(), clusterName)
	if err != nil {
		global.GVA_LOG.Error("获取Node列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(nodeList, "获取成功", c)
}

// GetNode 获取Node详情
// @Tags K8sNode
// @Summary 获取Node详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param nodeName query string true "Node名称"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/node/get [get]
func (a *K8sNodeApi) GetNode(c *gin.Context) {
	clusterName := c.Query("clusterName")
	nodeName := c.Query("nodeName")

	if clusterName == "" || nodeName == "" {
		response.FailWithMessage("集群名称和Node名称不能为空", c)
		return
	}

	node, err := service.ServiceGroupApp.K8sNodeService.GetNode(c.Request.Context(), clusterName, nodeName)
	if err != nil {
		global.GVA_LOG.Error("获取Node详情失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(node, "获取成功", c)
}

// CordonNode 设置Node调度状态
// @Tags K8sNode
// @Summary 设置Node调度状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CordonNodeRequest true "请求参数"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /k8s/node/cordon [post]
func (a *K8sNodeApi) CordonNode(c *gin.Context) {
	var req request.CordonNodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.K8sNodeService.CordonNode(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("设置Node调度状态失败!", zap.Error(err))
		response.FailWithMessage("操作失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("操作成功", c)
}
