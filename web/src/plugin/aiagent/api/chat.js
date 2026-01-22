import service from '@/utils/request'

/**
 * 发送消息并获取AI回复
 * @param {Object} data 对话请求
 * @param {number} data.conversationID 会话ID（可选）
 * @param {string} data.message 用户消息内容
 * @param {string} data.model 指定模型（可选）
 * @param {number} data.temperature 温度参数（可选）
 * @param {number} data.maxTokens 最大token数（可选）
 * @param {boolean} data.stream 是否流式输出（默认false）
 * @returns {Promise} 对话响应
 */
export const sendMessage = (data) => {
  return service({
    url: '/chat/sendMessage',
    method: 'post',
    data
  })
}
