package request

import "time"

type PcdnMetricsReport struct {
	NodeID               uint      `json:"nodeId" binding:"required"`
	ReportedAt           time.Time `json:"reportedAt"`
	RTTMs                float64   `json:"rttMs" binding:"required,gte=0"`
	PacketLossRate       float64   `json:"packetLossRate" binding:"required,gte=0,lte=1"`
	BandwidthUtilization float64   `json:"bandwidthUtilization" binding:"required,gte=0,lte=1"`
	CacheHitRate         float64   `json:"cacheHitRate" binding:"required,gte=0,lte=1"`
	ActiveConnections    int64     `json:"activeConnections" binding:"required,gte=0"`
}
