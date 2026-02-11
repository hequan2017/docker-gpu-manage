package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			Path:      "approval_flow",
			Name:      "approval_flow",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      12,
			Meta: model.Meta{
				Title: "审批流管理",
				Icon:  "stamp",
			},
		},
		{
			Path:      "approvalProcess",
			Name:      "approvalProcess",
			Hidden:    false,
			Component: "plugin/approval_flow/view/approvalprocess.vue",
			Sort:      1,
			Meta: model.Meta{
				Title: "发版申请",
				Icon:  "document",
			},
		},
	}
	utils.RegisterMenus(entities...)
}
