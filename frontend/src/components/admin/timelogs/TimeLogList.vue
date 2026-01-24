<script setup lang="ts">
import { ref } from 'vue'

import type { ITimeLog } from '../../../stores/mockTimeLogs'
import { Capsules, InputSelectGroup } from '../../common/ui'
import Button from '../../common/ui/Button.vue'

defineProps<{
  filteredLogs: ITimeLog[]
  filterLogStatus: string
}>()

const emit = defineEmits<{
  'update:filterLogStatus': [val: any] // eslint-disable-line @typescript-eslint/no-explicit-any
  logAction: [action: string, log: ITimeLog]
}>()

const activeMenuId = ref<string | null>(null)

const toggleMenu = (id: string) => {
  activeMenuId.value = activeMenuId.value === id ? null : id
}

const closeIds = () => (activeMenuId.value = null)

const getLogStatusColor = (status: ITimeLog['status']) => {
  switch (status) {
    case 'approved':
      return '#d1fae5' 
    case 'flagged':
      return '#fee2e2' 
    default:
      return '#f3f4f6' 
  }
}

const handleAction = (action: string, log: ITimeLog) => {
  activeMenuId.value = null
  emit('logAction', action, log)
}
</script>

<template>
  <div class="tab-content" @click="closeIds">
    <div class="filters-row" @click.stop>
      <InputSelectGroup
        :modelValue="filterLogStatus"
        @update:modelValue="(val) => $emit('update:filterLogStatus', (val as string) || 'all')"
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
              <div class="action-menu-container" @click.stop>
                <button class="action-icon" @click.stop="toggleMenu(log.id)">⋮</button>
                <div v-if="activeMenuId === log.id" class="dropdown-menu">
                  <button @click="handleAction('approve', log)">✅ Approve</button>
                  <button @click="handleAction('edit', log)">✏️ Edit Time</button>
                  <button @click="handleAction('flag', log)">🚩 Flag Issue</button>
                  <hr />
                  <button class="danger" @click="handleAction('delete', log)">🗑️ Delete</button>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="filteredLogs.length === 0" class="empty-state">No logs found.</div>
    </div>
  </div>
</template>

<style scoped>
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

.logs-table-container {
  background: #fff;
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
    background: hsl(from var(--color-neutral) h s 98%);
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 600;
    font-size: 0.85rem;
    text-transform: uppercase;
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: 16px;
    border-bottom: 1px solid var(--border-color);
    color: var(--text-primary);
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
  background: hsl(from var(--color-neutral) h s 98%);
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
  background: hsl(from var(--color-primary) h s 95%);
  color: var(--color-primary);
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
  color: hsl(from var(--color-neutral) h s 50%);

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
    color: var(--text-primary);
  }
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 10%);
  min-width: 160px;
  z-index: 10;
  display: flex;
  flex-direction: column;
  padding: 4px;

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
    color: var(--text-primary);
    cursor: pointer;
    border-radius: 4px;

    &:hover {
      background: hsl(from var(--color-neutral) h s 98%);
    }

    &.danger {
      color: var(--color-danger);

      &:hover {
        background: hsl(from var(--color-danger) h s 98%);
      }
    }
  }
}

.empty-state {
  text-align: center;
  padding: 40px;
  color: hsl(from var(--color-neutral) h s 50%);
  font-style: italic;
}
</style>
