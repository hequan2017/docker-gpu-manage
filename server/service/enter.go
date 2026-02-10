package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/imageregistry"
	"github.com/flipped-aurora/gin-vue-admin/server/service/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pcdn"
	"github.com/flipped-aurora/gin-vue-admin/server/service/product"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup        system.ServiceGroup
	ExampleServiceGroup       example.ServiceGroup
	ImageregistryServiceGroup imageregistry.ServiceGroup
	ComputenodeServiceGroup   computenode.ServiceGroup
	ProductServiceGroup       product.ServiceGroup
	InstanceServiceGroup      instance.ServiceGroup
	PcdnServiceGroup          pcdn.ServiceGroup
}
