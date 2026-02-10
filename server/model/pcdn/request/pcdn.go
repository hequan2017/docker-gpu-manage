package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PcdnNodeSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           string      `json:"name" form:"name"`
	Region         string      `json:"region" form:"region"`
	Online         *bool       `json:"online" form:"online"`
	request.PageInfo
}

type PcdnNodeOnlineRequest struct {
	ID     uint `json:"id" binding:"required"`
	Online bool `json:"online"`
}

type PcdnNodeWeightRequest struct {
	ID     uint `json:"id" binding:"required"`
	Weight int  `json:"weight" binding:"required"`
}

type PcdnPolicySearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           string      `json:"name" form:"name"`
	Enabled        *bool       `json:"enabled" form:"enabled"`
	Published      *bool       `json:"published" form:"published"`
	request.PageInfo
}

type PcdnPolicyGrayReleaseRequest struct {
	ID          uint   `json:"id" binding:"required"`
	GrayPercent int    `json:"grayPercent" binding:"required,gte=0,lte=100"`
	Note        string `json:"note"`
}

type PcdnPolicySwitchRequest struct {
	ID      uint `json:"id" binding:"required"`
	Enabled bool `json:"enabled"`
}

type PcdnDispatchTaskSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	TraceID        string      `json:"traceId" form:"traceId"`
	Status         string      `json:"status" form:"status"`
	request.PageInfo
}
