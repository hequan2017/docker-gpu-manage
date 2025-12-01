package product

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductSpecRouter struct {}

// InitProductSpecRouter 初始化 产品规格 路由信息
func (s *ProductSpecRouter) InitProductSpecRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	productSpecRouter := Router.Group("productSpec").Use(middleware.OperationRecord())
	productSpecRouterWithoutRecord := Router.Group("productSpec")
	productSpecRouterWithoutAuth := PublicRouter.Group("productSpec")
	{
		productSpecRouter.POST("createProductSpec", productSpecApi.CreateProductSpec)   // 新建产品规格
		productSpecRouter.DELETE("deleteProductSpec", productSpecApi.DeleteProductSpec) // 删除产品规格
		productSpecRouter.DELETE("deleteProductSpecByIds", productSpecApi.DeleteProductSpecByIds) // 批量删除产品规格
		productSpecRouter.PUT("updateProductSpec", productSpecApi.UpdateProductSpec)    // 更新产品规格
	}
	{
		productSpecRouterWithoutRecord.GET("findProductSpec", productSpecApi.FindProductSpec)        // 根据ID获取产品规格
		productSpecRouterWithoutRecord.GET("getProductSpecList", productSpecApi.GetProductSpecList)  // 获取产品规格列表
	}
	{
	    productSpecRouterWithoutAuth.GET("getProductSpecPublic", productSpecApi.GetProductSpecPublic)  // 产品规格开放接口
	}
}
