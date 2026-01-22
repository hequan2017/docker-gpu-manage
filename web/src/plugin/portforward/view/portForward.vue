<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="源IP地址" prop="sourceIP">
          <el-input
            v-model="searchInfo.sourceIP"
            placeholder="请输入源IP地址"
            clearable
          />
        </el-form-item>
        <el-form-item label="协议类型" prop="protocol">
          <el-select
            v-model="searchInfo.protocol"
            placeholder="请选择协议类型"
            clearable
          >
            <el-option label="TCP" value="tcp" />
            <el-option label="UDP" value="udp" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择状态"
            clearable
          >
            <el-option label="启用" :value="true" />
            <el-option label="禁用" :value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
          <el-button
            icon="refresh"
            @click="toggleAutoRefresh"
            :type="autoRefresh ? 'success' : ''"
          >
            {{ autoRefresh ? '停止自动刷新' : '自动刷新' }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 状态统计卡片 -->
    <div class="gva-card-box" style="padding: 0 20px; margin-bottom: 20px;">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #409eff;">
                <el-icon><document /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ forwarderStats.total_rules || 0 }}</div>
                <div class="stat-label">总规则数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #67c23a;">
                <el-icon><success-filled /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ forwarderStats.running_forwarders || 0 }}</div>
                <div class="stat-label">运行中</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #e6a23c;">
                <el-icon><circle-check /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ forwarderStats.enabled_count || 0 }}</div>
                <div class="stat-label">已启用</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #909399;">
                <el-icon><connection /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ totalConnections }}</div>
                <div class="stat-label">活跃连接</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          新增
        </el-button>
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >
          删除
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        v-loading="loading"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="创建日期" prop="CreatedAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="源地址" width="180">
          <template #default="scope">
            <span class="address-text">{{ scope.row.sourceIP }}:{{ scope.row.sourcePort }}</span>
          </template>
        </el-table-column>

        <el-table-column align="left" label="协议" prop="protocol" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.protocol === 'tcp' ? 'primary' : 'success'">
              {{ scope.row.protocol.toUpperCase() }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="目标地址" width="180">
          <template #default="scope">
            <span class="address-text">{{ scope.row.targetIP }}:{{ scope.row.targetPort }}</span>
          </template>
        </el-table-column>

        <el-table-column align="left" label="运行状态" width="120">
          <template #default="scope">
            <div v-if="forwarderStatus[scope.row.ID] !== undefined">
              <el-tag v-if="forwarderStatus[scope.row.ID].running" type="success">
                运行中
              </el-tag>
              <el-tag v-else type="info">
                已停止
              </el-tag>
            </div>
            <el-tag v-else type="warning">
              未知
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="连接数" width="100">
          <template #default="scope">
            <span v-if="forwarderStatus[scope.row.ID] !== undefined">
              <el-tag v-if="forwarderStatus[scope.row.ID].running" type="primary">
                {{ forwarderStatus[scope.row.ID].conn_count || 0 }}
              </el-tag>
              <span v-else>-</span>
            </span>
            <span v-else>-</span>
          </template>
        </el-table-column>

        <el-table-column align="left" label="开关状态" prop="status" width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.status"
              @change="handleStatusChange(scope.row)"
            />
          </template>
        </el-table-column>

        <el-table-column align="left" label="描述" prop="description" min-width="150" show-overflow-tooltip />

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          width="280"
        >
          <template #default="scope">
            <el-button
              type="info"
              link
              icon="view"
              @click="viewStatus(scope.row)"
            >
              状态
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updatePortForwardFunc(scope.row)"
            >
              修改
            </el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
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

    <!-- 创建/编辑抽屉 -->
    <el-drawer
      v-model="dialogFormVisible"
      destroy-on-close
      size="600"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加端口转发规则' : '修改端口转发规则' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog"> 确 定 </el-button>
            <el-button @click="closeDialog"> 取 消 </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="top"
        :rules="rule"
        label-width="120px"
      >
        <el-form-item label="源IP地址:" prop="sourceIP">
          <el-input
            v-model="formData.sourceIP"
            :clearable="true"
            placeholder="请输入源IP地址，如: 0.0.0.0"
          />
        </el-form-item>
        <el-form-item label="源端口:" prop="sourcePort">
          <el-input-number
            v-model="formData.sourcePort"
            :min="1"
            :max="65535"
            placeholder="请输入源端口"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="协议类型:" prop="protocol">
          <el-select
            v-model="formData.protocol"
            placeholder="请选择协议类型"
            style="width: 100%"
          >
            <el-option label="TCP" value="tcp" />
            <el-option label="UDP" value="udp" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标IP地址:" prop="targetIP">
          <el-input
            v-model="formData.targetIP"
            :clearable="true"
            placeholder="请输入目标IP地址"
          />
        </el-form-item>
        <el-form-item label="目标端口:" prop="targetPort">
          <el-input-number
            v-model="formData.targetPort"
            :min="1"
            :max="65535"
            placeholder="请输入目标端口"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-switch
            v-model="formData.status"
            active-text="启用"
            inactive-text="禁用"
          />
        </el-form-item>
        <el-form-item label="规则描述:" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入规则描述"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 状态详情对话框 -->
    <el-dialog
      v-model="statusDialogVisible"
      title="端口转发状态详情"
      width="600px"
    >
      <div v-if="currentStatus">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="规则ID">
            {{ currentRow?.ID }}
          </el-descriptions-item>
          <el-descriptions-item label="运行状态">
            <el-tag v-if="currentStatus.running" type="success">运行中</el-tag>
            <el-tag v-else type="info">已停止</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="协议类型">
            <el-tag :type="currentRow?.protocol === 'tcp' ? 'primary' : 'success'">
              {{ currentStatus.protocol?.toUpperCase() || currentRow?.protocol?.toUpperCase() }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="活跃连接">
            <el-tag type="primary">{{ currentStatus.conn_count || 0 }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="源地址" :span="2">
            {{ currentRow?.sourceIP }}:{{ currentRow?.sourcePort }}
          </el-descriptions-item>
          <el-descriptions-item label="目标地址" :span="2">
            {{ currentRow?.targetIP }}:{{ currentRow?.targetPort }}
          </el-descriptions-item>
          <el-descriptions-item label="规则描述" :span="2">
            {{ currentRow?.description || '-' }}
          </el-descriptions-item>
        </el-descriptions>

        <div style="margin-top: 20px;">
          <h4>连接说明</h4>
          <el-alert
            v-if="currentStatus.running"
            title="端口转发正在运行中"
            type="success"
            :closable="false"
          >
            <template #default>
              <p>当前活跃连接数: <strong>{{ currentStatus.conn_count || 0 }}</strong></p>
              <p>可以通过访问 <code>{{ currentRow?.sourceIP }}:{{ currentRow?.sourcePort }}</code> 来连接到目标地址 <code>{{ currentRow?.targetIP }}:{{ currentRow?.targetPort }}</code></p>
            </template>
          </el-alert>
          <el-alert
            v-else
            title="端口转发未运行"
            type="warning"
            :closable="false"
          >
            <template #default>
              <p>请确保规则开关状态为"启用"以启动端口转发</p>
            </template>
          </el-alert>
        </div>
      </div>
      <template #footer>
        <el-button @click="statusDialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="refreshCurrentStatus">刷新状态</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createPortForward,
  deletePortForward,
  deletePortForwardByIds,
  updatePortForward,
  findPortForward,
  getPortForwardList,
  updatePortForwardStatus,
  getServerIP,
  getForwarderStatus,
  getAllForwarderStatus
} from '@/plugin/portforward/api/portForward'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Document,
  SuccessFilled,
  CircleCheck,
  Connection
} from '@element-plus/icons-vue'
import { ref, reactive, onMounted, onBeforeUnmount, computed } from 'vue'

defineOptions({
  name: 'PortForward'
})

// 验证规则
const rule = reactive({
  sourceIP: [
    { required: true, message: '请输入源IP地址', trigger: 'blur' },
    {
      pattern: /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/,
      message: '请输入正确的IP地址格式',
      trigger: 'blur'
    }
  ],
  sourcePort: [
    { required: true, message: '请输入源端口', trigger: 'blur' },
    { type: 'number', min: 1, max: 65535, message: '端口范围为1-65535', trigger: 'blur' }
  ],
  protocol: [
    { required: true, message: '请选择协议类型', trigger: 'change' }
  ],
  targetIP: [
    { required: true, message: '请输入目标IP地址', trigger: 'blur' },
    {
      pattern: /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/,
      message: '请输入正确的IP地址格式',
      trigger: 'blur'
    }
  ],
  targetPort: [
    { required: true, message: '请输入目标端口', trigger: 'blur' },
    { type: 'number', min: 1, max: 65535, message: '端口范围为1-65535', trigger: 'blur' }
  ]
})

const elFormRef = ref()
const elSearchFormRef = ref()
const loading = ref(false)

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 转发器状态
const forwarderStatus = ref({})
const forwarderStats = ref({
  total_rules: 0,
  running_forwarders: 0,
  enabled_count: 0,
  running_ids: []
})

// 计算总连接数
const totalConnections = computed(() => {
  return Object.values(forwarderStatus.value).reduce((sum, status) => {
    return sum + (status.conn_count || 0)
  }, 0)
})

// 自动刷新
const autoRefresh = ref(false)
const autoRefreshTimer = ref(null)

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 获取转发器状态
const getForwarderStatusData = async () => {
  try {
    const res = await getAllForwarderStatus()
    if (res.code === 0) {
      forwarderStats.value = res.data
      // 获取每个运行中转发器的详细状态
      if (res.data.running_ids && res.data.running_ids.length > 0) {
        for (const id of res.data.running_ids) {
          try {
            const statusRes = await getForwarderStatus({ ID: String(id) })
            if (statusRes.code === 0) {
              forwarderStatus.value[id] = statusRes.data
            }
          } catch (error) {
            console.error(`获取转发器 ${id} 状态失败:`, error)
          }
        }
      }
      // 清理已停止的转发器状态
      const runningIds = (res.data.running_ids || []).map(String)
      Object.keys(forwarderStatus.value).forEach(id => {
        if (!runningIds.includes(id)) {
          delete forwarderStatus.value[id]
        }
      })
    }
  } catch (error) {
    console.error('获取转发器状态失败:', error)
  }
}

// 查询
const getTableData = async () => {
  loading.value = true
  try {
    const table = await getPortForwardList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
      // 获取转发器状态
      await getForwarderStatusData()
    }
  } finally {
    loading.value = false
  }
}

