import service from '@/utils/request'

// 创建PCDN节点
export const createPcdnNode = (data) => {
  return service({
    url: '/pcdn/createPcdnNode',
    method: 'post',
    data
  })
}

// 删除PCDN节点
export const deletePcdnNode = (params) => {
  return service({
    url: '/pcdn/deletePcdnNode',
    method: 'delete',
    params
  })
}

// 批量删除PCDN节点
export const deletePcdnNodeByIds = (params) => {
  return service({
    url: '/pcdn/deletePcdnNodeByIds',
    method: 'delete',
    params
  })
}

// 更新PCDN节点
export const updatePcdnNode = (data) => {
  return service({
    url: '/pcdn/updatePcdnNode',
    method: 'put',
    data
  })
}

// 查询PCDN节点
export const findPcdnNode = (params) => {
  return service({
    url: '/pcdn/findPcdnNode',
    method: 'get',
    params
  })
}

// 分页获取PCDN节点列表
export const getPcdnNodeList = (params) => {
  return service({
    url: '/pcdn/getPcdnNodeList',
    method: 'get',
    params
  })
}
