package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"gorm.io/gorm"
)

var Config = new(config)

type config struct{}

// CreateConfig 创建配置
func (s *config) CreateConfig(config *model.AgentConfig) (err error) {
	err = global.GVA_DB.Create(config).Error
	return err
}

// DeleteConfig 删除配置
func (s *config) DeleteConfig(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.AgentConfig{}, "id = ?", ID).Error
	return err
}

// UpdateConfig 更新配置
func (s *config) UpdateConfig(config model.AgentConfig) (err error) {
	err = global.GVA_DB.Model(&model.AgentConfig{}).Where("id = ?", config.ID).Updates(&config).Error
	return err
}

// GetConfig 根据ID获取配置
func (s *config) GetConfig(ID string) (config model.AgentConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&config).Error
	return
}

// GetConfigList 获取所有配置列表
func (s *config) GetConfigList() (list []model.AgentConfig, err error) {
	err = global.GVA_DB.Find(&list).Error
	return
}

// GetActiveConfig 获取激活的配置
func (s *config) GetActiveConfig() (config model.AgentConfig, err error) {
	err = global.GVA_DB.Where("is_active = ?", true).First(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config, errors.New("未找到激活的AI配置，请先配置AI Agent")
		}
		return
	}
	return
}

// SetConfigActive 设置配置为激活状态（同时将其他配置设为非激活）
func (s *config) SetConfigActive(ID string) (err error) {
	// 先将所有配置设为非激活
	err = global.GVA_DB.Model(&model.AgentConfig{}).Where("id != ?", ID).Update("is_active", false).Error
	if err != nil {
		return err
	}
	// 激活指定配置
	return global.GVA_DB.Model(&model.AgentConfig{}).Where("id = ?", ID).Update("is_active", true).Error
}
