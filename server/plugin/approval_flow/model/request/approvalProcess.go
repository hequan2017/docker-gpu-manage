
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ApprovalProcessSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Title  *string `json:"title" form:"title"` 
       Version  *string `json:"version" form:"version"` 
       Content  *string `json:"content" form:"content"` 
       TargetServer  *string `json:"targetServer" form:"targetServer"` 
       Command  *string `json:"command" form:"command"` 
       Status  string `json:"status" form:"status"` 
       ApplicantId  *int `json:"applicantId" form:"applicantId"` 
       ApproverId  *int `json:"approverId" form:"approverId"` 
       Logs  string `json:"logs" form:"logs"` 
    request.PageInfo
}
