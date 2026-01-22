<template>
  <div class="dashboard-container">
    <!-- 项目标题 -->
    <div class="project-header">
      <h1 class="project-title">天启算力管理平台</h1>
      <p class="project-subtitle">Docker GPU 算力资源管理平台</p>
      <div class="header-stats">
        <span class="stat-item">
          <i class="el-icon"><Server /></i>
          <span>算力节点管理</span>
        </span>
        <span class="stat-item">
          <i class="el-icon"><Box /></i>
          <span>容器实例调度</span>
        </span>
        <span class="stat-item">
          <i class="el-icon"><Monitor /></i>
          <span>资源实时监控</span>
        </span>
      </div>
    </div>

    <!-- 核心功能快速入口 -->
    <el-row :gutter="20" class="quick-nav">
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('instance')">
          <div class="nav-icon">🐳</div>
          <div class="nav-title">实例管理</div>
          <div class="nav-desc">GPU容器实例</div>
          <div class="nav-badge">热门</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('computeNode')">
          <div class="nav-icon">🖥️</div>
          <div class="nav-title">算力节点</div>
          <div class="nav-desc">节点管理</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('imageRegistry')">
          <div class="nav-icon">📦</div>
          <div class="nav-title">镜像库</div>
          <div class="nav-desc">镜像管理</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('productSpec')">
          <div class="nav-icon">💎</div>
          <div class="nav-title">产品规格</div>
          <div class="nav-desc">规格配置</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('portForward')">
          <div class="nav-icon">🔁</div>
          <div class="nav-title">端口转发</div>
          <div class="nav-desc">转发规则</div>
          <div class="nav-badge">新功能</div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4">
        <div class="nav-card" @click="goTo('project')">
          <div class="nav-icon">📖</div>
          <div class="nav-title">项目说明</div>
          <div class="nav-desc">完整文档</div>
        </div>
      </el-col>
    </el-row>

    <!-- 平台简介和特性 -->
    <el-card class="intro-card" shadow="hover">
      <template #header>
        <span class="card-title">🚀 平台简介</span>
      </template>
      <div class="intro-content">
        <p class="intro-text">
          <strong>天启算力管理平台</strong> 是一个企业级的 GPU 容器化资源管理和调度系统，
          采用现代化的微服务架构，提供从资源管理到容器实例全生命周期的完整解决方案。
          平台支持 <el-tag type="success" size="small">HAMi 显存切分</el-tag>、
          <el-tag type="warning" size="small">SSH 跳板机</el-tag>、
          <el-tag type="info" size="small">Web 终端</el-tag> 等企业级功能。
        </p>

        <el-divider content-position="left">核心特性</el-divider>

        <el-row :gutter="20">
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">🧠</div>
              <div class="feature-text">
                <h4>智能资源匹配</h4>
                <p>根据 GPU 需求、显存、CPU 等资源自动选择最优算力节点</p>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">✂️</div>
              <div class="feature-text">
                <h4>显存切分</h4>
                <p>支持 HAMi 显存虚拟化技术，提高 GPU 利用率</p>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">🖥️</div>
              <div class="feature-text">
                <h4>多节点管理</h4>
                <p>统一管理多个分布式算力节点，支持 TLS 安全连接</p>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">🔐</div>
              <div class="feature-text">
                <h4>SSH 跳板机</h4>
                <p>安全的 SSH 连接方式，支持跳板机访问容器</p>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">💻</div>
              <div class="feature-text">
                <h4>Web 终端</h4>
                <p>内置 Web 终端，浏览器中直接操作容器</p>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="8">
            <div class="feature-item">
              <div class="feature-icon">📊</div>
              <div class="feature-text">
                <h4>实时监控</h4>
                <p>CPU、内存、网络 I/O、进程数等实时监控</p>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-card>

    <!-- 应用场景 -->
    <el-card class="scenario-card" shadow="hover">
      <template #header>
        <span class="card-title">🎨 应用场景</span>
      </template>
      <el-row :gutter="15">
        <el-col :xs="12" :sm="8" :md="6" v-for="scenario in scenarios" :key="scenario.title">
          <div class="scenario-item">
            <div class="scenario-icon">{{ scenario.icon }}</div>
            <div class="scenario-content">
              <h5>{{ scenario.title }}</h5>
              <p>{{ scenario.desc }}</p>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!-- 技术架构和快速指南 -->
    <el-row :gutter="20">
      <!-- 技术栈 -->
      <el-col :xs="24" :md="12">
        <el-card class="tech-card" shadow="hover">
          <template #header>
            <span class="card-title">💻 技术栈</span>
          </template>
          <el-tabs v-model="activeTab">
            <el-tab-pane label="后端" name="backend">
              <div class="tech-list">
                <div class="tech-item" v-for="tech in backendTech" :key="tech.name">
                  <el-tag type="primary" size="large">{{ tech.name }}</el-tag>
                  <span>{{ tech.desc }}</span>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane label="前端" name="frontend">
              <div class="tech-list">
                <div class="tech-item" v-for="tech in frontendTech" :key="tech.name">
                  <el-tag type="success" size="large">{{ tech.name }}</el-tag>
                  <span>{{ tech.desc }}</span>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane label="数据库" name="database">
              <div class="tech-list">
                <div class="tech-item" v-for="tech in databaseTech" :key="tech.name">
                  <el-tag type="warning" size="large">{{ tech.name }}</el-tag>
                  <span>{{ tech.desc }}</span>
                </div>
              </div>
            </el-tab-pane>
          </el-tabs>
        </el-card>
      </el-col>

      <!-- 快速开始 -->
      <el-col :xs="24" :md="12">
        <el-card class="guide-card" shadow="hover">
          <template #header>
            <span class="card-title">🚀 快速开始</span>
          </template>
          <el-timeline>
            <el-timeline-item timestamp="步骤 1" color="#409eff">
              <h4>添加算力节点</h4>
              <p class="guide-text">在"算力节点"中添加您的 GPU 服务器，配置 Docker 连接信息</p>
              <el-button type="primary" link @click="goTo('computeNode')">前往添加 →</el-button>
            </el-timeline-item>
            <el-timeline-item timestamp="步骤 2" color="#67c23a">
              <h4>配置镜像库</h4>
              <p class="guide-text">添加 Docker 镜像仓库地址，支持私有仓库和显存切分配置</p>
              <el-button type="success" link @click="goTo('imageRegistry')">前往配置 →</el-button>
            </el-timeline-item>
            <el-timeline-item timestamp="步骤 3" color="#e6a23c">
              <h4>创建产品规格</h4>
              <p class="guide-text">定义 GPU 产品规格，包括显卡型号、数量、价格等</p>
              <el-button type="warning" link @click="goTo('productSpec')">前往创建 →</el-button>
            </el-timeline-item>
            <el-timeline-item timestamp="步骤 4" color="#f56c6c">
              <h4>创建实例</h4>
              <p class="guide-text">选择镜像和规格，系统自动匹配最优节点创建容器</p>
              <el-button type="danger" link @click="goTo('instance')">前往创建 →</el-button>
            </el-timeline-item>
          </el-timeline>
        </el-card>
      </el-col>
    </el-row>

    <!-- 功能亮点 -->
    <el-card class="highlight-card" shadow="hover">
      <template #header>
        <span class="card-title">✨ 功能亮点</span>
      </template>
      <el-row :gutter="20">
        <el-col :xs="24" :sm="12" :md="6" v-for="highlight in highlights" :key="highlight.title">
          <div class="highlight-item">
            <div class="highlight-number">{{ highlight.number }}</div>
            <div class="highlight-content">
              <h4>{{ highlight.title }}</h4>
              <p>{{ highlight.desc }}</p>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!-- 系统状态 -->
    <el-card class="status-card" shadow="hover">
      <template #header>
        <span class="card-title">📊 系统状态</span>
      </template>
      <el-row :gutter="20">
        <el-col :xs="12" :sm="6">
          <div class="status-item">
            <el-icon class="status-icon" color="#67c23a"><CircleCheck /></el-icon>
            <div class="status-info">
              <div class="status-label">系统状态</div>
              <div class="status-value">运行中</div>
            </div>
          </div>
        </el-col>
        <el-col :xs="12" :sm="6">
          <div class="status-item">
            <el-icon class="status-icon" color="#409eff"><Connection /></el-icon>
            <div class="status-info">
              <div class="status-label">Docker API</div>
              <div class="status-value">已连接</div>
            </div>
          </div>
        </el-col>
        <el-col :xs="12" :sm="6">
          <div class="status-item">
            <el-icon class="status-icon" color="#e6a23c"><Clock /></el-icon>
            <div class="status-info">
              <div class="status-label">定时任务</div>
              <div class="status-value">30秒/次</div>
            </div>
          </div>
        </el-col>
        <el-col :xs="12" :sm="6">
          <div class="status-item">
            <el-icon class="status-icon" color="#909399"><DocumentCopy /></el-icon>
            <div class="status-info">
              <div class="status-label">完整文档</div>
              <el-button type="text" @click="goTo('project')">查看 →</el-button>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!-- 页脚 -->
    <div class="footer">
      <p>© 2025 天启算力管理平台 v2.0 |
        <a href="https://github.com/hequan2017/docker-gpu-manage" target="_blank">GitHub</a> |
        <a href="javascript:void(0)" @click="goTo('project')">项目文档</a>
      </p>
      <p class="footer-links">
        基于 <el-tag size="small" type="primary">Go</el-tag> +
        <el-tag size="small" type="success">Vue 3</el-tag> +
        <el-tag size="small" type="info">Docker</el-tag> 构建
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  Server,
  Box,
  Monitor,
  CircleCheck,
  Connection,
  Clock,
  DocumentCopy
} from '@element-plus/icons-vue'

