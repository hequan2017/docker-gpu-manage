package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// PortForwardSearch 端口转发查询条件
type PortForwardSearch struct {
	SourceIP   string `json:"sourceIP" form:"sourceIP"`     // 源IP地址
	SourcePort int    `json:"sourcePort" form:"sourcePort"` // 源端口
	Protocol   string `json:"protocol" form:"protocol"`     // 协议类型
	TargetIP   string `json:"targetIP" form:"targetIP"`     // 目标IP地址
	Status     *bool  `json:"status" form:"status"`         // 状态
	request.PageInfo
}
