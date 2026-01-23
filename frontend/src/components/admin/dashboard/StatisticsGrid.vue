<script setup lang="ts">
import { computed } from 'vue'

import { useActiveVolunteersStore } from '../../../stores/activeVolunteers'

const props = defineProps<{
  pendingCount: number
  adoptablePetsCount: number
}>()

const volunteerStore = useActiveVolunteersStore()

const stats = computed(() => [
  { label: 'Pending Applications', value: props.pendingCount, color: 'orange', icon: '📝' },
  { label: 'Adoptable Pets', value: props.adoptablePetsCount, color: 'green', icon: '🐾' },
  { label: 'Volunteers', value: volunteerStore.activeCount, color: 'purple', icon: '🤝' },
  { label: 'Donations (Month)', value: '$3,250', color: 'blue', icon: '❤️' },
])
</script>

<template>
  <div class="stats-grid">
    <div
      v-if="volunteerStore.error"
      style="
        grid-column: 1 / -1;
        background: #fee2e2;
        color: #b91c1c;
        padding: 12px;
        border-radius: 8px;
        margin-bottom: 16px;
      "
    >
      Error loading data: {{ volunteerStore.error }}
    </div>
    <div v-for="stat in stats" :key="stat.label" class="stat-card" :class="`color-${stat.color}`">
      <div class="stat-icon">{{ stat.icon }}</div>
      <div class="stat-info">
        <span class="stat-value">{{ stat.value }}</span>
        <span class="stat-label">{{ stat.label }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
}

.stat-card {
  background: var(--text-inverse);
  padding: 24px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s;
}

.stat-card:hover {
  transform: translateY(-2px);
  border-color: #cbd5e1;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.stat-icon {
  font-size: 2rem;
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: hsl(from var(--color-neutral) h s 95%);
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
}

.stat-label {
  color: var(--text-secondary);
  font-size: 0.9rem;
}

.stat-card.color-orange .stat-icon {
  background-color: hsl(from var(--color-warning) h s 88%);
  color: var(--color-warning);
}
.stat-card.color-green .stat-icon {
  background-color: hsl(from var(--color-primary) h s 88%);
  color: var(--color-primary);
}
.stat-card.color-purple .stat-icon {
  background-color: hsl(from var(--color-secondary) h s 88%);
  color: var(--color-secondary);
}
.stat-card.color-blue .stat-icon {
  background-color: hsl(from var(--color-secondary) h s 85%);
  color: var(--color-secondary);
}
</style>
