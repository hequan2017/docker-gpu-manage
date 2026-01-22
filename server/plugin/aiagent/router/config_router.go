package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Config = new(config)

type config struct{}

// Init 初始化 配置 路由信息
func (r *config) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("config").Use(middleware.OperationRecord())
		group.POST("createConfig", apiConfig.CreateConfig)     // 新建配置
		group.DELETE("deleteConfig", apiConfig.DeleteConfig)  // 删除配置
		group.PUT("updateConfig", apiConfig.UpdateConfig)     // 更新配置
		group.POST("setActive", apiConfig.SetConfigActive)    // 设置激活状态
	}
	{
		group := private.Group("config")
		group.GET("findConfig", apiConfig.FindConfig)       // 根据ID获取配置
		group.GET("getConfigList", apiConfig.GetConfigList) // 获取配置列表
		group.GET("getActive", apiConfig.GetActiveConfig)   // 获取激活的配置
	}
}
