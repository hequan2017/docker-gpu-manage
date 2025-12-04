package instance

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"go.uber.org/zap"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// DockerService Docker服务
type DockerService struct{}

// ContainerConfig 容器配置
type ContainerConfig struct {
	Image              string // 镜像地址
	Name               string // 容器名称
	CPUCores           int64  // CPU核心数
	MemoryGB           int64  // 内存大小(GB)
	SystemDiskGB       int64  // 系统盘大小(GB)
	DataDiskGB         int64  // 数据盘大小(GB)
	GPUCount           int64  // GPU数量
	SupportMemorySplit bool   // 是否支持显存分割
	MemoryCapacity     int64  // 显存容量(GB)
	PerCardCapacity    int64  // 节点单卡显存容量(GB)
}

// CreateDockerClient 创建Docker客户端
func (d *DockerService) CreateDockerClient(node *computenode.ComputeNode) (*client.Client, error) {
	if node.DockerAddress == nil || *node.DockerAddress == "" {
		return nil, fmt.Errorf("节点Docker连接地址为空")
	}

	dockerHost := *node.DockerAddress

	// 检查是否使用TLS
	useTLS := node.UseTls != nil && *node.UseTls

	if useTLS {
		// 使用TLS连接
		if node.CaCert == nil || node.ClientCert == nil || node.ClientKey == nil {
			return nil, fmt.Errorf("TLS证书配置不完整")
		}

		// 创建TLS配置
		tlsConfig, err := d.createTLSConfig(*node.CaCert, *node.ClientCert, *node.ClientKey)
		if err != nil {
			return nil, fmt.Errorf("创建TLS配置失败: %v", err)
		}

		// 创建HTTP客户端
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
			Timeout: 30 * time.Second,
		}

		return client.NewClientWithOpts(
			client.WithHost(dockerHost),
			client.WithHTTPClient(httpClient),
			client.WithAPIVersionNegotiation(),
		)
	}

	// 不使用TLS
	return client.NewClientWithOpts(
		client.WithHost(dockerHost),
		client.WithAPIVersionNegotiation(),
	)
}

