package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/api"
)

type RouterGroup struct {
	portForward
}

var RouterGroupApp = new(RouterGroup)

// 初始化 API 引用
var (
	apiPortForward = api.Api.PortForward
)
