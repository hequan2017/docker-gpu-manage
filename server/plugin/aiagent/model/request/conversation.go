package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// ConversationSearch 会话搜索结构体
type ConversationSearch struct {
	Title    *string `json:"title" form:"title"`
	Model    *string `json:"model" form:"model"`
	IsActive *bool   `json:"isActive" form:"isActive"`
	request.PageInfo
}

// ChatRequest 对话请求结构体
type ChatRequest struct {
	ConversationID *uint    `json:"conversationID" form:"conversationID"` // 会话ID，为空则创建新会话
	Message        string   `json:"message" form:"message" binding:"required"` // 用户消息内容
	Model          *string  `json:"model" form:"model"`          // 指定模型（可选，优先使用会话配置）
	Temperature    *float64 `json:"temperature" form:"temperature"` // 温度参数（可选）
	MaxTokens      *int     `json:"maxTokens" form:"maxTokens"`   // 最大token数（可选）
	Stream         *bool    `json:"stream" form:"stream"`         // 是否流式输出（默认false）
}

// ChatResponse 对话响应结构体
type ChatResponse struct {
	ConversationID uint   `json:"conversationID"` // 会话ID
	MessageID      uint   `json:"messageID"`      // 消息ID
	Content        string `json:"content"`        // AI回复内容
	FinishReason   string `json:"finishReason"`   // 结束原因
	Usage          Usage   `json:"usage"`         // token使用情况
}

// Usage token使用情况
type Usage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
}
