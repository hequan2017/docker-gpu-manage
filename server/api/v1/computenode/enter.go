package computenode

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ComputeNodeApi }

var computeNodeService = service.ServiceGroupApp.ComputenodeServiceGroup.ComputeNodeService
