import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/scanUpload',
    name: 'ScanUpload',
    meta: {
      title: '扫码上传',
      client: true
    },
    component: () => import('@/view/example/upload/scanUpload.vue')
  },

  {
    path: '/pcdn',
    name: 'Pcdn',
    component: () => import('@/view/layout/index.vue'),
    redirect: '/pcdn/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'pcdnDashboard',
        component: () => import('@/view/pcdn/dashboard.vue')
      },
      {
        path: 'node-management',
        name: 'pcdnNodeManagement',
        component: () => import('@/view/pcdn/nodeManagement.vue')
      },
      {
        path: 'policy-management',
        name: 'pcdnPolicyManagement',
        component: () => import('@/view/pcdn/policyManagement.vue')
      },
      {
        path: 'dispatch-tasks',
        name: 'pcdnDispatchTasks',
        component: () => import('@/view/pcdn/dispatchTasks.vue')
      }
    ]
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
