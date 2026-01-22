package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager/utils"
	"gorm.io/gorm"
)

// K8sClusterService 集群服务结构体
type K8sClusterService struct{}

// CreateK8sCluster 创建K8s集群
func (s *K8sClusterService) CreateK8sCluster(ctx context.Context, cluster *model.K8sCluster) error {
	// 检查集群名称是否已存在
	var count int64
	err := global.GVA_DB.Model(&model.K8sCluster{}).Where("name = ?", cluster.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("集群名称已存在")
	}

	// 如果设置为默认集群，需要将其他集群的默认状态取消
	if cluster.IsDefault {
		err := global.GVA_DB.Model(&model.K8sCluster{}).Where("is_default = ?", true).Update("is_default", false).Error
		if err != nil {
			return err
		}
	}

	// 加密 kubeconfig
	encryptedConfig, err := utils.Encrypt(cluster.KubeConfig)
	if err != nil {
		return fmt.Errorf("加密kubeconfig失败: %w", err)
	}
	cluster.KubeConfig = encryptedConfig

	// 创建集群记录
	err = global.GVA_DB.Create(cluster).Error
	if err != nil {
		return err
	}

	// 验证集群连接并获取版本信息
	client, err := createClient(cluster)
	if err != nil {
		// 连接失败，更新状态为offline
		global.GVA_DB.Model(cluster).Update("status", "offline")
		return fmt.Errorf("集群连接验证失败: %w", err)
	}

	// 获取服务器版本
	version, err := client.Clientset.Discovery().ServerVersion()
	if err != nil {
		global.GVA_DB.Model(cluster).Update("status", "unknown")
		return fmt.Errorf("获取集群版本失败: %w", err)
	}

	// 获取节点数量
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx, nil)
	if err == nil {
		cluster.NodeCount = len(nodes.Items)
	}

	// 更新集群信息
	updates := map[string]interface{}{
		"version": version.GitVersion,
		"status":  "online",
		"endpoint": client.Config.Host,
	}
	if cluster.Endpoint == "" {
		updates["endpoint"] = client.Config.Host
	}

	err = global.GVA_DB.Model(cluster).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteK8sCluster 删除K8s集群
func (s *K8sClusterService) DeleteK8sCluster(ctx context.Context, id uint) error {
	// 查询集群
	var cluster model.K8sCluster
	err := global.GVA_DB.First(&cluster, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrClusterNotFound
		}
		return err
	}

	// 删除客户端缓存
	RemoveClient(cluster.Name)

	// 删除集群记录
	return global.GVA_DB.Delete(&cluster).Error
}

// DeleteK8sClusterByIds 批量删除K8s集群
func (s *K8sClusterService) DeleteK8sClusterByIds(ctx context.Context, ids []uint) error {
	// 查询集群列表
	var clusters []model.K8sCluster
	err := global.GVA_DB.Find(&clusters, ids).Error
	if err != nil {
		return err
	}

	// 删除客户端缓存
	for _, cluster := range clusters {
		RemoveClient(cluster.Name)
	}

	// 删除集群记录
	return global.GVA_DB.Delete(&clusters).Error
}

// UpdateK8sCluster 更新K8s集群
func (s *K8sClusterService) UpdateK8sCluster(ctx context.Context, cluster *model.K8sCluster) error {
	// 检查名称是否与其他集群冲突
	var count int64
	err := global.GVA_DB.Model(&model.K8sCluster{}).
		Where("name = ? AND id != ?", cluster.Name, cluster.ID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("集群名称已被其他集群使用")
	}

	// 如果设置为默认集群，需要将其他集群的默认状态取消
	if cluster.IsDefault {
		err := global.GVA_DB.Model(&model.K8sCluster{}).
			Where("is_default = ? AND id != ?", true, cluster.ID).
			Update("is_default", false).Error
		if err != nil {
			return err
		}
	}

	// 获取旧的集群信息
	var oldCluster model.K8sCluster
	err = global.GVA_DB.First(&oldCluster, cluster.ID).Error
	if err != nil {
		return err
	}

	// 如果kubeconfig发生变化，需要刷新客户端
	oldName := oldCluster.Name
	if oldCluster.KubeConfig != cluster.KubeConfig || oldCluster.Name != cluster.Name {
		RemoveClient(oldName)
	}

	// 加密 kubeconfig
	encryptedConfig, err := utils.Encrypt(cluster.KubeConfig)
	if err != nil {
		return fmt.Errorf("加密kubeconfig失败: %w", err)
	}

	// 更新集群信息
	err = global.GVA_DB.Model(cluster).Updates(map[string]interface{}{
		"name":        cluster.Name,
		"kube_config": encryptedConfig,
		"endpoint":    cluster.Endpoint,
		"description": cluster.Description,
		"region":      cluster.Region,
		"provider":    cluster.Provider,
		"is_default":  cluster.IsDefault,
		"status":      "unknown", // 重置状态，需要重新检查
	}).Error

	if err != nil {
		return err
	}

	// 重新验证集群连接
	client, err := createClient(cluster)
	if err != nil {
		return err
	}

	// 获取服务器版本
	version, err := client.Clientset.Discovery().ServerVersion()
	if err != nil {
		return err
	}

	// 获取节点数量
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx, nil)
	nodeCount := 0
	if err == nil {
		nodeCount = len(nodes.Items)
	}

	// 更新集群状态
	return global.GVA_DB.Model(cluster).Updates(map[string]interface{}{
		"version":    version.GitVersion,
		"status":     "online",
		"endpoint":   client.Config.Host,
		"node_count": nodeCount,
	}).Error
}