// createTLSConfig 创建TLS配置
func (d *DockerService) createTLSConfig(caCert, clientCert, clientKey string) (*tls.Config, error) {
	// 加载CA证书
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM([]byte(caCert)) {
		return nil, fmt.Errorf("无法解析CA证书")
	}

	// 加载客户端证书
	cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
	if err != nil {
		return nil, fmt.Errorf("加载客户端证书失败: %v", err)
	}

	return &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

// CreateContainer 创建容器
func (d *DockerService) CreateContainer(ctx context.Context, node *computenode.ComputeNode, config *ContainerConfig) (containerID string, err error) {
	// 创建Docker客户端
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return "", fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	// 构建容器配置
	containerConfig := &container.Config{
		Image: config.Image,
		Labels: map[string]string{
			"managed-by": "docker-gpu-manage",
			"instance":   config.Name,
		},
	}

	// 构建主机配置
	hostConfig := &container.HostConfig{}

	// CPU配置: --cpus=N
	if config.CPUCores > 0 {
		// NanoCPUs 是以纳秒为单位的CPU配额，1核 = 1e9 纳秒
		hostConfig.NanoCPUs = config.CPUCores * 1e9
	}

	// 内存配置: --memory=Ng
	if config.MemoryGB > 0 {
		// Memory 是以字节为单位
		hostConfig.Memory = config.MemoryGB * 1024 * 1024 * 1024
	}

	// 数据盘配置: 创建命名卷并挂载到 /data
	if config.DataDiskGB > 0 {
		volumeName := fmt.Sprintf("%s-data", config.Name)

		// 创建数据卷
		_, err = cli.VolumeCreate(ctx, volume.CreateOptions{
			Name: volumeName,
			Labels: map[string]string{
				"managed-by": "docker-gpu-manage",
				"instance":   config.Name,
			},
		})
		if err != nil {
			global.GVA_LOG.Warn("创建数据卷失败，可能已存在", zap.Error(err))
		}

		// 挂载数据卷到 /data
		hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
			Type:   mount.TypeVolume,
			Source: volumeName,
			Target: "/data",
		})
	}

	// 显存分割配置: 如果支持显存分割，添加相关卷挂载和环境变量
	if config.SupportMemorySplit && config.MemoryCapacity > 0 && config.PerCardCapacity > 0 {
		// 从节点读取 HAMi-core 目录路径，如果未配置则使用默认路径
		hamiCorePath := "/root/HAMi-core/build"
		if node.HamiCore != nil && *node.HamiCore != "" {
			hamiCorePath = *node.HamiCore
		}
		// 挂载 HAMi 库目录: -v {hamiCorePath}:/libvgpu/build
		hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: hamiCorePath,
			Target: "/libvgpu/build",
		})

		// 初始化环境变量（如果还没有）
		if containerConfig.Env == nil {
			containerConfig.Env = []string{}
		}

		// 添加环境变量: LD_PRELOAD=/libvgpu/build/libvgpu.so
		containerConfig.Env = append(containerConfig.Env, "LD_PRELOAD=/libvgpu/build/libvgpu.so")

		// 添加环境变量: CUDA_DEVICE_MEMORY_LIMIT=产品规格中的显存容量 g
		containerConfig.Env = append(containerConfig.Env, fmt.Sprintf("CUDA_DEVICE_MEMORY_LIMIT=%dg", config.MemoryCapacity))

		// 计算 CUDA_DEVICE_SM_LIMIT: 产品规格中的显存容量 / 主机本来的单卡总容量（整数）
		smLimit := int64(0)
		if config.PerCardCapacity > 0 {
			smLimit = config.MemoryCapacity / config.PerCardCapacity
		}
		// 添加环境变量: CUDA_DEVICE_SM_LIMIT
		containerConfig.Env = append(containerConfig.Env, fmt.Sprintf("CUDA_DEVICE_SM_LIMIT=%d", smLimit))

		global.GVA_LOG.Info("添加显存分割配置",
			zap.String("容器名称", config.Name),
			zap.Int64("显存容量", config.MemoryCapacity),
			zap.Int64("单卡容量", config.PerCardCapacity),
			zap.Int64("SM限制", smLimit))
	}

	// GPU配置: --gpus N
	if config.GPUCount > 0 {
		hostConfig.DeviceRequests = []container.DeviceRequest{
			{
				Driver: "nvidia",
				Count:  int(config.GPUCount),
				Capabilities: [][]string{
					{"gpu"},
				},
			},
		}
	}

	// 系统盘配置: --storage-opt overlay2.size=NG
	// 如果配置了系统盘大小，尝试添加存储选项
	if config.SystemDiskGB > 0 {
		hostConfig.StorageOpt = map[string]string{
			"overlay2.size": fmt.Sprintf("%dG", config.SystemDiskGB),
		}
	}

	// 尝试创建容器（先尝试带系统盘限制）
	var resp container.CreateResponse
	resp, err = cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, config.Name)

	// 如果创建失败且设置了系统盘参数，尝试不带系统盘参数重试
	if err != nil && config.SystemDiskGB > 0 {
		global.GVA_LOG.Warn("使用系统盘参数创建容器失败，尝试不带系统盘参数重试",
			zap.Error(err),
			zap.Int64("systemDiskGB", config.SystemDiskGB))

		// 清除系统盘参数
		hostConfig.StorageOpt = nil

		// 重试创建容器
		resp, err = cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, config.Name)
	}

	if err != nil {
		return "", fmt.Errorf("创建容器失败: %v", err)
	}

	// 启动容器
	if err = cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		// 启动失败，删除已创建的容器
		_ = cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
		return "", fmt.Errorf("启动容器失败: %v", err)
	}

	return resp.ID, nil
}

// DeleteContainer 删除容器及其数据卷
func (d *DockerService) DeleteContainer(ctx context.Context, node *computenode.ComputeNode, containerID string, containerName string) error {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	// 先获取容器信息，提取所有挂载的命名卷
	var volumeNames []string
	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err == nil {
		// 从容器挂载信息中提取所有命名卷
		for _, m := range inspect.Mounts {
			if m.Type == mount.TypeVolume && m.Name != "" {
				volumeNames = append(volumeNames, m.Name)
			}
		}
	} else {
		// 如果无法获取容器信息，尝试使用默认的命名规则
		global.GVA_LOG.Warn("获取容器信息失败，使用默认卷名", zap.Error(err))
		if containerName != "" {
			volumeNames = append(volumeNames, fmt.Sprintf("%s-data", containerName))
		}
	}

	// 先停止容器
	timeout := 10
	_ = cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeout})

	// 删除容器
	err = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{
		Force:         true,
		RemoveVolumes: true, // 同时删除匿名卷
	})
	if err != nil {
		global.GVA_LOG.Warn("删除容器失败", zap.Error(err))
		// 即使容器删除失败，也继续删除数据卷
	}

	// 删除所有挂载的命名数据卷
	for _, volumeName := range volumeNames {
		err = cli.VolumeRemove(ctx, volumeName, true)
		if err != nil {
			// 记录警告但不中断流程，因为卷可能已经被删除或不存在
			global.GVA_LOG.Warn("删除数据卷失败", zap.String("volume", volumeName), zap.Error(err))
		} else {
			global.GVA_LOG.Info("成功删除数据卷", zap.String("volume", volumeName))
		}
	}

	return nil
}

