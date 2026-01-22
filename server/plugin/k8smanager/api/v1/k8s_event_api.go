package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sEventApi struct{}

// GetEventList 获取Event列表
// @Tags K8sEvent
// @Summary 获取Event列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetEventsRequest true "查询条件"
// @Success 200 {object} response.Response{data=interface{},msg=string} "获取成功"
// @Router /k8s/event/list [post]
func (a *K8sEventApi) GetEventList(c *gin.Context) {
	var req request.GetEventsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	eventList, err := service.ServiceGroupApp.K8sEventService.GetEventList(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("获取Event列表失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(eventList, "获取成功", c)
}
