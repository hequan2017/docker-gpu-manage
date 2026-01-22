package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// K8sClusterSearch K8s集群搜索结构体
type K8sClusterSearch struct {
	request.PageInfo
	Name     string `json:"name" form:"name"`         // 集群名称
	Status   string `json:"status" form:"status"`     // 集群状态
	Provider string `json:"provider" form:"provider"` // 云服务商
	Region   string `json:"region" form:"region"`     // 区域
}

// CreateK8sClusterRequest 创建K8s集群请求
type CreateK8sClusterRequest struct {
	Name        string `json:"name" binding:"required"`         // 集群名称
	KubeConfig  string `json:"kubeConfig" binding:"required"`   // kubeconfig配置内容
	Endpoint    string `json:"endpoint"`                        // API Server地址（可选，从kubeconfig解析）
	Description string `json:"description"`                     // 集群描述
	Region      string `json:"region"`                           // 区域
	Provider    string `json:"provider"`                         // 云服务商
	IsDefault   bool   `json:"isDefault"`                        // 是否默认集群
}

// UpdateK8sClusterRequest 更新K8s集群请求
type UpdateK8sClusterRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	KubeConfig  string `json:"kubeConfig" binding:"required"`
	Endpoint    string `json:"endpoint"`
	Description string `json:"description"`
	Region      string `json:"region"`
	Provider    string `json:"provider"`
	IsDefault   bool   `json:"isDefault"`
}

// K8sResourceSearch K8s资源通用搜索结构体
type K8sResourceSearch struct {
	ClusterName string `json:"clusterName" form:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" form:"namespace"`                        // 命名空间，默认为all
	Label       string `json:"label" form:"label"`                                // 标签过滤
	Field       string `json:"field" form:"field"`                                // 字段选择器
}

// K8sPodSearch Pod搜索
type K8sPodSearch struct {
	K8sResourceSearch
	ShowAll bool `json:"showAll" form:"showAll"` // 是否显示所有命名空间
}

// K8sDeploymentSearch Deployment搜索
type K8sDeploymentSearch struct {
	K8sResourceSearch
}

// K8sServiceSearch Service搜索
type K8sServiceSearch struct {
	K8sResourceSearch
}

// K8sNamespaceSearch Namespace搜索
type K8sNamespaceSearch struct {
	K8sResourceSearch
}

// DeleteK8sResourceRequest 删除K8s资源请求
type DeleteK8sResourceRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" binding:"required"`   // 命名空间
	Name        string `json:"name" binding:"required"`       // 资源名称
	Resource    string `json:"resource" binding:"required"`   // 资源类型: pod, deployment, service等
}

// ScaleDeploymentRequest 扩缩容Deployment请求
type ScaleDeploymentRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" binding:"required"`   // 命名空间
	Name        string `json:"name" binding:"required"`       // Deployment名称
	Replicas    int32 `json:"replicas" binding:"required"`     // 副本数
}

// RestartDeploymentRequest 重启Deployment请求
type RestartDeploymentRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" binding:"required"`   // 命名空间
	Name        string `json:"name" binding:"required"`       // Deployment名称
}

// GetPodLogRequest 获取Pod日志请求
type GetPodLogRequest struct {
	ClusterName  string `json:"clusterName" binding:"required"`  // 集群名称
	Namespace    string `json:"namespace" binding:"required"`    // 命名空间
	PodName      string `json:"podName" binding:"required"`      // Pod名称
	Container    string `json:"container"`                       // 容器名称（默认为第一个容器）
	TailLines    int64 `json:"tailLines"`                        // 返回最近的行数
	Follow       bool   `json:"follow"`                          // 是否持续跟踪日志
	Previous     bool   `json:"previous"`                        // 是否查看之前容器实例的日志
	Timestamps   bool   `json:"timestamps"`                      // 是否显示时间戳
	SinceSeconds int64 `json:"sinceSeconds"`                     // 返回最近多少秒的日志
}

// ExecPodRequest 在Pod中执行命令请求
type ExecPodRequest struct {
	ClusterName string   `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string   `json:"namespace" binding:"required"`   // 命名空间
	PodName     string   `json:"podName" binding:"required"`     // Pod名称
	Container   string   `json:"container"`                      // 容器名称
	Command     []string `json:"command" binding:"required"`     // 要执行的命令
}

// GetPodTerminalRequest 获取Pod终端请求
type GetPodTerminalRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" binding:"required"`   // 命名空间
	PodName     string `json:"podName" binding:"required"`     // Pod名称
	Container   string `json:"container"`                      // 容器名称
}

// GetYamlRequest 获取资源YAML请求
type GetYamlRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace" binding:"required"`   // 命名空间
	Kind        string `json:"kind" binding:"required"`       // 资源类型
	Name        string `json:"name" binding:"required"`       // 资源名称
}

// ApplyYamlRequest 应用YAML请求
type ApplyYamlRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace"`                      // 目标命名空间
	Yaml        string `json:"yaml" binding:"required"`        // YAML内容
}

// GetNodeMetricsRequest 获取节点指标请求
type GetNodeMetricsRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	NodeName    string `json:"nodeName"`                       // 节点名称（不指定则返回所有）
}

// GetPodMetricsRequest 获取Pod指标请求
type GetPodMetricsRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace"`                      // 命名空间（不指定则返回所有）
	PodName     string `json:"podName"`                        // Pod名称
}

// GetEventsRequest 获取事件请求
type GetEventsRequest struct {
	ClusterName string `json:"clusterName" binding:"required"` // 集群名称
	Namespace   string `json:"namespace"`                      // 命名空间（不指定则返回所有）
	Kind        string `json:"kind"`                           // 资源类型过滤
	Name        string `json:"name"`                           // 资源名称过滤
	Limit       int64 `json:"limit"`                           // 返回数量限制
}
