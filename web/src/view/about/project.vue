<template>
  <div class="project-document-container">
    <!-- 项目标题 -->
    <div class="project-header">
      <h1 class="project-title">天启算力管理平台</h1>
      <p class="project-subtitle">Docker GPU 算力资源管理平台</p>
    </div>

    <!-- 项目介绍 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-document /></el-icon>
          <span>📖 项目介绍</span>
        </div>
      </template>
      <div class="content-text">
        <p>
          <strong>Docker GPU 算力资源管理平台</strong> 是一个企业级的 GPU 容器化资源管理和调度系统，旨在帮助组织高效、安全地管理和分配 GPU 算力资源。平台采用现代化的微服务架构，提供从资源管理到容器实例全生命周期的完整解决方案。
        </p>
      </div>

      <!-- 项目图片 -->
      <div class="project-images">
        <el-image
          v-for="(img, index) in projectImages"
          :key="index"
          :src="img.src"
          :preview-src-list="projectImages.map(i => i.src)"
          :initial-index="index"
          fit="cover"
          class="project-image"
          :preview-teleported="true"
        >
          <template #error>
            <div class="image-error">
              <el-icon><icon-picture /></el-icon>
            </div>
          </template>
        </el-image>
      </div>
    </el-card>

    <!-- 项目目标 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-aim /></el-icon>
          <span>🎯 项目目标</span>
        </div>
      </template>
      <div class="content-text">
        <p class="mb-3">
          随着人工智能、深度学习、科学计算等领域的快速发展，GPU 算力资源已成为稀缺且昂贵的计算资源。传统的 GPU 资源管理方式存在以下痛点：
        </p>
        <ul class="pain-points">
          <li>❌ <strong>资源利用率低</strong>：GPU 资源分配不灵活，难以实现细粒度的资源切分和共享</li>
          <li>❌ <strong>管理成本高</strong>：多节点、多 GPU 环境下的资源管理复杂，缺乏统一的管理界面</li>
          <li>❌ <strong>安全性不足</strong>：缺乏完善的权限控制和访问审计机制</li>
          <li>❌ <strong>运维效率低</strong>：容器创建、监控、维护等操作需要大量人工干预</li>
        </ul>
        <p class="mt-3"><strong>本平台致力于解决上述问题，提供：</strong></p>
        <ul class="solutions">
          <li>✅ <strong>统一的资源管理</strong>：集中管理多个 GPU 算力节点，实现资源的统一调度和分配</li>
          <li>✅ <strong>灵活的资源配置</strong>：支持 GPU 显存切分，实现更细粒度的资源分配，提高资源利用率</li>
          <li>✅ <strong>完整的生命周期管理</strong>：从容器创建到删除的全流程自动化管理</li>
          <li>✅ <strong>安全可靠的访问控制</strong>：基于 RBAC 的权限管理，支持 SSH 跳板机和 Web 终端</li>
          <li>✅ <strong>实时监控与运维</strong>：容器状态自动同步、资源使用率实时监控、日志查看等运维功能</li>
        </ul>
      </div>
    </el-card>

    <!-- 核心特色 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-star /></el-icon>
          <span>🌟 核心特色</span>
        </div>
      </template>
      <div class="feature-grid">
        <div v-for="feature in features" :key="feature.title" class="feature-item">
          <div class="feature-icon">{{ feature.icon }}</div>
          <div class="feature-content">
            <h4>{{ feature.title }}</h4>
            <p>{{ feature.desc }}</p>
          </div>
        </div>
      </div>
    </el-card>

    <!-- 应用场景 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-grid /></el-icon>
          <span>🎨 应用场景</span>
        </div>
      </template>
      <div class="scenario-list">
        <el-tag
          v-for="scenario in scenarios"
          :key="scenario.title"
          class="scenario-tag"
          size="large"
          :type="scenario.type"
        >
          {{ scenario.icon }} {{ scenario.title }}
        </el-tag>
      </div>
    </el-card>

    <!-- 技术亮点 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-lightbulb /></el-icon>
          <span>🚀 技术亮点</span>
        </div>
      </template>
      <div class="tech-highlights">
        <el-row :gutter="20">
          <el-col
            v-for="highlight in techHighlights"
            :key="highlight.title"
            :xs="24"
            :sm="12"
            :md="6"
          >
            <div class="highlight-card">
              <div class="highlight-icon">{{ highlight.icon }}</div>
              <h4>{{ highlight.title }}</h4>
              <p>{{ highlight.desc }}</p>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-card>

    <!-- 核心功能模块 -->
    <el-collapse v-model="activeNames" class="function-collapse">
      <el-collapse-item
        v-for="module in functionModules"
        :key="module.id"
        :title="module.title"
        :name="module.id"
      >
        <div class="module-content">
          <div class="module-desc">{{ module.desc }}</div>

          <!-- 功能特性列表 -->
          <div v-if="module.features" class="module-features">
            <h4>✨ 功能特性</h4>
            <ul>
              <li v-for="feature in module.features" :key="feature">{{ feature }}</li>
            </ul>
          </div>

          <!-- 数据字段表格 -->
          <div v-if="module.fields" class="module-fields">
            <h4>📋 字段说明</h4>
            <el-table :data="module.fields" border style="width: 100%">
              <el-table-column prop="field" label="字段" width="180" />
              <el-table-column prop="type" label="类型" width="100" />
              <el-table-column prop="required" label="必填" width="80" />
              <el-table-column prop="desc" label="说明" />
            </el-table>
          </div>

          <!-- 支持操作 -->
          <div v-if="module.operations" class="module-operations">
            <h4>🔧 支持操作</h4>
            <div class="operation-tags">
              <el-tag
                v-for="op in module.operations"
                :key="op"
                type="info"
                effect="plain"
              >
                {{ op }}
              </el-tag>
            </div>
          </div>

          <!-- API接口 -->
          <div v-if="module.apis" class="module-apis">
            <h4>🔌 API接口</h4>
            <div class="api-list">
              <code v-for="api in module.apis" :key="api" class="api-item">
                {{ api }}
              </code>
            </div>
          </div>

          <!-- 配置说明 -->
          <div v-if="module.config" class="module-config">
            <h4>⚙️ 配置说明</h4>
            <pre class="config-code"><code>{{ module.config }}</code></pre>
          </div>
        </div>
      </el-collapse-item>
    </el-collapse>

    <!-- 技术栈 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-cpu /></el-icon>
          <span>💻 技术栈</span>
        </div>
      </template>
      <el-tabs v-model="activeTab">
        <el-tab-pane label="后端技术" name="backend">
          <div class="tech-list">
            <div v-for="tech in backendTech" :key="tech.name" class="tech-item">
              <el-tag type="primary">{{ tech.name }}</el-tag>
              <span>{{ tech.desc }}</span>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="前端技术" name="frontend">
          <div class="tech-list">
            <div v-for="tech in frontendTech" :key="tech.name" class="tech-item">
              <el-tag type="success">{{ tech.name }}</el-tag>
              <span>{{ tech.desc }}</span>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="数据库" name="database">
          <div class="tech-list">
            <div v-for="tech in databaseTech" :key="tech.name" class="tech-item">
              <el-tag type="warning">{{ tech.name }}</el-tag>
              <span>{{ tech.desc }}</span>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="容器技术" name="container">
          <div class="tech-list">
            <div v-for="tech in containerTech" :key="tech.name" class="tech-item">
              <el-tag type="info">{{ tech.name }}</el-tag>
              <span>{{ tech.desc }}</span>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>

    <!-- 部署指南 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-download /></el-icon>
          <span>📦 部署指南</span>
        </div>
      </template>
      <el-collapse>
        <el-collapse-item title="🔧 环境要求" name="env">
          <div class="env-requirements">
            <h4>后端环境：</h4>
            <ul>
              <li>Go 1.23+</li>
              <li>MySQL 5.7+ / PostgreSQL / SQLite / MSSQL / Oracle</li>
              <li>Redis（可选，用于缓存和会话管理）</li>
              <li>Docker（用于管理GPU容器）</li>
            </ul>
            <h4>前端环境：</h4>
            <ul>
              <li>Node.js 20+</li>
              <li>npm 或 pnpm</li>
            </ul>
          </div>
        </el-collapse-item>

        <el-collapse-item title="💻 本地开发部署" name="local">
          <div class="deploy-steps">
            <h4>1. 克隆项目</h4>
            <pre class="code-block"><code>git clone https://github.com/hequan2017/docker-gpu-manage
