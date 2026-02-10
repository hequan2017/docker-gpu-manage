<template>
  <div class="pcdn-page">
    <el-card shadow="never">
      <template #header>调度任务流水</template>
      <el-table :data="tableData" v-loading="loading" border>
        <el-table-column prop="taskId" label="任务ID" min-width="180" />
        <el-table-column prop="startTime" label="开始时间" min-width="180" />
        <el-table-column prop="status" label="状态" min-width="100">
          <template #default="scope">
            <el-tag :type="statusTypeMap[scope.row.status] || 'info'">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reason" label="失败原因" min-width="220" />
        <el-table-column prop="retryStatus" label="重试状态" min-width="120" />
        <el-table-column label="操作" min-width="120" fixed="right">
          <template #default="scope">
            <el-button size="small" type="primary" :disabled="scope.row.status !== '失败'" @click="retry(scope.row)">重试</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getPcdnDispatchTaskList, retryPcdnDispatchTask } from '@/api/pcdn'

defineOptions({ name: 'PcdnDispatchTasks' })

const loading = ref(false)
const tableData = ref([])

const statusTypeMap = {
  成功: 'success',
  失败: 'danger',
  进行中: 'warning'
}

const mockData = [
  { taskId: 'task-20260210-001', startTime: '2026-02-10 10:20:10', status: '成功', reason: '-', retryStatus: '无需重试' },
  { taskId: 'task-20260210-002', startTime: '2026-02-10 10:21:36', status: '失败', reason: '节点连接超时', retryStatus: '待重试' },
  { taskId: 'task-20260210-003', startTime: '2026-02-10 10:22:01', status: '进行中', reason: '-', retryStatus: '处理中' }
]

const loadData = async () => {
  loading.value = true
  try {
    const res = await getPcdnDispatchTaskList()
    tableData.value = res?.data?.list || mockData
  } catch (e) {
    tableData.value = mockData
  } finally {
    loading.value = false
  }
}

const retry = async (row) => {
  await retryPcdnDispatchTask({ taskId: row.taskId })
  ElMessage.success('重试任务已提交')
  loadData()
}

onMounted(loadData)
</script>

<style scoped lang="scss">
.pcdn-page {
  padding: 20px;
}
</style>
