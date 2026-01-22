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
            <el-option label="全部" value="" />
            <el-option
              v-for="ns in namespaceList"
              :key="ns"
              :label="ns"
              :value="ns"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="资源类型">
          <el-select v-model="searchInfo.kind" placeholder="请选择资源类型" clearable>
            <el-option label="全部" value="" />
            <el-option label="Pod" value="Pod" />
            <el-option label="Deployment" value="Deployment" />
            <el-option label="Service" value="Service" />
            <el-option label="Node" value="Node" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table :data="tableData" style="width: 100%" :default-sort="{ prop: 'lastSeen', order: 'descending' }">
        <el-table-column prop="type" label="类型" width="120">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)" size="small">{{ scope.row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reason" label="原因" width="150" />
        <el-table-column prop="message" label="消息" min-width="300" show-overflow-tooltip />
        <el-table-column prop="object" label="对象" width="200" />
        <el-table-column prop="count" label="次数" width="100" />
        <el-table-column prop="lastSeen" label="最后发生时间" width="180" />
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getEventList, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage } from 'element-plus'

const searchInfo = reactive({
  clusterName: '',
  namespace: '',
  kind: '',
  limit: 100
})

const tableData = ref([])
const clusterList = ref([])
const namespaceList = ref(['default', 'kube-system', 'kube-public'])

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

// 获取事件列表
const getEventListData = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请先选择集群')
    return
  }

  const res = await getEventList({
    clusterName: searchInfo.clusterName,
    namespace: searchInfo.namespace,
    kind: searchInfo.kind,
    limit: searchInfo.limit
  })

  if (res.code === 0) {
    tableData.value = res.data.items.map(item => ({
      type: item.type,
      reason: item.reason,
      message: item.message,
      object: `${item.involvedObject.kind}/${item.involvedObject.name}`,
      count: item.count,
      lastSeen: formatTime(item.lastTimestamp)
    }))
  }
}

// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return '-'
  const date = new Date(timestamp)
  return date.toLocaleString('zh-CN')
}

// 获取事件类型标签颜色
const getTypeTagType = (type) => {
  switch (type) {
    case 'Normal':
      return 'success'
    case 'Warning':
      return 'warning'
    case 'Error':
      return 'danger'
    default:
      return 'info'
  }
}

// 集群切换
const handleClusterChange = () => {
  getEventListData()
}

// 查询
const onSubmit = () => {
  getEventListData()
}

// 重置
const reset = () => {
  searchInfo.namespace = ''
  searchInfo.kind = ''
  getEventListData()
}

onMounted(() => {
  getClusters()
})
</script>

<script>
export default {
  name: 'K8sEvent'
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
