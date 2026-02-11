
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/model/request"
)

var ServerAsset = new(asset)

type asset struct {}
// CreateServerAsset 创建服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) CreateServerAsset(ctx context.Context, asset *model.ServerAsset) (err error) {
	err = global.GVA_DB.Create(asset).Error
	return err
}

// DeleteServerAsset 删除服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) DeleteServerAsset(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.ServerAsset{},"id = ?",ID).Error
	return err
}

// DeleteServerAssetByIds 批量删除服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) DeleteServerAssetByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.ServerAsset{},"id in ?",IDs).Error
	return err
}

// UpdateServerAsset 更新服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) UpdateServerAsset(ctx context.Context, asset model.ServerAsset) (err error) {
	err = global.GVA_DB.Model(&model.ServerAsset{}).Where("id = ?",asset.ID).Updates(&asset).Error
	return err
}

// GetServerAsset 根据ID获取服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) GetServerAsset(ctx context.Context, ID string) (asset model.ServerAsset, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&asset).Error
	return
}
// GetServerAssetInfoList 分页获取服务器资产记录
// Author [yourname](https://github.com/yourname)
func (s *asset) GetServerAssetInfoList(ctx context.Context, info request.ServerAssetSearch) (list []model.ServerAsset, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ServerAsset{})
    var assets []model.ServerAsset
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.HostName != nil && *info.HostName != "" {
        db = db.Where("host_name LIKE ?", "%"+ *info.HostName+"%")
    }
    if info.IP != nil && *info.IP != "" {
        db = db.Where("ip LIKE ?", "%"+ *info.IP+"%")
    }
    if info.Sn != nil && *info.Sn != "" {
        db = db.Where("sn LIKE ?", "%"+ *info.Sn+"%")
    }
    if info.Configuration != nil && *info.Configuration != "" {
        db = db.Where("configuration LIKE ?", "%"+ *info.Configuration+"%")
    }
    if info.Status != "" {
        db = db.Where("status = ?", info.Status)
    }
    if info.ServiceType != nil && *info.ServiceType != "" {
        db = db.Where("service_type LIKE ?", "%"+ *info.ServiceType+"%")
    }
			if len(info.DeployTimeRange) == 2 {
				db = db.Where("deploy_time BETWEEN ? AND ? ", info.DeployTimeRange[0], info.DeployTimeRange[1])
			}
			if len(info.OfflineTimeRange) == 2 {
				db = db.Where("offline_time BETWEEN ? AND ? ", info.OfflineTimeRange[0], info.OfflineTimeRange[1])
			}
			if len(info.ScrapTimeRange) == 2 {
				db = db.Where("scrap_time BETWEEN ? AND ? ", info.ScrapTimeRange[0], info.ScrapTimeRange[1])
			}
    if info.ScrapReason != nil && *info.ScrapReason != "" {
        db = db.Where("scrap_reason LIKE ?", "%"+ *info.ScrapReason+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["deploy_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&assets).Error
	return  assets, total, err
}

func (s *asset)GetServerAssetPublic(ctx context.Context) {

}
