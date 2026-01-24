<script setup lang="ts">
import { computed } from 'vue'

import { Icon } from '../../common/ui'

const props = defineProps<{
  pendingCount: number
  adoptablePetsCount: number
}>()

const stats = computed(() => [
  {
    id: 'pending',
    label: 'Pending Applications',
    value: props.pendingCount,
    icon: 'fileText',
    color: 'var(--color-danger)',
  },
  {
    id: 'pets',
    label: 'Adoptable Pets',
    value: props.adoptablePetsCount,
    icon: 'pawPrint',
    viewBox: '0 0 128 128',
    color: 'var(--color-primary)',
  },
  {
    id: 'vols',
    label: 'Volunteers',
    value: '16',
    icon: 'users',
    color: 'var(--color-secondary)',
  },
  {
    id: 'money',
    label: 'Donations (Month)',
    value: '$3,250',
    icon: 'dollar',
    color: 'var(--color-tertiary)',
  },
])
</script>

<template>
  <div class="stats-grid">
    <div v-for="stat in stats" :key="stat.id" class="stat-card">
      <div class="stat-icon-wrapper" :style="{ color: stat.color }">
        <Icon :name="stat.icon" size="40" :viewBox="stat.viewBox" />
      </div>
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
  gap: 20px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  transition: transform 0.2s;
}

.stat-card:hover {
  transform: translateY(-2px);
  border-color: #cbd5e1;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 10%);
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;

  /* backgrounds removed as requested */
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

/* Unused color classes removed */
</style>
