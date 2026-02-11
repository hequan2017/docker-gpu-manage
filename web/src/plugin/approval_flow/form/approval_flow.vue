
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="申请标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入申请标题" />
       </el-form-item>
        <el-form-item label="版本号:" prop="version">
          <el-input v-model="formData.version" :clearable="true"  placeholder="请输入版本号" />
       </el-form-item>
        <el-form-item label="发版内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入发版内容" />
       </el-form-item>
        <el-form-item label="目标服务器:" prop="targetServer">
          <el-input v-model="formData.targetServer" :clearable="true"  placeholder="请输入目标服务器" />
       </el-form-item>
        <el-form-item label="执行命令:" prop="command">
          <el-input v-model="formData.command" :clearable="true"  placeholder="请输入执行命令" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="申请人:" prop="applicantId">
          <el-input v-model.number="formData.applicantId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批人:" prop="approverId">
          <el-input v-model.number="formData.approverId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执行日志:" prop="logs">
          <RichEdit v-model="formData.logs"/>
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
  createApprovalProcess,
  updateApprovalProcess,
  findApprovalProcess
} from '@/plugin/approval_flow/api/approval_flow'

defineOptions({
    name: 'ApprovalProcessForm'
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
const approval_statusOptions = ref([])
const formData = ref({
            title: '',
            version: '',
            content: '',
            targetServer: '',
            command: '',
            applicantId: 0,
            approverId: 0,
            logs: '',
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '标题不能为空',
                   trigger: ['input','blur'],
               }],
               version : [{
                   required: true,
                   message: '版本号不能为空',
                   trigger: ['input','blur'],
               }],
               targetServer : [{
                   required: true,
                   message: '目标服务器不能为空',
                   trigger: ['input','blur'],
               }],
               command : [{
                   required: true,
                   message: '执行命令不能为空',
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
      const res = await findApprovalProcess({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    approval_statusOptions.value = await getDictFunc('approval_status')
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
               res = await createApprovalProcess(formData.value)
               break
             case 'update':
               res = await updateApprovalProcess(formData.value)
               break
             default:
               res = await createApprovalProcess(formData.value)
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
