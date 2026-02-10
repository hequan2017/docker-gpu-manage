<template>
  <div class="pcdn-page">
    <el-card shadow="never">
      <template #header>
        <div class="card-header">节点管理</div>
      </template>
      <el-form :inline="true" class="filter-row">
        <el-form-item label="地域">
          <el-select v-model="query.region" placeholder="全部" clearable>
            <el-option v-for="item in regions" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="ISP">
          <el-select v-model="query.isp" placeholder="全部" clearable>
            <el-option v-for="item in isps" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-button type="primary" @click="loadData">筛选</el-button>
      </el-form>
      <el-table :data="tableData" v-loading="loading" border>
        <el-table-column prop="nodeId" label="节点ID" min-width="140" />
        <el-table-column prop="region" label="地域" min-width="120" />
        <el-table-column prop="isp" label="ISP" min-width="120" />
        <el-table-column prop="status" label="状态" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === '在线' ? 'success' : 'danger'">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="health" label="健康度" min-width="180">
          <template #default="scope">
            <el-progress :percentage="scope.row.health" :color="scope.row.health > 80 ? '#67c23a' : '#e6a23c'" />
          </template>
        </el-table-column>
        <el-table-column prop="bandwidth" label="带宽" min-width="120" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { getPcdnNodeList } from '@/api/pcdn'

defineOptions({ name: 'PcdnNodeManagement' })

const loading = ref(false)
const query = ref({ region: '', isp: '' })
const regions = ['华北', '华东', '华南', '西南']
const isps = ['电信', '联通', '移动', 'BGP']
const tableData = ref([])

const mockData = [
  { nodeId: 'node-hz-001', region: '华东', isp: '电信', status: '在线', health: 96, bandwidth: '12.5 Gbps' },
  { nodeId: 'node-bj-002', region: '华北', isp: '联通', status: '在线', health: 88, bandwidth: '9.8 Gbps' },
  { nodeId: 'node-gz-003', region: '华南', isp: '移动', status: '离线', health: 42, bandwidth: '0 Gbps' }
]

const loadData = async () => {
  loading.value = true
  try {
    const res = await getPcdnNodeList(query.value)
    tableData.value = res?.data?.list || mockData
  } catch (e) {
    tableData.value = mockData
  } finally {
    loading.value = false
  }
}

onMounted(loadData)
</script>

<style scoped lang="scss">
.pcdn-page {
  padding: 20px;
}

.card-header {
  font-size: 16px;
  font-weight: 600;
}

.filter-row {
  margin-bottom: 12px;
}
</style>
