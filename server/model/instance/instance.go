
// 自动生成模板Instance
package instance
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 实例管理 结构体  Instance
type Instance struct {
    global.GVA_MODEL
  ImageId  *int64 `json:"imageId" form:"imageId" gorm:"comment:镜像ID;column:image_id;" binding:"required"`  //镜像
  SpecId  *int64 `json:"specId" form:"specId" gorm:"comment:产品规格ID;column:spec_id;" binding:"required"`  //产品规格
  UserId  *int64 `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;"`  //用户
  NodeId  *int64 `json:"nodeId" form:"nodeId" gorm:"comment:算力节点ID;column:node_id;" binding:"required"`  //算力节点
  ContainerId  *string `json:"containerId" form:"containerId" gorm:"comment:Docker容器ID;column:container_id;size:255;"`  //Docker容器
  Name  *string `json:"name" form:"name" gorm:"comment:实例名称;column:name;size:255;" binding:"required"`  //实例名称
  ContainerStatus  *string `json:"containerStatus" form:"containerStatus" gorm:"comment:容器状态;column:container_status;size:50;"`  //容器状态
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注信息;column:remark;size:1000;"`  //备注
}


// TableName 实例管理 Instance自定义表名 instance
func (Instance) TableName() string {
    return "instance"
}





