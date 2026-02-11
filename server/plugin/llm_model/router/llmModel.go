package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var LlmModel = new(llm)

type llm struct {}

// Init 初始化 开源大模型 路由信息
func (r *llm) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("llm").Use(middleware.OperationRecord())
		group.POST("createLlmModel", apiLlmModel.CreateLlmModel)   // 新建开源大模型
		group.DELETE("deleteLlmModel", apiLlmModel.DeleteLlmModel) // 删除开源大模型
		group.DELETE("deleteLlmModelByIds", apiLlmModel.DeleteLlmModelByIds) // 批量删除开源大模型
		group.PUT("updateLlmModel", apiLlmModel.UpdateLlmModel)    // 更新开源大模型
	}
	{
	    group := private.Group("llm")
		group.GET("findLlmModel", apiLlmModel.FindLlmModel)        // 根据ID获取开源大模型
		group.GET("getLlmModelList", apiLlmModel.GetLlmModelList)  // 获取开源大模型列表
	}
	{
	    group := public.Group("llm")
	    group.GET("getLlmModelPublic", apiLlmModel.GetLlmModelPublic)  // 开源大模型开放接口
	}
}
