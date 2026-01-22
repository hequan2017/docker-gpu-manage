package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0, // 可根据需要调整为父级菜单ID
			Path:      "finetuning",
			Name:      "finetuning",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      10,
			Meta:      model.Meta{Title: "算法微调", Icon: "cpu"},
		},
		{
			ParentId:  0, // 这里需要根据实际创建的父级菜单ID进行调整
			Path:      "taskList",
			Name:      "finetuningTaskList",
			Hidden:    false,
			Component: "plugin/finetuning/view/taskList.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "微调任务", Icon: "list"},
		},
		{
			ParentId:  0, // 这里需要根据实际创建的父级菜单ID进行调整
			Path:      "taskDetail",
			Name:      "finetuningTaskDetail",
			Hidden:    true,
			Component: "plugin/finetuning/view/taskDetail.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "任务详情", Icon: "document"},
		},
	}
	utils.RegisterMenus(entities...)
}
