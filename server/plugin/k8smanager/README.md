# K8s 管理插件使用文档

## 📦 插件概述

这是一个基于 gin-vue-admin 框架的 Kubernetes 集群管理插件，提供了完整的 K8s 集群管理功能，包括集群连接、Pod 管理、Deployment 管理、Service 管理、Namespace 管理等核心功能。

## ✨ 核心功能

### 1. 集群管理
- ✅ 支持多集群管理
- ✅ KubeConfig 配置导入
- ✅ 集群状态监控（在线/离线）
- ✅ 自动获取集群版本和节点数
- ✅ 集群健康检查

### 2. Pod 管理
- ✅ Pod 列表查询
- ✅ Pod 详情查看
- ✅ Pod 删除
- ✅ Pod 日志查看
- ✅ Pod 容器列表
- ✅ Pod 事件查看

### 3. Deployment 管理
- ✅ Deployment 列表查询
- ✅ Deployment 详情查看
- ✅ Deployment 扩缩容
- ✅ Deployment 重启
- ✅ Deployment 删除
- ✅ 关联 Pods 查看

### 4. Service 管理
- ✅ Service 列表查询
- ✅ Service 详情查看
- ✅ Service 删除
- ✅ Endpoints 查看

### 5. Namespace 管理
- ✅ Namespace 列表查询
- ✅ Namespace 详情查看
- ✅ Namespace 创建
- ✅ Namespace 删除

### 6. 事件管理
- ✅ Event 列表查询
- ✅ 按资源类型过滤
- ✅ 按命名空间过滤

## 📁 项目结构

```
server/plugin/k8smanager/
├── api/                    # API 控制器层
│   └── v1/                # API v1 版本
│       ├── enter.go       # API 组入口
│       ├── k8s_cluster_api.go      # 集群 API
│       ├── k8s_pod_api.go          # Pod API
│       ├── k8s_deployment_api.go   # Deployment API
│       ├── k8s_service_api.go      # Service API
│       ├── k8s_namespace_api.go    # Namespace API
│       └── k8s_event_api.go        # Event API
├── initialize/            # 初始化模块
│   ├── api.go           # API 注册
│   ├── gorm.go          # 数据库初始化
│   ├── menu.go          # 菜单初始化
│   └── router.go        # 路由初始化
├── model/               # 数据模型层
│   ├── k8s_cluster.go  # 集群数据模型
│   └── request/        # 请求模型
│       └── k8s_cluster.go
├── router/              # 路由层
│   ├── enter.go        # 路由组入口
│   ├── k8s_cluster.go
│   ├── k8s_pod.go
│   ├── k8s_deployment.go
│   ├── k8s_service.go
│   ├── k8s_namespace.go
│   └── k8s_event.go
├── service/             # 服务层
│   ├── enter.go        # 服务组入口
│   ├── k8s_client.go   # K8s 客户端管理
│   ├── k8s_cluster_service.go
│   ├── k8s_pod_service.go
│   ├── k8s_deployment_service.go
│   ├── k8s_service.go
│   ├── k8s_namespace_service.go
│   └── k8s_event_service.go
├── plugin.go           # 插件入口
├── k8s_manager.sql     # SQL 脚本
└── README.md           # 使用文档

web/src/plugin/k8smanager/
├── api/               # API 接口封装
│   └── cluster.js    # 所有 K8s 相关 API
└── view/             # 页面组件
    └── cluster.vue   # 集群管理页面
```

## 🚀 安装步骤

### 1. 后端安装

#### 步骤 1：添加依赖

在 `server/go.mod` 中添加以下依赖：

```go
require (
    k8s.io/api v0.28.0
    k8s.io/apimachinery v0.28.0
    k8s.io/client-go v0.28.0
    k8s.io/metrics v0.28.0
)
```

然后运行：
```bash
cd server
go mod tidy
```

#### 步骤 2：注册插件

在 `server/main.go` 中注册插件：

```go
package main

import (
    "github.com/flipped-aurora/gin-vue-admin/server/core"
    "github.com/flipped-aurora/gin-vue-admin/server/initialize"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/k8smanager"
    // ... 其他导入
)

func main() {
    // ... 其他初始化代码

    // 注册 K8s 管理插件
    core.RegisterPlugin(k8smanager.Plugin)

    // ... 其他代码
}
```

#### 步骤 3：执行 SQL 脚本

执行 `server/plugin/k8smanager/k8s_manager.sql` 文件中的 SQL 语句，创建菜单和 API 权限。

**重要提示**：
- 根据实际情况调整菜单 ID（1000-1006）和 API ID（2001-2029）
- 根据实际情况调整 `authority_id`（示例中为 888）
- 根据需要调整 `parent_id`（决定菜单挂载位置）

#### 步骤 4：启动后端服务

```bash
cd server
go run main.go
```

