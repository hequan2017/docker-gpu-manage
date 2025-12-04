
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
    <el-option v-for="(item,key) in searchDataSource.nodeId" :key="key" :label="item.label" :value="item.value" />
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
          
          <el-form-item label="容器" prop="containerId">
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
            <el-button type="primary" icon="plus" @click="openCreateDialog">新增实例</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate template-id="instance_Instance" />
            <ExportExcel template-id="instance_Instance" filterDeleted/>
            <ImportExcel template-id="instance_Instance" @on-success="getTableData" />
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
        
        <el-table-column align="left" label="实例名称" prop="name" width="120" />

        <el-table-column align="left" label="创建用户" prop="userName" width="120">
          <template #default="scope">
            <span>{{ scope.row.userName || '-' }}</span>
          </template>
        </el-table-column>

            <el-table-column align="left" label="镜像" prop="imageId" width="180">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.imageId,scope.row.imageId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="产品规格" prop="specId" width="120">
    <template #default="scope">
        <span>{{ getSpecName(scope.row.specId) }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="算力节点" prop="nodeId" width="120">
    <template #default="scope">
        <span>{{ filterDataSource(searchDataSource.nodeId,scope.row.nodeId) }}</span>
    </template>
</el-table-column>
            <!-- <el-table-column align="left" label="容器" prop="containerId" width="150" show-overflow-tooltip>
    <template #default="scope">
        <span class="container-id">{{ scope.row.containerId ? scope.row.containerId.substring(0, 12) : '-' }}</span>
    </template>
</el-table-column> -->



            <el-table-column align="left" label="状态" prop="containerStatus" width="100">
              <template #default="scope">
                <el-tag :type="getStatusType(scope.row.containerStatus)" size="small">
                  {{ scope.row.containerStatus || '-' }}
                </el-tag>
              </template>
            </el-table-column>

            <!-- <el-table-column align="left" label="备注" prop="remark" width="120" /> -->

        <el-table-column align="left" label="详情" fixed="right" min-width="150">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button 
              type="success" 
              link 
              class="table-button" 
              @click="showSshConnectionInfo(scope.row)"
            ><el-icon style="margin-right: 3px"><Connection /></el-icon>SSH连接</el-button>
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="350">
            <template #default="scope">
            <el-button 
              type="success" 
              link 
              class="table-button" 
              :disabled="scope.row.containerStatus === 'running' || !scope.row.containerId"
              @click="handleStartContainer(scope.row)"
            ><el-icon style="margin-right: 3px"><VideoPlay /></el-icon>启动</el-button>
            <el-button 
              type="warning" 
              link 
              class="table-button" 
              :disabled="scope.row.containerStatus !== 'running'"
              @click="handleStopContainer(scope.row)"
            ><el-icon style="margin-right: 3px"><VideoPause /></el-icon>停止</el-button>
            <el-button 
              type="primary" 
              link 
              class="table-button" 
              :disabled="!scope.row.containerId"
              @click="handleRestartContainer(scope.row)"
            ><el-icon style="margin-right: 3px"><RefreshRight /></el-icon>重启</el-button>
            <el-button 
              type="info" 
              link 
              class="table-button" 
              :disabled="!scope.row.containerId"
              @click="openLogsDialog(scope.row)"
            ><el-icon style="margin-right: 3px"><Document /></el-icon>日志</el-button>
            <el-button 
              type="primary" 
              link 
              class="table-button" 
              :disabled="scope.row.containerStatus !== 'running'"
              @click="openTerminalDialog(scope.row)"
            ><el-icon style="margin-right: 3px"><Monitor /></el-icon>终端</el-button>
            <el-button type="danger" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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

    <!-- 新增实例弹窗 - 参考图片设计 -->
    <el-dialog 
      v-model="createDialogVisible" 
      title="新增实例" 
      width="1200px" 
      :close-on-click-modal="false"
      destroy-on-close
      class="create-instance-dialog"
    >
      <div class="create-instance-content">
        <!-- 步骤1: 选择镜像 -->
        <div class="step-section">
          <div class="step-header">
            <el-icon class="step-icon" :class="{ 'completed': createForm.imageId }"><CircleCheck /></el-icon>
            <span class="step-title">选择镜像</span>
          </div>
          <div class="step-content">
            <el-select 
              v-model="createForm.imageId" 
              placeholder="请选择镜像" 
              filterable 
              style="width: 100%"
              size="large"
              @change="onImageChange"
            >
              <el-option 
                v-for="item in dataSource.imageId" 
                :key="item.value" 
                :label="item.label" 
                :value="item.value" 
              />
            </el-select>
          </div>
        </div>

        <!-- 步骤2: 选择显卡规格 -->
        <div class="step-section">
          <div class="step-header">
            <el-icon class="step-icon" :class="{ 'completed': createForm.specId }"><CircleCheck /></el-icon>
            <span class="step-title">选择显卡规格</span>
          </div>
          <div class="step-content">
            <div v-if="!createForm.imageId" class="empty-tip">
              <el-icon><Warning /></el-icon>
              <span>请先选择镜像</span>
            </div>
            <div v-else class="spec-cards">
              <div 
                v-for="spec in filteredSpecs" 
                :key="spec.value"
                class="spec-card"
                :class="{ 'selected': createForm.specId === spec.value }"
                @click="selectSpec(spec)"
              >
                <div class="spec-card-header">
                  <el-icon class="gpu-icon"><Cpu /></el-icon>
                  <span class="spec-name">{{ spec.name }}</span>
                </div>
                <div class="spec-card-body">
                  <div class="spec-info">
                    <span class="spec-gpu">{{ spec.gpu_model }} x {{ spec.gpu_count || 1 }}</span>
                  </div>
                  <div class="spec-details">
                    <span v-if="spec.cpu_cores">CPU: {{ spec.cpu_cores }}核</span>
                    <span v-if="spec.memory_gb">内存: {{ spec.memory_gb }}GB</span>
                    <span v-if="spec.memory_capacity">显存: {{ spec.memory_capacity }}GB</span>
                    <span v-if="spec.system_disk_gb">系统盘: {{ spec.system_disk_gb }}GB</span>
                    <span v-if="spec.data_disk_gb">数据盘: {{ spec.data_disk_gb }}GB</span>
                  </div>
                  <div class="spec-price" v-if="spec.price_per_hour">
                    <span class="price">¥{{ spec.price_per_hour }}</span>
                    <span class="unit">/小时</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 步骤3: 选择节点 -->
        <div class="step-section">
          <div class="step-header">
            <el-icon class="step-icon" :class="{ 'completed': createForm.nodeId }"><CircleCheck /></el-icon>
            <span class="step-title">选择主机</span>
            <el-button 
              v-if="createForm.specId" 
              size="small" 
              @click="refreshNodes" 
              :loading="nodesLoading"
              style="margin-left: auto;"
            >
              刷新
            </el-button>
          </div>
          <div class="step-content">
            <div v-if="!createForm.specId" class="empty-tip">
              <el-icon><Warning /></el-icon>
              <span>请先选择显卡规格</span>
            </div>
            <div v-else-if="nodesLoading" class="loading-tip">
              <el-icon class="is-loading"><Loading /></el-icon>
              <span>正在查询可用节点...</span>
            </div>
            <div v-else-if="availableNodes.length === 0" class="empty-tip">
              <el-icon><Warning /></el-icon>
              <span>暂无满足条件的可用节点</span>
            </div>
            <el-table 
              v-else
              :data="availableNodes" 
              style="width: 100%"
              highlight-current-row
              @current-change="selectNode"
              :row-class-name="getRowClassName"
            >
              <el-table-column prop="name" label="主机(可用资源)" width="120" />
              <el-table-column prop="gpuName" label="GPU型号" width="120" />
              <el-table-column label="GPU数量" width="120">
                <template #default="scope">
                  <span class="available-count">{{ scope.row.availableGpu }}</span>
                </template>
              </el-table-column>
              <el-table-column label="CPU" width="110">
                <template #default="scope">
                  {{ scope.row.availableCpu }}
                </template>
              </el-table-column>
              <el-table-column label="内存" width="100">
                <template #default="scope">
                  {{ scope.row.availableMemory }}GB
                </template>
              </el-table-column>
              <el-table-column label="单卡可用显存" width="120">
                <template #default="scope">
                  {{ scope.row.memoryCapacity || 0 }}GB
                </template>
              </el-table-column>
              <el-table-column label="系统盘" width="100">
                <template #default="scope">
                  {{ scope.row.availableSystemDisk }}GB
                </template>
              </el-table-column>

              <el-table-column label="数据盘" width="100">
                <template #default="scope">
                  {{ scope.row.availableDataDisk }}GB
                </template>
              </el-table-column>

              <el-table-column prop="region" label="区域" width="100" />
              <el-table-column label="价格" width="120">
                <template #default="scope">
                  <span class="price-text">¥{{ scope.row.pricePerHour }}/时</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>

        <!-- 步骤4: 实例信息 -->
        <div class="step-section">
          <div class="step-header">
            <el-icon class="step-icon" :class="{ 'completed': createForm.name }"><CircleCheck /></el-icon>
            <span class="step-title">实例信息</span>
          </div>
          <div class="step-content">
            <el-form :model="createForm" :rules="createFormRules" ref="createFormRef" label-width="80px">
              <el-form-item label="实例名称" prop="name" required>
                <el-input 
                  v-model="createForm.name" 
                  placeholder="请输入实例名称（不支持中文）" 
                  @input="handleNameInput"
                  maxlength="50"
                />
                <div class="form-tip">提示：实例名称仅支持字母、数字、横线和下划线，不支持中文</div>
              </el-form-item>
              <el-form-item label="备注">
                <el-input v-model="createForm.remark" type="textarea" placeholder="请输入备注" :rows="2" />
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <div class="summary" v-if="selectedSpec">
            <span>费用: </span>
            <span class="total-price">¥{{ selectedSpec.price_per_hour || 0 }}</span>
            <span class="price-unit">/小时</span>
          </div>
          <div class="actions">
            <el-button @click="createDialogVisible = false">取消</el-button>
            <el-button 
              type="primary" 
              @click="submitCreate" 
              :loading="btnLoading"
              :disabled="!canSubmit"
            >
              创建实例
            </el-button>
          </div>
        </div>
      </template>
    </el-dialog>

    <!-- 编辑弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">编辑</span>
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
        <el-option v-for="(item,key) in searchDataSource.nodeId" :key="key" :label="item.label" :value="item.value" />
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
        <div>
          <div v-if="getSpecInfo(detailForm.specId)" style="font-size: 12px; color: #666;">
            <div v-if="getSpecInfo(detailForm.specId).gpu_model">
              GPU: {{ getSpecInfo(detailForm.specId).gpu_model }} x {{ getSpecInfo(detailForm.specId).gpu_count || 0 }}
            </div>
            <div v-if="getSpecInfo(detailForm.specId).cpu_cores">
              CPU: {{ getSpecInfo(detailForm.specId).cpu_cores }}核
            </div>
            <div v-if="getSpecInfo(detailForm.specId).memory_gb">
              内存: {{ getSpecInfo(detailForm.specId).memory_gb }}GB
            </div>
            <div v-if="getSpecInfo(detailForm.specId).memory_capacity">
              显存: {{ getSpecInfo(detailForm.specId).memory_capacity }}GB
            </div>
            <div v-if="getSpecInfo(detailForm.specId).system_disk_gb">
              系统盘: {{ getSpecInfo(detailForm.specId).system_disk_gb }}GB
            </div>
            <div v-if="getSpecInfo(detailForm.specId).data_disk_gb">
              数据盘: {{ getSpecInfo(detailForm.specId).data_disk_gb }}GB
            </div>
            <div v-if="getSpecInfo(detailForm.specId).price_per_hour" style="margin-top: 4px; color: #409EFF;">
              价格: ¥{{ getSpecInfo(detailForm.specId).price_per_hour }}/小时
            </div>
          </div>
        </div>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="算力节点">
    <template #default="scope">
        <span>{{ filterDataSource(searchDataSource.nodeId,detailForm.nodeId) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="容器ID">
    {{ detailForm.containerId || '-' }}
</el-descriptions-item>
                    <el-descriptions-item label="容器名称">
    {{ detailForm.containerName || '-' }}
</el-descriptions-item>
                    <el-descriptions-item label="实例名称">
    {{ detailForm.name }}
    </el-descriptions-item>
                    <el-descriptions-item label="创建用户">
    {{ detailForm.userName || '-' }}
    </el-descriptions-item>
                    <el-descriptions-item label="容器状态">
    {{ detailForm.containerStatus }}
    </el-descriptions-item>
                    <el-descriptions-item label="资源使用率" v-if="detailForm.containerId && detailForm.containerStatus === 'running'">
    <div class="stats-container">
      <div class="stats-item">
        <div class="stats-label">
          <span>CPU使用率</span>
          <span class="stats-value">{{ (detailForm.cpuUsagePercent ?? 0).toFixed(2) }}%</span>
        </div>
        <el-progress 
          :percentage="Number(((detailForm.cpuUsagePercent ?? 0).toFixed(2)))" 
          :color="getProgressColor(detailForm.cpuUsagePercent)"
          :stroke-width="12"
        />
      </div>
      <div class="stats-item" style="margin-top: 16px;">
        <div class="stats-label">
          <span>内存使用率</span>
          <span class="stats-value">{{ (detailForm.memoryUsagePercent ?? 0).toFixed(2) }}%</span>
        </div>
        <el-progress 
          :percentage="Number(((detailForm.memoryUsagePercent ?? 0).toFixed(2)))" 
          :color="getProgressColor(detailForm.memoryUsagePercent)"
          :stroke-width="12"
        />
      </div>
      
      <!-- GPU 显存 -->
      <div class="stats-item" style="margin-top: 16px;" v-if="shouldShowGpuStats">
        <div class="stats-label">
          <span>显存使用率</span>
          <span class="stats-value">{{ (detailForm.gpuMemoryUsageRate ?? 0).toFixed(2) }}%</span>
        </div>
        <el-progress 
          :percentage="Number(((detailForm.gpuMemoryUsageRate ?? 0).toFixed(2)))" 
          :color="getProgressColor(detailForm.gpuMemoryUsageRate)"
          :stroke-width="12"
        />
      </div>

      <el-button 
        size="small" 
        @click="refreshDetail" 
        :loading="statsLoading"
        style="margin-top: 12px;"
      >
        <el-icon><RefreshRight /></el-icon>刷新
      </el-button>
    </div>
    </el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

    <!-- 日志弹窗 -->
    <el-dialog 
      v-model="logsDialogVisible" 
      title="容器日志" 
      width="900px" 
      destroy-on-close
      class="logs-dialog"
      @close="stopLogsAutoRefresh"
    >
      <div class="logs-toolbar">
        <el-select v-model="logsTail" size="small" style="width: 120px" @change="refreshLogs">
          <el-option label="最近50行" value="50" />
          <el-option label="最近100行" value="100" />
          <el-option label="最近500行" value="500" />
          <el-option label="最近1000行" value="1000" />
        </el-select>
        <el-switch 
          v-model="logsAutoRefresh" 
          size="small"
          active-text="自动刷新" 
          inactive-text=""
          @change="toggleLogsAutoRefresh"
          style="margin-left: 12px;"
        />
        <el-button size="small" @click="refreshLogs" :loading="logsLoading" style="margin-left: 12px;">
          <el-icon><RefreshRight /></el-icon>手动刷新
        </el-button>
      </div>
      <div class="logs-content" ref="logsContainer">
        <pre v-if="logsContent">{{ logsContent }}</pre>
        <el-empty v-else-if="!logsLoading" description="暂无日志" />
        <div v-if="logsLoading && !logsContent" class="logs-loading">
          <el-icon class="is-loading"><Loading /></el-icon>
          <span>加载中...</span>
        </div>
      </div>
    </el-dialog>

    <!-- 终端弹窗 -->
    <el-dialog 
      v-model="terminalDialogVisible" 
      title="容器终端" 
      width="900px" 
      destroy-on-close
      class="terminal-dialog"
      :close-on-click-modal="true"
      :close-on-press-escape="true"
      @close="handleTerminalClose"
    >
      <div class="terminal-toolbar">
        <div class="toolbar-left">
          <div class="shell-select-wrapper">
            <el-icon class="toolbar-icon"><Monitor /></el-icon>
            <span class="toolbar-label">Shell类型</span>
            <el-select 
              v-model="terminalShell" 
              size="default" 
              class="shell-select"
              @change="reconnectTerminal" 
              :disabled="terminalConnected"
              placeholder="选择Shell"
            >
              <el-option label="/bin/bash" value="bash">
                <div class="shell-option">
                  <el-icon class="shell-icon"><Monitor /></el-icon>
                  <span class="shell-name">/bin/bash</span>
                  <span class="shell-desc">Bash Shell</span>
                </div>
              </el-option>
              <el-option label="/bin/sh" value="sh">
                <div class="shell-option">
                  <el-icon class="shell-icon"><Monitor /></el-icon>
                  <span class="shell-name">/bin/sh</span>
                  <span class="shell-desc">POSIX Shell</span>
                </div>
              </el-option>
            </el-select>
          </div>
        </div>
        <div class="toolbar-right">
          <el-tag 
            :type="terminalConnected ? 'success' : 'info'" 
            size="small"
            effect="plain"
          >
            <el-icon style="margin-right: 4px;">
              <component :is="terminalConnected ? 'CircleCheck' : 'Loading'" />
            </el-icon>
            {{ terminalConnected ? '已连接' : '未连接' }}
          </el-tag>
        </div>
      </div>
      <div class="terminal-container" ref="terminalContainer"></div>
    </el-dialog>

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
  getInstanceList,
  getAvailableNodes,
  startContainer,
  stopContainer,
  restartContainer,
  getContainerLogs,
  getTerminalWsUrl
} from '@/api/instance/instance'
import { getJumpboxConfig } from '@/api/system'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { useAppStore } from "@/pinia"
import { CircleCheck, Cpu, Warning, Loading, ArrowDown, VideoPlay, VideoPause, RefreshRight, Document, Monitor, Connection } from '@element-plus/icons-vue'

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
  const dataSource = ref({
    imageId: [],
    specId: [],
    nodeId: [],
    userId: []
  })
  // 搜索用的数据源（包含所有节点）
  const searchDataSource = ref({
    nodeId: []
  })
  
  const getDataSourceFunc = async()=>{
    const res = await getInstanceDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
      // 搜索用的节点数据从数据源中获取
      if (res.data.allNodes) {
        searchDataSource.value.nodeId = res.data.allNodes
      }
    }
  }
  
  getDataSourceFunc()
  
  // 获取产品规格名称（只显示名称，不显示详细信息）
  const getSpecName = (specId) => {
    if (!specId || !dataSource.value.specId) return '-'
    const spec = dataSource.value.specId.find(item => item.value === specId)
    return spec ? spec.name : '-'
  }

  // 获取产品规格详细信息
  const getSpecInfo = (specId) => {
    if (!specId || !dataSource.value.specId) return null
    return dataSource.value.specId.find(item => item.value === specId)
  }



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

// 详情刷新状态与定时器（每30秒拉一次实例详情，读取表中指标）
const statsLoading = ref(false)
let statsTimer = null

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
  // 如果容器正在运行，定时刷新实例详情（30秒）
  if (detailForm.value.containerId && detailForm.value.containerStatus === 'running') {
    refreshDetail()
    statsTimer = setInterval(() => {
      refreshDetail()
    }, 30000)
  }
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

// 刷新实例详情（从表里拉取指标字段）
const refreshDetail = async () => {
  if (!detailForm.value.ID) return
  statsLoading.value = true
  try {
    const res = await findInstance({ ID: detailForm.value.ID })
    if (res.code === 0 && res.data) {
      detailForm.value = res.data
    }
  } finally {
    statsLoading.value = false
  }
}

// 格式化字节数
const formatBytes = (bytes) => {
  if (!bytes || bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i]
}

// 获取进度条颜色
const getProgressColor = (percentage) => {
  if (!percentage) return '#909399'
  if (percentage < 50) return '#67c23a'
  if (percentage < 80) return '#e6a23c'
  return '#f56c6c'
}

// 是否展示GPU显存（当规格有显卡且容器运行中）
const shouldShowGpuStats = computed(() => {
  const spec = getSpecInfo(detailForm.value?.specId)
  const gpuCount = (spec && (spec.gpu_count || spec.gpuCount)) ? (spec.gpu_count || spec.gpuCount) : 0
  return gpuCount > 0 && detailForm.value?.containerStatus === 'running'
})

// 显示 GPU 已用显存/总显存
const gpuUsedText = computed(() => {
  const total = containerStats.value.gpuMemorySizeGB || 0
  const rate = containerStats.value.gpuMemoryUsageRate || 0
  if (!total) return ''
  const used = (total * rate) / 100
  return `${used.toFixed(2)}GB / ${total}GB`
})

// 关闭详情弹窗
const closeDetailShow = () => {
  // 清除定时器
  if (statsTimer) {
    clearInterval(statsTimer)
    statsTimer = null
  }
  detailShow.value = false
  detailForm.value = {}
  containerStats.value = {
    cpuUsagePercent: 0,
    memoryUsage: 0,
    memoryLimit: 0,
    memoryUsagePercent: 0,
    networkRx: 0,
    networkTx: 0,
    blockRead: 0,
    blockWrite: 0,
    pids: 0,
    gpuMemorySizeGB: 0,
    gpuMemoryUsageRate: 0
  }
}

// ============== 新增实例弹窗相关 ==============
const createDialogVisible = ref(false)
const createForm = ref({
  imageId: undefined,
  specId: undefined,
  nodeId: undefined,
  name: '',
  remark: ''
})
const availableNodes = ref([])
const nodesLoading = ref(false)
const selectedSpec = ref(null)
const createFormRef = ref(null)

// 实例名称输入处理：过滤中文
const handleNameInput = (value) => {
  // 只保留字母、数字、横线和下划线
  createForm.value.name = value.replace(/[^a-zA-Z0-9_-]/g, '')
}

// 实例名称校验规则
const createFormRules = reactive({
  name: [
    {
      required: true,
      message: '请输入实例名称',
      trigger: ['input', 'blur'],
    },
    {
      pattern: /^[a-zA-Z0-9_-]+$/,
      message: '实例名称仅支持字母、数字、横线和下划线，不支持中文',
      trigger: ['input', 'blur'],
    },
  ],
})

// 打开新增实例弹窗
const openCreateDialog = () => {
  createForm.value = {
    imageId: undefined,
    specId: undefined,
    nodeId: undefined,
    name: '',
    remark: ''
  }
  availableNodes.value = []
  selectedSpec.value = null
  createDialogVisible.value = true
}

// 获取当前选择的镜像信息
const selectedImage = computed(() => {
  if (!createForm.value.imageId || !dataSource.value.imageId) return null
  return dataSource.value.imageId.find(item => item.value === createForm.value.imageId)
})

// 根据镜像的显存分割支持情况过滤显卡规格
const filteredSpecs = computed(() => {
  if (!selectedImage.value || !dataSource.value.specId) {
    return []
  }
  const imageSupportMemorySplit = selectedImage.value.supportMemorySplit || false
  
  return dataSource.value.specId.filter(spec => {
    const specSupportMemorySplit = spec.supportMemorySplit || false
    // 如果镜像不支持显存分割，只显示不支持显存分割的规格
    // 如果镜像支持显存分割，显示所有规格
    if (!imageSupportMemorySplit) {
      return !specSupportMemorySplit
    } else {
      return true // 显示所有规格
    }
  })
})

// 镜像变更处理
const onImageChange = () => {
  // 重置规格和节点选择
  createForm.value.specId = undefined
  createForm.value.nodeId = undefined
  selectedSpec.value = null
  availableNodes.value = []
}

// 选择产品规格
const selectSpec = async (spec) => {
  createForm.value.specId = spec.value
  createForm.value.nodeId = undefined // 重置节点选择
  selectedSpec.value = spec
  await fetchAvailableNodes(spec.value)
}

// 获取可用节点
const fetchAvailableNodes = async (specId) => {
  if (!specId) return
  nodesLoading.value = true
  try {
    const res = await getAvailableNodes({ specId })
    if (res.code === 0) {
      availableNodes.value = res.data || []
    } else {
      availableNodes.value = []
    }
  } catch (e) {
    availableNodes.value = []
  } finally {
    nodesLoading.value = false
  }
}

// 刷新节点列表
const refreshNodes = () => {
  if (createForm.value.specId) {
    fetchAvailableNodes(createForm.value.specId)
  }
}

// 选择节点
const selectNode = (row) => {
  if (row) {
    createForm.value.nodeId = row.id
  }
}

// 获取行样式
const getRowClassName = ({ row }) => {
  return createForm.value.nodeId === row.id ? 'selected-row' : ''
}

// 是否可以提交
const canSubmit = computed(() => {
  return createForm.value.imageId && 
         createForm.value.specId && 
         createForm.value.nodeId && 
         createForm.value.name
})

// 提交创建
const submitCreate = async () => {
  if (!canSubmit.value) {
    ElMessage.warning('请完善所有必填信息')
    return
  }
  
  // 表单校验
  if (!createFormRef.value) {
    return
  }
  
  try {
    await createFormRef.value.validate()
  } catch (error) {
    ElMessage.warning('请检查输入信息，实例名称不支持中文')
    return
  }
  
  btnLoading.value = true
  try {
    const res = await createInstance(createForm.value)
    if (res.code === 0) {
      ElMessage.success('创建成功')
      createDialogVisible.value = false
      getTableData()
    }
  } finally {
    btnLoading.value = false
  }
}

// 监听对话框关闭，重置表单
watch(createDialogVisible, (newVal) => {
  if (!newVal) {
    // 关闭时重置表单和校验状态
    createForm.value = {
      imageId: undefined,
      specId: undefined,
      nodeId: undefined,
      name: '',
      remark: ''
    }
    availableNodes.value = []
    selectedSpec.value = null
    if (createFormRef.value) {
      createFormRef.value.clearValidate()
    }
  }
})

// 监听规格变化，自动获取节点
watch(() => createForm.value.specId, (newVal) => {
  if (newVal) {
    fetchAvailableNodes(newVal)
  }
})

// ============== 容器操作相关 ==============

// 获取状态标签类型
const getStatusType = (status) => {
  const statusMap = {
    'running': 'success',
    'exited': 'danger',
    'created': 'info',
    'paused': 'warning',
    'restarting': 'warning',
    'removing': 'danger',
    'dead': 'danger',
    'creating': 'info',
    'failed': 'danger',
    'unknown': 'info'
  }
  return statusMap[status] || 'info'
}

// 处理容器操作
const handleContainerAction = async (command, row) => {
  switch (command) {
    case 'start':
      await handleStartContainer(row)
      break
    case 'stop':
      await handleStopContainer(row)
      break
    case 'restart':
      await handleRestartContainer(row)
      break
    case 'logs':
      openLogsDialog(row)
      break
    case 'terminal':
      openTerminalDialog(row)
      break
  }
}

// 启动容器
const handleStartContainer = async (row) => {
  try {
    const res = await startContainer({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('启动成功')
      getTableData()
    }
  } catch (e) {
    ElMessage.error('启动失败')
  }
}

// 停止容器
const handleStopContainer = async (row) => {
  ElMessageBox.confirm('确定要停止该容器吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await stopContainer({ ID: row.ID })
      if (res.code === 0) {
        ElMessage.success('停止成功')
        getTableData()
      }
    } catch (e) {
      ElMessage.error('停止失败')
    }
  })
}

