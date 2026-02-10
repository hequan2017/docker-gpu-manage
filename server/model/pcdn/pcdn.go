package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// PcdnNode PCDN 节点基础信息
// 包含 ISP、地域、带宽上下行、在线状态、健康度、成本权重。
type PcdnNode struct {
	global.GVA_MODEL
	NodeID          string  `json:"nodeId" form:"nodeId" gorm:"column:node_id;size:128;not null;uniqueIndex:uk_pcdn_node_node_id;comment:节点ID"`
	Name            string  `json:"name" form:"name" gorm:"column:name;size:255;not null;comment:节点名称"`
	ISP             string  `json:"isp" form:"isp" gorm:"column:isp;size:64;not null;index:idx_pcdn_node_region_isp,priority:2;comment:运营商"`
	Region          string  `json:"region" form:"region" gorm:"column:region;size:128;not null;index:idx_pcdn_node_region_isp,priority:1;comment:地域"`
	UpstreamMbps    int64   `json:"upstreamMbps" form:"upstreamMbps" gorm:"column:upstream_mbps;not null;default:0;comment:上行带宽Mbps"`
	DownstreamMbps  int64   `json:"downstreamMbps" form:"downstreamMbps" gorm:"column:downstream_mbps;not null;default:0;comment:下行带宽Mbps"`
	OnlineStatus    string  `json:"onlineStatus" form:"onlineStatus" gorm:"column:online_status;size:32;not null;default:'offline';index;comment:在线状态"`
	HealthScore     float64 `json:"healthScore" form:"healthScore" gorm:"column:health_score;type:decimal(5,2);not null;default:0;comment:健康度"`
	CostWeight      float64 `json:"costWeight" form:"costWeight" gorm:"column:cost_weight;type:decimal(8,4);not null;default:1;comment:成本权重"`
	LastHeartbeatAt *int64  `json:"lastHeartbeatAt" form:"lastHeartbeatAt" gorm:"column:last_heartbeat_at;comment:最近心跳时间戳"`
}

func (PcdnNode) TableName() string {
	return "pcdn_node"
}

// PcdnResource PCDN 节点资源能力。
type PcdnResource struct {
	global.GVA_MODEL
	NodeID             string `json:"nodeId" form:"nodeId" gorm:"column:node_id;size:128;not null;index:idx_pcdn_resource_node_id;comment:节点ID"`
	ContentTypes       string `json:"contentTypes" form:"contentTypes" gorm:"column:content_types;type:text;not null;comment:可提供内容类型(JSON数组)"`
	CacheCapacityGB    int64  `json:"cacheCapacityGb" form:"cacheCapacityGb" gorm:"column:cache_capacity_gb;not null;default:0;comment:缓存容量GB"`
	MaxConcurrency     int64  `json:"maxConcurrency" form:"maxConcurrency" gorm:"column:max_concurrency;not null;default:0;comment:最大并发能力"`
	CurrentConcurrency int64  `json:"currentConcurrency" form:"currentConcurrency" gorm:"column:current_concurrency;not null;default:0;comment:当前并发"`
}

func (PcdnResource) TableName() string {
	return "pcdn_resource"
}

// PcdnPolicy PCDN 调度策略配置。
type PcdnPolicy struct {
	global.GVA_MODEL
	PolicyName    string  `json:"policyName" form:"policyName" gorm:"column:policy_name;size:128;not null;uniqueIndex:uk_pcdn_policy_name;comment:策略名称"`
	StrategyType  string  `json:"strategyType" form:"strategyType" gorm:"column:strategy_type;size:32;not null;default:'hybrid';comment:策略类型 latency/cost/hybrid"`
	LatencyWeight float64 `json:"latencyWeight" form:"latencyWeight" gorm:"column:latency_weight;type:decimal(6,4);not null;default:0.5;comment:延迟权重"`
	CostWeight    float64 `json:"costWeight" form:"costWeight" gorm:"column:cost_weight;type:decimal(6,4);not null;default:0.5;comment:成本权重"`
	RateLimitQPS  int64   `json:"rateLimitQps" form:"rateLimitQps" gorm:"column:rate_limit_qps;not null;default:0;comment:限流阈值QPS"`
	Whitelist     string  `json:"whitelist" form:"whitelist" gorm:"column:whitelist;type:text;comment:白名单(JSON数组)"`
	Blacklist     string  `json:"blacklist" form:"blacklist" gorm:"column:blacklist;type:text;comment:黑名单(JSON数组)"`
	IsEnabled     bool    `json:"isEnabled" form:"isEnabled" gorm:"column:is_enabled;not null;default:true;comment:是否启用"`
}

