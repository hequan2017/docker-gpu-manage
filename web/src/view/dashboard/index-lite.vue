<template>
  <div class="dashboard-container">
    <!-- 项目标题 -->
    <div class="project-header">
      <h1 class="project-title">天启算力管理平台</h1>
      <p class="project-subtitle">Docker GPU 算力资源管理平台</p>
    </div>

    <!-- 项目介绍 -->
    <el-card class="section-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><Document /></el-icon>
          <span>📖 项目介绍</span>
        </div>
      </template>
      <div class="content-text">
        <p>
          <strong>Docker GPU 算力资源管理平台</strong> 是一个企业级的 GPU 容器化资源管理和调度系统，旨在帮助组织高效、安全地管理和分配 GPU 算力资源。平台采用现代化的微服务架构，提供从资源管理到容器实例全生命周期的完整解决方案。
        </p>
        <p class="mt-3">
          <el-tag type="success" size="large">✅ 页面加载正常</el-tag>
          <el-tag type="info" size="large" class="ml-2">📊 响应式布局</el-tag>
          <el-tag type="warning" size="large" class="ml-2">🚀 性能优化版</el-tag>
        </p>
      </div>
    </el-card>

    <!-- 快速导航 -->
    <el-row :gutter="20" class="mt-4">
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="quick-nav-card">
          <div class="quick-nav-item" @click="goToPage('instance')">
            <div class="icon">🐳</div>
            <div class="text">实例管理</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="quick-nav-card">
          <div class="quick-nav-item" @click="goToPage('computeNode')">
            <div class="icon">🖥️</div>
            <div class="text">算力节点</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="quick-nav-card">
          <div class="quick-nav-item" @click="goToPage('imageRegistry')">
            <div class="icon">📦</div>
            <div class="text">镜像库</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card shadow="hover" class="quick-nav-card">
          <div class="quick-nav-item" @click="goToPage('portForward')">
            <div class="icon">🔁</div>
            <div class="text">端口转发</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 核心功能 -->
    <el-card class="section-card mt-4" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><Star /></el-icon>
          <span>🌟 核心功能</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :xs="24" :sm="12" :md="8" v-for="feature in features" :key="feature.title">
          <div class="feature-item-simple">
            <div class="feature-icon">{{ feature.icon }}</div>
            <div class="feature-content">
              <h4>{{ feature.title }}</h4>
              <p>{{ feature.desc }}</p>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!-- 技术栈 -->
    <el-card class="section-card mt-4" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><Cpu /></el-icon>
          <span>💻 技术栈</span>
        </div>
      </template>
      <el-tabs v-model="activeTab">
        <el-tab-pane label="后端" name="backend">
          <div class="tech-list-simple">
            <el-tag v-for="tech in backendTech" :key="tech.name" class="tech-tag" type="primary">
              {{ tech.name }}
            </el-tag>
          </div>
        </el-tab-pane>
        <el-tab-pane label="前端" name="frontend">
          <div class="tech-list-simple">
            <el-tag v-for="tech in frontendTech" :key="tech.name" class="tech-tag" type="success">
              {{ tech.name }}
            </el-tag>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>

    <!-- 页脚 -->
    <div class="footer">
      <p>© 2025 天启算力管理平台 | 版本 v2.0</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  Document,
  Star,
  Cpu
} from '@element-plus/icons-vue'

defineOptions({
  name: 'Dashboard'
})

const router = useRouter()

// 技术栈Tab
const activeTab = ref('backend')

// 核心功能
const features = ref([
  { icon: '🧠', title: '智能资源匹配', desc: '根据GPU需求智能匹配节点' },
  { icon: '✂️', title: '显存切分', desc: '支持HAMi显存切分技术' },
  { icon: '🖥️', title: '多节点管理', desc: '统一管理多个算力节点' },
  { icon: '🔐', title: 'SSH跳板机', desc: '安全的SSH连接方式' },
  { icon: '💻', title: 'Web终端', desc: '浏览器中直接操作容器' },
  { icon: '📊', title: '实时监控', desc: 'CPU/内存/网络实时监控' }
])

// 技术栈
const backendTech = ref(['Gin', 'GORM', 'Go 1.23+', 'Docker API', 'Zap'])
const frontendTech = ref(['Vue 3', 'Element Plus', 'Vite', 'Pinia', 'Vue Router'])

// 跳转页面
const goToPage = (path) => {
  router.push({ path: `/${path}` })
}
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

.quick-nav-card {
  cursor: pointer;
  transition: all 0.3s;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
  }

  .quick-nav-item {
    text-align: center;
    padding: 20px;

    .icon {
      font-size: 48px;
      margin-bottom: 10px;
    }

    .text {
      font-size: 16px;
      font-weight: 600;
      color: #303133;
    }
  }
}

.feature-item-simple {
  display: flex;
  gap: 15px;
  padding: 20px;
  background: linear-gradient(135deg, #f6f8f9 0%, #ffffff 100%);
  border-radius: 8px;
  border: 1px solid #e4e7ed;
  margin-bottom: 15px;
  transition: all 0.3s;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    transform: translateX(5px);
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

.tech-list-simple {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;

  .tech-tag {
    font-size: 15px;
    padding: 12px 20px;
  }
}

.footer {
  text-align: center;
  padding: 30px 20px;
  color: #909399;
  font-size: 14px;

  p {
    margin: 0;
  }
}

.ml-2 {
  margin-left: 10px;
}

.ml-3 {
  margin-left: 15px;
}

.mt-3 {
  margin-top: 15px;
}

.mt-4 {
  margin-top: 20px;
}
</style>
