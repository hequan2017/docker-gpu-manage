package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// K8s集群相关API
		{Path: "/k8s/cluster/create", Description: "创建K8s集群", ApiGroup: "K8s集群", Method: "POST"},
		{Path: "/k8s/cluster/delete", Description: "删除K8s集群", ApiGroup: "K8s集群", Method: "DELETE"},
		{Path: "/k8s/cluster/deleteByIds", Description: "批量删除K8s集群", ApiGroup: "K8s集群", Method: "DELETE"},
		{Path: "/k8s/cluster/update", Description: "更新K8s集群", ApiGroup: "K8s集群", Method: "PUT"},
		{Path: "/k8s/cluster/get", Description: "获取K8s集群详情", ApiGroup: "K8s集群", Method: "GET"},
		{Path: "/k8s/cluster/list", Description: "获取K8s集群列表", ApiGroup: "K8s集群", Method: "GET"},
		{Path: "/k8s/cluster/refresh", Description: "刷新K8s集群状态", ApiGroup: "K8s集群", Method: "POST"},
		{Path: "/k8s/cluster/all", Description: "获取所有K8s集群", ApiGroup: "K8s集群", Method: "GET"},

		// Pod相关API
		{Path: "/k8s/pod/list", Description: "获取Pod列表", ApiGroup: "K8s Pod", Method: "GET"},
		{Path: "/k8s/pod/get", Description: "获取Pod详情", ApiGroup: "K8s Pod", Method: "GET"},
		{Path: "/k8s/pod/delete", Description: "删除Pod", ApiGroup: "K8s Pod", Method: "DELETE"},
		{Path: "/k8s/pod/log", Description: "获取Pod日志", ApiGroup: "K8s Pod", Method: "POST"},
		{Path: "/k8s/pod/containers", Description: "获取Pod容器列表", ApiGroup: "K8s Pod", Method: "GET"},
		{Path: "/k8s/pod/events", Description: "获取Pod事件", ApiGroup: "K8s Pod", Method: "GET"},

		// Deployment相关API
		{Path: "/k8s/deployment/list", Description: "获取Deployment列表", ApiGroup: "K8s Deployment", Method: "GET"},
		{Path: "/k8s/deployment/get", Description: "获取Deployment详情", ApiGroup: "K8s Deployment", Method: "GET"},
		{Path: "/k8s/deployment/scale", Description: "扩缩容Deployment", ApiGroup: "K8s Deployment", Method: "POST"},
		{Path: "/k8s/deployment/restart", Description: "重启Deployment", ApiGroup: "K8s Deployment", Method: "POST"},
		{Path: "/k8s/deployment/delete", Description: "删除Deployment", ApiGroup: "K8s Deployment", Method: "DELETE"},
		{Path: "/k8s/deployment/pods", Description: "获取Deployment关联的Pods", ApiGroup: "K8s Deployment", Method: "GET"},

		// Service相关API
		{Path: "/k8s/service/list", Description: "获取Service列表", ApiGroup: "K8s Service", Method: "GET"},
		{Path: "/k8s/service/get", Description: "获取Service详情", ApiGroup: "K8s Service", Method: "GET"},
		{Path: "/k8s/service/delete", Description: "删除Service", ApiGroup: "K8s Service", Method: "DELETE"},
		{Path: "/k8s/service/endpoints", Description: "获取Service的Endpoints", ApiGroup: "K8s Service", Method: "GET"},

		// Namespace相关API
		{Path: "/k8s/namespace/list", Description: "获取Namespace列表", ApiGroup: "K8s Namespace", Method: "GET"},
		{Path: "/k8s/namespace/get", Description: "获取Namespace详情", ApiGroup: "K8s Namespace", Method: "GET"},
		{Path: "/k8s/namespace/create", Description: "创建Namespace", ApiGroup: "K8s Namespace", Method: "POST"},
		{Path: "/k8s/namespace/delete", Description: "删除Namespace", ApiGroup: "K8s Namespace", Method: "DELETE"},

		// Event相关API
		{Path: "/k8s/event/list", Description: "获取Event列表", ApiGroup: "K8s Event", Method: "POST"},
	}
	utils.RegisterApis(entities...)
}