defineOptions({
  name: 'Dashboard'
})

const router = useRouter()
const activeTab = ref('backend')

const goTo = (path) => {
  // 跳转到业务页面，这些页面都在 layout 下
  // 例如: instance -> /instance (动态路由会自动添加 layout 前缀)
  router.push({ name: path })
}

// 应用场景
const scenarios = ref([
  { icon: '🤖', title: 'AI/ML 训练', desc: '机器学习模型训练' },
  { icon: '🔬', title: '科研计算', desc: '科学计算和仿真' },
  { icon: '☁️', title: '云服务', desc: '云服务提供商' },
  { icon: '🏢', title: '企业管理', desc: '企业内部算力管理' },
  { icon: '🎓', title: '教育机构', desc: '教学实验环境' },
  { icon: '💡', title: '创新研发', desc: '新技术研发测试' }
])

// 技术栈
const backendTech = ref([
  { name: 'Go 1.23+', desc: '高性能开发语言' },
  { name: 'Gin', desc: 'Web 框架' },
  { name: 'GORM', desc: 'ORM 框架' },
  { name: 'Docker API', desc: 'Docker 客户端' },
  { name: 'Zap', desc: '高性能日志' }
])

const frontendTech = ref([
  { name: 'Vue 3', desc: '前端框架' },
  { name: 'Element Plus', desc: 'UI 组件库' },
  { name: 'Vite', desc: '构建工具' },
  { name: 'Pinia', desc: '状态管理' },
  { name: 'Vue Router', desc: '路由管理' }
])

