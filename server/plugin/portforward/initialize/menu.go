package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		// 父菜单：端口转发管理（顶级菜单）
		{
			ParentId:  0, // 0 表示顶级菜单
			Path:      "/portForward",
			Name:      "portForward",
			Hidden:    false,
			Component: "view/routerHolder.vue", // 路由容器组件
			Sort:      10,
			Meta:      model.Meta{Title: "端口转发", Icon: "position"},
		},
		// 子菜单：转发规则管理（会自动设置为第一个菜单的子菜单）
		{
			Path:      "portForwardRules",
			Name:      "portForwardRules",
			Hidden:    false,
			Component: "plugin/portforward/view/portForward.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "转发规则", Icon: "list"},
		},
	}
	// RegisterMenus 会自动将第一个菜单作为父菜单，其余菜单作为子菜单
	utils.RegisterMenus(entities...)
}
