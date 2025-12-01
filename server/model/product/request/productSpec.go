
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductSpecSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      GpuModel  *string `json:"gpuModel" form:"gpuModel"` 
      IsOnShelf  *bool `json:"isOnShelf" form:"isOnShelf"` 
    request.PageInfo
}