// StopContainer 停止容器
func (d *DockerService) StopContainer(ctx context.Context, node *computenode.ComputeNode, containerID string) error {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	timeout := 30
	return cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeout})
}

// StartContainer 启动容器
func (d *DockerService) StartContainer(ctx context.Context, node *computenode.ComputeNode, containerID string) error {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	return cli.ContainerStart(ctx, containerID, container.StartOptions{})
}

// RestartContainer 重启容器
func (d *DockerService) RestartContainer(ctx context.Context, node *computenode.ComputeNode, containerID string) error {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	timeout := 30
	return cli.ContainerRestart(ctx, containerID, container.StopOptions{Timeout: &timeout})
}

// GetContainerStatus 获取容器状态
func (d *DockerService) GetContainerStatus(ctx context.Context, node *computenode.ComputeNode, containerID string) (string, error) {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return "", fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return "", fmt.Errorf("获取容器信息失败: %v", err)
	}

	return inspect.State.Status, nil
}

// BuildContainerConfig 根据实例信息构建容器配置
func (d *DockerService) BuildContainerConfig(image *imageregistry.ImageRegistry, spec *product.ProductSpec, node *computenode.ComputeNode, instanceName string) *ContainerConfig {
	config := &ContainerConfig{
		Name: instanceName,
	}

	// 镜像地址
	if image.Address != nil {
		config.Image = *image.Address
	}

	// CPU核心数
	if spec.CpuCores != nil {
		config.CPUCores = *spec.CpuCores
	}

	// 内存大小
	if spec.MemoryGb != nil {
		config.MemoryGB = *spec.MemoryGb
	}

	// 系统盘大小
	if spec.SystemDiskGb != nil {
		config.SystemDiskGB = *spec.SystemDiskGb
	}

	// 数据盘大小
	if spec.DataDiskGb != nil {
		config.DataDiskGB = *spec.DataDiskGb
	}

	// GPU数量
	if spec.GpuCount != nil {
		config.GPUCount = *spec.GpuCount
	}

	// 显存分割相关配置
	if spec.SupportMemorySplit != nil && *spec.SupportMemorySplit {
		config.SupportMemorySplit = true
		if spec.MemoryCapacity != nil {
			config.MemoryCapacity = *spec.MemoryCapacity
		}
		// 计算节点单卡显存容量
		if node.MemoryCapacity != nil && node.GpuCount != nil && *node.GpuCount > 0 {
			config.PerCardCapacity = *node.MemoryCapacity / *node.GpuCount
		}
	}

	return config
}

// GenerateInstanceName 生成实例名称
func (d *DockerService) GenerateInstanceName(baseName string, instanceID uint) string {
	// 清理名称，只保留字母数字和横杠
	cleanName := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return '-'
	}, baseName)

	// 移除连续的横杠
	for strings.Contains(cleanName, "--") {
		cleanName = strings.ReplaceAll(cleanName, "--", "-")
	}

	// 移除首尾横杠
	cleanName = strings.Trim(cleanName, "-")

	if cleanName == "" {
		cleanName = "instance"
	}

	return fmt.Sprintf("%s-%d-%d", cleanName, instanceID, time.Now().Unix())
}

// SyncContainerStatus 同步容器状态到数据库
func (d *DockerService) SyncContainerStatus(ctx context.Context, instanceID uint) error {
	var inst instanceModel.Instance
	if err := global.GVA_DB.Where("id = ?", instanceID).First(&inst).Error; err != nil {
		return err
	}

	if inst.ContainerId == nil || *inst.ContainerId == "" || inst.NodeId == nil {
		return nil
	}

	var node computenode.ComputeNode
	if err := global.GVA_DB.Where("id = ?", *inst.NodeId).First(&node).Error; err != nil {
		return err
	}

	status, err := d.GetContainerStatus(ctx, &node, *inst.ContainerId)
	if err != nil {
		status = "unknown"
	}

	return global.GVA_DB.Model(&inst).Update("container_status", status).Error
}

// GetContainerLogs 获取容器日志
func (d *DockerService) GetContainerLogs(ctx context.Context, node *computenode.ComputeNode, containerID string, tail string) (string, error) {
	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return "", fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       tail,
		Timestamps: true,
	}

	logs, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return "", fmt.Errorf("获取容器日志失败: %v", err)
	}
	defer logs.Close()

	// 读取日志内容
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(logs)
	if err != nil {
		return "", fmt.Errorf("读取日志内容失败: %v", err)
	}

	return buf.String(), nil
}

// ptr 辅助函数：创建字符串指针
func ptr(s string) *string {
	return &s
}

// int64Ptr 辅助函数：创建int64指针
func int64Ptr(i int64) *int64 {
	return &i
}

