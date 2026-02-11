
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主机名:" prop="hostName">
          <el-input v-model="formData.hostName" :clearable="true"  placeholder="请输入主机名" />
       </el-form-item>
        <el-form-item label="IP地址:" prop="ip">
          <el-input v-model="formData.ip" :clearable="true"  placeholder="请输入IP地址" />
       </el-form-item>
        <el-form-item label="SN号:" prop="sn">
          <el-input v-model="formData.sn" :clearable="true"  placeholder="请输入SN号" />
       </el-form-item>
        <el-form-item label="硬件配置:" prop="configuration">
          <el-input v-model="formData.configuration" :clearable="true"  placeholder="请输入硬件配置" />
       </el-form-item>
        <el-form-item label="服务器状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="部署服务:" prop="serviceType">
          <el-input v-model="formData.serviceType" :clearable="true"  placeholder="请输入部署服务" />
       </el-form-item>
        <el-form-item label="部署时间:" prop="deployTime">
          <el-date-picker v-model="formData.deployTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="下机时间:" prop="offlineTime">
          <el-date-picker v-model="formData.offlineTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="报废时间:" prop="scrapTime">
          <el-date-picker v-model="formData.scrapTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="报废原因:" prop="scrapReason">
          <el-input v-model="formData.scrapReason" :clearable="true"  placeholder="请输入报废原因" />
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
  createServerAsset,
  updateServerAsset,
  findServerAsset
} from '@/plugin/server_lifecycle/api/server_lifecycle'

defineOptions({
    name: 'ServerAssetForm'
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
const server_statusOptions = ref([])
const formData = ref({
            hostName: '',
            ip: '',
            sn: '',
            configuration: '',
            serviceType: '',
            deployTime: new Date(),
            offlineTime: new Date(),
            scrapTime: new Date(),
            scrapReason: '',
        })
// 验证规则
const rule = reactive({
               hostName : [{
                   required: true,
                   message: '主机名不能为空',
                   trigger: ['input','blur'],
               }],
               ip : [{
                   required: true,
                   message: 'IP地址不能为空',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findServerAsset({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    server_statusOptions.value = await getDictFunc('server_status')
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
               res = await createServerAsset(formData.value)
               break
             case 'update':
               res = await updateServerAsset(formData.value)
               break
             default:
               res = await createServerAsset(formData.value)
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
