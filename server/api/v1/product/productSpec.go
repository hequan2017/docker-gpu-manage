package product

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/product"
    productReq "github.com/flipped-aurora/gin-vue-admin/server/model/product/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ProductSpecApi struct {}



// CreateProductSpec 创建产品规格
// @Tags ProductSpec
// @Summary 创建产品规格
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body product.ProductSpec true "创建产品规格"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /productSpec/createProductSpec [post]
func (productSpecApi *ProductSpecApi) CreateProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var productSpec product.ProductSpec
	err := c.ShouldBindJSON(&productSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productSpecService.CreateProductSpec(ctx,&productSpec)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProductSpec 删除产品规格
// @Tags ProductSpec
// @Summary 删除产品规格
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body product.ProductSpec true "删除产品规格"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /productSpec/deleteProductSpec [delete]
func (productSpecApi *ProductSpecApi) DeleteProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := productSpecService.DeleteProductSpec(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductSpecByIds 批量删除产品规格
// @Tags ProductSpec
// @Summary 批量删除产品规格
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /productSpec/deleteProductSpecByIds [delete]
func (productSpecApi *ProductSpecApi) DeleteProductSpecByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := productSpecService.DeleteProductSpecByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProductSpec 更新产品规格
// @Tags ProductSpec
// @Summary 更新产品规格
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body product.ProductSpec true "更新产品规格"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /productSpec/updateProductSpec [put]
func (productSpecApi *ProductSpecApi) UpdateProductSpec(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var productSpec product.ProductSpec
	err := c.ShouldBindJSON(&productSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productSpecService.UpdateProductSpec(ctx,productSpec)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProductSpec 用id查询产品规格
// @Tags ProductSpec
// @Summary 用id查询产品规格
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询产品规格"
// @Success 200 {object} response.Response{data=product.ProductSpec,msg=string} "查询成功"
// @Router /productSpec/findProductSpec [get]
func (productSpecApi *ProductSpecApi) FindProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reproductSpec, err := productSpecService.GetProductSpec(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reproductSpec, c)
}
// GetProductSpecList 分页获取产品规格列表
// @Tags ProductSpec
// @Summary 分页获取产品规格列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query productReq.ProductSpecSearch true "分页获取产品规格列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /productSpec/getProductSpecList [get]
func (productSpecApi *ProductSpecApi) GetProductSpecList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo productReq.ProductSpecSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := productSpecService.GetProductSpecInfoList(ctx,pageInfo)
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

// GetProductSpecPublic 不需要鉴权的产品规格接口
// @Tags ProductSpec
// @Summary 不需要鉴权的产品规格接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /productSpec/getProductSpecPublic [get]
func (productSpecApi *ProductSpecApi) GetProductSpecPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    productSpecService.GetProductSpecPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的产品规格接口信息",
    }, "获取成功", c)
}
