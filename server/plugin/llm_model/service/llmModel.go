
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/model/request"
)

var LlmModel = new(llm)

type llm struct {}
// CreateLlmModel 创建开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) CreateLlmModel(ctx context.Context, llm *model.LlmModel) (err error) {
	err = global.GVA_DB.Create(llm).Error
	return err
}

// DeleteLlmModel 删除开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) DeleteLlmModel(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.LlmModel{},"id = ?",ID).Error
	return err
}

// DeleteLlmModelByIds 批量删除开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) DeleteLlmModelByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.LlmModel{},"id in ?",IDs).Error
	return err
}

// UpdateLlmModel 更新开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) UpdateLlmModel(ctx context.Context, llm model.LlmModel) (err error) {
	err = global.GVA_DB.Model(&model.LlmModel{}).Where("id = ?",llm.ID).Updates(&llm).Error
	return err
}

// GetLlmModel 根据ID获取开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) GetLlmModel(ctx context.Context, ID string) (llm model.LlmModel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&llm).Error
	return
}
// GetLlmModelInfoList 分页获取开源大模型记录
// Author [yourname](https://github.com/yourname)
func (s *llm) GetLlmModelInfoList(ctx context.Context, info request.LlmModelSearch) (list []model.LlmModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.LlmModel{})
    var llms []model.LlmModel
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Publisher != nil && *info.Publisher != "" {
        db = db.Where("publisher LIKE ?", "%"+ *info.Publisher+"%")
    }
    if info.Type != "" {
        db = db.Where("type = ?", info.Type)
    }
    if info.Parameters != nil && *info.Parameters != "" {
        db = db.Where("parameters LIKE ?", "%"+ *info.Parameters+"%")
    }
    if info.Url != nil && *info.Url != "" {
        db = db.Where("url LIKE ?", "%"+ *info.Url+"%")
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
	err = db.Find(&llms).Error
	return  llms, total, err
}

func (s *llm)GetLlmModelPublic(ctx context.Context) {

}
