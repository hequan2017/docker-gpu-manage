package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/api"

var (
	Router = new(router)

	apiConversation = api.Api.Conversation
	apiMessage      = api.Api.Message
	apiChat         = api.Api.Chat
	apiConfig       = api.Api.Config
)

type router struct {
	Conversation conversation
	Message      message
	Chat         chat
	Config       config
}
