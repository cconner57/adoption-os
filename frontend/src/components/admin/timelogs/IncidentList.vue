<script setup lang="ts">
import { ref } from 'vue'

import type { IIncident } from '../../../stores/mockTimeLogs'
import { Capsules, InputSelectGroup } from '../../common/ui'
import Button from '../../common/ui/Button.vue'

defineProps<{
  filteredIncidents: IIncident[]
  filterSeverity: string
}>()

const emit = defineEmits<{
  'update:filterSeverity': [val: string]
  'incidentAction': [action: string, inc: IIncident]
}>()

const activeMenuId = ref<string | null>(null)

const toggleMenu = (id: string) => {
  activeMenuId.value = activeMenuId.value === id ? null : id
}

const closeIds = () => (activeMenuId.value = null)

const getSeverityColor = (severity: IIncident['severity']) => {
  switch (severity) {
    case 'critical':
      return '#b91c1c' 
    case 'high':
      return '#ef4444' 
    case 'medium':
      return '#f97316' 
    case 'low':
      return '#eab308' 
    default:
      return '#9ca3af'
  }
}

const handleAction = (action: string, inc: IIncident) => {
  activeMenuId.value = null
  emit('incidentAction', action, inc)
}
</script>

<template>
  <div class="tab-content" @click="closeIds">
    <div class="filters-row" @click.stop>
      <InputSelectGroup
        :modelValue="filterSeverity"
        @update:modelValue="(val) => $emit('update:filterSeverity', (val as string) || 'all')"
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
            <div class="action-menu-container" @click.stop>
              <Button
                title="Actions ▾"
                size="small"
                color="white"
                :onClick="() => toggleMenu(inc.id)"
              />
              <div v-if="activeMenuId === inc.id" class="dropdown-menu bottom-right">
                <button @click="handleAction('resolve', inc)">✅ Mark Resolved</button>
                <button @click="handleAction('report', inc)">📄 Download Report</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="filteredIncidents.length === 0" class="empty-state">
        No incidents matching filters.
      </div>
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
    color: var(--text-primary);
  }
}

.severity-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.inc-date {
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.inc-body {
  margin-bottom: 16px;
  p {
    margin: 0;
    line-height: 1.5;
    color: var(--text-primary);
  }
}

.inc-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid var(--border-color);
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);

  strong {
    color: var(--text-primary);
  }
}

.inc-actions {
  display: flex;
  gap: 8px;
}

.action-menu-container {
  position: relative;
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
  }
}

.empty-state {
  text-align: center;
  padding: 40px;
  color: hsl(from var(--color-neutral) h s 50%);
  font-style: italic;
}
</style>
