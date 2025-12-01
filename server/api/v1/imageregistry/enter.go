package imageregistry

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ImageRegistryApi }

var imageRegistryService = service.ServiceGroupApp.ImageregistryServiceGroup.ImageRegistryService
