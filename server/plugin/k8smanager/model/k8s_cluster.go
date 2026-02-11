package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// K8sCluster K8s集群配置模型
type K8sCluster struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"type:varchar(100);not null;uniqueIndex;comment:集群名称" binding:"required"` // 集群名称
	KubeConfig  string `json:"kubeConfig" gorm:"type:longtext;comment:kubeconfig配置内容" binding:"required"`               // kubeconfig配置内容
	Endpoint    string `json:"endpoint" gorm:"type:varchar(500);comment:API Server地址"`                                // API Server地址
	Version     string `json:"version" gorm:"type:varchar(50);comment:K8s版本"`                                          // K8s版本
	Status      string `json:"status" gorm:"type:varchar(20);default:unknown;comment:集群状态"`                          // 集群状态: online, offline, unknown
	Description string `json:"description" gorm:"type:varchar(500);comment:集群描述"`                                    // 集群描述
	Region      string `json:"region" gorm:"type:varchar(100);comment:区域"`                                            // 区域
	Provider    string `json:"provider" gorm:"type:varchar(50);comment:云服务商"`                                        // 云服务商: aliyun, aws, tencent, native
	IsDefault   bool   `json:"isDefault" gorm:"type:tinyint(1);default:0;comment:是否默认集群"`                           // 是否默认集群
	NodeCount   int    `json:"nodeCount" gorm:"type:int;default:0;comment:节点数量"`                                     // 节点数量
}

// TableName 指定表名
func (K8sCluster) TableName() string {
	return "k8s_clusters"
}
