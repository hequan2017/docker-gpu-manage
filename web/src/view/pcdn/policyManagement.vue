<template>
  <div class="pcdn-page">
    <el-row :gutter="16">
      <el-col :span="8">
        <el-card shadow="never">
          <template #header>策略创建</template>
          <el-form :model="policyForm" label-width="80px">
            <el-form-item label="策略名">
              <el-input v-model="policyForm.name" placeholder="请输入策略名称" />
            </el-form-item>
            <el-form-item label="说明">
              <el-input v-model="policyForm.desc" type="textarea" :rows="3" placeholder="请输入策略说明" />
            </el-form-item>
            <el-button type="primary" @click="handleCreate">创建策略</el-button>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-card shadow="never">
          <template #header>策略列表</template>
          <el-table :data="tableData" v-loading="loading" border>
            <el-table-column prop="name" label="策略名" min-width="140" />
            <el-table-column prop="version" label="版本" min-width="100" />
            <el-table-column prop="weight" label="权重" min-width="160">
              <template #default="scope">
                <el-slider v-model="scope.row.weight" :max="100" @change="(value) => updateWeight(scope.row, value)" />
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" min-width="100" />
            <el-table-column label="操作" min-width="180" fixed="right">
              <template #default="scope">
                <el-button size="small" type="success" @click="publish(scope.row)">发布</el-button>
                <el-button size="small" type="warning" @click="rollback(scope.row)">回滚</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {
  createPcdnPolicy,
  getPcdnPolicyList,
  publishPcdnPolicy,
  rollbackPcdnPolicy,
  updatePcdnPolicyWeight
} from '@/api/pcdn'

defineOptions({ name: 'PcdnPolicyManagement' })

const loading = ref(false)
const policyForm = ref({ name: '', desc: '' })
const tableData = ref([])

const mockData = [
  { id: 1, name: '默认调度策略', version: 'v1.3.0', weight: 65, status: '已发布' },
  { id: 2, name: '大促保障策略', version: 'v2.0.1', weight: 35, status: '草稿' }
]

const loadData = async () => {
  loading.value = true
  try {
    const res = await getPcdnPolicyList()
    tableData.value = res?.data?.list || mockData
  } catch (e) {
    tableData.value = mockData
  } finally {
    loading.value = false
  }
}

const handleCreate = async () => {
  if (!policyForm.value.name) return ElMessage.warning('请填写策略名')
  await createPcdnPolicy(policyForm.value)
  ElMessage.success('策略已创建')
  policyForm.value = { name: '', desc: '' }
  loadData()
}

const updateWeight = async (row, weight) => {
  await updatePcdnPolicyWeight({ id: row.id, weight })
  ElMessage.success('权重已更新')
}

const publish = async (row) => {
  await publishPcdnPolicy({ id: row.id })
  ElMessage.success(`已发布：${row.name}`)
  loadData()
}

const rollback = async (row) => {
  await rollbackPcdnPolicy({ id: row.id })
  ElMessage.success(`已回滚：${row.name}`)
  loadData()
}

onMounted(loadData)
</script>

<style scoped lang="scss">
.pcdn-page {
  padding: 20px;
}
</style>
