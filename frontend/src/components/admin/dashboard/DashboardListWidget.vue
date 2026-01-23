<script setup lang="ts">
import { Capsules } from '../../common/ui'

export interface IListItem {
  id: number | string
  title: string
  subtitle: string
  // Badge/Capsule props
  statusLabel?: string
  statusColor?: string
  statusTextColor?: string
  statusSize?: 'sm' | 'md'
}

defineProps<{
  title: string
  items: IListItem[]
}>()
</script>

<template>
  <div class="widget">
    <div class="widget-header">
      <h3>{{ title }}</h3>
    </div>
    <ul class="list-widget">
      <li v-for="item in items" :key="item.id" class="list-item">
        <div class="item-main">
          <span class="item-title">{{ item.title }}</span>
          <span class="item-subtitle">{{ item.subtitle }}</span>
        </div>
        <!-- Default Capsule display if statusLabel exists -->
        <slot name="action" :item="item">
          <Capsules
            v-if="item.statusLabel"
            :size="item.statusSize || 'sm'"
            :label="item.statusLabel"
            :color="item.statusColor"
            :textColor="item.statusTextColor"
          />
        </slot>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.widget {
  background: var(--text-inverse);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
}

.widget h3 {
  font-size: 1.1rem;
  margin: 0;
  font-weight: 700;
  color: var(--text-primary);
}

.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 12px;
}

.list-widget {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-color);
}

.list-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.item-main {
  display: flex;
  flex-direction: column;
}

.item-title {
  font-weight: 600;
  font-size: 0.95rem;
  color: var(--text-primary);
}

.item-subtitle {
  font-size: 0.85rem;
  color: var(--text-secondary);
}
</style>
