package pcdn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PcdnDispatchTaskApi struct{}

func (a *PcdnDispatchTaskApi) GetPcdnDispatchTaskList(c *gin.Context) {
	var req pcdnReq.PcdnDispatchTaskSearch
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pcdnDispatchTaskService.GetPcdnDispatchTaskList(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("获取PCDN调度任务列表失败", zap.Error(err))
		response.FailWithMessage("获取PCDN调度任务列表失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: req.Page, PageSize: req.PageSize}, "获取PCDN调度任务列表成功", c)
}
