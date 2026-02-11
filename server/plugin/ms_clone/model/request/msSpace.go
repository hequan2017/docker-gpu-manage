
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MsSpaceSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       Description  *string `json:"description" form:"description"` 
    request.PageInfo
}
