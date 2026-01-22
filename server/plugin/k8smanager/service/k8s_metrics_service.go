package service

import (
	"context"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"go.uber.org/zap"
)

// K8sMetricsService K8s监控指标服务
type K8sMetricsService struct{}

// MetricsData 指标数据
type MetricsData struct {
	ClusterName string                 `json:"clusterName"`
	MetricType  string                 `json:"metricType"` // cluster, node, pod
	Timestamp   time.Time              `json:"timestamp"`
	Data        map[string]interface{} `json:"data"`
}

// ClusterMetrics 集群级别的指标
type ClusterMetrics struct {
	ClusterName    string    `json:"clusterName"`
	Timestamp      time.Time `json:"timestamp"`
	TotalNodes     int       `json:"totalNodes"`
	ReadyNodes     int       `json:"readyNodes"`
	TotalPods      int       `json:"totalPods"`
	RunningPods    int       `json:"runningPods"`
	TotalCpu       float64   `json:"totalCpu"`       // 总CPU核心数
	UsedCpu        float64   `json:"usedCpu"`        // 已使用CPU
	TotalMemory    int64     `json:"totalMemory"`    // 总内存(字节)
	UsedMemory     int64     `json:"usedMemory"`     // 已使用内存
	TotalStorage   int64     `json:"totalStorage"`   // 总存储(字节)
	UsedStorage    int64     `json:"usedStorage"`    // 已使用存储
	NetworkRxBytes int64     `json:"networkRxBytes"` // 网络接收字节数
	NetworkTxBytes int64     `json:"networkTxBytes"` // 网络发送字节数
}

// NodeMetrics 节点指标
type NodeMetrics struct {
	ClusterName    string    `json:"clusterName"`
	NodeName       string    `json:"nodeName"`
	Timestamp      time.Time `json:"timestamp"`
	CpuUsage       float64   `json:"cpuUsage"`       // CPU使用率(百分比)
	MemoryUsage    float64   `json:"memoryUsage"`    // 内存使用率(百分比)
	DiskUsage      float64   `json:"diskUsage"`      // 磁盘使用率(百分比)
	PodCount       int       `json:"podCount"`       // Pod数量
	ReadyPodCount  int       `json:"readyPodCount"`  // 就绪Pod数量
	NetworkRxBytes int64     `json:"networkRxBytes"` // 网络接收字节数
	NetworkTxBytes int64     `json:"networkTxBytes"` // 网络发送字节数
}

// PodMetrics Pod指标
type PodMetrics struct {
	ClusterName string    `json:"clusterName"`
	Namespace   string    `json:"namespace"`
	PodName     string    `json:"podName"`
	Timestamp   time.Time `json:"timestamp"`
	CpuUsage    float64   `json:"cpuUsage"`    // CPU使用量(毫核)
	MemoryUsage int64     `json:"memoryUsage"` // 内存使用量(字节)
	RestartCount int      `json:"restartCount"` // 重启次数
}

// metricsCollector 指标收集器
type metricsCollector struct {
	mu                sync.RWMutex
	clusterMetrics    map[string]*ClusterMetrics
	nodeMetrics       map[string][]*NodeMetrics     // key: clusterName
	podMetrics        map[string][]*PodMetrics      // key: clusterName/namespace
	lastCollectTime   map[string]time.Time
	collectInterval   time.Duration
	stopChan          chan struct{}
	collectorOnce     sync.Once
}

var metricsCollectorInstance = &metricsCollector{
	clusterMetrics:   make(map[string]*ClusterMetrics),
	nodeMetrics:      make(map[string][]*NodeMetrics),
	podMetrics:       make(map[string][]*PodMetrics),
	lastCollectTime:  make(map[string]time.Time),
	collectInterval:  30 * time.Second, // 默认30秒收集一次
	stopChan:         make(chan struct{}),
}

