package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PcdnNodeApi
	PcdnPolicyApi
	PcdnDispatchTaskApi
}

var (
	pcdnNodeService         = service.ServiceGroupApp.PcdnServiceGroup.PcdnNodeService
	pcdnPolicyService       = service.ServiceGroupApp.PcdnServiceGroup.PcdnPolicyService
	pcdnDispatchTaskService = service.ServiceGroupApp.PcdnServiceGroup.PcdnDispatchTaskService
)
