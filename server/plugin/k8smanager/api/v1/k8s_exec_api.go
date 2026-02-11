package v1

import (
	"fmt"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ExecPod 在Pod中执行命令（WebSocket）
// @Tags K8sPod
// @Summary 在Pod中执行命令
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param clusterName query string true "集群名称"
// @Param namespace query string true "命名空间"
// @Param podName query string true "Pod名称"
// @Param container query string false "容器名称"
// @Param command query string false "执行命令"
// @Success 200 {string} string "WebSocket连接"
// @Router /k8s/pod/exec [get]
func (a *K8sPodApi) ExecPod(c *gin.Context) {
	clusterName := c.Query("clusterName")
	namespace := c.Query("namespace")
	podName := c.Query("podName")
	container := c.Query("container")
	command := c.Query("command")

	if clusterName == "" || namespace == "" || podName == "" {
		c.String(http.StatusBadRequest, "参数不完整")
		return
	}

	if command == "" {
		command = "/bin/sh"
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GVA_LOG.Error("WebSocket升级失败", zap.Error(err))
		return
	}
	defer ws.Close()

	err = service.ServiceGroupApp.K8sPodService.ExecPod(clusterName, namespace, podName, container, command, ws)
	if err != nil {
		global.GVA_LOG.Error("执行命令失败", zap.Error(err))
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
	}
}
