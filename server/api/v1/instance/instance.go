package instance

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	instanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/instance/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InstanceApi struct{}

// CreateInstance 创建实例管理
// @Tags Instance
// @Summary 创建实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body instanceModel.Instance true "创建实例管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /instance/createInstance [post]
func (instanceApi *InstanceApi) CreateInstance(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var inst instanceModel.Instance
	err := c.ShouldBindJSON(&inst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = instanceService.CreateInstance(ctx, &inst)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteInstance 删除实例管理
// @Tags Instance
// @Summary 删除实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body instanceModel.Instance true "删除实例管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /instance/deleteInstance [delete]
func (instanceApi *InstanceApi) DeleteInstance(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := instanceService.DeleteInstance(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInstanceByIds 批量删除实例管理
// @Tags Instance
// @Summary 批量删除实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /instance/deleteInstanceByIds [delete]
func (instanceApi *InstanceApi) DeleteInstanceByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := instanceService.DeleteInstanceByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInstance 更新实例管理
// @Tags Instance
// @Summary 更新实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body instanceModel.Instance true "更新实例管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /instance/updateInstance [put]
func (instanceApi *InstanceApi) UpdateInstance(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var inst instanceModel.Instance
	err := c.ShouldBindJSON(&inst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = instanceService.UpdateInstance(ctx, inst)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInstance 用id查询实例管理
// @Tags Instance
// @Summary 用id查询实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询实例管理"
// @Success 200 {object} response.Response{data=instanceModel.Instance,msg=string} "查询成功"
// @Router /instance/findInstance [get]
func (instanceApi *InstanceApi) FindInstance(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reinstance, err := instanceService.GetInstance(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinstance, c)
}

// GetInstanceList 分页获取实例管理列表
// @Tags Instance
// @Summary 分页获取实例管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query instanceReq.InstanceSearch true "分页获取实例管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /instance/getInstanceList [get]
func (instanceApi *InstanceApi) GetInstanceList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo instanceReq.InstanceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := instanceService.GetInstanceInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetInstanceDataSource 获取Instance的数据源
// @Tags Instance
// @Summary 获取Instance的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /instance/getInstanceDataSource [get]
func (instanceApi *InstanceApi) GetInstanceDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := instanceService.GetInstanceDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetInstancePublic 不需要鉴权的实例管理接口
// @Tags Instance
// @Summary 不需要鉴权的实例管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /instance/getInstancePublic [get]
func (instanceApi *InstanceApi) GetInstancePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	instanceService.GetInstancePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的实例管理接口信息",
	}, "获取成功", c)
}
