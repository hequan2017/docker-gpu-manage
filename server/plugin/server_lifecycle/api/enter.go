package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/service"

var (
	Api                = new(api)
	serviceServerAsset = service.Service.ServerAsset
)

type api struct{ ServerAsset asset }
