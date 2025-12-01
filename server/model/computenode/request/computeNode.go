
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ComputeNodeSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      Region  *string `json:"region" form:"region"` 
      PublicIp  *string `json:"publicIp" form:"publicIp"` 
      PrivateIp  *string `json:"privateIp" form:"privateIp"` 
      GpuName  *string `json:"gpuName" form:"gpuName"` 
      IsOnShelf  *bool `json:"isOnShelf" form:"isOnShelf"` 
    request.PageInfo
}
