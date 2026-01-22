<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="主机名">
          <el-input v-model="searchInfo.hostName" placeholder="请输入主机名" clearable />
        </el-form-item>
        <el-form-item label="服务标签">
          <el-input v-model="searchInfo.serviceTag" placeholder="请输入服务标签" clearable />
        </el-form-item>
        <el-form-item label="IP地址">
          <el-input v-model="searchInfo.ipAddress" placeholder="请输入IP地址" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="在线" value="online" />
            <el-option label="离线" value="offline" />
            <el-option label="维护中" value="maintenance" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 统计卡片 -->
    <div class="gva-card-box" style="padding: 0 20px; margin-bottom: 20px;">
      <el-row :gutter="20">
        <el-col :xs="12" :sm="8" :md="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #409eff;">
                <el-icon><server /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ statistics.total || 0 }}</div>
                <div class="stat-label">总服务器</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="8" :md="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #67c23a;">
                <el-icon><circle-check /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ statistics.online || 0 }}</div>
                <div class="stat-label">在线</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="8" :md="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #f56c6c;">
                <el-icon><circle-close /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ statistics.offline || 0 }}</div>
                <div class="stat-label">离线</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :xs="12" :sm="8" :md="6">
          <el-card shadow="hover">
            <div class="stat-card">
              <div class="stat-icon" style="background: #e6a23c;">
                <el-icon><warning /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ statistics.maintenance || 0 }}</div>
                <div class="stat-label">维护中</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
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
        <el-table-column align="left" label="主机名" prop="hostName" width="150" />
        <el-table-column align="left" label="服务标签" prop="serviceTag" width="140" />
        <el-table-column align="left" label="型号" prop="model" width="150" show-overflow-tooltip />
        <el-table-column align="left" label="CPU" width="120">
          <template #default="scope">
            <div>{{ scope.row.cpuModel }}</div>
            <div class="text-gray">{{ scope.row.cpuCores }}核/{{ scope.row.cpuThreads }}线程</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="内存" width="100">
          <template #default="scope">
            {{ scope.row.memoryCapacity }}GB
          </template>
        </el-table-column>
        <el-table-column align="left" label="IP地址" prop="ipAddress" width="140" />
        <el-table-column align="left" label="位置" width="120">
          <template #default="scope">
            <div>{{ scope.row.cabinet }}</div>
            <div class="text-gray">U{{ scope.row.rackPosition }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'online'" type="success">在线</el-tag>
            <el-tag v-else-if="scope.row.status === 'offline'" type="danger">离线</el-tag>
            <el-tag v-else type="warning">维护中</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="部门" prop="department" width="120" show-overflow-tooltip />
        <el-table-column align="left" label="负责人" prop="manager" width="100" />
        <el-table-column align="left" label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="viewDetails(scope.row)">详情</el-button>
            <el-button type="primary" link icon="edit" @click="updateDellAssetFunc(scope.row)">编辑</el-button>
            <el-button type="danger" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-drawer
      v-model="dialogFormVisible"
      destroy-on-close
      size="700px"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增戴尔服务器资产' : '编辑戴尔服务器资产' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确定</el-button>
            <el-button @click="closeDialog">取消</el-button>
          </div>
        </div>
      </template>

      <el-form ref="elFormRef" :model="formData" label-position="top" :rules="rules" label-width="120px">
        <el-tabs v-model="activeTab">
          <!-- 基本信息 -->
          <el-tab-pane label="基本信息" name="basic">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="主机名" prop="hostName">
                  <el-input v-model="formData.hostName" placeholder="请输入主机名" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="服务标签" prop="serviceTag">
                  <el-input v-model="formData.serviceTag" placeholder="请输入服务标签" clearable />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="资产编号">
                  <el-input v-model="formData.assetNumber" placeholder="请输入资产编号" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="型号">
                  <el-input v-model="formData.model" placeholder="请输入型号" clearable />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="序列号">
                  <el-input v-model="formData.serialNumber" placeholder="请输入序列号" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="状态">
                  <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
                    <el-option label="在线" value="online" />
                    <el-option label="离线" value="offline" />
                    <el-option label="维护中" value="maintenance" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
          </el-tab-pane>

          <!-- 硬件配置 -->
          <el-tab-pane label="硬件配置" name="hardware">
            <el-form-item label="CPU型号">
              <el-input v-model="formData.cpuModel" placeholder="例如: Intel Xeon E5-2680 v4" clearable />
            </el-form-item>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="CPU核心数">
                  <el-input-number v-model="formData.cpuCores" :min="0" style="width: 100%" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="CPU线程数">
                  <el-input-number v-model="formData.cpuThreads" :min="0" style="width: 100%" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="内存容量(GB)">
              <el-input-number v-model="formData.memoryCapacity" :min="0" style="width: 100%" />
            </el-form-item>
            <el-form-item label="磁盘信息">
              <el-input v-model="formData.diskInfo" type="textarea" :rows="3" placeholder="例如: 2x1TB SSD RAID1" />
            </el-form-item>
            <el-form-item label="网卡信息">
              <el-input v-model="formData.networkInfo" type="textarea" :rows="3" placeholder="例如: 4x10GbE" />
            </el-form-item>
          </el-tab-pane>

          <!-- 网络与位置 -->
          <el-tab-pane label="网络与位置" name="network">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="IP地址">
                  <el-input v-model="formData.ipAddress" placeholder="请输入IP地址" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="MAC地址">
                  <el-input v-model="formData.macAddress" placeholder="请输入MAC地址" clearable />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="机柜位置">
                  <el-input v-model="formData.cabinet" placeholder="例如: A01" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="机架位置(U位)">
                  <el-input v-model="formData.rackPosition" placeholder="例如: U5-U8" clearable />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="电源状态">
              <el-select v-model="formData.powerStatus" placeholder="请选择电源状态" style="width: 100%">
                <el-option label="开机" value="online" />
                <el-option label="关机" value="offline" />
              </el-select>
            </el-form-item>
          </el-tab-pane>

          <!-- 其他信息 -->
          <el-tab-pane label="其他信息" name="other">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="操作系统">
                  <el-input v-model="formData.os" placeholder="例如: CentOS 7.9" clearable />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="所属部门">
                  <el-input v-model="formData.department" placeholder="请输入所属部门" clearable />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="负责人">
              <el-input v-model="formData.manager" placeholder="请输入负责人" clearable />
            </el-form-item>
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="购买日期">
                  <el-date-picker v-model="formData.purchaseDate" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="保修到期日">
                  <el-date-picker v-model="formData.warrantyExpiry" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" style="width: 100%" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="备注">
              <el-input v-model="formData.remarks" type="textarea" :rows="4" placeholder="请输入备注信息" />
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </el-drawer>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="服务器详情" width="900px">
      <el-descriptions v-if="currentDetail" :column="2" border>
        <el-descriptions-item label="主机名">{{ currentDetail.hostName }}</el-descriptions-item>
        <el-descriptions-item label="服务标签">{{ currentDetail.serviceTag }}</el-descriptions-item>
        <el-descriptions-item label="型号">{{ currentDetail.model || '-' }}</el-descriptions-item>
        <el-descriptions-item label="序列号">{{ currentDetail.serialNumber || '-' }}</el-descriptions-item>
        <el-descriptions-item label="CPU型号">{{ currentDetail.cpuModel || '-' }}</el-descriptions-item>
        <el-descriptions-item label="CPU配置">
          {{ currentDetail.cpuCores }}核/{{ currentDetail.cpuThreads }}线程
        </el-descriptions-item>
        <el-descriptions-item label="内存容量">{{ currentDetail.memoryCapacity }}GB</el-descriptions-item>
        <el-descriptions-item label="IP地址">{{ currentDetail.ipAddress || '-' }}</el-descriptions-item>
        <el-descriptions-item label="MAC地址">{{ currentDetail.macAddress || '-' }}</el-descriptions-item>
        <el-descriptions-item label="机柜位置">{{ currentDetail.cabinet || '-' }}</el-descriptions-item>
        <el-descriptions-item label="机架位置">{{ currentDetail.rackPosition || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="currentDetail.status === 'online'" type="success">在线</el-tag>
          <el-tag v-else-if="currentDetail.status === 'offline'" type="danger">离线</el-tag>
          <el-tag v-else type="warning">维护中</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="操作系统">{{ currentDetail.os || '-' }}</el-descriptions-item>
        <el-descriptions-item label="所属部门">{{ currentDetail.department || '-' }}</el-descriptions-item>
        <el-descriptions-item label="负责人">{{ currentDetail.manager || '-' }}</el-descriptions-item>
        <el-descriptions-item label="购买日期">{{ currentDetail.purchaseDate || '-' }}</el-descriptions-item>
        <el-descriptions-item label="保修到期日">{{ currentDetail.warrantyExpiry || '-' }}</el-descriptions-item>
        <el-descriptions-item label="磁盘信息" :span="2">{{ currentDetail.diskInfo || '-' }}</el-descriptions-item>
        <el-descriptions-item label="网卡信息" :span="2">{{ currentDetail.networkInfo || '-' }}</el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">{{ currentDetail.remarks || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createDellAsset,
  deleteDellAsset,
  deleteDellAssetByIds,
  updateDellAsset,
  findDellAsset,
  getDellAssetList,
  getDellAssetStatistics
} from '@/plugin/dellasset/api/dellAsset'

import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Server,
  CircleCheck,
  CircleClose,
  Warning
} from '@element-plus/icons-vue'
import { ref, reactive, onMounted } from 'vue'

defineOptions({
  name: 'DellAsset'
})

// 表单验证规则
const rules = reactive({
  hostName: [{ required: true, message: '请输入主机名', trigger: 'blur' }],
  serviceTag: [{ required: true, message: '请输入服务标签', trigger: 'blur' }]
})

const searchFormRef = ref()
const elFormRef = ref()
const loading = ref(false)

// 分页
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})

