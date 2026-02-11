
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="空间名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false"  placeholder="请输入空间名称" />
       </el-form-item>
        <el-form-item label="封面图:" prop="cover">
          <SelectImage v-model="formData.cover" file-type="image"/>
       </el-form-item>
        <el-form-item label="简介:" prop="description">
          <el-input v-model="formData.description" :clearable="false"  placeholder="请输入简介" />
       </el-form-item>
        <el-form-item label="SDK类型(Gradio/Streamlit):" prop="sdk">
          <el-input v-model="formData.sdk" :clearable="false"  placeholder="请输入SDK类型(Gradio/Streamlit)" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in ms_space_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="入口文件路径:" prop="appFile">
          <el-input v-model="formData.appFile" :clearable="false"  placeholder="请输入入口文件路径" />
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
  createMsSpace,
  updateMsSpace,
  findMsSpace
} from '@/plugin/ms_clone/api/ms_space'

defineOptions({
    name: 'MsSpaceForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const ms_space_statusOptions = ref([])
const formData = ref({
            name: '',
            cover: "",
            description: '',
            sdk: '',
            status: '',
            appFile: '',
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
      const res = await findMsSpace({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ms_space_statusOptions.value = await getDictFunc('ms_space_status')
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
               res = await createMsSpace(formData.value)
               break
             case 'update':
               res = await updateMsSpace(formData.value)
               break
             default:
               res = await createMsSpace(formData.value)
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
