package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			Path:      "llm_model",
			Name:      "llm_model",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      11,
			Meta: model.Meta{
				Title: "大模型管理",
				Icon:  "box",
			},
		},
		{
			Path:      "llmModel",
			Name:      "llmModel",
			Hidden:    false,
			Component: "plugin/llm_model/view/llm_model.vue",
			Sort:      1,
			Meta: model.Meta{
				Title: "模型列表",
				Icon:  "menu",
			},
		},
	}
	utils.RegisterMenus(entities...)
}
