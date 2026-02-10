package instance

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

const (
	releaseDraft  = "draft"
	releaseCanary = "canary"
	releaseStable = "stable"
)

type SchedulingRequestMeta struct {
	Region   string
	ISP      string
	UserHash string
}

type CandidateScoreDetail struct {
	NodeID           uint    `json:"nodeId"`
	NodeName         string  `json:"nodeName"`
	TotalScore       float64 `json:"totalScore"`
	GPUScore         float64 `json:"gpuScore"`
	CPUScore         float64 `json:"cpuScore"`
	MemoryScore      float64 `json:"memoryScore"`
	DiskScore        float64 `json:"diskScore"`
	RuleMatched      string  `json:"ruleMatched"`
	StrategyVersion  string  `json:"strategyVersion"`
	StrategyState    string  `json:"strategyState"`
	CircuitBreakerOn bool    `json:"circuitBreakerOn"`
}

type schedulerRuntime struct {
	mu                sync.Mutex
	requestCount      int
	failureCount      int
	totalLatencyMs    int64
	circuitOpened     bool
	lastStableNodes   []AvailableNode
	lastStableVersion string
}

var pcdnSchedulerRuntime = &schedulerRuntime{}

func chooseStrategy(meta SchedulingRequestMeta) (version string, state string, rule string, forceStable bool) {
	cfg := global.GVA_CONFIG.PCDN
	version = strings.TrimSpace(cfg.Version)
	if version == "" {
		version = "v1"
	}
	state = strings.ToLower(strings.TrimSpace(cfg.ReleaseState))
	if state == "" {
		state = releaseStable
	}
	stableVersion := strings.TrimSpace(cfg.StableVersion)
	if stableVersion == "" {
		stableVersion = "v1"
	}

	// 熔断已打开时，无条件回切到上一稳定策略
	if pcdnSchedulerRuntime.isCircuitOpened() {
		return stableVersion, releaseStable, "circuit_breaker_fallback", true
	}

	switch state {
	case releaseStable:
		return version, releaseStable, "stable_release", false
	case releaseDraft:
		// draft 默认不对外放量，回退稳定版本
		return stableVersion, releaseStable, "draft_guard", true
	case releaseCanary:
		if !cfg.Enabled {
			return stableVersion, releaseStable, "pcdn_disabled", true
		}
		if canaryMatched(meta, cfg.Canary.Regions, cfg.Canary.ISPs, cfg.Canary.UserHashPercentage) {
			return version, releaseCanary, "canary_hit", false
		}
		return stableVersion, releaseStable, "canary_miss", true
	default:
		return stableVersion, releaseStable, "unknown_state_fallback", true
	}
}

func canaryMatched(meta SchedulingRequestMeta, regions []string, isps []string, hashPercent int) bool {
	regionMatched := len(regions) == 0 || containsIgnoreCase(regions, meta.Region)
	ispMatched := len(isps) == 0 || containsIgnoreCase(isps, meta.ISP)
	if !regionMatched || !ispMatched {
		return false
	}
	if hashPercent <= 0 {
		return false
	}
	if hashPercent >= 100 {
		return true
	}
	if strings.TrimSpace(meta.UserHash) == "" {
		return false
	}
	bucket := int(crc32.ChecksumIEEE([]byte(meta.UserHash)) % 100)
	return bucket < hashPercent
}

func containsIgnoreCase(arr []string, v string) bool {
	for _, item := range arr {
		if strings.EqualFold(strings.TrimSpace(item), strings.TrimSpace(v)) {
			return true
		}
	}
	return false
}

