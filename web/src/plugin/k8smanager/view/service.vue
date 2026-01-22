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
        <el-table-column prop="name" label="Service名称" width="250" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="type" label="类型" width="150">
          <template #default="scope">
            <el-tag :type="getServiceTypeColor(scope.row.type)">
              {{ scope.row.type }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="clusterIP" label="Cluster IP" width="150" />
        <el-table-column prop="externalIP" label="External IP" width="150">
          <template #default="scope">
            <span v-if="scope.row.externalIP === '<pending>'"">
              <el-tag type="warning">待分配</el-tag>
            </span>
            <span v-else>{{ scope.row.externalIP || '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="ports" label="端口" width="200">
          <template #default="scope">
            <div v-for="(port, index) in scope.row.ports.slice(0, 3)" :key="index">
              <el-tag size="small">{{ port }}</el-tag>
            </div>
            <el-tag v-if="scope.row.ports.length > 3" size="small" type="info">+{{ scope.row.ports.length - 3 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="selector" label="选择器" min-width="200" show-overflow-tooltip />
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="250">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getServiceDetail(scope.row)">详情</el-button>
            <el-button icon="document" type="primary" link @click="viewEndpoints(scope.row)">端点</el-button>
            <el-button icon="delete" type="danger" link @click="deleteService(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="Service 详情" width="80%" top="5vh">
      <el-descriptions v-if="serviceDetail" :column="2" border>
        <el-descriptions-item label="名称">{{ serviceDetail.metadata?.name }}</el-descriptions-item>
        <el-descriptions-item label="命名空间">{{ serviceDetail.metadata?.namespace }}</el-descriptions-item>
        <el-descriptions-item label="类型">{{ serviceDetail.spec?.type }}</el-descriptions-item>
        <el-descriptions-item label="Cluster IP">{{ serviceDetail.spec?.clusterIP || '-' }}</el-descriptions-item>
        <el-descriptions-item label="External IPs">
          {{ serviceDetail.spec?.externalIPs?.join(', ') || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="Session Affinity">{{ serviceDetail.spec?.sessionAffinity || 'None' }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatTime(serviceDetail.metadata?.creationTimestamp) }}</el-descriptions-item>
        <el-descriptions-item label="标签">
          <el-tag v-for="(value, key) in serviceDetail.metadata?.labels" :key="key" size="small" style="margin: 2px;">
            {{ key }}: {{ value }}
          </el-tag>
        </el-descriptions-item>
      </el-descriptions>

      <!-- YAML -->
      <el-tabs v-if="serviceYaml" style="margin-top: 20px;">
        <el-tab-pane label="YAML">
          <pre class="yaml-code">{{ serviceYaml }}</pre>
        </el-tab-pane>
        <el-tab-pane label="端口配置">
          <el-table :data="serviceDetail.spec?.ports" style="width: 100%">
            <el-table-column prop="name" label="端口名称" />
            <el-table-column prop="protocol" label="协议" />
            <el-table-column prop="port" label="端口" />
            <el-table-column prop="targetPort" label="目标端口" />
            <el-table-column prop="nodePort" label="NodePort" />
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>

    <!-- 端点对话框 -->
    <el-dialog v-model="endpointsDialogVisible" title="Service Endpoints" width="80%">
      <el-table :data="endpoints" v-loading="endpointsLoading">
        <el-table-column prop="type" label="类型" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.type === 'NodePort' ? 'success' : 'primary'">
              {{ scope.row.type }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="address" label="地址" width="200" />
        <el-table-column prop="ports" label="端口" width="200">
          <template #default="scope">
            <div v-for="(port, index) in scope.row.ports" :key="index">
              {{ port.port }}/{{ port.protocol }}
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'
import { getServiceList, getService as getServiceApi, deleteService as deleteServiceApi, getServiceEndpoints as getServiceEndpointsApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'K8sService'
})

const searchInfo = reactive({
  clusterName: '',
  namespace: 'default'
})

const tableData = ref([])
const clusterList = ref([])
const namespaceList = ref(['default', 'kube-system', 'kube-public'])
const loading = false

// 详情对话框
const detailDialogVisible = ref(false)
const serviceDetail = ref(null)
const serviceYaml = ref('')

// 端点对话框
const endpointsDialogVisible = ref(false)
const endpoints = ref([])
const endpointsLoading = ref(false)

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

// 获取 Service 列表
const getServiceListData = async() => {
  loading.value = true
  try {
    const res = await getServiceList({
      clusterName: searchInfo.clusterName,
      namespace: searchInfo.namespace
    })

    if (res.code === 0) {
      tableData.value = res.data.items.map(item => ({
        name: item.metadata.name,
        namespace: item.metadata.namespace,
        type: item.spec.type,
        clusterIP: item.spec.clusterIP || '-',
        externalIP: item.spec.externalIPs?.join(', ') || (item.spec.type === 'LoadBalancer' ? '<pending>' : '-'),
        ports: item.spec.ports?.map(p => `${p.port}/${p.protocol}`) || [],
        selector: formatSelector(item.spec.selector),
        age: calculateAge(item.metadata.creationTimestamp),
        raw: item
      }))
    }
  } finally {
    loading.value = false
  }
}

// 获取服务类型颜色
const getServiceTypeColor = (type) => {
  switch (type) {
    case 'ClusterIP':
      return ''
    case 'NodePort':
      return 'success'
    case 'LoadBalancer':
      return 'warning'
    case 'ExternalName':
      return 'info'
    default:
      return 'info'
  }
}

// 格式化选择器
const formatSelector = (selector) => {
  if (!selector) return '-'
  const labels = selector.matchLabels || {}
  return Object.entries(labels).map(([k, v]) => `${k}=${v}`).join(', ')
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
  getServiceListData()
}

// 查询
const onSubmit = () => {
  getServiceListData()
}

// 重置
const reset = () => {
  searchInfo.namespace = 'default'
  getServiceListData()
}

// 切换自动刷新
const toggleAutoRefresh = () => {
  autoRefresh.value = !autoRefresh.value
  if (autoRefresh.value) {
    autoRefreshTimer.value = setInterval(getServiceListData, 20000) // 每20秒刷新
    ElMessage.success('已开启自动刷新')
  } else {
    clearInterval(autoRefreshTimer.value)
    ElMessage.info('已停止自动刷新')
  }
}

// 查看 Service 详情
const getServiceDetail = async(row) => {
  loading.value = true
  try {
    const res = await getServiceApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      serviceDetail.value = res.data
      serviceYaml.value = JSON.stringify(res.data, null, 2)
      detailDialogVisible.value = true
    }
  } finally {
    loading.value = false
  }
}

// 查看端点
const viewEndpoints = async(row) => {
  endpointsDialogVisible.value = true
  endpointsLoading.value = true

  try {
    const res = await getServiceEndpointsApi({
      clusterName: searchInfo.clusterName,
      namespace: row.namespace,
      name: row.name
    })

    if (res.code === 0) {
      endpoints.value = res.data.subsets || []
      endpointsLoading.value = false
    }
  } finally {
    endpointsLoading.value = false
  }
}

// 删除 Service
const deleteService = (row) => {
  ElMessageBox.confirm(`确定要删除 Service "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    loading.value = true
    try {
      const res = await deleteServiceApi({
        clusterName: searchInfo.clusterName,
        namespace: row.namespace,
        name: row.name
      })

      if (res.code === 0) {
        ElMessage.success('删除成功')
        getServiceListData()
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
  name: 'K8sService'
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
