package service

import (
	"context"
	"io"

	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

// ExecPod 在Pod中执行命令
func (s *K8sPodService) ExecPod(clusterName, namespace, podName, container, command string, ws *websocket.Conn) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	cmd := []string{command}
	if command == "/bin/sh" {
		// 尝试 /bin/bash，如果失败则回退到 /bin/sh
		// 这里简单处理，直接使用用户指定的命令
	}

	req := client.Clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Container: container,
		Command:   cmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(client.Config, "POST", req.URL())
	if err != nil {
		return err
	}

	stream := &WsStream{ws: ws}
	return exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdin:  stream,
		Stdout: stream,
		Stderr: stream,
		Tty:    true,
	})
}

// WsStream WebSocket流适配器
type WsStream struct {
	ws *websocket.Conn
}

func (s *WsStream) Read(p []byte) (n int, err error) {
	_, message, err := s.ws.ReadMessage()
	if err != nil {
		return 0, err
	}
	copy(p, message)
	return len(message), nil
}

func (s *WsStream) Write(p []byte) (n int, err error) {
	err = s.ws.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
