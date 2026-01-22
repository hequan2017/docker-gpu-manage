package service

type ServiceGroup struct {
	K8sClusterService
	K8sPodService
	K8sDeploymentService
	K8sServiceService
	K8sNamespaceService
	K8sEventService
	K8sAuditService
	K8sPermissionService
	K8sMetricsService
}

var ServiceGroupApp = new(ServiceGroup)
