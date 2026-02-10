package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type PcdnNodeSearch struct {
	NodeID       *string `json:"nodeId" form:"nodeId"`
	Region       *string `json:"region" form:"region"`
	ISP          *string `json:"isp" form:"isp"`
	OnlineStatus *string `json:"onlineStatus" form:"onlineStatus"`
	request.PageInfo
}

type PcdnResourceSearch struct {
	NodeID *string `json:"nodeId" form:"nodeId"`
	request.PageInfo
}

type PcdnPolicySearch struct {
	PolicyName   *string `json:"policyName" form:"policyName"`
	StrategyType *string `json:"strategyType" form:"strategyType"`
	IsEnabled    *bool   `json:"isEnabled" form:"isEnabled"`
	request.PageInfo
}

type PcdnDispatchTaskSearch struct {
	TaskID   *string `json:"taskId" form:"taskId"`
	TraceID  *string `json:"traceId" form:"traceId"`
	Status   *string `json:"status" form:"status"`
	PolicyID *uint   `json:"policyId" form:"policyId"`
	request.PageInfo
}

type PcdnMetricSnapshotSearch struct {
	NodeID      *string `json:"nodeId" form:"nodeId"`
	StartAtUnix *int64  `json:"startAtUnix" form:"startAtUnix"`
	EndAtUnix   *int64  `json:"endAtUnix" form:"endAtUnix"`
	request.PageInfo
}
