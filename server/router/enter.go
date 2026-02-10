package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/imageregistry"
	"github.com/flipped-aurora/gin-vue-admin/server/router/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pcdn"
	"github.com/flipped-aurora/gin-vue-admin/server/router/product"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Imageregistry imageregistry.RouterGroup
	Computenode   computenode.RouterGroup
	Product       product.RouterGroup
	Instance      instance.RouterGroup
	Pcdn          pcdn.RouterGroup
}
