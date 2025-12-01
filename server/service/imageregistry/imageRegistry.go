
package imageregistry

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry"
    imageregistryReq "github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry/request"
)

type ImageRegistryService struct {}
// CreateImageRegistry 创建镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService) CreateImageRegistry(ctx context.Context, imageRegistry *imageregistry.ImageRegistry) (err error) {
	err = global.GVA_DB.Create(imageRegistry).Error
	return err
}

// DeleteImageRegistry 删除镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService)DeleteImageRegistry(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&imageregistry.ImageRegistry{},"id = ?",ID).Error
	return err
}

// DeleteImageRegistryByIds 批量删除镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService)DeleteImageRegistryByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]imageregistry.ImageRegistry{},"id in ?",IDs).Error
	return err
}

// UpdateImageRegistry 更新镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService)UpdateImageRegistry(ctx context.Context, imageRegistry imageregistry.ImageRegistry) (err error) {
	err = global.GVA_DB.Model(&imageregistry.ImageRegistry{}).Where("id = ?",imageRegistry.ID).Updates(&imageRegistry).Error
	return err
}

// GetImageRegistry 根据ID获取镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService)GetImageRegistry(ctx context.Context, ID string) (imageRegistry imageregistry.ImageRegistry, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&imageRegistry).Error
	return
}
// GetImageRegistryInfoList 分页获取镜像库记录
// Author [yourname](https://github.com/yourname)
func (imageRegistryService *ImageRegistryService)GetImageRegistryInfoList(ctx context.Context, info imageregistryReq.ImageRegistrySearch) (list []imageregistry.ImageRegistry, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&imageregistry.ImageRegistry{})
    var imageRegistrys []imageregistry.ImageRegistry
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Address != nil && *info.Address != "" {
        db = db.Where("address LIKE ?", "%"+ *info.Address+"%")
    }
    if info.Source != nil && *info.Source != "" {
        db = db.Where("source LIKE ?", "%"+ *info.Source+"%")
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

	err = db.Find(&imageRegistrys).Error
	return  imageRegistrys, total, err
}
func (imageRegistryService *ImageRegistryService)GetImageRegistryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