// 重启容器
const handleRestartContainer = async (row) => {
  try {
    const res = await restartContainer({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('重启成功')
      getTableData()
    }
  } catch (e) {
    ElMessage.error('重启失败')
  }
}

// ============== 日志相关 ==============
const logsDialogVisible = ref(false)
const logsContent = ref('')
const logsLoading = ref(false)
const logsTail = ref('100')
const currentLogsInstance = ref(null)
const logsContainer = ref(null)
const logsAutoRefresh = ref(true) // 默认开启自动刷新
let logsRefreshTimer = null

const openLogsDialog = (row) => {
  currentLogsInstance.value = row
  logsDialogVisible.value = true
  logsAutoRefresh.value = true
  refreshLogs()
  startLogsAutoRefresh()
}

const refreshLogs = async () => {
  if (!currentLogsInstance.value) return
  logsLoading.value = true
  try {
    const res = await getContainerLogs({ 
      ID: currentLogsInstance.value.ID, 
      tail: logsTail.value 
    })
    if (res.code === 0) {
      logsContent.value = res.data || ''
      // 滚动到底部
      nextTick(() => {
        if (logsContainer.value) {
          logsContainer.value.scrollTop = logsContainer.value.scrollHeight
        }
      })
    }
  } catch (e) {
    logsContent.value = ''
  } finally {
    logsLoading.value = false
  }
}

// 开始自动刷新日志
const startLogsAutoRefresh = () => {
  stopLogsAutoRefresh()
  if (logsAutoRefresh.value) {
    logsRefreshTimer = setInterval(() => {
      refreshLogs()
    }, 1000)
  }
}

// 停止自动刷新日志
const stopLogsAutoRefresh = () => {
  if (logsRefreshTimer) {
    clearInterval(logsRefreshTimer)
    logsRefreshTimer = null
  }
}

// 切换自动刷新
const toggleLogsAutoRefresh = (val) => {
  if (val) {
    startLogsAutoRefresh()
  } else {
    stopLogsAutoRefresh()
  }
}

// ============== 终端相关 ==============
const terminalDialogVisible = ref(false)
const terminalContainer = ref(null)
const currentTerminalInstance = ref(null)
const terminalShell = ref('bash') // 默认使用bash
const terminalConnected = ref(false) // 响应式连接状态
let terminal = null
let terminalWs = null
let fitAddon = null
let resizeHandler = null

// WebSocket 状态常量，用于模板
const WS_OPEN = WebSocket.OPEN

const openTerminalDialog = async (row) => {
  currentTerminalInstance.value = row
  terminalShell.value = 'bash' // 重置为默认值
  terminalDialogVisible.value = true
  
  await nextTick()
  initTerminal()
}

// 显示SSH连接信息
const showSshConnectionInfo = async (row) => {
  // 从配置文件获取跳板机配置
  let serverIp = '公网IP' // 默认值
  let sshPort = 2026 // 默认值
  
  try {
    const res = await getJumpboxConfig()
    if (res.code === 0 && res.data) {
      if (res.data['server-ip']) {
        serverIp = res.data['server-ip']
      }
      if (res.data.port) {
        sshPort = res.data.port
      }
    }
  } catch (error) {
    console.warn('获取跳板机配置失败，使用默认值', error)
  }
  
  // 获取用户名（从实例的创建用户获取，如果没有则提示用户输入）
  const username = row.userName || 'your_username'
  
  const sshCommand = `ssh ${username}@${serverIp} -p ${sshPort}`
  
  ElMessageBox.alert(
    `<div style="text-align: left;">
      <p style="margin-bottom: 10px; font-weight: bold;">请使用以下SSH命令连接跳板机：</p>
      <p style="margin-bottom: 10px;">
        <code style="background: #f5f5f5; padding: 8px 12px; border-radius: 4px; display: inline-block; font-size: 14px; color: #409eff;">${sshCommand}</code>
      </p>
      <p style="margin-bottom: 5px; color: #666; font-size: 12px;">提示：</p>
      <ul style="margin: 5px 0; padding-left: 20px; color: #666; font-size: 12px;">
        <li>连接后输入系统密码进行认证</li>
        <li>认证成功后可以看到您创建的所有容器列表</li>
        <li>输入序号即可连接到对应的容器</li>
      </ul>
    </div>`,
    'SSH连接信息',
    {
      dangerouslyUseHTMLString: true,
      confirmButtonText: '复制命令',
      cancelButtonText: '关闭',
      showCancelButton: true,
      beforeClose: (action, instance, done) => {
        if (action === 'confirm') {
          // 复制命令到剪贴板
          navigator.clipboard.writeText(sshCommand).then(() => {
            ElMessage.success('SSH命令已复制到剪贴板')
          }).catch(() => {
            ElMessage.error('复制失败，请手动复制')
          })
        }
        done()
      }
    }
  )
}

const initTerminal = async () => {
  if (!terminalContainer.value || !currentTerminalInstance.value) return
  
  // 如果终端已存在，先清理
  if (terminal) {
    terminal.dispose()
    terminal = null
  }
  
  // 动态导入xterm
  const { Terminal } = await import('xterm')
  const { FitAddon } = await import('xterm-addon-fit')
  await import('xterm/css/xterm.css')
  
  terminal = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    fontFamily: 'Menlo, Monaco, "Courier New", monospace',
    theme: {
      background: '#1e1e1e',
      foreground: '#d4d4d4',
      cursor: '#d4d4d4'
    }
  })
  
  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.open(terminalContainer.value)
  fitAddon.fit()
  
  // 连接WebSocket
  connectTerminal()
  
  // 处理输入
  terminal.onData((data) => {
    if (terminalWs && terminalConnected.value) {
      const msg = JSON.stringify({
        type: 'input',
        data: data
      })
      terminalWs.send(msg)
    }
  })
  
  // 处理大小变化
  terminal.onResize(({ cols, rows }) => {
    if (terminalWs && terminalConnected.value) {
      const msg = JSON.stringify({
        type: 'resize',
        cols: cols,
        rows: rows
      })
      terminalWs.send(msg)
    }
  })
  
  // 窗口大小变化时调整终端大小
  resizeHandler = () => {
    if (fitAddon) {
      fitAddon.fit()
      // 同步发送大小变化
      if (terminalWs && terminalConnected.value && terminal) {
        const msg = JSON.stringify({
          type: 'resize',
          cols: terminal.cols,
          rows: terminal.rows
        })
        terminalWs.send(msg)
      }
    }
  }
  window.addEventListener('resize', resizeHandler)
}

