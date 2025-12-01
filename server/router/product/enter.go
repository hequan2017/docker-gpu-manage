package product

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ProductSpecRouter }

var productSpecApi = api.ApiGroupApp.ProductApiGroup.ProductSpecApi
