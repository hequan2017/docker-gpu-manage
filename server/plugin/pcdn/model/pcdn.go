package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PcdnNode PCDN 节点 结构体
type PcdnNode struct {
	global.GVA_MODEL
	Name          string  `json:"name" form:"name" gorm:"column:name;comment:节点名称;type:varchar(100);"`
	IP            string  `json:"ip" form:"ip" gorm:"column:ip;comment:IP地址;type:varchar(50);"`
	Mac           string  `json:"mac" form:"mac" gorm:"column:mac;comment:MAC地址;type:varchar(50);"`
	OS            string  `json:"os" form:"os" gorm:"column:os;comment:操作系统;type:varchar(50);"`
	Status        string  `json:"status" form:"status" gorm:"column:status;comment:状态: online/offline/abnormal;type:varchar(20);default:offline;"`
	Bandwidth     int     `json:"bandwidth" form:"bandwidth" gorm:"column:bandwidth;comment:上行带宽(Mbps);"`
	UsedBandwidth int     `json:"usedBandwidth" form:"usedBandwidth" gorm:"column:used_bandwidth;comment:已用带宽;"`
	TotalStorage  int     `json:"totalStorage" form:"totalStorage" gorm:"column:total_storage;comment:总存储(GB);"`
	UsedStorage   int     `json:"usedStorage" form:"usedStorage" gorm:"column:used_storage;comment:已用存储;"`
	BusinessType  string  `json:"businessType" form:"businessType" gorm:"column:business_type;comment:业务类型;type:varchar(50);"`
	TodayIncome   float64 `json:"todayIncome" form:"todayIncome" gorm:"column:today_income;comment:今日收益;type:decimal(10,2);default:0;"`
	TotalIncome   float64 `json:"totalIncome" form:"totalIncome" gorm:"column:total_income;comment:总收益;type:decimal(10,2);default:0;"`
}

// TableName PcdnNode 自定义表名 gva_pcdn_nodes
func (PcdnNode) TableName() string {
	return "gva_pcdn_nodes"
}
