package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/approval_flow/api"

var (
	Router             = new(router)
	apiApprovalProcess = api.Api.ApprovalProcess
)

type router struct{ ApprovalProcess approval }
