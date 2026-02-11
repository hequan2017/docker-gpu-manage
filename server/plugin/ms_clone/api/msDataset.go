package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var MsDataset = new(msDataset)

type msDataset struct {}

// CreateMsDataset 创建数据集
// @Tags MsDataset
// @Summary 创建数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "创建数据集"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dataset/createMsDataset [post]
func (a *msDataset) CreateMsDataset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsDataset
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsDataset.CreateMsDataset(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMsDataset 删除数据集
// @Tags MsDataset
// @Summary 删除数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "删除数据集"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dataset/deleteMsDataset [delete]
func (a *msDataset) DeleteMsDataset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMsDataset.DeleteMsDataset(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMsDatasetByIds 批量删除数据集
// @Tags MsDataset
// @Summary 批量删除数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dataset/deleteMsDatasetByIds [delete]
func (a *msDataset) DeleteMsDatasetByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMsDataset.DeleteMsDatasetByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMsDataset 更新数据集
// @Tags MsDataset
// @Summary 更新数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "更新数据集"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dataset/updateMsDataset [put]
func (a *msDataset) UpdateMsDataset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MsDataset
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsDataset.UpdateMsDataset(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMsDataset 用id查询数据集
// @Tags MsDataset
// @Summary 用id查询数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据集"
// @Success 200 {object} response.Response{data=model.MsDataset,msg=string} "查询成功"
// @Router /dataset/findMsDataset [get]
func (a *msDataset) FindMsDataset(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redataset, err := serviceMsDataset.GetMsDataset(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(redataset, c)
}
// GetMsDatasetList 分页获取数据集列表
// @Tags MsDataset
// @Summary 分页获取数据集列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MsDatasetSearch true "分页获取数据集列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dataset/getMsDatasetList [get]
func (a *msDataset) GetMsDatasetList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MsDatasetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMsDataset.GetMsDatasetInfoList(ctx,pageInfo)
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
// GetMsDatasetPublic 不需要鉴权的数据集接口
// @Tags MsDataset
// @Summary 不需要鉴权的数据集接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dataset/getMsDatasetPublic [get]
func (a *msDataset) GetMsDatasetPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMsDataset.GetMsDatasetPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的数据集接口信息"}, "获取成功", c)
}
