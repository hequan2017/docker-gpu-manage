package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"go.uber.org/zap"
)

// K8sAuditService K8s审计服务
type K8sAuditService struct{}

// AuditRecord 审计记录
type AuditRecord struct {
	UserID      uint
	Username    string
	UserIP      string
	UserAgent   string
	Action      string
	Resource    string
	ResourceID  string
	ClusterName string
	Namespace   string
	Description string
	RequestData interface{}
	Response    interface{}
	Status      string
	ErrorMsg    string
	Metadata    map[string]interface{}
}

// LogOperation 记录操作日志
func (s *K8sAuditService) LogOperation(ctx context.Context, record *AuditRecord) error {
	startTime := time.Now()

	// 序列化请求数据
	var requestData string
	if record.RequestData != nil {
		if bytes, err := json.Marshal(record.RequestData); err == nil {
			requestData = string(bytes)
		}
	}

	// 序列化响应数据
	var responseData string
	if record.Response != nil {
		if bytes, err := json.Marshal(record.Response); err == nil {
			// 限制响应数据大小
			responseStr := string(bytes)
			if len(responseStr) > 10000 {
				responseStr = responseStr[:10000] + "...(truncated)"
			}
			responseData = responseStr
		}
	}

	// 序列化元数据
	var metadata string
	if record.Metadata != nil {
		if bytes, err := json.Marshal(record.Metadata); err == nil {
			metadata = string(bytes)
		}
	}

	// 计算执行时长（从 context 中获取开始时间）
	duration := time.Since(startTime).Milliseconds()

	auditLog := &model.K8sAuditLog{
		UserID:      record.UserID,
		Username:    record.Username,
		UserIP:      record.UserIP,
		UserAgent:   record.UserAgent,
		Action:      record.Action,
		Resource:    record.Resource,
		ResourceID:  record.ResourceID,
		ClusterName: record.ClusterName,
		Namespace:   record.Namespace,
		Description: record.Description,
		RequestData: requestData,
		Response:    responseData,
		Status:      record.Status,
		ErrorMsg:    record.ErrorMsg,
		Duration:    duration,
		Metadata:    metadata,
	}

	// 异步写入日志，避免影响主流程
	go func() {
		if err := global.GVA_DB.Create(auditLog).Error; err != nil {
			global.GVA_LOG.Error("写入审计日志失败", zap.Error(err))
		}
	}()

	return nil
}

// GetAuditLogs 获取审计日志列表
func (s *K8sAuditService) GetAuditLogs(ctx context.Context, info *model.K8sAuditLogSearch) ([]model.K8sAuditLog, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 构建查询条件
	db := global.GVA_DB.Model(&model.K8sAuditLog{})

	// 时间范围过滤
	if !info.StartTime.IsZero() {
		db = db.Where("created_at >= ?", info.StartTime)
	}
	if !info.EndTime.IsZero() {
		db = db.Where("created_at <= ?", info.EndTime)
	}

	// 用户过滤
	if info.UserID > 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}

	// 操作类型过滤
	if info.Action != "" {
		db = db.Where("action = ?", info.Action)
	}

	// 资源类型过滤
	if info.Resource != "" {
		db = db.Where("resource = ?", info.Resource)
	}

	// 集群过滤
	if info.Cluster != "" {
		db = db.Where("cluster_name = ?", info.Cluster)
	}

	// 状态过滤
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}

	// 关键词搜索
	if info.Keyword != "" {
		db = db.Where("description LIKE ? OR resource_id LIKE ? OR error_msg LIKE ?",
			"%"+info.Keyword+"%", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}

	// 获取总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取列表
	var logs []model.K8sAuditLog
	err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// GetAuditLogStats 获取审计日志统计信息
func (s *K8sAuditService) GetAuditLogStats(ctx context.Context, days int) (map[string]interface{}, error) {
	if days <= 0 {
		days = 7 // 默认7天
	}

	startTime := time.Now().AddDate(0, 0, -days)

	// 总操作数
	var totalOps int64
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Where("created_at >= ?", startTime).
		Count(&totalOps).Error; err != nil {
		return nil, err
	}

	// 按操作类型统计
	var actionStats []struct {
		Action string
		Count  int64
	}
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Select("action, count(*) as count").
		Where("created_at >= ?", startTime).
		Group("action").
		Order("count DESC").
		Find(&actionStats).Error; err != nil {
		return nil, err
	}

	// 按状态统计
	var statusStats []struct {
		Status string
		Count  int64
	}
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Select("status, count(*) as count").
		Where("created_at >= ?", startTime).
		Group("status").
		Find(&statusStats).Error; err != nil {
		return nil, err
	}

	// 按用户统计（Top 10）
	var userStats []struct {
		Username string
		Count    int64
	}
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Select("username, count(*) as count").
		Where("created_at >= ?", startTime).
		Group("username").
		Order("count DESC").
		Limit(10).
		Find(&userStats).Error; err != nil {
		return nil, err
	}

	// 失败操作数
	var failedOps int64
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Where("created_at >= ? AND status = ?", startTime, "failure").
		Count(&failedOps).Error; err != nil {
		return nil, err
	}

	// 平均执行时长
	var avgDuration float64
	if err := global.GVA_DB.Model(&model.K8sAuditLog{}).
		Select("AVG(duration) as avg_duration").
		Where("created_at >= ?", startTime).
		Scan(&avgDuration).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_operations":    totalOps,
		"failed_operations":   failedOps,
		"success_rate":        fmt.Sprintf("%.2f%%", float64(totalOps-failedOps)/float64(totalOps)*100),
		"avg_duration_ms":     fmt.Sprintf("%.2f", avgDuration),
		"action_stats":        actionStats,
		"status_stats":        statusStats,
		"top_users":           userStats,
		"period_days":         days,
	}, nil
}

// DeleteOldLogs 删除旧的审计日志
func (s *K8sAuditService) DeleteOldLogs(ctx context.Context, retainDays int) error {
	if retainDays <= 0 {
		retainDays = 90 // 默认保留90天
	}

	cutoffDate := time.Now().AddDate(0, 0, -retainDays)

	result := global.GVA_DB.Where("created_at < ?", cutoffDate).Delete(&model.K8sAuditLog{})
	if result.Error != nil {
		return result.Error
	}

	global.GVA_LOG.Info("清理旧审计日志",
		zap.Int64("deleted", result.RowsAffected),
		zap.Int("retain_days", retainDays))

	return nil
}

// ExportAuditLogs 导出审计日志（用于审计报告）
func (s *K8sAuditService) ExportAuditLogs(ctx context.Context, info *model.K8sAuditLogSearch) ([]model.K8sAuditLog, error) {
	// 构建查询条件（与 GetAuditLogs 相同）
	db := global.GVA_DB.Model(&model.K8sAuditLog{})

	if !info.StartTime.IsZero() {
		db = db.Where("created_at >= ?", info.StartTime)
	}
	if !info.EndTime.IsZero() {
		db = db.Where("created_at <= ?", info.EndTime)
	}
	if info.UserID > 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Action != "" {
		db = db.Where("action = ?", info.Action)
	}
	if info.Resource != "" {
		db = db.Where("resource = ?", info.Resource)
	}
	if info.Cluster != "" {
		db = db.Where("cluster_name = ?", info.Cluster)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}

	// 限制导出数量
	var logs []model.K8sAuditLog
	err := db.Order("created_at DESC").Limit(10000).Find(&logs).Error
	return logs, err
}
