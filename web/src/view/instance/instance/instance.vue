
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="镜像" prop="imageId">
  <el-select v-model="searchInfo.imageId" filterable placeholder="请选择镜像" :clearable="true">
    <el-option v-for="(item,key) in dataSource.imageId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="产品规格" prop="specId">
  <el-select v-model="searchInfo.specId" filterable placeholder="请选择产品规格" :clearable="true">
    <el-option v-for="(item,key) in dataSource.specId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="算力节点" prop="nodeId">
  <el-select v-model="searchInfo.nodeId" filterable placeholder="请选择算力节点" :clearable="true">
    <el-option v-for="(item,key) in dataSource.nodeId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="实例名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="容器状态" prop="containerStatus">
  <el-input v-model="searchInfo.containerStatus" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="用户" prop="userId">
  <el-select v-model="searchInfo.userId" filterable placeholder="请选择用户" :clearable="false">
    <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
          
          <el-form-item label="Docker容器" prop="containerId">
  <el-input v-model="searchInfo.containerId" placeholder="搜索条件" />
</el-form-item>
          
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="instance_Instance" />
            <ExportExcel  template-id="instance_Instance" filterDeleted/>
            <ImportExcel  template-id="instance_Instance" @on-success="getTableData" />
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
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="镜像" prop="imageId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.imageId,scope.row.imageId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="产品规格" prop="specId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.specId,scope.row.specId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="算力节点" prop="nodeId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.nodeId,scope.row.nodeId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="Docker容器" prop="containerId" width="120" />

            <el-table-column align="left" label="实例名称" prop="name" width="120" />

            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateInstanceFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="镜像">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.imageId,detailForm.imageId) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="产品规格">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.specId,detailForm.specId) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="算力节点">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.nodeId,detailForm.nodeId) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="Docker容器">
    {{ detailForm.containerId }}
</el-descriptions-item>
                    <el-descriptions-item label="实例名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="容器状态">
    {{ detailForm.containerStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getInstanceDataSource,
  createInstance,
  deleteInstance,
  deleteInstanceByIds,
  updateInstance,
  findInstance,
  getInstanceList
} from '@/api/instance/instance'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Instance'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            imageId: undefined,
            specId: undefined,
            nodeId: undefined,
            name: '',
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getInstanceDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               imageId : [{
                   required: true,
                   message: '请选择镜像',
                   trigger: ['input','blur'],
               },
              ],
               specId : [{
                   required: true,
                   message: '请选择产品规格',
                   trigger: ['input','blur'],
               },
              ],
               nodeId : [{
                   required: true,
                   message: '请选择算力节点',
                   trigger: ['input','blur'],
               },
              ],
               name : [{
                   required: true,
                   message: '请输入实例名称',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getInstanceList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteInstanceFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteInstanceByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateInstanceFunc = async(row) => {
    const res = await findInstance({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteInstanceFunc = async (row) => {
    const res = await deleteInstance({ ID: row.ID })
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
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        imageId: undefined,
        specId: undefined,
        nodeId: undefined,
        name: '',
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findInstance({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