cd docker-gpu-manage
mv server/config.yaml.bak server/config.yaml</code></pre>

            <h4>2. 启动后端服务</h4>
            <pre class="code-block"><code>cd server
go mod download
go run main.go</code></pre>

            <h4>3. 启动前端服务</h4>
            <pre class="code-block"><code>cd web
npm install
npm run dev</code></pre>

            <h4>4. 访问系统</h4>
            <ul>
              <li>前端地址：http://localhost:8080</li>
              <li>后端API：http://localhost:8890</li>
              <li>Swagger文档：http://localhost:8890/swagger/index.html</li>
              <li>默认账号：admin / 123456</li>
            </ul>
          </div>
        </el-collapse-item>

        <el-collapse-item title="🐳 Docker 部署" name="docker">
          <div class="deploy-steps">
            <pre class="code-block"><code>cd deploy/docker-compose
docker-compose up -d</code></pre>
          </div>
        </el-collapse-item>

        <el-collapse-item title="☸️ Kubernetes 部署" name="k8s">
          <div class="deploy-steps">
            <pre class="code-block"><code>cd deploy/kubernetes
kubectl apply -f server/
kubectl apply -f web/</code></pre>
          </div>
        </el-collapse-item>
      </el-collapse>
    </el-card>

    <!-- 项目结构 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><icon-folder-opened /></el-icon>
          <span>📁 项目结构</span>
        </div>
      </template>
      <pre class="structure-tree"><code>├── server/                 # 后端代码
