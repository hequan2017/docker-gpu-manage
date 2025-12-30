/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'

export const config = {
  appName: '天启算力平台',
  showViteLogo: true,
  keepAliveTabs: false,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(`> 欢迎使用天启算力平台`)
    console.log(`> 当前版本:v${packageInfo.version}`)
    console.log(
      `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
    )
    console.log(
      `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
    )
    console.log('\n')
  }
}

export default config
