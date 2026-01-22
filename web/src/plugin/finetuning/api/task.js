import service from '@/utils/request'

// @Tags FinetuningTask
// @Summary 创建微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateFinetuningTaskRequest true "创建微调任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /finetuning/createTask [post]
export const createFinetuningTask = (data) => {
  return service({
    url: '/finetuning/createTask',
    method: 'post',
    data
  })
}

// @Tags FinetuningTask
// @Summary 删除微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DeleteFinetuningTaskRequest true "删除微调任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /finetuning/deleteTask [delete]
export const deleteFinetuningTask = (params) => {
  return service({
    url: '/finetuning/deleteTask',
    method: 'delete',
    params
  })
}

// @Tags FinetuningTask
// @Summary 停止微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.StopFinetuningTaskRequest true "停止微调任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"停止成功"}"
// @Router /finetuning/stopTask [post]
export const stopFinetuningTask = (params) => {
  return service({
    url: '/finetuning/stopTask',
    method: 'post',
    params
  })
}

// @Tags FinetuningTask
// @Summary 根据ID获取微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetFinetuningTaskById true "根据ID获取微调任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /finetuning/getTask [get]
export const getFinetuningTask = (params) => {
  return service({
    url: '/finetuning/getTask',
    method: 'get',
    params
  })
}

// @Tags FinetuningTask
// @Summary 分页获取微调任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.FinetuningTaskSearch true "分页获取微调任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /finetuning/getTaskList [get]
export const getFinetuningTaskList = (params) => {
  return service({
    url: '/finetuning/getTaskList',
    method: 'get',
    params
  })
}

// @Tags FinetuningTask
// @Summary 获取微调任务日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "任务ID"
// @Param lines query int false "获取日志行数"
// @Param offset query int false "日志偏移量"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /finetuning/getTaskLog [get]
export const getFinetuningTaskLog = (params) => {
  return service({
    url: '/finetuning/getTaskLog',
    method: 'get',
    params
  })
}