// 统计数据
const statistics = ref({
  total: 0,
  online: 0,
  offline: 0,
  maintenance: 0
})

// 对话框
const dialogFormVisible = ref(false)
const detailDialogVisible = ref(false)
const type = ref('create')
const activeTab = ref('basic')
const currentDetail = ref(null)

// 表单数据
const formData = ref({
  hostName: '',
  serviceTag: '',
  assetNumber: '',
  model: '',
  serialNumber: '',
  cpuModel: '',
  cpuCores: 0,
  cpuThreads: 0,
  memoryCapacity: 0,
  diskInfo: '',
  networkInfo: '',
  ipAddress: '',
  macAddress: '',
  cabinet: '',
  rackPosition: '',
  powerStatus: 'offline',
  purchaseDate: '',
  warrantyExpiry: '',
  os: '',
  department: '',
  manager: '',
  status: 'offline',
  remarks: ''
})

// 多选
const multipleSelection = ref([])

// 重置搜索
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 查询
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 获取列表数据
const getTableData = async () => {
  loading.value = true
  try {
    const res = await getDellAssetList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (res.code === 0) {
      tableData.value = res.data.list
      total.value = res.data.total
      page.value = res.data.page
      pageSize.value = res.data.pageSize
    }
  } finally {
    loading.value = false
  }
}

