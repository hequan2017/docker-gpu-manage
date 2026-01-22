package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Conversation AI对话会话 结构体
type Conversation struct {
	global.GVA_MODEL
	Title       string  `json:"title" form:"title" gorm:"column:title;comment:会话标题;type:varchar(200);"`          // 会话标题
	UserID      *int    `json:"userID" form:"userID" gorm:"column:user_id;comment:所属用户;index"`                    // 所属用户
	Model       string  `json:"model" form:"model" gorm:"column:model;comment:使用的模型;type:varchar(50);default:glm-4-plus;"` // 使用的模型
	SystemPrompt *string `json:"systemPrompt" form:"systemPrompt" gorm:"column:system_prompt;comment:系统提示词;type:text;"` // 系统提示词
	Temperature *float64 `json:"temperature" form:"temperature" gorm:"column:temperature;comment:温度参数;default:0.7;"`  // 温度参数
	MaxTokens   *int    `json:"maxTokens" form:"maxTokens" gorm:"column:max_tokens;comment:最大token数;default:4096;"` // 最大token数
	IsActive    bool    `json:"isActive" form:"isActive" gorm:"column:is_active;comment:是否激活;default:true;"`       // 是否激活
}

// TableName Conversation 自定义表名 gva_aiagent_conversations
func (Conversation) TableName() string {
	return "gva_aiagent_conversations"
}
