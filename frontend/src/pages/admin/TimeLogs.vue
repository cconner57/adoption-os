<script setup lang="ts">
import { ref, computed } from 'vue'
import {
  mockTimeLogs,
  mockIncidents,
  type ITimeLog,
  type IIncident,
} from '../../stores/mockTimeLogs'
import { Capsules, InputSelectGroup, InputField } from '../../components/common/ui'
import Button from '../../components/common/ui/Button.vue'

const activeTab = ref<'logs' | 'incidents'>('logs')
const searchQuery = ref('')

// --- TIME LOGS LOGIC ---
const filterLogStatus = ref<'all' | 'approved' | 'pending'>('all')

const filteredLogs = computed(() => {
  return mockTimeLogs.value
    .filter((log) => {
      // 1. Text Search
      const searchMatch =
        !searchQuery.value ||
        log.volunteerName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        log.taskDescription.toLowerCase().includes(searchQuery.value.toLowerCase())

      // 2. Status Filter
      const statusMatch = filterLogStatus.value === 'all' || log.status === filterLogStatus.value

      return searchMatch && statusMatch
    })
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
})

const getLogStatusColor = (status: ITimeLog['status']) => {
  switch (status) {
    case 'approved':
      return '#d1fae5' // Green
    case 'flagged':
      return '#fee2e2' // Red
    default:
      return '#f3f4f6' // Gray (Pending)
  }
}

// --- INCIDENTS LOGIC ---
const filterSeverity = ref<'all' | 'low' | 'medium' | 'high' | 'critical'>('all')

const filteredIncidents = computed(() => {
  return mockIncidents.value
    .filter((inc) => {
      const searchMatch =
        !searchQuery.value ||
        inc.reportedBy.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        inc.description.toLowerCase().includes(searchQuery.value.toLowerCase())

      const severityMatch = filterSeverity.value === 'all' || inc.severity === filterSeverity.value

      return searchMatch && severityMatch
    })
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
})

const getSeverityColor = (severity: IIncident['severity']) => {
  switch (severity) {
    case 'critical':
      return '#b91c1c' // Deep Red
    case 'high':
      return '#ef4444' // Red
    case 'medium':
      return '#f97316' // Orange
    case 'low':
      return '#eab308' // Yellow
    default:
      return '#9ca3af'
  }
}

const getSeverityBg = (severity: IIncident['severity']) => {
  switch (severity) {
    case 'critical':
      return '#fef2f2'
    case 'high':
      return '#fef2f2'
    case 'medium':
      return '#fff7ed'
    case 'low':
      return '#fefce8'
    default:
      return '#f3f4f6'
  }
}

// --- ACTIONS LOGIC ---
const activeMenuId = ref<string | null>(null)

const toggleMenu = (id: string) => {
  activeMenuId.value = activeMenuId.value === id ? null : id
}

// Close menu when clicking outside (simple implementation)
const closeIds = () => (activeMenuId.value = null)

const handleLogAction = (action: 'approve' | 'flag' | 'edit' | 'delete', log: ITimeLog) => {
  activeMenuId.value = null
  if (action === 'approve') {
    log.status = 'approved'
    alert(`Approved time log for ${log.volunteerName}`)
  } else if (action === 'flag') {
    log.status = 'flagged'
    alert(`Flagged time log for review`)
  } else if (action === 'delete') {
    mockTimeLogs.value = mockTimeLogs.value.filter((l) => l.id !== log.id)
  } else if (action === 'edit') {
    const newDuration = prompt('Update duration (e.g. 4h 30m):', log.duration || '')
    if (newDuration) log.duration = newDuration
  }
}

const handleIncidentAction = (action: 'resolve' | 'archive' | 'report', inc: IIncident) => {
  activeMenuId.value = null
  if (action === 'resolve') {
    inc.status = 'resolved'
  } else if (action === 'report') {
    alert('Generating PDF report for insurance...')
  }
}
</script>

