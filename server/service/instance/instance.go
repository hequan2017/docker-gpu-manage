package instance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	instanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/instance/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"go.uber.org/zap"
)

// InstanceWithUser 包含用户信息的实例结构体
type InstanceWithUser struct {
	instanceModel.Instance
	UserName string `json:"userName"` // 创建用户名
}

type InstanceService struct{}

var dockerService = &DockerService{}

// CreateInstance 创建实例管理记录并创建Docker容器
func (instanceService *InstanceService) CreateInstance(ctx context.Context, inst *instanceModel.Instance) (err error) {
	// 1. 获取镜像信息
	var image imageregistry.ImageRegistry
	if inst.ImageId == nil {
		return fmt.Errorf("镜像ID不能为空")
	}
	if err = global.GVA_DB.Where("id = ?", *inst.ImageId).First(&image).Error; err != nil {
		return fmt.Errorf("获取镜像信息失败: %v", err)
	}

	// 2. 获取产品规格信息
	var spec product.ProductSpec
	if inst.SpecId == nil {
		return fmt.Errorf("产品规格ID不能为空")
	}
	if err = global.GVA_DB.Where("id = ?", *inst.SpecId).First(&spec).Error; err != nil {
		return fmt.Errorf("获取产品规格信息失败: %v", err)
	}

	// 3. 获取算力节点信息
	var node computenode.ComputeNode
	if inst.NodeId == nil {
		return fmt.Errorf("算力节点ID不能为空")
	}
	if err = global.GVA_DB.Where("id = ?", *inst.NodeId).First(&node).Error; err != nil {
		return fmt.Errorf("获取算力节点信息失败: %v", err)
	}

	// 4. 先创建数据库记录获取ID
	initialStatus := "creating"
	inst.ContainerStatus = &initialStatus
	if err = global.GVA_DB.Create(inst).Error; err != nil {
		return fmt.Errorf("创建实例记录失败: %v", err)
	}

	// 5. 生成容器名称
	instanceName := inst.Name
	if instanceName == nil || *instanceName == "" {
		name := fmt.Sprintf("instance-%d", inst.ID)
		instanceName = &name
	}
	containerName := dockerService.GenerateInstanceName(*instanceName, inst.ID)

	// 6. 构建容器配置
	containerConfig := dockerService.BuildContainerConfig(&image, &spec, containerName)

	// 7. 创建Docker容器
	containerID, err := dockerService.CreateContainer(ctx, &node, containerConfig)
	if err != nil {
		// 创建容器失败，更新状态
		failedStatus := "failed"
		global.GVA_DB.Model(inst).Updates(map[string]interface{}{
			"container_status": failedStatus,
		})
		global.GVA_LOG.Error("创建Docker容器失败", zap.Error(err))
		return fmt.Errorf("创建Docker容器失败: %v", err)
	}

	// 8. 更新实例记录
	runningStatus := "running"
	err = global.GVA_DB.Model(inst).Updates(map[string]interface{}{
		"container_id":     containerID,
		"container_status": runningStatus,
	}).Error
	if err != nil {
		global.GVA_LOG.Error("更新实例记录失败", zap.Error(err))
	}

	return nil
}

// DeleteInstance 删除实例管理记录并删除Docker容器
func (instanceService *InstanceService) DeleteInstance(ctx context.Context, ID string, userID uint, isAdmin bool) (err error) {
	// 1. 获取实例信息
	var inst instanceModel.Instance
	if err = global.GVA_DB.Where("id = ?", ID).First(&inst).Error; err != nil {
		return fmt.Errorf("获取实例信息失败: %v", err)
	}

	// 权限检查：普通用户只能删除自己创建的实例
	if !isAdmin {
		userIDInt64 := int64(userID)
		if inst.UserId == nil || *inst.UserId != userIDInt64 {
			return fmt.Errorf("无权删除此实例")
		}
	}

	// 2. 如果有容器ID，先删除Docker容器
	if inst.ContainerId != nil && *inst.ContainerId != "" && inst.NodeId != nil {
		var node computenode.ComputeNode
		if err = global.GVA_DB.Where("id = ?", *inst.NodeId).First(&node).Error; err == nil {
			containerName := ""
			if inst.Name != nil {
				containerName = dockerService.GenerateInstanceName(*inst.Name, inst.ID)
			}
			// 删除容器（忽略错误，因为容器可能已经不存在）
			if delErr := dockerService.DeleteContainer(ctx, &node, *inst.ContainerId, containerName); delErr != nil {
				global.GVA_LOG.Warn("删除Docker容器失败", zap.Error(delErr))
			}
		}
	}

	// 3. 删除数据库记录
	err = global.GVA_DB.Delete(&instanceModel.Instance{}, "id = ?", ID).Error
	return err
}

