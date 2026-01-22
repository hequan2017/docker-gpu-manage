package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/finetuning/createTask",
			Description: "创建微调任务",
			ApiGroup:    "算法微调",
			Method:      "POST",
		},
		{
			Path:        "/finetuning/deleteTask",
			Description: "删除微调任务",
			ApiGroup:    "算法微调",
			Method:      "DELETE",
		},
		{
			Path:        "/finetuning/stopTask",
			Description: "停止微调任务",
			ApiGroup:    "算法微调",
			Method:      "POST",
		},
		{
			Path:        "/finetuning/getTask",
			Description: "根据ID获取微调任务",
			ApiGroup:    "算法微调",
			Method:      "GET",
		},
		{
			Path:        "/finetuning/getTaskList",
			Description: "获取微调任务列表",
			ApiGroup:    "算法微调",
			Method:      "GET",
		},
		{
			Path:        "/finetuning/getTaskLog",
			Description: "获取微调任务日志",
			ApiGroup:    "算法微调",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
