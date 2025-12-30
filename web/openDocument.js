// 打开文档脚本
import child_process from 'child_process'

var url = 'https://github.com/hequan2017/docker-gpu-manage'
var cmd = ''
switch (process.platform) {
  case 'win32':
    cmd = 'start'
    child_process.exec(cmd + ' ' + url)
    break

  case 'darwin':
    cmd = 'open'
    child_process.exec(cmd + ' ' + url)
    break
}
