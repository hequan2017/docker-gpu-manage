
package computenode

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
    computenodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/computenode/request"
)

type ComputeNodeService struct {}
// CreateComputeNode 创建算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService) CreateComputeNode(ctx context.Context, computeNode *computenode.ComputeNode) (err error) {
	err = global.GVA_DB.Create(computeNode).Error
	return err
}

// DeleteComputeNode 删除算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService)DeleteComputeNode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&computenode.ComputeNode{},"id = ?",ID).Error
	return err
}

// DeleteComputeNodeByIds 批量删除算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService)DeleteComputeNodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]computenode.ComputeNode{},"id in ?",IDs).Error
	return err
}

// UpdateComputeNode 更新算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService)UpdateComputeNode(ctx context.Context, computeNode computenode.ComputeNode) (err error) {
	err = global.GVA_DB.Model(&computenode.ComputeNode{}).Where("id = ?",computeNode.ID).Updates(&computeNode).Error
	return err
}

// GetComputeNode 根据ID获取算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService)GetComputeNode(ctx context.Context, ID string) (computeNode computenode.ComputeNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&computeNode).Error
	return
}
// GetComputeNodeInfoList 分页获取算力节点记录
// Author [yourname](https://github.com/yourname)
func (computeNodeService *ComputeNodeService)GetComputeNodeInfoList(ctx context.Context, info computenodeReq.ComputeNodeSearch) (list []computenode.ComputeNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&computenode.ComputeNode{})
    var computeNodes []computenode.ComputeNode
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Region != nil && *info.Region != "" {
        db = db.Where("region LIKE ?", "%"+ *info.Region+"%")
    }
    if info.PublicIp != nil && *info.PublicIp != "" {
        db = db.Where("public_ip LIKE ?", "%"+ *info.PublicIp+"%")
    }
    if info.PrivateIp != nil && *info.PrivateIp != "" {
        db = db.Where("private_ip LIKE ?", "%"+ *info.PrivateIp+"%")
    }
    if info.GpuName != nil && *info.GpuName != "" {
        db = db.Where("gpu_name LIKE ?", "%"+ *info.GpuName+"%")
    }
    if info.IsOnShelf != nil {
        db = db.Where("is_on_shelf = ?", *info.IsOnShelf)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&computeNodes).Error
	return  computeNodes, total, err
}
func (computeNodeService *ComputeNodeService)GetComputeNodePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
