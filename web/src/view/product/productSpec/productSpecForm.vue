
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="显卡型号:" prop="gpuModel">
    <el-input v-model="formData.gpuModel" :clearable="true" placeholder="请输入显卡型号" />
</el-form-item>
        <el-form-item label="显卡数量:" prop="gpuCount">
    <el-input v-model.number="formData.gpuCount" :clearable="true" placeholder="请输入显卡数量" />
</el-form-item>
        <el-form-item label="CPU核心数:" prop="cpuCores">
    <el-input v-model.number="formData.cpuCores" :clearable="true" placeholder="请输入CPU核心数" />
</el-form-item>
        <el-form-item label="内存(GB):" prop="memoryGb">
    <el-input v-model.number="formData.memoryGb" :clearable="true" placeholder="请输入内存(GB)" />
</el-form-item>
        <el-form-item label="系统盘容量(GB):" prop="systemDiskGb">
    <el-input v-model.number="formData.systemDiskGb" :clearable="true" placeholder="请输入系统盘容量(GB)" />
</el-form-item>
        <el-form-item label="数据盘容量(GB):" prop="dataDiskGb">
    <el-input v-model.number="formData.dataDiskGb" :clearable="true" placeholder="请输入数据盘容量(GB)" />
</el-form-item>
        <el-form-item label="价格/小时:" prop="pricePerHour">
    <el-input-number v-model="formData.pricePerHour" style="width:100%" :precision="2" :clearable="true" />
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
  createProductSpec,
  updateProductSpec,
  findProductSpec
} from '@/api/product/productSpec'

defineOptions({
    name: 'ProductSpecForm'
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
            gpuModel: '',
            gpuCount: 0,
            cpuCores: 0,
            memoryGb: 0,
            systemDiskGb: 0,
            dataDiskGb: 0,
            pricePerHour: 0,
            isOnShelf: false,
            remark: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '请输入名称',
                   trigger: ['input','blur'],
               }],
               gpuModel : [{
                   required: true,
                   message: '请输入显卡型号',
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
      const res = await findProductSpec({ ID: route.query.id })
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
               res = await createProductSpec(formData.value)
               break
             case 'update':
               res = await updateProductSpec(formData.value)
               break
             default:
               res = await createProductSpec(formData.value)
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
