package openclaw

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/openclaw/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	initialize.Router(group)
}

func (p *plugin) RouterPath() string {
	return "openclaw"
}
