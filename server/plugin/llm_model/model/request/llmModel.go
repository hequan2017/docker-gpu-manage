
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type LlmModelSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       Publisher  *string `json:"publisher" form:"publisher"` 
       Type  string `json:"type" form:"type"` 
       Parameters  *string `json:"parameters" form:"parameters"` 
       Url  *string `json:"url" form:"url"` 
       Description  *string `json:"description" form:"description"` 
    request.PageInfo
}
