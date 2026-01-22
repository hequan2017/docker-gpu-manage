package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AgentConfig AI Agent配置 结构体
type AgentConfig struct {
	global.GVA_MODEL
	Name        string  `json:"name" form:"name" gorm:"column:name;comment:配置名称;type:varchar(100);uniqueIndex;not null;"` // 配置名称
	APIKey      string  `json:"apiKey" form:"apiKey" gorm:"column:api_key;comment:API Key;type:varchar(200);not null;"`    // API Key
	BaseURL     string  `json:"baseURL" form:"baseURL" gorm:"column:base_url;comment:API基础URL;type:varchar(500);default:https://open.bigmodel.cn/api/paas/v4/;"` // API基础URL
	Model       string  `json:"model" form:"model" gorm:"column:model;comment:默认模型;type:varchar(50);default:glm-4-plus;"` // 默认模型
	Temperature float64 `json:"temperature" form:"temperature" gorm:"column:temperature;comment:默认温度;default:0.7;"`      // 默认温度
	MaxTokens   int     `json:"maxTokens" form:"maxTokens" gorm:"column:max_tokens;comment:默认最大token数;default:4096;"`  // 默认最大token数
	IsActive    bool    `json:"isActive" form:"isActive" gorm:"column:is_active;comment:是否启用;default:true;"`            // 是否启用
}

// TableName AgentConfig 自定义表名 gva_aiagent_configs
func (AgentConfig) TableName() string {
	return "gva_aiagent_configs"
}
