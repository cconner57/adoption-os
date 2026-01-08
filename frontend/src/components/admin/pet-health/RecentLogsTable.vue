<script setup lang="ts">
import type { IHealthLogEntry } from '../../../models/common'

defineProps<{
  logs: IHealthLogEntry[]
}>()

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}

function formatWeight(grams?: number | null) {
  if (!grams) return '-'
  const kg = (grams / 1000).toFixed(2)
  const lbs = (grams * 0.00220462).toFixed(2)
  return `${kg}kg (${lbs}lbs)`
}

function getStatusColor(status?: string | null) {
  if (!status) return 'gray'
  if (['normal', 'energetic', 'calm'].includes(status)) return 'green'
  if (['low', 'lethargic'].includes(status)) return 'orange'
  if (['none', 'blood', 'diarrhea'].includes(status)) return 'red'
  return 'blue'
}
</script>

<template>
  <div class="logs-section">
    <h3>Recent Logs</h3>
    <div class="logs-table-wrapper">
      <table class="logs-table">
        <thead>
          <tr>
            <th>Date</th>
            <th>Weight</th>
            <th>Temp</th>
            <th>Eating</th>
            <th>Activity</th>
            <th>Notes</th>
            <th>Recorded By</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="log in logs" :key="log.id">
            <td>{{ formatDate(log.date) }}</td>
            <td>{{ formatWeight(log.weight) }}</td>
            <td>{{ log.temperature || '-' }}°F</td>
            <td>
              <span class="status-text" :class="getStatusColor(log.eating)">
                {{ log.eating || '-' }}
              </span>
            </td>
            <td>
              <span class="status-text" :class="getStatusColor(log.activity)">
                {{ log.activity || '-' }}
              </span>
            </td>
            <td class="notes-cell" :title="log.notes || ''">{{ log.notes || '-' }}</td>
            <td class="text-muted">{{ log.recordedBy }}</td>
          </tr>
          <tr v-if="logs.length === 0">
            <td colspan="7" class="text-center">No records found.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.logs-section {
  h3 {
    font-size: 1.2rem;
    margin-bottom: 16px;
    color: var(--font-color-dark);
  }
}

.logs-table-wrapper {
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden;
  max-height: 500px;
  overflow-y: auto;
}

.logs-table {
  width: 100%;
  border-collapse: collapse;

  th {
    background: #f9fafb;
    padding: 12px 16px;
    text-align: left;
    font-size: 0.85rem;
    font-weight: 600;
    color: var(--font-color-medium);
    position: sticky;
    top: 0;
    z-index: 10;
  }

  td {
    padding: 12px 16px;
    border-top: 1px solid #e5e7eb;
    font-size: 0.9rem;
    color: var(--font-color-dark);
  }

  tr:nth-child(even) {
    background-color: #f9fafb;
  }

  tr:hover {
    background: #f3f4f6;
  }
}

.status-text {
  font-weight: 500;
  text-transform: capitalize;

  &.green {
    color: var(--green);
  }
  &.orange {
    color: var(--yellow);
  }
  &.red {
    color: var(--red);
  }
}

.notes-cell {
  max-width: 200px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.text-muted {
  color: var(--font-color-medium);
}

.text-center {
  text-align: center;
  padding: 16px;
}
</style>
