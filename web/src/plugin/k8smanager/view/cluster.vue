<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="集群名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="请输入集群名称" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="在线" value="online" />
            <el-option label="离线" value="offline" />
            <el-option label="未知" value="unknown" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增集群</el-button>
        <el-button icon="refresh" @click="refreshAllClusters">刷新状态</el-button>
      </div>

      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="id" width="80" />
        <el-table-column align="left" label="集群名称" prop="name" width="200" />
        <el-table-column align="left" label="版本" prop="version" width="150" />
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'online'" type="success">在线</el-tag>
            <el-tag v-else-if="scope.row.status === 'offline'" type="danger">离线</el-tag>
            <el-tag v-else type="info">未知</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="云服务商" prop="provider" width="150" />
        <el-table-column align="left" label="区域" prop="region" width="150" />
        <el-table-column align="left" label="节点数" prop="nodeCount" width="100" />
        <el-table-column align="left" label="描述" prop="description" min-width="200" show-overflow-tooltip />
        <el-table-column align="left" label="默认" prop="isDefault" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.isDefault" type="warning">是</el-tag>
            <el-tag v-else type="info">否</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button
              icon="view"
              type="primary"
              link
              @click="getK8sClusterDetail(scope.row)"
            >详情</el-button>
            <el-button
              icon="edit"
              type="primary"
              link
              @click="updateK8sClusterFunc(scope.row)"
            >编辑</el-button>
            <el-button
              icon="refresh"
              type="primary"
              link
              @click="refreshClusterStatus(scope.row)"
            >刷新</el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteK8sClusterFunc(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="dialogTitle"
      width="600px"
    >
      <el-form
        ref="clusterForm"
        :model="formData"
        :rules="rules"
        label-width="120px"
      >
        <el-form-item label="集群名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入集群名称" clearable />
        </el-form-item>
        <el-form-item label="KubeConfig" prop="kubeConfig">
          <el-input
            v-model="formData.kubeConfig"
            type="textarea"
            :rows="10"
            placeholder="请粘贴kubeconfig内容"
          />
        </el-form-item>
        <el-form-item label="API Server地址" prop="endpoint">
          <el-input v-model="formData.endpoint" placeholder="自动从kubeconfig解析，可手动修改" clearable />
        </el-form-item>
        <el-form-item label="云服务商" prop="provider">
          <el-select v-model="formData.provider" placeholder="请选择云服务商">
            <el-option label="阿里云" value="aliyun" />
            <el-option label="腾讯云" value="tencent" />
            <el-option label="AWS" value="aws" />
            <el-option label="Azure" value="azure" />
            <el-option label="自建" value="native" />
            <el-option label="其他" value="other" />
          </el-select>
        </el-form-item>
        <el-form-item label="区域" prop="region">
          <el-input v-model="formData.region" placeholder="请输入区域" clearable />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入集群描述"
          />
        </el-form-item>
        <el-form-item label="设为默认集群" prop="isDefault">
          <el-switch v-model="formData.isDefault" />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="enterDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createK8sCluster,
  deleteK8sCluster,
  deleteK8sClusterByIds,
  updateK8sCluster,
  getK8sCluster,
  getK8sClusterList,
  refreshK8sClusterStatus
} from '@/plugin/k8smanager/api/cluster.js'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})
const dialogFormVisible = ref(false)
const dialogTitle = ref('新增集群')
const type = ref('')
const formData = reactive({
  name: '',
  kubeConfig: '',
  endpoint: '',
  description: '',
  region: '',
  provider: '',
  isDefault: false
})

const rules = reactive({
  name: [{ required: true, message: '请输入集群名称', trigger: 'blur' }],
  kubeConfig: [{ required: true, message: '请输入kubeconfig内容', trigger: 'blur' }]
})

// 获取列表
const getTableData = async() => {
  const table = await getK8sClusterList({
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

// 查询
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 重置
const reset = () => {
  searchInfo.value = {}
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

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    searchInfo.value.orderKey = toLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 多选
const handleSelectionChange = (val) => {
  // 处理多选逻辑
}

// 打开对话框
const openDialog = () => {
  type.value = 'create'
  dialogTitle.value = '新增集群'
  dialogFormVisible.value = true
}

// 关闭对话框
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.name = ''
  formData.kubeConfig = ''
  formData.endpoint = ''
  formData.description = ''
  formData.region = ''
  formData.provider = ''
  formData.isDefault = false
}

// 确认操作
const enterDialog = async() => {
  const clusterForm = ref(null)
  await clusterForm.value.validate()
  let res
  if (type.value === 'create') {
    res = await createK8sCluster(formData)
  } else {
    res = await updateK8sCluster(formData)
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: type.value === 'create' ? '创建成功' : '更新成功'
    })
    closeDialog()
    getTableData()
  }
}

// 查看详情
const getK8sClusterDetail = async(row) => {
  const res = await getK8sCluster({ id: row.id })
  if (res.code === 0) {
    // 显示详情对话框
    console.log(res.data)
  }
}

// 编辑
const updateK8sClusterFunc = (row) => {
  type.value = 'update'
  dialogTitle.value = '编辑集群'
  formData.id = row.id
  formData.name = row.name
  formData.kubeConfig = row.kubeConfig
  formData.endpoint = row.endpoint
  formData.description = row.description
  formData.region = row.region
  formData.provider = row.provider
  formData.isDefault = row.isDefault
  dialogFormVisible.value = true
}

// 刷新单个集群状态
const refreshClusterStatus = async(row) => {
  const res = await refreshK8sClusterStatus({ clusterName: row.name })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '刷新成功'
    })
    getTableData()
  }
}

// 刷新所有集群状态
const refreshAllClusters = async() => {
  ElMessage({
    type: 'info',
    message: '正在刷新所有集群状态...'
  })
  getTableData()
}

// 删除
const deleteK8sClusterFunc = (row) => {
  ElMessageBox.confirm('确定要删除该集群吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteK8sCluster({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

// 初始化
getTableData()
</script>

<script>
export default {
  name: 'K8sCluster'
}
</script>

<style scoped lang="scss">
.gva-search-box {
  padding: 20px;
  padding-bottom: 0;
}

.gva-table-box {
  padding: 20px;
}

.gva-btn-list {
  margin-bottom: 20px;
}

.gva-pagination {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}
</style>
