<script setup lang="ts">
import { useRouter } from 'vue-router'

import { Icon } from '../../common/ui'

const props = defineProps<{
  label: string
  value: string | number
  icon: string
  color: string
  viewBox?: string
  to?: string
}>()

const router = useRouter()

function handleClick() {
  if (props.to) {
    router.push(props.to)
  }
}
</script>

<template>
  <div
    class="stat-card"
    @click="handleClick"
    @keydown.enter="handleClick"
    role="button"
    :tabindex="to ? 0 : -1"
    :class="{ clickable: !!to }"
  >
    <div class="stat-icon-wrapper" :style="{ color: color }">
      <Icon :name="icon" size="40" :viewBox="viewBox" />
    </div>
    <div class="stat-info">
      <span class="stat-value">{{ value }}</span>
      <span class="stat-label">{{ label }}</span>
    </div>
  </div>
</template>

<style scoped>
.stat-card {
  background: var(--text-inverse);
  padding: 24px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  gap: 20px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  transition: transform 0.2s, box-shadow 0.2s;
  height: 100%;
}

.stat-card.clickable {
  cursor: pointer;
}

.stat-card.clickable:hover {
  transform: translateY(-2px);
  border-color: #cbd5e1;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 10%);
}

.stat-card.clickable:focus-visible {
  border-color: var(--color-primary);
  box-shadow: 0 0 0 2px oklch(from var(--color-primary) l c h / 20%);
  outline: none;
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
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

@media (width <= 768px) {
  .stat-card {
    padding: 16px;
    gap: 12px;
  }

  .stat-value {
    font-size: 1.25rem;
  }

  .stat-label {
    font-size: 0.8rem;
  }
}
</style>
