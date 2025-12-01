package imageregistry

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ImageRegistryRouter struct {}

// InitImageRegistryRouter 初始化 镜像库 路由信息
func (s *ImageRegistryRouter) InitImageRegistryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	imageRegistryRouter := Router.Group("imageRegistry").Use(middleware.OperationRecord())
	imageRegistryRouterWithoutRecord := Router.Group("imageRegistry")
	imageRegistryRouterWithoutAuth := PublicRouter.Group("imageRegistry")
	{
		imageRegistryRouter.POST("createImageRegistry", imageRegistryApi.CreateImageRegistry)   // 新建镜像库
		imageRegistryRouter.DELETE("deleteImageRegistry", imageRegistryApi.DeleteImageRegistry) // 删除镜像库
		imageRegistryRouter.DELETE("deleteImageRegistryByIds", imageRegistryApi.DeleteImageRegistryByIds) // 批量删除镜像库
		imageRegistryRouter.PUT("updateImageRegistry", imageRegistryApi.UpdateImageRegistry)    // 更新镜像库
	}
	{
		imageRegistryRouterWithoutRecord.GET("findImageRegistry", imageRegistryApi.FindImageRegistry)        // 根据ID获取镜像库
		imageRegistryRouterWithoutRecord.GET("getImageRegistryList", imageRegistryApi.GetImageRegistryList)  // 获取镜像库列表
	}
	{
	    imageRegistryRouterWithoutAuth.GET("getImageRegistryPublic", imageRegistryApi.GetImageRegistryPublic)  // 镜像库开放接口
	}
}
