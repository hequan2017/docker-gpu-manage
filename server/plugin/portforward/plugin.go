package portforward

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	global.GVA_LOG.Info("[端口转发插件] 开始注册...")
	ctx := context.Background()

	// 如果需要配置文件，请到config.Config中填充配置结构，且到下方方法中填入其在config.yaml中的key
	// initialize.Viper()

	global.GVA_LOG.Info("[端口转发插件] 正在注册API...")
	// 安装插件时候自动注册的api数据请到下方法.Api方法中实现
	initialize.Api(ctx)

	global.GVA_LOG.Info("[端口转发插件] 正在注册菜单...")
	// 安装插件时候自动注册的菜单数据请到下方法.Menu方法中实现
	initialize.Menu(ctx)

	global.GVA_LOG.Info("[端口转发插件] 正在创建数据库表...")
	initialize.Gorm(ctx)

	global.GVA_LOG.Info("[端口转发插件] 正在注册路由...")
	initialize.Router(group)

	global.GVA_LOG.Info("[端口转发插件] 注册完成！")
}
