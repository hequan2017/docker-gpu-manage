
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ImageRegistrySearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      Address  *string `json:"address" form:"address"` 
      Source  *string `json:"source" form:"source"` 
      IsOnShelf  *bool `json:"isOnShelf" form:"isOnShelf"` 
    request.PageInfo
}
