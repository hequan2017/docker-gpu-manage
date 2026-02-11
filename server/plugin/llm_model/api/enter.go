package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/service"

var (
	Api             = new(api)
	serviceLlmModel = service.Service.LlmModel
)

type api struct{ LlmModel llm }
