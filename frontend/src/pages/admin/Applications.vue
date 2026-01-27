<script setup lang="ts">
import { onMounted, ref } from 'vue'

import ApplicationCard from '../../components/admin/applications/ApplicationCard.vue'
import { InputSelectGroup, Select } from '../../components/common/ui'
import { useApplications } from '../../composables/useApplications'

const {
  activeTab,
  filterStatus,
  selectedYear,
  currentYear,
  applications,
  loading,
  filteredApplications,
  pendingGroup,
  approvedGroup,
  resendingAppId,
  resendSuccessAppId,
  fetchApplications,
  updateStatus,
  viewOriginal,
  resendEmail,
} = useApplications()

const tabs = [
  { id: 'adoption', label: 'Adoption', icon: '🐾' },
  { id: 'volunteer', label: 'Volunteer', icon: '🤝' },
  { id: 'surrender', label: 'Surrender', icon: '🏠' },
  { id: 'history', label: 'Past Years', icon: '🕰️' },
] as const

const expandedId = ref<string | null>(null)
const isPendingOpen = ref(true)
const isApprovedOpen = ref(true)
const isFiltersOpen = ref(false)

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
}

const isMounted = ref(false)

onMounted(() => {
  isMounted.value = true
  fetchApplications(true)
})
</script>

<template>
  <div class="applications-page">
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Applications</h1>
    </Teleport>

    <div class="page-header">
      <h1 class="desktop-only">Applications</h1>
      <div class="header-actions">
        <!-- Desktop Filter Group (Always Visible) / Mobile (Collapsible) -->
        <div class="filter-wrapper" :class="{ open: isFiltersOpen }">
          <InputSelectGroup
            :modelValue="filterStatus"
            @update:modelValue="(val) => (filterStatus = val as typeof filterStatus)"
            :options="[
              { label: 'View All', value: 'all' },
              { label: 'Pending', value: 'pending' },
              { label: 'Approved', value: 'approved' },
              { label: 'Denied', value: 'denied' },
            ]"
          />
        </div>

        <div class="mobile-actions">
          <button
            class="filter-toggle-btn"
            :class="{ active: isFiltersOpen }"
            @click="isFiltersOpen = !isFiltersOpen"
          >
            Filters
            <span v-if="filterStatus !== 'all'" class="badge">1</span>
          </button>

          <div v-if="activeTab === 'history'" class="year-select">
            <Select
              v-model="selectedYear"
              :options="Array.from({ length: 5 }, (_, i) => currentYear - i).map(y => ({ label: String(y), value: y }))"
              @update:modelValue="() => fetchApplications()"
              placeholder="Select Year"
            />
          </div>
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

      <template v-else>
        <!-- Pending / Denied / Needs Info Section -->
        <div class="section-container" v-if="pendingGroup.length > 0">
          <div class="section-header" @click="isPendingOpen = !isPendingOpen">
             <div class="section-title-group">
              <h2>Pending & Denied</h2>
              <span class="count-badge">{{ pendingGroup.length }}</span>
             </div>
             <div class="expand-arrow" :class="{ 'rotated': !isPendingOpen }">
               ▼
             </div>
          </div>

          <div v-if="isPendingOpen" class="section-content">
            <ApplicationCard
              v-for="app in pendingGroup"
              :key="app.id"
              :app="app"
              :expanded="expandedId === app.id"
              :isExpandedId="expandedId === app.id"
              :isResending="resendingAppId === app.id"
              :isResendSuccess="resendSuccessAppId === app.id"
              @toggle="toggleExpand(app.id)"
              @update-status="updateStatus"
              @view-original="viewOriginal"
              @resend-email="resendEmail"
            />
          </div>
        </div>

        <!-- Approved Section -->
        <div class="section-container" v-if="approvedGroup.length > 0">
           <div class="section-header" @click="isApprovedOpen = !isApprovedOpen">
             <div class="section-title-group">
              <h2>Approved</h2>
              <span class="count-badge">{{ approvedGroup.length }}</span>
             </div>
             <div class="expand-arrow" :class="{ 'rotated': !isApprovedOpen }">
               ▼
             </div>
          </div>

          <div v-if="isApprovedOpen" class="section-content">
            <ApplicationCard
              v-for="app in approvedGroup"
              :key="app.id"
              :app="app"
              :expanded="expandedId === app.id"
              :isExpandedId="expandedId === app.id"
              :isResending="resendingAppId === app.id"
              :isResendSuccess="resendSuccessAppId === app.id"
              @toggle="toggleExpand(app.id)"
              @update-status="updateStatus"
              @view-original="viewOriginal"
              @resend-email="resendEmail"
            />
          </div>
        </div>
      </template>
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
  flex-shrink: 0;
}

.page-header h1.desktop-only {
  font-size: 1.8rem;
  color: var(--text-primary);
  margin: 0;
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .page-header h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.tabs {
  display: flex;
  gap: 16px;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 24px;
  flex-shrink: 0;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 16px;
  font-size: 1rem;
  color: var(--color-neutral-text-soft);
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
  background-color: var(--color-neutral-weak);
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.tab-btn.active {
  color: var(--color-secondary);
  border-bottom-color: var(--color-secondary);
  font-weight: 600;
}

.count-badge {
  background: var(--color-neutral-weak);
  color: var(--color-neutral-text-soft);
  font-size: 0.75rem;
  padding: 2px 8px;
  border-radius: 12px;
  font-weight: 600;
}

.tab-btn.active .count-badge {
  background: var(--color-secondary-weak);
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

.section-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #f3f4f6;
  border-radius: 8px;
  cursor: pointer;
  user-select: none;
  transition: background-color 0.2s;
}

.section-header:hover {
  background: #e5e7eb;
}

.section-header h2 {
  font-size: 1rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.section-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.section-title-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.expand-arrow {
  color: var(--color-neutral-text-soft);
  font-size: 0.8rem;
  transition: transform 0.2s ease;
}

.expand-arrow.rotated {
  transform: rotate(-90deg);
}

.filter-toggle-btn {
  display: none;
}

.filter-wrapper {
  display: block;
}

@media (width <= 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .header-actions {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .mobile-actions {
    display: flex;
    width: 100%;
    justify-content: space-between;
    align-items: center;
  }

  /* Collapsible Wrapper */
  .filter-wrapper {
    display: none;
    width: 100%;
  }

  /* When open, show filters */
  .filter-wrapper.open {
    display: block;
    animation: fadeIn 0.2s ease-out;
  }

  /* Filter Button Styling (Matches Reference) */
  .filter-toggle-btn {
    all: unset;
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 20px;
    border-radius: 20px;
    background-color: #f3f4f6;
    color: var(--text-secondary);
    font-weight: 600;
    cursor: pointer;
    border: 1px solid transparent;
    transition: all 0.2s;
  }

  .filter-toggle-btn.active {
    background-color: var(--text-primary);
    color: var(--text-inverse);
  }

  .filter-toggle-btn .badge {
    background: var(--color-primary);
    color: #fff;
    border-radius: 50%;
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.75rem;
    font-weight: 700;
  }

  /* Make tabs scrollable on mobile */
  .tabs {
    overflow-x: auto;
    padding-bottom: 4px;
    margin-right: -16px;
    scrollbar-width: none;
  }

  .tabs::-webkit-scrollbar {
    display: none;
  }

  .tab-btn {
    white-space: nowrap;
    flex-shrink: 0;
  }
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-5px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
