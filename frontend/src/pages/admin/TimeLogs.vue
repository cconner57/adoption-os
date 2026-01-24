<script setup lang="ts">
import { computed,ref } from 'vue'

import IncidentList from '../../components/admin/timelogs/IncidentList.vue'
import TimeLogList from '../../components/admin/timelogs/TimeLogList.vue'
import { InputField } from '../../components/common/ui'
import {
  type IIncident,
  type ITimeLog,
  mockIncidents,
  mockTimeLogs,
} from '../../stores/mockTimeLogs'

const activeTab = ref<'logs' | 'incidents'>('logs')
const searchQuery = ref('')
const filterLogStatus = ref<'all' | 'approved' | 'pending'>('all')
const filterSeverity = ref<'all' | 'low' | 'medium' | 'high' | 'critical'>('all')

const filteredLogs = computed(() => {
  return mockTimeLogs.value
    .filter((log) => {
      
      const searchMatch =
        !searchQuery.value ||
        log.volunteerName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        log.taskDescription.toLowerCase().includes(searchQuery.value.toLowerCase())

      const statusMatch = filterLogStatus.value === 'all' || log.status === filterLogStatus.value

      return searchMatch && statusMatch
    })
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
})

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

const handleLogAction = (action: string, log: ITimeLog) => {
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

const handleIncidentAction = (action: string, inc: IIncident) => {
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

    <TimeLogList
      v-if="activeTab === 'logs'"
      :filtered-logs="filteredLogs"
      v-model:filterLogStatus="filterLogStatus"
      @logAction="handleLogAction"
    />

    <IncidentList
      v-if="activeTab === 'incidents'"
      :filtered-incidents="filteredIncidents"
      v-model:filterSeverity="filterSeverity"
      @incidentAction="handleIncidentAction"
    />
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
    color: var(--text-primary);
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
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  font-weight: 500;

  &.active {
    color: var(--text-primary);
    border-bottom-color: var(--color-secondary);
    font-weight: 700;
  }
}
</style>
