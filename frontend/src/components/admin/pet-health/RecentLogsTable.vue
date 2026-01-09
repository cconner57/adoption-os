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
    color: var(--text-primary);
  }
}

.logs-table-wrapper {
  border: 1px solid var(--border-color);
  border-radius: 12px;
  overflow: hidden;
  max-height: 500px;
  overflow-y: auto;
}

.logs-table {
  width: 100%;
  border-collapse: collapse;

  th {
    background: hsl(from var(--color-neutral) h s 98%);
    padding: 12px 16px;
    text-align: left;
    font-size: 0.85rem;
    font-weight: 600;
    color: hsl(from var(--color-neutral) h s 40%);
    position: sticky;
    top: 0;
    z-index: 10;
  }

  td {
    padding: 12px 16px;
    border-top: 1px solid var(--border-color);
    font-size: 0.9rem;
    color: var(--text-primary);
  }

  tr:nth-child(even) {
    background-color: hsl(from var(--color-neutral) h s 98%);
  }

  tr:hover {
    background: hsl(from var(--color-neutral) h s 95%);
  }
}

.status-text {
  font-weight: 500;
  text-transform: capitalize;

  &.green {
    color: var(--color-primary);
  }
  &.orange {
    color: var(--color-warning);
  }
  &.red {
    color: var(--color-danger);
  }
  &.blue {
    color: var(--color-secondary);
  }
}

.notes-cell {
  max-width: 200px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.text-muted {
  color: hsl(from var(--color-neutral) h s 50%);
}

.text-center {
  text-align: center;
  padding: 16px;
}
</style>
