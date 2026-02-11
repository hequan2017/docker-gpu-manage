
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="数据集名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入数据集名称" />
       </el-form-item>
        <el-form-item label="封面图:" prop="cover">
          <SelectImage v-model="formData.cover" file-type="image"/>
       </el-form-item>
        <el-form-item label="简介:" prop="description">
          <el-input v-model="formData.description" :clearable="false"  placeholder="请输入简介" />
       </el-form-item>
        <el-form-item label="数据集大小:" prop="size">
          <el-input v-model="formData.size" :clearable="false"  placeholder="请输入数据集大小" />
       </el-form-item>
        <el-form-item label="发布者:" prop="publisher">
          <el-input v-model="formData.publisher" :clearable="false"  placeholder="请输入发布者" />
       </el-form-item>
        <el-form-item label="详情文档:" prop="readme">
          <RichEdit v-model="formData.readme"/>
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
  createMsDataset,
  updateMsDataset,
  findMsDataset
} from '@/plugin/ms_clone/api/ms_dataset'

defineOptions({
    name: 'MsDatasetForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            cover: "",
            description: '',
            size: '',
            publisher: '',
            readme: '',
        })
// 验证规则
const rule = reactive({
               name : [{
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
      const res = await findMsDataset({ ID: route.query.id })
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
               res = await createMsDataset(formData.value)
               break
             case 'update':
               res = await updateMsDataset(formData.value)
               break
             default:
               res = await createMsDataset(formData.value)
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