│   ├── api/v1/            # API 控制器
│   ├── model/             # 数据模型
│   ├── service/           # 业务逻辑
│   ├── router/            # 路由配置
│   ├── plugin/            # 插件系统
│   │   ├── portforward/   # 端口转发插件
│   │   └── k8smanager/    # K8s管理插件
│   └── config/            # 配置文件
├── web/                    # 前端代码
│   ├── src/api/           # API 调用
│   ├── src/view/          # 页面组件
│   └── src/plugin/        # 前端插件
└── README.md</code></pre>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  Document as IconDocument,
  Aim as IconAim,
  Star as IconStar,
  Grid as IconGrid,
  Lightbulb as IconLightbulb,
  Cpu as IconCpu,
  Download as IconDownload,
  FolderOpened as IconFolderOpened,
  Picture as IconPicture
} from '@element-plus/icons-vue'

// 项目图片
const projectImages = ref([
  { src: '/docs/4.png' },
  { src: '/docs/1.png' },
  { src: '/docs/2.png' },
  { src: '/docs/3.png' },
  { src: '/docs/5.png' }
])

// 核心特色
const features = ref([
  {
    icon: '🧠',
    title: '智能资源匹配',
    desc: '根据产品规格的 GPU 需求、显存需求、CPU、内存、磁盘等资源进行智能匹配，自动选择最优算力节点'
  },
  {
    icon: '✂️',
    title: '显存切分支持',
    desc: '支持 HAMi 显存切分技术，可以将单块 GPU 的显存切分为多个虚拟 GPU，实现更灵活的资源配置'
  },
  {
    icon: '🖥️',
    title: '多节点管理',
    desc: '支持管理多个分布式 GPU 算力节点，支持 TLS 安全连接，自动检测节点状态'
  },
  {
    icon: '🔐',
    title: 'SSH 跳板机服务',
    desc: '提供安全的 SSH 跳板机功能，用户可以通过 SSH 连接跳板机后选择并连接到自己的容器实例'
  },
  {
    icon: '💻',
    title: 'Web 终端集成',
    desc: '内置 Web 终端功能，无需额外工具即可在浏览器中直接操作容器'
  },
  {
    icon: '📊',
    title: '实时资源监控',
    desc: '提供 CPU、内存、网络 I/O、块设备 I/O、进程数等实时监控指标'
  },
  {
    icon: '⚡',
    title: '自动化运维',
    desc: '定时任务自动检查容器状态，保持数据同步，减少人工干预'
  }
])

