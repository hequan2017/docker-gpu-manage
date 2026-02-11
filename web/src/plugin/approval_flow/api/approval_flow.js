import service from '@/utils/request'
// @Tags ApprovalProcess
// @Summary 创建发版申请
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApprovalProcess true "创建发版申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /approval/createApprovalProcess [post]
export const createApprovalProcess = (data) => {
  return service({
    url: '/approval/createApprovalProcess',
    method: 'post',
    data
  })
}

// @Tags ApprovalProcess
// @Summary 删除发版申请
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApprovalProcess true "删除发版申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /approval/deleteApprovalProcess [delete]
export const deleteApprovalProcess = (params) => {
  return service({
    url: '/approval/deleteApprovalProcess',
    method: 'delete',
    params
  })
}

// @Tags ApprovalProcess
// @Summary 批量删除发版申请
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除发版申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /approval/deleteApprovalProcess [delete]
export const deleteApprovalProcessByIds = (params) => {
  return service({
    url: '/approval/deleteApprovalProcessByIds',
    method: 'delete',
    params
  })
}

// @Tags ApprovalProcess
// @Summary 更新发版申请
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApprovalProcess true "更新发版申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /approval/updateApprovalProcess [put]
export const updateApprovalProcess = (data) => {
  return service({
    url: '/approval/updateApprovalProcess',
    method: 'put',
    data
  })
}

// @Tags ApprovalProcess
// @Summary 用id查询发版申请
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ApprovalProcess true "用id查询发版申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /approval/findApprovalProcess [get]
export const findApprovalProcess = (params) => {
  return service({
    url: '/approval/findApprovalProcess',
    method: 'get',
    params
  })
}

// @Tags ApprovalProcess
// @Summary 分页获取发版申请列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取发版申请列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /approval/getApprovalProcessList [get]
export const getApprovalProcessList = (params) => {
  return service({
    url: '/approval/getApprovalProcessList',
    method: 'get',
    params
  })
}
// @Tags ApprovalProcess
// @Summary 不需要鉴权的发版申请接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ApprovalProcessSearch true "分页获取发版申请列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /approval/getApprovalProcessPublic [get]
export const getApprovalProcessPublic = () => {
  return service({
    url: '/approval/getApprovalProcessPublic',
    method: 'get',
  })
}
