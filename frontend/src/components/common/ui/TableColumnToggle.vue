<script setup lang="ts">
import { ref } from 'vue'

import SettingsButton from './SettingsButton.vue'
import Toggle from './Toggle.vue'

const props = defineProps<{
  modelValue: Record<string, boolean>
  columns: { key: string; label: string }[]
}>()

const emit = defineEmits<{
  'update:modelValue': [val: Record<string, boolean>]
}>()

const isOpen = ref(false)

const updateCol = (key: string, val: boolean) => {
  const newVal = { ...props.modelValue, [key]: val }
  emit('update:modelValue', newVal)
}
</script>

<template>
  <div class="settings-dropdown-wrapper" @click.stop>
    <SettingsButton
      :active="isOpen"
      title="Table Settings"
      @click="isOpen = !isOpen"
    />

    <div v-if="isOpen" class="settings-dropdown">
      <div class="dropdown-header">Visible Columns</div>
      <div
        v-for="col in columns"
        :key="col.key"
        class="dropdown-item"
        @click="updateCol(col.key, !modelValue[col.key])"
      >
        <Toggle
          :model-value="modelValue[col.key]"
          :label="col.label"
          label-position="left"
          full-width
          style="pointer-events: none"
        />
      </div>
    </div>

    <div v-if="isOpen" class="overlay" @click="isOpen = false"></div>
  </div>
</template>

<style scoped>
.settings-dropdown-wrapper {
  position: relative;
  display: inline-block;
}

.overlay {
  position: fixed;
  inset: 0;
  z-index: 40;
}

.settings-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 8px;
  background: #fff;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 10%);
  width: 280px;
  padding: 8px 0;
  z-index: 50;
}


.dropdown-header {
  padding: 8px 16px;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--color-neutral-text-soft);
  text-transform: uppercase;
  border-bottom: 1px solid var(--border-color);
  margin-bottom: 4px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  font-size: 0.95rem;
  color: var(--text-primary);
  cursor: pointer;
  transition: background 0.2s;

  &:hover {
    background: var(--color-neutral-weak);
  }

  input[type='checkbox'] {
    accent-color: var(--color-secondary);
    cursor: pointer;
  }
}
</style>