// DeleteInstanceByIds 批量删除实例管理记录并删除Docker容器
func (instanceService *InstanceService) DeleteInstanceByIds(ctx context.Context, IDs []string, userID uint, isAdmin bool) (err error) {
	// 逐个删除以确保容器也被删除
	for _, id := range IDs {
		if delErr := instanceService.DeleteInstance(ctx, id, userID, isAdmin); delErr != nil {
			global.GVA_LOG.Warn("删除实例失败", zap.String("id", id), zap.Error(delErr))
			// 如果是权限错误，直接返回
			if delErr.Error() == "无权删除此实例" {
				return delErr
			}
		}
	}
	return nil
}

// UpdateInstance 更新实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) UpdateInstance(ctx context.Context, inst instanceModel.Instance) (err error) {
	err = global.GVA_DB.Model(&instanceModel.Instance{}).Where("id = ?", inst.ID).Updates(&inst).Error
	return err
}

// GetInstance 根据ID获取实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) GetInstance(ctx context.Context, ID string) (inst InstanceWithUser, err error) {
	var instance instanceModel.Instance
	err = global.GVA_DB.Where("id = ?", ID).First(&instance).Error
	if err != nil {
		return
	}

	inst.Instance = instance
	// 查询用户信息
	if instance.UserId != nil {
		var username string
		global.GVA_DB.Table("sys_users").Where("id = ?", *instance.UserId).Select("username").Scan(&username)
		inst.UserName = username
	}
	return
}

