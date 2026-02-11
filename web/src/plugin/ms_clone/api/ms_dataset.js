import service from '@/utils/request'
// @Tags MsDataset
// @Summary 创建数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "创建数据集"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataset/createMsDataset [post]
export const createMsDataset = (data) => {
  return service({
    url: '/dataset/createMsDataset',
    method: 'post',
    data
  })
}

// @Tags MsDataset
// @Summary 删除数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "删除数据集"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataset/deleteMsDataset [delete]
export const deleteMsDataset = (params) => {
  return service({
    url: '/dataset/deleteMsDataset',
    method: 'delete',
    params
  })
}

// @Tags MsDataset
// @Summary 批量删除数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据集"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataset/deleteMsDataset [delete]
export const deleteMsDatasetByIds = (params) => {
  return service({
    url: '/dataset/deleteMsDatasetByIds',
    method: 'delete',
    params
  })
}

// @Tags MsDataset
// @Summary 更新数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MsDataset true "更新数据集"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataset/updateMsDataset [put]
export const updateMsDataset = (data) => {
  return service({
    url: '/dataset/updateMsDataset',
    method: 'put',
    data
  })
}

// @Tags MsDataset
// @Summary 用id查询数据集
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MsDataset true "用id查询数据集"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataset/findMsDataset [get]
export const findMsDataset = (params) => {
  return service({
    url: '/dataset/findMsDataset',
    method: 'get',
    params
  })
}

// @Tags MsDataset
// @Summary 分页获取数据集列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据集列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataset/getMsDatasetList [get]
export const getMsDatasetList = (params) => {
  return service({
    url: '/dataset/getMsDatasetList',
    method: 'get',
    params
  })
}
// @Tags MsDataset
// @Summary 不需要鉴权的数据集接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MsDatasetSearch true "分页获取数据集列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dataset/getMsDatasetPublic [get]
export const getMsDatasetPublic = () => {
  return service({
    url: '/dataset/getMsDatasetPublic',
    method: 'get',
  })
}
