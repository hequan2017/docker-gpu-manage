package product

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ProductSpecApi }

var productSpecService = service.ServiceGroupApp.ProductServiceGroup.ProductSpecService