// GetInstanceInfoList 分页获取实例管理记录
// Author [yourname](https://github.com/yourname)
func (instanceService *InstanceService) GetInstanceInfoList(ctx context.Context, info instanceReq.InstanceSearch, userID uint, isAdmin bool) (list []InstanceWithUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db，使用 Left Join 关联用户表
	db := global.GVA_DB.Table("instance").
		Select("instance.*, sys_users.username as user_name").
		Joins("LEFT JOIN sys_users ON instance.user_id = sys_users.id")
	var instances []InstanceWithUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	// 权限控制：普通用户只能看到自己创建的实例
	if !isAdmin {
		userIDInt64 := int64(userID)
		db = db.Where("instance.user_id = ?", userIDInt64)
	}

	if info.ImageId != nil {
		db = db.Where("instance.image_id = ?", *info.ImageId)
	}
	if info.SpecId != nil {
		db = db.Where("instance.spec_id = ?", *info.SpecId)
	}
	if info.UserId != nil {
		db = db.Where("instance.user_id = ?", *info.UserId)
	}
	if info.NodeId != nil {
		db = db.Where("instance.node_id = ?", *info.NodeId)
	}
	if info.ContainerId != nil && *info.ContainerId != "" {
		db = db.Where("instance.container_id LIKE ?", "%"+*info.ContainerId+"%")
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("instance.name LIKE ?", "%"+*info.Name+"%")
	}
	if info.ContainerStatus != nil && *info.ContainerStatus != "" {
		db = db.Where("instance.container_status = ?", *info.ContainerStatus)
	}

	// 使用单独的查询来统计总数，避免 JOIN 影响计数
	countDB := global.GVA_DB.Model(&instanceModel.Instance{})
	// 应用相同的过滤条件
	if len(info.CreatedAtRange) == 2 {
		countDB = countDB.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if !isAdmin {
		userIDInt64 := int64(userID)
		countDB = countDB.Where("user_id = ?", userIDInt64)
	}
	if info.ImageId != nil {
		countDB = countDB.Where("image_id = ?", *info.ImageId)
	}
	if info.SpecId != nil {
		countDB = countDB.Where("spec_id = ?", *info.SpecId)
	}
	if info.UserId != nil {
		countDB = countDB.Where("user_id = ?", *info.UserId)
	}
	if info.NodeId != nil {
		countDB = countDB.Where("node_id = ?", *info.NodeId)
	}
	if info.ContainerId != nil && *info.ContainerId != "" {
		countDB = countDB.Where("container_id LIKE ?", "%"+*info.ContainerId+"%")
	}
	if info.Name != nil && *info.Name != "" {
		countDB = countDB.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.ContainerStatus != nil && *info.ContainerStatus != "" {
		countDB = countDB.Where("container_status = ?", *info.ContainerStatus)
	}
	err = countDB.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&instances).Error
	return instances, total, err
}
func (instanceService *InstanceService) GetInstanceDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	// 只返回已上架的镜像
	imageId := make([]map[string]any, 0)
	global.GVA_DB.Table("image_registry").Where("deleted_at IS NULL AND is_on_shelf = ?", true).Select("name as label,id as value").Scan(&imageId)
	res["imageId"] = imageId

	// 节点列表（初始为空，需要根据产品规格动态获取）
	nodeId := make([]map[string]any, 0)
	res["nodeId"] = nodeId

	// 所有节点列表（用于搜索过滤）
	allNodes := make([]map[string]any, 0)
	global.GVA_DB.Table("compute_node").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&allNodes)
	res["allNodes"] = allNodes

	// 只返回已上架的产品规格，并包含详细信息用于前端显示
	specId := make([]map[string]any, 0)
	global.GVA_DB.Table("product_spec").Where("deleted_at IS NULL AND is_on_shelf = ?", true).
		Select("id as value, name, gpu_model, gpu_count, cpu_cores, memory_gb, system_disk_gb, data_disk_gb, price_per_hour").
		Scan(&specId)
	// 为每个规格生成显示标签
	for i := range specId {
		item := specId[i]
		name := item["name"]
		gpuModel := item["gpu_model"]
		gpuCount := item["gpu_count"]
		cpuCores := item["cpu_cores"]
		memoryGb := item["memory_gb"]
		systemDiskGb := item["system_disk_gb"]
		dataDiskGb := item["data_disk_gb"]
		// 格式: 名称 | GPU型号 x 数量 | CPU核心 | 内存GB | 系统盘GB | 数据盘GB
		label := formatSpecLabel(name, gpuModel, gpuCount, cpuCores, memoryGb, systemDiskGb, dataDiskGb)
		specId[i]["label"] = label
	}
	res["specId"] = specId

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("username as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}

// formatSpecLabel 格式化产品规格标签
func formatSpecLabel(name, gpuModel, gpuCount, cpuCores, memoryGb, systemDiskGb, dataDiskGb any) string {
	label := ""
	if name != nil {
		label = name.(string)
	}
	if gpuModel != nil {
		label += " | " + gpuModel.(string)
		if gpuCount != nil {
			count := int64(0)
			switch v := gpuCount.(type) {
			case int64:
				count = v
			case int:
				count = int64(v)
			case float64:
				count = int64(v)
			}
			if count > 0 {
				label += " x " + strconv.FormatInt(count, 10)
			}
		}
	}
	if cpuCores != nil {
		cores := int64(0)
		switch v := cpuCores.(type) {
		case int64:
			cores = v
		case int:
			cores = int64(v)
		case float64:
			cores = int64(v)
		}
		if cores > 0 {
			label += " | " + strconv.FormatInt(cores, 10) + "核"
		}
	}
	if memoryGb != nil {
		mem := int64(0)
		switch v := memoryGb.(type) {
		case int64:
			mem = v
		case int:
			mem = int64(v)
		case float64:
			mem = int64(v)
		}
		if mem > 0 {
			label += " | " + strconv.FormatInt(mem, 10) + "G内存"
		}
	}
	// 系统盘不显示
	_ = systemDiskGb
	if dataDiskGb != nil {
		disk := int64(0)
		switch v := dataDiskGb.(type) {
		case int64:
			disk = v
		case int:
			disk = int64(v)
		case float64:
			disk = int64(v)
		}
		if disk > 0 {
			label += " | 数据盘" + strconv.FormatInt(disk, 10) + "G"
		}
	}
	return label
}
func (instanceService *InstanceService) GetInstancePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AvailableNode 可用节点信息
type AvailableNode struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	Region            string  `json:"region"`
	GpuName           string  `json:"gpuName"`
	GpuCount          int64   `json:"gpuCount"`     // 节点总GPU数量
	AvailableGpu      int64   `json:"availableGpu"` // 可用GPU数量
	Cpu               string  `json:"cpu"`
	AvailableCpu      int64   `json:"availableCpu"` // 可用CPU核心数
	Memory            string  `json:"memory"`
	AvailableMemory   int64   `json:"availableMemory"` // 可用内存(GB)
	SystemDisk        string  `json:"systemDisk"`
	DataDisk          string  `json:"dataDisk"`
	AvailableDataDisk int64   `json:"availableDataDisk"` // 可用数据盘(GB)
	PublicIp          string  `json:"publicIp"`
	PricePerHour      float64 `json:"pricePerHour"`
}

