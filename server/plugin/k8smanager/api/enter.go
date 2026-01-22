package v1

type ApiGroup struct {
	K8sClusterApi
	K8sPodApi
	K8sDeploymentApi
	K8sServiceApi
	K8sNamespaceApi
	K8sEventApi
}

var ApiGroupApp = new(ApiGroup)
