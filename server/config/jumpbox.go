package config

// Jumpbox SSH跳板机配置
type Jumpbox struct {
	Enabled  bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`   // 是否启用
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`           // SSH端口
	HostKey  string `mapstructure:"host-key" json:"host-key" yaml:"host-key"` // 主机私钥路径（可选，不设置则自动生成）
	Banner   string `mapstructure:"banner" json:"banner" yaml:"banner"`     // SSH欢迎信息
}

