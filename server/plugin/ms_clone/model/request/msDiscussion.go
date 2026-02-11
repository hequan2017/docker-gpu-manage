
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MsDiscussionSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    Title       string `json:"title" form:"title"`
    Type        *int    `json:"type" form:"type"`
    Pid         *uint   `json:"pid" form:"pid"`
    UserID      *int64  `json:"userId" form:"userId"`
    RelatedID   *int64  `json:"relatedId" form:"relatedId"`
    RelatedType string `json:"relatedType" form:"relatedType"`
    IsSolved    *bool   `json:"isSolved" form:"isSolved"`
    request.PageInfo
}
