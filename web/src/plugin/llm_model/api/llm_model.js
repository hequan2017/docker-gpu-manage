import service from '@/utils/request'
// @Tags LlmModel
// @Summary 创建开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "创建开源大模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /llm/createLlmModel [post]
export const createLlmModel = (data) => {
  return service({
    url: '/llm/createLlmModel',
    method: 'post',
    data
  })
}

// @Tags LlmModel
// @Summary 删除开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "删除开源大模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /llm/deleteLlmModel [delete]
export const deleteLlmModel = (params) => {
  return service({
    url: '/llm/deleteLlmModel',
    method: 'delete',
    params
  })
}

// @Tags LlmModel
// @Summary 批量删除开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除开源大模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /llm/deleteLlmModel [delete]
export const deleteLlmModelByIds = (params) => {
  return service({
    url: '/llm/deleteLlmModelByIds',
    method: 'delete',
    params
  })
}

// @Tags LlmModel
// @Summary 更新开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LlmModel true "更新开源大模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /llm/updateLlmModel [put]
export const updateLlmModel = (data) => {
  return service({
    url: '/llm/updateLlmModel',
    method: 'put',
    data
  })
}

// @Tags LlmModel
// @Summary 用id查询开源大模型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.LlmModel true "用id查询开源大模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /llm/findLlmModel [get]
export const findLlmModel = (params) => {
  return service({
    url: '/llm/findLlmModel',
    method: 'get',
    params
  })
}

// @Tags LlmModel
// @Summary 分页获取开源大模型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取开源大模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /llm/getLlmModelList [get]
export const getLlmModelList = (params) => {
  return service({
    url: '/llm/getLlmModelList',
    method: 'get',
    params
  })
}
// @Tags LlmModel
// @Summary 不需要鉴权的开源大模型接口
// @Accept application/json
// @Produce application/json
// @Param data query request.LlmModelSearch true "分页获取开源大模型列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /llm/getLlmModelPublic [get]
export const getLlmModelPublic = () => {
  return service({
    url: '/llm/getLlmModelPublic',
    method: 'get',
  })
}
