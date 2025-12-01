
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="镜像:" prop="imageId">
    <el-select v-model="formData.imageId" placeholder="请选择镜像" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.imageId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="产品规格:" prop="specId">
    <el-select v-model="formData.specId" placeholder="请选择产品规格" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.specId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="算力节点:" prop="nodeId">
    <el-select v-model="formData.nodeId" placeholder="请选择算力节点" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.nodeId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="实例名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入实例名称" />
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
    getInstanceDataSource,
  createInstance,
  updateInstance,
  findInstance
} from '@/api/instance/instance'

defineOptions({
    name: 'InstanceForm'
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
            imageId: undefined,
            specId: undefined,
            nodeId: undefined,
            name: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               imageId : [{
                   required: true,
                   message: '请选择镜像',
                   trigger: ['input','blur'],
               }],
               specId : [{
                   required: true,
                   message: '请选择产品规格',
                   trigger: ['input','blur'],
               }],
               nodeId : [{
                   required: true,
                   message: '请选择算力节点',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '请输入实例名称',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getInstanceDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInstance({ ID: route.query.id })
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
               res = await createInstance(formData.value)
               break
             case 'update':
               res = await updateInstance(formData.value)
               break
             default:
               res = await createInstance(formData.value)
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
