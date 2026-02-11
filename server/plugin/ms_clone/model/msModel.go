
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MsModel 模型库 结构体
type MsModel struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:模型名称;column:name;size:100;" binding:"required"`  //模型名称
  Cover  string `json:"cover" form:"cover" gorm:"comment:封面图;column:cover;"`  //封面图
  Description  *string `json:"description" form:"description" gorm:"comment:简介;column:description;size:500;"`  //简介
  TaskType  *string `json:"taskType" form:"taskType" gorm:"comment:任务类型;column:task_type;"`  //任务类型
  Publisher  *string `json:"publisher" form:"publisher" gorm:"comment:发布者;column:publisher;"`  //发布者
  Readme  *string `json:"readme" form:"readme" gorm:"comment:详情文档;column:readme;type:text;"`  //详情文档
  DownloadCount  *int64 `json:"downloadCount" form:"downloadCount" gorm:"default:0;comment:下载量;column:download_count;"`  //下载量
}


// TableName 模型库 MsModel自定义表名 gva_ms_models
func (MsModel) TableName() string {
    return "gva_ms_models"
}







