package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.AutoMigrate(
		&model.K8sCluster{},
	)
	if err != nil {
		global.GVA_LOG.Error("k8smanager plugin auto migrate failed", zap.String("error", err.Error()))
		panic(err)
	}
}
