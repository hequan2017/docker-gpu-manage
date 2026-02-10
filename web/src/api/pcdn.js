import service from '@/utils/request'

export const getPcdnNodeList = (params) => {
  return service({
    url: '/pcdn/node/list',
    method: 'get',
    params
  })
}

export const getPcdnPolicyList = (params) => {
  return service({
    url: '/pcdn/policy/list',
    method: 'get',
    params
  })
}

export const createPcdnPolicy = (data) => {
  return service({
    url: '/pcdn/policy/create',
    method: 'post',
    data
  })
}

export const updatePcdnPolicyWeight = (data) => {
  return service({
    url: '/pcdn/policy/weight',
    method: 'put',
    data
  })
}

export const publishPcdnPolicy = (data) => {
  return service({
    url: '/pcdn/policy/publish',
    method: 'post',
    data
  })
}

export const rollbackPcdnPolicy = (data) => {
  return service({
    url: '/pcdn/policy/rollback',
    method: 'post',
    data
  })
}

export const getPcdnDispatchTaskList = (params) => {
  return service({
    url: '/pcdn/dispatch/list',
    method: 'get',
    params
  })
}

export const retryPcdnDispatchTask = (data) => {
  return service({
    url: '/pcdn/dispatch/retry',
    method: 'post',
    data
  })
}

export const getPcdnDashboardMetrics = () => {
  return service({
    url: '/pcdn/dashboard/metrics',
    method: 'get'
  })
}
