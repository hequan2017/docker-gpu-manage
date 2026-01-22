package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/router"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.PortForward.Init(public, private)

	// 同步所有端口转发规则状态
	go func() {
		if err := service.PortForward.SyncAllPortForwards(); err != nil {
			global.GVA_LOG.Error("同步端口转发规则失败", zap.Error(err))
		} else {
			global.GVA_LOG.Info("端口转发规则同步完成")
		}
	}()
}
