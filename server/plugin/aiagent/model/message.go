package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Message AI对话消息 结构体
type Message struct {
	global.GVA_MODEL
	ConversationID uint   `json:"conversationID" form:"conversationID" gorm:"column:conversation_id;comment:会话ID;index:idx_conversation_id;not null;"` // 会话ID
	Role           string `json:"role" form:"role" gorm:"column:role;comment:角色(user/assistant/system);type:varchar(20);not null;"`                     // 角色
	Content        string `json:"content" form:"content" gorm:"column:content;comment:消息内容;type:text;not null;"`                                     // 消息内容
	TokenCount     *int   `json:"tokenCount" form:"tokenCount" gorm:"column:token_count;comment:token数量;"`                                            // token数量
	Metadata       datatypes.JSON `json:"metadata" form:"metadata" gorm:"column:metadata;comment:元数据(如finish_reason等);" swaggertype:"array,object"`  // 元数据
}

// TableName Message 自定义表名 gva_aiagent_messages
func (Message) TableName() string {
	return "gva_aiagent_messages"
}
