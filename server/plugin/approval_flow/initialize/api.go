package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/approval/createApprovalProcess", Description: "新建发版申请", ApiGroup: "approval_flow", Method: "POST"},
		{Path: "/approval/deleteApprovalProcess", Description: "删除发版申请", ApiGroup: "approval_flow", Method: "DELETE"},
		{Path: "/approval/deleteApprovalProcessByIds", Description: "批量删除发版申请", ApiGroup: "approval_flow", Method: "DELETE"},
		{Path: "/approval/updateApprovalProcess", Description: "更新发版申请", ApiGroup: "approval_flow", Method: "PUT"},
		{Path: "/approval/findApprovalProcess", Description: "根据ID获取发版申请", ApiGroup: "approval_flow", Method: "GET"},
		{Path: "/approval/getApprovalProcessList", Description: "获取发版申请列表", ApiGroup: "approval_flow", Method: "GET"},
		{Path: "/approval/approveRequest", Description: "批准发版申请", ApiGroup: "approval_flow", Method: "PUT"},
		{Path: "/approval/rejectRequest", Description: "驳回发版申请", ApiGroup: "approval_flow", Method: "PUT"},
	}
	utils.RegisterApis(entities...)
}
