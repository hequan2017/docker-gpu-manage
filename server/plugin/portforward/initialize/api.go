package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/portForward/createPortForward",
			Description: "新建端口转发规则",
			ApiGroup:    "端口转发",
			Method:      "POST",
		},
		{
			Path:        "/portForward/deletePortForward",
			Description: "删除端口转发规则",
			ApiGroup:    "端口转发",
			Method:      "DELETE",
		},
		{
			Path:        "/portForward/deletePortForwardByIds",
			Description: "批量删除端口转发规则",
			ApiGroup:    "端口转发",
			Method:      "DELETE",
		},
		{
			Path:        "/portForward/updatePortForward",
			Description: "更新端口转发规则",
			ApiGroup:    "端口转发",
			Method:      "PUT",
		},
		{
			Path:        "/portForward/updatePortForwardStatus",
			Description: "更新端口转发规则状态",
			ApiGroup:    "端口转发",
			Method:      "PUT",
		},
		{
			Path:        "/portForward/findPortForward",
			Description: "根据ID获取端口转发规则",
			ApiGroup:    "端口转发",
			Method:      "GET",
		},
		{
			Path:        "/portForward/getPortForwardList",
			Description: "获取端口转发规则列表",
			ApiGroup:    "端口转发",
			Method:      "GET",
		},
		{
			Path:        "/portForward/getServerIP",
			Description: "获取服务器IP地址",
			ApiGroup:    "端口转发",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
