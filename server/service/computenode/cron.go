package computenode

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	instanceSvc "github.com/flipped-aurora/gin-vue-admin/server/service/instance"
	"go.uber.org/zap"
)

// CheckAllNodeDockerStatus 检查所有节点的 Docker 状态并更新数据库
func CheckAllNodeDockerStatus(ctx context.Context) {
	var nodes []model.ComputeNode
	if err := global.GVA_DB.Where("deleted_at IS NULL").Find(&nodes).Error; err != nil {
		global.GVA_LOG.Error("查询算力节点失败", zap.Error(err))
		return
	}

	dockerSvc := instanceSvc.DockerService{}
	var success, failed int

	for i := range nodes {
		n := &nodes[i]

		// 为单个节点设置超时，避免长时间阻塞
		checkCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		connected, message := dockerSvc.TestDockerConnection(checkCtx, n)
		cancel()

		if connected {
			status := "connected"
			_ = global.GVA_DB.Model(n).Where("id = ?", n.ID).Update("docker_status", status).Error
			// global.GVA_LOG.Debug("Docker连接正常", zap.Uint("nodeId", n.ID), zap.String("name", safeStr(n.Name)))
			success++
		} else {
			status := "failed"
			_ = global.GVA_DB.Model(n).Where("id = ?", n.ID).Update("docker_status", status).Error
			global.GVA_LOG.Warn("Docker连接异常", zap.Uint("nodeId", n.ID), zap.String("name", safeStr(n.Name)), zap.String("error", message))
			failed++
		}

		// 防止对远端Docker造成瞬时高压
		time.Sleep(100 * time.Millisecond)
	}

	// 已移除定时任务汇总日志，避免控制台噪声
}

func safeStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
