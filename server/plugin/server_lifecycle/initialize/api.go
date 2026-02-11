package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/asset/createServerAsset", Description: "新建服务器资产", ApiGroup: "server_lifecycle", Method: "POST"},
		{Path: "/asset/deleteServerAsset", Description: "删除服务器资产", ApiGroup: "server_lifecycle", Method: "DELETE"},
		{Path: "/asset/deleteServerAssetByIds", Description: "批量删除服务器资产", ApiGroup: "server_lifecycle", Method: "DELETE"},
		{Path: "/asset/updateServerAsset", Description: "更新服务器资产", ApiGroup: "server_lifecycle", Method: "PUT"},
		{Path: "/asset/findServerAsset", Description: "根据ID获取服务器资产", ApiGroup: "server_lifecycle", Method: "GET"},
		{Path: "/asset/getServerAssetList", Description: "获取服务器资产列表", ApiGroup: "server_lifecycle", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
