package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/openclaw/api"
	"github.com/gin-gonic/gin"
)

func InitOpenClawRouter(Router *gin.RouterGroup) {
	openClawRouter := Router.Group("openclaw")
	{
		openClawRouter.GET("health", api.Api.HealthCheck)
	}
}