<template>
  <div class="time-logs-page">
    <div class="page-header">
      <h1>Time Logs & Incidents</h1>
      <div class="search-bar">
        <InputField v-model="searchQuery" placeholder="Search volunteers or descriptions..." />
      </div>
    </div>

    <!-- TABS -->
    <div class="tabs">
      <button class="tab-btn" :class="{ active: activeTab === 'logs' }" @click="activeTab = 'logs'">
        ⏱️ Time Sheets
      </button>
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'incidents' }"
        @click="activeTab = 'incidents'"
      >
        ⚠️ Incident Reports
      </button>
    </div>

    <!-- CONTENT: TIME LOGS -->
    <div v-if="activeTab === 'logs'" class="tab-content">
      <div class="filters-row">
        <InputSelectGroup
          :modelValue="filterLogStatus"
          @update:modelValue="(val) => (filterLogStatus = val as any)"
          :options="[
            { label: 'All Logs', value: 'all' },
            { label: 'Pending Review', value: 'pending' },
            { label: 'Approved', value: 'approved' },
          ]"
        />
        <Button title="Export CSV" size="small" color="white" />
      </div>

      <div class="logs-table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>Volunteer</th>
              <th>Date</th>
              <th>Shift Time</th>
              <th>Duration</th>
              <th>Task Description</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="log in filteredLogs" :key="log.id">
              <td class="fw-bold">{{ log.volunteerName }}</td>
              <td>{{ log.date }}</td>
              <td>
                <span class="time-pill">{{ log.timeIn }}</span>
                <span class="arrow">→</span>
                <span v-if="log.timeOut" class="time-pill">{{ log.timeOut }}</span>
                <span v-else class="active-badge">Active</span>
              </td>
              <td>{{ log.duration || '--' }}</td>
              <td class="desc-cell">{{ log.taskDescription }}</td>
              <td>
                <Capsules :label="log.status" :color="getLogStatusColor(log.status)" size="sm" />
              </td>
              <td>
                <div class="action-menu-container">
                  <button class="action-icon" @click.stop="toggleMenu(log.id)">⋮</button>
                  <div v-if="activeMenuId === log.id" class="dropdown-menu">
                    <button @click="handleLogAction('approve', log)">✅ Approve</button>
                    <button @click="handleLogAction('edit', log)">✏️ Edit Time</button>
                    <button @click="handleLogAction('flag', log)">🚩 Flag Issue</button>
                    <hr />
                    <button class="danger" @click="handleLogAction('delete', log)">
                      🗑️ Delete
                    </button>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <div v-if="filteredLogs.length === 0" class="empty-state" @click="closeIds">
          No logs found.
        </div>
      </div>
    </div>

    <!-- CONTENT: INCIDENTS -->
    <div v-if="activeTab === 'incidents'" class="tab-content">
      <div class="filters-row">
        <InputSelectGroup
          :modelValue="filterSeverity"
          @update:modelValue="(val) => (filterSeverity = val as any)"
          :options="[
            { label: 'All Severities', value: 'all' },
            { label: 'Low', value: 'low' },
            { label: 'Medium', value: 'medium' },
            { label: 'High', value: 'high' },
          ]"
        />
      </div>

      <div class="incidents-list">
        <div
          v-for="inc in filteredIncidents"
          :key="inc.id"
          class="incident-card"
          :style="{ borderLeftColor: getSeverityColor(inc.severity) }"
        >
          <div class="inc-header">
            <div class="inc-title">
              <span
                class="severity-dot"
                :style="{ background: getSeverityColor(inc.severity) }"
              ></span>
              <h3>{{ inc.type.toUpperCase() }} Incident</h3>
              <span class="inc-date">{{ inc.date }} at {{ inc.time }}</span>
            </div>
            <Capsules :label="inc.status" color="#f3f4f6" size="sm" />
          </div>

          <div class="inc-body">
            <p>{{ inc.description }}</p>
          </div>

          <div class="inc-footer">
            <span class="reporter"
              >Reported by: <strong>{{ inc.reportedBy }}</strong></span
            >
            <div class="inc-actions">
              <div class="action-menu-container">
                <Button
                  title="Actions ▾"
                  size="small"
                  color="white"
                  :onClick="() => toggleMenu(inc.id)"
                />
                <div v-if="activeMenuId === inc.id" class="dropdown-menu bottom-right">
                  <button @click="handleIncidentAction('resolve', inc)">✅ Mark Resolved</button>
                  <button @click="handleIncidentAction('report', inc)">📄 Download Report</button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="filteredIncidents.length === 0" class="empty-state" @click="closeIds">
          No incidents matching filters.
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.time-logs-page {
  display: flex;
  flex-direction: column;
  height: 100%;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  h1 {
    margin: 0;
    font-size: 1.8rem;
    color: var(--font-color-dark);
  }
}

