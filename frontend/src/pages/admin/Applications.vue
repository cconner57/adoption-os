<script setup lang="ts">
import { computed,ref } from 'vue'

import ApplicationCard, {
  type IApplicationItem,
} from '../../components/admin/applications/ApplicationCard.vue'
import { InputSelectGroup } from '../../components/common/ui'
import { API_ENDPOINTS } from '../../constants/api'

function formatRole(prefs: unknown): string | null {
  if (!prefs) return null
  return Array.isArray(prefs) ? prefs.join(', ') : String(prefs || '')
}
const activeTab = ref<'volunteer' | 'surrender' | 'adoption' | 'history'>('adoption')
const filterStatus = ref<'all' | 'pending' | 'approved' | 'denied' | 'needs_info' | 'autodeleted'>(
  'all',
)
const expandedId = ref<string | null>(null)
const currentYear = new Date().getFullYear()
const selectedYear = ref(currentYear)

const applications = ref<IApplicationItem[]>([])
const loading = ref(true)

const tabs = [
  { id: 'adoption', label: 'Adoption', icon: '🐾' },
  { id: 'volunteer', label: 'Volunteer', icon: '🤝' },
  { id: 'surrender', label: 'Surrender', icon: '🏠' },
  { id: 'history', label: 'Past Years', icon: '🕰️' },
] as const

// Used for base logic below
// const API_url = API_ENDPOINTS.APPLICATIONS
// Refactoring to use API_ENDPOINTS directly

const filteredApplications = computed(() => {
  return applications.value.filter((app) => {
    const typeMatch = activeTab.value === 'history' ? true : app.type === activeTab.value
    const statusMatch = filterStatus.value === 'all' || app.status === filterStatus.value
    return typeMatch && statusMatch
  })
})

const toggleExpand = (id: string) => {
  if (expandedId.value === id) {
    expandedId.value = null
  } else {
    expandedId.value = id
  }
}

const selectTab = (id: typeof activeTab.value) => {
  activeTab.value = id
  expandedId.value = null
  if (id !== 'history') {

  }
}

import { type IApplication,mockApplications } from '../../stores/mockApplications'

async function fetchApplications(autoFocus = false) {
  loading.value = true
  try {
    const token = localStorage.getItem('token')
    const yearParam = activeTab.value === 'history' ? selectedYear.value : currentYear

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    let rawApps: any[] = []

    // Fetch from API
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
      console.warn('API fetch failed, falling back to mock data', e)
    }

    // Merge with Mock Store (Frontend persistence)
    const mockApps = mockApplications.value.map(app => ({
      id: app.id,
      type: app.type,
      status: app.status,
      created_at: app.date,
      data: { ...app.details, ...app.fullApplication, email: app.email, firstName: app.applicantName.split(' ')[0], lastName: app.applicantName.split(' ')[1] || '' }
    }))

    // Combine (Mock apps first)
    rawApps = [...mockApps, ...rawApps]

    applications.value = rawApps.map((app: { id: string; type: 'adoption' | 'volunteer' | 'surrender'; status: IApplicationItem['status']; created_at: string; data: Record<string, unknown> }) => {
        const d = app.data || {}
        let name = 'Unknown'

        if (app.type === 'volunteer') {
          name =
            d.firstName && d.lastName
              ? `${d.firstName} ${d.lastName}`
              : String(d.nameFull || 'Volunteer Applicant')
        } else if (app.type === 'adoption') {
          name = d.firstName && d.lastName ? `${d.firstName} ${d.lastName}` : 'Adoption Applicant'
        } else if (app.type === 'surrender') {
          name = d.firstName && d.lastName ? `${d.firstName} ${d.lastName}` : 'Surrender Applicant'
        }

        return {
          id: app.id,
          type: app.type,
          status: app.status,
          date: app.created_at,
          applicantName: name,
          email: String(d.email || d.Email || ''),
          details: {
            petName:
              String(d.petName || d.catPreferenceName ||
              d.animalName ||
              (d.catPreferenceBreed ? `Breed: ${d.catPreferenceBreed}` : null) || ''),
            role: formatRole(d.positionPreferences),
            reason: String(d.interestReason || d.adoptionReason || d.animalWhySurrendered || ''),
          },
          fullApplication: { ...d, createdAt: app.created_at },
        }
      })

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

          activeTab.value = 'history'
          expandedId.value = null
          filterStatus.value = 'all'
          selectedYear.value = currentYear - 1
          fetchApplications()
          return
        }
      }

      if (autoFocus && activeTab.value === 'history' && applications.value.length === 0) {
        activeTab.value = 'adoption'
        selectedYear.value = currentYear
        fetchApplications(false)
        return
      }
  } catch (e) {
    console.error('Failed to process applications', e)
  } finally {
    loading.value = false
  }
}

