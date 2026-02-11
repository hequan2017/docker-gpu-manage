package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/api"

var (
	Router         = new(router)
	apiServerAsset = api.Api.ServerAsset
)

type router struct{ ServerAsset asset }
