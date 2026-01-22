# 算力节点管理

## 概述

算力节点管理功能用于统一管理多个 GPU 算力节点，支持 Docker TLS 安全连接，自动测试 Docker 连接状态，实现分布式 GPU 资源的集中化管理。

## 功能特性

### 核心功能

- ✅ **多节点管理**：统一管理多个 GPU 算力节点
- ✅ **TLS 安全连接**：支持 Docker TLS 安全连接
- ✅ **状态自动检测**：自动测试 Docker 连接状态
- ✅ **节点信息记录**：记录 GPU、CPU、内存、磁盘等硬件信息
- ✅ **显存切分配置**：支持 HAMi 显存切分配置

### 节点状态

| 状态 | 说明 |
|------|------|
| `connected` | 已连接 |
| `failed` | 连接失败 |
| `unknown` | 未知 |

## 数据模型

### 字段说明

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| 名字 | string | ✅ | 节点名称 |
| 区域 | string | | 节点区域 |
| CPU | string | | CPU 信息 |
| 内存 | string | | 内存信息 |
| 系统盘容量 | string | | 系统盘大小 |
| 数据盘容量 | string | | 数据盘大小 |
| IP地址公网 | string | ✅ | 公网 IP |
| IP地址内网 | string | ✅ | 内网 IP |
| SSH端口 | int | ✅ | 默认 22 |
| 用户名 | string | | SSH 用户名 |
| 密码 | string | | SSH 密码 |
| 显卡名称 | string | | GPU 型号 |
| 显卡数量 | int | | GPU 数量 |
| 显存容量 | int | | 单卡显存容量(GB) |
| HAMi-core目录 | string | | 节点 HAMi-core build 目录路径 |
| Docker连接地址 | string | | Docker API 地址 |
| 使用TLS | bool | | 默认启用 |
| CA证书 | text | | TLS CA 证书 |
| 客户端证书 | text | | TLS 客户端证书 |
| 客户端私钥 | text | | TLS 客户端私钥 |
| Docker状态 | string | | 自动测试 |
| 是否上架 | bool | ✅ | 默认上架 |
| 备注 | string | | 备注信息 |

## 使用指南

### 添加算力节点

1. 进入【算力节点】管理页面
2. 点击"新建节点"按钮
3. 填写节点基本信息：
   - 节点名称
   - 区域（可选）
   - IP 地址（公网和内网）
   - SSH 端口（默认 22）
   - SSH 用户名和密码（可选）
4. 填写 GPU 信息：
   - 显卡型号（如：NVIDIA GeForce RTX 3090）
   - 显卡数量
   - 显存容量（GB）
5. 配置 Docker 连接：
   - Docker API 地址（如：tcp://0.0.0.0:2375）
   - 是否使用 TLS（默认启用）
   - 如使用 TLS，上传证书文件：
     - CA 证书
     - 客户端证书
     - 客户端私钥
6. 配置 HAMi 显存切分（可选）：
   - HAMi-core 目录路径（如：/root/HAMi-core/build）
7. 点击"确定"创建节点

### Docker 连接测试

创建或更新节点时，系统会自动测试 Docker 连接状态：

1. **TLS 连接测试**（如果启用 TLS）：
   - 使用提供的证书尝试连接 Docker API
   - 验证证书有效性
   - 测试 API 可用性

2. **非 TLS 连接测试**：
   - 直接连接 Docker API
   - 测试基本 API 可用性

3. **状态更新**：
   - 连接成功：状态设为 `connected`
   - 连接失败：状态设为 `failed`
   - 无法测试：状态设为 `unknown`

### 编辑节点

1. 在节点列表中找到要编辑的节点
2. 点击"编辑"按钮
3. 修改节点信息
4. 点击"确定"保存更改
5. 系统会自动重新测试 Docker 连接

### 删除节点

1. 在节点列表中找到要删除的节点
2. 点击"删除"按钮
3. 确认删除操作

**注意**：
- 删除节点前，请确保没有正在运行的容器实例
- 删除节点后，无法恢复

## TLS 安全连接

### 证书配置

使用 TLS 连接时，需要提供以下证书：

1. **CA 证书（CA Certificate）**
   - 用于验证 Docker 服务端证书
   - PEM 格式

2. **客户端证书（Client Certificate）**
   - 用于客户端身份验证
   - PEM 格式

3. **客户端私钥（Client Key）**
   - 客户端私钥
   - PEM 格式

### 证书获取方式

