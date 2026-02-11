package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/approval_flow/service"

var (
	Api                    = new(api)
	serviceApprovalProcess = service.Service.ApprovalProcess
)

type api struct{ ApprovalProcess approval }
