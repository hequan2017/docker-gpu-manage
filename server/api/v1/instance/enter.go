package instance

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ InstanceApi }

var instanceService = service.ServiceGroupApp.InstanceServiceGroup.InstanceService
