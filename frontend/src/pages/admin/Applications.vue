<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockApplications, type IApplication } from '../../stores/mockApplications'
import { Capsules, InputSelectGroup } from '../../components/common/ui'
import Button from '../../components/common/ui/Button.vue'

const activeTab = ref<'volunteer' | 'surrender' | 'adoption' | 'history'>('adoption') // Reordered default? User said prioritize Adoption.
const filterStatus = ref<'all' | 'pending' | 'approved' | 'denied' | 'needs_info' | 'autodeleted'>(
  'all',
)
const expandedId = ref<string | null>(null)
const currentYear = new Date().getFullYear()
const selectedYear = ref(currentYear)

const applications = ref<any[]>([])
const loading = ref(true)

const tabs = [
  { id: 'adoption', label: 'Adoption', icon: '🐾' },
  { id: 'volunteer', label: 'Volunteer', icon: '🤝' },
  { id: 'surrender', label: 'Surrender', icon: '🏠' },
  { id: 'history', label: 'Past Years', icon: '🕰️' },
] as const

const API_url = import.meta.env.VITE_API_URL || ''

const filteredApplications = computed(() => {
  return applications.value.filter((app) => {
    // Map backend 'type' to tab id (lowercase is fine, but backend stores 'adoption', 'volunteer', 'surrender')
    // Backend status: 'pending', 'approved', 'denied', 'needs_info'
    const typeMatch = activeTab.value === 'history' ? true : app.type === activeTab.value
    const statusMatch = filterStatus.value === 'all' || app.status === filterStatus.value
    return typeMatch && statusMatch
  })
})

