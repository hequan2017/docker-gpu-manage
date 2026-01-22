# 天启算力管理平台 - 功能文档

欢迎来到天启算力管理平台的官方文档！本文档按照功能模块组织，详细介绍平台的各个功能模块。

## 📚 目录

### Docker GPU 管理
| 模块 | 描述 | 状态 |
|------|------|------|
| [容器实例管理](Container-Instances.md) | GPU 容器实例的完整生命周期管理 | ✅ |
| [算力节点管理](Compute-Nodes.md) | 多 GPU 算力节点的统一管理 | ✅ |
| [镜像库管理](Image-Registry.md) | Docker 镜像仓库管理 | ✅ |
| [产品规格管理](Product-Specs.md) | GPU 产品规格和定价定义 | ✅ |
| [端口转发管理](Port-Forwarding.md) | 灵活网络端口转发规则管理 | ✅ |
| [SSH 跳板机](SSH-Jumpbox.md) | 安全的 SSH 容器连接服务 | ✅ |

### Kubernetes 集群管理
| 模块 | 描述 | 状态 |
|------|------|------|
| [K8s 集群管理](Kubernetes-Management.md) | 多 K8s 集群统一管理 | ✅ |

### 系统功能
| 模块 | 描述 | 状态 |
|------|------|------|
| [权限管理](Authorization.md) | 基于 RBAC 的权限控制系统 | 📝 |
| [定时任务](Scheduled-Tasks.md) | 自动化定时任务管理 | 📝 |

### 资产管理
| 模块 | 描述 | 状态 |
|------|------|------|
| [戴尔资产管理](Dell-Asset-Management.md) | 物理服务器资产全生命周期管理 | ✅ |

### AI 智能化
| 模块 | 描述 | 状态 |
|------|------|------|
| [AI Agent 智能助手](AI-Agent.md) | 集成智谱 GLM-4.7 的智能对话助手 | ✅ |
| [模型微调](Model-Finetuning.md) | LLaMA 等大模型的微调任务管理 | ✅ |

## 🚀 快速开始

### 环境要求

**后端环境：**
- Go 1.23+
- MySQL 5.7+ / PostgreSQL / SQLite / MSSQL / Oracle
- Redis（可选，用于缓存和会话管理）
- Docker（用于管理GPU容器）

**前端环境：**
- Node.js 20+
- npm 或 pnpm

### 快速部署

```bash
# 克隆项目
git clone https://github.com/he7555/docker-gpu-manage.git
cd docker-gpu-manage

# 复制配置文件
cp server/config.yaml.bak server/config.yaml

# 启动后端
cd server
go mod download
go run main.go

# 启动前端（新终端）
cd web
npm install
npm run dev
```

### 访问系统

- 前端地址：`http://localhost:8080`
- 后端API：`http://localhost:8890`
- Swagger文档：`http://localhost:8890/swagger/index.html`

**默认管理员账号：**
- 用户名：`admin`
- 密码：`123456`

## 🎯 核心特性

### 智能资源调度
- 根据产品规格的GPU需求、显存需求、CPU、内存、磁盘等资源进行智能匹配
- 支持 HAMi 显存切分技术，实现更细粒度的资源分配
- 自动选择最优算力节点

### 安全可靠
- 基于 RBAC 的权限控制系统
- 支持 Docker TLS 安全连接
- SSH 跳板机安全访问
- 完整的操作审计日志

### 高效运维
- 实时资源监控（CPU、内存、网络I/O、块设备I/O、进程数）
- 容器状态自动同步
- 定时任务自动检查
- Web 终端在线操作

## 📖 文档说明

### 文档结构

本文档采用 GitHub Wiki 格式组织，每个功能模块都有独立的文档页面。

### 文档状态

- ✅ 已完成
- 🚧 进行中
- 📝 计划中

### 贡献指南

欢迎参与文档改进！如果您发现文档有任何问题或建议，欢迎提交 Issue 或 Pull Request。

## 📞 联系方式

- 技术支持：support@gpu-manage.io
- 商务合作：business@gpu-manage.io
- GitHub Issues：[提交问题](https://github.com/he7555/docker-gpu-manage/issues)

## 🔗 相关链接

- [项目主页](https://github.com/he7555/docker-gpu-manage)
- [在线官网](./website/index.html)
- [API 文档](http://localhost:8890/swagger/index.html)
- [更新日志](CHANGELOG.md)

---

**最后更新时间**：2025-01-22
