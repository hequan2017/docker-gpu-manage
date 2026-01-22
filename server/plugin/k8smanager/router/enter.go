package router

type RouterGroup struct {
	K8sClusterRouter
	K8sPodRouter
	K8sDeploymentRouter
	K8sServiceRouter
	K8sNamespaceRouter
	K8sEventRouter
}

var RouterGroupApp = new(RouterGroup)
