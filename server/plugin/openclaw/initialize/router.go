package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/openclaw/router"
	"github.com/gin-gonic/gin"
)

func Router(group *gin.Engine) {
	publicGroup := group.Group("")
	router.InitOpenClawRouter(publicGroup)
}
