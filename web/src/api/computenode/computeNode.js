import service from '@/utils/request'
// @Tags ComputeNode
// @Summary 创建算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ComputeNode true "创建算力节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /computeNode/createComputeNode [post]
export const createComputeNode = (data) => {
  return service({
    url: '/computeNode/createComputeNode',
    method: 'post',
    data
  })
}

// @Tags ComputeNode
// @Summary 删除算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ComputeNode true "删除算力节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /computeNode/deleteComputeNode [delete]
export const deleteComputeNode = (params) => {
  return service({
    url: '/computeNode/deleteComputeNode',
    method: 'delete',
    params
  })
}

// @Tags ComputeNode
// @Summary 批量删除算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除算力节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /computeNode/deleteComputeNode [delete]
export const deleteComputeNodeByIds = (params) => {
  return service({
    url: '/computeNode/deleteComputeNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags ComputeNode
// @Summary 更新算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ComputeNode true "更新算力节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /computeNode/updateComputeNode [put]
export const updateComputeNode = (data) => {
  return service({
    url: '/computeNode/updateComputeNode',
    method: 'put',
    data
  })
}

// @Tags ComputeNode
// @Summary 用id查询算力节点
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ComputeNode true "用id查询算力节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /computeNode/findComputeNode [get]
export const findComputeNode = (params) => {
  return service({
    url: '/computeNode/findComputeNode',
    method: 'get',
    params
  })
}

// @Tags ComputeNode
// @Summary 分页获取算力节点列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取算力节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /computeNode/getComputeNodeList [get]
export const getComputeNodeList = (params) => {
  return service({
    url: '/computeNode/getComputeNodeList',
    method: 'get',
    params
  })
}

// @Tags ComputeNode
// @Summary 不需要鉴权的算力节点接口
// @Accept application/json
// @Produce application/json
// @Param data query computenodeReq.ComputeNodeSearch true "分页获取算力节点列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /computeNode/getComputeNodePublic [get]
export const getComputeNodePublic = () => {
  return service({
    url: '/computeNode/getComputeNodePublic',
    method: 'get',
  })
}
