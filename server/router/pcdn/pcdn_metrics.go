package pcdn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PcdnMetricsRouter struct{}

func (r *PcdnMetricsRouter) InitPcdnMetricsRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	pcdnPrivate := privateRouter.Group("pcdnMetrics")
	pcdnPublic := publicRouter.Group("pcdnMetrics").Use(middleware.PcdnNodeAuth())
	{
		pcdnPublic.POST("report", pcdnMetricsApi.ReportMetrics)
	}
	{
		pcdnPrivate.GET("latest", pcdnMetricsApi.GetLatestSnapshot)
	}
}