// CollectClusterMetrics 收集集群指标
func (s *K8sMetricsService) CollectClusterMetrics(ctx context.Context, clusterName string) (*ClusterMetrics, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	metrics := &ClusterMetrics{
		ClusterName: clusterName,
		Timestamp:   time.Now(),
	}

	// 获取节点列表和状态
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx, nil)
	if err == nil {
		metrics.TotalNodes = len(nodes.Items)
		for _, node := range nodes.Items {
			ready := false
			for _, condition := range node.Status.Conditions {
				if condition.Type == "Ready" && condition.Status == "True" {
					ready = true
					break
				}
			}
			if ready {
				metrics.ReadyNodes++
			}
		}
	}

	// 获取Pod统计
	pods, err := client.Clientset.CoreV1().Pods("").List(ctx, nil)
	if err == nil {
		metrics.TotalPods = len(pods.Items)
		for _, pod := range pods.Items {
			if pod.Status.Phase == "Running" {
				metrics.RunningPods++
			}
		}
	}

	// 获取节点资源信息
	if err == nil {
		for _, node := range nodes.Items {
			// CPU
			metrics.TotalCpu += node.Status.Capacity.Cpu().AsApproximateFloat64()
			metrics.UsedCpu += node.Status.Allocated.Cpu().AsApproximateFloat64()

			// 内存
			metrics.TotalMemory += node.Status.Capacity.Memory().Value()
			metrics.UsedMemory += node.Status.Allocated.Memory().Value()

			// 存储
			metrics.TotalStorage += node.Status.Capacity.StorageEphemeral().Value()
			metrics.UsedStorage += node.Status.Allocated.StorageEphemeral().Value()
		}
	}

	// 使用metrics客户端获取实时指标
	if client.MetricsClient != nil {
		// 获取节点指标
		nodeMetricsList, err := client.MetricsClient.MetricsV1beta1().NodeMetricses().List(ctx, nil)
		if err == nil {
			for _, nodeMetric := range nodeMetricsList.Items {
				metrics.UsedCpu += nodeMetric.Usage.Cpu().AsApproximateFloat64()
				metrics.UsedMemory += nodeMetric.Usage.Memory().Value()
			}
		}
	}

	// 存储到收集器
	metricsCollectorInstance.mu.Lock()
	metricsCollectorInstance.clusterMetrics[clusterName] = metrics
	metricsCollectorInstance.lastCollectTime[clusterName] = time.Now()
	metricsCollectorInstance.mu.Unlock()

	return metrics, nil
}

// CollectNodeMetrics 收集节点指标
func (s *K8sMetricsService) CollectNodeMetrics(ctx context.Context, clusterName string) ([]*NodeMetrics, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	var nodeMetricsList []*NodeMetrics

	// 获取节点列表
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 获取节点指标
	var metricsList *map[string]interface{}
	if client.MetricsClient != nil {
		nodeMetrics, err := client.MetricsClient.MetricsV1beta1().NodeMetricses().List(ctx, nil)
		if err == nil {
			metricsMap := make(map[string]interface{})
			for _, m := range nodeMetrics.Items {
				metricsMap[m.Name] = m
			}
			metricsList = &metricsMap
		}
	}

	// 构建节点指标
	for _, node := range nodes.Items {
		metrics := &NodeMetrics{
			ClusterName: clusterName,
			NodeName:    node.Name,
			Timestamp:   time.Now(),
		}

		// 获取Pod数量
		pods, err := client.Clientset.CoreV1().Pods("").List(ctx, nil)
		if err == nil {
			for _, pod := range pods.Items {
				if pod.Spec.NodeName == node.Name {
					metrics.PodCount++
					if pod.Status.Phase == "Running" {
						metrics.ReadyPodCount++
					}
				}
			}
		}

		// 从metrics获取CPU和内存使用率
		if metricsList != nil {
			if nodeMetric, ok := (*metricsList)[node.Name]; ok {
				if nm, ok := nodeMetric.(interface{ GetUsage() interface{} }); ok {
					usage := nm.GetUsage()
					// 这里需要根据实际的metrics API结构解析
					// 简化处理，实际使用时需要完整解析
				}
			}
		}

		nodeMetricsList = append(nodeMetricsList, metrics)
	}

	// 存储到收集器
	metricsCollectorInstance.mu.Lock()
	metricsCollectorInstance.nodeMetrics[clusterName] = nodeMetricsList
	metricsCollectorInstance.mu.Unlock()

	return nodeMetricsList, nil
}

// CollectPodMetrics 收集Pod指标
func (s *K8sMetricsService) CollectPodMetrics(ctx context.Context, clusterName, namespace string) ([]*PodMetrics, error) {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return nil, err
	}

	var podMetricsList []*PodMetrics

	// 获取Pod列表
	pods, err := client.Clientset.CoreV1().Pods(namespace).List(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 获取Pod指标
	if client.MetricsClient != nil {
		podMetrics, err := client.MetricsClient.MetricsV1beta1().PodMetricses(namespace).List(ctx, nil)
		if err == nil {
			for _, podMetric := range podMetrics.Items {
				metrics := &PodMetrics{
					ClusterName: clusterName,
					Namespace:   namespace,
					PodName:     podMetric.Name,
					Timestamp:   time.Now(),
				}

				// 获取容器指标总和
				for _, container := range podMetric.Containers {
					metrics.CpuUsage += container.Usage.Cpu().AsApproximateFloat64()
					metrics.MemoryUsage += container.Usage.Memory().Value()
				}

				// 获取重启次数
				for _, pod := range pods.Items {
					if pod.Name == podMetric.Name {
						for _, cs := range pod.Status.ContainerStatuses {
							metrics.RestartCount += int(cs.RestartCount)
						}
						break
					}
				}

				podMetricsList = append(podMetricsList, metrics)
			}
		}
	}

	// 存储到收集器
	key := clusterName + "/" + namespace
	metricsCollectorInstance.mu.Lock()
	metricsCollectorInstance.podMetrics[key] = podMetricsList
	metricsCollectorInstance.mu.Unlock()

	return podMetricsList, nil
}

