package pcdn

import "math"

// RealtimeNodeMetric 节点实时指标。
type RealtimeNodeMetric struct {
	NodeID         uint
	LatencyMS      float64
	UnitCost       float64
	LoadPercent    float64
	HealthScore    float64
	Online         bool
	PolicyDisabled bool
	ISP            string
}

// ScoreWeights 策略权重。
type ScoreWeights struct {
	Latency float64
	Cost    float64
	Load    float64
	Health  float64
}

// PolicyEngine 执行加权评分。
type PolicyEngine struct {
	weights ScoreWeights
}

func NewPolicyEngine(weights ScoreWeights) *PolicyEngine {
	if weights == (ScoreWeights{}) {
		weights = ScoreWeights{Latency: 0.35, Cost: 0.2, Load: 0.2, Health: 0.25}
	}
	return &PolicyEngine{weights: weights}
}

// Score 将指标归一化后进行加权，分数越高越优。
func (p *PolicyEngine) Score(metric RealtimeNodeMetric) float64 {
	latencyScore := 1 / (1 + math.Max(metric.LatencyMS, 0)/50)
	costScore := 1 / (1 + math.Max(metric.UnitCost, 0))
	loadScore := 1 - math.Min(math.Max(metric.LoadPercent, 0), 100)/100
	healthScore := math.Min(math.Max(metric.HealthScore, 0), 100) / 100

	return p.weights.Latency*latencyScore +
		p.weights.Cost*costScore +
		p.weights.Load*loadScore +
		p.weights.Health*healthScore
}
