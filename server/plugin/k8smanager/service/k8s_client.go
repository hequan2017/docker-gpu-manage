package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

var (
	// ErrClusterNotFound 集群不存在错误
	ErrClusterNotFound = errors.New("集群不存在")
	// ErrClusterOffline 集群离线错误
	ErrClusterOffline = errors.New("集群离线")
	// ErrInvalidKubeConfig 无效的kubeconfig错误
	ErrInvalidKubeConfig = errors.New("无效的kubeconfig配置")
)

// K8sClient K8s客户端管理器
type K8sClient struct {
	Cluster       *model.K8sCluster
	Clientset     *kubernetes.Clientset
	MetricsClient *metricsv.Clientset
	Config        *rest.Config
}

// k8sClientManager K8s客户端管理器单例
type k8sClientManager struct {
	clients map[string]*K8sClient
	mu      sync.RWMutex
}

var clientManager = &k8sClientManager{
	clients: make(map[string]*K8sClient),
}

// GetClusterClient 获取集群客户端
func GetClusterClient(clusterName string) (*K8sClient, error) {
	clientManager.mu.RLock()
	client, exists := clientManager.clients[clusterName]
	clientManager.mu.RUnlock()

	if exists {
		return client, nil
	}

	// 从数据库加载集群配置
	var cluster model.K8sCluster
	err := global.GVA_DB.Where("name = ?", clusterName).First(&cluster).Error
	if err != nil {
		return nil, ErrClusterNotFound
	}

	// 创建新客户端
	return createClient(&cluster)
}

// createClient 创建新的K8s客户端
func createClient(cluster *model.K8sCluster) (*K8sClient, error) {
	// 使用kubeconfig创建配置
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidKubeConfig, err)
	}

	// 创建clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("创建kubernetes客户端失败: %w", err)
	}

	// 创建metrics客户端
	metricsClient, err := metricsv.NewForConfig(config)
	if err != nil {
		// metrics客户端创建失败不影响主功能
		global.GVA_LOG.Warn("创建metrics客户端失败", zap.Error(err))
	}

	k8sClient := &K8sClient{
		Cluster:       cluster,
		Clientset:     clientset,
		MetricsClient: metricsClient,
		Config:        config,
	}

	// 缓存客户端
	clientManager.mu.Lock()
	clientManager.clients[cluster.Name] = k8sClient
	clientManager.mu.Unlock()

	return k8sClient, nil
}

// RemoveClient 移除集群客户端缓存
func RemoveClient(clusterName string) {
	clientManager.mu.Lock()
	delete(clientManager.clients, clusterName)
	clientManager.mu.Unlock()
}

// RefreshClient 刷新集群客户端
func RefreshClient(clusterName string) (*K8sClient, error) {
	RemoveClient(clusterName)
	return GetClusterClient(clusterName)
}

// CheckClusterHealth 检查集群健康状态
func CheckClusterHealth(client *K8sClient) error {
	// 通过尝试访问服务器版本API来检查健康状态
	_, err := client.Clientset.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("集群健康检查失败: %w", err)
	}
	return nil
}
