package pcdn

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PcdnMetricSnapshot PCDN 节点周期指标快照
type PcdnMetricSnapshot struct {
	global.GVA_MODEL
	NodeID               uint      `json:"nodeId" gorm:"comment:节点ID;column:node_id;index:idx_node_reported_at,priority:1"`
	ReportedAt           time.Time `json:"reportedAt" gorm:"comment:节点上报时间;column:reported_at;index:idx_node_reported_at,priority:2"`
	RTTMs                float64   `json:"rttMs" gorm:"comment:RTT毫秒;column:rtt_ms;type:decimal(10,3)"`
	PacketLossRate       float64   `json:"packetLossRate" gorm:"comment:丢包率(0~1);column:packet_loss_rate;type:decimal(10,6)"`
	BandwidthUtilization float64   `json:"bandwidthUtilization" gorm:"comment:带宽占用率(0~1);column:bandwidth_utilization;type:decimal(10,6)"`
	CacheHitRate         float64   `json:"cacheHitRate" gorm:"comment:缓存命中率(0~1);column:cache_hit_rate;type:decimal(10,6)"`
	ActiveConnections    int64     `json:"activeConnections" gorm:"comment:活跃连接数;column:active_connections"`
	HealthScore          float64   `json:"healthScore" gorm:"comment:健康分(0~100);column:health_score;type:decimal(10,3)"`
	Window1mScore        float64   `json:"window1mScore" gorm:"comment:1分钟窗口健康分;column:window_1m_score;type:decimal(10,3)"`
	Window5mScore        float64   `json:"window5mScore" gorm:"comment:5分钟窗口健康分;column:window_5m_score;type:decimal(10,3)"`
	Window15mScore       float64   `json:"window15mScore" gorm:"comment:15分钟窗口健康分;column:window_15m_score;type:decimal(10,3)"`
}

func (PcdnMetricSnapshot) TableName() string {
	return "pcdn_metric_snapshot"
}