const databaseTech = ref([
  { name: 'MySQL', desc: '主数据库（推荐）' },
  { name: 'PostgreSQL', desc: '支持' },
  { name: 'SQLite', desc: '支持' },
  { name: 'Redis', desc: '缓存（可选）' }
])

// 功能亮点
const highlights = ref([
  { number: '01', title: '智能调度', desc: '自动匹配最优节点' },
  { number: '02', title: '显存切分', desc: '提高 GPU 利用率' },
  { number: '03', title: '安全可靠', desc: 'TLS + RBAC 权限' },
  { number: '04', title: '实时监控', desc: '资源使用率监控' }
])
</script>

<style scoped lang="scss">
.dashboard-container {
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
    animation: pulse 15s ease-in-out infinite;
  }

  @keyframes pulse {
    0%, 100% {
      transform: translate(0, 0) scale(1);
    }
    50% {
      transform: translate(-10%, -10%) scale(1.05);
    }
  }

  .project-title {
    font-size: 42px;
    font-weight: bold;
    margin: 0 0 10px 0;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
    position: relative;
    z-index: 1;
  }

  .project-subtitle {
    font-size: 18px;
    margin: 0 0 20px 0;
    opacity: 0.95;
    position: relative;
    z-index: 1;
  }

  .header-stats {
    display: flex;
    justify-content: center;
    gap: 30px;
    flex-wrap: wrap;
    position: relative;
    z-index: 1;

    .stat-item {
      display: flex;
      align-items: center;
      gap: 8px;
      font-size: 14px;
      opacity: 0.9;
    }
  }
}

.quick-nav {
  margin-bottom: 20px;

  .nav-card {
    background: white;
    border-radius: 12px;
    padding: 20px;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
    height: 100%;
    margin-bottom: 20px;
    position: relative;
    overflow: hidden;

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      opacity: 0;
      transition: opacity 0.3s ease;
      z-index: 0;
    }

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);

      &::before {
        opacity: 1;
      }

      .nav-icon {
        transform: scale(1.1);
        position: relative;
        z-index: 1;
      }

      .nav-title,
      .nav-desc {
        color: white;
        position: relative;
        z-index: 1;
      }
    }

    .nav-icon {
      font-size: 48px;
      margin-bottom: 10px;
      transition: transform 0.3s ease;
      position: relative;
      z-index: 1;
    }

    .nav-title {
      font-size: 16px;
      font-weight: 600;
      color: #303133;
      margin-bottom: 5px;
      transition: color 0.3s ease;
    }

    .nav-desc {
      font-size: 13px;
      color: #909399;
      transition: color 0.3s ease;
    }

    .nav-badge {
      position: absolute;
      top: 10px;
      right: 10px;
      background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
      color: white;
      font-size: 11px;
      padding: 2px 8px;
      border-radius: 10px;
      font-weight: 600;
    }
  }
}

