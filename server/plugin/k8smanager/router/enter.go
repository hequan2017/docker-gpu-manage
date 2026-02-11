package router

type RouterGroup struct {
	K8sClusterRouter
	K8sPodRouter
	K8sDeploymentRouter
	K8sServiceRouter
	K8sNamespaceRouter
	K8sEventRouter
	K8sAuditRouter
	K8sMetricsRouter
	K8sNodeRouter
}

var RouterGroupApp = new(RouterGroup)
