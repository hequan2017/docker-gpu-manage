package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/approvalProcess/createApprovalProcess", Description: "新建发版申请", ApiGroup: "approval_flow", Method: "POST"},
		{Path: "/approvalProcess/deleteApprovalProcess", Description: "删除发版申请", ApiGroup: "approval_flow", Method: "DELETE"},
		{Path: "/approvalProcess/deleteApprovalProcessByIds", Description: "批量删除发版申请", ApiGroup: "approval_flow", Method: "DELETE"},
		{Path: "/approvalProcess/updateApprovalProcess", Description: "更新发版申请", ApiGroup: "approval_flow", Method: "PUT"},
		{Path: "/approvalProcess/findApprovalProcess", Description: "根据ID获取发版申请", ApiGroup: "approval_flow", Method: "GET"},
		{Path: "/approvalProcess/getApprovalProcessList", Description: "获取发版申请列表", ApiGroup: "approval_flow", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
