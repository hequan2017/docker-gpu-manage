import service from '@/utils/request'

/**
 * 获取消息列表
 * @param {Object} params 查询参数
 * @param {number} params.conversationID 会话ID
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.role 角色过滤（可选）
 * @returns {Promise} 消息列表
 */
export const getMessageList = (params) => {
  return service({
    url: '/message/getMessageList',
    method: 'get',
    params
  })
}

/**
 * 删除消息
 * @param {Object} params 查询参数
 * @param {string} params.ID 消息ID
 * @returns {Promise}
 */
export const deleteMessage = (params) => {
  return service({
    url: '/message/deleteMessage',
    method: 'delete',
    params
  })
}
