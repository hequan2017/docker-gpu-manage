<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="节点名称">
          <el-input v-model="searchInfo.name" placeholder="搜索节点名称" />
        </el-form-item>
        <el-form-item label="IP地址">
          <el-input v-model="searchInfo.ip" placeholder="搜索IP地址" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="节点名称" prop="name" width="120" />
        <el-table-column align="left" label="IP地址" prop="ip" width="120" />
        <el-table-column align="left" label="MAC地址" prop="mac" width="120" />
        <el-table-column align="left" label="操作系统" prop="os" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
                <el-tag :type="scope.row.status === 'online' ? 'success' : 'danger'">{{ scope.row.status }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="类型" prop="type" width="120" />
        <el-table-column align="left" label="带宽(Mbps)" prop="bandwidth" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updatePcdnNodeFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="节点名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="IP地址:">
          <el-input v-model="formData.ip" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="MAC地址:">
          <el-input v-model="formData.mac" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作系统:">
          <el-input v-model="formData.os" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPU核心:">
          <el-input-number v-model="formData.cpu" :min="1" />
        </el-form-item>
        <el-form-item label="内存(GB):">
            <el-input-number v-model="formData.memory" :min="1" />
        </el-form-item>
        <el-form-item label="磁盘(GB):">
            <el-input-number v-model="formData.disk" :min="1" />
        </el-form-item>
        <el-form-item label="状态:">
           <el-select v-model="formData.status" placeholder="请选择">
               <el-option label="Online" value="online" />
               <el-option label="Offline" value="offline" />
           </el-select>
        </el-form-item>
        <el-form-item label="类型:">
            <el-select v-model="formData.type" placeholder="请选择">
                <el-option label="Edge" value="edge" />
                <el-option label="Core" value="core" />
            </el-select>
        </el-form-item>
        <el-form-item label="带宽(Mbps):">
            <el-input-number v-model="formData.bandwidth" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createPcdnNode,
  deletePcdnNode,
  deletePcdnNodeByIds,
  updatePcdnNode,
  findPcdnNode,
  getPcdnNodeList
} from '@/plugin/pcdn/api/pcdn'

import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = reactive({})

const onReset = () => {
  searchInfo.name = ""
  searchInfo.ip = ""
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const getTableData = async() => {
  const table = await getPcdnNodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deletePcdnNode({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

const deleteVisible = ref(false)
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value.forEach(item => {
    ids.push(item.ID)
  })
  const res = await deletePcdnNodeByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

const dialogFormVisible = ref(false)
const type = ref('')
const formData = ref({
  name: '',
  ip: '',
  mac: '',
  os: '',
  cpu: 0,
  memory: 0,
  disk: 0,
  status: 'online',
  type: 'edge',
  bandwidth: 0
})

const updatePcdnNodeFunc = async(row) => {
  const res = await findPcdnNode({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rePcdnNode
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    ip: '',
    mac: '',
    os: '',
    cpu: 0,
    memory: 0,
    disk: 0,
    status: 'online',
    type: 'edge',
    bandwidth: 0
  }
}

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createPcdnNode(formData.value)
      break
    case 'update':
      res = await updatePcdnNode(formData.value)
      break
    default:
      res = await createPcdnNode(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}

const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

</script>

<style>
</style>