.search-bar {
  width: 300px;
}

.tabs {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 4px;
  font-size: 1rem;
  color: var(--font-color-medium);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  font-weight: 500;

  &.active {
    color: var(--font-color-dark);
    border-bottom-color: var(--purple);
    font-weight: 700;
  }
}

.tab-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
  flex: 1;
  min-height: 0;
  overflow-y: auto;
}

.filters-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* DATA TABLE STYLES */
.logs-table-container {
  background: white;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.data-table {
  width: 100%;
  border-collapse: collapse;

  th {
    text-align: left;
    padding: 16px;
    background: #f9fafb;
    color: var(--font-color-medium);
    font-weight: 600;
    font-size: 0.85rem;
    text-transform: uppercase;
    border-bottom: 1px solid #e5e7eb;
  }

  td {
    padding: 16px;
    border-bottom: 1px solid #f3f4f6;
    color: var(--font-color-dark);
    font-size: 0.95rem;
  }

  tr:last-child td {
    border-bottom: none;
  }
}

.fw-bold {
  font-weight: 600;
}
.desc-cell {
  max-width: 300px;
}

.time-pill {
  background: #f3f4f6;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 0.85rem;
  font-family: monospace;
}

.arrow {
  margin: 0 4px;
  color: #9ca3af;
}

.active-badge {
  background: #dcfce7;
  color: #166534;
  font-size: 0.75rem;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: 700;
  text-transform: uppercase;
  margin-left: 4px;
}

.action-menu-container {
  position: relative;
}

.action-icon {
  background: none;
  border: none;
  font-size: 1.2rem;
  padding: 4px 8px;
  border-radius: 4px;
  cursor: pointer;
  color: var(--font-color-medium);

  &:hover {
    background: #f3f4f6;
    color: var(--font-color-dark);
  }
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  min-width: 160px;
  z-index: 10;
  display: flex;
  flex-direction: column;
  padding: 4px;

  &.bottom-right {
    bottom: 100%;
    top: auto;
    margin-bottom: 8px;
  }

  hr {
    border: none;
    border-top: 1px solid #f3f4f6;
    margin: 4px 0;
  }

  button {
    background: none;
    border: none;
    text-align: left;
    padding: 8px 12px;
    font-size: 0.9rem;
    color: var(--font-color-dark);
    cursor: pointer;
    border-radius: 4px;

    &:hover {
      background: #f9fafb;
    }

    &.danger {
      color: #ef4444;
      &:hover {
        background: #fef2f2;
      }
    }
  }
}

/* INCIDENTS CARD STYLES */
.incidents-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.incident-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-left-width: 4px;
  border-radius: 8px;
  padding: 20px;
}

.inc-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.inc-title {
  display: flex;
  align-items: center;
  gap: 12px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--font-color-dark);
  }
}

.severity-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.inc-date {
  font-size: 0.9rem;
  color: var(--font-color-medium);
}

.inc-body {
  margin-bottom: 16px;
  p {
    margin: 0;
    line-height: 1.5;
    color: var(--font-color-dark);
  }
}

.inc-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid #f3f4f6;
  font-size: 0.9rem;
  color: var(--font-color-medium);

  strong {
    color: var(--font-color-dark);
  }
}

.inc-actions {
  display: flex;
  gap: 8px;
}

.empty-state {
  text-align: center;
  padding: 40px;
  color: var(--font-color-medium);
  font-style: italic;
}
</style>
