
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MsDataset 数据集 结构体
type MsDataset struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:数据集名称;column:name;size:100;" binding:"required"`  //数据集名称
  Cover  string `json:"cover" form:"cover" gorm:"comment:封面图;column:cover;"`  //封面图
  Description  *string `json:"description" form:"description" gorm:"comment:简介;column:description;size:500;"`  //简介
  Size  *string `json:"size" form:"size" gorm:"comment:数据集大小;column:size;"`  //数据集大小
  Publisher  *string `json:"publisher" form:"publisher" gorm:"comment:发布者;column:publisher;"`  //发布者
  Readme  *string `json:"readme" form:"readme" gorm:"comment:详情文档;column:readme;type:text;"`  //详情文档
}


// TableName 数据集 MsDataset自定义表名 gva_ms_datasets
func (MsDataset) TableName() string {
    return "gva_ms_datasets"
}







