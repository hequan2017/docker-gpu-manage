package dellasset

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/router"
)

type DellAssetPlugin struct{}

// Register 注册路由
func (p *DellAssetPlugin) Register(RouterGroup interface{}) {
	// 路由注册在 initialize/router.go 中处理
}

// RouterPath 获取路由路径
func (p *DellAssetPlugin) RouterPath() string {
	return "dellAsset"
}

func init() {
	// 数据库迁移
	global.GVA_DB.AutoMigrate(&model.DellAsset{})

	// 路由注册会在系统初始化时通过 initialize/router.go 处理
}
