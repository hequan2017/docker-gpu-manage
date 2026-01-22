# 镜像库管理

## 概述

镜像库管理功能用于统一管理 Docker 镜像仓库，支持多镜像源配置，为容器实例提供镜像资源。同时支持显存切分功能配置，实现更灵活的 GPU 资源分配。

## 功能特性

### 核心功能

- ✅ **多镜像源支持**：统一管理多个镜像仓库
- ✅ **显存切分配置**：为镜像库配置显存切分支持
- ✅ **上架/下架管理**：控制镜像库的可用状态
- ✅ **镜像信息记录**：记录镜像库的详细信息和描述

## 数据模型

### 字段说明

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| 名字 | string | ✅ | 镜像库名称 |
| 地址 | string | ✅ | 镜像库地址 |
| 描述 | string | | 镜像库描述 |
| 来源 | string | | 镜像库来源 |
| 是否支持显存切分 | bool | ✅ | 默认否 |
| 是否上架 | bool | ✅ | 默认上架 |
| 备注 | string | | 备注信息 |

## 使用指南

### 添加镜像库

1. 进入【镜像库】管理页面
2. 点击"新建镜像库"按钮
3. 填写镜像库信息：
   - 镜像库名称
   - 镜像库地址（如：registry.example.com/project/image）
   - 描述（可选）
   - 来源（可选）
   - 是否支持显存切分（默认否）
   - 是否上架（默认上架）
   - 备注（可选）
4. 点击"确定"创建镜像库

### 编辑镜像库

1. 在镜像库列表中找到要编辑的镜像库
2. 点击"编辑"按钮
3. 修改镜像库信息
4. 点击"确定"保存更改

### 删除镜像库

1. 在镜像库列表中找到要删除的镜像库
2. 点击"删除"按钮
3. 确认删除操作

**注意**：
- 删除镜像库前，请确保没有正在使用该镜像库的容器实例
- 删除镜像库后，无法恢复

### 上架/下架镜像库

1. 在镜像库列表中找到要操作的镜像库
2. 点击"上架"或"下架"按钮
3. 确认操作

- **上架**：镜像库可以在创建容器时选择
- **下架**：镜像库不可见，无法创建新容器

## 显存切分配置

### 什么是显存切分？

显存切分是指将单块 GPU 的显存切分为多个虚拟 GPU，实现更灵活的显存资源分配。HAMi 是一个开源的 GPU 显存虚拟化方案。

### 配置步骤

1. **在镜像库中启用显存切分**
   - 创建或编辑镜像库时
   - 将"是否支持显存切分"设置为"是"

2. **在产品规格中启用显存切分**
   - 创建或编辑产品规格时
   - 将"是否支持显存切分"设置为"是"

3. **创建容器时自动应用**
   - 选择支持显存切分的镜像和产品规格
   - 系统自动配置显存切分相关参数

### 显存切分工作原理

使用支持显存切分的镜像和产品规格创建容器时：

1. 系统自动从算力节点的 HAMi-core 目录读取配置
2. 将 HAMi-core 目录挂载到容器的 `/libvgpu/build`
3. 自动注入环境变量：
   ```bash
   LD_PRELOAD=/libvgpu/build/libvgpu.so
   CUDA_DEVICE_MEMORY_LIMIT=<显存大小>
   CUDA_DEVICE_SM_LIMIT=<SM数量>
   ```

## 镜像库地址格式

### Docker Hub 镜像

```
# 格式
nginx:latest
python:3.9-slim

# 完整格式
docker.io/library/nginx:latest
```

### 私有镜像仓库

```
# 格式
registry.example.com/project/image:tag

# 示例
registry.company.com/ml/tensorflow:2.10.0-gpu
harbor.company.com/project/image:v1.0
```

### 完整 URL 格式

```
# 带认证的镜像仓库
username:password@registry.example.com/project/image:tag

# 示例
admin:password123@registry.company.com/project/image:latest
```

## 镜像拉取

### 自动拉取

创建容器实例时，系统会自动从配置的镜像库拉取镜像：
- 如果镜像不存在于本地，会自动从镜像库拉取
- 支持私有镜像仓库（需配置认证信息）
- 显示拉取进度和状态

### 手动拉取

如需手动拉取镜像测试：

```bash
# 拉取公共镜像
docker pull nginx:latest

# 拉取私有镜像
docker login registry.example.com
docker pull registry.example.com/project/image:tag
```

## 镜像管理最佳实践

### 镜像版本管理

1. **使用明确的版本标签**
   - ✅ 推荐：`tensorflow:2.10.0-gpu`
   - ❌ 不推荐：`tensorflow:latest`

2. **版本标签规范**
   - 主版本：`v1.0`
   - 次版本：`v1.2.3`
   - 构建号：`v1.0.0-build.123`

3. **多版本支持**
   - 为同一镜像创建多个镜像库条目
   - 使用不同的标签区分版本

### 镜像优化

1. **减小镜像体积**
   - 使用 alpine 基础镜像
   - 清理不必要的文件
   - 多阶段构建

2. **安全性**
   - 及时更新基础镜像
   - 扫描镜像漏洞
   - 最小化镜像权限

## 常见问题

### Q1: 镜像拉取失败怎么办？

A: 请检查以下项目：
1. 镜像库地址是否正确
2. 网络连接是否正常
3. 镜像是否存在
4. 私有镜像仓库是否已配置认证

### Q2: 如何配置私有镜像仓库？

A:
1. 在 GPU 节点上登录镜像仓库：
   ```bash
   docker login registry.example.com
   ```
2. 输入用户名和密码
3. 认证信息会保存在 `~/.docker/config.json`

### Q3: 显存切分镜像和普通镜像有什么区别？

A:
- **显存切分镜像**：
  - 需要配置"是否支持显存切分"为"是"
  - 创建容器时会自动配置显存切分参数
  - 适用于需要 GPU 显存切分的场景

- **普通镜像**：
  - 不支持显存切分或默认关闭
  - 创建容器时使用常规 GPU 配置

### Q4: 如何测试镜像库是否可用？

A:
1. 在 GPU 节点上执行：
   ```bash
   docker pull <镜像名称>
   ```
2. 检查镜像是否成功拉取
3. 尝试使用镜像创建容器测试

### Q5: 镜像库可以修改吗？

A:
- 可以修改镜像库的描述、状态等信息
- 修改镜像库地址后，已创建的容器不受影响
- 需要更新镜像时，建议创建新的镜像库条目

## API 接口

### 创建镜像库
```
POST /imageRegistry/createImageRegistry
```

### 更新镜像库
```
PUT /imageRegistry/updateImageRegistry
```

### 删除镜像库
```
DELETE /imageRegistry/deleteImageRegistry
```

### 查询镜像库
```
GET /imageRegistry/findImageRegistry
```

### 获取镜像库列表
```
GET /imageRegistry/getImageRegistryList
```

## 相关文档

- [容器实例管理](Container-Instances.md)
- [算力节点管理](Compute-Nodes.md)
- [产品规格管理](Product-Specs.md)

---

**最后更新**：2025-01-22
