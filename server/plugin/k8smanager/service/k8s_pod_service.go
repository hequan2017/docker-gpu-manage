package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sPodService Pod服务结构体
type K8sPodService struct{}

// GetPodList 获取Pod列表
func (s *K8sPodService) GetPodList(ctx context.Context, info *request.K8sPodSearch) (*corev1.PodList, error) {
	client, err := GetClusterClient(info.ClusterName)
	if err != nil {
		return nil, err
	}

	// 构建列表选项
	listOptions := metav1.ListOptions{}
	if info.Label != "" {
		listOptions.LabelSelector = info.Label
	}
	if info.Field != "" {
		listOptions.FieldSelector = info.Field
	}

	var podList *corev1.PodList
	if info.ShowAll || info.Namespace == "all" {
		podList, err = client.Clientset.CoreV1().Pods("").List(ctx, listOptions)
	} else {
		podList, err = client.Clientset.CoreV1().Pods(info.Namespace).List(ctx, listOptions)
	}

	if err != nil {
		return nil, fmt.Errorf("获取Pod列表失败: %w", err)
	}

	return podList, nil
}

// GetPod 获取单个Pod详情
func (s *K8sPodService) GetPod(ctx context.Context, clusterName, namespace, podName string) (*corev1.Pod, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	pod, err := client.Clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Pod详情失败: %w", err)
	}

	return pod, nil
}

// DeletePod 删除Pod
func (s *K8sPodService) DeletePod(ctx context.Context, clusterName, namespace, podName string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	err = client.Clientset.CoreV1().Pods(namespace).Delete(ctx, podName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("删除Pod失败: %w", err)
	}

	return nil
}

// GetPodLog 获取Pod日志
func (s *K8sPodService) GetPodLog(ctx context.Context, req *request.GetPodLogRequest) (string, error) {
	client, err := GetClusterClient(req.ClusterName)
	if err != nil {
		return "", err
	}

	// 构建日志选项
	logOptions := &corev1.PodLogOptions{
		Follow:     req.Follow,
		Previous:   req.Previous,
		Timestamps: req.Timestamps,
	}

	if req.TailLines > 0 {
		logOptions.TailLines = &req.TailLines
	}
	if req.SinceSeconds > 0 {
		logOptions.SinceSeconds = &req.SinceSeconds
	}
	if req.Container != "" {
		logOptions.Container = req.Container
	}

	// 获取日志流
	reqLogs := client.Clientset.CoreV1().Pods(req.Namespace).GetLogs(req.PodName, logOptions)

	// 读取日志内容
	logData, err := reqLogs.DoRaw(ctx)
	if err != nil {
		return "", fmt.Errorf("获取Pod日志失败: %w", err)
	}

	return string(logData), nil
}

// GetPodContainers 获取Pod中的容器列表
func (s *K8sPodService) GetPodContainers(ctx context.Context, clusterName, namespace, podName string) ([]corev1.Container, []corev1.Container, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, nil, err
	}

	pod, err := client.Clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("获取Pod信息失败: %w", err)
	}

	return pod.Spec.Containers, pod.Spec.InitContainers, nil
}

// GetPodEvents 获取Pod相关事件
func (s *K8sPodService) GetPodEvents(ctx context.Context, clusterName, namespace, podName string) (*corev1.EventList, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	// 构建字段选择器，只获取该Pod的事件
	fieldSelector := fmt.Sprintf("involvedObject.name=%s,involvedObject.kind=Pod", podName)

	events, err := client.Clientset.CoreV1().Events(namespace).List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("获取Pod事件失败: %w", err)
	}

	return events, nil
}
