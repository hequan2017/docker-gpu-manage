package service

type ServiceGroup struct {
	K8sClusterService
	K8sPodService
	K8sDeploymentService
	K8sServiceService
	K8sNamespaceService
	K8sEventService
}

var ServiceGroupApp = new(ServiceGroup)
