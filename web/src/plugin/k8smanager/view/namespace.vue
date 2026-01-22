<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="集群">
          <el-select v-model="searchInfo.clusterName" placeholder="请选择集群" clearable @change="handleClusterChange">
            <el-option
              v-for="cluster in clusterList"
              :key="cluster.id"
              :label="cluster.name"
              :value="cluster.name"
            />
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
        <el-button type="primary" icon="plus" @click="openCreateDialog">新建 Namespace</el-button>
      </div>

      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="name" label="Namespace 名称" width="300" />
        <el-table-column prop="status" label="状态" width="150">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'Active'" type="success">活跃</el-tag>
            <el-tag v-else type="info">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="age" label="运行时间" width="200" />
        <el-table-column label="标签" min-width="300">
          <template #default="scope">
            <el-tag
              v-for="(value, key) in scope.row.labels"
              :key="key"
              size="small"
              style="margin-right: 5px"
            >
              {{ key }}: {{ value }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getNamespaceDetail(scope.row)">详情</el-button>
            <el-button icon="delete" type="danger" link @click="deleteNamespace(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 创建 Namespace 对话框 -->
    <el-dialog v-model="createDialogVisible" title="创建 Namespace" width="500px">
      <el-form :model="createForm" label-width="120px">
        <el-form-item label="名称">
          <el-input v-model="createForm.name" placeholder="请输入 Namespace 名称" />
        </el-form-item>
        <el-form-item label="标签">
          <el-input v-model="createForm.labelsStr" type="textarea" :rows="3" placeholder="格式：key1=value1,key2=value2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmCreate">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getNamespaceList, createNamespace as createNamespaceApi, deleteNamespace as deleteNamespaceApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchInfo = reactive({
  clusterName: ''
})

const tableData = ref([])
const clusterList = ref([])
const createDialogVisible = ref(false)
const createForm = reactive({
  name: '',
  labelsStr: ''
})

// 获取集群列表
const getClusters = async() => {
  const res = await getAllK8sClusters()
  if (res.code === 0) {
    clusterList.value = res.data
    if (clusterList.value.length > 0) {
      searchInfo.clusterName = clusterList.value[0].name
      onSubmit()
    }
  }
}

// 获取 Namespace 列表
const getNamespaceListData = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请先选择集群')
    return
  }

  const res = await getNamespaceList({
    clusterName: searchInfo.clusterName
  })

  if (res.code === 0) {
    tableData.value = res.data.items.map(item => ({
      name: item.metadata.name,
      status: item.status.phase,
      labels: item.metadata.labels || {},
      age: calculateAge(item.metadata.creationTimestamp)
    }))
  }
}

// 计算运行时间
const calculateAge = (creationTimestamp) => {
  const now = new Date()
  const created = new Date(creationTimestamp)
  const diff = Math.floor((now - created) / 1000)

  if (diff < 60) return `${diff}s`
  if (diff < 3600) return `${Math.floor(diff / 60)}m`
  if (diff < 86400) return `${Math.floor(diff / 3600)}h`
  return `${Math.floor(diff / 86400)}d`
}

// 集群切换
const handleClusterChange = () => {
  getNamespaceListData()
}

// 查询
const onSubmit = () => {
  getNamespaceListData()
}

// 重置
const reset = () => {
  getNamespaceListData()
}

// 打开创建对话框
const openCreateDialog = () => {
  createForm.name = ''
  createForm.labelsStr = ''
  createDialogVisible.value = true
}

// 确认创建
const confirmCreate = async() => {
  if (!createForm.name) {
    ElMessage.warning('请输入 Namespace 名称')
    return
  }

  // 解析标签
  const labels = {}
  if (createForm.labelsStr) {
    try {
      const pairs = createForm.labelsStr.split(',')
      pairs.forEach(pair => {
        const [key, value] = pair.split('=')
        if (key && value) {
          labels[key.trim()] = value.trim()
        }
      })
    } catch (e) {
      ElMessage.error('标签格式错误，请使用格式：key1=value1,key2=value2')
      return
    }
  }

  const res = await createNamespaceApi({
    clusterName: searchInfo.clusterName,
    name: createForm.name,
    labels: labels
  })

  if (res.code === 0) {
    ElMessage.success('创建成功')
    createDialogVisible.value = false
    getNamespaceListData()
  } else {
    ElMessage.error('创建失败: ' + res.msg)
  }
}

// 查看 Namespace 详情
const getNamespaceDetail = (row) => {
  ElMessage.info('Namespace 详情功能开发中...')
}

// 删除 Namespace
const deleteNamespace = (row) => {
  ElMessageBox.confirm(`确定要删除 Namespace "${row.name}" 吗？删除后该命名空间下的所有资源都会被删除！`, '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteNamespaceApi({
      clusterName: searchInfo.clusterName,
      name: row.name
    })

    if (res.code === 0) {
      ElMessage.success('删除成功')
      getNamespaceListData()
    } else {
      ElMessage.error('删除失败: ' + res.msg)
    }
  })
}

onMounted(() => {
  getClusters()
})
</script>

<script>
export default {
  name: 'K8sNamespace'
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
</style>