### 2. 前端安装

#### 步骤 1：创建插件目录

确保前端文件已正确放置在 `web/src/plugin/k8smanager/` 目录下。

#### 步骤 2：安装依赖（如果需要）

```bash
cd web
npm install
# 或
yarn install
```

#### 步骤 3：启动前端服务

```bash
cd web
npm run dev
# 或
yarn dev
```

## 📝 使用说明

### 1. 添加 K8s 集群

1. 登录系统后，进入 "K8s管理" -> "集群管理"
2. 点击 "新增集群" 按钮
3. 填写集群信息：
   - **集群名称**：自定义名称，如 "生产集群"
   - **KubeConfig**：粘贴完整的 kubeconfig 文件内容
   - **API Server地址**：自动解析，可手动修改
   - **云服务商**：选择对应的云服务商
   - **区域**：填写集群所在区域
   - **描述**：可选，填写集群描述
   - **设为默认集群**：是否作为默认集群
4. 点击 "确定" 保存

系统会自动验证集群连接并获取版本信息。

### 2. 查看 Pod 列表

1. 进入 "K8s管理" -> "Pod管理"
2. 选择集群和命名空间
3. 可以按标签、字段选择器过滤
4. 查看 Pod 的状态、IP、节点等信息

### 3. 扩缩容 Deployment

1. 进入 "K8s管理" -> "Deployment管理"
2. 找到要操作的 Deployment
3. 点击 "扩缩容" 按钮
4. 设置目标副本数
5. 点击 "确定" 执行扩缩容

### 4. 查看 Pod 日志

1. 在 Pod 列表中点击 "详情"
2. 选择容器
3. 设置日志参数（行数、是否跟踪等）
4. 查看实时日志

## 🔧 API 接口文档

所有 API 都有完整的 Swagger 文档注解，启动后端服务后访问：

```
http://your-host:port/swagger/index.html
```

主要 API 路径：

- **集群管理**：`/k8s/cluster/*`
- **Pod 管理**：`/k8s/pod/*`
- **Deployment 管理**：`/k8s/deployment/*`
- **Service 管理**：`/k8s/service/*`
- **Namespace 管理**：`/k8s/namespace/*`
- **Event 管理**：`/k8s/event/*`

## 🔐 权限配置

插件中的所有 API 都已经注册到权限系统，需要为对应的角色（authority_id）分配权限：

1. **菜单权限**：通过 `sys_authority_menus` 表配置
2. **API 权限**：通过 `sys_authority_apis` 表配置

SQL 脚本中已为管理员角色（ID=888）添加了所有权限。

## 🛠️ 常见问题

### 1. 集群连接失败

**问题**：添加集群后显示 "离线" 状态

**解决方案**：
- 检查 kubeconfig 内容是否完整
- 检查网络连接是否正常
- 检查 API Server 地址是否可访问
- 查看后端日志获取详细错误信息

### 2. 无法查看 Pod 列表

**问题**：点击 Pod 管理页面无数据

**解决方案**：
- 确认集群状态为 "在线"
- 检查是否选择了正确的命名空间
- 检查用户是否有足够的权限访问该集群

### 3. 日志无法显示

**问题**：点击查看日志显示失败

**解决方案**：
- 检查 Pod 是否在运行
- 检查容器名称是否正确
- 检查是否有多行日志需要滚动显示

## 📊 数据库表结构

### k8s_clusters 表

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | bigint unsigned | 主键 ID |
| name | varchar(100) | 集群名称（唯一） |
| kube_config | longtext | kubeconfig 配置内容 |
| endpoint | varchar(500) | API Server 地址 |
| version | varchar(50) | K8s 版本 |
| status | varchar(20) | 集群状态（online/offline/unknown） |
| description | varchar(500) | 集群描述 |
| region | varchar(100) | 区域 |
| provider | varchar(50) | 云服务商 |
| is_default | tinyint(1) | 是否默认集群 |
| node_count | int | 节点数量 |
| created_at | datetime(3) | 创建时间 |
| updated_at | datetime(3) | 更新时间 |
| deleted_at | datetime(3) | 删除时间 |

## 🔄 后续开发计划

- [ ] StatefulSet 管理
- [ ] DaemonSet 管理
- [ ] ConfigMap/Secret 管理
- [ ] PVC/PV 管理
- [ ] Ingress 管理
- [ ] Pod 终端（WebSocket）
- [ ] 实时监控指标
- [ ] YAML 编辑器
- [ ] Helm Chart 支持
- [ ] 多集群统一视图

## 📄 开源协议

本项目采用 MIT 开源协议。

## 🙏 致谢

感谢 gin-vue-admin 框架和 Kubernetes 社区。

---

**注意**：本插件需要与 gin-vue-admin 框架配合使用，请确保框架版本兼容。
