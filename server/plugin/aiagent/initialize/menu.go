package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0, // 0 表示顶级菜单
			Path:      "aiagent",
			Name:      "aiagent",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "AI Agent", Icon: "chat-dot-square"},
		},
		{
			// ParentId 会被 RegisterMenus 自动设置为父菜单ID
			Path:      "chat",
			Name:      "aiagentChat",
			Hidden:    false,
			Component: "plugin/aiagent/view/chat.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "AI 对话", Icon: "chat-line-round"},
		},
		{
			// ParentId 会被 RegisterMenus 自动设置为父菜单ID
			Path:      "config",
			Name:      "aiagentConfig",
			Hidden:    false,
			Component: "plugin/aiagent/view/config.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "AI 配置", Icon: "setting"},
		},
	}

	// RegisterMenus 会自动将第一个菜单作为父菜单，其余菜单作为子菜单
	utils.RegisterMenus(entities...)
}
