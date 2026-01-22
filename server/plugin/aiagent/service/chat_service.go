package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	"go.uber.org/zap"
)

var Chat = new(chat)

type chat struct{}

// GLMMessage GLM API消息格式
type GLMMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GLMChatRequest GLM API请求格式
type GLMChatRequest struct {
	Model       string      `json:"model"`
	Messages    []GLMMessage `json:"messages"`
	Temperature *float64    `json:"temperature,omitempty"`
	MaxTokens   *int        `json:"max_tokens,omitempty"`
	Stream      bool        `json:"stream,omitempty"`
}

// GLMChatResponse GLM API响应格式
type GLMChatResponse struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// GLMErrorResponse GLM API错误响应
type GLMErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

// SendMessage 发送消息并获取AI回复
func (s *chat) SendMessage(chatReq request.ChatRequest, userID int) (resp request.ChatResponse, err error) {
	// 1. 获取或创建会话
	var conversation model.Conversation
	if chatReq.ConversationID != nil && *chatReq.ConversationID > 0 {
		// 使用现有会话
		conversation, err = Conversation.GetConversation(fmt.Sprint(*chatReq.ConversationID))
		if err != nil {
			global.GVA_LOG.Error("获取会话失败", zap.Error(err))
			return resp, err
		}
	} else {
		// 创建新会话
		// 获取激活的配置
		config, err := Config.GetActiveConfig()
		if err != nil {
			global.GVA_LOG.Error("获取AI配置失败", zap.Error(err))
			return resp, errors.New("获取AI配置失败，请联系管理员")
		}

		modelName := "glm-4-plus"
		if chatReq.Model != nil && *chatReq.Model != "" {
			modelName = *chatReq.Model
		} else {
			modelName = config.Model
		}

		temperature := config.Temperature
		if chatReq.Temperature != nil {
			temperature = *chatReq.Temperature
		}

		maxTokens := config.MaxTokens
		if chatReq.MaxTokens != nil {
			maxTokens = *chatReq.MaxTokens
		}

		conversation = model.Conversation{
			Title:       s.generateTitle(chatReq.Message),
			UserID:      &userID,
			Model:       modelName,
			Temperature: &temperature,
			MaxTokens:   &maxTokens,
			IsActive:    true,
		}
		err = Conversation.CreateConversation(&conversation)
		if err != nil {
			global.GVA_LOG.Error("创建会话失败", zap.Error(err))
			return resp, err
		}
	}

	// 2. 保存用户消息
	userMessage := model.Message{
		ConversationID: conversation.ID,
		Role:           "user",
		Content:        chatReq.Message,
	}
	err = Message.CreateMessage(&userMessage)
	if err != nil {
		global.GVA_LOG.Error("保存用户消息失败", zap.Error(err))
		return resp, err
	}

	// 3. 获取会话历史
	history, err := Message.GetMessagesByConversationID(conversation.ID)
	if err != nil {
		global.GVA_LOG.Error("获取会话历史失败", zap.Error(err))
		return resp, err
	}

	// 4. 构建GLM API请求
	config, err := Config.GetActiveConfig()
	if err != nil {
		global.GVA_LOG.Error("获取AI配置失败", zap.Error(err))
		return resp, errors.New("获取AI配置失败，请联系管理员")
	}

	glmMessages := []GLMMessage{}
	if conversation.SystemPrompt != nil && *conversation.SystemPrompt != "" {
		glmMessages = append(glmMessages, GLMMessage{
			Role:    "system",
			Content: *conversation.SystemPrompt,
		})
	}
	for _, msg := range history {
		glmMessages = append(glmMessages, GLMMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 获取温度参数，优先使用请求中的值
	var temperature *float64
	if chatReq.Temperature != nil {
		temperature = chatReq.Temperature
	} else {
		temperature = conversation.Temperature
	}

	// 获取最大token数，优先使用请求中的值
	var maxTokens *int
	if chatReq.MaxTokens != nil {
		maxTokens = chatReq.MaxTokens
	} else {
		maxTokens = conversation.MaxTokens
	}

	// 获取模型名称
	modelName := conversation.Model
	if chatReq.Model != nil && *chatReq.Model != "" {
		modelName = *chatReq.Model
	}

	glmReq := GLMChatRequest{
		Model:       modelName,
		Messages:    glmMessages,
		Temperature: temperature,
		MaxTokens:   maxTokens,
		Stream:      chatReq.Stream != nil && *chatReq.Stream,
	}

	// 5. 调用GLM API
	glmResp, err := s.callGLMAPI(glmReq, config)
	if err != nil {
		global.GVA_LOG.Error("调用GLM API失败", zap.Error(err))
		return resp, errors.New("调用AI服务失败: " + err.Error())
	}

	// 6. 保存AI回复
	if len(glmResp.Choices) > 0 {
		assistantMessage := model.Message{
			ConversationID: conversation.ID,
			Role:           "assistant",
			Content:        glmResp.Choices[0].Message.Content,
			TokenCount:     &glmResp.Usage.TotalTokens,
		}
		err = Message.CreateMessage(&assistantMessage)
		if err != nil {
			global.GVA_LOG.Error("保存AI消息失败", zap.Error(err))
			return resp, err
		}

		// 更新会话标题（如果是第一条消息）
		if len(history) == 0 {
			conversation.Title = s.generateTitle(chatReq.Message)
			Conversation.UpdateConversation(conversation)
		}

		resp = request.ChatResponse{
			ConversationID: conversation.ID,
			MessageID:      assistantMessage.ID,
			Content:        glmResp.Choices[0].Message.Content,
			FinishReason:   glmResp.Choices[0].FinishReason,
			Usage: request.Usage{
				PromptTokens:     glmResp.Usage.PromptTokens,
				CompletionTokens: glmResp.Usage.CompletionTokens,
				TotalTokens:      glmResp.Usage.TotalTokens,
			},
		}
	}

	return resp, nil
}

// callGLMAPI 调用GLM API
func (s *chat) callGLMAPI(req GLMChatRequest, config model.AgentConfig) (*GLMChatResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", config.BaseURL+"chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+config.APIKey)

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		var errResp GLMErrorResponse
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return nil, errors.New("API调用失败: " + string(body))
		}
		return nil, errors.New(errResp.Error.Message)
	}

	var glmResp GLMChatResponse
	err = json.Unmarshal(body, &glmResp)
	if err != nil {
		return nil, err
	}

	return &glmResp, nil
}

// generateTitle 生成会话标题
func (s *chat) generateTitle(message string) string {
	// 截取前30个字符作为标题
	message = strings.TrimSpace(message)
	if len(message) > 30 {
		return message[:30] + "..."
	}
	if len(message) == 0 {
		return "新对话"
	}
	return message
}
