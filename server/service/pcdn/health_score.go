package pcdn

import (
	"context"
	"math"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type HealthScoreService struct{}

func (h *HealthScoreService) CalculateFromMetrics(rttMs, packetLossRate, bandwidthUtilization, cacheHitRate float64, activeConnections int64) float64 {
	rttScore := metricLinearScore(rttMs, 20, 300, true)
	packetLossScore := metricLinearScore(packetLossRate, 0, 0.1, true)
	bandwidthScore := metricLinearScore(bandwidthUtilization, 0.2, 0.95, true)
	cacheScore := metricLinearScore(cacheHitRate, 0.3, 0.95, false)
	connectionsScore := metricLinearScore(float64(activeConnections), 50, 5000, true)

	raw := rttScore*0.28 + packetLossScore*0.24 + bandwidthScore*0.18 + cacheScore*0.22 + connectionsScore*0.08
	return round2(clamp(raw, 0, 100))
}

func (h *HealthScoreService) CalculateWindowScores(ctx context.Context, nodeID uint, fallbackScore float64) (window1m, window5m, window15m float64, err error) {
	window1m, err = h.windowAvg(ctx, nodeID, time.Minute, fallbackScore)
	if err != nil {
		return 0, 0, 0, err
	}
	window5m, err = h.windowAvg(ctx, nodeID, 5*time.Minute, fallbackScore)
	if err != nil {
		return 0, 0, 0, err
	}
	window15m, err = h.windowAvg(ctx, nodeID, 15*time.Minute, fallbackScore)
	if err != nil {
		return 0, 0, 0, err
	}
	return round2(window1m), round2(window5m), round2(window15m), nil
}

func (h *HealthScoreService) SchedulerWeight(window1m, window5m, window15m float64) float64 {
	weighted := window1m*0.5 + window5m*0.3 + window15m*0.2
	return round4(clamp(weighted/100, 0.05, 1.0))
}

func (h *HealthScoreService) windowAvg(ctx context.Context, nodeID uint, window time.Duration, fallback float64) (float64, error) {
	type result struct {
		Avg   float64
		Count int64
	}
	var res result
	err := global.GVA_DB.WithContext(ctx).
		Table("pcdn_metric_snapshot").
		Select("COALESCE(AVG(health_score),0) as avg, COUNT(1) as count").
		Where("node_id = ? AND reported_at >= ?", nodeID, time.Now().Add(-window)).
		Scan(&res).Error
	if err != nil {
		return 0, err
	}
	if res.Count == 0 {
		return fallback, nil
	}
	return res.Avg, nil
}

func metricLinearScore(value, good, bad float64, lowerBetter bool) float64 {
	if lowerBetter {
		if value <= good {
			return 100
		}
		if value >= bad {
			return 0
		}
		return (bad - value) / (bad - good) * 100
	}
	if value >= bad {
		return 100
	}
	if value <= good {
		return 0
	}
	return (value - good) / (bad - good) * 100
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

func round4(v float64) float64 {
	return math.Round(v*10000) / 10000
}
