
// 自动生成模板ImageRegistry
package imageregistry
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 镜像库 结构体  ImageRegistry
type ImageRegistry struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:镜像库名称;column:name;size:255;" binding:"required"`  //名字
  Address  *string `json:"address" form:"address" gorm:"comment:镜像库地址;column:address;size:500;" binding:"required"`  //地址
  Description  *string `json:"description" form:"description" gorm:"comment:镜像库描述;column:description;size:1000;"`  //描述
  Source  *string `json:"source" form:"source" gorm:"comment:镜像库来源;column:source;size:255;"`  //来源
  IsOnShelf  *bool `json:"isOnShelf" form:"isOnShelf" gorm:"default:true;comment:是否上架;column:is_on_shelf;" binding:"required"`  //是否上架
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注信息;column:remark;size:1000;"`  //备注
}


// TableName 镜像库 ImageRegistry自定义表名 image_registry
func (ImageRegistry) TableName() string {
    return "image_registry"
}





