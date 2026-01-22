package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// MessageSearch 消息搜索结构体
type MessageSearch struct {
	ConversationID *uint   `json:"conversationID" form:"conversationID"` // 会话ID
	Role           *string `json:"role" form:"role"`                    // 角色过滤
	request.PageInfo
}
