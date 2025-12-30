<template>
  <div class="mt-2">
    <div class="flex flex-col md:flex-row gap-4">
      <div class="w-full md:w-1/2">
        <el-card class="min-w-96">
          <template #header>
            <el-divider>天启算力平台</el-divider>
          </template>
          <div>
            <div class="w-full flex items-center justify-center mb-6">
              <!-- Docker GPU Logo -->
              <div class="w-32 h-32 rounded-full bg-blue-500 flex items-center justify-center">
                <svg viewBox="0 0 640 512" class="w-16 h-16 text-white" fill="currentColor">
                  <path d="M349.9 236.3h-66.1v-59.4h66.1v59.4zm0-204.3h-66.1v60.7h66.1V32zm78.2 144.8H362v59.4h66.1v-59.4zm-156.3-72.1h-66.1v60.1h66.1v-60.1zm78.1 0h-66.1v60.1h66.1v-60.1zm276.8 100c-14.4-9.7-47.6-13.2-73.1-8.4-3.3-24-16.7-44.9-41.1-63.7l-14-9.3-9.3 14c-18.4 27.8-23.4 73.6-3.7 103.8-8.7 4.7-25.8 11.1-48.4 10.7H2.4c-8.7 50.8 5.8 116.8 44 162.1 37.1 43.9 92.7 66.2 165.4 66.2 157.4 0 273.9-72.5 328.4-204.2 21.4.4 67.6.1 91.3-45.2 1.5-2.5 6.6-13.2 8.5-17.1l-13.3-8.9zm-511.1-27.9h-66v59.4h66.1v-59.4zm78.1 0h-66.1v59.4h66.1v-59.4zm78.1 0h-66.1v59.4h66.1v-59.4zm-78.1-72.1h-66.1v60.1h66.1v-60.1z"/>
                </svg>
              </div>
            </div>
            <div class="text-center">
              <h2 class="text-2xl font-bold mb-4">Docker GPU 算力管理平台</h2>
              <p class="text-gray-600 dark:text-gray-300 mb-2">高效、灵活、智能的容器化 GPU 资源调度平台</p>
            </div>
          </div>
        </el-card>
        <el-card class="min-w-96 mt-5">
          <template #header>
            <div>平台特性</div>
          </template>
          <div class="space-y-4">
            <div class="flex items-center">
              <div class="w-10 h-10 rounded-full bg-green-500 flex items-center justify-center mr-4">
                <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-7 14l-5-5 1.41-1.41L12 14.17l4.59-4.59L18 11l-6 6z"/>
                </svg>
              </div>
              <div>
                <h3 class="font-semibold">GPU 资源监控</h3>
                <p class="text-gray-500 text-sm">实时监控 GPU 使用率、显存、温度等关键指标</p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="w-10 h-10 rounded-full bg-blue-500 flex items-center justify-center mr-4">
                <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>
                </svg>
              </div>
              <div>
                <h3 class="font-semibold">容器化部署</h3>
                <p class="text-gray-500 text-sm">基于 Docker 的一键部署，快速启动 AI 训练环境</p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="w-10 h-10 rounded-full bg-purple-500 flex items-center justify-center mr-4">
                <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>
              </div>
              <div>
                <h3 class="font-semibold">智能调度</h3>
                <p class="text-gray-500 text-sm">自动分配空闲 GPU 资源，最大化算力利用率</p>
              </div>
            </div>
          </div>
        </el-card>
      </div>
      <div class="w-full md:w-1/2">
        <el-card>
          <template #header>
            <div>系统信息</div>
          </template>
          <div class="space-y-4">
            <div class="flex justify-between items-center py-2 border-b dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-300">平台名称</span>
              <span class="font-semibold">天启算力平台</span>
            </div>
            <div class="flex justify-between items-center py-2 border-b dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-300">前端框架</span>
              <span class="font-semibold">Vue 3 + Element Plus</span>
            </div>
            <div class="flex justify-between items-center py-2 border-b dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-300">后端框架</span>
              <span class="font-semibold">Gin (Golang)</span>
            </div>
            <div class="flex justify-between items-center py-2 border-b dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-300">容器技术</span>
              <span class="font-semibold">Docker</span>
            </div>
            <div class="flex justify-between items-center py-2">
              <span class="text-gray-600 dark:text-gray-300">GPU 支持</span>
              <span class="font-semibold">NVIDIA CUDA</span>
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
  defineOptions({
    name: 'About'
  })
</script>

<style scoped>
  .org-img {
    height: 150px;
    width: 150px;
  }

  .dom-center {
    margin-left: 50%;
    transform: translateX(-50%);
  }
</style>
