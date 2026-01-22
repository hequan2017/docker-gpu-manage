<template>
  <div>
    <div class="gva-form-box">
      <el-form
        :model="formData"
        ref="elFormRef"
        label-position="right"
        :rules="rule"
        label-width="120px"
      >
        <el-form-item label="源IP地址:" prop="sourceIP">
          <el-input
            v-model="formData.sourceIP"
            :clearable="true"
            placeholder="请输入源IP地址，如: 0.0.0.0"
          />
        </el-form-item>
        <el-form-item label="源端口:" prop="sourcePort">
          <el-input-number
            v-model="formData.sourcePort"
            :min="1"
            :max="65535"
            placeholder="请输入源端口"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="协议类型:" prop="protocol">
          <el-select
            v-model="formData.protocol"
            placeholder="请选择协议类型"
            style="width: 100%"
          >
            <el-option label="TCP" value="tcp" />
            <el-option label="UDP" value="udp" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标IP地址:" prop="targetIP">
          <el-input
            v-model="formData.targetIP"
            :clearable="true"
            placeholder="请输入目标IP地址"
          />
        </el-form-item>
        <el-form-item label="目标端口:" prop="targetPort">
          <el-input-number
            v-model="formData.targetPort"
            :min="1"
            :max="65535"
            placeholder="请输入目标端口"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-switch
            v-model="formData.status"
            active-text="启用"
            inactive-text="禁用"
          />
        </el-form-item>
        <el-form-item label="规则描述:" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入规则描述"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPortForward,
  updatePortForward,
  findPortForward
} from '@/plugin/portforward/api/portForward'

defineOptions({
  name: 'PortForwardForm'
})

import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  sourceIP: '0.0.0.0',
  sourcePort: 8080,
  protocol: 'tcp',
  targetIP: '',
  targetPort: 8080,
  status: true,
  description: ''
})

// 验证规则
const rule = reactive({
  sourceIP: [
    { required: true, message: '请输入源IP地址', trigger: 'blur' },
    {
      pattern: /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/,
      message: '请输入正确的IP地址格式',
      trigger: 'blur'
    }
  ],
  sourcePort: [
    { required: true, message: '请输入源端口', trigger: 'blur' },
    { type: 'number', min: 1, max: 65535, message: '端口范围为1-65535', trigger: 'blur' }
  ],
  protocol: [
    { required: true, message: '请选择协议类型', trigger: 'change' }
  ],
  targetIP: [
    { required: true, message: '请输入目标IP地址', trigger: 'blur' },
    {
      pattern: /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/,
      message: '请输入正确的IP地址格式',
      trigger: 'blur'
    }
  ],
  targetPort: [
    { required: true, message: '请输入目标端口', trigger: 'blur' },
    { type: 'number', min: 1, max: 65535, message: '端口范围为1-65535', trigger: 'blur' }
  ]
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  if (route.query.id) {
    const res = await findPortForward({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()

// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createPortForward(formData.value)
        break
      case 'update':
        res = await updatePortForward(formData.value)
        break
      default:
        res = await createPortForward(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}
</script>

<style></style>
