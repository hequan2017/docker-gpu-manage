package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model/request"
)

var PcdnNodeService = new(pcdnNodeService)

type pcdnNodeService struct{}

// CreatePcdnNode 创建PCDN节点
func (s *pcdnNodeService) CreatePcdnNode(pcdnNode *model.PcdnNode) (err error) {
	err = global.GVA_DB.Create(pcdnNode).Error
	return err
}

// DeletePcdnNode 删除PCDN节点
func (s *pcdnNodeService) DeletePcdnNode(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.PcdnNode{}, "id = ?", ID).Error
	return err
}

// DeletePcdnNodeByIds 批量删除PCDN节点
func (s *pcdnNodeService) DeletePcdnNodeByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.PcdnNode{}, "id in ?", ids).Error
	return err
}

// UpdatePcdnNode 更新PCDN节点
func (s *pcdnNodeService) UpdatePcdnNode(pcdnNode model.PcdnNode) (err error) {
	err = global.GVA_DB.Model(&model.PcdnNode{}).Where("id = ?", pcdnNode.ID).Updates(&pcdnNode).Error
	return err
}

// GetPcdnNode 根据id获取PCDN节点
func (s *pcdnNodeService) GetPcdnNode(ID string) (pcdnNode model.PcdnNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pcdnNode).Error
	return
}

// GetPcdnNodeInfoList 分页获取PCDN节点列表
func (s *pcdnNodeService) GetPcdnNodeInfoList(info request.PcdnNodeSearch) (list []model.PcdnNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.PcdnNode{})
	var pcdnNodes []model.PcdnNode

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.IP != "" {
		db = db.Where("ip LIKE ?", "%"+info.IP+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&pcdnNodes).Error
	return pcdnNodes, total, err
}
