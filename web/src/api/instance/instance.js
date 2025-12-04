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

// @Tags Instance
// @Summary 根据产品规格获取可用的算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param specId query int true "产品规格ID"
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /instance/getAvailableNodes [get]
export const getAvailableNodes = (params) => {
  return service({
    url: '/instance/getAvailableNodes',
    method: 'get',
    params
  })
}

// @Tags Instance
// @Summary 启动容器
// @Security ApiKeyAuth
// @Router /instance/startContainer [post]
export const startContainer = (params) => {
  return service({
    url: '/instance/startContainer',
    method: 'post',
    params
  })
}

// @Tags Instance
// @Summary 停止容器
// @Security ApiKeyAuth
// @Router /instance/stopContainer [post]
export const stopContainer = (params) => {
  return service({
    url: '/instance/stopContainer',
    method: 'post',
    params
  })
}

// @Tags Instance
// @Summary 重启容器
// @Security ApiKeyAuth
// @Router /instance/restartContainer [post]
export const restartContainer = (params) => {
  return service({
    url: '/instance/restartContainer',
    method: 'post',
    params
  })
}

// @Tags Instance
// @Summary 获取容器日志
// @Security ApiKeyAuth
// @Router /instance/getContainerLogs [get]
export const getContainerLogs = (params) => {
  return service({
    url: '/instance/getContainerLogs',
    method: 'get',
    params
  })
}


// 获取终端WebSocket地址
export const getTerminalWsUrl = (ID, shell = 'bash') => {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const host = window.location.host
  // 使用 /api 前缀，vite 代理会自动转发到后端
  const baseApi = import.meta.env.VITE_BASE_API || '/api'
  return `${protocol}//${host}${baseApi}/instance/terminal?ID=${ID}&shell=${shell}`
}