// 应用场景
const scenarios = ref([
  { icon: '🤖', title: 'AI/ML 训练平台', type: '' },
  { icon: '🔬', title: '科研计算平台', type: 'success' },
  { icon: '☁️', title: '云服务提供商', type: 'info' },
  { icon: '🏢', title: '企业内部算力管理', type: 'warning' },
  { icon: '🎓', title: '教育机构', type: 'danger' }
])

// 技术亮点
const techHighlights = ref([
  { icon: '🔗', title: '前后端分离', desc: '采用 Gin + Vue3 的现代化技术栈' },
  { icon: '🧩', title: '微服务设计', desc: '模块化设计，易于扩展和维护' },
  { icon: '🔒', title: '安全可靠', desc: '支持 Docker TLS 安全连接' },
  { icon: '⚡', title: '高性能', desc: '基于 Go 语言开发，性能优异' }
])

// 核心功能模块
const activeNames = ref(['1'])
const functionModules = ref([
  {
    id: '1',
    title: '🐳 容器实例管理',
    desc: '管理 GPU 容器实例的完整生命周期，提供丰富的容器操作功能。',
    features: [
      '智能主机匹配：根据产品规格的GPU需求、显存需求、CPU、内存、磁盘等资源进行智能匹配',
      '显存切分支持：根据镜像的显存切分支持情况，自动过滤可用的产品规格',
      '实例名称校验：实例名称仅支持字母、数字、横线和下划线',
      '资源分配优化：支持按卡分配显存，更精确地管理GPU资源',
      '容器创建时如启用显存切分：从"算力节点"的 HAMi-core 目录字段读取路径'
    ],
    operations: [
      '查看详情', 'SSH连接', '启动', '停止', '重启',
      '日志查看', 'Web终端', '删除实例'
    ]
  },
  {
    id: '2',
    title: '🖥️ 算力节点管理',
    desc: '管理 GPU 算力节点，支持 Docker TLS 安全连接，自动测试 Docker 连接状态。',
    fields: [
      { field: '名字', type: 'string', required: '✅', desc: '节点名称' },
      { field: 'IP地址公网', type: 'string', required: '✅', desc: '公网 IP' },
      { field: '显卡数量', type: 'int', required: '', desc: 'GPU 数量' },
      { field: '显存容量', type: 'int', required: '', desc: '单卡显存容量(GB)' }
    ]
  },
  {
    id: '3',
    title: '🔐 SSH跳板机服务',
    desc: '提供SSH跳板机功能，用户可以通过SSH连接跳板机，然后选择并连接到自己的容器实例。',
    config: `# 在 server/config.yaml 中配置
jumpbox:
  enabled: true          # 是否启用SSH跳板机
  port: 2026            # SSH监听端口（默认2026）
  server-ip: "192.168.112.148"
  banner: "欢迎使用SSH跳板机服务\\r\\n"`
  },
  {
    id: '4',
    title: '🔁 端口转发管理',
    desc: '提供端口转发规则管理功能，支持TCP/UDP协议的端口映射，自动获取本机IP地址。',
    features: [
      '支持TCP和UDP协议',
      '自动获取服务器所有非127.0.0.1的IP地址作为默认源IP',
      '灵活的端口映射配置（源IP:端口 → 目标IP:端口）',
      '启用/禁用状态切换，实时生效',
      '批量删除支持，提高管理效率'
    ],
    apis: [
      'POST /api/portForward/createPortForward - 创建端口转发规则',
      'PUT /api/portForward/updatePortForward - 更新端口转发规则',
      'DELETE /api/portForward/deletePortForward - 删除端口转发规则',
      'GET /api/portForward/getPortForwardList - 获取规则列表'
    ]
  }
])