// parseInt64 解析字符串为int64
func parseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// cpuSetCount 统计 cpuset 字符串中的CPU个数，例如 "0-3,5" -> 5
func cpuSetCount(set string) int {
	set = strings.TrimSpace(set)
	if set == "" {
		return 0
	}
	total := 0
	parts := strings.Split(set, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.Contains(p, "-") {
			// 范围
			rangeParts := strings.SplitN(p, "-", 2)
			if len(rangeParts) == 2 {
				start, err1 := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
				end, err2 := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
				if err1 == nil && err2 == nil && end >= start {
					total += (end - start + 1)
				}
			}
		} else {
			// 单个
			if _, err := strconv.Atoi(p); err == nil {
				total += 1
			}
		}
	}
	return total
}

// getAssignedCPUs 计算容器分配（可用）的 CPU 核数
// 优先顺序：cpuset -> NanoCPUs -> CPUQuota/CPUPeriod -> CPUCount -> 0(未知)
func (d *DockerService) getAssignedCPUs(ctx context.Context, cli *client.Client, containerID string) float64 {
	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err != nil || inspect.HostConfig == nil {
		return 0
	}
	hc := inspect.HostConfig
	// 1) cpuset 限定（最严格、最准确）
	if hc.CpusetCpus != "" {
		c := cpuSetCount(hc.CpusetCpus)
		if c > 0 {
			return float64(c)
		}
	}
	// 2) NanoCPUs（以 1e9 为 1 核）
	if hc.NanoCPUs > 0 {
		return float64(hc.NanoCPUs) / 1e9
	}
	// 3) CPUQuota/CPUPeriod（cgroup v1 常见限制）
	if hc.CPUQuota > 0 && hc.CPUPeriod > 0 {
		return float64(hc.CPUQuota) / float64(hc.CPUPeriod)
	}
	// 4) CPUCount（Windows 或部分平台）
	if hc.CPUCount > 0 {
		return float64(hc.CPUCount)
	}
	return 0
}

// TestDockerConnection 测试Docker连接
func (d *DockerService) TestDockerConnection(ctx context.Context, node *computenode.ComputeNode) (bool, string) {
	if node.DockerAddress == nil || *node.DockerAddress == "" {
		return false, "Docker连接地址为空"
	}

	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return false, fmt.Sprintf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	// 尝试ping Docker服务
	_, err = cli.Ping(ctx)
	if err != nil {
		return false, fmt.Sprintf("Docker连接失败: %v", err)
	}

	return true, "连接成功"
}

// ContainerStats 容器统计信息
type ContainerStats struct {
	// CPUUsagePercent: 归一化到0-100%（相对于可用CPU核数），更贴近直觉
	CPUUsagePercent float64 `json:"cpuUsagePercent"` // CPU使用率百分比(0-100)
	// CPUUsagePercentRaw: 原始百分比，可超过100%，= 使用核数合计 * 100
	CPUUsagePercentRaw float64 `json:"cpuUsagePercentRaw,omitempty"` // 原始CPU百分比(可能>100%)
	MemoryUsage        int64   `json:"memoryUsage"`                  // 内存使用量（字节）
	MemoryLimit        int64   `json:"memoryLimit"`                  // 内存限制（字节）
	MemoryUsagePercent float64 `json:"memoryUsagePercent"`           // 内存使用率百分比
	Pids               uint64  `json:"pids"`                         // 进程数
	GPUMemorySizeGB    float64 `json:"gpuMemorySizeGB"`              // GPU显存大小(GB) - 当产品规格中显卡数量>0时返回
	GPUMemoryUsageRate float64 `json:"gpuMemoryUsageRate"`           // GPU显存使用率(%) - 当产品规格中显卡数量>0时返回
}

// 显存采集缓存（无TTL校验）
var (
	statsCache sync.Map // key: cacheKey(node, containerID) -> cachedStats
)

type cachedStats struct {
	at    time.Time
	stats *ContainerStats
}

func cacheKey(node *computenode.ComputeNode, containerID string) string {
	addr := ""
	if node != nil && node.DockerAddress != nil {
		addr = *node.DockerAddress
	}
	return addr + "|" + containerID
}

