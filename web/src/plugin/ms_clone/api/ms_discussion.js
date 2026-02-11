import service from '@/utils/request'
// @Tags MsDiscussion
// @Summary 创建社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "创建社区讨论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /discussion/createMsDiscussion [post]
export const createMsDiscussion = (data) => {
  return service({
    url: '/discussion/createMsDiscussion',
    method: 'post',
    data
  })
}

// @Tags MsDiscussion
// @Summary 删除社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "删除社区讨论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /discussion/deleteMsDiscussion [delete]
export const deleteMsDiscussion = (params) => {
  return service({
    url: '/discussion/deleteMsDiscussion',
    method: 'delete',
    params
  })
}

// @Tags MsDiscussion
// @Summary 批量删除社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除社区讨论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /discussion/deleteMsDiscussion [delete]
export const deleteMsDiscussionByIds = (params) => {
  return service({
    url: '/discussion/deleteMsDiscussionByIds',
    method: 'delete',
    params
  })
}

// @Tags MsDiscussion
// @Summary 更新社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDiscussion true "更新社区讨论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /discussion/updateMsDiscussion [put]
export const updateMsDiscussion = (data) => {
  return service({
    url: '/discussion/updateMsDiscussion',
    method: 'put',
    data
  })
}

// @Tags MsDiscussion
// @Summary 用id查询社区讨论
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MsDiscussion true "用id查询社区讨论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /discussion/findMsDiscussion [get]
export const findMsDiscussion = (params) => {
  return service({
    url: '/discussion/findMsDiscussion',
    method: 'get',
    params
  })
}

// @Tags MsDiscussion
// @Summary 分页获取社区讨论列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取社区讨论列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /discussion/getMsDiscussionList [get]
export const getMsDiscussionList = (params) => {
  return service({
    url: '/discussion/getMsDiscussionList',
    method: 'get',
    params
  })
}
// @Tags MsDiscussion
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /discussion/findMsDiscussionDataSource [get]
export const getMsDiscussionDataSource = () => {
  return service({
    url: '/discussion/getMsDiscussionDataSource',
    method: 'get',
  })
}
// @Tags MsDiscussion
// @Summary 不需要鉴权的社区讨论接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MsDiscussionSearch true "分页获取社区讨论列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /discussion/getMsDiscussionPublic [get]
export const getMsDiscussionPublic = () => {
  return service({
    url: '/discussion/getMsDiscussionPublic',
    method: 'get',
  })
}
