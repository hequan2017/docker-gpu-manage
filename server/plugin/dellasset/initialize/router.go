package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/router"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	DellAssetRouter
}

// InitRouter 初始化路由
func InitRouter(Router *gin.RouterGroup) {
	routerGroup := new(RouterGroup)
	{
		routerGroup.DellAssetRouter.InitDellAssetRouter(Router)
	}
}
