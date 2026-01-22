<template>
  <div class="ai-config-container">
    <div class="config-header">
      <h2>AI Agent 配置管理</h2>
      <el-button type="primary" icon="Plus" @click="openDialog()">新建配置</el-button>
    </div>

    <el-table :data="tableData" style="width: 100%" stripe>
      <el-table-column prop="name" label="配置名称" width="200" />
      <el-table-column prop="model" label="默认模型" width="150" />
      <el-table-column prop="temperature" label="温度" width="100">
        <template #default="scope">
          {{ scope.row.temperature }}
        </template>
      </el-table-column>
      <el-table-column prop="maxTokens" label="最大Token" width="120" />
      <el-table-column prop="baseURL" label="API地址" min-width="200" show-overflow-tooltip />
      <el-table-column prop="apiKey" label="API Key" width="150">
        <template #default="scope">
          <span class="api-key-masked">{{ scope.row.apiKey }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.isActive ? 'success' : 'info'" size="small">
            {{ scope.row.isActive ? '已启用' : '未启用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" fixed="right">
        <template #default="scope">
          <el-button
            v-if="!scope.row.isActive"
            type="success"
            icon="Check"
            size="small"
            link
            @click="setActive(scope.row)"
          >
            启用
          </el-button>
          <el-button
            type="primary"
            icon="Edit"
            size="small"
            link
            @click="openDialog(scope.row)"
          >
            编辑
          </el-button>
          <el-button
            type="danger"
            icon="Delete"
            size="small"
            link
            @click="deleteConfig(scope.row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="type === 'create' ? '新建配置' : '编辑配置'"
      width="600px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-width="120px"
      >
        <el-form-item label="配置名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入配置名称"
            :disabled="type === 'update'"
          />
        </el-form-item>

        <el-form-item label="API Key" prop="apiKey">
          <el-input
            v-model="formData.apiKey"
            type="password"
            show-password
            placeholder="请输入API Key"
          />
        </el-form-item>

        <el-form-item label="API地址" prop="baseURL">
          <el-input
            v-model="formData.baseURL"
            placeholder="请输入API基础URL"
          />
        </el-form-item>

        <el-form-item label="默认模型" prop="model">
          <el-select
            v-model="formData.model"
            placeholder="请选择默认模型"
            style="width: 100%"
          >
            <el-option label="GLM-4-Plus" value="glm-4-plus" />
            <el-option label="GLM-4-Air" value="glm-4-air" />
            <el-option label="GLM-4-Flash" value="glm-4-flash" />
            <el-option label="GLM-3-Turbo" value="glm-3-turbo" />
          </el-select>
        </el-form-item>

        <el-form-item label="温度参数" prop="temperature">
          <el-slider
            v-model="formData.temperature"
            :min="0"
            :max="2"
            :step="0.1"
            :marks="{ 0: '精确', 1: '平衡', 2: '创意' }"
            show-input
          />
        </el-form-item>

        <el-form-item label="最大Token数" prop="maxTokens">
          <el-input-number
            v-model="formData.maxTokens"
            :min="128"
            :max="32768"
            :step="128"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="启用状态" prop="isActive">
          <el-switch v-model="formData.isActive" />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm" :loading="submitLoading">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getConfigList,
  createConfig,
  updateConfig,
  deleteConfig as deleteConfigApi,
  setConfigActive
} from '@/plugin/aiagent/api/config'

defineOptions({
  name: 'AIAgentConfig'
})

const tableData = ref([])
const dialogVisible = ref(false)
const type = ref('create')
const formRef = ref(null)
const submitLoading = ref(false)

const formData = ref({
  name: '',
  apiKey: '',
  baseURL: 'https://open.bigmodel.cn/api/paas/v4/',
  model: 'glm-4-plus',
  temperature: 0.7,
  maxTokens: 4096,
  isActive: false
})

const rules = reactive({
  name: [
    { required: true, message: '请输入配置名称', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  apiKey: [
    { required: true, message: '请输入API Key', trigger: 'blur' }
  ],
  baseURL: [
    { required: true, message: '请输入API地址', trigger: 'blur' },
    { type: 'url', message: '请输入有效的URL', trigger: 'blur' }
  ],
  model: [
    { required: true, message: '请选择默认模型', trigger: 'change' }
  ],
  temperature: [
    { required: true, message: '请输入温度参数', trigger: 'blur' }
  ],
  maxTokens: [
    { required: true, message: '请输入最大Token数', trigger: 'blur' }
  ]
})

// 加载配置列表
const loadConfigs = async () => {
  const res = await getConfigList()
  if (res.code === 0) {
    tableData.value = res.data || []
  }
}

// 打开对话框
const openDialog = (row = null) => {
  if (row) {
    type.value = 'update'
    formData.value = {
      ID: row.ID,
      name: row.name,
      apiKey: '', // 编辑时不显示原API Key，需要重新输入
      baseURL: row.baseURL,
      model: row.model,
      temperature: row.temperature,
      maxTokens: row.maxTokens,
      isActive: row.isActive
    }
  } else {
    type.value = 'create'
    formData.value = {
      name: '',
      apiKey: '',
      baseURL: 'https://open.bigmodel.cn/api/paas/v4/',
      model: 'glm-4-plus',
      temperature: 0.7,
      maxTokens: 4096,
      isActive: false
    }
  }
  dialogVisible.value = true
}

// 提交表单
const submitForm = async () => {
  await formRef.value?.validate()
  submitLoading.value = true

  try {
    let res
    if (type.value === 'create') {
      res = await createConfig(formData.value)
    } else {
      res = await updateConfig(formData.value)
    }

    if (res.code === 0) {
      ElMessage.success(type.value === 'create' ? '创建成功' : '更新成功')
      dialogVisible.value = false
      await loadConfigs()
    }
  } catch (error) {
    ElMessage.error('操作失败: ' + error.message)
  } finally {
    submitLoading.value = false
  }
}

// 设置启用
const setActive = async (row) => {
  const res = await setConfigActive({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('已启用该配置')
    await loadConfigs()
  }
}

// 删除配置
const deleteConfig = (row) => {
  ElMessageBox.confirm(
    `确定要删除配置"${row.name}"吗？删除后无法恢复。`,
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    const res = await deleteConfigApi({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await loadConfigs()
    }
  })
}

onMounted(() => {
  loadConfigs()
})
</script>

<style scoped>
.ai-config-container {
  padding: 20px;
  background: #fff;
  border-radius: 4px;
}

.config-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.config-header h2 {
  margin: 0;
  font-size: 20px;
  font-weight: 500;
}

.api-key-masked {
  font-family: monospace;
  color: #909399;
}
</style>
