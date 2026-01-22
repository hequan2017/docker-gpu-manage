<template>
  <div class="task-detail">
    <el-page-header @back="goBack" content="任务详情" />

    <el-card v-loading="loading" class="detail-card">
      <template #header>
        <div class="card-header">
          <span>{{ taskData.name || '加载中...' }}</span>
          <el-tag :type="getStatusType(taskData.status)" effect="dark">
            {{ getStatusLabel(taskData.status) }}
          </el-tag>
        </div>
      </template>

      <el-descriptions :column="2" border>
        <el-descriptions-item label="任务ID">
          {{ taskData.ID }}
        </el-descriptions-item>
        <el-descriptions-item label="任务名称">
          {{ taskData.name }}
        </el-descriptions-item>
        <el-descriptions-item label="基础模型">
          {{ taskData.baseModel }}
        </el-descriptions-item>
        <el-descriptions-item label="数据集路径">
          {{ taskData.datasetPath }}
        </el-descriptions-item>
        <el-descriptions-item label="输出路径">
          {{ taskData.outputPath || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="进度">
          <el-progress
            :percentage="taskData.progress || 0"
            :status="getProgressStatus(taskData.status)"
          />
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDate(taskData.CreatedAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="开始时间">
          {{ formatDate(taskData.startedAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="结束时间">
          {{ formatDate(taskData.finishedAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="进程ID">
          {{ taskData.pid || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="执行命令" :span="2">
          <el-input
            v-model="taskData.command"
            type="textarea"
            :rows="3"
            readonly
          />
        </el-descriptions-item>
        <el-descriptions-item v-if="taskData.errorMessage" label="错误信息" :span="2">
          <el-alert :title="taskData.errorMessage" type="error" :closable="false" />
        </el-descriptions-item>
      </el-descriptions>

      <div class="action-buttons">
        <el-button
          v-if="taskData.status === 'running'"
          type="warning"
          icon="video-pause"
          @click="stopTask"
        >
          停止任务
        </el-button>
        <el-button icon="refresh" @click="refresh"> 刷新 </el-button>
        <el-button type="danger" icon="delete" @click="deleteTask">
          删除任务
        </el-button>
      </div>
    </el-card>

    <!-- 训练日志 -->
    <el-card class="log-card">
      <template #header>
        <div class="card-header">
          <span>训练日志</span>
          <div class="header-actions">
            <el-checkbox v-model="autoScroll">自动滚动</el-checkbox>
            <el-button
              size="small"
              icon="refresh"
              :loading="logLoading"
              @click="fetchLog"
            >
              刷新日志
            </el-button>
          </div>
        </div>
      </template>
      <div ref="logContainer" class="log-container">
        <pre v-if="logContent" class="log-content">{{ logContent }}</pre>
        <el-empty v-else description="暂无日志" />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import {
  deleteFinetuningTask,
  getFinetuningTask,
  getFinetuningTaskLog,
  stopFinetuningTask
} from '@/plugin/finetuning/api/task'
import { ElMessage, ElMessageBox } from 'element-plus'
import { onMounted, onUnmounted, ref, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'

defineOptions({
  name: 'FinetuningTaskDetail'
})

const route = useRoute()
const router = useRouter()

const taskId = ref(route.params.id || route.query.id)
const loading = ref(false)
const logLoading = ref(false)
const taskData = ref({})
const logContent = ref('')
const logContainer = ref(null)
const autoScroll = ref(true)

let refreshTimer = null

// 获取任务详情
const fetchTaskDetail = async () => {
  loading.value = true
  try {
    const res = await getFinetuningTask({ id: taskId.value })
    if (res.code === 0) {
      taskData.value = res.data
    }
  } finally {
    loading.value = false
  }
}

// 获取日志
const fetchLog = async () => {
  logLoading.value = true
  try {
    const res = await getFinetuningTaskLog({
      id: taskId.value,
      lines: 1000
    })
    if (res.code === 0) {
      logContent.value = res.data || ''
      if (autoScroll.value) {
        await nextTick()
        scrollToBottom()
      }
    }
  } finally {
    logLoading.value = false
  }
}

// 滚动到底部
const scrollToBottom = () => {
  if (logContainer.value) {
    logContainer.value.scrollTop = logContainer.value.scrollHeight
  }
}

// 刷新
const refresh = () => {
  fetchTaskDetail()
  fetchLog()
}

// 停止任务
const stopTask = async () => {
  await ElMessageBox.confirm('确定要停止该任务吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  const res = await stopFinetuningTask({ id: taskId.value })
  if (res.code === 0) {
    ElMessage.success('任务已停止')
    fetchTaskDetail()
  }
}

// 删除任务
const deleteTask = async () => {
  await ElMessageBox.confirm('确定要删除该任务吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  const res = await deleteFinetuningTask({ id: taskId.value })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    goBack()
  }
}

// 返回
const goBack = () => {
  router.back()
}

// 状态相关方法
const getStatusType = (status) => {
  const typeMap = {
    pending: 'info',
    running: 'primary',
    completed: 'success',
    failed: 'danger',
    stopped: 'warning'
  }
  return typeMap[status] || 'info'
}

const getStatusLabel = (status) => {
  const labelMap = {
    pending: '待执行',
    running: '执行中',
    completed: '已完成',
    failed: '失败',
    stopped: '已停止'
  }
  return labelMap[status] || status
}

const getProgressStatus = (status) => {
  if (status === 'completed') return 'success'
  if (status === 'failed') return 'exception'
  return null
}

// 格式化日期
const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleString('zh-CN')
}

// 启动自动刷新
const startAutoRefresh = () => {
  // 任务运行中时每5秒刷新一次
  refreshTimer = setInterval(() => {
    if (taskData.value.status === 'running') {
      fetchTaskDetail()
      fetchLog()
    } else if (taskData.value.status === 'pending') {
      fetchTaskDetail()
    }
  }, 5000)
}

// 停止自动刷新
const stopAutoRefresh = () => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
    refreshTimer = null
  }
}

// 初始化
onMounted(async () => {
  await fetchTaskDetail()
  await fetchLog()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<style scoped lang="scss">
.task-detail {
  padding: 20px;
}

.detail-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .header-actions {
    display: flex;
    align-items: center;
    gap: 12px;
  }
}

.action-buttons {
  margin-top: 20px;
  display: flex;
  gap: 12px;
}

.log-card {
  .log-container {
    height: 500px;
    overflow-y: auto;
    background: #1e1e1e;
    border-radius: 4px;
    padding: 16px;
  }

  .log-content {
    margin: 0;
    color: #d4d4d4;
    font-family: 'Courier New', monospace;
    font-size: 13px;
    line-height: 1.6;
    white-space: pre-wrap;
    word-wrap: break-word;
  }
}
</style>
