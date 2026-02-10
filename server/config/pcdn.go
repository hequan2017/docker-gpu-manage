package config

// PCDN 调度配置
// 支持策略发布状态、灰度规则以及熔断阈值动态调整
// 可通过热更新配置文件实时生效
// 发布状态：draft/canary/stable
// 灰度规则：地域、ISP、用户哈希比例
// 熔断规则：失败率、延迟阈值
// 注意：阈值单位详见字段注释
//
// 该结构由配置映射驱动，字段命名与 YAML 保持一致
// 为了避免歧义，这里保留较直观的命名。
//
//nolint:revive
type PCDN struct {
	Enabled        bool         `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Version        string       `mapstructure:"version" json:"version" yaml:"version"`
	ReleaseState   string       `mapstructure:"release-state" json:"release-state" yaml:"release-state"`
	Canary         Canary       `mapstructure:"canary" json:"canary" yaml:"canary"`
	CircuitBreaker CircuitBreak `mapstructure:"circuit-breaker" json:"circuit-breaker" yaml:"circuit-breaker"`
	ScoreWeight    ScoreWeight  `mapstructure:"score-weight" json:"score-weight" yaml:"score-weight"`
	StableVersion  string       `mapstructure:"stable-version" json:"stable-version" yaml:"stable-version"`
}

type Canary struct {
	Regions            []string `mapstructure:"regions" json:"regions" yaml:"regions"`
	ISPs               []string `mapstructure:"isps" json:"isps" yaml:"isps"`
	UserHashPercentage int      `mapstructure:"user-hash-percentage" json:"user-hash-percentage" yaml:"user-hash-percentage"`
}

type CircuitBreak struct {
	FailureRateThreshold float64 `mapstructure:"failure-rate-threshold" json:"failure-rate-threshold" yaml:"failure-rate-threshold"`
	LatencyThresholdMs   int64   `mapstructure:"latency-threshold-ms" json:"latency-threshold-ms" yaml:"latency-threshold-ms"`
	MinSamples           int     `mapstructure:"min-samples" json:"min-samples" yaml:"min-samples"`
}

type ScoreWeight struct {
	Gpu    float64 `mapstructure:"gpu" json:"gpu" yaml:"gpu"`
	CPU    float64 `mapstructure:"cpu" json:"cpu" yaml:"cpu"`
	Memory float64 `mapstructure:"memory" json:"memory" yaml:"memory"`
	Disk   float64 `mapstructure:"disk" json:"disk" yaml:"disk"`
}
