import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '@/views/Home.vue'

export default createRouter({
    history: createWebHashHistory(),
    routes: [
      {
        path: '/',
        name: 'Home',
        component: Home,
      },
      {
        path: '/show/:id',
        name: 'Button',
        component: () => import('@/views/Button.vue')
      },
      {
        path: '/show=true/:id',
        name: 'Show',
        component: () => import('@/views/Show.vue')
      }
    ],
  })
  

