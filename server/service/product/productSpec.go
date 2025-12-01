
package product

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
    productReq "github.com/flipped-aurora/gin-vue-admin/server/model/product/request"
)

type ProductSpecService struct {}
// CreateProductSpec 创建产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService) CreateProductSpec(ctx context.Context, productSpec *product.ProductSpec) (err error) {
	err = global.GVA_DB.Create(productSpec).Error
	return err
}

// DeleteProductSpec 删除产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService)DeleteProductSpec(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&product.ProductSpec{},"id = ?",ID).Error
	return err
}

// DeleteProductSpecByIds 批量删除产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService)DeleteProductSpecByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]product.ProductSpec{},"id in ?",IDs).Error
	return err
}

// UpdateProductSpec 更新产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService)UpdateProductSpec(ctx context.Context, productSpec product.ProductSpec) (err error) {
	err = global.GVA_DB.Model(&product.ProductSpec{}).Where("id = ?",productSpec.ID).Updates(&productSpec).Error
	return err
}

// GetProductSpec 根据ID获取产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService)GetProductSpec(ctx context.Context, ID string) (productSpec product.ProductSpec, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&productSpec).Error
	return
}
// GetProductSpecInfoList 分页获取产品规格记录
// Author [yourname](https://github.com/yourname)
func (productSpecService *ProductSpecService)GetProductSpecInfoList(ctx context.Context, info productReq.ProductSpecSearch) (list []product.ProductSpec, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&product.ProductSpec{})
    var productSpecs []product.ProductSpec
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.GpuModel != nil && *info.GpuModel != "" {
        db = db.Where("gpu_model LIKE ?", "%"+ *info.GpuModel+"%")
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

	err = db.Find(&productSpecs).Error
	return  productSpecs, total, err
}
func (productSpecService *ProductSpecService)GetProductSpecPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
