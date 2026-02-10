package pcdn

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pcdn"
	pcdnReq "github.com/flipped-aurora/gin-vue-admin/server/model/pcdn/request"
)

type PcdnNodeService struct{}

func (s *PcdnNodeService) CreatePcdnNode(_ context.Context, in *pcdn.PcdnNode) error {
	return global.GVA_DB.Create(in).Error
}

func (s *PcdnNodeService) DeletePcdnNode(_ context.Context, id uint) error {
	return global.GVA_DB.Delete(&pcdn.PcdnNode{}, id).Error
}

func (s *PcdnNodeService) UpdatePcdnNode(_ context.Context, in pcdn.PcdnNode) error {
	return global.GVA_DB.Model(&pcdn.PcdnNode{}).Where("id = ?", in.ID).Updates(&in).Error
}

func (s *PcdnNodeService) GetPcdnNode(_ context.Context, id uint) (out pcdn.PcdnNode, err error) {
	err = global.GVA_DB.First(&out, id).Error
	return
}

func (s *PcdnNodeService) GetPcdnNodeList(_ context.Context, info pcdnReq.PcdnNodeSearch) (list []pcdn.PcdnNode, total int64, err error) {
	db := global.GVA_DB.Model(&pcdn.PcdnNode{})
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Region != "" {
		db = db.Where("region LIKE ?", "%"+info.Region+"%")
	}
	if info.Online != nil {
		db = db.Where("online = ?", *info.Online)
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

func (s *PcdnNodeService) SwitchOnline(_ context.Context, in pcdnReq.PcdnNodeOnlineRequest) error {
	return global.GVA_DB.Model(&pcdn.PcdnNode{}).Where("id = ?", in.ID).Updates(map[string]any{"online": in.Online}).Error
}

func (s *PcdnNodeService) UpdateWeight(_ context.Context, in pcdnReq.PcdnNodeWeightRequest) error {
	return global.GVA_DB.Model(&pcdn.PcdnNode{}).Where("id = ?", in.ID).Updates(map[string]any{"weight": in.Weight}).Error
}
