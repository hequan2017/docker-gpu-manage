package pcdn

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
)

type MetricsService struct{}

func (m *MetricsService) ReportMetrics(ctx context.Context, req pcdnReq.PcdnMetricsReport) (pcdn.PcdnMetricSnapshot, float64, error) {
	reportedAt := req.ReportedAt
	if reportedAt.IsZero() {
		reportedAt = time.Now()
	}

	healthSvc := HealthScoreService{}
	healthScore := healthSvc.CalculateFromMetrics(req.RTTMs, req.PacketLossRate, req.BandwidthUtilization, req.CacheHitRate, req.ActiveConnections)
	window1m, window5m, window15m, err := healthSvc.CalculateWindowScores(ctx, req.NodeID, healthScore)
	if err != nil {
		return pcdn.PcdnMetricSnapshot{}, 0, err
	}

	snapshot := pcdn.PcdnMetricSnapshot{
		NodeID:               req.NodeID,
		ReportedAt:           reportedAt,
		RTTMs:                req.RTTMs,
		PacketLossRate:       req.PacketLossRate,
		BandwidthUtilization: req.BandwidthUtilization,
		CacheHitRate:         req.CacheHitRate,
		ActiveConnections:    req.ActiveConnections,
		HealthScore:          healthScore,
		Window1mScore:        window1m,
		Window5mScore:        window5m,
		Window15mScore:       window15m,
	}
	if err = global.GVA_DB.WithContext(ctx).Create(&snapshot).Error; err != nil {
		return pcdn.PcdnMetricSnapshot{}, 0, err
	}

	weight := healthSvc.SchedulerWeight(window1m, window5m, window15m)
	return snapshot, weight, nil
}

func (m *MetricsService) GetLatestSnapshot(ctx context.Context, nodeID uint) (pcdn.PcdnMetricSnapshot, error) {
	var snapshot pcdn.PcdnMetricSnapshot
	err := global.GVA_DB.WithContext(ctx).Where("node_id = ?", nodeID).Order("reported_at desc").First(&snapshot).Error
	return snapshot, err
}
