package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ServerAsset = new(asset)

type asset struct {}

// Init 初始化 服务器资产 路由信息
func (r *asset) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("asset").Use(middleware.OperationRecord())
		group.POST("createServerAsset", apiServerAsset.CreateServerAsset)   // 新建服务器资产
		group.DELETE("deleteServerAsset", apiServerAsset.DeleteServerAsset) // 删除服务器资产
		group.DELETE("deleteServerAssetByIds", apiServerAsset.DeleteServerAssetByIds) // 批量删除服务器资产
		group.PUT("updateServerAsset", apiServerAsset.UpdateServerAsset)    // 更新服务器资产
	}
	{
	    group := private.Group("asset")
		group.GET("findServerAsset", apiServerAsset.FindServerAsset)        // 根据ID获取服务器资产
		group.GET("getServerAssetList", apiServerAsset.GetServerAssetList)  // 获取服务器资产列表
	}
	{
	    group := public.Group("asset")
	    group.GET("getServerAssetPublic", apiServerAsset.GetServerAssetPublic)  // 服务器资产开放接口
	}
}
