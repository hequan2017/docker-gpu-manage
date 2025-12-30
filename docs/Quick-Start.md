# 🚀 快速开始

本指南帮助你快速部署和运行天启算力管理平台。

## 环境要求

### 后端环境
- Go 1.23+
- MySQL 5.7+ / PostgreSQL / SQLite / MSSQL / Oracle
- Redis（可选，用于缓存和会话管理）
- Docker（用于管理 GPU 容器）

### 前端环境
- Node.js 20+
- npm 或 pnpm

## 方式一：本地开发部署

### 1. 克隆项目

```bash
git clone https://github.com/hequan2017/docker-gpu-manage
cd docker-gpu-manage
mv server/config.yaml.bak server/config.yaml
```

### 2. 启动后端服务

```bash
cd server

# 安装依赖
go mod download

# 启动服务
go run main.go
```

后端服务默认运行在 `http://localhost:8888`

### 3. 启动前端服务

```bash
cd web

# 安装依赖
npm install
# 或使用 pnpm
pnpm install

# 启动开发服务器
npm run dev
```

前端服务默认运行在 `http://localhost:8080`

### 4. 初始化数据库

1. 访问 `http://localhost:8080`
2. 系统会自动检测数据库是否已初始化
3. 如果未初始化，会跳转到数据库初始化页面
4. 填写数据库连接信息：
   - 数据库类型（MySQL/PostgreSQL/SQLite/MSSQL/Oracle）
   - 数据库地址和端口
   - 数据库名称（如果不存在会自动创建）
   - 用户名和密码
5. 点击「初始化」按钮
6. 初始化完成后，会自动创建默认管理员账号

### 5. 登录系统

**默认管理员账号：**
- 用户名：`admin`
- 密码：`123456`

> ⚠️ **首次登录后请及时修改密码！**

## 方式二：Docker Compose 部署

### 1. 启动服务

```bash
cd deploy/docker-compose

# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

默认端口：
- MySQL：13306
- Redis：16379
- 后端服务：8888
- 前端服务：8080

### 2. 初始化数据库

1. 等待所有容器启动完成（约 1-2 分钟）
2. 访问 `http://localhost:8080`
3. 按照 Web 界面提示完成数据库初始化

### 3. 停止服务

```bash
docker-compose down

# 如需删除数据卷（会清空数据库）
docker-compose down -v
```

## 方式三：Kubernetes 部署

项目提供了 Kubernetes 部署配置文件，位于 `deploy/kubernetes/` 目录。

```bash
cd deploy/kubernetes

# 部署后端服务
kubectl apply -f server/

# 部署前端服务
kubectl apply -f web/
```

## 开发调试（可选）

### 后端热重载

```bash
go install github.com/silenceper/gowatch@latest
cd server && gowatch
```

### 常用地址

| 服务 | 地址 |
|------|------|
| Swagger 文档 | http://127.0.0.1:8888/swagger/index.html |
| SSE 端点 | http://127.0.0.1:8888/sse |
| Message 端点 | http://127.0.0.1:8888/message |
| 前端服务 | http://127.0.0.1:8080 |

### 添加 MCP 配置

```json
{
  "mcpServers": {
    "GVA Helper": {
      "url": "http://127.0.0.1:8888/sse"
    }
  }
}
```

## 配置显存切分（HAMi）

如需使用显存切分功能：

### 1. 部署 HAMi-core

在 GPU 算力节点上部署 HAMi-core：

```bash
git clone https://github.com/Project-HAMi/HAMi-core
cd HAMi-core
# 按照 HAMi 文档进行编译
```

### 2. 配置算力节点

在「算力节点管理」中，填写 HAMi-core 目录路径：

```
/root/HAMi-core/build
```

系统会自动：
- 挂载 HAMi 库目录到容器 `/libvgpu/build`
- 注入 `LD_PRELOAD`、`CUDA_DEVICE_MEMORY_LIMIT`、`CUDA_DEVICE_SM_LIMIT` 环境变量

## 下一步

- [📦 功能模块](./Features.md) - 了解详细功能
- [⚙️ 配置说明](./Configuration.md) - 深入了解配置项
- [❓ 常见问题](./FAQ.md) - 解决部署问题

