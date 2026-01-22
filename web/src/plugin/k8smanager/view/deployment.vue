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
        <el-form-item label="命名空间">
          <el-select v-model="searchInfo.namespace" placeholder="请选择命名空间" clearable>
            <el-option label="全部" value="all" />
            <el-option
              v-for="ns in namespaceList"
              :key="ns"
              :label="ns"
              :value="ns"
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
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="name" label="Deployment名称" width="250" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="replicas" label="副本数" width="200">
          <template #default="scope">
            <span>期望: {{ scope.row.replicas }} / 就绪: {{ scope.row.readyReplicas }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="upToDate" label="最新" width="100">
          <template #default="scope">
            <span>{{ scope.row.upToDate }} / {{ scope.row.replicas }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="400">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getDeploymentDetail(scope.row)">详情</el-button>
            <el-button icon="refresh" type="primary" link @click="restartDeployment(scope.row)">重启</el-button>
            <el-button icon="edit" type="primary" link @click="scaleDeployment(scope.row)">扩缩容</el-button>
            <el-button icon="delete" type="danger" link @click="deleteDeployment(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 扩缩容对话框 -->
    <el-dialog v-model="scaleDialogVisible" title="扩缩容 Deployment" width="500px">
      <el-form :model="scaleForm" label-width="100px">
        <el-form-item label="副本数">
          <el-input-number v-model="scaleForm.replicas" :min="0" :max="100" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="scaleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmScale">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getDeploymentList, restartDeployment as restartDeploymentApi, deleteDeployment as deleteDeploymentApi, scaleDeployment as scaleDeploymentApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchInfo = reactive({
  clusterName: '',
  namespace: 'default'
})

const tableData = ref([])
const clusterList = ref([])
const namespaceList = ref(['default', 'kube-system', 'kube-public'])
const scaleDialogVisible = ref(false)
const scaleForm = reactive({
  name: '',
  namespace: '',
  replicas: 1
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

// 获取 Deployment 列表
const getDeploymentListData = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请先选择集群')
    return
  }

  const res = await getDeploymentList({
    clusterName: searchInfo.clusterName,
    namespace: searchInfo.namespace
  })

  if (res.code === 0) {
    tableData.value = res.data.items.map(item => ({
      name: item.metadata.name,
      namespace: item.metadata.namespace,
      replicas: item.spec.replicas || 0,
      readyReplicas: item.status.readyReplicas || 0,
      upToDate: item.status.updatedReplicas || 0,
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
  getDeploymentListData()
}

// 查询
const onSubmit = () => {
  getDeploymentListData()
}

// 重置
const reset = () => {
  searchInfo.namespace = 'default'
  getDeploymentListData()
}

// 查看 Deployment 详情
const getDeploymentDetail = (row) => {
  ElMessage.info('Deployment 详情功能开发中...')
}

// 重启 Deployment
const restartDeployment = (row) => {
  ElMessageBox.confirm(`确定要重启 Deployment "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await restartDeploymentApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      ElMessage.success('重启成功')
      getDeploymentListData()
    } else {
      ElMessage.error('重启失败: ' + res.msg)
    }
  })
}

// 扩缩容 Deployment
const scaleDeployment = (row) => {
  scaleForm.name = row.name
  scaleForm.namespace = row.namespace
  scaleForm.replicas = row.replicas
  scaleDialogVisible.value = true
}

// 确认扩缩容
const confirmScale = async() => {
  const res = await scaleDeploymentApi({
    clusterName: searchInfo.clusterName,
    namespace: scaleForm.namespace,
    name: scaleForm.name,
    replicas: scaleForm.replicas
  })

  if (res.code === 0) {
    ElMessage.success('扩缩容成功')
    scaleDialogVisible.value = false
    getDeploymentListData()
  } else {
    ElMessage.error('扩缩容失败: ' + res.msg)
  }
}

// 删除 Deployment
const deleteDeployment = (row) => {
  ElMessageBox.confirm(`确定要删除 Deployment "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteDeploymentApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      ElMessage.success('删除成功')
      getDeploymentListData()
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
  name: 'K8sDeployment'
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
</style>
