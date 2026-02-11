package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.AutoMigrate(
		new(model.PcdnNode),
	)
	if err != nil {
		global.GVA_LOG.Error("PCDN plugin auto migrate failed", zap.Error(err))
	}
}

func Api(ctx context.Context) {
	utils.RegisterApis(
		model2.SysApi{
			Path:        "/pcdn/createPcdnNode",
			Description: "创建PCDN节点",
			ApiGroup:    "PCDN",
			Method:      "POST",
		},
		model2.SysApi{
			Path:        "/pcdn/deletePcdnNode",
			Description: "删除PCDN节点",
			ApiGroup:    "PCDN",
			Method:      "DELETE",
		},
		model2.SysApi{
			Path:        "/pcdn/deletePcdnNodeByIds",
			Description: "批量删除PCDN节点",
			ApiGroup:    "PCDN",
			Method:      "DELETE",
		},
		model2.SysApi{
			Path:        "/pcdn/updatePcdnNode",
			Description: "更新PCDN节点",
			ApiGroup:    "PCDN",
			Method:      "PUT",
		},
		model2.SysApi{
			Path:        "/pcdn/findPcdnNode",
			Description: "根据ID获取PCDN节点",
			ApiGroup:    "PCDN",
			Method:      "GET",
		},
		model2.SysApi{
			Path:        "/pcdn/getPcdnNodeList",
			Description: "获取PCDN节点列表",
			ApiGroup:    "PCDN",
			Method:      "GET",
		},
	)
}

func Menu(ctx context.Context) {
	utils.RegisterMenus(
		model2.SysBaseMenu{
			Path:      "pcdn",
			Name:      "pcdn",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      9,
			Meta: model2.Meta{
				Title: "PCDN管理",
				Icon:  "monitor",
			},
		},
		model2.SysBaseMenu{
			Path:      "pcdnNode",
			Name:      "pcdnNode",
			Hidden:    false,
			Component: "plugin/pcdn/view/node.vue",
			Sort:      1,
			Meta: model2.Meta{
				Title: "节点管理",
				Icon:  "menu",
			},
		},
	)
}
