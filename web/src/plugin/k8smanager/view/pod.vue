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
          <el-select v-model="searchInfo.namespace" placeholder="请选择命名空间" clearable @change="onSubmit">
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
          <el-button icon="refresh" @click="toggleAutoRefresh" :type="autoRefresh ? 'success' : ''">
            {{ autoRefresh ? '停止自动刷新' : '自动刷新' }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table :data="tableData" style="width: 100%" v-loading="loading" :row-class-name="getRowClassName">
        <el-table-column prop="name" label="Pod名称" width="200" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'Running'" type="success">运行中</el-tag>
            <el-tag v-else-if="scope.row.status === 'Pending'" type="warning">等待中</el-tag>
            <el-tag v-else-if="scope.row.status === 'Succeeded'" type="info">成功</el-tag>
            <el-tag v-else-if="scope.row.status === 'Failed'" type="danger">失败</el-tag>
            <el-tag v-else type="info">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="Pod IP" width="150" />
        <el-table-column prop="node" label="节点" width="150" />
        <el-table-column prop="restarts" label="重启次数" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.restarts > 0" type="warning">{{ scope.row.restarts }}</el-tag>
            <span v-else>{{ scope.row.restarts }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="350">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getPodDetail(scope.row)">详情</el-button>
            <el-button icon="document" type="primary" link @click="getPodLog(scope.row)">日志</el-button>
            <el-button icon="terminal" type="primary" link @click="openTerminal(scope.row)">终端</el-button>
            <el-button icon="delete" type="danger" link @click="deletePod(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- Pod 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="Pod 详情" width="80%" top="5vh">
      <el-descriptions v-if="podDetail" :column="2" border>
        <el-descriptions-item label="名称">{{ podDetail.metadata?.name }}</el-descriptions-item>
        <el-descriptions-item label="命名空间">{{ podDetail.metadata?.namespace }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ podDetail.status?.phase }}</el-descriptions-item>
        <el-descriptions-item label="Pod IP">{{ podDetail.status?.podIP }}</el-descriptions-item>
        <el-descriptions-item label="节点">{{ podDetail.spec?.nodeName }}</el-descriptions-item>
        <el-descriptions-item label="重启次数">{{ getRestartCount(podDetail) }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatTime(podDetail.metadata?.creationTimestamp) }}</el-descriptions-item>
        <el-descriptions-item label="标签">
          <el-tag v-for="(value, key) in podDetail.metadata?.labels" :key="key" size="small" style="margin: 2px;">
            {{ key }}: {{ value }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>

      <!-- YAML -->
      <el-tabs v-if="podDetailYaml" style="margin-top: 20px;">
        <el-tab-pane label="YAML">
          <pre class="yaml-code">{{ podDetailYaml }}</pre>
        </el-tab-pane>
        <el-tab-pane label="容器">
          <el-table :data="podDetail.spec?.containers" style="width: 100%">
            <el-table-column prop="name" label="容器名称" />
            <el-table-column prop="image" label="镜像" />
            <el-table-column label="资源限制">
              <template #default="scope">
                <div v-if="scope.row.resources">
                  <div>CPU: {{ scope.row.resources.limits?.cpu || '未设置' }}</div>
                  <div>内存: {{ scope.row.resources.limits?.memory || '未设置' }}</div>
                </div>
                <span v-else>未设置</span>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>

    <!-- 日志对话框 -->
    <el-dialog v-model="logDialogVisible" title="Pod 日志" width="80%" top="5vh">
      <div class="log-toolbar">
        <el-button size="small" @click="getPodLog(currentPod)">刷新</el-button>
        <el-button size="small" @click="clearLog">清空</el-button>
        <el-checkbox v-model="logFollow" style="margin-left: 10px;">跟踪日志</el-checkbox>
        <el-input-number v-model="logTailLines" :min="10" :max="1000" :step="10" size="small" style="width: 120px; margin-left: 10px;" />
        <span style="margin-left: 5px;">行</span>
      </div>
      <el-input
        v-model="podLog"
        type="textarea"
        :rows="20"
        readonly
        placeholder="日志内容"
        class="log-content"
      />
    </el-dialog>

    <!-- 终端对话框 -->
    <el-dialog v-model="terminalDialogVisible" :title="`Pod 终端 - ${currentPod?.name}`" width="80%" top="5vh" @close="closeTerminal">
      <div class="terminal-container" ref="terminalContainer">
        <div class="terminal-placeholder">
          <el-icon><Connection /></el-icon>
          <p>终端功能开发中...</p>
          <p>请使用 kubectl 命令行工具连接到 Pod：</p>
          <pre class="command-example">kubectl exec -it {{ currentPod?.namespace }}/{{ currentPod?.name }} -- sh</pre>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { getPodList as getPodListApi, deletePod as deletePodApi, getPodLog as getPodLogApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Connection } from '@element-plus/icons-vue'

defineOptions({
  name: 'K8sPod'
})

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
const detailDialogVisible = ref(false)
const terminalDialogVisible = ref(false)
const podDetail = ref(null)
const podDetailYaml = ref('')
const currentPod = ref(null)
const loading = ref(false)

// 自动刷新
const autoRefresh = ref(false)
const autoRefreshTimer = ref(null)
const logFollow = ref(false)
const logTailLines = ref(100)

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
  loading.value = true
  try {
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
        restarts: getContainerRestartCount(item),
        age: calculateAge(item.metadata.creationTimestamp),
        raw: item
      }))
    }
  } finally {
    loading.value = false
  }
}