func scoreAndSortNodes(nodes []AvailableNode, version, state, rule string, breakerOn bool) ([]AvailableNode, []CandidateScoreDetail) {
	weights := global.GVA_CONFIG.PCDN.ScoreWeight
	if weights.Gpu == 0 && weights.CPU == 0 && weights.Memory == 0 && weights.Disk == 0 {
		weights.Gpu = 0.4
		weights.CPU = 0.25
		weights.Memory = 0.25
		weights.Disk = 0.1
	}
	maxGpu, maxCPU, maxMem, maxDisk := int64(1), int64(1), int64(1), int64(1)
	for _, n := range nodes {
		if n.AvailableGpu > maxGpu {
			maxGpu = n.AvailableGpu
		}
		if n.AvailableCpu > maxCPU {
			maxCPU = n.AvailableCpu
		}
		if n.AvailableMemory > maxMem {
			maxMem = n.AvailableMemory
		}
		disk := n.AvailableSystemDisk + n.AvailableDataDisk
		if disk > maxDisk {
			maxDisk = disk
		}
	}

	details := make([]CandidateScoreDetail, 0, len(nodes))
	type scoredNode struct {
		node  AvailableNode
		score float64
	}
	scored := make([]scoredNode, 0, len(nodes))

	for _, n := range nodes {
		gpu := float64(n.AvailableGpu) / float64(maxGpu)
		cpu := float64(n.AvailableCpu) / float64(maxCPU)
		mem := float64(n.AvailableMemory) / float64(maxMem)
		diskVal := n.AvailableSystemDisk + n.AvailableDataDisk
		disk := float64(diskVal) / float64(maxDisk)
		total := gpu*weights.Gpu + cpu*weights.CPU + mem*weights.Memory + disk*weights.Disk
		scored = append(scored, scoredNode{node: n, score: total})
		details = append(details, CandidateScoreDetail{
			NodeID:           n.ID,
			NodeName:         n.Name,
			TotalScore:       total,
			GPUScore:         gpu,
			CPUScore:         cpu,
			MemoryScore:      mem,
			DiskScore:        disk,
			RuleMatched:      rule,
			StrategyVersion:  version,
			StrategyState:    state,
			CircuitBreakerOn: breakerOn,
		})
	}

	sort.SliceStable(scored, func(i, j int) bool {
		if scored[i].score == scored[j].score {
			return scored[i].node.ID < scored[j].node.ID
		}
		return scored[i].score > scored[j].score
	})
	result := make([]AvailableNode, 0, len(scored))
	for _, s := range scored {
		result = append(result, s.node)
	}
	return result, details
}

func (r *schedulerRuntime) isCircuitOpened() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.circuitOpened
}

func (r *schedulerRuntime) recordAndEvaluate(version string, state string, nodes []AvailableNode, err error, latency time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.requestCount++
	r.totalLatencyMs += latency.Milliseconds()
	if err != nil || len(nodes) == 0 {
		r.failureCount++
	}
	// 刷新稳定策略快照
	if state == releaseStable && err == nil && len(nodes) > 0 {
		r.lastStableNodes = append([]AvailableNode(nil), nodes...)
		r.lastStableVersion = version
	}

	cfg := global.GVA_CONFIG.PCDN.CircuitBreaker
	minSamples := cfg.MinSamples
	if minSamples <= 0 {
		minSamples = 10
	}
	if r.requestCount < minSamples {
		return
	}

	failureRate := float64(r.failureCount) / float64(r.requestCount)
	avgLatency := float64(r.totalLatencyMs) / float64(r.requestCount)
	failThreshold := cfg.FailureRateThreshold
	if failThreshold <= 0 {
		failThreshold = 0.3
	}
	latencyThreshold := float64(cfg.LatencyThresholdMs)
	if latencyThreshold <= 0 {
		latencyThreshold = 800
	}

	if failureRate >= failThreshold || avgLatency >= latencyThreshold {
		if !r.circuitOpened {
			global.GVA_LOG.Warn("PCDN调度熔断触发，自动回切稳定策略",
				zap.Float64("failureRate", failureRate),
				zap.Float64("avgLatencyMs", avgLatency),
				zap.Float64("failureRateThreshold", failThreshold),
				zap.Float64("latencyThresholdMs", latencyThreshold),
			)
		}
		r.circuitOpened = true
	} else if r.circuitOpened {
		r.circuitOpened = false
		global.GVA_LOG.Info("PCDN调度熔断恢复，重新开放灰度")
	}
}

func (r *schedulerRuntime) stableFallbackNodes() []AvailableNode {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.lastStableNodes) == 0 {
		return nil
	}
	return append([]AvailableNode(nil), r.lastStableNodes...)
}

func writeSchedulingTraceLog(specID string, version string, state string, rule string, meta SchedulingRequestMeta, details []CandidateScoreDetail, nodeCount int, err error) {
	if err != nil {
		global.GVA_LOG.Warn("任务调度日志",
			zap.String("specId", specID),
			zap.String("strategyVersion", version),
			zap.String("releaseState", state),
			zap.String("hitRule", rule),
			zap.String("region", meta.Region),
			zap.String("isp", meta.ISP),
			zap.Int("candidateCount", nodeCount),
			zap.String("error", err.Error()),
		)
		return
	}
	global.GVA_LOG.Info("任务调度日志",
		zap.String("specId", specID),
		zap.String("strategyVersion", version),
		zap.String("releaseState", state),
		zap.String("hitRule", rule),
		zap.String("region", meta.Region),
		zap.String("isp", meta.ISP),
		zap.Int("candidateCount", nodeCount),
		zap.String("candidateScoreDetail", fmt.Sprintf("%+v", details)),
	)
}
