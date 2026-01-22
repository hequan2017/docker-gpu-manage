package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model"
	"go.uber.org/zap"
)

// InitDB 初始化数据库表
func InitDB() {
	err := global.GVA_DB.AutoMigrate(
		&model.DellAsset{},
	)
	if err != nil {
		global.GVA_LOG.Error("dellAsset plugin init table failed", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("dellAsset plugin init table success")
}