// 技术栈数据
const activeTab = ref('backend')
const backendTech = ref([
  { name: 'Gin', desc: 'Go Web框架' },
  { name: 'GORM', desc: 'Go ORM库' },
  { name: 'Go 1.23+', desc: '开发语言' },
  { name: 'golang.org/x/crypto/ssh', desc: 'SSH服务' },
  { name: 'Docker API', desc: 'Docker客户端' },
  { name: 'Zap', desc: '高性能日志库' }
])

const frontendTech = ref([
  { name: 'Vue 3', desc: '前端框架' },
  { name: 'Element Plus', desc: 'UI组件库' },
  { name: 'Vite', desc: '构建工具' },
  { name: 'Pinia', desc: '状态管理' },
  { name: 'Vue Router', desc: '路由管理' }
])

const databaseTech = ref([
  { name: 'MySQL', desc: '主数据库（推荐）' },
  { name: 'PostgreSQL', desc: '支持' },
  { name: 'SQLite', desc: '支持' },
  { name: 'MSSQL', desc: '支持' },
  { name: 'Oracle', desc: '支持' }
])

const containerTech = ref([
  { name: 'Docker', desc: '容器引擎' },
  { name: 'Docker TLS', desc: '安全连接' },
  { name: 'HAMi', desc: '显存切分' }
])
</script>

<style scoped lang="scss">
.project-document-container {
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
  background: #f5f7fa;
  min-height: 100vh;
}

.project-header {
  text-align: center;
  margin-bottom: 30px;
  padding: 40px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  color: white;

  .project-title {
    font-size: 48px;
    font-weight: bold;
    margin: 0 0 10px 0;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
  }

  .project-subtitle {
    font-size: 20px;
    margin: 0;
    opacity: 0.95;
  }
}

.section-card {
  margin-bottom: 20px;
  border-radius: 8px;

  :deep(.el-card__header) {
    background: linear-gradient(to right, #f6f8f9, #ffffff);
    border-bottom: 2px solid #e4e7ed;
  }
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 600;
  color: #303133;

  .el-icon {
    font-size: 20px;
    color: #409eff;
  }
}

.content-text {
  line-height: 1.8;
  color: #606266;
  font-size: 15px;

  p {
    margin-bottom: 15px;
  }

  strong {
    color: #303133;
    font-weight: 600;
  }
}

.pain-points {
  list-style: none;
  padding: 0;
  margin: 20px 0;

  li {
    padding: 12px 20px;
    margin-bottom: 10px;
    background: #fef0f0;
    border-left: 4px solid #f56c6c;
    border-radius: 4px;
    transition: all 0.3s;

    &:hover {
      background: #fde2e2;
      transform: translateX(5px);
    }
  }
}

.solutions {
  list-style: none;
  padding: 0;
  margin: 20px 0;

  li {
    padding: 12px 20px;
    margin-bottom: 10px;
    background: #f0f9ff;
    border-left: 4px solid #67c23a;
    border-radius: 4px;
    transition: all 0.3s;

    &:hover {
      background: #e1f3d8;
      transform: translateX(5px);
    }
  }
}

.project-images {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
  margin-top: 20px;

  .project-image {
    width: 100%;
    height: 200px;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    transition: all 0.3s;
    cursor: pointer;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
    }

    .image-error {
      display: flex;
      align-items: center;
      justify-content: center;
      height: 100%;
      background: #f5f7fa;
      color: #909399;
      font-size: 48px;
    }
  }
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;

  .feature-item {
    display: flex;
    gap: 15px;
    padding: 20px;
    background: linear-gradient(135deg, #f6f8f9 0%, #ffffff 100%);
    border-radius: 8px;
    border: 1px solid #e4e7ed;
    transition: all 0.3s;

    &:hover {
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      transform: translateY(-2px);
    }

    .feature-icon {
      font-size: 40px;
      flex-shrink: 0;
    }

    .feature-content {
      flex: 1;

      h4 {
        margin: 0 0 8px 0;
        font-size: 16px;
        color: #303133;
      }

      p {
        margin: 0;
        font-size: 14px;
        color: #606266;
        line-height: 1.6;
      }
    }
  }
}

.scenario-list {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;

  .scenario-tag {
    font-size: 15px;
    padding: 12px 20px;
    cursor: default;
  }
}

.tech-highlights {
  .highlight-card {
    text-align: center;
    padding: 25px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 12px;
    color: white;
    margin-bottom: 15px;
    transition: all 0.3s;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
    }

    .highlight-icon {
      font-size: 48px;
      margin-bottom: 15px;
    }

    h4 {
      margin: 0 0 10px 0;
      font-size: 18px;
    }

    p {
      margin: 0;
      font-size: 14px;
      opacity: 0.9;
    }
  }
}

