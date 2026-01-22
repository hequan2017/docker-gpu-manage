<template>
  <div class="dashboard-container">
    <!-- 项目标题 -->
    <div class="project-header">
      <h1 class="project-title">天启算力管理平台</h1>
      <p class="project-subtitle">Docker GPU 算力资源管理平台</p>
    </div>

    <!-- 快速导航 -->
    <el-row :gutter="20">
      <el-col :xs="12" :sm="8" :md="6" :lg="4" v-for="nav in quickNavs" :key="nav.name">
        <div class="nav-card" @click="goTo(nav.path)">
          <div class="nav-icon">{{ nav.icon }}</div>
          <div class="nav-title">{{ nav.name }}</div>
          <div class="nav-desc">{{ nav.desc }}</div>
        </div>
      </el-col>
    </el-row>

    <!-- 核心功能 -->
    <el-card class="feature-card" shadow="hover">
      <template #header>
        <span class="card-title">🌟 核心功能</span>
      </template>
      <el-row :gutter="15">
        <el-col :xs="24" :sm="12" :md="8" v-for="feature in features" :key="feature.title">
          <div class="feature-item">
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
    <el-row :gutter="20">
      <el-col :xs="24" :md="12">
        <el-card class="tech-card" shadow="hover">
          <template #header>
            <span class="card-title">💻 后端技术</span>
          </template>
          <div class="tech-tags">
            <el-tag type="primary" v-for="tech in backendTech" :key="tech">{{ tech }}</el-tag>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :md="12">
        <el-card class="tech-card" shadow="hover">
          <template #header>
            <span class="card-title">🎨 前端技术</span>
          </template>
          <div class="tech-tags">
            <el-tag type="success" v-for="tech in frontendTech" :key="tech">{{ tech }}</el-tag>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'Dashboard'
})

const router = useRouter()

// 快速导航
const quickNavs = ref([
  { name: '实例管理', icon: '🐳', desc: 'GPU容器实例', path: 'instance' },
  { name: '算力节点', icon: '🖥️', desc: '节点管理', path: 'computeNode' },
  { name: '镜像库', icon: '📦', desc: '镜像管理', path: 'imageRegistry' },
  { name: '产品规格', icon: '💎', desc: '规格配置', path: 'productSpec' },
  { name: '端口转发', icon: '🔁', desc: '转发规则', path: 'portForward' },
  { name: 'K8s集群', icon: '☸️', desc: '集群管理', path: 'k8sCluster' }
])

// 核心功能
const features = ref([
  { icon: '🧠', title: '智能资源匹配', desc: '根据GPU需求智能匹配最优节点' },
  { icon: '✂️', title: '显存切分', desc: '支持HAMi显存虚拟化技术' },
  { icon: '🖥️', title: '多节点管理', desc: '统一管理多个分布式算力节点' },
  { icon: '🔐', title: 'SSH跳板机', desc: '安全的SSH连接访问容器' },
  { icon: '💻', title: 'Web终端', desc: '浏览器中直接操作容器' },
  { icon: '📊', title: '实时监控', desc: 'CPU/内存/网络实时监控' }
])

// 技术栈
const backendTech = ref(['Go 1.23+', 'Gin', 'GORM', 'Docker API', 'Zap'])
const frontendTech = ref(['Vue 3', 'Element Plus', 'Vite', 'Pinia'])

// 跳转
const goTo = (path) => {
  router.push({ name: path })
}
</script>

<style scoped lang="scss">
.dashboard-container {
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
  background: #f5f7fa;
  min-height: calc(100vh - 100px);
}

.project-header {
  text-align: center;
  margin-bottom: 30px;
  padding: 40px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  color: white;

  .project-title {
    font-size: 42px;
    font-weight: bold;
    margin: 0 0 10px 0;
  }

  .project-subtitle {
    font-size: 18px;
    margin: 0;
    opacity: 0.95;
  }
}

.nav-card {
  background: white;
  border-radius: 12px;
  padding: 20px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  margin-bottom: 20px;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);

    .nav-icon,
    .nav-title,
    .nav-desc {
      color: white;
    }
  }

  .nav-icon {
    font-size: 48px;
    margin-bottom: 10px;
  }

  .nav-title {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
    margin-bottom: 5px;
  }

  .nav-desc {
    font-size: 13px;
    color: #909399;
  }
}

.feature-card {
  margin-bottom: 20px;
  border-radius: 12px;

  :deep(.el-card__header) {
    background: linear-gradient(to right, #f6f8f9, #ffffff);
  }

  .card-title {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.feature-item {
  display: flex;
  gap: 12px;
  padding: 15px;
  background: #f5f7fa;
  border-radius: 8px;
  margin-bottom: 12px;
  transition: all 0.3s;

  &:hover {
    background: linear-gradient(135deg, #667eea15 0%, #764ba215 100%);
    transform: translateX(5px);
  }

  .feature-icon {
    font-size: 36px;
    flex-shrink: 0;
  }

  .feature-content {
    flex: 1;

    h4 {
      margin: 0 0 6px 0;
      font-size: 15px;
      color: #303133;
    }

    p {
      margin: 0;
      font-size: 13px;
      color: #606266;
      line-height: 1.5;
    }
  }
}

.tech-card {
  margin-bottom: 20px;
  border-radius: 12px;

  :deep(.el-card__header) {
    background: linear-gradient(to right, #f6f8f9, #ffffff);
  }

  .card-title {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;

  .el-tag {
    font-size: 14px;
    padding: 8px 15px;
  }
}
</style>
