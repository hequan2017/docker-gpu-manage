package v1

type ApiGroup struct {
	K8sClusterApi
	K8sPodApi
	K8sDeploymentApi
	K8sServiceApi
	K8sNamespaceApi
	K8sEventApi
	K8sAuditApi
	K8sMetricsApi
}

var ApiGroupApp = new(ApiGroup)
