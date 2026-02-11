package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/api"

var Router = new(router)

type router struct {
	PcdnNodeRouter
}

var pcdnNodeApi = api.Api.PcdnNodeApi