// getGPUMemoryInfo 获取GPU显存信息
// 返回值: (显存大小GB, 显存使用率%)
func (d *DockerService) getGPUMemoryInfo(ctx context.Context, cli *client.Client, containerID string) (float64, float64) {
	start := time.Now()

	// 优先尝试无单位输出
	gpuSizeGB, gpuRate, ok := d.tryParseNvidiaSmi(ctx, cli, containerID, []string{
		"nvidia-smi", "--query-gpu=memory.total,memory.used", "--format=csv,noheader,nounits",
	})
	if ok {
		return gpuSizeGB, gpuRate
	}
	// 次选：允许带单位输出
	gpuSizeGB, gpuRate, ok = d.tryParseNvidiaSmi(ctx, cli, containerID, []string{
		"nvidia-smi", "--query-gpu=memory.total,memory.used", "--format=csv,noheader",
	})
	if ok {
		return gpuSizeGB, gpuRate
	}
	// 再次尝试绝对路径
	gpuSizeGB, gpuRate, ok = d.tryParseNvidiaSmi(ctx, cli, containerID, []string{
		"/usr/bin/nvidia-smi", "--query-gpu=memory.total,memory.used", "--format=csv,noheader,nounits",
	})
	if ok {
		return gpuSizeGB, gpuRate
	}
	// 容器内不可用时，尝试宿主机 nvidia-smi 配合 NVIDIA_VISIBLE_DEVICES
	gpuSizeGB, gpuRate, ok = d.tryHostNvidiaSmiWithVisibleDevices(ctx, cli, containerID)
	if ok {
		return gpuSizeGB, gpuRate
	}

	// Fallback：如果容器设置了 CUDA_DEVICE_MEMORY_LIMIT 环境变量，至少返回总显存大小
	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err == nil {
		limitGB := int64(0)
		for _, e := range inspect.Config.Env {
			if strings.HasPrefix(e, "CUDA_DEVICE_MEMORY_LIMIT=") {
				val := strings.TrimPrefix(e, "CUDA_DEVICE_MEMORY_LIMIT=")
				val = strings.TrimSpace(strings.TrimSuffix(val, "g"))
				if v, err := strconv.ParseInt(val, 10, 64); err == nil && v > 0 {
					limitGB = v
					break
				}
			}
		}
		if limitGB > 0 {
			return float64(limitGB), 0.0
		}
	}

	global.GVA_LOG.Warn("无法获取GPU显存信息（容器可能未安装nvidia-smi或未分配GPU）",
		zap.String("containerID", containerID), zap.Duration("cost", time.Since(start)))
	return 0.0, 0.0
}

// tryParseNvidiaSmi 在容器内执行 nvidia-smi 并解析（兼容多GPU、多行与带单位）
func (d *DockerService) tryParseNvidiaSmi(ctx context.Context, cli *client.Client, containerID string, cmd []string) (float64, float64, bool) {
	execConfig := container.ExecOptions{Cmd: cmd, AttachStdout: true, AttachStderr: true}
	execResp, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return 0, 0, false
	}
	a, err := cli.ContainerExecAttach(ctx, execResp.ID, container.ExecStartOptions{})
	if err != nil {
		return 0, 0, false
	}
	defer a.Close()
	var outBuf, errBuf bytes.Buffer
	if _, err := stdcopy.StdCopy(&outBuf, &errBuf, a.Reader); err != nil {
		return 0, 0, false
	}
	out := strings.TrimSpace(outBuf.String())
	if out == "" {
		return 0, 0, false
	}
	lines := strings.Split(out, "\n")
	var totalMB, usedMB float64
	var parsed bool
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// 兼容：
		// 1) nounits: "48000, 12000"
		// 2) 带单位: "48000 MiB, 12000 MiB"
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])
		lm := parseMBAllowUnits(left)
		rm := parseMBAllowUnits(right)
		if lm > 0 {
			totalMB += lm
			usedMB += rm
			parsed = true
		}
	}
	if !parsed || totalMB <= 0 {
		return 0, 0, false
	}
	gpuSizeGB := totalMB / 1024.0
	usage := 0.0
	if totalMB > 0 {
		usage = (usedMB / totalMB) * 100.0
	}
	return gpuSizeGB, usage, true
}

// parseMBAllowUnits 解析形如 "48000", "48000 MB", "48000 MiB" → 返回MB数
func parseMBAllowUnits(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	// 去除常见单位
	ls := strings.ToLower(s)
	ls = strings.ReplaceAll(ls, "mib", "mb")
	ls = strings.ReplaceAll(ls, " mib", " mb")
	ls = strings.TrimSpace(ls)
	// 提取前缀数字
	var numPart string
	for i, r := range ls {
		if !(r == '+' || r == '-' || r == '.' || (r >= '0' && r <= '9')) {
			numPart = strings.TrimSpace(ls[:i])
			break
		}
	}
	if numPart == "" {
		numPart = ls
	}
	v, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return 0
	}
	// 基本按MB处理
	return v
}