const fetchApplications = async (autoFocus = false) => {
  loading.value = true
  try {
    const token = localStorage.getItem('token')
    // Pass Year Filter
    const yearParam = activeTab.value === 'history' ? selectedYear.value : currentYear

    const res = await fetch(`${API_url}/v1/applications?page_size=100&year=${yearParam}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    if (res.ok) {
      const json = await res.json()
      const rawApps = json.applications || []

      applications.value = rawApps.map((app: any) => {
        const d = app.data || {}
        let name = 'Unknown'
        // ... (normalization) ...
        if (app.type === 'volunteer') {
          name =
            d.firstName && d.lastName
              ? `${d.firstName} ${d.lastName}`
              : d.nameFull || 'Volunteer Applicant'
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
          email: d.email || d.Email || '',
          details: {
            petName:
              d.catPreferenceName ||
              d.animalName ||
              (d.catPreferenceBreed ? `Breed: ${d.catPreferenceBreed}` : null),
            role: d.positionPreferences
              ? Array.isArray(d.positionPreferences)
                ? d.positionPreferences.join(', ')
                : d.positionPreferences
              : null,
            reason: d.interestReason || d.adoptionReason || d.animalWhySurrendered,
          },
          fullApplication: { ...d, createdAt: app.created_at },
        }
      })

      // Auto-Focus Logic (Only on initial load or explicit request)
      if (autoFocus && activeTab.value !== 'history') {
        // Don't redirect if we are specifically looking at history
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
          // Default to Past Years if all current year tabs are empty
          selectTab('history')
          return // selectTab triggers fetch, so return
        }
      }

      // Auto-Focus for History Tab (If switched to history and it's empty?)
      if (autoFocus && activeTab.value === 'history' && applications.value.length === 0) {
        // "If the past years tab has no applications, then lets default to focus on the adoption tab"
        // This prevents being stuck on empty history if auto-switched.
        activeTab.value = 'adoption'
        // Note: We might infinite loop if Adoption is also empty.
        // But our first block handles Adoption=Empty -> ... -> History.
        // Logic:
        // 1. Load Current Year. A=0, V=0, S=0. Switch to History.
        // 2. Fetch History (Previous Year).
        // 3. If History=0. Switch back to Adoption?
        // This creates a cycle if 2025 and 2024 are both empty.
        // Safety: Don't auto-switch back to Adoption if we *just* came from there?
        // Simplification: Just stay on History if it's empty, or default to Adoption if *specifically* requested.
        // The user requirement: "If the past years tab has no applications, then lets default to focus on the adoption tab"
        // Let's implement it but be careful.

        // Actually, if we just switched to history via `selectTab('history')`, we triggered a fetch.
        // If that fetch comes back empty, we go to Adoption.
        // Ideally we stop there.
        // But if Adoption is empty, the first block triggers again?
        // We need to avoid infinite loop.
        // Pass autoFocus=false when falling back to Adoption to stop the loop.
        // We can't pass params to activeTab change.
        // We just set activeTab='adoption' and call fetchApplications(false).

        activeTab.value = 'adoption'
        selectedYear.value = currentYear // Reset year?
        fetchApplications(false)
        return
      }
    }
  } catch (e) {
    console.error('Failed to fetch applications', e)
  } finally {
    loading.value = false
  }
}

import { onMounted } from 'vue'
onMounted(() => {
  fetchApplications(true)
})

const getStatusColor = (status: IApplication['status']) => {
  switch (status) {
    case 'approved':
      return '#d1fae5' // green-100
    case 'denied':
      return '#fee2e2' // red-100
    case 'needs_info':
      return '#ffedd5' // orange-100
    default:
      return '#f3f4f6' // gray-100
  }
}

const getStatusText = (status: IApplication['status']) => {
  switch (status) {
    case 'approved':
      return 'Approved'
    case 'denied':
      return 'Denied'
    case 'needs_info':
      return 'Needs Info'
    default:
      return 'Pending'
  }
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString(undefined, {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

const toggleExpand = (id: string) => {
  if (expandedId.value === id) {
    expandedId.value = null
  } else {
    expandedId.value = id
  }
}

const selectTab = (tabId: typeof activeTab.value) => {
  activeTab.value = tabId
  expandedId.value = null
  filterStatus.value = 'all' // Reset filter when changing tabs

  if (tabId === 'history') {
    selectedYear.value = currentYear - 1
  } else {
    selectedYear.value = currentYear
  }
  fetchApplications() // Re-fetch based on new tab rules (year)
}

// Mock Actions
const updateStatus = (app: IApplication, status: IApplication['status']) => {
  // app.status = status
  // Call API
  const token = localStorage.getItem('token')
  fetch(`${API_url}/v1/applications/${app.id}`, {
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

const getDaysPending = (dateStr: string) => {
  if (!dateStr) return 0
  const created = new Date(dateStr)
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - created.getTime())
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24))
}

const viewOriginal = async (appId: string) => {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_url}/v1/applications/${appId}/original`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    if (res.ok) {
      const text = await res.text()
      const blob = new Blob([text], { type: 'text/html' })
      const url = URL.createObjectURL(blob)
      window.open(url, '_blank')
      // Cleanup after a delay? Browsers handle blob URLs well enough usually if opened.
    } else {
      console.error('Failed to fetch original application')
    }
  } catch (e) {
    console.error(e)
  }
}

const formatKey = (key: string) => {
  // Split camelCase and underscores, capitalize words
  const withSpaces = key.replace(/([a-z])([A-Z])/g, '$1 $2').replace(/_/g, ' ')
  return withSpaces
    .split(' ')
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(' ')
}

const getDisplayFields = (data: any) => {
  if (!data) return []
  const entries = Object.entries(data)

  // Filter
  const filtered = entries.filter(([key, value]) => {
    // Hidden fields
    if (key === 'id') return false
    // Hide parent signature fields if empty
    if ((key === 'parentSignatureData' || key === 'parentSignatureDate') && !value) return false
    // Hide signatures from text list (shown as images)
    if (key === 'signatureData' || key === 'parentSignatureData') return false

    // Legacy generic empty check (keep or relax?)
    // User requested specifically to hide parent fields if not existing.
    // But previously I had `value !== null && value !== ''`.
    // I'll stick to hiding empty values generally to keep it clean, as per previous logic.
    if (value === null || value === '' || value === undefined) return false
    return true
  })

  // Sort
  filtered.sort(([keyA], [keyB]) => {
    const priority = [
      'firstName',
      'lastName',
      'age',
      'birthday',
      'address',
      'city',
      'zip',
      'createdAt',
      'nameFull',
      'signatureDate',
      'allergies',
      'phoneNumber',
      'availability',
      'positionPreferences',
      'interestReason',
      'volunteerExperience',
      'emergencyContactName',
      'emergencyContactPhone',
    ]
    const idxA = priority.indexOf(keyA)
    const idxB = priority.indexOf(keyB)

    if (idxA !== -1 && idxB !== -1) return idxA - idxB
    if (idxA !== -1) return -1
    if (idxB !== -1) return 1

    return 0 // Keep original order for rest
  })

  return filtered.map(([key, value]) => {
    let label = formatKey(key)
    let displayValue: any = value

    // Rename
    if (key === 'createdAt') {
      label = 'Submitted At'
    }

    // Format Dates
    const dateKeys = ['birthday', 'signatureDate', 'parentSignatureDate']
    if (dateKeys.includes(key) && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else {
        displayValue = formatDate(value)
      }
    }

    if (key === 'createdAt' && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else {
        const date = new Date(value)
        // "Jan 17, 2026 at 10:24 AM"
        displayValue = `${formatDate(value)} at ${date.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit' })}`
      }
    }

    // Format Booleans
    if (value === true || value === 'true') displayValue = 'Yes'
    if (value === false || value === 'false') displayValue = 'No'

    return {
      key,
      label,
      value: displayValue,
    }
  })
}

const getSignatureSrc = (data: string) => {
  if (!data) return ''
  // If it's the test data string, return a mock signature (SVG squiggle)
  if (data === 'base64data') {
    return 'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iNTAiPjxwYXRoIGQ9Ik0xMCwyNSBRNDAsNSA3MCwyNSBUMTMwLDI1IiBzdHJva2U9ImJsYWNrIiBzdHJva2Utd2lkdGg9IjIiIGZpbGw9Im5vbmUiLz48L3N2Zz4='
  }

  if (data.length < 50) return '' // Ignore other short dummy text
  if (data.startsWith('data:image')) return data
  return `data:image/png;base64,${data}`
}
</script>

<template>
  <div class="applications-page">
    <div class="page-header">
      <h1>Applications</h1>
      <div class="header-actions">
        <InputSelectGroup
          :modelValue="filterStatus"
          @update:modelValue="(val) => (filterStatus = val as any)"
          :options="[
            { label: 'View All', value: 'all' },
            { label: 'Pending', value: 'pending' },
            { label: 'Approved', value: 'approved' },
            { label: 'Denied', value: 'denied' },
            { label: 'Needs Info', value: 'needs_info' },
            { label: 'Auto-Deleted', value: 'autodeleted' },
          ]"
        />
        <!-- Year Selector for History Tab -->
        <div v-if="activeTab === 'history'" class="ml-4">
          <select v-model="selectedYear" @change="fetchApplications" class="p-2 border rounded-md">
            <option v-for="y in 5" :key="y" :value="currentYear - y">{{ currentYear - y }}</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Tabs -->
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

    <!-- List View -->
    <div class="applications-list">
      <div v-if="filteredApplications.length === 0" class="empty-state">
        <p>No applications found.</p>
      </div>

      <div
        v-for="app in filteredApplications"
        :key="app.id"
        class="app-card"
        :class="{
          unread: app.status === 'pending',
          expanded: expandedId === app.id,
        }"
        @click="toggleExpand(app.id)"
      >
        <!-- Card Header / Summary -->
        <div class="card-summary">
          <div class="app-main">
            <div class="app-header">
              <h3>{{ app.applicantName }}</h3>
              <span class="date">{{ formatDate(app.date) }}</span>
            </div>
            <p class="email">{{ app.email }}</p>

            <div class="details-preview" v-if="!expandedId || expandedId !== app.id">
              <span v-if="app.details.petName"
                ><strong>Pet:</strong> {{ app.details.petName }}</span
              >
              <span v-if="app.details.role"><strong>Role:</strong> {{ app.details.role }}</span>
              <span v-if="app.details.reason"
                ><strong>Reason:</strong> {{ app.details.reason }}</span
              >
            </div>
          </div>

          <div class="app-status">
            <Capsules
              :label="getStatusText(app.status)"
              :color="getStatusColor(app.status)"
              size="sm"
            />
            <div class="expand-icon" :class="{ rotated: expandedId === app.id }">▼</div>
          </div>
        </div>

        <!-- Expanded Content -->
        <div v-if="expandedId === app.id" class="expanded-content" @click.stop>
          <hr class="divider" />

          <div class="action-bar">
            <Button
              title="Approve"
              size="small"
              variant="primary"
              theme="primary"
              :disabled="app.status === 'approved'"
              :onClick="() => updateStatus(app, 'approved')"
            />
            <Button
              title="Request Info"
              size="small"
              variant="primary"
              theme="warning"
              :disabled="app.status === 'needs_info'"
              :onClick="() => updateStatus(app, 'needs_info')"
            />
            <Button
              title="Deny"
              size="small"
              variant="secondary"
              theme="danger"
              :disabled="app.status === 'denied'"
              :onClick="() => updateStatus(app, 'denied')"
            />
            <Button
              title="View Original Application"
              size="small"
              variant="tertiary"
              :onClick="() => viewOriginal(app.id)"
            />
          </div>

          <!-- Auto-Delete Warning -->
          <div
            v-if="app.status === 'pending' && getDaysPending(app.date) > 5"
            class="mt-4 p-3 bg-red-50 border border-red-200 rounded-md text-sm text-red-700 flex items-center gap-2"
            style="margin-bottom: 20px"
          >
            ⚠️ <strong>Warning:</strong> This application will be automatically denied and deleted
            in {{ 7 - getDaysPending(app.date) }} day(s) due to the 7-day retention policy.
          </div>

          <div class="full-application" v-if="app.fullApplication">
            <h4>Application Information</h4>
            <div class="qa-grid">
              <div
                v-for="field in getDisplayFields(app.fullApplication)"
                :key="field.key"
                class="qa-item"
              >
                <span class="question">{{ field.label }}</span>
                <span class="answer">
                  {{ Array.isArray(field.value) ? field.value.join(', ') : field.value }}
                </span>
              </div>

              <!-- Signatures Section (Moved inside Grid) -->
              <div v-if="app.fullApplication.signatureData" class="qa-item span-full">
                <span class="question">Signature</span>
                <img
                  v-if="getSignatureSrc(app.fullApplication.signatureData)"
                  :src="getSignatureSrc(app.fullApplication.signatureData)"
                  class="max-h-24 w-auto border rounded bg-white p-2 mt-1"
                />
                <p v-else class="text-sm text-gray-400 italic">Invalid Signature Data</p>
              </div>

              <div v-if="app.fullApplication.parentSignatureData" class="qa-item span-full">
                <span class="question">Parent Signature</span>
                <img
                  v-if="getSignatureSrc(app.fullApplication.parentSignatureData)"
                  :src="getSignatureSrc(app.fullApplication.parentSignatureData)"
                  class="max-h-24 w-auto border rounded bg-white p-2 mt-1"
                />
                <p v-else class="text-sm text-gray-400 italic">Invalid Signature Data</p>
              </div>
            </div>
          </div>
        </div>
      </div>
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
  h1 {
    font-size: 1.8rem;
    h1 {
      font-size: 1.8rem;
      color: var(--text-primary);
      margin: 0;
    }
    margin: 0;
  }
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Tabs */
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

  &:hover {
    color: var(--text-primary);
    background-color: hsl(from var(--color-neutral) h s 95%);
    border-top-left-radius: 8px;
    border-top-right-radius: 8px;
  }

  &.active {
    color: var(--color-secondary);
    border-bottom-color: var(--color-secondary);
    font-weight: 600;
  }
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

