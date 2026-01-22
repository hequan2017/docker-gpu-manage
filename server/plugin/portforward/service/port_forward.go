package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model/request"
)

var PortForward = new(portForward)

type portForward struct{}

// CreatePortForward 创建端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) CreatePortForward(portForward *model.PortForward) (err error) {
	err = global.GVA_DB.Create(portForward).Error
	return err
}

// DeletePortForward 删除端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) DeletePortForward(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.PortForward{}, "id = ?", ID).Error
	return err
}

// DeletePortForwardByIds 批量删除端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) DeletePortForwardByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.PortForward{}, "id in ?", IDs).Error
	return err
}

// UpdatePortForward 更新端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) UpdatePortForward(portForward model.PortForward) (err error) {
	err = global.GVA_DB.Model(&model.PortForward{}).Where("id = ?", portForward.ID).Updates(&portForward).Error
	return err
}

// GetPortForward 根据ID获取端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) GetPortForward(ID string) (portForward model.PortForward, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&portForward).Error
	return
}

// GetPortForwardList 分页获取端口转发规则列表
// Author [BMad Orchestrator]
func (s *portForward) GetPortForwardList(info request.PortForwardSearch) (list []model.PortForward, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.PortForward{})
	var portForwards []model.PortForward

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.SourceIP != "" {
		db = db.Where("source_ip LIKE ?", "%"+info.SourceIP+"%")
	}
	if info.SourcePort != 0 {
		db = db.Where("source_port = ?", info.SourcePort)
	}
	if info.Protocol != "" {
		db = db.Where("protocol = ?", info.Protocol)
	}
	if info.TargetIP != "" {
		db = db.Where("target_ip LIKE ?", "%"+info.TargetIP+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&portForwards).Error
	return portForwards, total, err
}

// UpdatePortForwardStatus 更新端口转发规则状态
// Author [BMad Orchestrator]
func (s *portForward) UpdatePortForwardStatus(ID string, status bool) (err error) {
	err = global.GVA_DB.Model(&model.PortForward{}).Where("id = ?", ID).Update("status", status).Error
	return err
}