// 获取统计数据
const getStatistics = async () => {
  try {
    const res = await getDellAssetStatistics()
    if (res.code === 0) {
      statistics.value = res.data
    }
  } catch (error) {
    console.error('获取统计数据失败:', error)
  }
}

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
    deleteDellAssetFunc(row)
  })
}

// 批量删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({ type: 'warning', message: '请选择要删除的数据' })
      return
    }
    multipleSelection.value.forEach((item) => {
      IDs.push(item.ID)
    })
    const res = await deleteDellAssetByIds({ IDs })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '删除成功' })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
      getStatistics()
    }
  })
}

// 更新
const updateDellAssetFunc = async (row) => {
  const res = await findDellAsset({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    activeTab.value = 'basic'
    dialogFormVisible.value = true
  }
}

// 删除函数
const deleteDellAssetFunc = async (row) => {
  const res = await deleteDellAsset({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '删除成功' })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
    getStatistics()
  }
}

// 打开对话框
const openDialog = () => {
  type.value = 'create'
  activeTab.value = 'basic'
  dialogFormVisible.value = true
}

// 关闭对话框
const closeDialog = () => {
  dialogFormVisible.value = false
  resetForm()
}

// 重置表单
const resetForm = () => {
  formData.value = {
    hostName: '',
    serviceTag: '',
    assetNumber: '',
    model: '',
    serialNumber: '',
    cpuModel: '',
    cpuCores: 0,
    cpuThreads: 0,
    memoryCapacity: 0,
    diskInfo: '',
    networkInfo: '',
    ipAddress: '',
    macAddress: '',
    cabinet: '',
    rackPosition: '',
    powerStatus: 'offline',
    purchaseDate: '',
    warrantyExpiry: '',
    os: '',
    department: '',
    manager: '',
    status: 'offline',
    remarks: ''
  }
}

// 确定对话框
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    if (type.value === 'create') {
      res = await createDellAsset(formData.value)
    } else {
      res = await updateDellAsset(formData.value)
    }
    if (res.code === 0) {
      ElMessage({ type: 'success', message: type.value === 'create' ? '创建成功' : '更新成功' })
      closeDialog()
      getTableData()
      getStatistics()
    }
  })
}

// 查看详情
const viewDetails = (row) => {
  currentDetail.value = row
  detailDialogVisible.value = true
}

onMounted(() => {
  getTableData()
  getStatistics()
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

.text-gray {
  color: #909399;
  font-size: 12px;
}

.gva-btn-list {
  margin-bottom: 15px;
}
</style>
