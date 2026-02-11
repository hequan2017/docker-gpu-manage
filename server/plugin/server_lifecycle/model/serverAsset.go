
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ServerAsset 服务器资产 结构体
type ServerAsset struct {
    global.GVA_MODEL
  HostName  *string `json:"hostName" form:"hostName" gorm:"comment:主机名;column:host_name;size:100;" binding:"required"`  //主机名
  IP  *string `json:"ip" form:"ip" gorm:"comment:IP地址;column:ip;size:50;" binding:"required"`  //IP地址
  Sn  *string `json:"sn" form:"sn" gorm:"comment:SN号;column:sn;size:100;"`  //SN号
  Configuration  *string `json:"configuration" form:"configuration" gorm:"comment:CPU/内存/磁盘等配置信息;column:configuration;size:255;"`  //硬件配置
  Status  string `json:"status" form:"status" gorm:"default:deploy_pending;comment:待部署/运行中/已下机/已报废;column:status;type:varchar(50);" binding:"required"`  //服务器状态
  ServiceType  *string `json:"serviceType" form:"serviceType" gorm:"comment:部署的服务类型，如Nginx;column:service_type;size:100;"`  //部署服务
  DeployTime  *time.Time `json:"deployTime" form:"deployTime" gorm:"comment:服务部署时间;column:deploy_time;"`  //部署时间
  OfflineTime  *time.Time `json:"offlineTime" form:"offlineTime" gorm:"comment:服务器下机时间;column:offline_time;"`  //下机时间
  ScrapTime  *time.Time `json:"scrapTime" form:"scrapTime" gorm:"comment:服务器报废时间;column:scrap_time;"`  //报废时间
  ScrapReason  *string `json:"scrapReason" form:"scrapReason" gorm:"comment:服务器报废原因;column:scrap_reason;size:255;"`  //报废原因
}


// TableName 服务器资产 ServerAsset自定义表名 sl_server_assets
func (ServerAsset) TableName() string {
    return "sl_server_assets"
}







