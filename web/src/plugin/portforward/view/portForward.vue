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
        </el-form-item>
      </el-form>
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
            {{ scope.row.sourceIP }}:{{ scope.row.sourcePort }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="协议" prop="protocol" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.protocol === 'tcp' ? 'primary' : 'success'">
              {{ scope.row.protocol.toUpperCase() }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="目标地址" width="180">
          <template #default="scope">
            {{ scope.row.targetIP }}:{{ scope.row.targetPort }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="状态" prop="status" width="100">
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
          width="200"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
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
  getServerIP
} from '@/plugin/portforward/api/portForward'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'

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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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

// 查询
const getTableData = async () => {
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
      // 使用第一个非127.0.0.1的IP地址
      formData.value.sourceIP = res.data.ips[0]
    } else {
      // 如果获取失败，使用默认值
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
  // 获取服务器IP作为默认值
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
  } else {
    // 如果更新失败，恢复原状态
    row.status = !row.status
  }
}
</script>

<style></style>
