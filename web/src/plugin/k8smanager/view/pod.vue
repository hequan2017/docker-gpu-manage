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
        <el-form-item label="标签">
          <el-input v-model="searchInfo.label" placeholder="请输入标签" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="name" label="Pod名称" width="200" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'Running'" type="success">运行中</el-tag>
            <el-tag v-else-if="scope.row.status === 'Pending'" type="warning">等待中</el-tag>
            <el-tag v-else-if="scope.row.status === 'Failed'" type="danger">失败</el-tag>
            <el-tag v-else type="info">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="Pod IP" width="150" />
        <el-table-column prop="node" label="节点" width="150" />
        <el-table-column prop="restarts" label="重启次数" width="100" />
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="300">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getPodDetail(scope.row)">详情</el-button>
            <el-button icon="document" type="primary" link @click="getPodLog(scope.row)">日志</el-button>
            <el-button icon="delete" type="danger" link @click="deletePod(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 日志对话框 -->
    <el-dialog v-model="logDialogVisible" title="Pod 日志" width="80%" top="5vh">
      <el-input
        v-model="podLog"
        type="textarea"
        :rows="20"
        readonly
        placeholder="日志内容"
        style="font-family: monospace; font-size: 12px;"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getPodList as getPodListApi, deletePod as deletePodApi, getPodLog as getPodLogApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchInfo = reactive({
  clusterName: '',
  namespace: 'default',
  label: '',
  showAll: false
})

const tableData = ref([])
const clusterList = ref([])
const namespaceList = ref(['default', 'kube-system', 'kube-public'])
const logDialogVisible = ref(false)
const podLog = ref('')

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

// 获取 Pod 列表
const getPodList = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请先选择集群')
    return
  }

  const res = await getPodListApi({
    clusterName: searchInfo.clusterName,
    namespace: searchInfo.namespace,
    label: searchInfo.label,
    showAll: searchInfo.showAll
  })

  if (res.code === 0) {
    tableData.value = res.data.items.map(item => ({
      name: item.metadata.name,
      namespace: item.metadata.namespace,
      status: item.status.phase,
      ip: item.status.podIP || '-',
      node: item.spec.nodeName || '-',
      restarts: item.status.containerStatuses?.reduce((sum, c) => sum + (c.restartCount || 0), 0) || 0,
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
  getPodList()
}

// 查询
const onSubmit = () => {
  getPodList()
}

// 重置
const reset = () => {
  searchInfo.namespace = 'default'
  searchInfo.label = ''
  getPodList()
}

// 查看 Pod 详情
const getPodDetail = (row) => {
  ElMessage.info('Pod 详情功能开发中...')
}

// 查看 Pod 日志
const getPodLog = async(row) => {
  const res = await getPodLogApi({
    clusterName: searchInfo.clusterName,
    namespace: row.namespace,
    podName: row.name,
    tailLines: 100
  })

  if (res.code === 0) {
    podLog.value = res.data
    logDialogVisible.value = true
  } else {
    ElMessage.error('获取日志失败: ' + res.msg)
  }
}

// 删除 Pod
const deletePod = (row) => {
  ElMessageBox.confirm(`确定要删除 Pod "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deletePodApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      ElMessage.success('删除成功')
      getPodList()
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
  name: 'K8sPod'
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
