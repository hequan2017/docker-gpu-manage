package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/llm/createLlmModel", Description: "新建开源大模型", ApiGroup: "llm_model", Method: "POST"},
		{Path: "/llm/deleteLlmModel", Description: "删除开源大模型", ApiGroup: "llm_model", Method: "DELETE"},
		{Path: "/llm/deleteLlmModelByIds", Description: "批量删除开源大模型", ApiGroup: "llm_model", Method: "DELETE"},
		{Path: "/llm/updateLlmModel", Description: "更新开源大模型", ApiGroup: "llm_model", Method: "PUT"},
		{Path: "/llm/findLlmModel", Description: "根据ID获取开源大模型", ApiGroup: "llm_model", Method: "GET"},
		{Path: "/llm/getLlmModelList", Description: "获取开源大模型列表", ApiGroup: "llm_model", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
