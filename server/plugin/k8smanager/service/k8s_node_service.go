package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sNodeService Node服务结构体
type K8sNodeService struct{}

// GetNodeList 获取Node列表
func (s *K8sNodeService) GetNodeList(ctx context.Context, clusterName string) (*corev1.NodeList, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	nodeList, err := client.Clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Node列表失败: %w", err)
	}

	return nodeList, nil
}

// GetNode 获取单个Node详情
func (s *K8sNodeService) GetNode(ctx context.Context, clusterName, nodeName string) (*corev1.Node, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	node, err := client.Clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Node详情失败: %w", err)
	}

	return node, nil
}

// CordonNode 设置Node调度状态 (Cordon/Uncordon)
func (s *K8sNodeService) CordonNode(ctx context.Context, req request.CordonNodeRequest) error {
	client, err := GetClusterClient(req.ClusterName)
	if err != nil {
		return err
	}

	node, err := client.Clientset.CoreV1().Nodes().Get(ctx, req.NodeName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取Node失败: %w", err)
	}

	// 如果状态没有改变，直接返回
	if node.Spec.Unschedulable == req.Unschedulable {
		return nil
	}

	node.Spec.Unschedulable = req.Unschedulable
	_, err = client.Clientset.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("更新Node状态失败: %w", err)
	}

	return nil
}
