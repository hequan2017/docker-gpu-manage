package pcdn

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// PcdnPolicy PCDN策略
type PcdnPolicy struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" gorm:"column:name;size:128;comment:策略名称" binding:"required"`
	Template      string `json:"template" form:"template" gorm:"column:template;type:text;comment:策略模板" binding:"required"`
	GrayPercent   int    `json:"grayPercent" form:"grayPercent" gorm:"column:gray_percent;default:0;comment:灰度比例"`
	Enabled       bool   `json:"enabled" form:"enabled" gorm:"column:enabled;default:false;comment:是否启用"`
	Published     bool   `json:"published" form:"published" gorm:"column:published;default:false;comment:是否已发布"`
	Version       string `json:"version" form:"version" gorm:"column:version;size:64;comment:版本"`
	PublishedNote string `json:"publishedNote" form:"publishedNote" gorm:"column:published_note;size:255;comment:发布说明"`
}

func (PcdnPolicy) TableName() string {
	return "pcdn_policy"
}
