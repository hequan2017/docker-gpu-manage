package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// DellAssetSearch 戴尔服务器资产搜索结构体
type DellAssetSearch struct {
	request.PageInfo
	HostName     string `form:"hostName"`     // 主机名
	ServiceTag   string `form:"serviceTag"`   // 服务标签
	AssetNumber  string `form:"assetNumber"`  // 资产编号
	Model        string `form:"model"`        // 型号
	IPAddress    string `form:"ipAddress"`    // IP地址
	Cabinet      string `form:"cabinet"`      // 机柜位置
	Department   string `form:"department"`   // 所属部门
	Manager      string `form:"manager"`      // 负责人
	PowerStatus  string `form:"powerStatus"`  // 电源状态
	Status       string `form:"status"`       // 状态
	OS           string `form:"os"`           // 操作系统
}
