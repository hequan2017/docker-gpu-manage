
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ServerAssetSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       HostName  *string `json:"hostName" form:"hostName"` 
       IP  *string `json:"ip" form:"ip"` 
       Sn  *string `json:"sn" form:"sn"` 
       Configuration  *string `json:"configuration" form:"configuration"` 
       Status  string `json:"status" form:"status"` 
       ServiceType  *string `json:"serviceType" form:"serviceType"` 
       DeployTimeRange  []time.Time  `json:"deployTimeRange" form:"deployTimeRange[]"`
       OfflineTimeRange  []time.Time  `json:"offlineTimeRange" form:"offlineTimeRange[]"`
       ScrapTimeRange  []time.Time  `json:"scrapTimeRange" form:"scrapTimeRange[]"`
       ScrapReason  *string `json:"scrapReason" form:"scrapReason"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
