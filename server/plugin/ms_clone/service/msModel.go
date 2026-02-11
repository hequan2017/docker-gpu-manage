
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
)

var MsModel = new(msModel)

type msModel struct {}

// CreateMsModel 创建模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) CreateMsModel(ctx context.Context, msModel *model.MsModel) (err error) {
	err = global.GVA_DB.Create(msModel).Error
	return err
}

// DeleteMsModel 删除模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) DeleteMsModel(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MsModel{},"id = ?",ID).Error
	return err
}

// DeleteMsModelByIds 批量删除模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) DeleteMsModelByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MsModel{},"id in ?",IDs).Error
	return err
}

// UpdateMsModel 更新模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) UpdateMsModel(ctx context.Context, msModel model.MsModel) (err error) {
	err = global.GVA_DB.Model(&model.MsModel{}).Where("id = ?",msModel.ID).Updates(&msModel).Error
	return err
}

// GetMsModel 根据ID获取模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) GetMsModel(ctx context.Context, ID string) (msModel model.MsModel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msModel).Error
	return
}
// GetMsModelInfoList 分页获取模型库记录
// Author [yourname](https://github.com/yourname)
func (s *msModel) GetMsModelInfoList(ctx context.Context, info request.MsModelSearch) (list []model.MsModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.MsModel{})
    var msModels []model.MsModel
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
	err = db.Find(&msModels).Error
	return  msModels, total, err
}

func (s *msModel)GetMsModelPublic(ctx context.Context) {

}
