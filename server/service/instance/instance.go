package instance

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/imageregistry"
	instanceModel "github.com/flipped-aurora/gin-vue-admin/server/model/instance"
	instanceReq "github.com/flipped-aurora/gin-vue-admin/server/model/instance/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	containerConfig := dockerService.BuildContainerConfig(&image, &spec, &node, containerName)

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
		"container_name":   containerName,
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
	global.GVA_DB.Table("image_registry").Where("deleted_at IS NULL AND is_on_shelf = ?", true).
		Select("name as label, id as value, support_memory_split as supportMemorySplit").Scan(&imageId)
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
		Select("id as value, name, gpu_model, gpu_count, memory_capacity, cpu_cores, memory_gb, system_disk_gb, data_disk_gb, price_per_hour, support_memory_split as supportMemorySplit").
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
	// 系统盘
	if systemDiskGb != nil {
		disk := int64(0)
		switch v := systemDiskGb.(type) {
		case int64:
			disk = v
		case int:
			disk = int64(v)
		case float64:
			disk = int64(v)
		}
		if disk > 0 {
			label += " | 系统盘" + strconv.FormatInt(disk, 10) + "G"
		}
	}
	// 数据盘
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
	ID                  uint    `json:"id"`
	Name                string  `json:"name"`
	Region              string  `json:"region"`
	GpuName             string  `json:"gpuName"`
	GpuCount            int64   `json:"gpuCount"`     // 节点总GPU数量
	AvailableGpu        int64   `json:"availableGpu"` // 可用GPU数量
	Cpu                 string  `json:"cpu"`
	AvailableCpu        int64   `json:"availableCpu"` // 可用CPU核心数
	Memory              string  `json:"memory"`
	AvailableMemory     int64   `json:"availableMemory"` // 可用内存(GB)
	MemoryCapacity      int64   `json:"memoryCapacity"`  // 显存容量(GB)
	SystemDisk          string  `json:"systemDisk"`
	AvailableSystemDisk int64   `json:"availableSystemDisk"` // 可用系统盘(GB)
	DataDisk            string  `json:"dataDisk"`
	AvailableDataDisk   int64   `json:"availableDataDisk"` // 可用数据盘(GB)
	PublicIp            string  `json:"publicIp"`
	PricePerHour        float64 `json:"pricePerHour"`
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
		NodeId             int64   `json:"nodeId"`
		GpuUsed            int64   `json:"gpuUsed"`
		CpuUsed            int64   `json:"cpuUsed"`
		MemUsed            int64   `json:"memUsed"`
		MemoryCapacityUsed int64   `json:"memoryCapacityUsed"` // 已使用的显存容量（累加值）
		CardMemoryUsage    []int64 `json:"cardMemoryUsage"`    // 每张卡已使用的显存容量
		SystemDiskUsed     int64   `json:"systemDiskUsed"`
		DataDiskUsed       int64   `json:"dataDiskUsed"`
	}
	usedResources := make(map[int64]UsedResource)

	// 查询每个节点已创建的实例所占用的资源
	var instances []instanceModel.Instance
	global.GVA_DB.Where("deleted_at IS NULL").Find(&instances)

	// 先获取所有节点信息，用于计算每张卡的显存容量
	nodeInfoMap := make(map[int64]*computenode.ComputeNode)
	for i := range allNodes {
		nodeInfoMap[int64(allNodes[i].ID)] = &allNodes[i]
	}

	for _, inst := range instances {
		if inst.NodeId == nil || inst.SpecId == nil {
			continue
		}
		// 获取该实例使用的规格
		var instSpec product.ProductSpec
		if err := global.GVA_DB.Where("id = ?", *inst.SpecId).First(&instSpec).Error; err != nil {
			continue
		}
		// 获取该实例使用的镜像，检查是否支持显存切分
		var image imageregistry.ImageRegistry
		supportMemorySplit := false
		if inst.ImageId != nil {
			if err := global.GVA_DB.Where("id = ?", *inst.ImageId).First(&image).Error; err == nil {
				if image.SupportMemorySplit != nil {
					supportMemorySplit = *image.SupportMemorySplit
				}
			}
		}
		nodeId := *inst.NodeId
		used := usedResources[nodeId]
		used.NodeId = nodeId

		// 记录分配前的状态
		beforeCardUsage := make([]int64, len(used.CardMemoryUsage))
		copy(beforeCardUsage, used.CardMemoryUsage)

		// 初始化卡显存使用数组
		if used.CardMemoryUsage == nil {
			used.CardMemoryUsage = make([]int64, 0)
		}

		// 获取节点信息
		nodeInfo, hasNodeInfo := nodeInfoMap[nodeId]

		if instSpec.GpuCount != nil {
			used.GpuUsed += *instSpec.GpuCount
		}
		if instSpec.CpuCores != nil {
			used.CpuUsed += *instSpec.CpuCores
		}
		if instSpec.MemoryGb != nil {
			used.MemUsed += *instSpec.MemoryGb
		}

		// 显存容量计算：按卡分配
		if instSpec.MemoryCapacity != nil && instSpec.GpuCount != nil && *instSpec.GpuCount > 0 {
			memoryNeeded := *instSpec.MemoryCapacity
			gpuCount := *instSpec.GpuCount

			// 计算每张卡需要的显存
			memoryPerCard := memoryNeeded / gpuCount

			// 获取节点每张卡的显存容量（MemoryCapacity 存储的是单卡容量）
			perCardCapacity := int64(0)
			if hasNodeInfo && nodeInfo.MemoryCapacity != nil {
				perCardCapacity = *nodeInfo.MemoryCapacity
			}

			// 确保卡显存使用数组有足够的长度
			if hasNodeInfo && nodeInfo.GpuCount != nil {
				totalCards := int(*nodeInfo.GpuCount)
				for len(used.CardMemoryUsage) < totalCards {
					used.CardMemoryUsage = append(used.CardMemoryUsage, 0)
				}
			}

			// 分配显存到卡上
			if supportMemorySplit && perCardCapacity > 0 {
				// 支持显存切分：可以分配到任意卡上
				// 优先分配到已有使用但未满的卡上，如果都满了，再分配到新卡上
				for i := 0; i < int(gpuCount); i++ {
					// 找到可以分配的卡（有剩余空间的卡）
					found := false
					// 优先查找已有使用但未满的卡
					for j := range used.CardMemoryUsage {
						remaining := perCardCapacity - used.CardMemoryUsage[j]
						if remaining >= memoryPerCard {
							used.CardMemoryUsage[j] += memoryPerCard
							used.MemoryCapacityUsed += memoryPerCard
							found = true
							break
						}
					}
					// 如果没找到可用卡，分配到新卡上（完全未使用的卡）
					if !found && len(used.CardMemoryUsage) < int(*nodeInfo.GpuCount) {
						used.CardMemoryUsage = append(used.CardMemoryUsage, memoryPerCard)
						used.MemoryCapacityUsed += memoryPerCard
					} else if !found {
						// 没有可用卡，累加到总使用量（这种情况不应该发生，但为了安全）
						used.MemoryCapacityUsed += memoryPerCard
						global.GVA_LOG.Warn("无法分配到卡，累加到总使用量",
							zap.Int64("显存", memoryPerCard))
					}
				}

				// 重新计算总使用显存，确保与卡使用情况一致
				totalCardUsage := int64(0)
				for _, usage := range used.CardMemoryUsage {
					totalCardUsage += usage
				}
				used.MemoryCapacityUsed = totalCardUsage
			} else {
				// 不支持显存切分：每张卡必须完全分配给一个实例，不能部分使用
				// 如果每张卡需要的显存等于每张卡的容量，可以分配
				if perCardCapacity > 0 && memoryPerCard == perCardCapacity {
					for i := 0; i < int(gpuCount); i++ {
						// 找到完全未使用的卡（使用量为0）
						found := false
						for j := range used.CardMemoryUsage {
							if used.CardMemoryUsage[j] == 0 {
								used.CardMemoryUsage[j] = perCardCapacity
								used.MemoryCapacityUsed += perCardCapacity
								found = true
								break
							}
						}
						if !found && len(used.CardMemoryUsage) < int(*nodeInfo.GpuCount) {
							used.CardMemoryUsage = append(used.CardMemoryUsage, perCardCapacity)
							used.MemoryCapacityUsed += perCardCapacity
						}
					}
				} else if perCardCapacity > 0 && memoryPerCard < perCardCapacity {
					// 如果每张卡需要的显存小于每张卡的容量，不支持显存切分时无法分配
					// 但为了统计，累加到总使用量（实际无法分配）
					used.MemoryCapacityUsed += memoryNeeded
					global.GVA_LOG.Warn("不支持切分且显存不匹配，无法分配",
						zap.Int64("每张卡需要", memoryPerCard),
						zap.Int64("每张卡容量", perCardCapacity))
				} else {
					// 如果每张卡需要的显存大于每张卡的容量，无法分配，但累加到总使用量
					used.MemoryCapacityUsed += memoryNeeded
					global.GVA_LOG.Warn("显存需求超过卡容量，无法分配",
						zap.Int64("每张卡需要", memoryPerCard),
						zap.Int64("每张卡容量", perCardCapacity))
				}

				// 重新计算总使用显存，确保与卡使用情况一致
				totalCardUsage := int64(0)
				for _, usage := range used.CardMemoryUsage {
					totalCardUsage += usage
				}
				used.MemoryCapacityUsed = totalCardUsage
			}
		} else if instSpec.MemoryCapacity != nil {
			// 如果没有GPU数量，直接累加
			used.MemoryCapacityUsed += *instSpec.MemoryCapacity
		}

		if instSpec.SystemDiskGb != nil {
			used.SystemDiskUsed += *instSpec.SystemDiskGb
		}
		if instSpec.DataDiskGb != nil {
			used.DataDiskUsed += *instSpec.DataDiskGb
		}
		usedResources[nodeId] = used
	}

	// 4. 筛选满足要求的节点
	nodes = make([]AvailableNode, 0)

	// 判断是否需要GPU
	requiredGpu := int64(0)
	needGpu := false
	if spec.GpuCount != nil && *spec.GpuCount > 0 {
		requiredGpu = *spec.GpuCount
		needGpu = true
	}

	for _, node := range allNodes {
		// 如果需要GPU，检查显卡型号是否匹配
		if needGpu {
			if spec.GpuModel != nil && node.GpuName != nil {
				if *spec.GpuModel != *node.GpuName {
					continue
				}
			}
		}

		// 计算可用资源
		nodeId := int64(node.ID)
		used := usedResources[nodeId]

		// 节点总GPU数量（仅在需要GPU时计算）
		totalGpu := int64(0)
		availableGpu := int64(0)
		if needGpu {
			if node.GpuCount != nil {
				totalGpu = *node.GpuCount
			}
			availableGpu = totalGpu - used.GpuUsed
		}

		// 获取节点CPU
		totalCpu := int64(0)
		if node.Cpu != nil {
			totalCpu = *node.Cpu
		}
		availableCpu := totalCpu - used.CpuUsed

		// 获取节点内存
		totalMem := int64(0)
		if node.Memory != nil {
			totalMem = *node.Memory
		}
		availableMem := totalMem - used.MemUsed

		// 获取节点系统盘
		totalSystemDisk := int64(0)
		if node.SystemDisk != nil {
			totalSystemDisk = *node.SystemDisk
		}
		availableSystemDisk := totalSystemDisk - used.SystemDiskUsed

		// 获取节点数据盘
		totalDataDisk := int64(0)
		if node.DataDisk != nil {
			totalDataDisk = *node.DataDisk
		}
		availableDataDisk := totalDataDisk - used.DataDiskUsed

		// 检查是否满足规格要求
		requiredCpu := int64(0)
		if spec.CpuCores != nil {
			requiredCpu = *spec.CpuCores
		}
		requiredMem := int64(0)
		if spec.MemoryGb != nil {
			requiredMem = *spec.MemoryGb
		}
		requiredSystemDisk := int64(0)
		if spec.SystemDiskGb != nil {
			requiredSystemDisk = *spec.SystemDiskGb
		}
		requiredDataDisk := int64(0)
		if spec.DataDiskGb != nil {
			requiredDataDisk = *spec.DataDiskGb
		}
		requiredMemoryCapacity := int64(0)
		if spec.MemoryCapacity != nil {
			requiredMemoryCapacity = *spec.MemoryCapacity
		}

		// 获取节点显存容量（MemoryCapacity 存储的是单卡容量）
		perCardCapacity := int64(0)
		if node.MemoryCapacity != nil {
			perCardCapacity = *node.MemoryCapacity
		}

		// 计算总显存容量
		totalMemoryCapacity := int64(0)
		if node.GpuCount != nil && *node.GpuCount > 0 && perCardCapacity > 0 {
			totalMemoryCapacity = perCardCapacity * *node.GpuCount
		}

		// 检查显存容量是否满足要求（按卡分配方式）
		if requiredMemoryCapacity > 0 && needGpu && requiredGpu > 0 && perCardCapacity > 0 {
			// 计算每张卡需要的显存
			requiredMemoryPerCard := requiredMemoryCapacity / requiredGpu

			// 初始化卡显存使用数组（如果还没有）
			if used.CardMemoryUsage == nil {
				used.CardMemoryUsage = make([]int64, 0)
			}

			// 确保数组长度足够
			totalCards := int64(0)
			if node.GpuCount != nil {
				totalCards = *node.GpuCount
			}
			for len(used.CardMemoryUsage) < int(totalCards) {
				used.CardMemoryUsage = append(used.CardMemoryUsage, 0)
			}

			// 获取产品规格是否支持显存分割
			specSupportMemorySplit := false
			if spec.SupportMemorySplit != nil {
				specSupportMemorySplit = *spec.SupportMemorySplit
			}

			global.GVA_LOG.Info("显存匹配检查",
				zap.Int64("节点ID", int64(node.ID)),
				zap.Int64("总显存容量", totalMemoryCapacity),
				zap.Int64("总卡数", totalCards),
				zap.Int64("每张卡容量", perCardCapacity),
				zap.Int64("需要显存", requiredMemoryCapacity),
				zap.Int64("需要GPU数", requiredGpu),
				zap.Int64("每张卡需要显存", requiredMemoryPerCard),
				zap.Bool("规格支持显存分割", specSupportMemorySplit),
				zap.Any("卡使用情况", used.CardMemoryUsage))

			// 创建一个临时数组来模拟分配，检查是否可以满足需求
			tempCardUsage := make([]int64, len(used.CardMemoryUsage))
			copy(tempCardUsage, used.CardMemoryUsage)

			canAllocate := true
			if specSupportMemorySplit {
				// 支持显存分割：根据每个卡的剩余可用显存容量进行判断
				for i := int64(0); i < requiredGpu; i++ {
					found := false
					// 查找可以分配的卡（有足够剩余空间的卡）
					for j := range tempCardUsage {
						remaining := perCardCapacity - tempCardUsage[j]
						if remaining >= requiredMemoryPerCard {
							tempCardUsage[j] += requiredMemoryPerCard
							global.GVA_LOG.Info("找到可用卡（支持显存分割）",
								zap.Int64("节点ID", int64(node.ID)),
								zap.Int("卡索引", j),
								zap.Int64("卡剩余显存", remaining),
								zap.Int64("需要显存", requiredMemoryPerCard))
							found = true
							break
						}
					}
					// 如果没找到可用卡，检查是否有新卡可用
					if !found {
						if len(tempCardUsage) < int(totalCards) {
							// 有新卡可用
							tempCardUsage = append(tempCardUsage, requiredMemoryPerCard)
							global.GVA_LOG.Info("使用新卡（支持显存分割）",
								zap.Int64("节点ID", int64(node.ID)),
								zap.Int("新卡索引", len(tempCardUsage)-1),
								zap.Int64("分配显存", requiredMemoryPerCard))
							found = true
						}
					}
					if !found {
						global.GVA_LOG.Warn("无法找到可用卡（支持显存分割）",
							zap.Int64("节点ID", int64(node.ID)),
							zap.Int64("需要GPU索引", i),
							zap.Int64("需要显存", requiredMemoryPerCard),
							zap.Any("当前卡使用情况", tempCardUsage))
						canAllocate = false
						break
					}
				}
			} else {
				// 不支持显存分割：需要整个卡的计算（每张卡必须完全未使用）
				// 需要 requiredGpu 张完全未使用的卡
				availableUnusedCards := int64(0)
				for _, cardUsage := range tempCardUsage {
					if cardUsage == 0 {
						availableUnusedCards++
					}
				}
				// 如果还有未使用的卡槽，也加上
				if len(tempCardUsage) < int(totalCards) {
					availableUnusedCards += totalCards - int64(len(tempCardUsage))
				}

				global.GVA_LOG.Info("检查未使用卡（不支持显存分割）",
					zap.Int64("节点ID", int64(node.ID)),
					zap.Int64("可用未使用卡数", availableUnusedCards),
					zap.Int64("需要GPU数", requiredGpu))

				if availableUnusedCards < requiredGpu {
					global.GVA_LOG.Warn("未使用卡数不足（不支持显存分割）",
						zap.Int64("节点ID", int64(node.ID)),
						zap.Int64("可用未使用卡数", availableUnusedCards),
						zap.Int64("需要GPU数", requiredGpu))
					canAllocate = false
				} else {
					// 模拟分配：标记 requiredGpu 张卡为已使用
					allocated := int64(0)
					for j := range tempCardUsage {
						if tempCardUsage[j] == 0 && allocated < requiredGpu {
							tempCardUsage[j] = perCardCapacity // 标记为完全使用
							allocated++
						}
					}
					// 如果还有未使用的卡槽，也标记
					for allocated < requiredGpu && len(tempCardUsage) < int(totalCards) {
						tempCardUsage = append(tempCardUsage, perCardCapacity)
						allocated++
					}
				}
			}

			// 如果无法分配，跳过该节点
			if !canAllocate {
				continue
			}

			// 显存检查通过（按卡分配），根据卡的使用情况计算实际可用GPU数量
			if needGpu {
				if specSupportMemorySplit {
					// 支持显存分割：计算有多少张卡有足够的剩余显存
					availableCards := int64(0)
					for _, cardUsage := range used.CardMemoryUsage {
						remaining := perCardCapacity - cardUsage
						if remaining >= requiredMemoryPerCard {
							availableCards++
						}
					}
					// 如果还有未使用的卡槽，也加上
					if len(used.CardMemoryUsage) < int(totalCards) {
						availableCards += totalCards - int64(len(used.CardMemoryUsage))
					}
					availableGpu = availableCards
					global.GVA_LOG.Info("根据卡使用情况计算可用GPU（支持显存分割）",
						zap.Int64("节点ID", int64(node.ID)),
						zap.Int64("可用卡数", availableCards),
						zap.Int64("可用GPU", availableGpu))
				} else {
					// 不支持显存分割：计算有多少张完全未使用的卡
					availableCards := int64(0)
					for _, cardUsage := range used.CardMemoryUsage {
						if cardUsage == 0 {
							availableCards++
						}
					}
					// 如果还有未使用的卡槽，也加上
					if len(used.CardMemoryUsage) < int(totalCards) {
						availableCards += totalCards - int64(len(used.CardMemoryUsage))
					}
					availableGpu = availableCards
					global.GVA_LOG.Info("根据卡使用情况计算可用GPU（不支持显存分割）",
						zap.Int64("节点ID", int64(node.ID)),
						zap.Int64("可用未使用卡数", availableCards),
						zap.Int64("可用GPU", availableGpu))
				}
			}
		} else if requiredMemoryCapacity > 0 {
			// 如果没有GPU需求，使用简单累加方式
			availableMemoryCapacity := totalMemoryCapacity - used.MemoryCapacityUsed
			if availableMemoryCapacity < requiredMemoryCapacity {
				continue
			}
		}

		// 资源不足则跳过（如果需要GPU才检查GPU资源）
		if needGpu && availableGpu < requiredGpu {
			continue
		}
		if availableCpu < requiredCpu {
			continue
		}
		if availableMem < requiredMem {
			continue
		}
		if availableSystemDisk < requiredSystemDisk {
			continue
		}
		if availableDataDisk < requiredDataDisk {
			continue
		}

		// 计算单卡的可用显存容量（找到所有卡中剩余显存最大的那张卡的剩余显存）
		availableMemoryCapacityPerCard := int64(0)
		if node.MemoryCapacity != nil {
			perCardCapacity := *node.MemoryCapacity

			if node.GpuCount != nil && *node.GpuCount > 0 {
				// 如果显存是按卡分配的，找到单卡的最大可用显存
				if len(used.CardMemoryUsage) > 0 && perCardCapacity > 0 {
					// 计算每张卡的剩余显存，找到最大的
					maxAvailablePerCard := int64(0)
					for _, cardUsage := range used.CardMemoryUsage {
						if cardUsage < perCardCapacity {
							remaining := perCardCapacity - cardUsage
							if remaining > maxAvailablePerCard {
								maxAvailablePerCard = remaining
							}
						}
					}
					availableMemoryCapacityPerCard = maxAvailablePerCard
					// 如果还有未使用的卡槽，未使用的卡有完整的单卡容量
					totalCards := int64(*node.GpuCount)
					if len(used.CardMemoryUsage) < int(totalCards) {
						// 未使用的卡有完整的单卡容量，这肯定比已使用卡的剩余显存大
						availableMemoryCapacityPerCard = perCardCapacity
					}
				} else {
					// 如果没有按卡分配，单卡可用显存就是单卡总容量
					availableMemoryCapacityPerCard = perCardCapacity
				}
			} else {
				// 如果没有GPU，单卡可用显存就是单卡总容量
				availableMemoryCapacityPerCard = perCardCapacity
			}
		}

		// 构建可用节点信息
		availableNode := AvailableNode{
			ID:                  node.ID,
			AvailableGpu:        availableGpu,
			AvailableCpu:        availableCpu,
			AvailableMemory:     availableMem,
			AvailableSystemDisk: availableSystemDisk,
			AvailableDataDisk:   availableDataDisk,
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
			availableNode.Cpu = strconv.FormatInt(*node.Cpu, 10)
		}
		if node.Memory != nil {
			availableNode.Memory = strconv.FormatInt(*node.Memory, 10)
		}
		// 返回单卡的可用显存容量
		availableNode.MemoryCapacity = availableMemoryCapacityPerCard
		if node.SystemDisk != nil {
			availableNode.SystemDisk = strconv.FormatInt(*node.SystemDisk, 10)
		}
		if node.DataDisk != nil {
			availableNode.DataDisk = strconv.FormatInt(*node.DataDisk, 10)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("实例不存在或已被删除")
		}
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
