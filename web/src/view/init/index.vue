<template>
  <div
    class="rounded-lg flex items-center justify-evenly w-full h-full relative md:w-screen md:h-screen md:bg-[#194bfb] overflow-hidden"
  >
    <div
      class="rounded-md w-full h-full flex items-center justify-center overflow-hidden"
    >
      <div
        class="oblique h-[130%] w-3/5 bg-white dark:bg-slate-900 transform -rotate-12 absolute -ml-80"
      />
      <div
        v-if="!page.showForm"
        :class="[page.showReadme ? 'slide-out-right' : 'slide-in-fwd-top']"
      >
        <div class="text-lg">
          <div
            class="font-sans text-4xl font-bold text-center mb-4 dark:text-white"
          >
            天启算力平台
          </div>
          <p class="text-gray-600 dark:text-gray-300 mb-2">初始化须知</p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            1.您需有一定的 VUE 和 GOLANG 基础
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            2.请确认是否了解后续的配置流程
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            3.如果您使用 MySQL 数据库，请确认数据库引擎为<span
              class="text-red-600 font-bold text-3xl ml-2"
              >InnoDB</span
            >
          </p>
          <p class="flex items-center justify-between mt-8">
            <el-button type="primary" size="large" @click="showNext">
              开始初始化
            </el-button>
          </p>
        </div>
      </div>
      <div
        v-if="page.showForm"
        :class="[page.showForm ? 'slide-in-left' : 'slide-out-right']"
        class="w-96"
      >
        <el-form ref="formRef" :model="form" label-width="100px" size="large">
          <el-form-item label="管理员密码">
            <el-input
              v-model="form.adminPassword"
              placeholder="admin账号的默认密码"
            ></el-input>
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-select
              v-model="form.dbType"
              placeholder="请选择"
              class="w-full"
              @change="changeDB"
            >
              <el-option key="mysql" label="mysql" value="mysql" />
              <el-option key="pgsql" label="pgsql" value="pgsql" />
              <el-option key="oracle" label="oracle" value="oracle" />
              <el-option key="mssql" label="mssql" value="mssql" />
              <el-option key="sqlite" label="sqlite" value="sqlite" />
            </el-select>
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="host">
            <el-input v-model="form.host" placeholder="请输入数据库链接" />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="port">
            <el-input v-model="form.port" placeholder="请输入数据库端口" />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="userName">
            <el-input
              v-model="form.userName"
              placeholder="请输入数据库用户名"
            />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="password">
            <el-input
              v-model="form.password"
              placeholder="请输入数据库密码（没有则为空）"
            />
          </el-form-item>
          <el-form-item label="dbName">
            <el-input v-model="form.dbName" placeholder="请输入数据库名称" />
          </el-form-item>
          <el-form-item v-if="form.dbType === 'sqlite'" label="dbPath">
            <el-input
              v-model="form.dbPath"
              placeholder="请输入sqlite数据库文件存放路径"
            />
          </el-form-item>
          <el-form-item v-if="form.dbType === 'pgsql'" label="template">
            <el-input
              v-model="form.template"
              placeholder="请输入postgresql指定template"
            />
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">立即初始化</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </div>

    <div class="hidden md:flex w-1/2 h-full float-right bg-[#194bfb] flex-col items-center justify-center px-12">
      <!-- Docker Logo SVG -->
      <div class="mb-8">
        <svg viewBox="0 0 640 512" class="w-32 h-32 text-white" fill="currentColor">
          <path d="M349.9 236.3h-66.1v-59.4h66.1v59.4zm0-204.3h-66.1v60.7h66.1V32zm78.2 144.8H362v59.4h66.1v-59.4zm-156.3-72.1h-66.1v60.1h66.1v-60.1zm78.1 0h-66.1v60.1h66.1v-60.1zm276.8 100c-14.4-9.7-47.6-13.2-73.1-8.4-3.3-24-16.7-44.9-41.1-63.7l-14-9.3-9.3 14c-18.4 27.8-23.4 73.6-3.7 103.8-8.7 4.7-25.8 11.1-48.4 10.7H2.4c-8.7 50.8 5.8 116.8 44 162.1 37.1 43.9 92.7 66.2 165.4 66.2 157.4 0 273.9-72.5 328.4-204.2 21.4.4 67.6.1 91.3-45.2 1.5-2.5 6.6-13.2 8.5-17.1l-13.3-8.9zm-511.1-27.9h-66v59.4h66.1v-59.4zm78.1 0h-66.1v59.4h66.1v-59.4zm78.1 0h-66.1v59.4h66.1v-59.4zm-78.1-72.1h-66.1v60.1h66.1v-60.1z"/>
        </svg>
      </div>
      
      <!-- 标题 -->
      <h2 class="text-white text-3xl font-bold mb-4 text-center">Docker GPU 算力管理</h2>
      <p class="text-blue-200 text-center mb-10 text-lg">高效、灵活、智能的容器化 GPU 资源调度平台</p>
      
      <!-- 特性列表 -->
      <div class="space-y-6 w-full max-w-md">
        <div class="flex items-center text-white">
          <div class="w-12 h-12 rounded-full bg-white/20 flex items-center justify-center mr-4 flex-shrink-0">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-7 14l-5-5 1.41-1.41L12 14.17l4.59-4.59L18 11l-6 6z"/>
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-lg">GPU 资源监控</h3>
            <p class="text-blue-200 text-sm">实时监控 GPU 使用率、显存、温度等关键指标</p>
          </div>
        </div>
        
        <div class="flex items-center text-white">
          <div class="w-12 h-12 rounded-full bg-white/20 flex items-center justify-center mr-4 flex-shrink-0">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-lg">容器化部署</h3>
            <p class="text-blue-200 text-sm">基于 Docker 的一键部署，快速启动 AI 训练环境</p>
          </div>
        </div>
        
        <div class="flex items-center text-white">
          <div class="w-12 h-12 rounded-full bg-white/20 flex items-center justify-center mr-4 flex-shrink-0">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </div>
          <div>
            <h3 class="font-semibold text-lg">智能调度</h3>
            <p class="text-blue-200 text-sm">自动分配空闲 GPU 资源，最大化算力利用率</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  // @ts-ignore
  import { initDB } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'

  defineOptions({
    name: 'Init'
  })

  const router = useRouter()

  const page = reactive({
    showReadme: false,
    showForm: false
  })

  const showNext = () => {
    page.showReadme = false
    setTimeout(() => {
      page.showForm = true
    }, 20)
  }

  const out = ref(false)

  const form = reactive({
    adminPassword: '123456',
    dbType: 'mysql',
    host: '127.0.0.1',
    port: '3306',
    userName: 'root',
    password: '',
    dbName: 'tianqi',
    dbPath: ''
  })

  const changeDB = (val) => {
    switch (val) {
      case 'mysql':
        Object.assign(form, {
          adminPassword: '123456',
          reAdminPassword: '',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'tianqi',
          dbPath: ''
        })
        break
      case 'pgsql':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'pgsql',
          host: '127.0.0.1',
          port: '5432',
          userName: 'postgres',
          password: '',
          dbName: 'tianqi',
          dbPath: '',
          template: 'template0'
        })
        break
      case 'oracle':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'oracle',
          host: '127.0.0.1',
          port: '1521',
          userName: 'oracle',
          password: '',
          dbName: 'tianqi',
          dbPath: ''
        })
        break
      case 'mssql':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'mssql',
          host: '127.0.0.1',
          port: '1433',
          userName: 'mssql',
          password: '',
          dbName: 'tianqi',
          dbPath: ''
        })
        break
      case 'sqlite':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'sqlite',
          host: '',
          port: '',
          userName: '',
          password: '',
          dbName: 'tianqi',
          dbPath: ''
        })
        break
      default:
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'tianqi',
          dbPath: ''
        })
    }
  }
  const onSubmit = async () => {
    if (form.adminPassword.length < 6) {
      ElMessage({
        type: 'error',
        message: '密码长度不能小于6位'
      })
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在初始化数据库，请稍候',
      spinner: 'loading',
      background: 'rgba(0, 0, 0, 0.7)'
    })
    try {
      const res = await initDB(form)
      if (res.code === 0) {
        out.value = true
        ElMessage({
          type: 'success',
          message: res.msg
        })
        router.push({ name: 'Login' })
      }
      loading.close()
    } catch (_) {
      loading.close()
    }
  }
