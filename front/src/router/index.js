import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'nodes',
      component: () => import('../views/NodesView.vue')
    },
    {
      path: '/nodes',
      name: 'nodes',
      component: () => import('../views/NodesView.vue')
    },
    {
      path: '/vms',
      name: 'vms',
      component: () => import('../views/VMsView.vue')
    },
    {
      path: '/containers',
      name: 'containers',
      component: () => import('../views/ContainersView.vue')
    },
    {
      path: '/files',
      name: 'files',
      component: () => import('../views/FilesView.vue')
    },
    {
      path: '/k8s',
      name: 'k8s',
      component: () => import('../views/K8sView.vue')
    }
  ]
})

export default router
