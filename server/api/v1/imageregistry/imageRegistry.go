package imageregistry

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry"
    imageregistryReq "github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ImageRegistryApi struct {}



// CreateImageRegistry 创建镜像库
// @Tags ImageRegistry
// @Summary 创建镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body imageregistry.ImageRegistry true "创建镜像库"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /imageRegistry/createImageRegistry [post]
func (imageRegistryApi *ImageRegistryApi) CreateImageRegistry(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var imageRegistry imageregistry.ImageRegistry
	err := c.ShouldBindJSON(&imageRegistry)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = imageRegistryService.CreateImageRegistry(ctx,&imageRegistry)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteImageRegistry 删除镜像库
// @Tags ImageRegistry
// @Summary 删除镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body imageregistry.ImageRegistry true "删除镜像库"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /imageRegistry/deleteImageRegistry [delete]
func (imageRegistryApi *ImageRegistryApi) DeleteImageRegistry(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := imageRegistryService.DeleteImageRegistry(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteImageRegistryByIds 批量删除镜像库
// @Tags ImageRegistry
// @Summary 批量删除镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /imageRegistry/deleteImageRegistryByIds [delete]
func (imageRegistryApi *ImageRegistryApi) DeleteImageRegistryByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := imageRegistryService.DeleteImageRegistryByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateImageRegistry 更新镜像库
// @Tags ImageRegistry
// @Summary 更新镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body imageregistry.ImageRegistry true "更新镜像库"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /imageRegistry/updateImageRegistry [put]
func (imageRegistryApi *ImageRegistryApi) UpdateImageRegistry(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var imageRegistry imageregistry.ImageRegistry
	err := c.ShouldBindJSON(&imageRegistry)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = imageRegistryService.UpdateImageRegistry(ctx,imageRegistry)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindImageRegistry 用id查询镜像库
// @Tags ImageRegistry
// @Summary 用id查询镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询镜像库"
// @Success 200 {object} response.Response{data=imageregistry.ImageRegistry,msg=string} "查询成功"
// @Router /imageRegistry/findImageRegistry [get]
func (imageRegistryApi *ImageRegistryApi) FindImageRegistry(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reimageRegistry, err := imageRegistryService.GetImageRegistry(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reimageRegistry, c)
}
// GetImageRegistryList 分页获取镜像库列表
// @Tags ImageRegistry
// @Summary 分页获取镜像库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query imageregistryReq.ImageRegistrySearch true "分页获取镜像库列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /imageRegistry/getImageRegistryList [get]
func (imageRegistryApi *ImageRegistryApi) GetImageRegistryList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo imageregistryReq.ImageRegistrySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := imageRegistryService.GetImageRegistryInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetImageRegistryPublic 不需要鉴权的镜像库接口
// @Tags ImageRegistry
// @Summary 不需要鉴权的镜像库接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /imageRegistry/getImageRegistryPublic [get]
func (imageRegistryApi *ImageRegistryApi) GetImageRegistryPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    imageRegistryService.GetImageRegistryPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的镜像库接口信息",
    }, "获取成功", c)
}
