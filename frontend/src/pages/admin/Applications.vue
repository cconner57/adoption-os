<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockApplications, type IApplication } from '../../stores/mockApplications'
import { Capsules, InputSelectGroup } from '../../components/common/ui'
import Button from '../../components/common/ui/Button.vue'

const activeTab = ref<'volunteer' | 'surrender' | 'adoption'>('volunteer')
const filterStatus = ref<'all' | 'pending' | 'approved' | 'denied' | 'needs_info'>('all')
const expandedId = ref<string | null>(null)

const tabs = [
  { id: 'volunteer', label: 'Volunteer', icon: '🤝' },
  { id: 'surrender', label: 'Surrender', icon: '🏠' },
  { id: 'adoption', label: 'Adoption', icon: '🐾' },
] as const

const filteredApplications = computed(() => {
  return mockApplications.value.filter((app) => {
    const typeMatch = app.type === activeTab.value
    const statusMatch = filterStatus.value === 'all' || app.status === filterStatus.value
    return typeMatch && statusMatch
  })
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
}

// Mock Actions
const updateStatus = (app: IApplication, status: IApplication['status']) => {
  app.status = status
  // In real app, call API here
}

const formatKey = (key: string) => {
  // Simple check to just return the key, or could capitalize/format
  return key
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
          ]"
        />
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
          {{ mockApplications.filter((a) => a.type === tab.id && a.status === 'pending').length }}
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
              color="green"
              :disabled="app.status === 'approved'"
              :onClick="() => updateStatus(app, 'approved')"
            />
            <Button
              title="Request Info"
              size="small"
              color="orange"
              :disabled="app.status === 'needs_info'"
              :onClick="() => updateStatus(app, 'needs_info')"
            />
            <Button
              title="Deny"
              size="small"
              color="white"
              style="color: #ef4444; border-color: #fee2e2"
              :disabled="app.status === 'denied'"
              :onClick="() => updateStatus(app, 'denied')"
            />
          </div>

          <div class="full-application">
            <h4>Original Application</h4>
            <div class="qa-grid">
              <div v-for="(value, key) in app.fullApplication" :key="key" class="qa-item">
                <span class="question">{{ formatKey(key) }}</span>
                <span class="answer">{{ value }}</span>
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
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.qa-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
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
