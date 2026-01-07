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
        {
          path: 'kiosk',
          name: 'admin-kiosk',
          component: () => import('../pages/admin/KioskManagement.vue'),
        },
        {
          path: 'kennel-displays',
          name: 'admin-kennel-displays',
          component: () => import('../pages/admin/KennelDisplays.vue'),
        },
        {
          path: 'event-displays',
          name: 'admin-event-displays',
          component: () => import('../pages/admin/EventDisplays.vue'),
        },
        {
          path: 'time-logs',
          name: 'admin-time-logs',
          component: () => import('../pages/admin/TimeLogs.vue'),
        },
        {
          path: 'donations',
          name: 'admin-donations',
          component: () => import('../pages/admin/Donations.vue'),
        },
        {
          path: 'inventory',
          name: 'admin-inventory',
          component: () => import('../pages/admin/Inventory.vue'),
        },
        {
          path: 'marketing',
          name: 'admin-marketing',
          component: () => import('../pages/admin/Marketing.vue'),
        },
        {
          path: 'transport',
          name: 'admin-transport',
          component: () => import('../pages/admin/Transport.vue'),
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