const connectTerminal = () => {
  // 如果已有连接，先关闭
  if (terminalWs) {
    terminalWs.close()
    terminalWs = null
  }
  
  // 重置连接状态
  terminalConnected.value = false
  
  if (!currentTerminalInstance.value) return
  
  // 连接WebSocket
  const wsUrl = getTerminalWsUrl(currentTerminalInstance.value.ID, terminalShell.value)
  terminalWs = new WebSocket(wsUrl)
  
  terminalWs.onopen = () => {
    // 更新连接状态
    terminalConnected.value = true
    if (terminal) {
      terminal.writeln(`\r\n连接成功，使用 ${terminalShell.value === 'bash' ? '/bin/bash' : '/bin/sh'}`)
      // 发送初始大小
      if (terminal.cols && terminal.rows) {
        const msg = JSON.stringify({
          type: 'resize',
          cols: terminal.cols,
          rows: terminal.rows
        })
        terminalWs.send(msg)
      }
    }
  }
  
  terminalWs.onmessage = (event) => {
    if (terminal) {
      if (event.data instanceof Blob) {
        event.data.text().then(text => {
          terminal.write(text)
        })
      } else if (typeof event.data === 'string') {
        // 如果是文本消息，可能是错误信息
        try {
          const data = JSON.parse(event.data)
          if (data.type === 'error') {
            terminal.writeln('\r\n错误: ' + data.message)
          }
        } catch (e) {
          // 不是JSON，直接写入
          terminal.write(event.data)
        }
      } else {
        // ArrayBuffer或其他二进制数据
        const reader = new FileReader()
        reader.onload = () => {
          terminal.write(reader.result)
        }
        reader.readAsText(event.data)
      }
    }
  }
  
  terminalWs.onerror = (error) => {
    // 更新连接状态
    terminalConnected.value = false
    if (terminal) {
      terminal.writeln('\r\n连接错误: ' + (error.message || '未知错误'))
    }
    console.error('WebSocket error:', error)
  }
  
  terminalWs.onclose = (event) => {
    // 更新连接状态
    terminalConnected.value = false
    if (terminal) {
      terminal.writeln('\r\n连接已关闭')
    }
    terminalWs = null
  }
}

