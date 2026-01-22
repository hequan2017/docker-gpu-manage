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
      <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column prop="name" label="Deployment名称" width="250" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="replicas" label="副本数" width="200">
          <template #default="scope">
            <div>
              <div>期望: <el-tag type="success">{{ scope.row.replicas }}</el-tag></div>
              <div style="margin-top: 4px;">就绪: <el-tag :type="scope.row.readyReplicas === scope.row.replicas ? 'success' : 'warning'">{{ scope.row.readyReplicas }}</el-tag></div>
            </div>
            <div style="margin-top: 4px;">
              <el-progress
                :percentage="Math.round((scope.row.readyReplicas / scope.row.replicas) * 100)"
                :status="scope.row.readyReplicas === scope.row.replicas ? 'success' : 'warning'"
                :stroke-width="8"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="upToDate" label="最新" width="150">
          <template #default="scope">
            <el-tag :type="scope.row.upToDate === scope.row.replicas ? 'success' : 'warning'">
              {{ scope.row.upToDate }} / {{ scope.row.replicas }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="images" label="镜像" min-width="300" show-overflow-tooltip>
          <template #default="scope">
            <el-tag v-for="(img, index) in scope.row.images.slice(0, 2)" :key="index" size="small" style="margin: 2px;">
              {{ truncateImage(img) }}
            </el-tag>
            <el-tag v-if="scope.row.images.length > 2" size="small" type="info">+{{ scope.row.images.length - 2 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="400">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getDeploymentDetail(scope.row)">详情</el-button>
            <el-button icon="document" type="primary" link @click="viewPods(scope.row)">Pods</el-button>
            <el-button icon="refresh" type="primary" link @click="restartDeployment(scope.row)">重启</el-button>
            <el-button icon="edit" type="primary" link @click="scaleDeployment(scope.row)">扩缩容</el-button>
            <el-button icon="delete" type="danger" link @click="deleteDeployment(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="Deployment 详情" width="80%" top="5vh">
      <el-descriptions v-if="deploymentDetail" :column="2" border>
        <el-descriptions-item label="名称">{{ deploymentDetail.metadata?.name }}</el-descriptions-item>
        <el-descriptions-item label="命名空间">{{ deploymentDetail.metadata?.namespace }}</el-descriptions-item>
        <el-descriptions-item label="副本数">{{ deploymentDetail.spec?.replicas || 0 }}</el-descriptions-item>
        <el-descriptions-item label="选择器">{{ deploymentDetail.spec?.selector?.matchLabels || '{}' }}</el-descriptions-item>
        <el-descriptions-item label="策略类型">{{ deploymentDetail.spec?.strategy?.type || 'RollingUpdate' }}</el-descriptions-item>
        <el-descriptions-item label="最小就绪">{{ deploymentDetail.spec?.minReadySeconds || 0 }}s</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatTime(deploymentDetail.metadata?.creationTimestamp) }}</el-descriptions-item>
      </el-descriptions>

      <!-- YAML -->
      <el-tabs v-if="deploymentYaml" style="margin-top: 20px;">
        <el-tab-pane label="YAML">
          <pre class="yaml-code">{{ deploymentYaml }}</pre>
        </el-tab-pane>
        <el-tab-pane label="容器">
          <el-table :data="deploymentDetail.spec?.template?.spec?.containers" style="width: 100%">
            <el-table-column prop="name" label="容器名称" />
            <el-table-column prop="image" label="镜像" />
            <el-table-column label="端口">
              <template #default="scope">
                <div v-for="port in scope.row.ports" :key="port.containerPort">
                  {{ port.containerPort }}/{{ port.protocol }}
                </div>
              </template>
            </el-table-column>
            <el-table-column label="环境变量">
              <template #default="scope">
                <el-tag v-if="scope.row.env" size="small" type="info">{{ scope.row.env.length }} 个</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>

    <!-- Pods 对话框 -->
    <el-dialog v-model="podsDialogVisible" :title="`Deployment Pods - ${currentDeployment?.name}`" width="80%">
      <el-table :data="deploymentPods" v-loading="podsLoading">
        <el-table-column prop="name" label="Pod名称" width="200" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 'Running'" type="success">运行中</el-tag>
            <el-tag v-else-if="scope.row.status === 'Pending'" type="warning">等待中</el-tag>
            <el-tag v-else type="info">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="Pod IP" width="150" />
        <el-table-column prop="node" label="节点" width="150" />
        <el-table-column prop="restarts" label="重启次数" width="100" />
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="viewPodDetail(scope.row)">详情</el-button>
            <el-button icon="document" type="primary" link @click="viewPodLog(scope.row)">日志</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 扩缩容对话框 -->
    <el-dialog v-model="scaleDialogVisible" title="扩缩容 Deployment" width="500px">
      <el-form :model="scaleForm" label-width="100px">
        <el-form-item label="当前副本数">
          <el-tag size="large">{{ currentDeployment?.replicas || 0 }}</el-tag>
        </el-form-item>
        <el-form-item label="目标副本数">
          <el-input-number v-model="scaleForm.replicas" :min="0" :max="100" />
        </el-form-item>
        <el-alert title="提示" type="info" :closable="false" style="margin-top: 10px;">
          调整副本数将会创建或删除 Pod 以匹配目标数量
        </el-alert>
      </el-form>
      <template #footer>
        <el-button @click="scaleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmScale" :loading="scaleLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'
import { getDeploymentList, getDeployment as getDeploymentApi, getDeploymentPods as getDeploymentPodsApi, restartDeployment as restartDeploymentApi, deleteDeployment as deleteDeploymentApi, scaleDeployment as scaleDeploymentApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'K8sDeployment'
})