// tryHostNvidiaSmiWithVisibleDevices 尝试在宿主机执行 nvidia-smi 作为兜底方案（仅当服务与Docker宿主同机时可用）
func (d *DockerService) tryHostNvidiaSmiWithVisibleDevices(ctx context.Context, cli *client.Client, containerID string) (float64, float64, bool) {
	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return 0, 0, false
	}
	visible := ""
	for _, e := range inspect.Config.Env {
		if strings.HasPrefix(e, "NVIDIA_VISIBLE_DEVICES=") {
			visible = strings.TrimSpace(strings.TrimPrefix(e, "NVIDIA_VISIBLE_DEVICES="))
			break
		}
	}
	if visible == "" || visible == "none" || visible == "void" {
		return 0, 0, false
	}
	// 宿主机 nvidia-smi
	bin, err := exec.LookPath("nvidia-smi")
	if err != nil {
		return 0, 0, false
	}
	args := []string{"--query-gpu=memory.total,memory.used", "--format=csv,noheader,nounits"}
	// 如果不是all，则限定设备
	if strings.ToLower(visible) != "all" {
		// 直接传入 -i 可接受 index 或 UUID
		args = append([]string{"-i", visible}, args...)
	}
	cmd := exec.CommandContext(ctx, bin, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		// 若限定失败，回退到不限定
		cmd2 := exec.CommandContext(ctx, bin, "--query-gpu=memory.total,memory.used", "--format=csv,noheader,nounits")
		var o2, e2 bytes.Buffer
		cmd2.Stdout = &o2
		cmd2.Stderr = &e2
		if err2 := cmd2.Run(); err2 != nil {
			return 0, 0, false
		}
		outBuf = o2
	}
	out := strings.TrimSpace(outBuf.String())
	if out == "" {
		return 0, 0, false
	}
	lines := strings.Split(out, "\n")
	var totalMB, usedMB float64
	var parsed bool
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}
		lm, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		rm, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if lm > 0 {
			totalMB += lm
			usedMB += rm
			parsed = true
		}
	}
	if !parsed || totalMB <= 0 {
		return 0, 0, false
	}
	sizeGB := totalMB / 1024.0
	rate := (usedMB / totalMB) * 100.0
	return sizeGB, rate, true
}

// getContainerStatsViaCLI 通过 docker stats --no-stream --format 获取统计信息
func (d *DockerService) getContainerStatsViaCLI(ctx context.Context, node *computenode.ComputeNode, containerID string) (*ContainerStats, error) {
	// 确保 docker 可用
	path, err := exec.LookPath("docker")
	if err != nil {
		return nil, fmt.Errorf("找不到docker命令: %v", err)
	}

	// 准备环境变量
	env := os.Environ()
	// 设置 DOCKER_HOST（例如 tcp://1.2.3.4:2376 或 unix:///var/run/docker.sock）
	if node.DockerAddress != nil && *node.DockerAddress != "" {
		env = append(env, fmt.Sprintf("DOCKER_HOST=%s", *node.DockerAddress))
	}
	// TLS 处理
	var tmpDir string
	cleanup := func() {}
	if node.UseTls != nil && *node.UseTls {
		// 创建临时目录并写入证书
		dir, err := os.MkdirTemp("", "docker-cert-*")
		if err != nil {
			return nil, fmt.Errorf("创建临时证书目录失败: %v", err)
		}
		tmpDir = dir
		cleanup = func() {
			os.RemoveAll(tmpDir)
		}
		// 写入证书
		if node.CaCert == nil || node.ClientCert == nil || node.ClientKey == nil {
			cleanup()
			return nil, fmt.Errorf("TLS证书配置不完整")
		}
		if err := os.WriteFile(filepath.Join(tmpDir, "ca.pem"), []byte(*node.CaCert), 0600); err != nil {
			cleanup()
			return nil, fmt.Errorf("写入ca.pem失败: %v", err)
		}
		if err := os.WriteFile(filepath.Join(tmpDir, "cert.pem"), []byte(*node.ClientCert), 0600); err != nil {
			cleanup()
			return nil, fmt.Errorf("写入cert.pem失败: %v", err)
		}
		if err := os.WriteFile(filepath.Join(tmpDir, "key.pem"), []byte(*node.ClientKey), 0600); err != nil {
			cleanup()
			return nil, fmt.Errorf("写入key.pem失败: %v", err)
		}
		env = append(env, "DOCKER_TLS_VERIFY=1")
		env = append(env, fmt.Sprintf("DOCKER_CERT_PATH=%s", tmpDir))
	}
	defer cleanup()

	// 使用 --format 输出JSON，便于解析
	// 注意：docker stats --no-stream --format '{{json .}}' <container>
	cmd := exec.CommandContext(ctx, path, "stats", "--no-stream", "--format", "{{json .}}", containerID)
	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("执行docker stats失败: %v, stderr=%s", err, stderr.String())
	}
	line := strings.TrimSpace(stdout.String())
	if line == "" {
		return nil, fmt.Errorf("docker stats 输出为空")
	}

	// 解析JSON行
	type statsLine struct {
		Container string `json:"Container"`
		Name      string `json:"Name"`
		CPUPerc   string `json:"CPUPerc"`
		MemUsage  string `json:"MemUsage"`
		MemPerc   string `json:"MemPerc"`
		NetIO     string `json:"NetIO"`
		BlockIO   string `json:"BlockIO"`
		PIDs      string `json:"PIDs"`
	}
	var sl statsLine
	if err := json.Unmarshal([]byte(line), &sl); err != nil {
		return nil, fmt.Errorf("解析docker stats JSON失败: %v", err)
	}

	// 转换字段
	cpuHostPercent := parsePercent(sl.CPUPerc) // 相对于整机（所有核）的百分比
	memUsedBytes, memLimitBytes := parseUsedTotal(sl.MemUsage)
	memPerc := parsePercent(sl.MemPerc)
	// 不再采集网络和块设备I/O
	pids := parseInt64(sl.PIDs)

	// docker CLI 的 CPUPerc 已经是“核百分比”（1核满载≈100，2核≈200）
	// 需要根据容器分配的 CPU 数进行归一化
	var assigned float64 = 0
	if node != nil {
		if cliTmp, err := d.CreateDockerClient(node); err == nil {
			defer cliTmp.Close()
			assigned = d.getAssignedCPUs(ctx, cliTmp, containerID)
		}
	}
	if assigned <= 0 {
		assigned = 1 // 兜底，防止除零
	}
	raw := cpuHostPercent
	norm := raw / assigned
	if norm < 0 {
		norm = 0
	}
	if norm > 100 {
		norm = 100
	}

	return &ContainerStats{
		CPUUsagePercent:    norm,
		CPUUsagePercentRaw: raw,
		MemoryUsage:        memUsedBytes,
		MemoryLimit:        memLimitBytes,
		MemoryUsagePercent: memPerc,
		Pids:               uint64(pids),
	}, nil
}

