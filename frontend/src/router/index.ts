import { createRouter, createWebHistory } from 'vue-router'
// Pages are now lazy-loaded inline to support code splitting
const KioskLayout = () => import('../layouts/KioskLayout.vue')
const KioskHome = () => import('../pages/kiosk/KioskHome.vue')
const KioskDailyCare = () => import('../pages/kiosk/KioskDailyCare.vue')
const KioskPetList = () => import('../pages/kiosk/KioskPetList.vue')
const KioskVetSchedule = () => import('../pages/kiosk/KioskVetSchedule.vue')

import { nextTick } from 'vue'

import { useMetrics } from '../composables/useMetrics'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('../pages/Home.vue'),
    },
    {
      path: '/login',
      component: () => import('../pages/Login.vue'),
      meta: { hideNavbar: true },
    },
    {
      path: '/kiosk',
      meta: { hideNavbar: true },
      component: KioskLayout,
      children: [
        { path: '', name: 'kiosk-home', component: KioskHome },
        { path: 'daily-care', name: 'kiosk-daily-care', component: KioskDailyCare },
        { path: 'pets', name: 'kiosk-pets', component: KioskPetList },
        { path: 'vet', name: 'kiosk-vet', component: KioskVetSchedule },
      ],
    },
    {
      path: '/admin',
      meta: { hideNavbar: true },
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
        {
          path: 'transport',
          name: 'admin-transport',
          component: () => import('../pages/admin/Transport.vue'),
        },
        {
          path: 'applications',
          name: 'admin-applications',
          component: () => import('../pages/admin/Applications.vue'),
        },
      ],
    },
    {
      path: '/about',
      component: () => import('../pages/About.vue'),
    },
    { path: '/adopt', name: 'adopt-list', component: () => import('../pages/Adopt.vue') },
    {
      path: '/adopt/:id',
      name: 'adopt-pet',
      component: () => import('../pages/Adopt.vue'),
      props: true,
    },
    {
      path: '/donate',
      component: () => import('../pages/Donate.vue'),
    },
    {
      path: '/volunteer',
      component: () => import('../pages/Volunteer.vue'),
    },
    {
      path: '/surrender',
      component: () => import('../pages/SurrenderPet.vue'),
    },
    {
      path: '/pet-adoption/:id',
      component: () => import('../pages/PetAdoption.vue'),
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('../pages/NotFound.vue'),
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
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  types?: any
}

/* eslint-disable no-unused-vars */
interface ViewTransitionTypeSet {
  add(type: string): void
  delete(type: string): boolean
  has(type: string): boolean
  clear(): void
  readonly size: number
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  forEach(callbackfn: (value: string, value2: string, set: Set<string>) => void, thisArg?: any): void
  entries(): IterableIterator<[string, string]>
  keys(): IterableIterator<string>
  values(): IterableIterator<string>
  [Symbol.iterator](): IterableIterator<string>
  readonly [Symbol.toStringTag]: string
}

type StartViewTransitionOptions = {
  update?: ViewTransitionUpdateCallback
  types?: string[]
}

interface ViewTransitionDocument extends Document {
  startViewTransition(
    callback?: ViewTransitionUpdateCallback,
    options?: StartViewTransitionOptions,
  ): ViewTransition
}
/* eslint-enable no-unused-vars */

import { useAuthStore } from '../stores/auth'
import { useUIStore } from '../stores/ui'

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()
  const uiStore = useUIStore()

  // Start global page loader
  uiStore.startLoading()

  // Ensure auth is initialized explicitly
  if (!authStore.initialized) {
    await authStore.initialize()
  }

  // Ensure we have the latest auth state (optional, but good for refresh)
  if (!authStore.user && authStore.isAuthenticated) {
    // If pinia persistence was used this would be different,
    // but currently we rely on `checkAuth` on app mount usually.
    // For now, assume state is correct.
  }

  // 1. Prevent logged-in users from seeing Login page
  if (to.path === '/login' && authStore.isAuthenticated) {
    next('/admin')
    return
  }

  // 2. Check if route requires admin access
  if (to.path.startsWith('/admin')) {
    if (!authStore.isAuthenticated) {
      // Not logged in -> Redirect to login
      next('/login')
      return
    }
    // We removed the role check redirect to '/' as requested.
    // Authorized users stay here.
  }

  // 3. Allow navigation
  next()
})

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
  const uiStore = useUIStore()

  uiStore.stopLoading()
  submitMetric('page_view', { path: to.fullPath })
})

router.onError(() => {
  const uiStore = useUIStore()
  uiStore.stopLoading()
})

export default router
