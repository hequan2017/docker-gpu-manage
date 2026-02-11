package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ApprovalProcess = new(approval)

type approval struct{}

// Init 初始化 发版申请 路由信息
func (r *approval) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("approval").Use(middleware.OperationRecord())
		group.POST("createApprovalProcess", apiApprovalProcess.CreateApprovalProcess)             // 新建发版申请
		group.DELETE("deleteApprovalProcess", apiApprovalProcess.DeleteApprovalProcess)           // 删除发版申请
		group.DELETE("deleteApprovalProcessByIds", apiApprovalProcess.DeleteApprovalProcessByIds) // 批量删除发版申请
		group.PUT("updateApprovalProcess", apiApprovalProcess.UpdateApprovalProcess)              // 更新发版申请
		group.PUT("approveRequest", apiApprovalProcess.ApproveRequest)                            // 批准发版申请
		group.PUT("rejectRequest", apiApprovalProcess.RejectRequest)                              // 驳回发版申请
	}
	{
		group := private.Group("approval")
		group.GET("findApprovalProcess", apiApprovalProcess.FindApprovalProcess)       // 根据ID获取发版申请
		group.GET("getApprovalProcessList", apiApprovalProcess.GetApprovalProcessList) // 获取发版申请列表
	}
	{
		group := public.Group("approval")
		group.GET("getApprovalProcessPublic", apiApprovalProcess.GetApprovalProcessPublic) // 发版申请开放接口
	}
}
