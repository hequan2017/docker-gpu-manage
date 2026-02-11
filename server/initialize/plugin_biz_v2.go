package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/approval_flow"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dellasset"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/portforward"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *gin.Engine) {
	PluginInitV2(engine, announcement.Plugin, aiagent.Plugin, finetuning.Plugin, dellasset.Plugin, portforward.Plugin, k8smanager.Plugin, pcdn.Plugin)
	PluginInitV2(engine, server_lifecycle.Plugin)
	PluginInitV2(engine, llm_model.Plugin)
	PluginInitV2(engine, approval_flow.Plugin)
}