const searchInfo = reactive({
  clusterName: '',
  namespace: 'default',
  label: ''
})

const tableData = ref([])
const clusterList = ref([])
const namespaceList = ref(['default', 'kube-system', 'kube-public'])
const loading = ref(false)

// 详情对话框
const detailDialogVisible = ref(false)
const deploymentDetail = ref(null)
const deploymentYaml = ref('')

// Pods 对话框
const podsDialogVisible = ref(false)
const deploymentPods = ref([])
const podsLoading = ref(false)

// 扩缩容
const scaleDialogVisible = ref(false)
const scaleForm = reactive({
  name: '',
  namespace: '',
  replicas: 1
})
const scaleLoading = ref(false)
const currentDeployment = ref(null)

// 自动刷新
const autoRefresh = ref(false)
const autoRefreshTimer = ref(null)

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
  loading.value = true
  try {
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
        availableReplicas: item.status.availableReplicas || 0,
        images: item.spec.template.spec.containers.map(c => c.image),
        age: calculateAge(item.metadata.creationTimestamp),
        raw: item
      }))
    }
  } finally {
    loading.value = false
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

// 截断镜像名称
const truncateImage = (image) => {
  if (image.length > 50) {
    const parts = image.split('/')
    return '...' + parts[parts.length - 1]
  }
  return image
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
  searchInfo.label = ''
  getDeploymentListData()
}

// 切换自动刷新
const toggleAutoRefresh = () => {
  autoRefresh.value = !autoRefresh.value
  if (autoRefresh.value) {
    autoRefreshTimer.value = setInterval(getDeploymentListData, 15000) // 每15秒刷新
    ElMessage.success('已开启自动刷新')
  } else {
    clearInterval(autoRefreshTimer.value)
    ElMessage.info('已停止自动刷新')
  }
}

// 查看 Deployment 详情
const getDeploymentDetail = async(row) => {
  currentDeployment.value = row
  loading.value = true
  try {
    const res = await getDeploymentApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      deploymentDetail.value = res.data
      deploymentYaml.value = JSON.stringify(res.data, null, 2)
      detailDialogVisible.value = true
    }
  } finally {
    loading.value = false
  }
}

// 查看 Pods
const viewPods = async(row) => {
  currentDeployment.value = row
  podsDialogVisible.value = true
  podsLoading.value = true

  try {
    const res = await getDeploymentPodsApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      deploymentPods.value = res.data.items.map(item => ({
        name: item.metadata.name,
        namespace: item.metadata.namespace,
        status: item.status.phase,
        ip: item.status.podIP || '-',
        node: item.spec.nodeName || '-',
        restarts: getContainerRestartCount(item),
        raw: item
      }))
    }
  } finally {
    podsLoading.value = false
  }
}

// 查看 Pod 详情
const viewPodDetail = (row) => {
  ElMessage.info(`查看 Pod 详情: ${row.name}`)
  // 可以跳转到 Pod 页面或打开详情对话框
}

// 查看 Pod 日志
const viewPodLog = (row) => {
  ElMessage.info(`查看 Pod 日志: ${row.name}`)
  // 可以打开日志对话框
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

// 重启 Deployment
const restartDeployment = (row) => {
  ElMessageBox.confirm(`确定要重启 Deployment "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    loading.value = true
    try {
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
    } finally {
      loading.value = false
    }
  })
}

// 扩缩容 Deployment
const scaleDeployment = (row) => {
  scaleForm.name = row.name
  scaleForm.namespace = row.namespace
  scaleForm.replicas = row.replicas
  currentDeployment.value = row
  scaleDialogVisible.value = true
}

// 确认扩缩容
const confirmScale = async() => {
  scaleLoading.value = true
  try {
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
  } finally {
    scaleLoading.value = false
  }
}

// 删除 Deployment
const deleteDeployment = (row) => {
  ElMessageBox.confirm(`确定要删除 Deployment "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    loading.value = true
    try {
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
    } finally {
      loading.value = false
    }
  })
}

// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return '-'
  return new Date(timestamp).toLocaleString()
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

.yaml-code {
  background: #f5f5f5;
  padding: 15px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  overflow-x: auto;
  max-height: 500px;
  overflow-y: auto;
}
</style>