在 Docker 服务器上执行以下命令获取证书：

```bash
# 查看 Docker TLS 配置
cat /etc/docker/daemon.json

# 通常证书位置
# CA 证书：/etc/docker/tls/ca.pem
# 服务器证书：/etc/docker/tls/cert.pem
# 服务器密钥：/etc/docker/tls/key.pem
# 客户端证书：/etc/docker/tls/client-cert.pem
# 客户端密钥：/etc/docker/tls/client-key.pem
```

### 启用 Docker TLS

1. 编辑 Docker 配置文件：
```bash
sudo vi /etc/docker/daemon.json
```

2. 添加以下配置：
```json
{
  "hosts": ["tcp://0.0.0.0:2375"],
  "tls": {
    "ca": "/etc/docker/tls/ca.pem",
    "cert": "/etc/docker/tls/server-cert.pem",
    "key": "/etc/docker/tls/server-key.pem",
    "client-cert": "/etc/docker/tls/client-cert.pem",
    "client-key": "/etc/docker/tls/client-key.pem"
  }
}
```

3. 重启 Docker 服务：
```bash
sudo systemctl restart docker
```

## HAMi 显存切分配置

### 什么是 HAMi？

HAMi 是一个开源的 GPU 显存虚拟化方案，支持将单块 GPU 的显存切分为多个虚拟 GPU，实现更灵活的资源配置。

### 配置步骤

1. **部署 HAMi**
   - 参考 [HAMi 官方文档](https://github.com/Project-HAMi/HAMi)
   - 在 GPU 服务器上部署 HAMi

2. **获取 HAMi-core 路径**
   ```bash
   # 查找 HAMi-core 目录
   find /root -name "HAMi-core" -type d
   # 通常路径：/root/HAMi-core/build
   ```

3. **在算力节点中配置**
   - 将 HAMi-core 目录路径填入"HAMi-core目录"字段
   - 例如：`/root/HAMi-core/build`

### 显存切分工作原理

当使用支持显存切分的产品规格创建容器时：

1. 系统自动从算力节点的 HAMi-core 目录读取配置
2. 将 HAMi-core 目录挂载到容器的 `/libvgpu/build`
3. 自动注入环境变量：
   - `LD_PRELOAD=/libvgpu/build/libvgpu.so`
   - `CUDA_DEVICE_MEMORY_LIMIT=显存大小`
   - `CUDA_DEVICE_SM_LIMIT=SM数量`

## 定时任务

### 自动状态检测

系统每 30 秒自动执行以下任务：

1. **检查所有算力节点的 Docker 连接状态**
   - 更新 DockerStatus 字段
   - 连接成功：`connected`
   - 连接失败：`failed`
   - 无法测试：`unknown`

2. **记录检测结果**
   - 仅记录失败的节点（Error 级别）
   - 降低控制台日志噪声

## 常见问题

### Q1: Docker 状态显示 "failed" 怎么办？

A: 请检查以下项目：
1. Docker 服务是否正在运行
2. Docker API 地址是否正确
3. 防火墙是否开放了 Docker 端口
4. TLS 证书是否正确配置

### Q2: 如何启用 Docker TLS 连接？

A:
1. 在 GPU 服务器上配置 Docker TLS
2. 获取必要的证书文件
3. 在添加节点时上传证书
4. 确保"使用TLS"选项已启用

### Q3: HAMi-core 目录路径应该填什么？

A: 填写 HAMi-core 项目编译后的 build 目录路径，通常是：
- `/root/HAMi-core/build`

### Q4: 如何测试节点是否可用？

A:
1. 创建节点时会自动测试
2. 可以手动点击"测试连接"按钮
3. 查看"状态"列了解连接情况

### Q5: 节点删除后还能恢复吗？

A:
- 删除节点后无法恢复
- 需要重新添加节点
- 建议删除前先确认

## API 接口

### 创建算力节点
```
POST /computeNode/createComputeNode
```

### 更新算力节点
```
PUT /computeNode/updateComputeNode
```

### 删除算力节点
```
DELETE /computeNode/deleteComputeNode
```

### 查询算力节点
```
GET /computeNode/getComputeNode
```

### 获取算力节点列表
```
GET /computeNode/getComputeNodeList
```

### 测试 Docker 连接
```
POST /computeNode/testConnection
```

## 相关文档

- [容器实例管理](Container-Instances.md)
- [产品规格管理](Product-Specs.md)
- [镜像库管理](Image-Registry.md)

---

**最后更新**：2025-01-22
