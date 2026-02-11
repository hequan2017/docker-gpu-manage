
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模型名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入模型名称" />
       </el-form-item>
        <el-form-item label="发布者:" prop="publisher">
          <el-input v-model="formData.publisher" :clearable="true"  placeholder="请输入发布者" />
       </el-form-item>
        <el-form-item label="模型类型:" prop="type">
        <el-select v-model="formData.type" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="参数量:" prop="parameters">
          <el-input v-model="formData.parameters" :clearable="true"  placeholder="请输入参数量" />
       </el-form-item>
        <el-form-item label="魔搭地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true"  placeholder="请输入魔搭地址" />
       </el-form-item>
        <el-form-item label="模型简介:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入模型简介" />
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
  createLlmModel,
  updateLlmModel,
  findLlmModel
} from '@/plugin/llm_model/api/llm_model'

defineOptions({
    name: 'LlmModelForm'
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
const llm_typeOptions = ref([])
const formData = ref({
            name: '',
            publisher: '',
            parameters: '',
            url: '',
            description: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '模型名称不能为空',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
                   required: true,
                   message: '地址不能为空',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLlmModel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    llm_typeOptions.value = await getDictFunc('llm_type')
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
               res = await createLlmModel(formData.value)
               break
             case 'update':
               res = await updateLlmModel(formData.value)
               break
             default:
               res = await createLlmModel(formData.value)
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
