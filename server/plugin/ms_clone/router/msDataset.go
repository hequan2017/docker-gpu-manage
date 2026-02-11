package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MsDataset = new(msDataset)

type msDataset struct {}

// Init 初始化 数据集 路由信息
func (r *msDataset) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("dataset").Use(middleware.OperationRecord())
		group.POST("createMsDataset", apiMsDataset.CreateMsDataset)   // 新建数据集
		group.DELETE("deleteMsDataset", apiMsDataset.DeleteMsDataset) // 删除数据集
		group.DELETE("deleteMsDatasetByIds", apiMsDataset.DeleteMsDatasetByIds) // 批量删除数据集
		group.PUT("updateMsDataset", apiMsDataset.UpdateMsDataset)    // 更新数据集
	}
	{
	    group := private.Group("dataset")
		group.GET("findMsDataset", apiMsDataset.FindMsDataset)        // 根据ID获取数据集
		group.GET("getMsDatasetList", apiMsDataset.GetMsDatasetList)  // 获取数据集列表
	}
	{
	    group := public.Group("dataset")
	    group.GET("getMsDatasetPublic", apiMsDataset.GetMsDatasetPublic)  // 数据集开放接口
	}
}