func (PcdnPolicy) TableName() string {
	return "pcdn_policy"
}

// PcdnDispatchTask PCDN 调度任务。
// 包含任务状态机、失败重试计数、trace_id。
type PcdnDispatchTask struct {
	global.GVA_MODEL
	TaskID        string `json:"taskId" form:"taskId" gorm:"column:task_id;size:128;not null;uniqueIndex:uk_pcdn_dispatch_task_id;comment:任务ID"`
	TraceID       string `json:"traceId" form:"traceId" gorm:"column:trace_id;size:128;not null;index;comment:链路追踪ID"`
	PolicyID      uint   `json:"policyId" form:"policyId" gorm:"column:policy_id;not null;index;comment:关联策略ID"`
	SourceRegion  string `json:"sourceRegion" form:"sourceRegion" gorm:"column:source_region;size:128;not null;comment:请求源地域"`
	TargetContent string `json:"targetContent" form:"targetContent" gorm:"column:target_content;size:255;not null;comment:目标内容标识"`
	Status        string `json:"status" form:"status" gorm:"column:status;size:32;not null;default:'pending';index:idx_pcdn_dispatch_status_created,priority:1;comment:任务状态 pending/running/success/failed/cancelled"`
	RetryCount    int32  `json:"retryCount" form:"retryCount" gorm:"column:retry_count;not null;default:0;comment:失败重试计数"`
	MaxRetryCount int32  `json:"maxRetryCount" form:"maxRetryCount" gorm:"column:max_retry_count;not null;default:3;comment:最大重试次数"`
	LastError     string `json:"lastError" form:"lastError" gorm:"column:last_error;type:text;comment:最后一次错误信息"`
	StateVersion  int64  `json:"stateVersion" form:"stateVersion" gorm:"column:state_version;not null;default:1;comment:状态机版本"`
	CreatedAtUnix int64  `json:"createdAtUnix" form:"createdAtUnix" gorm:"column:created_at_unix;not null;default:0;index:idx_pcdn_dispatch_status_created,priority:2;comment:创建时间戳"`
}

func (PcdnDispatchTask) TableName() string {
	return "pcdn_dispatch_task"
}

// PcdnMetricSnapshot PCDN 节点实时指标快照。
type PcdnMetricSnapshot struct {
	global.GVA_MODEL
	NodeID               string  `json:"nodeId" form:"nodeId" gorm:"column:node_id;size:128;not null;index:idx_pcdn_metric_node_time,priority:1;comment:节点ID"`
	SnapshotAt           int64   `json:"snapshotAt" form:"snapshotAt" gorm:"column:snapshot_at;not null;index:idx_pcdn_metric_node_time,priority:2;comment:快照时间戳"`
	RttMs                float64 `json:"rttMs" form:"rttMs" gorm:"column:rtt_ms;type:decimal(8,3);not null;default:0;comment:RTT毫秒"`
	PacketLossRate       float64 `json:"packetLossRate" form:"packetLossRate" gorm:"column:packet_loss_rate;type:decimal(7,4);not null;default:0;comment:丢包率"`
	CacheHitRate         float64 `json:"cacheHitRate" form:"cacheHitRate" gorm:"column:cache_hit_rate;type:decimal(7,4);not null;default:0;comment:命中率"`
	BandwidthUtilization float64 `json:"bandwidthUtilization" form:"bandwidthUtilization" gorm:"column:bandwidth_utilization;type:decimal(7,4);not null;default:0;comment:带宽利用率"`
}

func (PcdnMetricSnapshot) TableName() string {
	return "pcdn_metric_snapshot"
}
