package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MsModel = new(msModel)

type msModel struct {}

// Init 初始化 模型库 路由信息
func (r *msModel) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("model").Use(middleware.OperationRecord())
		group.POST("createMsModel", apiMsModel.CreateMsModel)   // 新建模型库
		group.DELETE("deleteMsModel", apiMsModel.DeleteMsModel) // 删除模型库
		group.DELETE("deleteMsModelByIds", apiMsModel.DeleteMsModelByIds) // 批量删除模型库
		group.PUT("updateMsModel", apiMsModel.UpdateMsModel)    // 更新模型库
	}
	{
	    group := private.Group("model")
		group.GET("findMsModel", apiMsModel.FindMsModel)        // 根据ID获取模型库
		group.GET("getMsModelList", apiMsModel.GetMsModelList)  // 获取模型库列表
	}
	{
	    group := public.Group("model")
	    group.GET("getMsModelPublic", apiMsModel.GetMsModelPublic)  // 模型库开放接口
	}
}
