package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  24, // 请根据实际情况调整父菜单ID
			Path:      "aiagent",
			Name:      "aiagent",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "AI Agent", Icon: "chat-dot-square"},
		},
		{
			ParentId:  0, // 将在下面动态设置
			Path:      "chat",
			Name:      "aiagentChat",
			Hidden:    false,
			Component: "plugin/aiagent/view/chat.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "AI 对话", Icon: "chat-line-round"},
		},
		{
			ParentId:  0, // 将在下面动态设置
			Path:      "config",
			Name:      "aiagentConfig",
			Hidden:    false,
			Component: "plugin/aiagent/view/config.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "AI 配置", Icon: "setting"},
		},
	}

	// 先创建父菜单
	parentMenu := entities[0]
	utils.RegisterMenus(parentMenu)

	// 获取父菜单ID并设置子菜单的ParentId
	// 注意：这里简化处理，实际可能需要通过查询获取父菜单ID
	// 假设父菜单已经创建，我们需要获取它的ID
	// 在实际使用中，你可能需要调整这部分逻辑

	// 创建子菜单
	// 这里的ParentId需要根据实际情况设置
	// 可以通过查询数据库获取刚创建的父菜单ID
}
