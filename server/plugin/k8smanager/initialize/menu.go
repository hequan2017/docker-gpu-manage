package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	// 清理重复的菜单
	cleanupDuplicateK8sMenus()

	// 注意：ParentId 需要根据实际的系统菜单结构进行调整
	// 通常可以设置为系统管理或其他父菜单的ID
	// 这里假设 ParentId = 0 表示顶级菜单，或者根据实际情况调整
	parentMenu := model.SysBaseMenu{
		ParentId:  0, // 可以调整为实际的父菜单ID
		Path:      "k8s",
		Name:      "k8s",
		Hidden:    false,
		Component: "view/routerHolder.vue",
		Sort:      8,
		Meta:      model.Meta{Title: "K8s管理", Icon: "cpu-line"},
	}

	// 子菜单
	childMenus := []model.SysBaseMenu{
		{
			Path:      "k8sCluster",
			Name:      "k8sCluster",
			Hidden:    false,
			Component: "plugin/k8smanager/view/cluster.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "集群管理", Icon: "server-line"},
		},
		{
			Path:      "k8sNode",
			Name:      "k8sNode",
			Hidden:    false,
			Component: "plugin/k8smanager/view/node.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "Node管理", Icon: "monitor"},
		},
		{
			Path:      "k8sPod",
			Name:      "k8sPod",
			Hidden:    false,
			Component: "plugin/k8smanager/view/pod.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "Pod管理", Icon: "apps-line"},
		},
		{
			Path:      "k8sDeployment",
			Name:      "k8sDeployment",
			Hidden:    false,
			Component: "plugin/k8smanager/view/deployment.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "Deployment管理", Icon: "stack-line"},
		},
		{
			Path:      "k8sService",
			Name:      "k8sService",
			Hidden:    false,
			Component: "plugin/k8smanager/view/service.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "Service管理", Icon: "links-line"},
		},
		{
			Path:      "k8sNamespace",
			Name:      "k8sNamespace",
			Hidden:    false,
			Component: "plugin/k8smanager/view/namespace.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "Namespace管理", Icon: "folder-line"},
		},
		{
			Path:      "k8sEvent",
			Name:      "k8sEvent",
			Hidden:    false,
			Component: "plugin/k8smanager/view/event.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "事件管理", Icon: "notification-line"},
		},
	}

	// 合并父菜单和子菜单
	menus := append([]model.SysBaseMenu{parentMenu}, childMenus...)

	// 使用工具函数注册菜单
	utils.RegisterMenus(menus...)

	// 确保 Node 菜单存在 (因为 RegisterMenus 可能会跳过已存在的父菜单)
	ensureNodeMenu(parentMenu.Name)
}

func ensureNodeMenu(parentName string) {
	var parentMenu model.SysBaseMenu
	if err := global.GVA_DB.Where("name = ?", parentName).First(&parentMenu).Error; err != nil {
		return
	}

	var count int64
	global.GVA_DB.Model(&model.SysBaseMenu{}).Where("name = ?", "k8sNode").Count(&count)
	if count == 0 {
		nodeMenu := model.SysBaseMenu{
			ParentId:  parentMenu.ID,
			Path:      "k8sNode",
			Name:      "k8sNode",
			Hidden:    false,
			Component: "plugin/k8smanager/view/node.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "Node管理", Icon: "monitor"},
		}
		global.GVA_DB.Create(&nodeMenu)
		
		// 添加菜单权限给管理员 (ID: 888)
		// 注意: 这里假设管理员ID为888, 实际情况可能不同
		global.GVA_DB.Exec("INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?, ?)", 888, nodeMenu.ID)
	}
}

func cleanupDuplicateK8sMenus() {
	var menus []model.SysBaseMenu
	// 查找所有名为 k8s 的父菜单
	if err := global.GVA_DB.Where("name = ? AND parent_id = ?", "k8s", 0).Order("id asc").Find(&menus).Error; err != nil {
		return
	}

	if len(menus) > 1 {
		// 保留第一个（ID最小的），删除其他的
		// 注意：这里的删除逻辑是物理删除，为了彻底解决重复问题
		for _, m := range menus[1:] {
			// 1. 删除关联的子菜单
			global.GVA_DB.Unscoped().Where("parent_id = ?", m.ID).Delete(&model.SysBaseMenu{})

			// 2. 删除菜单与角色的关联
			global.GVA_DB.Exec("DELETE FROM sys_authority_menus WHERE sys_base_menu_id = ?", m.ID)

			// 3. 删除父菜单本身
			global.GVA_DB.Unscoped().Delete(&m)
		}
	}
}
