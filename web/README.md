# 前端（Web）说明

基于 Vue 3 + Vite + Element Plus 的前端工程，配合 gin-vue-admin 后端使用。

## 环境要求
- Node.js 20+
- npm 或 pnpm

## 快速开始
```bash
# 安装依赖（任选其一）
npm install
# 或
pnpm install

# 启动开发服务（默认 http://127.0.0.1:8080）
npm run dev

# 生产构建
npm run build

# 本地预览打包产物
npm run preview
```

## 环境变量（.env.development 示例）
在项目根目录 web/ 下新建 .env.development（或根据需要新建 .env.production 等），并填写：

```env
# 前端调用后端的基础前缀，开发模式下将由 Vite 代理到后端
VITE_BASE_API=/api

# 后端地址与端口（供开发代理与日志展示使用）
VITE_BASE_PATH=http://127.0.0.1
VITE_SERVER_PORT=8888

# 前端开发服务端口
VITE_CLI_PORT=8080

# 静态文件/上传文件基础路径（与后端保持一致）
VITE_FILE_API=/uploads/file
```

说明：
- 开发模式下，请求以 VITE_BASE_API（默认 /api）为前缀的接口将被 Vite 代理到 `${VITE_BASE_PATH}:${VITE_SERVER_PORT}/`，配置位于 vite.config.js 的 server.proxy。
- 生产部署时，可将前端静态资源由 Web 服务器（如 Nginx）托管，并将 /api 反向代理到后端服务。

## 可用脚本
- npm run dev：启动开发服务器（vite）
- npm run build：生产构建
- npm run preview：本地预览构建产物
- npm run serve：同 dev（等价启动参数）
- npm run limit-build：在低内存环境下的构建辅助脚本

## 目录结构（简要）
```text
web
 ├── Dockerfile
 ├── index.html
 ├── package.json
 ├── vite.config.js
 ├── src
 │   ├── api                   # 后端接口封装
 │   ├── assets                # 静态资源与图标
 │   ├── components            # 公共组件
 │   ├── core                  # GVA 配置/启动横幅等
 │   ├── directive             # 指令（如 v-auth）
 │   ├── hooks                 # 通用 hooks
 │   ├── main.js               # 入口文件
 │   ├── pinia                 # 状态管理
 │   ├── router                # 路由
 │   ├── style                 # 全局样式
 │   ├── utils                 # 工具库（request、format 等）
 │   └── view                  # 页面视图
 └── .env.*                    # 环境变量文件（可选）
```

## 开发代理说明
- 源码中统一通过 `import.meta.env.VITE_BASE_API` 作为接口前缀（默认 `/api`）。
- Vite 在开发模式将 `/api` 代理到 `${VITE_BASE_PATH}:${VITE_SERVER_PORT}/`。
- 示例：前端调用 `/api/instance/getContainerStats`，实际转发到 `http://127.0.0.1:8888/instance/getContainerStats`。

## 常见问题
- 若前端无法访问后端，请检查：
  - 后端是否已在 8888 端口启动；
  - .env.development 的 VITE_BASE_PATH、VITE_SERVER_PORT 是否与后端一致；
  - 浏览器控制台与网络面板中的请求地址是否带有 `/api` 前缀；
  - vite.config.js 中代理配置是否生效。