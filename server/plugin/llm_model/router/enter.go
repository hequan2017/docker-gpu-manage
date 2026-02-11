package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/api"

var (
	Router      = new(router)
	apiLlmModel = api.Api.LlmModel
)

type router struct{ LlmModel llm }
