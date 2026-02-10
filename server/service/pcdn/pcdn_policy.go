package pcdn

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
)

type PcdnPolicyService struct{}

func (s *PcdnPolicyService) CreatePcdnPolicy(_ context.Context, in *pcdn.PcdnPolicy) error {
	return global.GVA_DB.Create(in).Error
}

func (s *PcdnPolicyService) DeletePcdnPolicy(_ context.Context, id uint) error {
	return global.GVA_DB.Delete(&pcdn.PcdnPolicy{}, id).Error
}

func (s *PcdnPolicyService) UpdatePcdnPolicy(_ context.Context, in pcdn.PcdnPolicy) error {
	return global.GVA_DB.Model(&pcdn.PcdnPolicy{}).Where("id = ?", in.ID).Updates(&in).Error
}

func (s *PcdnPolicyService) GetPcdnPolicy(_ context.Context, id uint) (out pcdn.PcdnPolicy, err error) {
	err = global.GVA_DB.First(&out, id).Error
	return
}

func (s *PcdnPolicyService) GetPcdnPolicyList(_ context.Context, info pcdnReq.PcdnPolicySearch) (list []pcdn.PcdnPolicy, total int64, err error) {
	db := global.GVA_DB.Model(&pcdn.PcdnPolicy{})
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Enabled != nil {
		db = db.Where("enabled = ?", *info.Enabled)
	}
	if info.Published != nil {
		db = db.Where("published = ?", *info.Published)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset(info.PageSize * (info.Page - 1))
	}
	err = db.Order("id desc").Find(&list).Error
	return
}

func (s *PcdnPolicyService) GrayRelease(_ context.Context, in pcdnReq.PcdnPolicyGrayReleaseRequest) error {
	return global.GVA_DB.Model(&pcdn.PcdnPolicy{}).Where("id = ?", in.ID).Updates(map[string]any{
		"gray_percent":   in.GrayPercent,
		"published":      true,
		"published_note": in.Note,
	}).Error
}

func (s *PcdnPolicyService) SwitchPolicy(_ context.Context, in pcdnReq.PcdnPolicySwitchRequest) error {
	return global.GVA_DB.Model(&pcdn.PcdnPolicy{}).Where("id = ?", in.ID).Updates(map[string]any{"enabled": in.Enabled}).Error
}