/* List */
.applications-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-bottom: 40px;
}

.app-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden; /* For inner expansion */
  transition: all 0.2s;
  cursor: pointer;
  border-left: 4px solid transparent;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    border-color: #d1d5db;
  }

  &.unread {
    border-left-color: var(--color-secondary);
    background: hsl(from var(--color-secondary) h s 98%);
  }

  &.expanded {
    background: white; /* Reset bg on expand */
    border-color: #d1d5db;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
    cursor: default; /* Inner content handles clicks */
  }
}

.card-summary {
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  cursor: pointer; /* Explicit pointer for the summary area */
}

.app-main {
  flex: 1;
}

.app-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 4px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--text-primary);
  }

  .date {
    font-size: 0.85rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.email {
  margin: 0 0 8px 0;
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.details-preview {
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
  display: flex;
  gap: 16px;

  strong {
    color: var(--text-primary);
    font-weight: 600;
  }
}

.app-status {
  display: flex;
  align-items: center;
  gap: 16px;
}

.expand-icon {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  transition: transform 0.2s;
}

.rotated {
  transform: rotate(180deg);
}

/* Expanded Section */
.expanded-content {
  padding: 0 24px 24px 24px;
  animation: slideDown 0.2s ease-out;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.divider {
  border: 0;
  border-top: 1px solid #f3f4f6;
  margin: 0 0 20px 0;
}

.action-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
  align-items: center; /* Ensure alignment */
}

/* Push the last button (View Original) to the right */
.action-bar > :last-child {
  margin-left: auto;
}

.full-application {
  background: #f9fafb;
  border-radius: 8px;
  padding: 24px;
  border: 1px solid #f3f4f6;

  h4 {
    margin: 0 0 16px 0;
    font-size: 1rem;
    color: var(--text-primary);
    font-weight: 600;
  }
}

.qa-grid {
  display: grid;
  display: grid;
  grid-template-columns: repeat(4, 1fr); /* Force 4 columns for explicit row layout */
  gap: 20px;

  @media (max-width: 1200px) {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
}

.qa-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.span-full {
  grid-column: 1 / span 2; /* Start at col 1 (new row), span 2 cols */
  margin-top: 16px;
}

.question {
  font-size: 0.85rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
}

.answer {
  font-size: 0.95rem;
  color: var(--text-primary);
  line-height: 1.4;
}

.empty-state {
  text-align: center;
  padding: 60px;
  color: hsl(from var(--color-neutral) h s 50%);
  background: var(--text-inverse);
  border-radius: 12px;
  border: 1px dashed #e5e7eb;
}
</style>
