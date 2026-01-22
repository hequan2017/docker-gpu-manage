package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sServiceService K8s Service服务结构体
type K8sServiceService struct{}

// GetServiceList 获取Service列表
func (s *K8sServiceService) GetServiceList(ctx context.Context, info *request.K8sServiceSearch) (*corev1.ServiceList, error) {
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

	var svcList *corev1.ServiceList
	if info.Namespace == "all" {
		svcList, err = client.Clientset.CoreV1().Services("").List(ctx, &listOptions)
	} else {
		svcList, err = client.Clientset.CoreV1().Services(info.Namespace).List(ctx, &listOptions)
	}

	if err != nil {
		return nil, fmt.Errorf("获取Service列表失败: %w", err)
	}

	return svcList, nil
}

// GetService 获取单个Service详情
func (s *K8sServiceService) GetService(ctx context.Context, clusterName, namespace, name string) (*corev1.Service, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	svc, err := client.Clientset.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Service详情失败: %w", err)
	}

	return svc, nil
}

// DeleteService 删除Service
func (s *K8sServiceService) DeleteService(ctx context.Context, clusterName, namespace, name string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	err = client.Clientset.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("删除Service失败: %w", err)
	}

	return nil
}

// GetServiceEndpoints 获取Service的Endpoints
func (s *K8sServiceService) GetServiceEndpoints(ctx context.Context, clusterName, namespace, serviceName string) (*corev1.Endpoints, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	endpoints, err := client.Clientset.CoreV1().Endpoints(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Endpoints失败: %w", err)
	}

	return endpoints, nil
}
