package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
)

var MsDiscussion = new(msDiscussion)

type msDiscussion struct{}

// CreateMsDiscussion 创建社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) CreateMsDiscussion(ctx context.Context, discussion *model.MsDiscussion) (err error) {
	err = global.GVA_DB.Create(discussion).Error
	return err
}

// DeleteMsDiscussion 删除社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) DeleteMsDiscussion(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MsDiscussion{}, "id = ?", ID).Error
	return err
}

// DeleteMsDiscussionByIds 批量删除社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) DeleteMsDiscussionByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MsDiscussion{}, "id in ?", IDs).Error
	return err
}

// UpdateMsDiscussion 更新社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) UpdateMsDiscussion(ctx context.Context, discussion model.MsDiscussion) (err error) {
	err = global.GVA_DB.Model(&model.MsDiscussion{}).Where("id = ?", discussion.ID).Updates(&discussion).Error
	return err
}

// GetMsDiscussion 根据ID获取社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) GetMsDiscussion(ctx context.Context, ID string) (discussion model.MsDiscussion, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&discussion).Error
	return
}

// GetMsDiscussionInfoList 分页获取社区讨论记录
// Author [yourname](https://github.com/yourname)
func (s *msDiscussion) GetMsDiscussionInfoList(ctx context.Context, info request.MsDiscussionSearch) (list []model.MsDiscussion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.MsDiscussion{})
	var discussions []model.MsDiscussion
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Pid != nil {
		db = db.Where("pid = ?", info.Pid)
	}
	if info.UserID != nil {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.RelatedID != nil {
		db = db.Where("related_id = ?", info.RelatedID)
	}
	if info.RelatedType != "" {
		db = db.Where("related_type = ?", info.RelatedType)
	}
	if info.IsSolved != nil {
		db = db.Where("is_solved = ?", info.IsSolved)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&discussions).Error
	return discussions, total, err
}
func (s *msDiscussion) GetMsDiscussionDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_users").Select("nickName as label,ID as value").Scan(&userId)
	res["userId"] = userId
	return
}

func (s *msDiscussion) GetMsDiscussionPublic(ctx context.Context) {

}