.intro-card,
.scenario-card,
.highlight-card,
.status-card {
  margin-bottom: 20px;
  border-radius: 12px;

  :deep(.el-card__header) {
    background: linear-gradient(to right, #f6f8f9, #ffffff);
    border-bottom: 2px solid #e4e7ed;
    border-radius: 12px 12px 0 0;
  }

  .card-title {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.intro-card {
  .intro-content {
    .intro-text {
      line-height: 1.8;
      color: #606266;
      font-size: 15px;
      margin: 0 0 20px 0;
    }
  }

  .feature-item {
    display: flex;
    align-items: flex-start;
    gap: 15px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 8px;
    transition: all 0.3s;
    margin-bottom: 15px;

    &:hover {
      background: linear-gradient(135deg, #667eea15 0%, #764ba215 100%);
      transform: translateX(5px);
      box-shadow: 0 2px 8px rgba(102, 126, 234, 0.2);
    }

    .feature-icon {
      font-size: 36px;
      flex-shrink: 0;
    }

    .feature-text {
      flex: 1;

      h4 {
        margin: 0 0 8px 0;
        font-size: 15px;
        color: #303133;
      }

      p {
        margin: 0;
        font-size: 13px;
        color: #606266;
        line-height: 1.6;
      }
    }
  }
}

.scenario-card {
  .scenario-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 8px;
    transition: all 0.3s;
    margin-bottom: 15px;

    &:hover {
      background: white;
      box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
      transform: translateY(-2px);
    }

    .scenario-icon {
      font-size: 32px;
      flex-shrink: 0;
    }

    .scenario-content {
      flex: 1;

      h5 {
        margin: 0 0 4px 0;
        font-size: 14px;
        color: #303133;
        font-weight: 600;
      }

      p {
        margin: 0;
        font-size: 12px;
        color: #909399;
      }
    }
  }
}

.tech-card,
.guide-card {
  margin-bottom: 20px;
  height: 100%;

  :deep(.el-card__header) {
    background: linear-gradient(to right, #f6f8f9, #ffffff);
    border-bottom: 2px solid #e4e7ed;
    border-radius: 12px 12px 0 0;
  }

  .card-title {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }

  .tech-list {
    .tech-item {
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 12px 0;
      border-bottom: 1px solid #f0f0f0;

      &:last-child {
        border-bottom: none;
      }

      span {
        flex: 1;
        font-size: 14px;
        color: #606266;
      }
    }
  }

  .guide-text {
    font-size: 14px;
    color: #606266;
    margin: 5px 0 10px 0;
    line-height: 1.6;
  }
}

.highlight-card {
  .highlight-item {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 20px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 12px;
    color: white;
    margin-bottom: 15px;
    transition: all 0.3s;

    &:hover {
      transform: translateY(-3px);
      box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
    }

    .highlight-number {
      font-size: 32px;
      font-weight: bold;
      opacity: 0.8;
      flex-shrink: 0;
    }

    .highlight-content {
      flex: 1;

      h4 {
        margin: 0 0 5px 0;
        font-size: 16px;
      }

      p {
        margin: 0;
        font-size: 13px;
        opacity: 0.9;
      }
    }
  }
}

.status-card {
  .status-item {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 8px;
    transition: all 0.3s;

    &:hover {
      background: white;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    }

    .status-icon {
      font-size: 32px;
      flex-shrink: 0;
    }

    .status-info {
      flex: 1;

      .status-label {
        font-size: 12px;
        color: #909399;
        margin-bottom: 4px;
      }

      .status-value {
        font-size: 16px;
        font-weight: 600;
        color: #303133;
      }
    }
  }
}

.footer {
  text-align: center;
  padding: 30px 20px;
  color: #909399;
  font-size: 14px;

  p {
    margin: 5px 0;
  }

  a {
    color: #409eff;
    text-decoration: none;
    margin: 0 5px;

    &:hover {
      text-decoration: underline;
    }
  }

  .footer-links {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 8px;
    flex-wrap: wrap;
  }
}
</style>
