package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/service"

var (
	Api = new(api)

	serviceConversation = service.Service.Conversation
	serviceMessage      = service.Service.Message
	serviceChat         = service.Service.Chat
	serviceConfig       = service.Service.Config
)

type api struct {
	Conversation conversation
	Message      message
	Chat         chat
	Config       config
}
