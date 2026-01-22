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
			Path:      "finetuning",
			Name:      "finetuning",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      7,
			Meta:      model.Meta{Title: "算法微调", Icon: "cpu"},
		},
		{
			// ParentId 会被 RegisterMenus 自动设置为父菜单ID
			Path:      "taskList",
			Name:      "finetuningTaskList",
			Hidden:    false,
			Component: "plugin/finetuning/view/taskList.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "微调任务", Icon: "list"},
		},
		{
			// ParentId 会被 RegisterMenus 自动设置为父菜单ID
			Path:      "taskDetail",
			Name:      "finetuningTaskDetail",
			Hidden:    true,
			Component: "plugin/finetuning/view/taskDetail.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "任务详情", Icon: "document"},
		},
	}

	// RegisterMenus 会自动将第一个菜单作为父菜单，其余菜单作为子菜单
	utils.RegisterMenus(entities...)
}