// 重新连接终端（切换shell时）
const reconnectTerminal = () => {
  if (terminal) {
    terminal.writeln(`\r\n正在切换到 ${terminalShell.value === 'bash' ? '/bin/bash' : '/bin/sh'}...`)
  }
  connectTerminal()
}

// 处理终端对话框关闭（在对话框完全关闭后触发）
const handleTerminalClose = () => {
  // 清理所有资源
  closeTerminal()
}

const closeTerminal = () => {
  // 重置连接状态
  terminalConnected.value = false
  // 清理WebSocket连接
  if (terminalWs) {
    try {
      terminalWs.close()
    } catch (e) {
      console.error('关闭WebSocket失败:', e)
    }
    terminalWs = null
  }
  // 清理终端实例
  if (terminal) {
    try {
      terminal.dispose()
    } catch (e) {
      console.error('清理终端失败:', e)
    }
    terminal = null
  }
  // 清理窗口大小监听
  if (resizeHandler) {
    window.removeEventListener('resize', resizeHandler)
    resizeHandler = null
  }
  fitAddon = null
  currentTerminalInstance.value = null
}

// 组件卸载时清理
onUnmounted(() => {
  closeTerminal()
  stopLogsAutoRefresh()
})

</script>

<style scoped>
.create-instance-dialog :deep(.el-dialog__body) {
  padding: 0;
}

