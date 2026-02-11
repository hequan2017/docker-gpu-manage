package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// PcdnNodeSearch 节点搜索结构体
type PcdnNodeSearch struct {
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
	IP     string `json:"ip" form:"ip"`
	request.PageInfo
}