// GetClusterMetrics 获取缓存的集群指标
func (s *K8sMetricsService) GetClusterMetrics(clusterName string) (*ClusterMetrics, error) {
	metricsCollectorInstance.mu.RLock()
	defer metricsCollectorInstance.mu.RUnlock()

	if metrics, ok := metricsCollectorInstance.clusterMetrics[clusterName]; ok {
		return metrics, nil
	}
	return nil, global.GVA_ERROR{ErrorMsg: "集群指标不存在"}
}

// GetNodeMetrics 获取缓存的节点指标
func (s *K8sMetricsService) GetNodeMetrics(clusterName string) ([]*NodeMetrics, error) {
	metricsCollectorInstance.mu.RLock()
	defer metricsCollectorInstance.mu.RUnlock()

	if metrics, ok := metricsCollectorInstance.nodeMetrics[clusterName]; ok {
		return metrics, nil
	}
	return nil, global.GVA_ERROR{ErrorMsg: "节点指标不存在"}
}

// GetPodMetrics 获取缓存的Pod指标
func (s *K8sMetricsService) GetPodMetrics(clusterName, namespace string) ([]*PodMetrics, error) {
	metricsCollectorInstance.mu.RLock()
	defer metricsCollectorInstance.mu.RUnlock()

	key := clusterName + "/" + namespace
	if metrics, ok := metricsCollectorInstance.podMetrics[key]; ok {
		return metrics, nil
	}
	return nil, global.GVA_ERROR{ErrorMsg: "Pod指标不存在"}
}

// StartAutoCollector 启动自动收集协程
func (s *K8sMetricsService) StartAutoCollector() {
	metricsCollectorInstance.collectorOnce.Do(func() {
		go s.autoCollect()
		global.GVA_LOG.Info("K8s监控指标自动收集已启动")
	})
}

// autoCollect 自动收集指标
func (s *K8sMetricsService) autoCollect() {
	ticker := time.NewTicker(metricsCollectorInstance.collectInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.collectAllClusterMetrics()
		case <-metricsCollectorInstance.stopChan:
			global.GVA_LOG.Info("K8s监控指标自动收集已停止")
			return
		}
	}
}

// collectAllClusterMetrics 收集所有集群的指标
func (s *K8sMetricsService) collectAllClusterMetrics() {
	ctx := context.Background()

	// 获取所有集群
	var clusters []model.K8sCluster
	err := global.GVA_DB.Find(&clusters).Error
	if err != nil {
		global.GVA_LOG.Error("获取集群列表失败", zap.Error(err))
		return
	}

	// 为每个集群收集指标
	for _, cluster := range clusters {
		if cluster.Status != "online" {
			continue
		}

		go func(clusterName string) {
			if _, err := s.CollectClusterMetrics(ctx, clusterName); err != nil {
				global.GVA_LOG.Warn("收集集群指标失败",
					zap.String("cluster", clusterName),
					zap.Error(err))
			}
		}(cluster.Name)
	}
}

// StopAutoCollector 停止自动收集
func (s *K8sMetricsService) StopAutoCollector() {
	close(metricsCollectorInstance.stopChan)
}

// GetMetricsSummary 获取指标摘要
func (s *K8sMetricsService) GetMetricsSummary() map[string]interface{} {
	metricsCollectorInstance.mu.RLock()
	defer metricsCollectorInstance.mu.RUnlock()

	summary := map[string]interface{}{
		"total_clusters": len(metricsCollectorInstance.clusterMetrics),
		"total_nodes":    0,
		"total_pods":     0,
		"last_collect":   metricsCollectorInstance.lastCollectTime,
	}

	for _, metrics := range metricsCollectorInstance.clusterMetrics {
		summary["total_nodes"] = summary["total_nodes"].(int) + metrics.TotalNodes
		summary["total_pods"] = summary["total_pods"].(int) + metrics.TotalPods
	}

	return summary
}
