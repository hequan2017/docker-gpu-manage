package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/service"
)

type ApiGroup struct {
	DellAssetApi
}

var ApiGroupApp = new(ApiGroup)

// 初始化 Service 引用
var (
	serviceDellAsset = service.ServiceGroupApp.DellAssetService
)
