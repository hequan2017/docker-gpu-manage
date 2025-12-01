import service from '@/utils/request'
// @Tags ImageRegistry
// @Summary 创建镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ImageRegistry true "创建镜像库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /imageRegistry/createImageRegistry [post]
export const createImageRegistry = (data) => {
  return service({
    url: '/imageRegistry/createImageRegistry',
    method: 'post',
    data
  })
}

// @Tags ImageRegistry
// @Summary 删除镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ImageRegistry true "删除镜像库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imageRegistry/deleteImageRegistry [delete]
export const deleteImageRegistry = (params) => {
  return service({
    url: '/imageRegistry/deleteImageRegistry',
    method: 'delete',
    params
  })
}

// @Tags ImageRegistry
// @Summary 批量删除镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除镜像库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /imageRegistry/deleteImageRegistry [delete]
export const deleteImageRegistryByIds = (params) => {
  return service({
    url: '/imageRegistry/deleteImageRegistryByIds',
    method: 'delete',
    params
  })
}

// @Tags ImageRegistry
// @Summary 更新镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ImageRegistry true "更新镜像库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /imageRegistry/updateImageRegistry [put]
export const updateImageRegistry = (data) => {
  return service({
    url: '/imageRegistry/updateImageRegistry',
    method: 'put',
    data
  })
}

// @Tags ImageRegistry
// @Summary 用id查询镜像库
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ImageRegistry true "用id查询镜像库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /imageRegistry/findImageRegistry [get]
export const findImageRegistry = (params) => {
  return service({
    url: '/imageRegistry/findImageRegistry',
    method: 'get',
    params
  })
}

// @Tags ImageRegistry
// @Summary 分页获取镜像库列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取镜像库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /imageRegistry/getImageRegistryList [get]
export const getImageRegistryList = (params) => {
  return service({
    url: '/imageRegistry/getImageRegistryList',
    method: 'get',
    params
  })
}

// @Tags ImageRegistry
// @Summary 不需要鉴权的镜像库接口
// @Accept application/json
// @Produce application/json
// @Param data query imageregistryReq.ImageRegistrySearch true "分页获取镜像库列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /imageRegistry/getImageRegistryPublic [get]
export const getImageRegistryPublic = () => {
  return service({
    url: '/imageRegistry/getImageRegistryPublic',
    method: 'get',
  })
}
