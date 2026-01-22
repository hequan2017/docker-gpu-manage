import service from '@/utils/request'

// @Tags PortForward
// @Summary 创建端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortForward true "源IP、源端口、协议、目标IP、目标端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /portForward/createPortForward [post]
export const createPortForward = (data) => {
  return service({
    url: '/portForward/createPortForward',
    method: 'post',
    data
  })
}

// @Tags PortForward
// @Summary 删除端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortForward true "删除端口转发规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /portForward/deletePortForward [delete]
export const deletePortForward = (params) => {
  return service({
    url: '/portForward/deletePortForward',
    method: 'delete',
    params
  })
}

// @Tags PortForward
// @Summary 批量删除端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除端口转发规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /portForward/deletePortForwardByIds [delete]
export const deletePortForwardByIds = (params) => {
  return service({
    url: '/portForward/deletePortForwardByIds',
    method: 'delete',
    params
  })
}

// @Tags PortForward
// @Summary 更新端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PortForward true "更新端口转发规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /portForward/updatePortForward [put]
export const updatePortForward = (data) => {
  return service({
    url: '/portForward/updatePortForward',
    method: 'put',
    data
  })
}

// @Tags PortForward
// @Summary 用id查询端口转发规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortForward true "用id查询端口转发规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /portForward/findPortForward [get]
export const findPortForward = (params) => {
  return service({
    url: '/portForward/findPortForward',
    method: 'get',
    params
  })
}

// @Tags PortForward
// @Summary 分页获取端口转发规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取端口转发规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /portForward/getPortForwardList [get]
export const getPortForwardList = (params) => {
  return service({
    url: '/portForward/getPortForwardList',
    method: 'get',
    params
  })
}

// @Tags PortForward
// @Summary 更新端口转发规则状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PortForward true "更新端口转发规则状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /portForward/updatePortForwardStatus [put]
export const updatePortForwardStatus = (data) => {
  return service({
    url: '/portForward/updatePortForwardStatus',
    method: 'put',
    data
  })
}

// @Tags PortForward
// @Summary 获取服务器IP地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /portForward/getServerIP [get]
export const getServerIP = () => {
  return service({
    url: '/portForward/getServerIP',
    method: 'get'
  })
}

// @Tags PortForward
// @Summary 获取端口转发运行状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query string true "规则ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /portForward/getForwarderStatus [get]
export const getForwarderStatus = (params) => {
  return service({
    url: '/portForward/getForwarderStatus',
    method: 'get',
    params
  })
}

// @Tags PortForward
// @Summary 获取所有端口转发运行状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /portForward/getAllForwarderStatus [get]
export const getAllForwarderStatus = () => {
  return service({
    url: '/portForward/getAllForwarderStatus',
    method: 'get'
  })
}
