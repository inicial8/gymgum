import {storeToRefs} from "pinia";
import {useMainStore} from "@/stores/main";
import {createRouter, createWebHistory, Router} from "vue-router";

const NotFound: {template: string} = {template: '<h2>Page Not Found</h2>'}

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    beforeEnter: (to: any, from: any) => {
      const {isAuthenticated} = storeToRefs(useMainStore())
      if (to.name === 'Login' && isAuthenticated.value) return {name: from.name}
    },
  },
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue')
      },
    ],
  },
  {
    path: '/clients',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Clients',
        component: () => import('@/views/Clients.vue'),
      },
    ],
    beforeEnter: (to: any, from: any) => {
      const {user} = storeToRefs(useMainStore())
      if (!(user.value.role === 'admin')) return {name: from.name}
    },
  },
  {
    path: '/account',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Account',
        component: () => import('@/views/Account.vue'),
      },
    ],
    beforeEnter: (to: any, from: any) => {
      const {user} = storeToRefs(useMainStore())
      if (!(user.value.role === 'client' || user.value.role === 'admin')) return {name: from.name}
    },
  },
  {path: '/:pathMatch(.*)*', name: 'not-found', component: NotFound},
]

const router: Router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

router.beforeEach((to, from) => {
  const {isAuthenticated} = storeToRefs(useMainStore())
  if (to.name !== 'Login' && !isAuthenticated.value) return {name: 'Login'}
})

export default router
