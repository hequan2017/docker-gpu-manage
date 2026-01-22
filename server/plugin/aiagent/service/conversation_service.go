package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	"gorm.io/gorm"
)

var Conversation = new(conversation)

type conversation struct{}

// CreateConversation 创建会话
func (s *conversation) CreateConversation(conversation *model.Conversation) (err error) {
	err = global.GVA_DB.Create(conversation).Error
	return err
}

// DeleteConversation 删除会话（级联删除消息）
func (s *conversation) DeleteConversation(ID string) (err error) {
	// 先删除关联的消息
	err = global.GVA_DB.Where("conversation_id = ?", ID).Delete(&model.Message{}).Error
	if err != nil {
		return err
	}
	// 再删除会话
	err = global.GVA_DB.Delete(&model.Conversation{}, "id = ?", ID).Error
	return err
}

// UpdateConversation 更新会话
func (s *conversation) UpdateConversation(conversation model.Conversation) (err error) {
	err = global.GVA_DB.Model(&model.Conversation{}).Where("id = ?", conversation.ID).Updates(&conversation).Error
	return err
}

// GetConversation 根据ID获取会话
func (s *conversation) GetConversation(ID string) (conversation model.Conversation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&conversation).Error
	return
}

// GetConversationList 分页获取会话列表
func (s *conversation) GetConversationList(info request.ConversationSearch) (list []model.Conversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Conversation{})
	var conversations []model.Conversation

	// 条件搜索
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Model != nil && *info.Model != "" {
		db = db.Where("model = ?", *info.Model)
	}
	if info.IsActive != nil {
		db = db.Where("is_active = ?", *info.IsActive)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 按创建时间倒序
	err = db.Order("created_at DESC").Find(&conversations).Error
	return conversations, total, err
}

// SetConversationActive 设置会话激活状态（同时将其他会话设为非激活）
func (s *conversation) SetConversationActive(ID string, userID int, isActive bool) (err error) {
	// 如果激活当前会话，需要将其他会话设为非激活
	if isActive {
		err = global.GVA_DB.Model(&model.Conversation{}).
			Where("user_id = ? AND id != ?", userID, ID).
			Update("is_active", false).Error
		if err != nil {
			return err
		}
	}
	return global.GVA_DB.Model(&model.Conversation{}).
		Where("id = ? AND user_id = ?", ID, userID).
		Update("is_active", isActive).Error
}

// GetActiveConversation 获取用户的激活会话
func (s *conversation) GetActiveConversation(userID int) (conversation model.Conversation, err error) {
	err = global.GVA_DB.Where("user_id = ? AND is_active = ?", userID, true).
		Order("created_at DESC").
		First(&conversation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return conversation, nil // 没有激活的会话，返回空而不是错误
		}
		return
	}
	return
}
