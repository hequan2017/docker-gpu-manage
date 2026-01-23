package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/dellAsset/createDellAsset",
			Description: "创建戴尔服务器资产",
			ApiGroup:    "戴尔资产管理",
			Method:      "POST",
		},
		{
			Path:        "/dellAsset/deleteDellAsset",
			Description: "删除戴尔服务器资产",
			ApiGroup:    "戴尔资产管理",
			Method:      "DELETE",
		},
		{
			Path:        "/dellAsset/deleteDellAssetByIds",
			Description: "批量删除戴尔服务器资产",
			ApiGroup:    "戴尔资产管理",
			Method:      "DELETE",
		},
		{
			Path:        "/dellAsset/updateDellAsset",
			Description: "更新戴尔服务器资产",
			ApiGroup:    "戴尔资产管理",
			Method:      "PUT",
		},
		{
			Path:        "/dellAsset/findDellAsset",
			Description: "根据ID获取戴尔服务器资产",
			ApiGroup:    "戴尔资产管理",
			Method:      "GET",
		},
		{
			Path:        "/dellAsset/getDellAssetList",
			Description: "获取戴尔服务器资产列表",
			ApiGroup:    "戴尔资产管理",
			Method:      "GET",
		},
		{
			Path:        "/dellAsset/setDellAssetImport",
			Description: "设置戴尔服务器资产导入状态",
			ApiGroup:    "戴尔资产管理",
			Method:      "POST",
		},
	}
	utils.RegisterApis(entities...)
}
