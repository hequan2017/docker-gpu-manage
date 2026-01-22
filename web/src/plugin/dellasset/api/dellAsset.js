import service from '@/utils/request'

// @Tags DellAsset
// @Summary 创建戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DellAsset true "服务器资产信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dellAsset/createDellAsset [post]
export const createDellAsset = (data) => {
  return service({
    url: '/dellAsset/createDellAsset',
    method: 'post',
    data
  })
}

// @Tags DellAsset
// @Summary 删除戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DellAsset true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dellAsset/deleteDellAsset [delete]
export const deleteDellAsset = (params) => {
  return service({
    url: '/dellAsset/deleteDellAsset',
    method: 'delete',
    params
  })
}

// @Tags DellAsset
// @Summary 批量删除戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除戴尔服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dellAsset/deleteDellAssetByIds [delete]
export const deleteDellAssetByIds = (params) => {
  return service({
    url: '/dellAsset/deleteDellAssetByIds',
    method: 'delete',
    params
  })
}

// @Tags DellAsset
// @Summary 更新戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DellAsset true "服务器资产信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dellAsset/updateDellAsset [put]
export const updateDellAsset = (data) => {
  return service({
    url: '/dellAsset/updateDellAsset',
    method: 'put',
    data
  })
}

// @Tags DellAsset
// @Summary 用id查询戴尔服务器资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DellAsset true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dellAsset/findDellAsset [get]
export const findDellAsset = (params) => {
  return service({
    url: '/dellAsset/findDellAsset',
    method: 'get',
    params
  })
}

// @Tags DellAsset
// @Summary 分页获取戴尔服务器资产列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取戴尔服务器资产列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dellAsset/getDellAssetList [get]
export const getDellAssetList = (params) => {
  return service({
    url: '/dellAsset/getDellAssetList',
    method: 'get',
    params
  })
}

// @Tags DellAsset
// @Summary 获取资产统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dellAsset/getStatistics [get]
export const getDellAssetStatistics = () => {
  return service({
    url: '/dellAsset/getStatistics',
    method: 'get'
  })
}
