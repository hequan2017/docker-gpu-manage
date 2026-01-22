package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sNamespaceService Namespace服务结构体
type K8sNamespaceService struct{}

// GetNamespaceList 获取Namespace列表
func (s *K8sNamespaceService) GetNamespaceList(ctx context.Context, info *request.K8sNamespaceSearch) (*corev1.NamespaceList, error) {
	client, err := GetClusterClient(info.ClusterName)
	if err != nil {
		return nil, err
	}

	listOptions := metav1.ListOptions{}
	if info.Label != "" {
		listOptions.LabelSelector = info.Label
	}
	if info.Field != "" {
		listOptions.FieldSelector = info.Field
	}

	namespaceList, err := client.Clientset.CoreV1().Namespaces().List(ctx, &listOptions)
	if err != nil {
		return nil, fmt.Errorf("获取Namespace列表失败: %w", err)
	}

	return namespaceList, nil
}

// GetNamespace 获取单个Namespace详情
func (s *K8sNamespaceService) GetNamespace(ctx context.Context, clusterName, name string) (*corev1.Namespace, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	ns, err := client.Clientset.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Namespace详情失败: %w", err)
	}

	return ns, nil
}

// CreateNamespace 创建Namespace
func (s *K8sNamespaceService) CreateNamespace(ctx context.Context, clusterName, name string, labels map[string]string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
	}

	_, err = client.Clientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("创建Namespace失败: %w", err)
	}

	return nil
}

// DeleteNamespace 删除Namespace
func (s *K8sNamespaceService) DeleteNamespace(ctx context.Context, clusterName, name string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	err = client.Clientset.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("删除Namespace失败: %w", err)
	}

	return nil
}