.function-collapse {
  margin-bottom: 20px;

  :deep(.el-collapse-item__header) {
    font-size: 16px;
    font-weight: 600;
    background: linear-gradient(to right, #f6f8f9, #ffffff);
    padding: 0 20px;
  }
}

.module-content {
  padding: 20px;

  h4 {
    margin: 20px 0 15px 0;
    font-size: 16px;
    color: #303133;
    border-left: 4px solid #409eff;
    padding-left: 10px;
  }

  .module-desc {
    color: #606266;
    line-height: 1.8;
    margin-bottom: 20px;
  }

  .module-features {
    ul {
      list-style: none;
      padding: 0;

      li {
        padding: 10px 15px;
        margin-bottom: 8px;
        background: #f0f9ff;
        border-left: 3px solid #409eff;
        border-radius: 4px;
        color: #606266;
      }
    }
  }

  .module-operations {
    .operation-tags {
      display: flex;
      flex-wrap: wrap;
      gap: 10px;
    }
  }

  .module-apis {
    .api-list {
      display: flex;
      flex-direction: column;
      gap: 10px;

      .api-item {
        padding: 10px 15px;
        background: #f5f7fa;
        border-radius: 4px;
        font-family: 'Courier New', monospace;
        font-size: 13px;
        color: #409eff;
        border-left: 3px solid #409eff;
      }
    }
  }

  .module-config {
    .config-code {
      background: #282c34;
      color: #abb2bf;
      padding: 20px;
      border-radius: 8px;
      overflow-x: auto;
      font-size: 14px;
      line-height: 1.6;
    }
  }
}

.tech-list {
  display: flex;
  flex-direction: column;
  gap: 15px;

  .tech-item {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 15px 20px;
    background: #f5f7fa;
    border-radius: 8px;
    transition: all 0.3s;

    &:hover {
      background: #e4e7ed;
      transform: translateX(5px);
    }

    span {
      flex: 1;
      color: #606266;
      font-size: 15px;
    }
  }
}

.env-requirements {
  h4 {
    margin: 20px 0 10px 0;
    color: #303133;
  }

  ul {
    margin: 10px 0 20px 20px;
    li {
      margin-bottom: 8px;
      color: #606266;
    }
  }
}

.deploy-steps {
  h4 {
    margin: 20px 0 10px 0;
    color: #303133;
  }

  .code-block {
    background: #282c34;
    color: #abb2bf;
    padding: 15px;
    border-radius: 6px;
    overflow-x: auto;
    margin: 10px 0 20px 0;
    font-size: 14px;
    line-height: 1.6;
  }

  ul {
    margin: 10px 0 20px 20px;
    li {
      margin-bottom: 8px;
      color: #606266;
    }
  }
}

.structure-tree {
  background: #282c34;
  color: #abb2bf;
  padding: 25px;
  border-radius: 8px;
  overflow-x: auto;
  font-size: 14px;
  line-height: 1.8;
  font-family: 'Courier New', monospace;
}

.mb-3 {
  margin-bottom: 15px;
}

.mt-3 {
  margin-top: 15px;
}
</style>
