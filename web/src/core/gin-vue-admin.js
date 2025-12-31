import { register } from './global'
import { viteLogo } from './config'
import { initErrorHandler } from './error-handel'

// Vue 插件对象
const ginVueAdmin = {
  install(app) {
    // 注册图标和全局配置
    register(app)
    
    // 初始化错误处理
    initErrorHandler(app)
    
    // 显示启动信息
    if (import.meta.env) {
      viteLogo(import.meta.env)
    }
  }
}

export default ginVueAdmin

