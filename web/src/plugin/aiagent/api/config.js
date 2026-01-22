import service from '@/utils/request'

/**
 * 创建AI配置
 * @param {Object} data 配置数据
 * @param {string} data.name 配置名称
 * @param {string} data.apiKey API Key
 * @param {string} data.baseURL API基础URL
 * @param {string} data.model 默认模型
 * @param {number} data.temperature 默认温度
 * @param {number} data.maxTokens 默认最大token数
 * @param {boolean} data.isActive 是否启用
 * @returns {Promise}
 */
export const createConfig = (data) => {
  return service({
    url: '/config/createConfig',
    method: 'post',
    data
  })
}

/**
 * 删除AI配置
 * @param {Object} params 查询参数
 * @param {string} params.ID 配置ID
 * @returns {Promise}
 */
export const deleteConfig = (params) => {
  return service({
    url: '/config/deleteConfig',
    method: 'delete',
    params
  })
}

/**
 * 更新AI配置
 * @param {Object} data 配置数据
 * @returns {Promise}
 */
export const updateConfig = (data) => {
  return service({
    url: '/config/updateConfig',
    method: 'put',
    data
  })
}

/**
 * 根据ID获取AI配置
 * @param {Object} params 查询参数
 * @param {string} params.ID 配置ID
 * @returns {Promise} 配置数据
 */
export const findConfig = (params) => {
  return service({
    url: '/config/findConfig',
    method: 'get',
    params
  })
}

/**
 * 获取AI配置列表
 * @returns {Promise} 配置列表
 */
export const getConfigList = () => {
  return service({
    url: '/config/getConfigList',
    method: 'get'
  })
}

/**
 * 设置AI配置激活状态
 * @param {Object} params 查询参数
 * @param {string} params.ID 配置ID
 * @returns {Promise}
 */
export const setConfigActive = (params) => {
  return service({
    url: '/config/setActive',
    method: 'post',
    params
  })
}

/**
 * 获取激活的AI配置
 * @returns {Promise} 激活的配置数据
 */
export const getActiveConfig = () => {
  return service({
    url: '/config/getActive',
    method: 'get'
  })
}
