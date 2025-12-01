package computenode

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ComputeNodeRouter }

var computeNodeApi = api.ApiGroupApp.ComputenodeApiGroup.ComputeNodeApi
