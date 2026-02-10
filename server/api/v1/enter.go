package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/imageregistry"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pcdn"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/product"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	ImageregistryApiGroup imageregistry.ApiGroup
	ComputenodeApiGroup   computenode.ApiGroup
	ProductApiGroup       product.ApiGroup
	InstanceApiGroup      instance.ApiGroup
	PcdnApiGroup          pcdn.ApiGroup
}
