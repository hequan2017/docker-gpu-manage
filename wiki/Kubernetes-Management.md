# Kubernetes 集群管理

## 概述

Kubernetes 集群管理功能提供完整的 K8s 集群管理能力，支持多集群统一管理，包含工作负载、Service、Pod 等资源的全生命周期管理，同时提供基于 RBAC 的多级权限控制和操作审计功能。

## 功能特性

### 核心功能

- ✅ **多集群管理**：支持添加、编辑、删除多个 K8s 集群
- ✅ **Kubeconfig 加密存储**：使用 AES-256-GCM 加密存储 kubeconfig
- ✅ **连接池管理**：自动清理过期连接，避免资源泄漏
- ✅ **工作负载管理**：管理 Deployment、StatefulSet、DaemonSet
- ✅ **Pod 管理**：列表、详情、日志查看、自动刷新
- ✅ **Service 管理**：管理集群服务，查看端点信息
- ✅ **监控指标**：收集集群、节点、Pod 资源使用率
- ✅ **RBAC 权限控制**：全局/集群/命名空间三级权限
- ✅ **操作审计**：记录所有 K8s 操作日志

## 集群配置

### 数据模型

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| 名称 | string | ✅ | 集群名称 |
| Kubeconfig | text | ✅ | K8s 集群配置（加密存储） |
| 描述 | string | | 集群描述 |
| 状态 | bool | ✅ | 是否启用 |

### Kubeconfig 格式

Kubeconfig 应包含完整的集群配置：

```yaml
apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority-data: <base64-encoded-ca>
    server: https://<cluster-endpoint>
  name: <cluster-name>
contexts:
- context:
    cluster: <cluster-name>
    user: <user-name>
    namespace: <default-namespace>
  name: <context-name>
current-context: <context-name>
users:
- name: <user-name>
  user:
    client-certificate-data: <base64-encoded-cert>
    client-key-data: <base64-encoded-key>
```

## 使用指南

### 添加集群

1. 进入【K8s 集群】管理页面
2. 点击"新建集群"按钮
3. 填写集群信息：
   - 集群名称
   - 上传或粘贴 Kubeconfig 内容
   - 集群描述（可选）
4. 点击"确定"添加集群

### 编辑集群

1. 在集群列表中找到要编辑的集群
2. 点击"编辑"按钮
3. 修改集群信息
4. 点击"确定"保存更改

### 删除集群

1. 在集群列表中找到要删除的集群
2. 点击"删除"按钮
3. 确认删除操作

**注意**：删除集群后，所有相关的配置和权限数据也会被删除。

### 启用/禁用集群

1. 在集群列表中找到要操作的集群
2. 点击"启用"或"禁用"按钮
3. 确认操作

## 工作负载管理

### Deployment 管理

**列表查看：**
- 副本数、镜像、运行时间
- 命名空间、标签选择器
- 按集群和命名空间分组

**详情查看：**
- YAML 配置
- 容器配置
- 环境变量
- 资源配额

**操作功能：**
- 扩缩容：调整副本数
- 重启：滚动更新
- 删除：删除 Deployment
- Pods 查看：关联的 Pod 列表

### StatefulSet 管理

**列表查看：**
- 副本数、服务名称、运行时间
- 存储卷声明
- 按集群和命名空间分组

**详情查看：**
- YAML 配置
- 存储卷配置
- 服务名称

**操作功能：**
- 扩缩容：调整副本数
- 重启：滚动更新
- 删除：删除 StatefulSet

### DaemonSet 管理

**列表查看：**
- 节点选择器
- 当前调度数、期望调度数
- 就绪节点数
- 按集群和命名空间分组

**详情查看：**
- YAML 配置
- 节点调度器
- 更新策略

**操作功能：**
- 删除：删除 DaemonSet
- 节点管理：管理调度节点

## Pod 管理

### 列表查看

- 状态：Running、Pending、Succeeded、Failed
- IP 地址、节点名称
- 重启次数、创建时间
- 按集群和命名空间分组
- 支持自动刷新（5秒间隔）

### 详情查看

- YAML 配置
- 容器配置
- 资源使用情况
- 事件记录

### 日志查看

- 支持配置行数
- 实时查看日志
- 支持日志过滤

### Pod 操作

- 删除 Pod
- 查看 Pod 详情
- 查看 Pod 事件

## Service 管理

### 列表查看

- 服务类型：ClusterIP、NodePort、LoadBalancer
- 端口映射：ClusterIP、NodePort
- 选择器标签
- 端点信息

### 详情查看

- YAML 配置
- 端口配置
- 选择器配置
- 关联的 Pods

