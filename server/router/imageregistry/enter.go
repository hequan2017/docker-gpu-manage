package imageregistry

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ImageRegistryRouter }

var imageRegistryApi = api.ApiGroupApp.ImageregistryApiGroup.ImageRegistryApi
