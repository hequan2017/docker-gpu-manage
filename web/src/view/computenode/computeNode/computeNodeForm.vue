
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名字:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名字" />
</el-form-item>
        <el-form-item label="区域:" prop="region">
    <el-input v-model="formData.region" :clearable="true" placeholder="请输入区域" />
</el-form-item>
        <el-form-item label="CPU:" prop="cpu">
    <el-input v-model="formData.cpu" :clearable="true" placeholder="请输入CPU" />
</el-form-item>
        <el-form-item label="内存:" prop="memory">
    <el-input v-model="formData.memory" :clearable="true" placeholder="请输入内存" />
</el-form-item>
        <el-form-item label="系统盘容量:" prop="systemDisk">
    <el-input v-model="formData.systemDisk" :clearable="true" placeholder="请输入系统盘容量" />
</el-form-item>
        <el-form-item label="数据盘容量:" prop="dataDisk">
    <el-input v-model="formData.dataDisk" :clearable="true" placeholder="请输入数据盘容量" />
</el-form-item>
        <el-form-item label="IP地址公网:" prop="publicIp">
    <el-input v-model="formData.publicIp" :clearable="true" placeholder="请输入IP地址公网" />
</el-form-item>
        <el-form-item label="IP地址内网:" prop="privateIp">
    <el-input v-model="formData.privateIp" :clearable="true" placeholder="请输入IP地址内网" />
</el-form-item>
        <el-form-item label="SSH端口:" prop="sshPort">
    <el-input v-model.number="formData.sshPort" :clearable="true" placeholder="请输入SSH端口" />
</el-form-item>
        <el-form-item label="用户名:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户名" />
</el-form-item>
        <el-form-item label="密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="显卡名称:" prop="gpuName">
    <el-input v-model="formData.gpuName" :clearable="true" placeholder="请输入显卡名称" />
</el-form-item>
        <el-form-item label="显卡数量:" prop="gpuCount">
    <el-input v-model.number="formData.gpuCount" :clearable="true" placeholder="请输入显卡数量" />
</el-form-item>
        <el-form-item label="Docker连接地址:" prop="dockerAddress">
    <el-input v-model="formData.dockerAddress" :clearable="true" placeholder="请输入Docker连接地址" />
</el-form-item>
        <el-form-item label="使用TLS:" prop="useTls">
    <el-switch v-model="formData.useTls" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="CA证书:" prop="caCert">
    <el-input v-model="formData.caCert" :clearable="true" placeholder="请输入CA证书" />
</el-form-item>
        <el-form-item label="客户端证书:" prop="clientCert">
    <el-input v-model="formData.clientCert" :clearable="true" placeholder="请输入客户端证书" />
</el-form-item>
        <el-form-item label="客户端私钥:" prop="clientKey">
    <el-input v-model="formData.clientKey" :clearable="true" placeholder="请输入客户端私钥" />
</el-form-item>
        <el-form-item label="是否上架:" prop="isOnShelf">
    <el-switch v-model="formData.isOnShelf" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createComputeNode,
  updateComputeNode,
  findComputeNode
} from '@/api/computenode/computeNode'

defineOptions({
    name: 'ComputeNodeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            region: '',
            cpu: '',
            memory: '',
            systemDisk: '',
            dataDisk: '',
            publicIp: '',
            privateIp: '',
            sshPort: 0,
            username: '',
            password: '',
            gpuName: '',
            gpuCount: 0,
            dockerAddress: '',
            useTls: false,
            caCert: '',
            clientCert: '',
            clientKey: '',
            isOnShelf: false,
            remark: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '请输入名字',
                   trigger: ['input','blur'],
               }],
               publicIp : [{
                   required: true,
                   message: '请输入公网IP地址',
                   trigger: ['input','blur'],
               }],
               privateIp : [{
                   required: true,
                   message: '请输入内网IP地址',
                   trigger: ['input','blur'],
               }],
               sshPort : [{
                   required: true,
                   message: '请输入SSH端口',
                   trigger: ['input','blur'],
               }],
               isOnShelf : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findComputeNode({ ID: route.query.id })
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
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createComputeNode(formData.value)
               break
             case 'update':
               res = await updateComputeNode(formData.value)
               break
             default:
               res = await createComputeNode(formData.value)
               break
           }
           btnLoading.value = false
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

<style>
</style>
