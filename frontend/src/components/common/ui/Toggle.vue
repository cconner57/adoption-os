<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
  defineProps<{
    modelValue: boolean
    label?: string
    disabled?: boolean
    labelPosition?: 'left' | 'right'
  }>(),
  {
    modelValue: false,
    disabled: false,
    labelPosition: 'right',
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

function toggle() {
  if (!props.disabled) {
    emit('update:modelValue', !props.modelValue)
  }
}
</script>

<template>
  <div
    class="toggle-wrapper"
    :class="{
      disabled: props.disabled,
      'label-left': props.labelPosition === 'left',
    }"
    @click="toggle"
  >
    <div
      class="toggle-switch"
      :class="{ active: props.modelValue }"
      role="switch"
      :aria-checked="props.modelValue"
      :aria-label="props.label"
    >
      <div class="toggle-thumb"></div>
    </div>
    <span v-if="props.label" class="toggle-label">{{ props.label }}</span>
  </div>
</template>

<style scoped>
.toggle-wrapper {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  user-select: none;
}

.toggle-wrapper.label-left {
  flex-direction: row-reverse;
  justify-content: flex-end;
}

.toggle-wrapper.disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.toggle-switch {
  width: 44px; /* standard width */
  height: 24px;
  background-color: hsl(from var(--color-neutral) h s 80%);
  border-radius: 999px;
  padding: 2px;
  transition: background-color 0.2s ease;
  position: relative;
}

.toggle-switch.active {
  background-color: var(--color-primary);
}

.toggle-thumb {
  width: 20px;
  height: 20px;
  background-color: white;
  border-radius: 50%;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
  transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  transform: translateX(0);
}

.toggle-switch.active .toggle-thumb {
  transform: translateX(20px);
}

.toggle-label {
  font-size: 0.95rem;
  font-weight: 500;
  color: var(--text-primary);
}
</style>
