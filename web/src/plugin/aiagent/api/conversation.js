import service from '@/utils/request'

/**
 * 创建会话
 * @param {Object} data 会话数据
 * @returns {Promise} 会话数据
 */
export const createConversation = (data) => {
  return service({
    url: '/conversation/createConversation',
    method: 'post',
    data
  })
}

/**
 * 删除会话
 * @param {Object} params 查询参数
 * @param {string} params.ID 会话ID
 * @returns {Promise}
 */
export const deleteConversation = (params) => {
  return service({
    url: '/conversation/deleteConversation',
    method: 'delete',
    params
  })
}

/**
 * 更新会话
 * @param {Object} data 会话数据
 * @returns {Promise}
 */
export const updateConversation = (data) => {
  return service({
    url: '/conversation/updateConversation',
    method: 'put',
    data
  })
}

/**
 * 根据ID获取会话
 * @param {Object} params 查询参数
 * @param {string} params.ID 会话ID
 * @returns {Promise} 会话数据
 */
export const findConversation = (params) => {
  return service({
    url: '/conversation/findConversation',
    method: 'get',
    params
  })
}

/**
 * 获取会话列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.title 标题搜索（可选）
 * @param {string} params.model 模型搜索（可选）
 * @param {boolean} params.isActive 是否激活（可选）
 * @returns {Promise} 会话列表
 */
export const getConversationList = (params) => {
  return service({
    url: '/conversation/getConversationList',
    method: 'get',
    params
  })
}

/**
 * 设置会话激活状态
 * @param {Object} params 查询参数
 * @param {string} params.ID 会话ID
 * @param {boolean} params.isActive 是否激活
 * @returns {Promise}
 */
export const setConversationActive = (params) => {
  return service({
    url: '/conversation/setActive',
    method: 'post',
    params
  })
}

/**
 * 获取激活的会话
 * @returns {Promise} 激活的会话数据
 */
export const getActiveConversation = () => {
  return service({
    url: '/conversation/getActive',
    method: 'get'
  })
}
