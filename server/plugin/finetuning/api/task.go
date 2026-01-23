package api

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	finetuningModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model"
	finetuningRequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model/request"
	finetuningService "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FinetuningTaskApi struct{}

var finetuningTaskService = finetuningService.ServiceGroupApp.FinetuningTaskService

// CreateFinetuningTask 创建微调任务
// @Tags FinetuningTask
// @Summary 创建微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body finetuningRequest.CreateFinetuningTaskRequest true "创建微调任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /finetuning/createTask [post]
func (a *FinetuningTaskApi) CreateFinetuningTask(c *gin.Context) {
	var req finetuningRequest.CreateFinetuningTaskRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从上下文获取用户ID（uint -> int 转换）
	userIDUint := c.GetUint("user_id")
	userID := int(userIDUint)

	// 构建任务对象
	task := &finetuningModel.FinetuningTask{
		Name:        req.Name,
		BaseModel:   req.BaseModel,
		DatasetPath: req.DatasetPath,
		UserID:      &userID,
	}

	// 设置可选字段
	if req.Description != "" {
		task.Description = &req.Description
	}
	if req.OutputPath != "" {
		task.OutputPath = &req.OutputPath
	}
	if req.Command != "" {
		task.Command = &req.Command
	}

	// 序列化训练参数
	if len(req.TrainingArgs) > 0 {
		trainingArgsJSON, err := json.Marshal(req.TrainingArgs)
		if err != nil {
			global.GVA_LOG.Error("序列化训练参数失败", zap.Error(err))
			response.FailWithMessage("序列化训练参数失败", c)
			return
		}
		trainingArgsStr := string(trainingArgsJSON)
		task.TrainingArgs = &trainingArgsStr
	}

	// 序列化GPU配置
	if len(req.GPUConfig) > 0 {
		gpuConfigJSON, err := json.Marshal(req.GPUConfig)
		if err != nil {
			global.GVA_LOG.Error("序列化GPU配置失败", zap.Error(err))
			response.FailWithMessage("序列化GPU配置失败", c)
			return
		}
		gpuConfigStr := string(gpuConfigJSON)
		task.GPUConfig = &gpuConfigStr
	}

	// 创建任务
	err = finetuningTaskService.CreateFinetuningTask(context.Background(), task)
	if err != nil {
		global.GVA_LOG.Error("创建任务失败!", zap.Error(err))
		response.FailWithMessage("创建任务失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("任务创建成功，已在后台开始执行", c)
}

// DeleteFinetuningTask 删除微调任务
// @Tags FinetuningTask
// @Summary 删除微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query finetuningRequest.DeleteFinetuningTaskRequest true "删除微调任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /finetuning/deleteTask [delete]
func (a *FinetuningTaskApi) DeleteFinetuningTask(c *gin.Context) {
	var req finetuningRequest.DeleteFinetuningTaskRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = finetuningTaskService.DeleteFinetuningTask(req.ID)
	if err != nil {
		global.GVA_LOG.Error("删除任务失败!", zap.Error(err))
		response.FailWithMessage("删除任务失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// StopFinetuningTask 停止微调任务
// @Tags FinetuningTask
// @Summary 停止微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query finetuningRequest.StopFinetuningTaskRequest true "停止微调任务"
// @Success 200 {object} response.Response{msg=string} "停止成功"
// @Router /finetuning/stopTask [post]
func (a *FinetuningTaskApi) StopFinetuningTask(c *gin.Context) {
	var req finetuningRequest.StopFinetuningTaskRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = finetuningTaskService.StopFinetuningTask(req.ID)
	if err != nil {
		global.GVA_LOG.Error("停止任务失败!", zap.Error(err))
		response.FailWithMessage("停止任务失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("任务已停止", c)
}

// GetFinetuningTask 用id查询微调任务
// @Tags FinetuningTask
// @Summary 用id查询微调任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query finetuningRequest.GetFinetuningTaskById true "用id查询微调任务"
// @Success 200 {object} response.Response{data=finetuningModel.FinetuningTask,msg=string} "查询成功"
// @Router /finetuning/getTask [get]
func (a *FinetuningTaskApi) GetFinetuningTask(c *gin.Context) {
	var req finetuningRequest.GetFinetuningTaskById
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	task, err := finetuningTaskService.GetFinetuningTaskById(req.ID)
	if err != nil {
		global.GVA_LOG.Error("查询任务失败!", zap.Error(err))
		response.FailWithMessage("查询任务失败", c)
		return
	}

	response.OkWithData(task, c)
}

// GetFinetuningTaskList 分页获取微调任务列表
// @Tags FinetuningTask
// @Summary 分页获取微调任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query finetuningRequest.FinetuningTaskSearch true "分页获取微调任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /finetuning/getTaskList [get]
func (a *FinetuningTaskApi) GetFinetuningTaskList(c *gin.Context) {
	var pageInfo finetuningRequest.FinetuningTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := finetuningTaskService.GetFinetuningTaskList(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取任务列表失败!", zap.Error(err))
		response.FailWithMessage("获取任务列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetFinetuningTaskLog 获取微调任务日志
// @Tags FinetuningTask
// @Summary 获取微调任务日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "任务ID"
// @Param lines query int false "获取日志行数"
// @Param offset query int false "日志偏移量"
// @Success 200 {object} response.Response{data=string,msg=string} "获取成功"
// @Router /finetuning/getTaskLog [get]
func (a *FinetuningTaskApi) GetFinetuningTaskLog(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("任务ID格式错误", c)
		return
	}

	var lines, offset *int
	if linesStr := c.Query("lines"); linesStr != "" {
		l, err := strconv.Atoi(linesStr)
		if err == nil {
			lines = &l
		}
	}
	if offsetStr := c.Query("offset"); offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err == nil {
			offset = &o
		}
	}

	logContent, err := finetuningTaskService.GetTaskLog(uint(id), lines, offset)
	if err != nil {
		global.GVA_LOG.Error("获取任务日志失败!", zap.Error(err))
		response.FailWithMessage("获取任务日志失败: "+err.Error(), c)
		return
	}

	response.OkWithData(logContent, c)
}
