package pcdn

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PcdnMetricsApi struct{}

func (p *PcdnMetricsApi) ReportMetrics(c *gin.Context) {
	ctx := c.Request.Context()
	var req pcdnReq.PcdnMetricsReport
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if nodeIDFromToken, ok := c.Get("pcdnNodeID"); ok {
		if id, ok := nodeIDFromToken.(uint); ok && id != req.NodeID {
			response.FailWithMessage("节点身份与上报节点不匹配", c)
			return
		}
	}

	snapshot, weight, err := pcdnMetricsService.ReportMetrics(ctx, req)
	if err != nil {
		global.GVA_LOG.Error("上报PCDN指标失败", zap.Error(err), zap.Uint("nodeId", req.NodeID))
		response.FailWithMessage("上报失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"snapshot":        snapshot,
		"schedulerWeight": weight,
	}, "上报成功", c)
}

func (p *PcdnMetricsApi) GetLatestSnapshot(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		NodeID uint `form:"nodeId" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	snapshot, err := pcdnMetricsService.GetLatestSnapshot(ctx, req.NodeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("节点暂无指标快照", c)
			return
		}
		global.GVA_LOG.Error("查询PCDN指标失败", zap.Error(err), zap.Uint("nodeId", req.NodeID))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	weight := healthScoreService.SchedulerWeight(snapshot.Window1mScore, snapshot.Window5mScore, snapshot.Window15mScore)
	response.OkWithData(gin.H{
		"snapshot":        snapshot,
		"schedulerWeight": weight,
	}, c)
}
