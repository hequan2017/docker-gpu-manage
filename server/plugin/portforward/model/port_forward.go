package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PortForward 端口转发 结构体
type PortForward struct {
	global.GVA_MODEL
	SourceIP      string `json:"sourceIP" form:"sourceIP" gorm:"column:source_ip;comment:源IP地址;type:varchar(64);not null"`                      // 源IP地址
	SourcePort    int    `json:"sourcePort" form:"sourcePort" gorm:"column:source_port;comment:源端口;not null"`                                 // 源端口
	Protocol      string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:协议类型;type:varchar(10);not null;default:tcp"`          // 协议类型: tcp/udp
	TargetIP      string `json:"targetIP" form:"targetIP" gorm:"column:target_ip;comment:目标IP地址;type:varchar(64);not null"`                    // 目标IP地址
	TargetPort    int    `json:"targetPort" form:"targetPort" gorm:"column:target_port;comment:目标端口;not null"`                               // 目标端口
	Status        bool   `json:"status" form:"status" gorm:"column:status;comment:状态;default:true"`                                           // 状态: true-启用, false-禁用
	Description   string `json:"description" form:"description" gorm:"column:description;comment:规则描述;type:varchar(255)"`                    // 规则描述
}

// TableName PortForward 自定义表名 gva_port_forward
func (PortForward) TableName() string {
	return "gva_port_forward"
}
