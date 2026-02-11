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
      <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column prop="metadata.name" label="名称" width="200" />
        <el-table-column label="状态" width="150">
          <template #default="scope">
            <el-tag v-if="isReady(scope.row)" type="success">Ready</el-tag>
            <el-tag v-else type="danger">NotReady</el-tag>
            <el-tag v-if="scope.row.spec.unschedulable" type="warning" style="margin-left: 5px">Cordon</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="角色" width="150">
          <template #default="scope">
            {{ getRoles(scope.row) }}
          </template>
        </el-table-column>
        <el-table-column label="版本" width="150">
          <template #default="scope">
            {{ scope.row.status.nodeInfo.kubeletVersion }}
          </template>
        </el-table-column>
        <el-table-column label="IP地址" width="150">
          <template #default="scope">
            {{ getInternalIP(scope.row) }}
          </template>
        </el-table-column>
        <el-table-column label="CPU/内存" width="200">
          <template #default="scope">
            <div>CPU: {{ scope.row.status.capacity.cpu }}</div>
            <div>Mem: {{ scope.row.status.capacity.memory }}</div>
          </template>
        </el-table-column>
        <el-table-column label="GPU资源" width="150">
            <template #default="scope">
                <div v-if="getGPUCount(scope.row) > 0">
                    <el-tag type="success">{{ getGPUCount(scope.row) }} GPU</el-tag>
                </div>
                <div v-else>
                    <span style="color: #909399">无</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="创建时间" width="180">
          <template #default="scope">
            {{ formatTime(scope.row.metadata.creationTimestamp) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getNodeDetail(scope.row)">详情</el-button>
            <el-button 
                v-if="!scope.row.spec.unschedulable" 
                icon="lock" 
                type="warning" 
                link 
                @click="cordonNodeFunc(scope.row, true)"
            >停止调度</el-button>
            <el-button 
                v-else 
                icon="unlock" 
                type="success" 
                link 
                @click="cordonNodeFunc(scope.row, false)"
            >恢复调度</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- Node 详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="Node 详情" width="80%" top="5vh">
      <el-descriptions v-if="nodeDetail" :column="2" border>
        <el-descriptions-item label="名称">{{ nodeDetail.metadata?.name }}</el-descriptions-item>
        <el-descriptions-item label="状态">
             <el-tag v-if="isReady(nodeDetail)" type="success">Ready</el-tag>
             <el-tag v-else type="danger">NotReady</el-tag>
             <el-tag v-if="nodeDetail.spec.unschedulable" type="warning" style="margin-left: 5px">Cordon</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="架构">{{ nodeDetail.status?.nodeInfo?.architecture }}</el-descriptions-item>
        <el-descriptions-item label="内核版本">{{ nodeDetail.status?.nodeInfo?.kernelVersion }}</el-descriptions-item>
        <el-descriptions-item label="容器运行时">{{ nodeDetail.status?.nodeInfo?.containerRuntimeVersion }}</el-descriptions-item>
        <el-descriptions-item label="OS镜像">{{ nodeDetail.status?.nodeInfo?.osImage }}</el-descriptions-item>
        <el-descriptions-item label="Internal IP">{{ getInternalIP(nodeDetail) }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ formatTime(nodeDetail.metadata?.creationTimestamp) }}</el-descriptions-item>
      </el-descriptions>

      <div style="margin-top: 20px;">
          <h4>标签</h4>
          <el-tag v-for="(value, key) in nodeDetail?.metadata?.labels" :key="key" size="small" style="margin: 2px;">
            {{ key }}: {{ value }}
          </el-tag>
      </div>

      <div style="margin-top: 20px;">
          <h4>容量</h4>
          <el-descriptions :column="3" border>
              <el-descriptions-item label="CPU">{{ nodeDetail?.status?.capacity?.cpu }}</el-descriptions-item>
              <el-descriptions-item label="Memory">{{ nodeDetail?.status?.capacity?.memory }}</el-descriptions-item>
              <el-descriptions-item label="Pods">{{ nodeDetail?.status?.capacity?.pods }}</el-descriptions-item>
              <el-descriptions-item label="Ephemeral Storage">{{ nodeDetail?.status?.capacity?.['ephemeral-storage'] }}</el-descriptions-item>
              <el-descriptions-item label="NVIDIA GPU" v-if="getGPUCount(nodeDetail) > 0">
                  {{ getGPUCount(nodeDetail) }}
              </el-descriptions-item>
          </el-descriptions>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getNodeList, cordonNode, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'K8sNode'
})

const searchInfo = reactive({
  clusterName: ''
})

const tableData = ref([])
const clusterList = ref([])
const loading = ref(false)
const detailDialogVisible = ref(false)
const nodeDetail = ref(null)

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

// 获取列表
const onSubmit = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请选择集群')
    return
  }
  loading.value = true
  try {
    const res = await getNodeList(searchInfo)
    if (res.code === 0) {
      tableData.value = res.data.items
    }
  } finally {
    loading.value = false
  }
}

const reset = () => {
    if (clusterList.value.length > 0) {
      searchInfo.clusterName = clusterList.value[0].name
    }
    onSubmit()
}

const handleClusterChange = () => {
  onSubmit()
}

const isReady = (node) => {
    const readyCondition = node.status.conditions.find(c => c.type === 'Ready')
    return readyCondition && readyCondition.status === 'True'
}

const getRoles = (node) => {
    const roles = []
    for (const key in node.metadata.labels) {
        if (key.startsWith('node-role.kubernetes.io/')) {
            roles.push(key.replace('node-role.kubernetes.io/', ''))
        }
    }
    if (roles.length === 0) return 'worker'
    return roles.join(', ')
}

const getInternalIP = (node) => {
    const address = node.status.addresses.find(a => a.type === 'InternalIP')
    return address ? address.address : ''
}

const formatTime = (time) => {
    return formatDate(time)
}

const getGPUCount = (node) => {
    // 尝试获取常见的 GPU 资源键
    if (!node.status.capacity) return 0
    
    // NVIDIA GPU
    if (node.status.capacity['nvidia.com/gpu']) {
        return parseInt(node.status.capacity['nvidia.com/gpu'])
    }
    
    // 其他可能的 GPU 键...
    return 0
}

const getNodeDetail = (row) => {
    nodeDetail.value = row
    detailDialogVisible.value = true
}

const cordonNodeFunc = (row, unschedulable) => {
    const action = unschedulable ? '停止调度' : '恢复调度'
    ElMessageBox.confirm(`确定要${action}该节点吗?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(async () => {
        const res = await cordonNode({
            clusterName: searchInfo.clusterName,
            nodeName: row.metadata.name,
            unschedulable: unschedulable
        })
        if (res.code === 0) {
            ElMessage.success(`${action}成功`)
            onSubmit()
        }
    })
}

onMounted(() => {
  getClusters()
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
</style>
