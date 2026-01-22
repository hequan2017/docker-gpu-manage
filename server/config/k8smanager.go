package config

// K8sManager K8s管理器配置
type K8sManager struct {
	// EncryptionKey 用于加密 kubeconfig 的密钥（32字节）
	// 建议从环境变量 K8S_MANAGER_ENCRYPTION_KEY 获取
	// 如果不配置，将使用默认密钥（不推荐用于生产环境）
	EncryptionKey string `mapstructure:"encryption-key" json:"encryption-key" yaml:"encryption-key" env:"K8S_MANAGER_ENCRYPTION_KEY"`

	// ClientTTL 客户端连接缓存时间（秒）
	// 默认 300 秒（5分钟）
	ClientTTL int `mapstructure:"client-ttl" json:"client-ttl" yaml:"client-ttl"`

	// MaxClients 最大客户端连接数
	// 默认 100
	MaxClients int `mapstructure:"max-clients" json:"max-clients" yaml:"max-clients"`
}
