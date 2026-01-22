package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset/api/v1"
	"github.com/gin-gonic/gin"
)

type DellAssetRouter struct{}

// InitDellAssetRouter 初始化戴尔资产路由
func (dar *DellAssetRouter) InitDellAssetRouter(Router *gin.RouterGroup) {
	dellAssetApi := v1.ApiGroupApp.DellAssetApi
	dellAssetRouter := Router.Group("dellAsset")
	{
		dellAssetRouter.POST("createDellAsset", dellAssetApi.CreateDellAsset)           // 创建戴尔资产
		dellAssetRouter.DELETE("deleteDellAsset", dellAssetApi.DeleteDellAsset)        // 删除戴尔资产
		dellAssetRouter.DELETE("deleteDellAssetByIds", dellAssetApi.DeleteDellAssetByIds) // 批量删除戴尔资产
		dellAssetRouter.PUT("updateDellAsset", dellAssetApi.UpdateDellAsset)           // 更新戴尔资产
		dellAssetRouter.GET("findDellAsset", dellAssetApi.FindDellAsset)              // 根据ID获取戴尔资产
		dellAssetRouter.GET("getDellAssetList", dellAssetApi.GetDellAssetList)        // 获取戴尔资产列表
		dellAssetRouter.GET("getStatistics", dellAssetApi.GetDellAssetStatistics)     // 获取统计信息
	}
}