.create-instance-content {
  max-height: 70vh;
  overflow-y: auto;
  padding: 20px;
}

.step-section {
  margin-bottom: 24px;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  overflow: hidden;
}

.step-header {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  background: #f5f7fa;
  border-bottom: 1px solid #e4e7ed;
}

.step-icon {
  font-size: 20px;
  margin-right: 8px;
  color: #c0c4cc;
}

.step-icon.completed {
  color: #67c23a;
}

.step-title {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.step-content {
  padding: 16px;
}

/* 规格卡片样式 */
.spec-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.spec-card {
  border: 2px solid #e4e7ed;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.3s;
}

.spec-card:hover {
  border-color: #409eff;
  box-shadow: 0 2px 12px rgba(64, 158, 255, 0.2);
}

.spec-card.selected {
  border-color: #67c23a;
  background: #f0f9eb;
}

.spec-card-header {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.gpu-icon {
  font-size: 24px;
  color: #67c23a;
  margin-right: 8px;
}

.spec-name {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.spec-card-body {
  font-size: 13px;
}

.spec-info {
  margin-bottom: 8px;
}

.spec-gpu {
  color: #409eff;
  font-weight: 500;
}

.spec-details {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  color: #909399;
  margin-bottom: 8px;
}

.spec-details span {
  background: #f4f4f5;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.spec-price {
  margin-top: 8px;
  text-align: right;
}

.spec-price .price {
  font-size: 18px;
  font-weight: 600;
  color: #f56c6c;
}

.spec-price .unit {
  font-size: 12px;
  color: #909399;
}

/* 空提示 */
.empty-tip,
.loading-tip {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px;
  color: #909399;
}

.empty-tip .el-icon,
.loading-tip .el-icon {
  margin-right: 8px;
  font-size: 20px;
}

/* 可用数量高亮 */
.available-count {
  color: #67c23a;
  font-weight: 600;
}

.price-text {
  color: #f56c6c;
  font-weight: 500;
}

/* 选中行样式 */
:deep(.selected-row) {
  background-color: #ecf5ff !important;
}

:deep(.el-table__body tr.current-row > td) {
  background-color: #ecf5ff !important;
}

/* 底部 */
.dialog-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-top: 1px solid #e4e7ed;
}

.summary {
  font-size: 14px;
  color: #606266;
}

.total-price {
  font-size: 24px;
  font-weight: 600;
  color: #f56c6c;
}

.price-unit {
  font-size: 14px;
  color: #909399;
}

.actions {
  display: flex;
  gap: 12px;
}

/* 容器ID样式 */
.container-id {
  font-family: 'Courier New', monospace;
  font-size: 12px;
  color: #909399;
}

/* 日志弹窗样式 */
.logs-dialog :deep(.el-dialog__body) {
  padding: 0;
}

.logs-toolbar {
  display: flex;
  gap: 12px;
  padding: 12px 20px;
  border-bottom: 1px solid #e4e7ed;
  background: #f5f7fa;
}

.logs-content {
  height: 500px;
  overflow-y: auto;
  padding: 16px;
  background: #1e1e1e;
}

.logs-content pre {
  margin: 0;
  font-family: 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.5;
  color: #d4d4d4;
  white-space: pre-wrap;
  word-break: break-all;
}

.logs-loading {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #d4d4d4;
}

.logs-loading .el-icon {
  margin-right: 8px;
}

/* 终端弹窗样式 */
.terminal-dialog :deep(.el-dialog__body) {
  padding: 0;
}

.terminal-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.toolbar-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.shell-select-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.15);
  padding: 6px 12px;
  border-radius: 8px;
  backdrop-filter: blur(10px);
}

