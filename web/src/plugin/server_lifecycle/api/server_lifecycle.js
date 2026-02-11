import service from '@/utils/request'
// @Tags ServerAsset
// @Summary 创建服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "创建服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /asset/createServerAsset [post]
export const createServerAsset = (data) => {
  return service({
    url: '/asset/createServerAsset',
    method: 'post',
    data
  })
}

// @Tags ServerAsset
// @Summary 删除服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "删除服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset/deleteServerAsset [delete]
export const deleteServerAsset = (params) => {
  return service({
    url: '/asset/deleteServerAsset',
    method: 'delete',
    params
  })
}

// @Tags ServerAsset
// @Summary 批量删除服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset/deleteServerAsset [delete]
export const deleteServerAssetByIds = (params) => {
  return service({
    url: '/asset/deleteServerAssetByIds',
    method: 'delete',
    params
  })
}

// @Tags ServerAsset
// @Summary 更新服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerAsset true "更新服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asset/updateServerAsset [put]
export const updateServerAsset = (data) => {
  return service({
    url: '/asset/updateServerAsset',
    method: 'put',
    data
  })
}

// @Tags ServerAsset
// @Summary 用id查询服务器资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ServerAsset true "用id查询服务器资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset/findServerAsset [get]
export const findServerAsset = (params) => {
  return service({
    url: '/asset/findServerAsset',
    method: 'get',
    params
  })
}

// @Tags ServerAsset
// @Summary 分页获取服务器资产列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取服务器资产列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset/getServerAssetList [get]
export const getServerAssetList = (params) => {
  return service({
    url: '/asset/getServerAssetList',
    method: 'get',
    params
  })
}
// @Tags ServerAsset
// @Summary 不需要鉴权的服务器资产接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ServerAssetSearch true "分页获取服务器资产列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /asset/getServerAssetPublic [get]
export const getServerAssetPublic = () => {
  return service({
    url: '/asset/getServerAssetPublic',
    method: 'get',
  })
}