// GetAvailableNodes 根据产品规格获取可用的算力节点
func (instanceService *InstanceService) GetAvailableNodes(ctx context.Context, specIdStr string) (nodes []AvailableNode, err error) {
	specId, err := strconv.ParseUint(specIdStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// 1. 获取产品规格信息
	var spec product.ProductSpec
	if err = global.GVA_DB.Where("id = ? AND deleted_at IS NULL", specId).First(&spec).Error; err != nil {
		return nil, err
	}

	// 2. 获取所有已上架的算力节点
	var allNodes []computenode.ComputeNode
	if err = global.GVA_DB.Where("deleted_at IS NULL AND is_on_shelf = ?", true).Find(&allNodes).Error; err != nil {
		return nil, err
	}

	// 3. 统计每个节点已被实例占用的资源
	type UsedResource struct {
		NodeId       int64 `json:"nodeId"`
		GpuUsed      int64 `json:"gpuUsed"`
		CpuUsed      int64 `json:"cpuUsed"`
		MemUsed      int64 `json:"memUsed"`
		DataDiskUsed int64 `json:"dataDiskUsed"`
	}
	usedResources := make(map[int64]UsedResource)

	// 查询每个节点已创建的实例所占用的资源
	var instances []instanceModel.Instance
	global.GVA_DB.Where("deleted_at IS NULL").Find(&instances)

	for _, inst := range instances {
		if inst.NodeId == nil || inst.SpecId == nil {
			continue
		}
		// 获取该实例使用的规格
		var instSpec product.ProductSpec
		if err := global.GVA_DB.Where("id = ?", *inst.SpecId).First(&instSpec).Error; err != nil {
			continue
		}
		nodeId := *inst.NodeId
		used := usedResources[nodeId]
		used.NodeId = nodeId
		if instSpec.GpuCount != nil {
			used.GpuUsed += *instSpec.GpuCount
		}
		if instSpec.CpuCores != nil {
			used.CpuUsed += *instSpec.CpuCores
		}
		if instSpec.MemoryGb != nil {
			used.MemUsed += *instSpec.MemoryGb
		}
		if instSpec.DataDiskGb != nil {
			used.DataDiskUsed += *instSpec.DataDiskGb
		}
		usedResources[nodeId] = used
	}

	// 4. 筛选满足要求的节点
	nodes = make([]AvailableNode, 0)
	for _, node := range allNodes {
		// 检查显卡型号是否匹配
		if spec.GpuModel != nil && node.GpuName != nil {
			if *spec.GpuModel != *node.GpuName {
				continue
			}
		}

		// 计算可用资源
		nodeId := int64(node.ID)
		used := usedResources[nodeId]

		// 节点总GPU数量
		totalGpu := int64(0)
		if node.GpuCount != nil {
			totalGpu = *node.GpuCount
		}
		availableGpu := totalGpu - used.GpuUsed

		// 解析节点CPU (假设格式为数字或"8核"这样的格式)
		totalCpu := parseResourceValue(node.Cpu)
		availableCpu := totalCpu - used.CpuUsed

		// 解析节点内存
		totalMem := parseResourceValue(node.Memory)
		availableMem := totalMem - used.MemUsed

		// 解析节点数据盘
		totalDataDisk := parseResourceValue(node.DataDisk)
		availableDataDisk := totalDataDisk - used.DataDiskUsed

		// 检查是否满足规格要求
		requiredGpu := int64(0)
		if spec.GpuCount != nil {
			requiredGpu = *spec.GpuCount
		}
		requiredCpu := int64(0)
		if spec.CpuCores != nil {
			requiredCpu = *spec.CpuCores
		}
		requiredMem := int64(0)
		if spec.MemoryGb != nil {
			requiredMem = *spec.MemoryGb
		}
		requiredDataDisk := int64(0)
		if spec.DataDiskGb != nil {
			requiredDataDisk = *spec.DataDiskGb
		}

		// 资源不足则跳过
		if availableGpu < requiredGpu || availableCpu < requiredCpu || availableMem < requiredMem || availableDataDisk < requiredDataDisk {
			continue
		}

		// 构建可用节点信息
		availableNode := AvailableNode{
			ID:                node.ID,
			AvailableGpu:      availableGpu,
			AvailableCpu:      availableCpu,
			AvailableMemory:   availableMem,
			AvailableDataDisk: availableDataDisk,
		}
		if node.Name != nil {
			availableNode.Name = *node.Name
		}
		if node.Region != nil {
			availableNode.Region = *node.Region
		}
		if node.GpuName != nil {
			availableNode.GpuName = *node.GpuName
		}
		if node.GpuCount != nil {
			availableNode.GpuCount = *node.GpuCount
		}
		if node.Cpu != nil {
			availableNode.Cpu = *node.Cpu
		}
		if node.Memory != nil {
			availableNode.Memory = *node.Memory
		}
		if node.SystemDisk != nil {
			availableNode.SystemDisk = *node.SystemDisk
		}
		if node.DataDisk != nil {
			availableNode.DataDisk = *node.DataDisk
		}
		if node.PublicIp != nil {
			availableNode.PublicIp = *node.PublicIp
		}
		if spec.PricePerHour != nil {
			availableNode.PricePerHour = *spec.PricePerHour
		}

		nodes = append(nodes, availableNode)
	}

	return nodes, nil
}

// parseResourceValue 解析资源值，如 "8核" -> 8, "16G" -> 16, "8" -> 8
func parseResourceValue(val *string) int64 {
	if val == nil {
		return 0
	}
	s := *val
	// 移除非数字字符
	numStr := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			numStr += string(c)
		} else if len(numStr) > 0 {
			break
		}
	}
	if numStr == "" {
		return 0
	}
	n, _ := strconv.ParseInt(numStr, 10, 64)
	return n
}

