package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// PcdnNode PCDN节点
type PcdnNode struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;size:128;comment:节点名称" binding:"required"`
	Region  string `json:"region" form:"region" gorm:"column:region;size:128;comment:节点区域"`
	Address string `json:"address" form:"address" gorm:"column:address;size:255;comment:节点地址" binding:"required"`
	Weight  int    `json:"weight" form:"weight" gorm:"column:weight;default:100;comment:节点权重"`
	Online  bool   `json:"online" form:"online" gorm:"column:online;default:true;comment:上下线状态"`
	Status  string `json:"status" form:"status" gorm:"column:status;size:32;default:'active';comment:节点状态"`
	Remark  string `json:"remark" form:"remark" gorm:"column:remark;size:512;comment:备注"`
}

func (PcdnNode) TableName() string {
	return "pcdn_node"
}