.toolbar-icon {
  font-size: 18px;
  color: #fff;
}

.toolbar-label {
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  white-space: nowrap;
}

.shell-select {
  min-width: 200px;
}

.shell-select :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.95) !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  border-radius: 6px;
}

.shell-select :deep(.el-input__inner) {
  color: #303133;
  font-weight: 500;
}

.shell-select :deep(.el-input__suffix) {
  color: #667eea;
}

.shell-select :deep(.el-input.is-disabled .el-input__wrapper) {
  background: rgba(255, 255, 255, 0.5) !important;
  cursor: not-allowed;
}

.shell-select :deep(.el-input.is-disabled .el-input__inner) {
  color: rgba(255, 255, 255, 0.8);
}

/* Shell选项样式 */
.shell-option {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 4px 0;
}

.shell-icon {
  font-size: 16px;
  color: #667eea;
  flex-shrink: 0;
}

.shell-name {
  font-weight: 600;
  color: #303133;
  font-size: 14px;
  min-width: 90px;
}

.shell-desc {
  font-size: 12px;
  color: #909399;
  margin-left: auto;
}

/* 下拉选项样式 */
.shell-select :deep(.el-select-dropdown__item) {
  padding: 10px 16px;
  height: auto;
}

.shell-select :deep(.el-select-dropdown__item:hover) {
  background: #f5f7fa;
}

.shell-select :deep(.el-select-dropdown__item.is-selected) {
  background: #ecf5ff;
  color: #409eff;
}

.shell-select :deep(.el-select-dropdown__item.is-selected .shell-name) {
  color: #409eff;
}

.toolbar-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.terminal-container {
  height: 500px;
  background: #1e1e1e;
  border-radius: 0 0 4px 4px;
}

/* 容器统计信息样式 */
.stats-container {
  padding: 8px 0;
}

.stats-item {
  margin-bottom: 8px;
}

.stats-label {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 14px;
  color: #606266;
}

.stats-value {
  font-weight: 600;
  color: #303133;
  font-size: 14px;
}

.stats-detail {
  margin-top: 4px;
  font-size: 12px;
  color: #909399;
}
.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
  line-height: 1.5;
}
</style>
