
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MsSpace 创空间 结构体
type MsSpace struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:空间名称;column:name;size:100;" binding:"required"`  //空间名称
  Cover  string `json:"cover" form:"cover" gorm:"comment:封面图;column:cover;"`  //封面图
  Description  *string `json:"description" form:"description" gorm:"comment:简介;column:description;size:500;"`  //简介
  Sdk  *string `json:"sdk" form:"sdk" gorm:"comment:SDK类型;column:sdk;"`  //SDK类型(Gradio/Streamlit)
  Status  *string `json:"status" form:"status" gorm:"comment:状态;column:status;"`  //状态
  AppFile  *string `json:"appFile" form:"appFile" gorm:"comment:入口文件;column:app_file;"`  //入口文件路径
}


// TableName 创空间 MsSpace自定义表名 gva_ms_spaces
func (MsSpace) TableName() string {
    return "gva_ms_spaces"
}







