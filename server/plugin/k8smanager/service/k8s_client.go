package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/utils"
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
	CreatedAt     time.Time // 连接创建时间
}

// k8sClientManager K8s客户端管理器单例
type k8sClientManager struct {
	clients    map[string]*K8sClient
	createdAt  map[string]time.Time // 连接创建时间映射
	mu         sync.RWMutex
	stopChan   chan struct{}      // 停止清理协程的通道
	cleanerOnce sync.Once         // 确保清理协程只启动一次
}

var clientManager = &k8sClientManager{
	clients:   make(map[string]*K8sClient),
	createdAt: make(map[string]time.Time),
	stopChan:  make(chan struct{}),
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
	// 解密 kubeconfig
	decryptedConfig, err := utils.Decrypt(cluster.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("解密kubeconfig失败: %w", err)
	}

	// 使用kubeconfig创建配置
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(decryptedConfig))
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

	now := time.Now()
	k8sClient := &K8sClient{
		Cluster:       cluster,
		Clientset:     clientset,
		MetricsClient: metricsClient,
		Config:        config,
		CreatedAt:     now,
	}

	// 缓存客户端
	clientManager.mu.Lock()
	clientManager.clients[cluster.Name] = k8sClient
	clientManager.createdAt[cluster.Name] = now
	clientManager.mu.Unlock()

	// 启动清理协程（如果还未启动）
	startCleaner()

	return k8sClient, nil
}

// RemoveClient 移除集群客户端缓存
func RemoveClient(clusterName string) {
	clientManager.mu.Lock()
	delete(clientManager.clients, clusterName)
	delete(clientManager.createdAt, clusterName)
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

// getClientTTL 获取客户端连接的TTL（生存时间）
func getClientTTL() time.Duration {
	// 从配置获取，默认为5分钟
	ttl := global.GVA_CONFIG.K8sManager.ClientTTL
	if ttl <= 0 {
		ttl = 300 // 默认5分钟
	}
	return time.Duration(ttl) * time.Second
}

// startCleaner 启动清理过期连接的协程
func startCleaner() {
	clientManager.cleanerOnce.Do(func() {
		go cleanupExpiredClients()
		global.GVA_LOG.Info("K8s客户端连接池清理协程已启动")
	})
}

// cleanupExpiredClients 定期清理过期的客户端连接
func cleanupExpiredClients() {
	ticker := time.NewTicker(1 * time.Minute) // 每分钟检查一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cleanupOnce()
		case <-clientManager.stopChan:
			global.GVA_LOG.Info("K8s客户端连接池清理协程已停止")
			return
		}
	}
}

// cleanupOnce 执行一次清理操作
func cleanupOnce() {
	ttl := getClientTTL()
	now := time.Now()
	var expiredCount int

	clientManager.mu.Lock()
	defer clientManager.mu.Unlock()

	for clusterName, createdAt := range clientManager.createdAt {
		if now.Sub(createdAt) > ttl {
			// 连接已过期，删除
			delete(clientManager.clients, clusterName)
			delete(clientManager.createdAt, clusterName)
			expiredCount++
		}
	}

	if expiredCount > 0 {
		global.GVA_LOG.Info("清理过期的K8s客户端连接",
			zap.Int("expired", expiredCount),
			zap.Int("remaining", len(clientManager.clients)))
	}
}

// StopCleaner 停止清理协程（通常在应用关闭时调用）
func StopCleaner() {
	close(clientManager.stopChan)
}

// GetClientStats 获取客户端连接池统计信息
func GetClientStats() map[string]interface{} {
	clientManager.mu.RLock()
	defer clientManager.mu.RUnlock()

	stats := map[string]interface{}{
		"total_connections": len(clientManager.clients),
		"ttl_seconds":       int(getClientTTL().Seconds()),
	}

	connections := make([]map[string]interface{}, 0, len(clientManager.clients))
	for name, client := range clientManager.clients {
		connections = append(connections, map[string]interface{}{
			"cluster_name": name,
			"created_at":   client.CreatedAt.Format(time.RFC3339),
			"age_seconds":  int(time.Since(client.CreatedAt).Seconds()),
		})
	}
	stats["connections"] = connections

	return stats
}
