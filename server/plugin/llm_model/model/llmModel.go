
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LlmModel 开源大模型 结构体
type LlmModel struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:模型名称;column:name;size:100;" binding:"required"`  //模型名称
  Publisher  *string `json:"publisher" form:"publisher" gorm:"comment:发布者/机构;column:publisher;size:100;"`  //发布者
  Type  string `json:"type" form:"type" gorm:"default:general_llm;comment:模型类型;column:type;type:enum(50);" binding:"required"`  //模型类型
  Parameters  *string `json:"parameters" form:"parameters" gorm:"comment:参数量(如7B, 13B);column:parameters;size:50;"`  //参数量
  Url  *string `json:"url" form:"url" gorm:"comment:魔搭社区地址;column:url;size:255;" binding:"required"`  //魔搭地址
  Description  *string `json:"description" form:"description" gorm:"comment:模型简介;column:description;size:500;"`  //模型简介
}


// TableName 开源大模型 LlmModel自定义表名 llm_models
func (LlmModel) TableName() string {
    return "llm_models"
}







