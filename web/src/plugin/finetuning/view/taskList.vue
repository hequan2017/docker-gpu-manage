<template>
  <div>
    <!-- 统计信息卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="总任务数" :value="stats.total" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="运行中" :value="stats.running">
            <template #suffix>
              <el-icon class="running-icon"><Loading /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="已完成" :value="stats.completed">
            <template #suffix>
              <el-icon class="success-icon"><CircleCheck /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="失败" :value="stats.failed">
            <template #suffix>
              <el-icon class="error-icon"><CircleClose /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
    </el-row>

    <!-- 搜索表单 -->
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="任务名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="请输入任务名称" clearable />
        </el-form-item>
        <el-form-item label="任务状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择任务状态"
            clearable
          >
            <el-option label="待执行" value="pending" />
            <el-option label="执行中" value="running" />
            <el-option label="已完成" value="completed" />
            <el-option label="失败" value="failed" />
            <el-option label="已停止" value="stopped" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格区域 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          新建任务
        </el-button>
        <el-button
          icon="delete"
          :disabled="!multipleSelection.length"
          @click="batchDelete"
        >
          批量删除
        </el-button>
        <el-button icon="refresh" @click="getTableData">
          刷新
        </el-button>
      </el-table>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column
          align="left"
          label="任务名称"
          prop="name"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="基础模型"
          prop="baseModel"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag
              :type="getStatusType(scope.row.status)"
              effect="dark"
            >
              {{ getStatusLabel(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="进度"
          prop="progress"
          width="150"
        >
          <template #default="scope">
            <el-progress
              :percentage="Math.round(scope.row.progress || 0)"
              :status="getProgressStatus(scope.row.status)"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="创建时间"
          prop="CreatedAt"
          width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="280"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="view"
              @click="viewDetail(scope.row)"
            >
              详情
            </el-button>
            <el-button
              v-if="scope.row.status === 'running'"
              type="warning"
              link
              icon="video-pause"
              @click="stopTask(scope.row)"
            >
              停止
            </el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteTask(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 创建任务对话框 -->
    <el-dialog
      v-model="dialogFormVisible"
      title="创建微调任务"
      width="900px"
      :close-on-click-modal="false"
      @close="closeDialog"
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-width="120px"
        :rules="rules"
      >
        <el-tabs v-model="activeTab">
          <!-- 基本信息 -->
          <el-tab-pane label="基本信息" name="basic">
            <el-form-item label="任务名称" prop="name">
              <el-input
                v-model="formData.name"
                placeholder="请输入任务名称"
                clearable
              />
            </el-form-item>
            <el-form-item label="任务描述" prop="description">
              <el-input
                v-model="formData.description"
                type="textarea"
                :rows="3"
                placeholder="请输入任务描述"
              />
            </el-form-item>
            <el-form-item label="基础模型" prop="baseModel">
              <el-input
                v-model="formData.baseModel"
                placeholder="例如: /path/to/model 或 llama2-7b"
                clearable
              />
              <div class="form-tip">
                支持本地路径或HuggingFace模型名称（如: meta-llama/Llama-2-7b）
              </div>
            </el-form-item>
            <el-form-item label="数据集路径" prop="datasetPath">
              <el-input
                v-model="formData.datasetPath"
                placeholder="请输入数据集路径"
                clearable
              />
            </el-form-item>
            <el-form-item label="输出路径">
              <el-input
                v-model="formData.outputPath"
                placeholder="留空则使用默认路径"
                clearable
              />
            </el-form-item>
          </el-tab-pane>

          <!-- GPU配置 -->
          <el-tab-pane label="GPU配置" name="gpu">
            <el-form-item label="CUDA设备">
              <el-input
                v-model="formData.gpuConfig.cuda_visible_devices"
                placeholder="例如: 0,1 或 0-2"
              />
              <div class="form-tip">
                指定使用的GPU设备，多个设备用逗号分隔
              </div>
            </el-form-item>
          </el-tab-pane>

          <!-- 训练参数 -->
          <el-tab-pane label="训练参数" name="training">
            <el-form-item label="配置预设">
              <el-select
                v-model="selectedPreset"
                placeholder="选择预设配置"
                @change="applyPreset"
              >
                <el-option label="快速测试" value="quick_test" />
                <el-option label="标准训练" value="standard" />
                <el-option label="完整微调" value="full_finetune" />
                <el-option label="自定义" value="custom" />
              </el-select>
            </el-form-item>
            <el-form-item label="学习率">
              <el-input-number
                v-model="formData.trainingArgs.learning_rate"
                :min="0"
                :max="1"
                :step="0.00001"
                :precision="5"
                :controls="true"
              />
              <div class="form-tip">
                常用值: 1e-4 到 1e-5
              </div>
            </el-form-item>
            <el-form-item label="批次大小">
              <el-input-number
                v-model="formData.trainingArgs.batch_size"
                :min="1"
                :max="1024"
                :step="1"
              />
              <div class="form-tip">
                根据GPU显存调整，常用值: 16, 32, 64
              </div>
            </el-form-item>
            <el-form-item label="训练轮数">
              <el-input-number
                v-model="formData.trainingArgs.num_train_epochs"
                :min="1"
                :max="100"
              />
            </el-form-item>
            <el-form-item label="最大步数">
              <el-input-number
                v-model="formData.trainingArgs.max_steps"
                :min="-1"
              />
              <div class="form-tip">
                -1 表示不限制步数，按轮数训练
              </div>
            </el-form-item>
            <el-form-item label="预热步数">
              <el-input-number
                v-model="formData.trainingArgs.warmup_steps"
                :min="0"
              />
            </el-form-item>
            <el-form-item label="日志间隔">
              <el-input-number
                v-model="formData.trainingArgs.logging_steps"
                :min="1"
              />
            </el-form-item>
            <el-form-item label="保存间隔">
              <el-input-number
                v-model="formData.trainingArgs.save_steps"
                :min="1"
              />
            </el-form-item>
          </el-tab-pane>

          <!-- 高级设置 -->
          <el-tab-pane label="高级设置" name="advanced">
            <el-form-item label="自定义命令">
              <el-input
                v-model="formData.command"
                type="textarea"
                :rows="6"
                placeholder="留空则使用默认命令模板，输入自定义命令将覆盖所有参数配置"
              />
              <div class="form-tip warning">
                ⚠️ 自定义命令将忽略上述所有参数配置
              </div>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
      <template #footer>
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitting" @click="enterDialog">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createFinetuningTask,
  deleteFinetuningTask,
  stopFinetuningTask,
  getFinetuningTaskList
} from '@/plugin/finetuning/api/task'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Loading, CircleCheck, CircleClose } from '@element-plus/icons-vue'

defineOptions({
  name: 'FinetuningTaskList'
})

const router = useRouter()

// 训练预设配置
const presets = {
  quick_test: {
    learning_rate: 0.0001,
    batch_size: 16,
    num_train_epochs: 1,
    max_steps: 100,
    warmup_steps: 10,
    logging_steps: 10,
    save_steps: 50
  },
  standard: {
    learning_rate: 0.0002,
    batch_size: 32,
    num_train_epochs: 3,
    max_steps: -1,
    warmup_steps: 100,
    logging_steps: 50,
    save_steps: 500
  },
  full_finetune: {
    learning_rate: 0.00005,
    batch_size: 64,
    num_train_epochs: 10,
    max_steps: -1,
    warmup_steps: 500,
    logging_steps: 100,
    save_steps: 1000
  }
}

// 统计信息
const stats = ref({
  total: 0,
  running: 0,
  completed: 0,
  failed: 0
})

// 搜索条件
const searchInfo = ref({
  name: '',
  status: '',
  baseModel: ''
})

// 表格数据
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const multipleSelection = ref([])

// 对话框
const dialogFormVisible = ref(false)
const submitting = ref(false)
const activeTab = ref('basic')
const selectedPreset = ref('standard')
const formData = ref({
  name: '',
  description: '',
  baseModel: '',
  datasetPath: '',
  outputPath: '',
  trainingArgs: {
    learning_rate: 0.0002,
    batch_size: 32,
    num_train_epochs: 3,
    max_steps: -1,
    warmup_steps: 100,
    logging_steps: 50,
    save_steps: 500
  },
  gpuConfig: {
    cuda_visible_devices: '0'
  },
  command: ''
})

// 表单验证规则
const rules = reactive({
  name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
  baseModel: [{ required: true, message: '请输入基础模型', trigger: 'blur' }],
  datasetPath: [
    { required: true, message: '请输入数据集路径', trigger: 'blur' }
  ]
})

// 应用预设配置
const applyPreset = (preset) => {
  if (preset !== 'custom' && presets[preset]) {
    formData.value.trainingArgs = { ...presets[preset] }
  }
}

// 获取任务列表
const getTableData = async () => {
  const res = await getFinetuningTaskList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    updateStats(res.data.list)
  }
}

// 更新统计信息
const updateStats = (list) => {
  stats.value = {
    total: list.length,
    running: list.filter(t => t.status === 'running').length,
    completed: list.filter(t => t.status === 'completed').length,
    failed: list.filter(t => t.status === 'failed').length
  }
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {
    name: '',
    status: '',
    baseModel: ''
  }
  getTableData()
}

// 分页
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 批量删除
const batchDelete = async () => {
  await ElMessageBox.confirm(
    `确定要删除选中的 ${multipleSelection.value.length} 个任务吗?`,
    '批量删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )

  const promises = multipleSelection.value.map(item =>
    deleteFinetuningTask({ id: item.ID })
  )

  await Promise.all(promises)
  ElMessage.success('批量删除成功')
  getTableData()
}

// 打开对话框
const openDialog = () => {
  dialogFormVisible.value = true
  applyPreset(selectedPreset.value)
}

// 关闭对话框
const closeDialog = () => {
  dialogFormVisible.value = false
  activeTab.value = 'basic'
  formData.value = {
    name: '',
    description: '',
    baseModel: '',
    datasetPath: '',
    outputPath: '',
    trainingArgs: {
      learning_rate: 0.0002,
      batch_size: 32,
    num_train_epochs: 3,
      max_steps: -1,
      warmup_steps: 100,
      logging_steps: 50,
      save_steps: 500
    },
    gpuConfig: {
      cuda_visible_devices: '0'
    },
    command: ''
  }
}

// 创建任务
const enterDialog = async () => {
  submitting.value = true
  try {
    const res = await createFinetuningTask(formData.value)
    if (res.code === 0) {
      ElMessage.success('任务创建成功')
      closeDialog()
      getTableData()
    }
  } finally {
    submitting.value = false
  }
}

// 停止任务
const stopTask = async (row) => {
  await ElMessageBox.confirm('确定要停止该任务吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  const res = await stopFinetuningTask({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('任务已停止')
    getTableData()
  }
}

// 删除任务
const deleteTask = async (row) => {
  await ElMessageBox.confirm('确定要删除该任务吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
  const res = await deleteFinetuningTask({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

// 查看详情
const viewDetail = (row) => {
  router.push({
    name: 'finetuningTaskDetail',
    params: { id: row.ID }
  })
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
  if (!date) return ''
  return new Date(date).toLocaleString('zh-CN')
}

// 自动刷新
let refreshTimer = null

const startAutoRefresh = () => {
  refreshTimer = setInterval(() => {
    getTableData()
  }, 10000) // 每10秒刷新一次
}

const stopAutoRefresh = () => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
    refreshTimer = null
  }
}

// 初始化
onMounted(() => {
  getTableData()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<style scoped lang="scss">
.stats-row {
  margin-bottom: 20px;

  .running-icon {
    color: #409eff;
    animation: rotate 2s linear infinite;
  }

  .success-icon {
    color: #67c23a;
  }

  .error-icon {
    color: #f56c6c;
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.gva-search-box {
  padding: 20px;
  background: #fff;
  margin-bottom: 20px;
  border-radius: 4px;
}

.gva-table-box {
  padding: 20px;
  background: #fff;
  border-radius: 4px;
}

.gva-btn-list {
  margin-bottom: 20px;
  display: flex;
  gap: 10px;
}

.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;

  &.warning {
    color: #e6a23c;
  }
}

:deep(.el-form-item__content) {
  flex-direction: column;
  align-items: flex-start;
}

:deep(.el-input-number) {
  width: 100%;
}
</style>
