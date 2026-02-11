
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
      
            <el-form-item label="主机名" prop="hostName">
  <el-input v-model="searchInfo.hostName" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="IP地址" prop="ip">
  <el-input v-model="searchInfo.ip" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="SN号" prop="sn">
  <el-input v-model="searchInfo.sn" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="服务器状态" prop="status">
    <el-tree-select v-model="searchInfo.status" placeholder="请选择服务器状态" :data="server_statusOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
           
            <el-form-item label="部署服务" prop="serviceType">
  <el-input v-model="searchInfo.serviceType" placeholder="搜索条件" />
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          <el-form-item label="硬件配置" prop="configuration">
  <el-input v-model="searchInfo.configuration" placeholder="搜索条件" />
</el-form-item>
          
          <el-form-item label="部署时间" prop="deployTime">
  <template #label>
    <span>
      部署时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.deployTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
          
          <el-form-item label="下机时间" prop="offlineTime">
  <template #label>
    <span>
      下机时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.offlineTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
          
          <el-form-item label="报废时间" prop="scrapTime">
  <template #label>
    <span>
      报废时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="!w-380px" v-model="searchInfo.scrapTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
          
          <el-form-item label="报废原因" prop="scrapReason">
  <el-input v-model="searchInfo.scrapReason" placeholder="搜索条件" />
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
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="server_lifecycle_ServerAsset" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="server_lifecycle_ServerAsset" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="server_lifecycle_ServerAsset" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="主机名" prop="hostName" width="120" />

            <el-table-column align="left" label="IP地址" prop="ip" width="120" />

            <el-table-column align="left" label="SN号" prop="sn" width="120" />

            <el-table-column align="left" label="硬件配置" prop="configuration" width="120" />

            <el-table-column align="left" label="服务器状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,server_statusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="部署服务" prop="serviceType" width="120" />

            <el-table-column sortable align="left" label="部署时间" prop="deployTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.deployTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="下机时间" prop="offlineTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.offlineTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="报废时间" prop="scrapTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.scrapTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="报废原因" prop="scrapReason" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateServerAssetFunc(scope.row)">编辑</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            
            <el-button v-if="scope.row.status === 'deploy_pending'" type="success" link icon="video-play" @click="handleDeploy(scope.row)">部署</el-button>
            <el-button v-if="scope.row.status === 'running'" type="warning" link icon="video-pause" @click="handleOffline(scope.row)">下机</el-button>
            <el-button v-if="scope.row.status === 'offline'" type="danger" link icon="delete-filled" @click="handleScrap(scope.row)">报废</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
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
             <el-form-item label="主机名:" prop="hostName">
    <el-input v-model="formData.hostName" :clearable="true" placeholder="请输入主机名" />
</el-form-item>
             <el-form-item label="IP地址:" prop="ip">
    <el-input v-model="formData.ip" :clearable="true" placeholder="请输入IP地址" />
</el-form-item>
             <el-form-item label="SN号:" prop="sn">
    <el-input v-model="formData.sn" :clearable="true" placeholder="请输入SN号" />
</el-form-item>
             <el-form-item label="硬件配置:" prop="configuration">
    <el-input v-model="formData.configuration" :clearable="true" placeholder="请输入硬件配置" />
</el-form-item>
             <el-form-item label="服务器状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择服务器状态" style="width:100%" filterable :clearable="true">
       <el-option v-for="item in [50]" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
             <el-form-item label="部署服务:" prop="serviceType">
    <el-input v-model="formData.serviceType" :clearable="true" placeholder="请输入部署服务" />
</el-form-item>
             <el-form-item label="部署时间:" prop="deployTime">
    <el-date-picker v-model="formData.deployTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="下机时间:" prop="offlineTime">
    <el-date-picker v-model="formData.offlineTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="报废时间:" prop="scrapTime">
    <el-date-picker v-model="formData.scrapTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="报废原因:" prop="scrapReason">
    <el-input v-model="formData.scrapReason" :clearable="true" placeholder="请输入报废原因" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="主机名">
    {{ detailForm.hostName }}
</el-descriptions-item>
                 <el-descriptions-item label="IP地址">
    {{ detailForm.ip }}
</el-descriptions-item>
                 <el-descriptions-item label="SN号">
    {{ detailForm.sn }}
</el-descriptions-item>
                 <el-descriptions-item label="硬件配置">
    {{ detailForm.configuration }}
</el-descriptions-item>
                 <el-descriptions-item label="服务器状态">
    {{ detailForm.status }}
</el-descriptions-item>
                 <el-descriptions-item label="部署服务">
    {{ detailForm.serviceType }}
</el-descriptions-item>
                 <el-descriptions-item label="部署时间">
    {{ detailForm.deployTime }}
</el-descriptions-item>
                 <el-descriptions-item label="下机时间">
    {{ detailForm.offlineTime }}
</el-descriptions-item>
                 <el-descriptions-item label="报废时间">
    {{ detailForm.scrapTime }}
</el-descriptions-item>
                 <el-descriptions-item label="报废原因">
    {{ detailForm.scrapReason }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createServerAsset,
  deleteServerAsset,
  deleteServerAssetByIds,
  updateServerAsset,
  findServerAsset,
  getServerAssetList
} from '@/plugin/server_lifecycle/api/server_lifecycle'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'ServerAsset'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const server_statusOptions = ref([])
const formData = ref({
            hostName: '',
            ip: '',
            sn: '',
            configuration: '',
            status: null,
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               ip : [{
                   required: true,
                   message: 'IP地址不能为空',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            deployTime: 'deploy_time',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
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
  const table = await getServerAssetList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    server_statusOptions.value = await getDictFunc('server_status')
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
            deleteServerAssetFunc(row)
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
      const res = await deleteServerAssetByIds({ IDs })
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
const updateServerAssetFunc = async(row) => {
    const res = await findServerAsset({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteServerAssetFunc = async (row) => {
    const res = await deleteServerAsset({ ID: row.ID })
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
        hostName: '',
        ip: '',
        sn: '',
        configuration: '',
        status: null,
        serviceType: '',
        deployTime: new Date(),
        offlineTime: new Date(),
        scrapTime: new Date(),
        scrapReason: '',
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
  const res = await findServerAsset({ ID: row.ID })
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

// 业务流程操作
const handleDeploy = (row) => {
  ElMessageBox.prompt('请输入部署的服务类型（如 Nginx）', '服务部署', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '服务类型不能为空'
  }).then(async ({ value }) => {
    const data = {
      ...row,
      status: 'running',
      serviceType: value,
      deployTime: new Date()
    }
    const res = await updateServerAsset(data)
    if (res.code === 0) {
      ElMessage.success('部署成功')
      getTableData()
    }
  }).catch(() => {})
}

const handleOffline = (row) => {
  ElMessageBox.confirm('确定要将该服务器下机吗？', '下机确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const data = {
      ...row,
      status: 'offline',
      offlineTime: new Date()
    }
    const res = await updateServerAsset(data)
    if (res.code === 0) {
      ElMessage.success('下机成功')
      getTableData()
    }
  }).catch(() => {})
}

const handleScrap = (row) => {
  ElMessageBox.prompt('请输入报废原因', '报废登记', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '报废原因不能为空'
  }).then(async ({ value }) => {
    const data = {
      ...row,
      status: 'scrapped',
      scrapReason: value,
      scrapTime: new Date()
    }
    const res = await updateServerAsset(data)
    if (res.code === 0) {
      ElMessage.success('报废登记成功')
      getTableData()
    }
  }).catch(() => {})
}


</script>

<style>

</style>
