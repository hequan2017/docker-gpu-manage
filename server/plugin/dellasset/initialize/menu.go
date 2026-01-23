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
			Path:      "dellAsset",
			Name:      "dellAsset",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      8,
			Meta:      model.Meta{Title: "戴尔资产管理", Icon: "cpu"},
		},
		{
			// ParentId 会被 RegisterMenus 自动设置为父菜单ID
			Path:      "assetList",
			Name:      "dellAssetList",
			Hidden:    false,
			Component: "plugin/dellasset/view/dellAsset.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "资产列表", Icon: "list"},
		},
	}

	// RegisterMenus 会自动将第一个菜单作为父菜单，其余菜单作为子菜单
	utils.RegisterMenus(entities...)
}
