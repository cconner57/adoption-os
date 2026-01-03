<script setup lang="ts">
import { computed } from 'vue'
const props = withDefaults(defineProps<{
  label: string
  modelValue?: string | number | boolean | null
  trueValue?: string | number | boolean | null
  falseValue?: string | number | boolean | null
  trueLabel?: string
  falseLabel?: string
}>(), {
  modelValue: 'No',
  trueValue: 'Yes',
  falseValue: 'No'
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number | boolean | null]
}>()

const displayTrueLabel = computed(() => {
  if (props.trueLabel) return props.trueLabel
  if (typeof props.trueValue === 'string') return props.trueValue
  return 'Yes'
})

const displayFalseLabel = computed(() => {
  if (props.falseLabel) return props.falseLabel
  if (typeof props.falseValue === 'string') return props.falseValue
  return 'No'
})
</script>

<template>
  <div class="field-group">
    <label class="field-label">{{ props.label }}</label>
    <div class="toggle-container">
      <button
        type="button"
        class="toggle-btn"
        :class="{ active: props.modelValue === props.trueValue }"
        @click="emit('update:modelValue', props.trueValue)"
      >
        {{ displayTrueLabel }}
      </button>
      <button
        type="button"
        class="toggle-btn"
        :class="{ active: props.modelValue === props.falseValue }"
        @click="emit('update:modelValue', props.falseValue)"
      >
        {{ displayFalseLabel }}
      </button>
    </div>
  </div>
</template>

<style scoped lang="css">
.field-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
  margin-bottom: 12px;
}

.field-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--font-color-dark);
  margin-bottom: 4px;
}

.toggle-container {
  display: flex;
  background: #fff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 4px;
  height: 48px;
  width: 100%;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.toggle-btn {
  flex: 1;
  border: none;
  background: transparent;
  border-radius: 6px;
  font-weight: 500;
  color: #64748b;
  cursor: pointer;
  padding: 0 4px;
  transition: all 0.2s;
  font-size: 1rem;

  &.active {
    background: color-mix(in srgb, var(--green) 10%, white);
    border: 1px solid var(--green);
    box-shadow: 0 0 0 1px var(--green) inset;
    color: var(--font-color-dark);
    font-weight: 600;
  }

  &:hover:not(.active) {
    background-color: #f1f5f9;
  }
}
</style>