// 切换自动刷新
const toggleAutoRefresh = () => {
  autoRefresh.value = !autoRefresh.value
  if (autoRefresh.value) {
    autoRefreshTimer.value = setInterval(getForwarderStatusData, 5000) // 每5秒刷新
    ElMessage.success('已开启自动刷新')
  } else {
    clearInterval(autoRefreshTimer.value)
    ElMessage.info('已停止自动刷新')
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deletePortForwardFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID)
      })
    const res = await deletePortForwardByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePortForwardFunc = async (row) => {
  const res = await findPortForward({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deletePortForwardFunc = async (row) => {
  const res = await deletePortForward({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = async () => {
  type.value = 'create'
  // 获取服务器IP并设置为默认值
  try {
    const res = await getServerIP()
    if (res.code === 0 && res.data.ips && res.data.ips.length > 0) {
      formData.value.sourceIP = res.data.ips[0]
    } else {
      formData.value.sourceIP = '0.0.0.0'
    }
  } catch (error) {
    console.error('获取服务器IP失败:', error)
    formData.value.sourceIP = '0.0.0.0'
  }
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = async () => {
  dialogFormVisible.value = false
  // 重置表单数据
  try {
    const res = await getServerIP()
    if (res.code === 0 && res.data.ips && res.data.ips.length > 0) {
      formData.value = {
        sourceIP: res.data.ips[0],
        sourcePort: 8080,
        protocol: 'tcp',
        targetIP: '',
        targetPort: 8080,
        status: true,
        description: ''
      }
    } else {
      formData.value = {
        sourceIP: '0.0.0.0',
        sourcePort: 8080,
        protocol: 'tcp',
        targetIP: '',
        targetPort: 8080,
        status: true,
        description: ''
      }
    }
  } catch (error) {
    formData.value = {
      sourceIP: '0.0.0.0',
      sourcePort: 8080,
      protocol: 'tcp',
      targetIP: '',
      targetPort: 8080,
      status: true,
      description: ''
    }
  }
}

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createPortForward(formData.value)
        break
      case 'update':
        res = await updatePortForward(formData.value)
        break
      default:
        res = await createPortForward(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 表单数据
const formData = ref({
  sourceIP: '0.0.0.0',
  sourcePort: 8080,
  protocol: 'tcp',
  targetIP: '',
  targetPort: 8080,
  status: true,
  description: ''
})

// 状态变更处理
const handleStatusChange = async (row) => {
  const res = await updatePortForwardStatus({
    ID: row.ID,
    status: row.status
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '状态更新成功'
    })
    // 刷新转发器状态
    await getForwarderStatusData()
  } else {
    // 如果更新失败，恢复原状态
    row.status = !row.status
  }
}

// 状态详情对话框
const statusDialogVisible = ref(false)
const currentStatus = ref(null)
const currentRow = ref(null)

// 查看状态
const viewStatus = async (row) => {
  currentRow.value = row
  statusDialogVisible.value = true
  await refreshCurrentStatus()
}

// 刷新当前状态
const refreshCurrentStatus = async () => {
  if (!currentRow.value) return
  try {
    const res = await getForwarderStatus({ ID: String(currentRow.value.ID) })
    if (res.code === 0) {
      currentStatus.value = res.data
    }
  } catch (error) {
    console.error('获取转发器状态失败:', error)
  }
}

onMounted(() => {
  // 初始获取状态
  getForwarderStatusData()
})

onBeforeUnmount(() => {
  if (autoRefreshTimer.value) {
    clearInterval(autoRefreshTimer.value)
  }
})
</script>

<style scoped lang="scss">
.gva-search-box {
  padding: 20px;
  padding-bottom: 0;
}

.gva-table-box {
  padding: 20px;
}

.gva-card-box {
  .stat-card {
    display: flex;
    align-items: center;
    padding: 10px;

    .stat-icon {
      width: 60px;
      height: 60px;
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-size: 24px;
      margin-right: 15px;
    }

    .stat-content {
      flex: 1;

      .stat-value {
        font-size: 28px;
        font-weight: bold;
        color: #303133;
        line-height: 1;
      }

      .stat-label {
        font-size: 14px;
        color: #909399;
        margin-top: 8px;
      }
    }
  }
}

.address-text {
  font-family: 'Courier New', monospace;
  font-weight: 500;
}

code {
  background: #f5f5f5;
  padding: 2px 6px;
  border-radius: 3px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
}
</style>
