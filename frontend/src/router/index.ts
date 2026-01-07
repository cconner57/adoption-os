import { createRouter, createWebHistory } from 'vue-router'
import {
  About,
  Adopt,
  Donate,
  Home,
  SurrenderPet,
  Volunteer,
  PetAdoption,
  NotFound,
  Login,
} from '../pages/index.ts'
import { useMetrics } from '../composables/useMetrics'
import { nextTick } from 'vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Home,
    },
    {
      path: '/login',
      component: Login,
    },
    {
      path: '/admin',
      component: () => import('../layouts/AdminLayout.vue'),
      children: [
        {
          path: '',
          name: 'admin-overview',
          component: () => import('../pages/admin/DashboardHome.vue'),
        },
        {
          path: 'calendar',
          name: 'admin-calendar',
          component: () => import('../pages/admin/Calendar.vue'),
        },
        {
          path: 'pets',
          name: 'admin-pets',
          component: () => import('../pages/admin/Pets.vue'),
        },
        {
          path: 'pet-health',
          name: 'admin-pet-health',
          component: () => import('../pages/admin/PetHealth.vue'),
        },
        {
          path: 'volunteers',
          name: 'admin-volunteers',
          component: () => import('../pages/admin/Volunteers.vue'),
        },
        {
          path: 'messages',
          name: 'admin-messages',
          component: () => import('../pages/admin/Messages.vue'),
        },
        {
          path: 'settings',
          name: 'admin-settings',
          component: () => import('../pages/admin/Settings.vue'),
        },
      ],
    },
    {
      path: '/about',
      component: About,
    },
    { path: '/adopt', name: 'adopt-list', component: Adopt },
    {
      path: '/adopt/:id',
      name: 'adopt-pet',
      component: Adopt,
      props: true,
    },
    {
      path: '/donate',
      component: Donate,
    },
    {
      path: '/volunteer',
      component: Volunteer,
    },
    {
      path: '/surrender',
      component: SurrenderPet,
    },
    {
      path: '/pet-adoption/:id',
      component: PetAdoption,
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: NotFound,
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition
    return { top: 0 }
  },
})

interface ViewTransition {
  finished: Promise<void>
  ready: Promise<void>
  updateCallbackDone: Promise<void>
  skipTransition(): void
}

type ViewTransitionUpdateCallback = () => Promise<void> | void
type StartViewTransitionOptions = Record<string, unknown>

interface ViewTransitionDocument extends Document {
  startViewTransition(
    callback?: ViewTransitionUpdateCallback,
    options?: StartViewTransitionOptions,
  ): ViewTransition
}

router.beforeResolve((to, from, next) => {
  const doc = document as unknown as ViewTransitionDocument

  if (!doc.startViewTransition) {
    next()
    return
  }

  doc.startViewTransition(async () => {
    next()
    await nextTick()
  })
})

router.afterEach((to) => {
  const { submitMetric } = useMetrics()
  submitMetric('page_view', { path: to.fullPath })
})

export default router
