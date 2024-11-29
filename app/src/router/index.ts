import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../views/LogoutView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    {
      path: '/analysisCenter',
      name: 'main',
      component: () => import('../views/MainView.vue')
    },
    {
      path: '/contactUs',
      name: 'contactUs',
      component: () => import('../views/ContactUsView.vue')
    },
    {
      path: '/statistics',
      name: 'statistics',
      component: () => import('../views/StatisticsView.vue')
    }
  ]
})

export default router

router.beforeEach((to, from, next) => {
  const { isLoggedIn } = useAuthStore()
  if (to.name == 'logout' && !isLoggedIn) {
    next({ name: 'home' })
  } else if (to.name == 'main' && !isLoggedIn) {
    next({ name: 'home' })
  } else {
    next()
  }
})
