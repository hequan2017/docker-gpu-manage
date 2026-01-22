package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/service"

var (
	Api                = new(api)
	servicePortForward = service.Service.PortForward
)

type api struct{ PortForward portForward }
