// 自动生成模板ComputeNode
package computenode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 算力节点 结构体  ComputeNode
type ComputeNode struct {
	global.GVA_MODEL
	Name           *string `json:"name" form:"name" gorm:"comment:节点名称;column:name;size:255;" binding:"required"`                              //名字
	Region         *string `json:"region" form:"region" gorm:"comment:节点区域;column:region;size:255;"`                                           //区域
	Cpu            *int64  `json:"cpu" form:"cpu" gorm:"comment:CPU信息;column:cpu;"`                                                            //CPU
	Memory         *int64  `json:"memory" form:"memory" gorm:"comment:内存信息;column:memory;"`                                                    //内存
	SystemDisk     *int64  `json:"systemDisk" form:"systemDisk" gorm:"comment:系统盘容量;column:system_disk;"`                                      //系统盘容量
	DataDisk       *int64  `json:"dataDisk" form:"dataDisk" gorm:"comment:数据盘容量;column:data_disk;"`                                            //数据盘容量
	PublicIp       *string `json:"publicIp" form:"publicIp" gorm:"comment:公网IP地址;column:public_ip;size:255;" binding:"required"`               //IP地址公网
	PrivateIp      *string `json:"privateIp" form:"privateIp" gorm:"comment:内网IP地址;column:private_ip;size:255;" binding:"required"`            //IP地址内网
	SshPort        *int64  `json:"sshPort" form:"sshPort" gorm:"default:22;comment:SSH端口;column:ssh_port;"`                                    //SSH端口
	Username       *string `json:"username" form:"username" gorm:"comment:登录用户名;column:username;size:255;"`                                    //用户名
	Password       *string `json:"password" form:"password" gorm:"comment:登录密码;column:password;size:255;"`                                     //密码
	GpuName        *string `json:"gpuName" form:"gpuName" gorm:"comment:显卡名称;column:gpu_name;size:255;"`                                       //显卡名称
	GpuCount       *int64  `json:"gpuCount" form:"gpuCount" gorm:"comment:显卡数量;column:gpu_count;"`                                             //显卡数量
	MemoryCapacity *int64  `json:"memoryCapacity" form:"memoryCapacity" gorm:"comment:显存容量;column:memory_capacity;"`                           //显存容量
	DockerAddress  *string `json:"dockerAddress" form:"dockerAddress" gorm:"comment:Docker连接地址;column:docker_address;size:500;"`               //Docker连接地址
	UseTls         *bool   `json:"useTls" form:"useTls" gorm:"default:true;comment:是否使用TLS;column:use_tls;"`                                   //使用TLS
	CaCert         *string `json:"caCert" form:"caCert" gorm:"comment:CA证书内容;column:ca_cert;type:text;"`                                       //CA证书
	ClientCert     *string `json:"clientCert" form:"clientCert" gorm:"comment:客户端证书内容;column:client_cert;type:text;"`                          //客户端证书
	ClientKey      *string `json:"clientKey" form:"clientKey" gorm:"comment:客户端私钥内容;column:client_key;type:text;"`                             //客户端私钥
	IsOnShelf      *bool   `json:"isOnShelf" form:"isOnShelf" gorm:"default:true;comment:是否上架;column:is_on_shelf;" binding:"required"`         //是否上架
	Remark         *string `json:"remark" form:"remark" gorm:"comment:备注信息;column:remark;size:1000;"`                                          //备注
	DockerStatus   *string `json:"dockerStatus" form:"dockerStatus" gorm:"comment:Docker连接状态;column:docker_status;size:50;default:'unknown';"` //Docker连接状态
}

// TableName 算力节点 ComputeNode自定义表名 compute_node
func (ComputeNode) TableName() string {
	return "compute_node"
}
