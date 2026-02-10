package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PcdnMetricsApi }

var (
	pcdnMetricsService = service.ServiceGroupApp.PcdnServiceGroup.MetricsService
	healthScoreService = service.ServiceGroupApp.PcdnServiceGroup.HealthScoreService
)
