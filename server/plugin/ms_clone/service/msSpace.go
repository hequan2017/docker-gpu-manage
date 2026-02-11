
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
)

var MsSpace = new(msSpace)

type msSpace struct {}

// CreateMsSpace 创建创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) CreateMsSpace(ctx context.Context, msSpace *model.MsSpace) (err error) {
	err = global.GVA_DB.Create(msSpace).Error
	return err
}

// DeleteMsSpace 删除创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) DeleteMsSpace(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MsSpace{},"id = ?",ID).Error
	return err
}

// DeleteMsSpaceByIds 批量删除创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) DeleteMsSpaceByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MsSpace{},"id in ?",IDs).Error
	return err
}

// UpdateMsSpace 更新创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) UpdateMsSpace(ctx context.Context, msSpace model.MsSpace) (err error) {
	err = global.GVA_DB.Model(&model.MsSpace{}).Where("id = ?",msSpace.ID).Updates(&msSpace).Error
	return err
}

// GetMsSpace 根据ID获取创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) GetMsSpace(ctx context.Context, ID string) (msSpace model.MsSpace, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msSpace).Error
	return
}
// GetMsSpaceInfoList 分页获取创空间记录
// Author [yourname](https://github.com/yourname)
func (s *msSpace) GetMsSpaceInfoList(ctx context.Context, info request.MsSpaceSearch) (list []model.MsSpace, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.MsSpace{})
    var msSpaces []model.MsSpace
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&msSpaces).Error
	return  msSpaces, total, err
}

func (s *msSpace)GetMsSpacePublic(ctx context.Context) {

}
