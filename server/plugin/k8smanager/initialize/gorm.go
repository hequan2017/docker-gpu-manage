package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
)

func Gorm(ctx context.Context) {
	// 检查并删除过时的 api_server 字段
	if global.GVA_DB.Migrator().HasColumn(&model.K8sCluster{}, "api_server") {
		err := global.GVA_DB.Migrator().DropColumn(&model.K8sCluster{}, "api_server")
		if err != nil {
			global.GVA_LOG.Error("Failed to drop api_server column", zap.Error(err))
		} else {
			global.GVA_LOG.Info("Dropped obsolete column: api_server")
		}
	}

	err := global.GVA_DB.AutoMigrate(
		&model.K8sCluster{},
		&model.K8sAuditLog{},
		&model.K8sPermission{},
		&model.K8sRolePermission{},
		&model.K8sUserPermission{},
	)
	if err != nil {
		global.GVA_LOG.Error("k8smanager plugin auto migrate failed", zap.String("error", err.Error()))
		panic(err)
	}
}