// StartContainer 启动容器
func (instanceService *InstanceService) StartContainer(ctx context.Context, ID string) error {
	inst, node, err := instanceService.getInstanceAndNode(ID)
	if err != nil {
		return err
	}
	if inst.ContainerId == nil || *inst.ContainerId == "" {
		return fmt.Errorf("容器ID为空")
	}

	err = dockerService.StartContainer(ctx, node, *inst.ContainerId)
	if err != nil {
		return err
	}

	// 更新状态
	status := "running"
	return global.GVA_DB.Model(&inst).Update("container_status", status).Error
}

// StopContainer 停止容器
func (instanceService *InstanceService) StopContainer(ctx context.Context, ID string) error {
	inst, node, err := instanceService.getInstanceAndNode(ID)
	if err != nil {
		return err
	}
	if inst.ContainerId == nil || *inst.ContainerId == "" {
		return fmt.Errorf("容器ID为空")
	}

	err = dockerService.StopContainer(ctx, node, *inst.ContainerId)
	if err != nil {
		return err
	}

	// 更新状态
	status := "exited"
	return global.GVA_DB.Model(&inst).Update("container_status", status).Error
}

// RestartContainer 重启容器
func (instanceService *InstanceService) RestartContainer(ctx context.Context, ID string) error {
	inst, node, err := instanceService.getInstanceAndNode(ID)
	if err != nil {
		return err
	}
	if inst.ContainerId == nil || *inst.ContainerId == "" {
		return fmt.Errorf("容器ID为空")
	}

	err = dockerService.RestartContainer(ctx, node, *inst.ContainerId)
	if err != nil {
		return err
	}

	// 更新状态
	status := "running"
	return global.GVA_DB.Model(&inst).Update("container_status", status).Error
}

// GetContainerLogs 获取容器日志
func (instanceService *InstanceService) GetContainerLogs(ctx context.Context, ID string, tail string) (string, error) {
	inst, node, err := instanceService.getInstanceAndNode(ID)
	if err != nil {
		return "", err
	}
	if inst.ContainerId == nil || *inst.ContainerId == "" {
		return "", fmt.Errorf("容器ID为空")
	}

	return dockerService.GetContainerLogs(ctx, node, *inst.ContainerId, tail)
}

// getInstanceAndNode 获取实例和节点信息
func (instanceService *InstanceService) getInstanceAndNode(ID string) (*instanceModel.Instance, *computenode.ComputeNode, error) {
	var inst instanceModel.Instance
	if err := global.GVA_DB.Where("id = ?", ID).First(&inst).Error; err != nil {
		return nil, nil, fmt.Errorf("获取实例信息失败: %v", err)
	}

	if inst.NodeId == nil {
		return nil, nil, fmt.Errorf("实例未关联节点")
	}

	var node computenode.ComputeNode
	if err := global.GVA_DB.Where("id = ?", *inst.NodeId).First(&node).Error; err != nil {
		return nil, nil, fmt.Errorf("获取节点信息失败: %v", err)
	}

	return &inst, &node, nil
}
