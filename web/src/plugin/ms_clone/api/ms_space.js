import service from '@/utils/request'
// @Tags MsSpace
// @Summary 创建创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "创建创空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /space/createMsSpace [post]
export const createMsSpace = (data) => {
  return service({
    url: '/space/createMsSpace',
    method: 'post',
    data
  })
}

// @Tags MsSpace
// @Summary 删除创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "删除创空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /space/deleteMsSpace [delete]
export const deleteMsSpace = (params) => {
  return service({
    url: '/space/deleteMsSpace',
    method: 'delete',
    params
  })
}

// @Tags MsSpace
// @Summary 批量删除创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除创空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /space/deleteMsSpace [delete]
export const deleteMsSpaceByIds = (params) => {
  return service({
    url: '/space/deleteMsSpaceByIds',
    method: 'delete',
    params
  })
}

// @Tags MsSpace
// @Summary 更新创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsSpace true "更新创空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /space/updateMsSpace [put]
export const updateMsSpace = (data) => {
  return service({
    url: '/space/updateMsSpace',
    method: 'put',
    data
  })
}

// @Tags MsSpace
// @Summary 用id查询创空间
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MsSpace true "用id查询创空间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /space/findMsSpace [get]
export const findMsSpace = (params) => {
  return service({
    url: '/space/findMsSpace',
    method: 'get',
    params
  })
}

// @Tags MsSpace
// @Summary 分页获取创空间列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取创空间列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /space/getMsSpaceList [get]
export const getMsSpaceList = (params) => {
  return service({
    url: '/space/getMsSpaceList',
    method: 'get',
    params
  })
}
// @Tags MsSpace
// @Summary 不需要鉴权的创空间接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MsSpaceSearch true "分页获取创空间列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /space/getMsSpacePublic [get]
export const getMsSpacePublic = () => {
  return service({
    url: '/space/getMsSpacePublic',
    method: 'get',
  })
}
