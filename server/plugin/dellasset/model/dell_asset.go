package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DellAsset 戴尔服务器资产 结构体
type DellAsset struct {
	global.GVA_MODEL
	HostName       string  `json:"hostName" form:"hostName" gorm:"column:host_name;comment:主机名;type:varchar(100);not null"`                                        // 主机名
	ServiceTag     string  `json:"serviceTag" form:"serviceTag" gorm:"column:service_tag;comment:服务标签;type:varchar(20);uniqueIndex;not null"`                    // 服务标签(戴尔唯一标识)
	AssetNumber    string  `json:"assetNumber" form:"assetNumber" gorm:"column:asset_number;comment:资产编号;type:varchar(50)"`                                     // 资产编号
	Model          string  `json:"model" form:"model" gorm:"column:model;comment:型号;type:varchar(100)"`                                                        // 型号
	SerialNumber   string  `json:"serialNumber" form:"serialNumber" gorm:"column:serial_number;comment:序列号;type:varchar(50)"`                              // 序列号
	CPUModel       string  `json:"cpuModel" form:"cpuModel" gorm:"column:cpu_model;comment:CPU型号;type:varchar(100)"`                                      // CPU型号
	CPUCores       int     `json:"cpuCores" form:"cpuCores" gorm:"column:cpu_cores;comment:CPU核心数;default:0"`                                          // CPU核心数
	CPUThreads     int     `json:"cpuThreads" form:"cpuThreads" gorm:"column:cpu_threads;comment:CPU线程数;default:0"`                                      // CPU线程数
	MemoryCapacity int     `json:"memoryCapacity" form:"memoryCapacity" gorm:"column:memory_capacity;comment:内存容量(GB);default:0"`                          // 内存容量(GB)
	DiskInfo       string  `json:"diskInfo" form:"diskInfo" gorm:"column:disk_info;comment:磁盘信息;type:varchar(500)"`                                      // 磁盘信息
	NetworkInfo    string  `json:"networkInfo" form:"networkInfo" gorm:"column:network_info;comment:网卡信息;type:varchar(500)"`                                // 网卡信息
	IPAddress      string  `json:"ipAddress" form:"ipAddress" gorm:"column:ip_address;comment:IP地址;type:varchar(50)"`                                      // IP地址
	MACAddress     string  `json:"macAddress" form:"macAddress" gorm:"column:mac_address;comment:MAC地址;type:varchar(50)"`                                  // MAC地址
	Cabinet        string  `json:"cabinet" form:"cabinet" gorm:"column:cabinet;comment:机柜位置;type:varchar(50)"`                                          // 机柜位置
	RackPosition   string  `json:"rackPosition" form:"rackPosition" gorm:"column:rack_position;comment:机架位置;type:varchar(20)"`                             // 机架位置(U位)
	PowerStatus    string  `json:"powerStatus" form:"powerStatus" gorm:"column:power_status;comment:电源状态;type:varchar(20);default:'offline'"`               // 电源状态: online/offline/maintenance
	PurchaseDate   string  `json:"purchaseDate" form:"purchaseDate" gorm:"column:purchase_date;comment:购买日期;type:varchar(20)"`                             // 购买日期
	WarrantyExpiry string  `json:"warrantyExpiry" form:"warrantyExpiry" gorm:"column:warranty_expiry;comment:保修到期日;type:varchar(20)"`                      // 保修到期日
	OS             string  `json:"os" form:"os" gorm:"column:os;comment:操作系统;type:varchar(100)"`                                                          // 操作系统
	Department     string  `json:"department" form:"department" gorm:"column:department;comment:所属部门;type:varchar(100)"`                                 // 所属部门
	Manager        string  `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;type:varchar(50)"`                                            // 负责人
	Status         string  `json:"status" form:"status" gorm:"column:status;comment:状态;type:varchar(20);default:'offline'"`                             // 状态: online/offline/maintenance
	Remarks        string  `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;type:text"`                                                    // 备注
}

// TableName DellAsset 自定义表名
func (DellAsset) TableName() string {
	return "gva_dell_asset"
}
