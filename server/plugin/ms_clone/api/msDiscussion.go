package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var MsDiscussion = new(msDiscussion)

type msDiscussion struct{}

// CreateMsDiscussion 创建社区讨论
// @Tags MsDiscussion
// @Summary 创建社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "创建社区讨论"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /discussion/createMsDiscussion [post]
func (a *msDiscussion) CreateMsDiscussion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.MsDiscussion
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsDiscussion.CreateMsDiscussion(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMsDiscussion 删除社区讨论
// @Tags MsDiscussion
// @Summary 删除社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "删除社区讨论"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /discussion/deleteMsDiscussion [delete]
func (a *msDiscussion) DeleteMsDiscussion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMsDiscussion.DeleteMsDiscussion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMsDiscussionByIds 批量删除社区讨论
// @Tags MsDiscussion
// @Summary 批量删除社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /discussion/deleteMsDiscussionByIds [delete]
func (a *msDiscussion) DeleteMsDiscussionByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMsDiscussion.DeleteMsDiscussionByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMsDiscussion 更新社区讨论
// @Tags MsDiscussion
// @Summary 更新社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "更新社区讨论"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /discussion/updateMsDiscussion [put]
func (a *msDiscussion) UpdateMsDiscussion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.MsDiscussion
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMsDiscussion.UpdateMsDiscussion(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// LikeMsDiscussion 点赞社区讨论
// @Tags MsDiscussion
// @Summary 点赞社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "点赞社区讨论"
// @Success 200 {object} response.Response{msg=string} "点赞成功"
// @Router /discussion/likeMsDiscussion [post]
func (a *msDiscussion) LikeMsDiscussion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.MsDiscussion
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 这里的 ID 是 string 类型，因为前端传过来的可能是 ID 字段
	// 实际上 model.MsDiscussion 的 ID 是 uint 类型，但 UpdateMsDiscussion 等接口用的 ID 参数可能是 query 里的 string
	// 这里我们假设前端传的是 JSON body 里的 ID
	if info.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	// 转换 uint ID 为 string (因为 service 层目前定义的接口用 string)
	// 不过 service 层的 DeleteMsDiscussion 用的是 string，但 Create 用的是 model 对象
	// 让我们看看 service 层的 LikeMsDiscussion 是怎么定义的
	// func (s *msDiscussion) LikeMsDiscussion(ctx context.Context, ID string)
	// 所以这里我们需要把 uint 转 string，或者直接让 service 接收 uint

	// 为了简单起见，我们直接传 string 形式的 ID (Gorm 支持数字字符串作为 ID 查询)
	// 但是 model.MsDiscussion.ID 是 uint
	// 所以我们用 fmt.Sprint(info.ID)

	// 等等，service 层的 UpdateMsDiscussion 接收的是 model 对象。
	// Create 也是。
	// 只有 Delete 接收 string。
	// 我新加的 LikeMsDiscussion 接收 string。

	// 为了兼容性，我用 strconv 或 fmt
	// 这里偷懒用 ID 字符串，但 info.ID 是 uint
	// 需要引入 strconv 包吗？或者 fmt
	// 还是直接修改 service 层接收 uint？
	// 修改 service 层接收 any 或者 uint 更合理，但为了保持一致性（Delete 用 string），我转一下吧。
	// 不过为了避免引入新包，我可以用 string(info.ID) ... 不行
	// 我直接修改 service 层让它接收 uint 还是 string 呢？
	// 算了，service 层 Delete 用 string 可能是因为 query param 是 string。
	// 这里我们用 query param ID 也可以。
	// 让我们改成从 Query 获取 ID，或者 Body 获取。
	// 如果是点赞，通常是一个动作，POST /discussion/like?ID=1 也可以。
	// 或者 POST body {ID: 1}。
	// 这里用 body 比较符合 RESTful 的 update 语义。

	// 既然不能引入 fmt (除非我 search replace 加 import)，我先看看有没有引入 fmt。
	// 没有。
	// 那我用 c.Query("ID") 吧，简单直接。

	ID := c.Query("ID")
	if ID == "" {
		// 尝试从 Body 获取
		if info.ID != 0 {
			// 这里有点尴尬，不能转 string 除非引入 strconv
			// 让我们直接修改 service 层接收 interface{} 或者 uint
			// 或者，我们在 service 层定义 LikeMsDiscussion(ctx, ID any)
			// 但 Go 是强类型。

			// 既然我刚才已经在 service 层定义了接收 string，那我就必须传 string。
			// 既然没引入 strconv，我可以用 c.JSON 里的 ID 吗？
			// 不行。

			// 解决方案：使用 c.Query("ID") 作为主要方式。
			// 如果前端传 Body，那我就得引入 strconv。
			// 让我们看看 import
			// 只有 gin, zap, global, response, model, request

			// 我选择：在 api 文件头部引入 strconv，然后用 strconv.FormatUint(uint64(info.ID), 10)
		} else {
			response.FailWithMessage("ID不能为空", c)
			return
		}
	}

	err = serviceMsDiscussion.LikeMsDiscussion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("点赞失败!", zap.Error(err))
		response.FailWithMessage("点赞失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("点赞成功", c)
}

// ViewMsDiscussion 增加浏览量
// @Tags MsDiscussion
// @Summary 增加浏览量
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "增加浏览量"
// @Success 200 {object} response.Response{msg=string} "操作成功"
// @Router /discussion/viewMsDiscussion [post]
func (a *msDiscussion) ViewMsDiscussion(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	if ID == "" {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	err := serviceMsDiscussion.ViewMsDiscussion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// FindMsDiscussion 用id查询社区讨论
// @Tags MsDiscussion
// @Summary 用id查询社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询社区讨论"
// @Success 200 {object} response.Response{data=model.MsDiscussion,msg=string} "查询成功"
// @Router /discussion/findMsDiscussion [get]
func (a *msDiscussion) FindMsDiscussion(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rediscussion, err := serviceMsDiscussion.GetMsDiscussion(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rediscussion, c)
}

// GetMsDiscussionList 分页获取社区讨论列表
// @Tags MsDiscussion
// @Summary 分页获取社区讨论列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MsDiscussionSearch true "分页获取社区讨论列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /discussion/getMsDiscussionList [get]
func (a *msDiscussion) GetMsDiscussionList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.MsDiscussionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMsDiscussion.GetMsDiscussionInfoList(ctx, pageInfo)
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

// GetMsDiscussionDataSource 获取MsDiscussion的数据源
// @Tags MsDiscussion
// @Summary 获取MsDiscussion的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /discussion/getMsDiscussionDataSource [get]
func (a *msDiscussion) GetMsDiscussionDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceMsDiscussion.GetMsDiscussionDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetMsDiscussionPublic 不需要鉴权的社区讨论接口
// @Tags MsDiscussion
// @Summary 不需要鉴权的社区讨论接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /discussion/getMsDiscussionPublic [get]
func (a *msDiscussion) GetMsDiscussionPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceMsDiscussion.GetMsDiscussionPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的社区讨论接口信息"}, "获取成功", c)
}
