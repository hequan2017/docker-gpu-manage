import service from '@/utils/request'

// ==================== K8s 集群管理 ====================

/**
 * 创建K8s集群
 * @param {Object} data 集群信息
 * @returns {Promise} 创建结果
 */
export const createK8sCluster = (data) => {
  return service({
    url: '/k8s/cluster/create',
    method: 'post',
    data: data
  })
}

/**
 * 删除K8s集群
 * @param {Object} data 集群信息
 * @returns {Promise} 删除结果
 */
export const deleteK8sCluster = (data) => {
  return service({
    url: '/k8s/cluster/delete',
    method: 'delete',
    data: data
  })
}

/**
 * 批量删除K8s集群
 * @param {Object} data 集群ID列表
 * @returns {Promise} 删除结果
 */
export const deleteK8sClusterByIds = (data) => {
  return service({
    url: '/k8s/cluster/deleteByIds',
    method: 'delete',
    data: data
  })
}

/**
 * 更新K8s集群
 * @param {Object} data 集群信息
 * @returns {Promise} 更新结果
 */
export const updateK8sCluster = (data) => {
  return service({
    url: '/k8s/cluster/update',
    method: 'put',
    data: data
  })
}

/**
 * 获取K8s集群详情
 * @param {Object} params 查询参数
 * @returns {Promise} 集群详情
 */
export const getK8sCluster = (params) => {
  return service({
    url: '/k8s/cluster/get',
    method: 'get',
    params: params
  })
}

/**
 * 获取K8s集群列表
 * @param {Object} params 查询参数
 * @returns {Promise} 集群列表
 */
export const getK8sClusterList = (params) => {
  return service({
    url: '/k8s/cluster/list',
    method: 'get',
    params: params
  })
}

/**
 * 刷新K8s集群状态
 * @param {Object} params 查询参数
 * @returns {Promise} 刷新结果
 */
export const refreshK8sClusterStatus = (params) => {
  return service({
    url: '/k8s/cluster/refresh',
    method: 'post',
    params: params
  })
}

/**
 * 获取所有K8s集群（用于下拉选择）
 * @returns {Promise} 集群列表
 */
export const getAllK8sClusters = () => {
  return service({
    url: '/k8s/cluster/all',
    method: 'get'
  })
}

// ==================== Pod 管理 ====================

/**
 * 获取Pod列表
 * @param {Object} params 查询参数
 * @returns {Promise} Pod列表
 */
export const getPodList = (params) => {
  return service({
    url: '/k8s/pod/list',
    method: 'get',
    params: params
  })
}

/**
 * 删除Pod
 * @param {Object} params 删除参数
 * @returns {Promise} 删除结果
 */
export const deletePod = (params) => {
  return service({
    url: '/k8s/pod/delete',
    method: 'delete',
    params: params
  })
}

/**
 * 获取Pod日志
 * @param {Object} data 查询参数
 * @returns {Promise} 日志内容
 */
export const getPodLog = (data) => {
  return service({
    url: '/k8s/pod/log',
    method: 'post',
    data: data
  })
}

/**
 * 获取Pod容器列表
 * @param {Object} params 查询参数
 * @returns {Promise} 容器列表
 */
export const getPodContainers = (params) => {
  return service({
    url: '/k8s/pod/containers',
    method: 'get',
    params: params
  })
}

/**
 * 获取Pod事件
 * @param {Object} params 查询参数
 * @returns {Promise} 事件列表
 */
export const getPodEvents = (params) => {
  return service({
    url: '/k8s/pod/events',
    method: 'get',
    params: params
  })
}

/**
 * AI诊断Pod
 * @param {Object} params 查询参数
 * @returns {Promise} 诊断结果
 */
export const diagnosePod = (params) => {
  return service({
    url: '/k8s/pod/diagnose',
    method: 'post',
    params: params
  })
}

// ==================== Deployment 管理 ====================

/**
 * 获取Deployment列表
 * @param {Object} params 查询参数
 * @returns {Promise} Deployment列表
 */
export const getDeploymentList = (params) => {
  return service({
    url: '/k8s/deployment/list',
    method: 'get',
    params: params
  })
}

/**
 * 获取Deployment详情
 * @param {Object} params 查询参数
 * @returns {Promise} Deployment详情
 */
export const getDeployment = (params) => {
  return service({
    url: '/k8s/deployment/get',
    method: 'get',
    params: params
  })
}

/**
 * 扩缩容Deployment
 * @param {Object} data 扩缩容参数
 * @returns {Promise} 扩缩容结果
 */
export const scaleDeployment = (data) => {
  return service({
    url: '/k8s/deployment/scale',
    method: 'post',
    data: data
  })
}

/**
 * 重启Deployment
 * @param {Object} data 重启参数
 * @returns {Promise} 重启结果
 */
export const restartDeployment = (data) => {
  return service({
    url: '/k8s/deployment/restart',
    method: 'post',
    data: data
  })
}

/**
 * 删除Deployment
 * @param {Object} data 删除参数
 * @returns {Promise} 删除结果
 */
export const deleteDeployment = (data) => {
  return service({
    url: '/k8s/deployment/delete',
    method: 'delete',
    data: data
  })
}

/**
 * 获取Deployment关联的Pods
 * @param {Object} params 查询参数
 * @returns {Promise} Pod列表
 */
export const getDeploymentPods = (params) => {
  return service({
    url: '/k8s/deployment/pods',
    method: 'get',
    params: params
  })
}
