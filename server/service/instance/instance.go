package instance

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	instanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/instance/request"
)

type InstanceService struct{}

// CreateInstance 创建实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) CreateInstance(ctx context.Context, inst *instanceModel.Instance) (err error) {
	err = global.GVA_DB.Create(inst).Error
	return err
}

// DeleteInstance 删除实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) DeleteInstance(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&instanceModel.Instance{}, "id = ?", ID).Error
	return err
}

// DeleteInstanceByIds 批量删除实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) DeleteInstanceByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]instanceModel.Instance{}, "id in ?", IDs).Error
	return err
}

// UpdateInstance 更新实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) UpdateInstance(ctx context.Context, inst instanceModel.Instance) (err error) {
	err = global.GVA_DB.Model(&instanceModel.Instance{}).Where("id = ?", inst.ID).Updates(&inst).Error
	return err
}

// GetInstance 根据ID获取实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) GetInstance(ctx context.Context, ID string) (inst instanceModel.Instance, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&inst).Error
	return
}

// GetInstanceInfoList 分页获取实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) GetInstanceInfoList(ctx context.Context, info instanceReq.InstanceSearch) (list []instanceModel.Instance, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&instanceModel.Instance{})
	var instances []instanceModel.Instance
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.ImageId != nil {
		db = db.Where("image_id = ?", *info.ImageId)
	}
	if info.SpecId != nil {
		db = db.Where("spec_id = ?", *info.SpecId)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.NodeId != nil {
		db = db.Where("node_id = ?", *info.NodeId)
	}
	if info.ContainerId != nil && *info.ContainerId != "" {
		db = db.Where("container_id LIKE ?", "%"+*info.ContainerId+"%")
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.ContainerStatus != nil && *info.ContainerStatus != "" {
		db = db.Where("container_status = ?", *info.ContainerStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&instances).Error
	return instances, total, err
}
func (instanceService *InstanceService) GetInstanceDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	imageId := make([]map[string]any, 0)

	global.GVA_DB.Table("image_registry").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&imageId)
	res["imageId"] = imageId
	nodeId := make([]map[string]any, 0)

	global.GVA_DB.Table("compute_node").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&nodeId)
	res["nodeId"] = nodeId
	specId := make([]map[string]any, 0)

	global.GVA_DB.Table("product_spec").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&specId)
	res["specId"] = specId
	userId := make([]map[string]any, 0)

	global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("username as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
func (instanceService *InstanceService) GetInstancePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