const updateStatus = (app: IApplicationItem, status: IApplicationItem['status']) => {
  // Check if it's a mock application
  if (app.id.startsWith('a-app-') || app.id.startsWith('v-app-') || app.id.startsWith('s-app-')) {
    const mockApp = mockApplications.value.find(a => a.id === app.id)
    if (mockApp) {
      // Cast standard status to mock status if compatible, or ignore incompatibilities for mock
      mockApp.status = status as IApplication['status']
      app.status = status

      // If approved, you would technically save to DB here
      if (status === 'approved') {
        console.log('Application approved! Saving to permanent database (Mock Action)')
        // Example: Update pet status
        // const petId = app.details.petId
        // updatePetStatus(petId, 'Adopted')
      }
    }
    return
  }

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
    }
  })
}

const viewOriginal = async (appId: string) => {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_ENDPOINTS.APPLICATIONS}/${appId}/original`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
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

</script>

<template>
  <div class="applications-page">
    <div class="page-header">
      <h1>Applications</h1>
      <div class="header-actions">
        <InputSelectGroup
          :modelValue="filterStatus"
          @update:modelValue="(val) => (filterStatus = val as typeof filterStatus)"
          :options="[
            { label: 'View All', value: 'all' },
            { label: 'Pending', value: 'pending' },
            { label: 'Approved', value: 'approved' },
            { label: 'Denied', value: 'denied' },
            { label: 'Needs Info', value: 'needs_info' },
            { label: 'Auto-Deleted', value: 'autodeleted' },
          ]"
        />
        <div v-if="activeTab === 'history'" class="ml-4">
          <select v-model="selectedYear" @change="() => fetchApplications()" class="p-2 border rounded-md">
            <option v-for="y in 5" :key="y" :value="currentYear - y">{{ currentYear - y }}</option>
          </select>
        </div>
      </div>
    </div>

    <div class="tabs">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        class="tab-btn"
        :class="{ active: activeTab === tab.id }"
        @click="selectTab(tab.id)"
      >
        <span class="tab-icon">{{ tab.icon }}</span>
        {{ tab.label }}
        <span class="count-badge">
          {{ applications.filter((a) => a.type === tab.id && a.status === 'pending').length }}
        </span>
      </button>
    </div>

    <div class="applications-list">
      <div v-if="loading" class="loading-state">
        <p>Loading applications...</p>
      </div>
      <div v-else-if="filteredApplications.length === 0" class="empty-state">
        <p>No applications found.</p>
      </div>

      <ApplicationCard
        v-else
        v-for="app in filteredApplications"
        :key="app.id"
        :app="app"
        :expanded="expandedId === app.id"
        :isExpandedId="expandedId === app.id"
        @toggle="toggleExpand(app.id)"
        @update-status="updateStatus"
        @view-original="viewOriginal"
      />
    </div>
  </div>
</template>

<style scoped>
.applications-page {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.page-header {
  margin-bottom: 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-header h1 {
  font-size: 1.8rem;
  color: var(--text-primary);
  margin: 0;
}

.tabs {
  display: flex;
  gap: 16px;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 24px;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 16px;
  font-size: 1rem;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  transition: all 0.2s;
}

.tab-btn:hover {
  color: var(--text-primary);
  background-color: hsl(from var(--color-neutral) h s 95%);
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.tab-btn.active {
  color: var(--color-secondary);
  border-bottom-color: var(--color-secondary);
  font-weight: 600;
}

.count-badge {
  background: hsl(from var(--color-neutral) h s 95%);
  color: hsl(from var(--color-neutral) h s 50%);
  font-size: 0.75rem;
  padding: 2px 8px;
  border-radius: 12px;
  font-weight: 600;
}

.tab-btn.active .count-badge {
  background: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
}

.applications-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-bottom: 40px;
}

.empty-state,
.loading-state {
  text-align: center;
  padding: 40px;
  color: var(--text-secondary);
  background: #f9fafb;
  border-radius: 12px;
}
</style>
