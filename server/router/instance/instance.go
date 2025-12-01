package instance

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InstanceRouter struct {}

// InitInstanceRouter 初始化 实例管理 路由信息
func (s *InstanceRouter) InitInstanceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	instanceRouter := Router.Group("instance").Use(middleware.OperationRecord())
	instanceRouterWithoutRecord := Router.Group("instance")
	instanceRouterWithoutAuth := PublicRouter.Group("instance")
	{
		instanceRouter.POST("createInstance", instanceApi.CreateInstance)   // 新建实例管理
		instanceRouter.DELETE("deleteInstance", instanceApi.DeleteInstance) // 删除实例管理
		instanceRouter.DELETE("deleteInstanceByIds", instanceApi.DeleteInstanceByIds) // 批量删除实例管理
		instanceRouter.PUT("updateInstance", instanceApi.UpdateInstance)    // 更新实例管理
	}
	{
		instanceRouterWithoutRecord.GET("findInstance", instanceApi.FindInstance)        // 根据ID获取实例管理
		instanceRouterWithoutRecord.GET("getInstanceList", instanceApi.GetInstanceList)  // 获取实例管理列表
	}
	{
	    instanceRouterWithoutAuth.GET("getInstanceDataSource", instanceApi.GetInstanceDataSource)  // 获取实例管理数据源
	    instanceRouterWithoutAuth.GET("getInstancePublic", instanceApi.GetInstancePublic)  // 实例管理开放接口
	}
}
