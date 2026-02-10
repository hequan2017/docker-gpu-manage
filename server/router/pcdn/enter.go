package pcdn

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PcdnRouter
}

var pcdnApi = api.ApiGroupApp.PcdnApiGroup
