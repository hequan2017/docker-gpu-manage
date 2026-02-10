package initialize

import (
	"context"
	"fmt"

	computeNodeSvc "github.com/flipped-aurora/gin-vue-admin/server/service/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 其他定时任务定在这里 参考上方使用方法

		_, err = global.GVA_Timer.AddTaskByFuncWithSecond("PCDNDispatch", "*/10 * * * * *", func() {
			task.ProcessPcdnDispatchTasks()
		}, "PCDN调度任务执行", option...)
		if err != nil {
			fmt.Println("add PCDN dispatch timer error:", err)
		}

		_, err = global.GVA_Timer.AddTaskByFuncWithSecond("PCDNDispatch", "*/12 * * * * *", func() {
			task.SyncPcdnDispatchStatus()
		}, "PCDN调度状态同步", option...)
		if err != nil {
			fmt.Println("add PCDN sync timer error:", err)
		}

		//_, err := global.GVA_Timer.AddTaskByFunc("定时任务标识", "corn表达式", func() {
		//	具体执行内容...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()

	// 合并一个定时任务：每30秒执行一次，依次检查Docker状态和容器指标
	_, err := gcron.AddSingleton(context.Background(), "*/30 * * * * *", func(ctx context.Context) {
		// 先检查节点 Docker 状态，确保后续容器检查的依赖健康
		computeNodeSvc.CheckAllNodeDockerStatus(ctx)
		// 再检查容器状态与指标
		instance.CheckAllContainerStatusAndMetrics(ctx)
	}, "system-health-check")
	if err != nil {
		global.GVA_LOG.Error("启动合并定时任务失败", zap.Error(err))
	} else {
		global.GVA_LOG.Info("合并定时任务已启动，每30秒执行一次")
	}
}
