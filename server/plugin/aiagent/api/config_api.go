package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Config = new(config)

type config struct{}

// CreateConfig 创建配置
// @Tags Config
// @Summary 创建配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AgentConfig true "创建配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /config/createConfig [post]
func (a *config) CreateConfig(c *gin.Context) {
	var config model.AgentConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceConfig.CreateConfig(&config)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteConfig 删除配置
// @Tags Config
// @Summary 删除配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "配置ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /config/deleteConfig [delete]
func (a *config) DeleteConfig(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceConfig.DeleteConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateConfig 更新配置
// @Tags Config
// @Summary 更新配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AgentConfig true "更新配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /config/updateConfig [put]
func (a *config) UpdateConfig(c *gin.Context) {
	var config model.AgentConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// API Key 不允许修改为空
	if config.APIKey == "" {
		response.FailWithMessage("API Key不能为空", c)
		return
	}
	err = serviceConfig.UpdateConfig(config)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindConfig 用id查询配置
// @Tags Config
// @Summary 用id查询配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "配置ID"
// @Success 200 {object} response.Response{data=model.AgentConfig,msg=string} "查询成功"
// @Router /config/findConfig [get]
func (a *config) FindConfig(c *gin.Context) {
	ID := c.Query("ID")
	config, err := serviceConfig.GetConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	// 隐藏API Key（只显示前8位）
	if len(config.APIKey) > 8 {
		config.APIKey = config.APIKey[:8] + "****"
	}
	response.OkWithData(config, c)
}

// GetConfigList 获取所有配置列表
// @Tags Config
// @Summary 获取所有配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]model.AgentConfig,msg=string} "获取成功"
// @Router /config/getConfigList [get]
func (a *config) GetConfigList(c *gin.Context) {
	list, err := serviceConfig.GetConfigList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 隐藏API Key
	for i := range list {
		if len(list[i].APIKey) > 8 {
			list[i].APIKey = list[i].APIKey[:8] + "****"
		}
	}
	response.OkWithData(list, c)
}

// SetConfigActive 设置配置为激活状态
// @Tags Config
// @Summary 设置配置为激活状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "配置ID"
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /config/setActive [post]
func (a *config) SetConfigActive(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceConfig.SetConfigActive(ID)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetActiveConfig 获取激活的配置
// @Tags Config
// @Summary 获取激活的配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=model.AgentConfig,msg=string} "获取成功"
// @Router /config/getActive [get]
func (a *config) GetActiveConfig(c *gin.Context) {
	config, err := serviceConfig.GetActiveConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 隐藏API Key
	if len(config.APIKey) > 8 {
		config.APIKey = config.APIKey[:8] + "****"
	}
	response.OkWithData(config, c)
}
