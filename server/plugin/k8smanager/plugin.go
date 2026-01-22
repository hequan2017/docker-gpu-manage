package k8smanager

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()

	// 初始化数据库表
	initialize.Gorm(ctx)

	// 注册API权限
	initialize.Api(ctx)

	// 注册菜单
	initialize.Menu(ctx)

	// 注册路由
	initialize.Router(group)
}

func (p *plugin) RouterPath() string {
	return "/k8s"
}
