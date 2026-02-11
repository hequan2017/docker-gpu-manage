import service from '@/utils/request'
// @Tags MsModel
// @Summary 创建模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "创建模型库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /model/createMsModel [post]
export const createMsModel = (data) => {
  return service({
    url: '/model/createMsModel',
    method: 'post',
    data
  })
}

// @Tags MsModel
// @Summary 删除模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "删除模型库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /model/deleteMsModel [delete]
export const deleteMsModel = (params) => {
  return service({
    url: '/model/deleteMsModel',
    method: 'delete',
    params
  })
}

// @Tags MsModel
// @Summary 批量删除模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模型库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /model/deleteMsModel [delete]
export const deleteMsModelByIds = (params) => {
  return service({
    url: '/model/deleteMsModelByIds',
    method: 'delete',
    params
  })
}

// @Tags MsModel
// @Summary 更新模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsModel true "更新模型库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /model/updateMsModel [put]
export const updateMsModel = (data) => {
  return service({
    url: '/model/updateMsModel',
    method: 'put',
    data
  })
}

// @Tags MsModel
// @Summary 用id查询模型库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MsModel true "用id查询模型库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /model/findMsModel [get]
export const findMsModel = (params) => {
  return service({
    url: '/model/findMsModel',
    method: 'get',
    params
  })
}

// @Tags MsModel
// @Summary 分页获取模型库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取模型库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /model/getMsModelList [get]
export const getMsModelList = (params) => {
  return service({
    url: '/model/getMsModelList',
    method: 'get',
    params
  })
}
// @Tags MsModel
// @Summary 不需要鉴权的模型库接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MsModelSearch true "分页获取模型库列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /model/getMsModelPublic [get]
export const getMsModelPublic = () => {
  return service({
    url: '/model/getMsModelPublic',
    method: 'get',
  })
}
