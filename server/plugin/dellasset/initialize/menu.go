package initialize

const (
	DellAssetMenuID = 6030 // 戴尔资产管理菜单ID
)

// InitMenu 初始化菜单（配合 dellasset_install.sql 使用）
// 说明：本方法返回菜单配置，实际菜单创建请执行 SQL 文件 dellasset_install.sql
func InitMenu() map[uint]interface{} {
	return map[uint]interface{}{
		DellAssetMenuID: map[string]interface{}{
			"name":      "dellAsset",
			"path":      "dellAsset",
			"flag":      "gva",
			"component": "plugin/dellasset/view/dellAsset.vue",
			"meta": map[string]interface{}{
				"title":       "戴尔资产管理",
				"icon":        "cpu",
				"iconStyle":   "",
				"keepAlive":   false,
				"defaultMenu": false,
			},
			"parent_id": 3, // 放在资源管理下(根据实际情况调整，资源管理的ID通常是3)
			"btns": []map[string]interface{}{
				{
					"name": "create",
					"desc": "新增",
				},
				{
					"name": "update",
					"desc": "编辑",
				},
				{
					"name": "delete",
					"desc": "删除",
				},
			},
		},
	}
}

// GetInstallSQLPath 获取安装SQL文件路径
func GetInstallSQLPath() string {
	return "plugin/dellasset/dellasset_install.sql"
}
