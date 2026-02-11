
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="用户:" prop="userId">
        <el-select  v-model="formData.userId" placeholder="请选择用户" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="关联ID:" prop="relatedId">
          <el-input v-model.number="formData.relatedId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联类型(Model/Dataset/Space):" prop="relatedType">
          <el-input v-model="formData.relatedType" :clearable="false"  placeholder="请输入关联类型(Model/Dataset/Space)" />
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
    getMsDiscussionDataSource,
  createMsDiscussion,
  updateMsDiscussion,
  findMsDiscussion
} from '@/plugin/ms_clone/api/ms_discussion'

defineOptions({
    name: 'MsDiscussionForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            content: '',
            userId: undefined,
            relatedId: 0,
            relatedType: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getMsDiscussionDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMsDiscussion({ ID: route.query.id })
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
               res = await createMsDiscussion(formData.value)
               break
             case 'update':
               res = await updateMsDiscussion(formData.value)
               break
             default:
               res = await createMsDiscussion(formData.value)
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
