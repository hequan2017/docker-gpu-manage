package instance

import (
	"context"
	"math"
	"sync"
	"sync/atomic"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"github.com/gogf/gf/v2/os/gcron"
	"go.uber.org/zap"
)

// StartContainerStatusCheckCron 启动容器状态检查定时任务
// 每30秒检查一次所有容器的运行状态，并刷新CPU/内存/GPU显存使用率
func StartContainerStatusCheckCron() {
	// 使用 gcron 创建定时任务，每30秒执行一次
	// cron 表达式格式：秒 分钟 小时 日 月 星期
	_, err := gcron.AddSingleton(context.Background(), "*/30 * * * * *", func(ctx context.Context) {
		checkAllContainerStatusAndMetrics(ctx)
	}, "container-status-metrics-check")
	if err != nil {
		global.GVA_LOG.Error("启动容器状态/指标检查定时任务失败", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("容器状态/指标检查定时任务已启动，每30秒检查一次")
}

// checkAllContainerStatusAndMetrics 检查所有容器的状态并刷新监控指标（并发版本）
func checkAllContainerStatusAndMetrics(ctx context.Context) {
	// 获取所有有容器ID的实例（未删除的）
	var instances []instanceModel.Instance
	if err := global.GVA_DB.Where("deleted_at IS NULL AND container_id IS NOT NULL AND container_id != ''").Find(&instances).Error; err != nil {
		global.GVA_LOG.Error("查询实例列表失败", zap.Error(err))
		return
	}

	total := len(instances)
	if total == 0 {
		global.GVA_LOG.Info("没有需要检查的容器实例")
		return
	}

	// 使用原子操作计数
	var successCount, failCount int64

	// 创建 WaitGroup 用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 使用信号量控制并发数，避免过多 goroutine 对系统造成压力
	// 建议并发数为 10-20，可根据实际情况调整
	maxConcurrency := 30
	semaphore := make(chan struct{}, maxConcurrency)

	for _, inst := range instances {
		if inst.ContainerId == nil || *inst.ContainerId == "" || inst.NodeId == nil {
			continue
		}

		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量

		go func(instance instanceModel.Instance) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			// 1) 同步容器状态
			if err := dockerService.SyncContainerStatus(ctx, instance.ID); err != nil {
				global.GVA_LOG.Error("同步容器状态失败",
					zap.Uint("实例ID", instance.ID),
					zap.String("容器ID", *instance.ContainerId),
					zap.Error(err))
				atomic.AddInt64(&failCount, 1)
				return
			}
			atomic.AddInt64(&successCount, 1)

			// 2) 刷新CPU/内存/GPU显存使用率
			var node computenode.ComputeNode
			if err := global.GVA_DB.Where("id = ?", *instance.NodeId).First(&node).Error; err != nil {
				global.GVA_LOG.Warn("获取节点信息失败，跳过指标刷新",
					zap.Uint("实例ID", instance.ID), zap.Error(err))
				return
			}

			stats, err := dockerService.GetContainerStats(ctx, &node, *instance.ContainerId)
			if err != nil {
				global.GVA_LOG.Warn("获取容器统计信息失败，跳过指标刷新",
					zap.Uint("实例ID", instance.ID), zap.String("容器ID", *instance.ContainerId), zap.Error(err))
				return
			}

			// 将 CPU 和内存使用率保留 2 位小数
			cpu := roundToTwoDecimals(stats.CPUUsagePercent)
			mem := roundToTwoDecimals(stats.MemoryUsagePercent)

			// GPU 显存采集校验：根据规格是否需要GPU来决定是否记录GPU指标
			gpu := stats.GPUMemoryUsageRate
			if instance.SpecId != nil {
				var spec product.ProductSpec
				if err := global.GVA_DB.Where("id = ?", *instance.SpecId).First(&spec).Error; err == nil {
					expectGPU := spec.GpuCount != nil && *spec.GpuCount > 0
					if !expectGPU {
						// 无GPU需求则强制置零
						gpu = 0
					} else {
						// 有GPU需求但未检测到显存信息时给出告警
						if stats.GPUMemorySizeGB <= 0 {
							global.GVA_LOG.Warn("GPU显存采集异常：未检测到显存信息",
								zap.Uint("实例ID", instance.ID),
								zap.String("容器ID", *instance.ContainerId),
								zap.Any("规格GPU数量", *spec.GpuCount))
						}
					}
				}
			}
			// GPU 使用率保留 2 位小数
			gpu = roundToTwoDecimals(gpu)

			if err := global.GVA_DB.Model(&instanceModel.Instance{}).
				Where("id = ?", instance.ID).
				Updates(map[string]any{
					"cpu_usage_percent":     cpu,
					"memory_usage_percent":  mem,
					"gpu_memory_usage_rate": gpu,
				}).Error; err != nil {
				global.GVA_LOG.Warn("写入实例指标失败",
					zap.Uint("实例ID", instance.ID), zap.Error(err))
			}
		}(inst)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	global.GVA_LOG.Info("容器状态/指标检查完成（并发模式）",
		zap.Int("总数", total),
		zap.Int64("成功", atomic.LoadInt64(&successCount)),
		zap.Int64("失败", atomic.LoadInt64(&failCount)))
}

// roundToTwoDecimals 将浮点数保留 2 位小数
func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