// GetK8sCluster 根据ID获取K8s集群
func (s *K8sClusterService) GetK8sCluster(ctx context.Context, id uint) (*model.K8sCluster, error) {
	var cluster model.K8sCluster
	err := global.GVA_DB.First(&cluster, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClusterNotFound
		}
		return nil, err
	}

	// 解密 kubeconfig（仅用于返回给前端）
	decryptedConfig, err := utils.Decrypt(cluster.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("解密kubeconfig失败: %w", err)
	}
	cluster.KubeConfig = decryptedConfig

	return &cluster, nil
}

// GetK8sClusterByName 根据名称获取K8s集群
func (s *K8sClusterService) GetK8sClusterByName(ctx context.Context, name string) (*model.K8sCluster, error) {
	var cluster model.K8sCluster
	err := global.GVA_DB.Where("name = ?", name).First(&cluster).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClusterNotFound
		}
		return nil, err
	}

	// 解密 kubeconfig
	decryptedConfig, err := utils.Decrypt(cluster.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("解密kubeconfig失败: %w", err)
	}
	cluster.KubeConfig = decryptedConfig

	return &cluster, nil
}

// GetK8sClusterInfoList 获取K8s集群列表
func (s *K8sClusterService) GetK8sClusterInfoList(ctx context.Context, info *request.K8sClusterSearch) ([]model.K8sCluster, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 构建查询条件
	db := global.GVA_DB.Model(&model.K8sCluster{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.Provider != "" {
		db = db.Where("provider = ?", info.Provider)
	}
	if info.Region != "" {
		db = db.Where("region LIKE ?", "%"+info.Region+"%")
	}

	// 获取总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取列表
	var clusters []model.K8sCluster
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&clusters).Error
	if err != nil {
		return nil, 0, err
	}

	return clusters, total, nil
}

// CheckClusterHealth 检查指定集群的健康状态
func (s *K8sClusterService) CheckClusterHealth(ctx context.Context, clusterName string) error {
	client, err := GetClusterClient(clusterName)
	if err != nil {
		return err
	}

	return CheckClusterHealth(client)
}

// RefreshClusterStatus 刷新集群状态
func (s *K8sClusterService) RefreshClusterStatus(ctx context.Context, clusterName string) error {
	var cluster model.K8sCluster
	err := global.GVA_DB.Where("name = ?", clusterName).First(&cluster).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrClusterNotFound
		}
		return err
	}

	// 移除旧客户端
	RemoveClient(clusterName)

	// 创建新客户端
	client, err := createClient(&cluster)
	if err != nil {
		// 连接失败
		return global.GVA_DB.Model(&cluster).Update("status", "offline").Error
	}

	// 检查健康状态
	err = CheckClusterHealth(client)
	if err != nil {
		return global.GVA_DB.Model(&cluster).Update("status", "offline").Error
	}

	// 获取服务器版本
	version, err := client.Clientset.Discovery().ServerVersion()
	if err != nil {
		return global.GVA_DB.Model(&cluster).Updates(map[string]interface{}{
			"status":  "unknown",
			"version": "",
		}).Error
	}

	// 获取节点数量
	nodes, err := client.Clientset.CoreV1().Nodes().List(ctx, nil)
	nodeCount := 0
	if err == nil {
		nodeCount = len(nodes.Items)
	}

	// 更新为在线状态
	return global.GVA_DB.Model(&cluster).Updates(map[string]interface{}{
		"status":     "online",
		"version":    version.GitVersion,
		"node_count": nodeCount,
	}).Error
}

// GetAllClusters 获取所有集群（用于下拉选择等）
func (s *K8sClusterService) GetAllClusters(ctx context.Context) ([]model.K8sCluster, error) {
	var clusters []model.K8sCluster
	err := global.GVA_DB.Select("id, name, status, provider").Find(&clusters).Error
	return clusters, err
}