### 端点查看

- 端口 IP 地址
- NodePort 地址
- 关联的 Pod 列表

### Service 操作

- 删除 Service
- 查看 Service 详情
- 查看端点

## 权限控制

### 三级权限模型

#### 1. 全局权限

拥有所有集群的所有资源的完全访问权限。

**适用角色：** 超级管理员

**权限范围：**
- 所有集群的所有资源
- 所有命名空间
- 创建、删除集群

#### 2. 集群权限

拥有指定集群的所有资源的完全访问权限。

**适用角色：** 集群管理员

**权限范围：**
- 指定集群的所有资源
- 该集群的所有命名空间
- 创建、删除命名空间

#### 3. 命名空间权限

拥有指定集群的指定命名空间的有限访问权限。

**适用角色：** 开发人员、运维人员

**权限范围：**
- 指定集群的指定命名空间
- 创建、删除资源（根据角色配置）

### 权限配置

权限在后台的 `sys_casbin` 表中配置：

```
p, 888, k8s:cluster:*, *, *, *
p, 888, k8s:node:*, *, *
p, 888, k8s:pod:*, *, *
```

## 监控指标

### 集群级别

- 节点总数
- Pod 总数
- 资源使用率（CPU、内存）

### 节点级别

- CPU 使用率
- 内存使用率
- 磁盘使用率
- 网络流量

### Pod 级别

- CPU 使用率
- 内存使用率
- 网络流量
- 磁盘 I/O

## 操作审计

### 审计日志

系统会记录以下操作：

- 集群的添加、编辑、删除
- 资源的创建、更新、删除
- 权限的变更
- 配置的修改

### 日志查看

1. 进入【操作日志】页面
2. 筛选模块：选择 "Kubernetes"
3. 查看详细操作记录

## 常见问题

### Q1: 如何获取 Kubeconfig 文件？

A:
```bash
# 默认位置
~/.kube/config

# 复制 kubeconfig
cat ~/.kube/config

# 或指定命名空间的 kubeconfig
kubectl config view --raw
```

### Q2: 连接集群失败怎么办？

A: 请检查以下项目：
1. Kubeconfig 配置是否正确
2. 集群 API 地址是否可达
3. 网络连接是否正常
4. 证书是否有效
5. 权限配置是否正确

### Q3: 如何查看 Pod 日志？

A:
1. 在 Pod 列表中找到目标 Pod
2. 点击"日志"按钮
3. 选择日志行数（默认 100 行）
4. 查看实时日志输出

### Q4: 如何扩缩容 Deployment？

A:
1. 在 Deployment 详情页面
2. 点击"扩缩容"按钮
3. 输入期望的副本数
4. 确认扩缩容

### Q5: 权限不足怎么办？

A:
1. 联系管理员分配相应权限
2. 确认用户角色和权限范围
3. 检查 RBAC 配置是否正确

## API 接口

### 集群管理

```
POST /k8s/cluster/create      # 创建集群
PUT /k8s/cluster/update       # 更新集群
DELETE /k8s/cluster/delete     # 删除集群
GET  /k8s/cluster/list        # 获取集群列表
```

### 工作负载管理

```
GET  /k8s/deployment/list      # 获取 Deployment 列表
GET  /k8s/deployment/detail     # 获取 Deployment 详情
POST /k8s/deployment/scale     # 扩缩容 Deployment
POST /k8s/deployment/restart    # 重启 Deployment
```

### Pod 管理

```
GET  /k8s/pod/list              # 获取 Pod 列表
GET  /k8s/pod/detail            # 获取 Pod 详情
GET  /k8s/pod/logs              # 获取 Pod 日志
DELETE /k8s/pod/delete          # 删除 Pod
```

### Service 管理

```
GET  /k8s/service/list           # 获取 Service 列表
GET  /k8s/service/detail         # 获取 Service 详情
DELETE /k8s/service/delete       # 删除 Service
```

### 监控指标

```
GET  /k8s/metrics/cluster      # 获取集群指标
GET  /k8s/metrics/nodes         # 获取节点指标
GET  /k8s/metrics/pods          # 获取 Pod 指标
```

## 技术实现

- **后端插件**：`server/plugin/k8smanager/`
- **前端插件**：`web/src/plugin/k8smanager/`
- **K8s 客户端**：`k8s.io/client-go`
- **加密算法**：AES-256-GCM
- **连接池**：支持 TTL 和自动清理

## 相关文档

- [权限管理](Authorization.md)
- [定时任务](Scheduled-Tasks.md)

---

**最后更新**：2025-01-22