// parsePercent 将 "12.34%" 转为 12.34
func parsePercent(s string) float64 {
	s = strings.TrimSpace(strings.TrimSuffix(s, "%"))
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// parseUsedTotal 解析 "X / Y"，两端的带单位字符串转为字节数，若只有一个值，返回该值与0
func parseUsedTotal(s string) (int64, int64) {
	parts := strings.Split(s, "/")
	if len(parts) == 0 {
		return 0, 0
	}
	first := parseBytes(strings.TrimSpace(parts[0]))
	if len(parts) < 2 {
		return first, 0
	}
	second := parseBytes(strings.TrimSpace(parts[1]))
	return first, second
}

// parseBytes 将带单位字符串转为字节数，支持 B, KB/kB, MB/MiB, GB/GiB, TB/TiB
func parseBytes(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	// 拆出数值与单位
	// 例如 "824KiB", "1.23GB", "12.5 MB"
	var numPart, unitPart string
	for i, r := range s {
		if !(r == '+' || r == '-' || r == '.' || (r >= '0' && r <= '9')) {
			numPart = strings.TrimSpace(s[:i])
			unitPart = strings.TrimSpace(s[i:])
			break
		}
	}
	if numPart == "" {
		numPart = s
		unitPart = "B"
	}
	val, _ := strconv.ParseFloat(numPart, 64)
	unit := strings.ToUpper(unitPart)
	// 兼容空格和iB
	unit = strings.ReplaceAll(unit, "IB", "B")
	unit = strings.ReplaceAll(unit, "I", "")
	unit = strings.TrimSpace(unit)
	switch unit {
	case "B", "":
		return int64(val)
	case "KB", "KIB", "K":
		return int64(val * 1024)
	case "MB", "MIB", "M":
		return int64(val * 1024 * 1024)
	case "GB", "GIB", "G":
		return int64(val * 1024 * 1024 * 1024)
	case "TB", "TIB", "T":
		return int64(val * 1024 * 1024 * 1024 * 1024)
	default:
		// 也可能是 "kB" 小写k
		lu := strings.ToLower(unitPart)
		if strings.HasPrefix(lu, "kb") || lu == "k" {
			return int64(val * 1000)
		}
		if strings.HasPrefix(lu, "mb") || lu == "m" {
			return int64(val * 1000 * 1000)
		}
		if strings.HasPrefix(lu, "gb") || lu == "g" {
			return int64(val * 1000 * 1000 * 1000)
		}
		if strings.HasPrefix(lu, "tb") || lu == "t" {
			return int64(val * 1000 * 1000 * 1000 * 1000)
		}
	}
	return int64(val)
}

// GetContainerStats 获取容器统计信息
func (d *DockerService) GetContainerStats(ctx context.Context, node *computenode.ComputeNode, containerID string) (*ContainerStats, error) {

	// 优先走 docker stats --no-stream（与Docker CLI一致）
	if stats, err := d.getContainerStatsViaCLI(ctx, node, containerID); err == nil && stats != nil {
		// 追加GPU信息（通过SDK exec nvidia-smi）
		if node != nil {
			if cliTmp, err := d.CreateDockerClient(node); err == nil {
				defer cliTmp.Close()
				gm, gr := d.getGPUMemoryInfo(ctx, cliTmp, containerID)
				stats.GPUMemorySizeGB = gm
				stats.GPUMemoryUsageRate = gr
			}
		}
		statsCache.Store(cacheKey(node, containerID), cachedStats{at: time.Now(), stats: stats})
		return stats, nil
	}

	cli, err := d.CreateDockerClient(node)
	if err != nil {
		return nil, fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer cli.Close()

	statsCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	stream, err := cli.ContainerStats(statsCtx, containerID, true)
	if err != nil {
		return nil, fmt.Errorf("获取容器统计信息失败: %v", err)
	}
	defer stream.Body.Close()

	dec := json.NewDecoder(stream.Body)
	var prev, curr types.StatsJSON
	if err := dec.Decode(&prev); err != nil {
		return nil, fmt.Errorf("解析统计信息失败(首样本): %v", err)
	}
	var hasSecond bool
	if err := dec.Decode(&curr); err == nil {
		hasSecond = true
	} else {
		curr = prev
	}

	// CPU 计算（multi-core 原始百分比 raw 和按容器配额归一化的 norm）
	var rawCPUPercent float64
	numCPUs := curr.CPUStats.OnlineCPUs
	if numCPUs == 0 {
		numCPUs = uint32(len(curr.CPUStats.CPUUsage.PercpuUsage))
	}
	if hasSecond {
		cpuDelta := float64(curr.CPUStats.CPUUsage.TotalUsage - prev.CPUStats.CPUUsage.TotalUsage)
		var elapsedNs float64
		if !curr.Read.IsZero() && !prev.Read.IsZero() {
			elapsedNs = float64(curr.Read.Sub(prev.Read).Nanoseconds())
		}
		if elapsedNs > 0 && cpuDelta >= 0 {
			rawCPUPercent = (cpuDelta / elapsedNs) * 100.0
		} else {
			systemDelta := float64(curr.CPUStats.SystemUsage - prev.CPUStats.SystemUsage)
			if cpuDelta > 0 && systemDelta > 0 {
				if numCPUs == 0 {
					numCPUs = uint32(len(curr.CPUStats.CPUUsage.PercpuUsage))
				}
				rawCPUPercent = (cpuDelta / systemDelta) * float64(numCPUs) * 100.0
			}
		}
	} else if curr.CPUStats.CPUUsage.TotalUsage > 0 && curr.PreCPUStats.CPUUsage.TotalUsage > 0 {
		cpuDelta := float64(curr.CPUStats.CPUUsage.TotalUsage - curr.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta := float64(curr.CPUStats.SystemUsage - curr.PreCPUStats.SystemUsage)
		if systemDelta > 0 {
			if numCPUs == 0 {
				numCPUs = uint32(len(curr.CPUStats.CPUUsage.PercpuUsage))
			}
			rawCPUPercent = (cpuDelta / systemDelta) * float64(numCPUs) * 100.0
		}
	}

	// 根据容器配额计算 0-100 的归一化百分比
	assigned := d.getAssignedCPUs(ctx, cli, containerID)
	normCPUPercent := rawCPUPercent
	if assigned > 0 {
		normCPUPercent = rawCPUPercent / assigned
	}
	if normCPUPercent < 0 {
		normCPUPercent = 0
	}
	if normCPUPercent > 100 {
		normCPUPercent = 100
	}

	// 内存
	var memoryPercent float64
	memoryUsage := curr.MemoryStats.Usage
	memoryLimit := curr.MemoryStats.Limit
	if memoryLimit > 0 {
		memoryPercent = float64(memoryUsage) / float64(memoryLimit) * 100.0
	}

	// 不再采集网络与块设备I/O
	pids := curr.PidsStats.Current

	gm, gr := d.getGPUMemoryInfo(ctx, cli, containerID)

	res := &ContainerStats{
		CPUUsagePercent:    normCPUPercent,
		CPUUsagePercentRaw: rawCPUPercent,
		MemoryUsage:        int64(memoryUsage),
		MemoryLimit:        int64(memoryLimit),
		MemoryUsagePercent: memoryPercent,
		Pids:               pids,
		GPUMemorySizeGB:    gm,
		GPUMemoryUsageRate: gr,
	}
	statsCache.Store(cacheKey(node, containerID), cachedStats{at: time.Now(), stats: res})
	return res, nil
}
