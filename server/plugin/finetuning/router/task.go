package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/api"
	"github.com/gin-gonic/gin"
)

type FinetuningTaskRouter struct{}

var apiGroupApp = api.ApiGroupApp

// InitFinetuningTaskRouter 初始化微调任务路由信息
func (r *FinetuningTaskRouter) InitFinetuningTaskRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("finetuning").Use(middleware.OperationRecord())
		group.POST("createTask", apiGroupApp.FinetuningTaskApi.CreateFinetuningTask)  // 创建微调任务
		group.DELETE("deleteTask", apiGroupApp.FinetuningTaskApi.DeleteFinetuningTask) // 删除微调任务
		group.POST("stopTask", apiGroupApp.FinetuningTaskApi.StopFinetuningTask)       // 停止微调任务
	}
	{
		group := private.Group("finetuning")
		group.GET("getTask", apiGroupApp.FinetuningTaskApi.GetFinetuningTask)       // 根据ID获取微调任务
		group.GET("getTaskList", apiGroupApp.FinetuningTaskApi.GetFinetuningTaskList) // 获取微调任务列表
		group.GET("getTaskLog", apiGroupApp.FinetuningTaskApi.GetFinetuningTaskLog)    // 获取任务日志
	}
}
