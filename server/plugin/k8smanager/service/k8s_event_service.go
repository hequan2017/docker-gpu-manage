package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sEventService Event服务结构体
type K8sEventService struct{}

// GetEventList 获取Event列表
func (s *K8sEventService) GetEventList(ctx context.Context, req *request.GetEventsRequest) (*corev1.EventList, error) {
	client, err := GetClusterClient(req.ClusterName)
	if err != nil {
		return nil, err
	}

	listOptions := metav1.ListOptions{}
	if req.Kind != "" && req.Name != "" {
		// 构建字段选择器
		listOptions.FieldSelector = fmt.Sprintf("involvedObject.kind=%s,involvedObject.name=%s", req.Kind, req.Name)
	}
	if req.Limit > 0 {
		listOptions.Limit = req.Limit
	}

	var eventList *corev1.EventList
	if req.Namespace == "" || req.Namespace == "all" {
		eventList, err = client.Clientset.CoreV1().Events("").List(ctx, listOptions)
	} else {
		eventList, err = client.Clientset.CoreV1().Events(req.Namespace).List(ctx, listOptions)
	}

	if err != nil {
		return nil, fmt.Errorf("获取Event列表失败: %w", err)
	}

	return eventList, nil
}
