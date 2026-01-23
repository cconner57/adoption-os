<script setup lang="ts">
defineProps<{
  items: { id: string; label: string }[]
  modelValue: string
}>()

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div class="tabs">
    <button
      v-for="item in items"
      :key="item.id"
      class="tab-btn"
      :class="{ active: modelValue === item.id }"
      @click="emit('update:modelValue', item.id)"
    >
      {{ item.label }}
    </button>
  </div>
</template>

<style scoped>
.tabs {
  display: flex;
  border-bottom: 1px solid var(--border-color);
  margin-bottom: 24px;
  overflow-x: auto;
  scrollbar-width: none;
}

.tabs::-webkit-scrollbar {
  display: none;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 0;
  margin-right: 24px; 
  font-size: 1rem;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  white-space: nowrap;
  font-weight: 500;
  transition: all 0.2s ease;

  &:last-child {
    margin-right: 0;
  }

  &.active {
    color: var(--color-secondary);
    border-bottom-color: var(--color-secondary);
    font-weight: 600;
  }

  &:hover:not(.active) {
    color: var(--text-primary);
  }
}
</style>
