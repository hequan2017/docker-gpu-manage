package pcdn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/router"
	"github.com/gin-gonic/gin"
)

type PcdnPlugin struct{}

var Plugin = new(PcdnPlugin)

func CreatePcdnPlug() *PcdnPlugin {
	return &PcdnPlugin{}
}

func (*PcdnPlugin) Register(group *gin.Engine) {
	router.Router.InitPcdnNodeRouter(group.Group(""))
}

func (*PcdnPlugin) RouterPath() string {
	return "pcdn"
}

func (*PcdnPlugin) Init() {
	initialize.Gorm(nil)
	initialize.Api(nil)
	initialize.Menu(nil)
}
