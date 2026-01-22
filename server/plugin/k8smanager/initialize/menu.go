package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
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
}
