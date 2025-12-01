
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InstanceSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      ImageId  *int `json:"imageId" form:"imageId"` 
      SpecId  *int `json:"specId" form:"specId"` 
      UserId  *int `json:"userId" form:"userId"` 
      NodeId  *int `json:"nodeId" form:"nodeId"` 
      ContainerId  *string `json:"containerId" form:"containerId"` 
      Name  *string `json:"name" form:"name"` 
      ContainerStatus  *string `json:"containerStatus" form:"containerStatus"` 
    request.PageInfo
}