</script>

<style lang="scss" scoped>
  .slide-in-fwd-top {
    -webkit-animation: slide-in-fwd-top 0.4s
      cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-fwd-top 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }
  .slide-out-right {
    -webkit-animation: slide-out-right 0.5s
      cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
    animation: slide-out-right 0.5s cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
  }
  .slide-in-left {
    -webkit-animation: slide-in-left 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94)
      both;
    animation: slide-in-left 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }
  @-webkit-keyframes slide-in-fwd-top {
    0% {
      transform: translateZ(-1400px) translateY(-800px);
      opacity: 0;
    }
    100% {
      transform: translateZ(0) translateY(0);
      opacity: 1;
    }
  }
  @keyframes slide-in-fwd-top {
    0% {
      transform: translateZ(-1400px) translateY(-800px);
      opacity: 0;
    }
    100% {
      transform: translateZ(0) translateY(0);
      opacity: 1;
    }
  }
  @-webkit-keyframes slide-out-right {
    0% {
      transform: translateX(0);
      opacity: 1;
    }
    100% {
      transform: translateX(1000px);
      opacity: 0;
    }
  }
  @keyframes slide-out-right {
    0% {
      transform: translateX(0);
      opacity: 1;
    }
    100% {
      transform: translateX(1000px);
      opacity: 0;
    }
  }
  @-webkit-keyframes slide-in-left {
    0% {
      transform: translateX(-1000px);
      opacity: 0;
    }
    100% {
      transform: translateX(0);
      opacity: 1;
    }
  }
  @keyframes slide-in-left {
    0% {
      transform: translateX(-1000px);
      opacity: 0;
    }
    100% {
      transform: translateX(0);
      opacity: 1;
    }
  }
  @media (max-width: 750px) {
    .form {
      width: 94vw !important;
      padding: 0;
    }
  }
</style>
