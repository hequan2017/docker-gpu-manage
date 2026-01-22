package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// 会话相关API
		{Path: "/conversation/createConversation", Description: "创建会话", ApiGroup: "AI Agent", Method: "POST"},
		{Path: "/conversation/deleteConversation", Description: "删除会话", ApiGroup: "AI Agent", Method: "DELETE"},
		{Path: "/conversation/updateConversation", Description: "更新会话", ApiGroup: "AI Agent", Method: "PUT"},
		{Path: "/conversation/findConversation", Description: "根据ID获取会话", ApiGroup: "AI Agent", Method: "GET"},
		{Path: "/conversation/getConversationList", Description: "获取会话列表", ApiGroup: "AI Agent", Method: "GET"},
		{Path: "/conversation/setActive", Description: "设置会话激活状态", ApiGroup: "AI Agent", Method: "POST"},
		{Path: "/conversation/getActive", Description: "获取激活的会话", ApiGroup: "AI Agent", Method: "GET"},
		// 消息相关API
		{Path: "/message/getMessageList", Description: "获取消息列表", ApiGroup: "AI Agent", Method: "GET"},
		{Path: "/message/deleteMessage", Description: "删除消息", ApiGroup: "AI Agent", Method: "DELETE"},
		// 聊天相关API
		{Path: "/chat/sendMessage", Description: "发送消息", ApiGroup: "AI Agent", Method: "POST"},
		// 配置相关API
		{Path: "/config/createConfig", Description: "创建AI配置", ApiGroup: "AI Agent", Method: "POST"},
		{Path: "/config/deleteConfig", Description: "删除AI配置", ApiGroup: "AI Agent", Method: "DELETE"},
		{Path: "/config/updateConfig", Description: "更新AI配置", ApiGroup: "AI Agent", Method: "PUT"},
		{Path: "/config/findConfig", Description: "根据ID获取AI配置", ApiGroup: "AI Agent", Method: "GET"},
		{Path: "/config/getConfigList", Description: "获取AI配置列表", ApiGroup: "AI Agent", Method: "GET"},
		{Path: "/config/setActive", Description: "设置AI配置激活状态", ApiGroup: "AI Agent", Method: "POST"},
		{Path: "/config/getActive", Description: "获取激活的AI配置", ApiGroup: "AI Agent", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
