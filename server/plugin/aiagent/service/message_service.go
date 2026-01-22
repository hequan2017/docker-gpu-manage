package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
)

var Message = new(message)

type message struct{}

// CreateMessage 创建消息
func (s *message) CreateMessage(message *model.Message) (err error) {
	err = global.GVA_DB.Create(message).Error
	return err
}

// DeleteMessage 删除消息
func (s *message) DeleteMessage(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Message{}, "id = ?", ID).Error
	return err
}

// GetMessage 根据ID获取消息
func (s *message) GetMessage(ID string) (message model.Message, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&message).Error
	return
}

// GetMessageList 分页获取消息列表
func (s *message) GetMessageList(info request.MessageSearch) (list []model.Message, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Message{})
	var messages []model.Message

	// 必须指定会话ID
	if info.ConversationID == nil {
		return messages, 0, nil
	}

	db = db.Where("conversation_id = ?", *info.ConversationID)

	// 角色过滤
	if info.Role != nil && *info.Role != "" {
		db = db.Where("role = ?", *info.Role)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 按创建时间正序
	err = db.Order("created_at ASC").Find(&messages).Error
	return messages, total, err
}

// GetMessagesByConversationID 获取会话的所有消息（不分页，用于对话上下文）
func (s *message) GetMessagesByConversationID(conversationID uint) (messages []model.Message, err error) {
	err = global.GVA_DB.Where("conversation_id = ?", conversationID).
		Order("created_at ASC").
		Find(&messages).Error
	return
}
