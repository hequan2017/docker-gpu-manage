package service

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sDeploymentService Deployment服务结构体
type K8sDeploymentService struct{}

// GetDeploymentList 获取Deployment列表
func (s *K8sDeploymentService) GetDeploymentList(ctx context.Context, info *request.K8sDeploymentSearch) (*appsv1.DeploymentList, error) {
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

	var deployList *appsv1.DeploymentList
	if info.Namespace == "all" {
		deployList, err = client.Clientset.AppsV1().Deployments("").List(ctx, listOptions)
	} else {
		deployList, err = client.Clientset.AppsV1().Deployments(info.Namespace).List(ctx, listOptions)
	}

	if err != nil {
		return nil, fmt.Errorf("获取Deployment列表失败: %w", err)
	}

	return deployList, nil
}

// GetDeployment 获取单个Deployment详情
func (s *K8sDeploymentService) GetDeployment(ctx context.Context, clusterName, namespace, name string) (*appsv1.Deployment, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	deploy, err := client.Clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Deployment详情失败: %w", err)
	}

	return deploy, nil
}

// ScaleDeployment 扩缩容Deployment
func (s *K8sDeploymentService) ScaleDeployment(ctx context.Context, req *request.ScaleDeploymentRequest) (*appsv1.Deployment, error) {
	client, err := GetClusterClient(req.ClusterName)
	if err != nil {
		return nil, err
	}

	// 获取当前Deployment
	deploy, err := client.Clientset.AppsV1().Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Deployment失败: %w", err)
	}

	// 设置副本数
	deploy.Spec.Replicas = &req.Replicas

	// 更新Deployment
	updatedDeploy, err := client.Clientset.AppsV1().Deployments(req.Namespace).Update(ctx, deploy, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("扩缩容Deployment失败: %w", err)
	}

	return updatedDeploy, nil
}

// RestartDeployment 重启Deployment（通过触发Rollout）
func (s *K8sDeploymentService) RestartDeployment(ctx context.Context, req *request.RestartDeploymentRequest) error {
	client, err := GetClusterClient(req.ClusterName)
	if err != nil {
		return err
	}

	// 获取当前Deployment
	deploy, err := client.Clientset.AppsV1().Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取Deployment失败: %w", err)
	}

	// 通过添加注解触发重启
	if deploy.Spec.Template.Annotations == nil {
		deploy.Spec.Template.Annotations = make(map[string]string)
	}
	deploy.Spec.Template.Annotations["kubectl.kubernetes.io/restartedAt"] = fmt.Sprintf("%d", metav1.Now().Unix())

	_, err = client.Clientset.AppsV1().Deployments(req.Namespace).Update(ctx, deploy, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("重启Deployment失败: %w", err)
	}

	return nil
}

// DeleteDeployment 删除Deployment
func (s *K8sDeploymentService) DeleteDeployment(ctx context.Context, clusterName, namespace, name string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	err = client.Clientset.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("删除Deployment失败: %w", err)
	}

	return nil
}

// GetDeploymentPods 获取Deployment关联的Pods
func (s *K8sDeploymentService) GetDeploymentPods(ctx context.Context, clusterName, namespace, deploymentName string) (*corev1.PodList, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	// 获取Deployment以获取标签选择器
	deploy, err := client.Clientset.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Deployment失败: %w", err)
	}

	// 使用标签选择器获取Pods
	listOptions := metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(deploy.Spec.Selector),
	}

	pods, err := client.Clientset.CoreV1().Pods(namespace).List(ctx, listOptions)
	if err != nil {
		return nil, fmt.Errorf("获取Pod列表失败: %w", err)
	}

	return pods, nil
}
