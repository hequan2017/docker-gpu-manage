package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/model/request"
)

type DellAssetService struct{}

// CreateDellAsset 创建戴尔服务器资产
// Author [BMad Orchestrator]
func (das *DellAssetService) CreateDellAsset(dellAsset *model.DellAsset) (err error) {
	err = global.GVA_DB.Create(dellAsset).Error
	return err
}

// DeleteDellAsset 删除戴尔服务器资产
// Author [BMad Orchestrator]
func (das *DellAssetService) DeleteDellAsset(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.DellAsset{}, "id = ?", ID).Error
	return err
}

// DeleteDellAssetByIds 批量删除戴尔服务器资产
// Author [BMad Orchestrator]
func (das *DellAssetService) DeleteDellAssetByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.DellAsset{}, "id in ?", IDs).Error
	return err
}

// UpdateDellAsset 更新戴尔服务器资产
// Author [BMad Orchestrator]
func (das *DellAssetService) UpdateDellAsset(dellAsset model.DellAsset) (err error) {
	err = global.GVA_DB.Model(&model.DellAsset{}).Where("id = ?", dellAsset.ID).Updates(&dellAsset).Error
	return err
}

// GetDellAsset 根据ID获取戴尔服务器资产
// Author [BMad Orchestrator]
func (das *DellAssetService) GetDellAsset(ID string) (dellAsset model.DellAsset, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dellAsset).Error
	return
}

// GetDellAssetList 分页获取戴尔服务器资产列表
// Author [BMad Orchestrator]
func (das *DellAssetService) GetDellAssetList(info request.DellAssetSearch) (list []model.DellAsset, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.DellAsset{})
	var dellAssets []model.DellAsset

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.HostName != "" {
		db = db.Where("host_name LIKE ?", "%"+info.HostName+"%")
	}
	if info.ServiceTag != "" {
		db = db.Where("service_tag LIKE ?", "%"+info.ServiceTag+"%")
	}
	if info.AssetNumber != "" {
		db = db.Where("asset_number LIKE ?", "%"+info.AssetNumber+"%")
	}
	if info.Model != "" {
		db = db.Where("model LIKE ?", "%"+info.Model+"%")
	}
	if info.IPAddress != "" {
		db = db.Where("ip_address LIKE ?", "%"+info.IPAddress+"%")
	}
	if info.Cabinet != "" {
		db = db.Where("cabinet LIKE ?", "%"+info.Cabinet+"%")
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	if info.Manager != "" {
		db = db.Where("manager LIKE ?", "%"+info.Manager+"%")
	}
	if info.PowerStatus != "" {
		db = db.Where("power_status = ?", info.PowerStatus)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.OS != "" {
		db = db.Where("os LIKE ?", "%"+info.OS+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&dellAssets).Error
	return dellAssets, total, err
}

// GetDellAssetStatistics 获取资产统计信息
// Author [BMad Orchestrator]
func (das *DellAssetService) GetDellAssetStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总数统计
	var totalCount int64
	global.GVA_DB.Model(&model.DellAsset{}).Count(&totalCount)
	stats["total"] = totalCount

	// 按状态统计
	var onlineCount, offlineCount, maintenanceCount int64
	global.GVA_DB.Model(&model.DellAsset{}).Where("status = ?", "online").Count(&onlineCount)
	global.GVA_DB.Model(&model.DellAsset{}).Where("status = ?", "offline").Count(&offlineCount)
	global.GVA_DB.Model(&model.DellAsset{}).Where("status = ?", "maintenance").Count(&maintenanceCount)

	stats["online"] = onlineCount
	stats["offline"] = offlineCount
	stats["maintenance"] = maintenanceCount

	// 按部门统计
	var departmentStats []struct {
		Department string
		Count      int64
	}
	global.GVA_DB.Model(&model.DellAsset{}).Select("department, count(*) as count").
		Where("department != ''").
		Group("department").
		Scan(&departmentStats)
	stats["by_department"] = departmentStats

	// 按型号统计
	var modelStats []struct {
		Model string
		Count int64
	}
	global.GVA_DB.Model(&model.DellAsset{}).Select("model, count(*) as count").
		Where("model != ''").
		Group("model").
		Order("count DESC").
		Limit(10).
		Scan(&modelStats)
	stats["by_model"] = modelStats

	return stats, nil
}
