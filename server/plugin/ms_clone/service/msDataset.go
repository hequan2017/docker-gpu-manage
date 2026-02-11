
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
)

var MsDataset = new(msDataset)

type msDataset struct {}

// CreateMsDataset 创建数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) CreateMsDataset(ctx context.Context, msDataset *model.MsDataset) (err error) {
	err = global.GVA_DB.Create(msDataset).Error
	return err
}

// DeleteMsDataset 删除数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) DeleteMsDataset(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MsDataset{},"id = ?",ID).Error
	return err
}

// DeleteMsDatasetByIds 批量删除数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) DeleteMsDatasetByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MsDataset{},"id in ?",IDs).Error
	return err
}

// UpdateMsDataset 更新数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) UpdateMsDataset(ctx context.Context, msDataset model.MsDataset) (err error) {
	err = global.GVA_DB.Model(&model.MsDataset{}).Where("id = ?",msDataset.ID).Updates(&msDataset).Error
	return err
}

// GetMsDataset 根据ID获取数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) GetMsDataset(ctx context.Context, ID string) (msDataset model.MsDataset, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msDataset).Error
	return
}
// GetMsDatasetInfoList 分页获取数据集记录
// Author [yourname](https://github.com/yourname)
func (s *msDataset) GetMsDatasetInfoList(ctx context.Context, info request.MsDatasetSearch) (list []model.MsDataset, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.MsDataset{})
    var msDatasets []model.MsDataset
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
	err = db.Find(&msDatasets).Error
	return  msDatasets, total, err
}

func (s *msDataset)GetMsDatasetPublic(ctx context.Context) {

}
