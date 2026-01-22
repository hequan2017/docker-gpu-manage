package service

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward/model/request"
	"go.uber.org/zap"
)

var (
	PortForward        = new(portForward)
	forwarderMgr       *ForwarderManager
	forwarderMgrInited bool
)

type portForward struct{}

// initForwarderManager 初始化转发器管理器
func initForwarderManager() {
	if !forwarderMgrInited {
		forwarderMgr = GetForwarderManager(global.GVA_LOG)
		forwarderMgrInited = true
	}
}

// CreatePortForward 创建端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) CreatePortForward(portForward *model.PortForward) (err error) {
	err = global.GVA_DB.Create(portForward).Error
	if err != nil {
		return err
	}

	// 如果规则状态为启用，启动转发
	if portForward.Status {
		initForwarderManager()
		if err := forwarderMgr.StartPortForward(portForward); err != nil {
			global.GVA_LOG.Error("启动端口转发失败",
				zap.Uint("id", portForward.ID),
				zap.Error(err),
			)
			// 不返回错误，因为规则已经创建成功
		}
	}

	return nil
}

// DeletePortForward 删除端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) DeletePortForward(ID string) (err error) {
	// 先停止转发
	initForwarderManager()
	id, _ := strconv.ParseUint(ID, 10, 32)
	if forwarderMgr.IsRunning(uint(id)) {
		_ = forwarderMgr.StopPortForward(uint(id))
	}

	err = global.GVA_DB.Delete(&model.PortForward{}, "id = ?", ID).Error
	return err
}

// DeletePortForwardByIds 批量删除端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) DeletePortForwardByIds(IDs []string) (err error) {
	// 先停止所有转发
	initForwarderManager()
	for _, ID := range IDs {
		id, _ := strconv.ParseUint(ID, 10, 32)
		if forwarderMgr.IsRunning(uint(id)) {
			_ = forwarderMgr.StopPortForward(uint(id))
		}
	}

	err = global.GVA_DB.Delete(&[]model.PortForward{}, "id in ?", IDs).Error
	return err
}

// UpdatePortForward 更新端口转发规则
// Author [BMad Orchestrator]
func (s *portForward) UpdatePortForward(portForward model.PortForward) (err error) {
	// 先停止旧的转发
	initForwarderManager()
	if forwarderMgr.IsRunning(portForward.ID) {
		_ = forwarderMgr.StopPortForward(portForward.ID)
	}

	err = global.GVA_DB.Model(&model.PortForward{}).Where("id = ?", portForward.ID).Updates(&portForward).Error
	if err != nil {
		return err
	}

	// 如果规则状态为启用，启动新的转发
	if portForward.Status {
		if err := forwarderMgr.StartPortForward(&portForward); err != nil {
			global.GVA_LOG.Error("启动端口转发失败",
				zap.Uint("id", portForward.ID),
				zap.Error(err),
			)
		}
	}

	return nil
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
	if err != nil {
		return err
	}

	// 根据状态启动或停止转发
	initForwarderManager()
	id, _ := strconv.ParseUint(ID, 10, 32)
	ruleID := uint(id)

	if status {
		// 启用规则，启动转发
		rule, err := s.GetPortForward(ID)
		if err == nil {
			if err := forwarderMgr.StartPortForward(&rule); err != nil {
				global.GVA_LOG.Error("启动端口转发失败",
					zap.Uint("id", ruleID),
					zap.Error(err),
				)
			}
		}
	} else {
		// 禁用规则，停止转发
		if forwarderMgr.IsRunning(ruleID) {
			_ = forwarderMgr.StopPortForward(ruleID)
		}
	}

	return nil
}

// SyncAllPortForwards 同步所有端口转发规则状态
// 在系统启动时调用，确保数据库状态与实际运行状态一致
func (s *portForward) SyncAllPortForwards() error {
	initForwarderManager()

	var rules []model.PortForward
	err := global.GVA_DB.Find(&rules).Error
	if err != nil {
		return err
	}

	forwarderMgr.SyncPortForwardStatus(rules)

	return nil
}

// GetForwarderStatus 获取端口转发运行状态
func (s *portForward) GetForwarderStatus(ID string) map[string]interface{} {
	initForwarderManager()
	id, _ := strconv.ParseUint(ID, 10, 32)
	return forwarderMgr.GetForwarderStatus(uint(id))
}

// GetAllForwarderStatus 获取所有端口转发运行状态
func (s *portForward) GetAllForwarderStatus() map[string]interface{} {
	initForwarderManager()

	runningIDs := forwarderMgr.GetRunningForwarders()
	totalCount := int64(0)
	runningCount := 0

	// 获取总规则数
	global.GVA_DB.Model(&model.PortForward{}).Count(&totalCount)

	// 获取启用的规则数
	var enabledRules []model.PortForward
	global.GVA_DB.Where("status = ?", true).Find(&enabledRules)

	return map[string]interface{}{
		"total_rules":     totalCount,
		"running_forwarders": len(runningIDs),
		"running_ids":      runningIDs,
		"enabled_count":    len(enabledRules),
	}
}


