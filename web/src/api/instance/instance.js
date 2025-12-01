import service from '@/utils/request'
// @Tags Instance
// @Summary 创建实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instance true "创建实例管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /instance/createInstance [post]
export const createInstance = (data) => {
  return service({
    url: '/instance/createInstance',
    method: 'post',
    data
  })
}

// @Tags Instance
// @Summary 删除实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instance true "删除实例管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /instance/deleteInstance [delete]
export const deleteInstance = (params) => {
  return service({
    url: '/instance/deleteInstance',
    method: 'delete',
    params
  })
}

// @Tags Instance
// @Summary 批量删除实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除实例管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /instance/deleteInstance [delete]
export const deleteInstanceByIds = (params) => {
  return service({
    url: '/instance/deleteInstanceByIds',
    method: 'delete',
    params
  })
}

// @Tags Instance
// @Summary 更新实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instance true "更新实例管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /instance/updateInstance [put]
export const updateInstance = (data) => {
  return service({
    url: '/instance/updateInstance',
    method: 'put',
    data
  })
}

// @Tags Instance
// @Summary 用id查询实例管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Instance true "用id查询实例管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /instance/findInstance [get]
export const findInstance = (params) => {
  return service({
    url: '/instance/findInstance',
    method: 'get',
    params
  })
}

// @Tags Instance
// @Summary 分页获取实例管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取实例管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /instance/getInstanceList [get]
export const getInstanceList = (params) => {
  return service({
    url: '/instance/getInstanceList',
    method: 'get',
    params
  })
}
// @Tags Instance
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /instance/findInstanceDataSource [get]
export const getInstanceDataSource = () => {
  return service({
    url: '/instance/getInstanceDataSource',
    method: 'get',
  })
}

// @Tags Instance
// @Summary 不需要鉴权的实例管理接口
// @Accept application/json
// @Produce application/json
// @Param data query instanceReq.InstanceSearch true "分页获取实例管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /instance/getInstancePublic [get]
export const getInstancePublic = () => {
  return service({
    url: '/instance/getInstancePublic',
    method: 'get',
  })
}
