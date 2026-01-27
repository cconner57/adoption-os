import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

import type { IApplicationItem } from '../components/admin/applications/ApplicationCard.vue'
import { API_ENDPOINTS } from '../constants/api'
import { mapApplicationToItem } from '../pages/admin/utils'

export const useApplicationsStore = defineStore('applications', () => {
  const activeTab = ref<'volunteer' | 'surrender' | 'adoption' | 'history'>('adoption')
  const filterStatus = ref<'all' | 'pending' | 'approved' | 'denied'>('all')
  const currentYear = new Date().getFullYear()
  const selectedYear = ref(currentYear)
  const applications = ref<IApplicationItem[]>([])
  const loading = ref(true)
  const lastUpdated = ref(0) // Cache timestamp
  const CACHE_DURATION = 5 * 60 * 1000 // 5 minutes

  // Resend Email State
  const resendingAppId = ref<string | null>(null)
  const resendSuccessAppId = ref<string | null>(null)

  const filteredApplications = computed(() => {
    return applications.value.filter((app) => {
      const typeMatch = activeTab.value === 'history' ? true : app.type === activeTab.value
      const statusMatch = filterStatus.value === 'all' || app.status === filterStatus.value
      return typeMatch && statusMatch
    })
  })

  // Group helpers
  const pendingGroup = computed(() => {
    return filteredApplications.value.filter(app => ['pending', 'denied', 'needs_info'].includes(app.status))
  })

  const approvedGroup = computed(() => {
    return filteredApplications.value.filter(app => app.status === 'approved')
  })

  const deletedGroup = computed(() => {
    return filteredApplications.value.filter(app => app.status === 'autodeleted')
  })

  const fetchApplications = async (autoFocus = false, forceRefresh = false) => {
    // Cache check
    const now = Date.now()
    if (!forceRefresh && applications.value.length > 0 && (now - lastUpdated.value < CACHE_DURATION)) {
      console.log('Using cached applications')
      loading.value = false
      return
    }

    loading.value = true
    try {
      const token = localStorage.getItem('token')
      const yearParam = activeTab.value === 'history' ? selectedYear.value : currentYear

      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      let rawApps: any[] = []

      try {
        const res = await fetch(`${API_ENDPOINTS.APPLICATIONS}?page_size=100&year=${yearParam}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        if (res.ok) {
          const json = await res.json()
          rawApps = json.applications || []
        }
      } catch (e) {
        console.warn('API fetch failed', e)
      }

      // Sort Newest first
      rawApps = rawApps.sort((a, b) => {
        return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      })

      applications.value = rawApps.map(mapApplicationToItem)
      lastUpdated.value = Date.now()

      // Update App Badge
      const pendingCount = applications.value.filter((a) => a.status === 'pending').length
      if ('setAppBadge' in navigator) {
        navigator.setAppBadge(pendingCount).catch((e) => console.error('Error setting badge', e))
      }

      // Auto-switch tabs logic - ONLY on initial load or explicit autoFocus request, not regular refreshes
      if (autoFocus && activeTab.value !== 'history') {
        const adoptionCount = applications.value.filter((a) => a.type === 'adoption').length
        const volunteerCount = applications.value.filter((a) => a.type === 'volunteer').length
        const surrenderCount = applications.value.filter((a) => a.type === 'surrender').length

        if (adoptionCount > 0) {
          activeTab.value = 'adoption'
        } else if (volunteerCount > 0) {
          activeTab.value = 'volunteer'
        } else if (surrenderCount > 0) {
          activeTab.value = 'surrender'
        } else {
          // If empty, default to history last year
          activeTab.value = 'history'
          filterStatus.value = 'all'
          selectedYear.value = currentYear - 1
          // Recursively fetch for history without autoFocus to avoid infinite loop
          fetchApplications(false, true)
          return
        }
      }

      // Auto-switch back to adoption if history empty
      if (autoFocus && activeTab.value === 'history' && applications.value.length === 0) {
        activeTab.value = 'adoption'
        selectedYear.value = currentYear
        fetchApplications(false, true)
        return
      }

    } catch (e) {
      console.error('Failed to process applications', e)
    } finally {
      loading.value = false
    }
  }

  const updateStatus = (app: IApplicationItem, status: IApplicationItem['status']) => {
    const token = localStorage.getItem('token')
    fetch(`${API_ENDPOINTS.APPLICATIONS}/${app.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ status }),
    }).then((res) => {
      if (res.ok) {
        app.status = status
        // Update badge after status change
        const pendingCount = applications.value.filter((a) => a.status === 'pending').length
        if ('setAppBadge' in navigator) {
          navigator.setAppBadge(pendingCount).catch((e) => console.error('Error setting badge', e))
        }
      }
    })
  }

  const viewOriginal = async (appId: string) => {
    try {
      const token = localStorage.getItem('token')
      const res = await fetch(`${API_ENDPOINTS.APPLICATIONS}/${appId}/original`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      if (res.ok) {
        const text = await res.text()
        const blob = new Blob([text], { type: 'text/html' })
        const url = URL.createObjectURL(blob)
        window.open(url, '_blank')
      } else {
        console.error('Failed to fetch original application')
      }
    } catch (e) {
      console.error(e)
    }
  }

  const resendEmail = async (appId: string) => {
    resendingAppId.value = appId
    try {
      const token = localStorage.getItem('token')
      const res = await fetch(`${API_ENDPOINTS.APPLICATIONS}/${appId}/resend-email`, {
        method: 'POST',
        headers: { Authorization: `Bearer ${token}` },
      })
      if (res.ok) {
        resendSuccessAppId.value = appId
        setTimeout(() => {
          resendSuccessAppId.value = null
        }, 2000)
      } else {
        const err = await res.json()
        console.error(`Failed to resend email: ${err.error || 'Unknown error'}`)
      }
    } catch (e) {
      console.error('Failed to resend email due to network error', e)
    } finally {
      resendingAppId.value = null
    }
  }

  return {
    activeTab,
    filterStatus,
    selectedYear,
    currentYear,
    applications,
    loading,
    filteredApplications,
    pendingGroup,
    approvedGroup,
    deletedGroup,
    resendingAppId,
    resendSuccessAppId,
    fetchApplications,
    updateStatus,
    viewOriginal,
    resendEmail,
  }
})
