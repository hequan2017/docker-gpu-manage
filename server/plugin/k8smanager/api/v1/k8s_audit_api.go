package v1

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type K8sAuditApi struct{}

// GetAuditLogs 获取审计日志列表
// @Tags K8sAudit
// @Summary 获取审计日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sAuditLogSearch true "查询条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8s/audit/list [get]
func (a *K8sAuditApi) GetAuditLogs(c *gin.Context) {
	var info model.K8sAuditLogSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	logs, total, err := service.ServiceGroupApp.K8sAuditService.GetAuditLogs(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("获取审计日志失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     logs,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

// GetAuditLogStats 获取审计日志统计
// @Tags K8sAudit
// @Summary 获取审计日志统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param days query int false "统计天数" default(7)
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /k8s/audit/stats [get]
func (a *K8sAuditApi) GetAuditLogStats(c *gin.Context) {
	days := 7
	if d := c.Query("days"); d != "" {
		if num, err := parseInt(d); err == nil && num > 0 {
			days = num
		}
	}

	stats, err := service.ServiceGroupApp.K8sAuditService.GetAuditLogStats(c.Request.Context(), days)
	if err != nil {
		global.GVA_LOG.Error("获取审计统计失败!", zap.String("error", err.Error()))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetClientStats 获取客户端连接统计
// @Tags K8sAudit
// @Summary 获取客户端连接统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /k8s/audit/client-stats [get]
func (a *K8sAuditApi) GetClientStats(c *gin.Context) {
	stats := service.GetClientStats()
	response.OkWithDetailed(stats, "获取成功", c)
}

// DeleteOldLogs 删除旧的审计日志
// @Tags K8sAudit
// @Summary 删除旧的审计日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param retainDays query int false "保留天数" default(90)
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8s/audit/cleanup [delete]
func (a *K8sAuditApi) DeleteOldLogs(c *gin.Context) {
	retainDays := 90
	if rd := c.Query("retainDays"); rd != "" {
		if num, err := parseInt(rd); err == nil && num > 0 {
			retainDays = num
		}
	}

	err := service.ServiceGroupApp.K8sAuditService.DeleteOldLogs(c.Request.Context(), retainDays)
	if err != nil {
		global.GVA_LOG.Error("清理审计日志失败!", zap.String("error", err.Error()))
		response.FailWithMessage("清理失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("清理成功", c)
}

// ExportAuditLogs 导出审计日志
// @Tags K8sAudit
// @Summary 导出审计日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sAuditLogSearch true "查询条件"
// @Success 200 {object} response.Response{data=[]model.K8sAuditLog,msg=string} "导出成功"
// @Router /k8s/audit/export [get]
func (a *K8sAuditApi) ExportAuditLogs(c *gin.Context) {
	var info model.K8sAuditLogSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	logs, err := service.ServiceGroupApp.K8sAuditService.ExportAuditLogs(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("导出审计日志失败!", zap.String("error", err.Error()))
		response.FailWithMessage("导出失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(logs, "导出成功", c)
}

func parseInt(s string) (int, error) {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	return num, err
}
