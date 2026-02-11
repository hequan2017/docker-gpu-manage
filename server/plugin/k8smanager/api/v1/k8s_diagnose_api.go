package v1

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	aiagentModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model"
	aiagentReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/model/request"
	aiagentService "github.com/flipped-aurora/gin-vue-admin/server/plugin/aiagent/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DiagnosePod 使用AI诊断Pod
// @Tags K8sPod
// @Summary 使用AI诊断Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param podName query string true "Pod名称"
// @Success 200 {object} response.Response{data=string,msg=string} "诊断结果"
// @Router /k8s/pod/diagnose [post]
func (a *K8sPodApi) DiagnosePod(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	podName := c.Query("podName")

	if clusterName == "" || namespace == "" || podName == "" {
		response.FailWithMessage("参数不完整", c)
		return
	}

	// 1. 获取Pod详情和事件
	pod, err := service.ServiceGroupApp.K8sPodService.GetPod(c.Request.Context(), clusterName, namespace, podName)
	if err != nil {
		global.GVA_LOG.Error("获取Pod详情失败", zap.Error(err))
		response.FailWithMessage("获取Pod详情失败: "+err.Error(), c)
		return
	}

	events, err := service.ServiceGroupApp.K8sPodService.GetPodEvents(c.Request.Context(), clusterName, namespace, podName)
	if err != nil {
		global.GVA_LOG.Error("获取Pod事件失败", zap.Error(err))
		response.FailWithMessage("获取Pod事件失败: "+err.Error(), c)
		return
	}

	// 2. 获取最近的日志（如果有错误）
	logs := ""
	if pod.Status.Phase != "Running" && pod.Status.Phase != "Succeeded" {
		// 尝试获取日志
		// ... 这里简化处理，实际可以复用 GetPodLog
	}

	// 3. 构建 Prompt
	prompt := fmt.Sprintf(`请作为 Kubernetes 专家分析以下 Pod 的状态并给出诊断建议：
Pod 名称: %s
命名空间: %s
状态: %s
重启次数: %d

Events:
`, pod.Name, pod.Namespace, pod.Status.Phase, len(pod.Status.ContainerStatuses))

	for _, e := range events.Items {
		prompt += fmt.Sprintf("- [%s] %s: %s\n", e.Type, e.Reason, e.Message)
	}

	if logs != "" {
		prompt += "\n最近日志:\n" + logs
	}

	// 4. 调用 AI Agent
	// 获取激活的配置
	config, err := aiagentService.Config.GetActiveConfig()
	if err != nil {
		response.FailWithMessage("未找到激活的AI配置", c)
		return
	}

	// 创建临时会话
	userID := int(utils.GetUserID(c))
	conversation := aiagentModel.Conversation{
		Title:     fmt.Sprintf("诊断 Pod: %s", podName),
		Model:     config.Model,
		IsActive:  true,
		MaxTokens: &config.MaxTokens,
		UserID:    &userID,
	}
	err = aiagentService.Conversation.CreateConversation(&conversation)
	if err != nil {
		response.FailWithMessage("创建AI会话失败", c)
		return
	}

	// 发送消息
	chatReq := aiagentReq.ChatRequest{
		ConversationID: &conversation.ID,
		Message:        prompt,
		Stream:         new(bool), // default false
	}

	chatResp, err := aiagentService.Chat.SendMessage(chatReq, int(utils.GetUserID(c)))
	if err != nil {
		global.GVA_LOG.Error("调用AI服务失败", zap.Error(err))
		response.FailWithMessage("调用AI服务失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(chatResp.Content, "AI诊断完成", c)
}
