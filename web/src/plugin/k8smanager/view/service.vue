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
        <el-table-column prop="name" label="Service名称" width="250" />
        <el-table-column prop="namespace" label="命名空间" width="150" />
        <el-table-column prop="type" label="类型" width="150" />
        <el-table-column prop="clusterIP" label="Cluster IP" width="150" />
        <el-table-column prop="externalIP" label="External IP" width="150" />
        <el-table-column prop="ports" label="端口" width="200" />
        <el-table-column prop="age" label="运行时间" width="150" />
        <el-table-column label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button icon="view" type="primary" link @click="getServiceDetail(scope.row)">详情</el-button>
            <el-button icon="delete" type="danger" link @click="deleteService(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getServiceList, deleteService as deleteServiceApi, getAllK8sClusters } from '@/plugin/k8smanager/api/cluster.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchInfo = reactive({
  clusterName: '',
  namespace: 'default'
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

// 获取 Service 列表
const getServiceListData = async() => {
  if (!searchInfo.clusterName) {
    ElMessage.warning('请先选择集群')
    return
  }

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
      externalIP: item.spec.externalIPs?.join(', ') || item.spec.type === 'LoadBalancer' ? '<pending>' : '-',
      ports: item.spec.ports?.map(p => `${p.port}/${p.protocol}`).join(', ') || '-',
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

// 查看 Service 详情
const getServiceDetail = (row) => {
  ElMessage.info('Service 详情功能开发中...')
}

// 删除 Service
const deleteService = (row) => {
  ElMessageBox.confirm(`确定要删除 Service "${row.name}" 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
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
  })
}

onMounted(() => {
  getClusters()
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
</style>