// 获取容器重启次数
const getContainerRestartCount = (pod) => {
  let count = 0
  if (pod.status?.containerStatuses) {
    for (const cs of pod.status.containerStatuses) {
      count += cs.restartCount || 0
    }
  }
  return count
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
  // 更新命名空间列表
  getNamespaces()
  getPodList()
}

// 获取命名空间列表
const getNamespaces = async() => {
  // 这里可以调用 API 获取实际的命名空间列表
  // 暂时使用默认列表
  if (searchInfo.namespace === 'all') {
    searchInfo.namespace = 'default'
  }
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

// 切换自动刷新
const toggleAutoRefresh = () => {
  autoRefresh.value = !autoRefresh.value
  if (autoRefresh.value) {
    autoRefreshTimer.value = setInterval(getPodList, 10000) // 每10秒刷新
    ElMessage.success('已开启自动刷新')
  } else {
    clearInterval(autoRefreshTimer.value)
    ElMessage.info('已停止自动刷新')
  }
}

// 查看 Pod 详情
const getPodDetail = async(row) => {
  currentPod.value = row
  podDetail.value = row.raw
  podDetailYaml.value = JSON.stringify(row.raw, null, 2)
  detailDialogVisible.value = true
}

// 查看 Pod 日志
const getPodLog = async(row) => {
  if (row) {
    currentPod.value = row
  }

  loading.value = true
  try {
    const res = await getPodLogApi({
      clusterName: searchInfo.clusterName,
      namespace: currentPod.value.namespace,
      podName: currentPod.value.name,
      tailLines: logTailLines.value,
      follow: false
    })

    if (res.code === 0) {
      podLog.value = res.data
      logDialogVisible.value = true
    } else {
      ElMessage.error('获取日志失败: ' + res.msg)
    }
  } finally {
    loading.value = false
  }
}

// 清空日志
const clearLog = () => {
  podLog.value = ''
}

// 打开终端
const openTerminal = (row) => {
  currentPod.value = row
  terminalDialogVisible.value = true
}

// 关闭终端
const closeTerminal = () => {
  // 清理终端连接
  currentPod.value = null
}

// 获取重启次数
const getRestartCount = (podDetail) => {
  return getContainerRestartCount(podDetail)
}

// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return '-'
  return new Date(timestamp).toLocaleString()
}

// 获取行样式
const getRowClassName = ({ row }) => {
  if (row.status === 'Failed') {
    return 'error-row'
  } else if (row.status === 'Pending') {
    return 'warning-row'
  }
  return ''
}

// 删除 Pod
const deletePod = (row) => {
  ElMessageBox.confirm(`确定要删除 Pod "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    loading.value = true
    try {
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
    } finally {
      loading.value = false
    }
  })
}

onMounted(() => {
  getClusters()
})

onBeforeUnmount(() => {
  if (autoRefreshTimer.value) {
    clearInterval(autoRefreshTimer.value)
  }
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

.log-content {
  font-family: 'Courier New', monospace;
  font-size: 12px;
  line-height: 1.4;
}

.log-toolbar {
  margin-bottom: 10px;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.yaml-code {
  background: #f5f5f5;
  padding: 15px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  overflow-x: auto;
}

.terminal-container {
  background: #1e1e1e;
  min-height: 400px;
  border-radius: 4px;
  padding: 20px;
}

.terminal-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #888;
  text-align: center;
}

.terminal-placeholder .el-icon {
  font-size: 48px;
  margin-bottom: 20px;
}

.command-example {
  background: #333;
  color: #fff;
  padding: 15px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  margin-top: 20px;
}

:deep(.el-table .error-row) {
  background-color: #fef0f0;
}

:deep(.el-table .warning-row) {
  background-color: #fdf6ec;
}
</style>
