<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  label?: string
  modelValue: string | string[]
  options: string[] | { label: string; value: string }[]
  multiple?: boolean
  hasError?: boolean
}>(), {
  multiple: false,
  hasError: false
})

const emit = defineEmits<{
  'update:modelValue': [value: string | string[]]
  'blur': []
}>()

const normalizedOptions = computed(() => {
  return props.options.map(opt => {
    if (typeof opt === 'string') {
      return { label: opt, value: opt }
    }
    return opt
  })
})

const isSelected = (value: string) => {
  if (props.multiple) {
    return Array.isArray(props.modelValue) ? props.modelValue.includes(value) : false
  }
  return props.modelValue === value
}

const toggleOption = (value: string) => {
  if (props.multiple) {
    const current = Array.isArray(props.modelValue) ? [...props.modelValue] : []
    const index = current.indexOf(value)

    if (index === -1) {
      current.push(value)
    } else {
      current.splice(index, 1)
    }
    emit('update:modelValue', current)
  } else {
    emit('update:modelValue', value)
  }
  emit('blur')
}
</script>

<template>
  <fieldset class="field" :class="{ 'has-error': hasError }" :aria-labelledby="label ? `legend-${label.replace(/\s+/g, '-')}` : undefined">
    <legend v-if="label" :id="`legend-${label.replace(/\s+/g, '-')}`" class="label">{{ label }}</legend>
    <div class="chips">
      <label
        v-for="option in normalizedOptions"
        :key="option.value"
        class="chip"
        :class="{ 'selected': isSelected(option.value) }"
      >
        <input
          :type="multiple ? 'checkbox' : 'radio'"
          :value="option.value"
          :checked="isSelected(option.value)"
          @change="toggleOption(option.value)"
          @blur="emit('blur')"
        />
        <span>{{ option.label }}</span>
      </label>
    </div>
  </fieldset>
</template>

<style scoped lang="css">
.field {
  border: 0;
  padding: 0;
  margin: 0;
  width: 100%;
}

.label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  font-size: 0.875rem;
  color: var(--font-color-dark);
}

.chips {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  border: none;
  padding: 0;
  margin: 0;

  @media (max-width: 440px) {
    flex-direction: column;
    text-align: center;
    /* Center chips content */
    align-items: stretch;
  }
}

.chip {
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border-radius: 999px;
  border: 1px solid #e7ebf0;
  background-color: #fff;
  cursor: pointer;
  user-select: none;
  font-size: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition:
    background 0.2s,
    border-color 0.2s,
    box-shadow 0.2s;

  @media (max-width: 440px) {
      width: 100%;
      justify-content: center;
  }

  span {
    font-weight: 600;
    color: var(--text-900);
    line-height: 1.5;
  }

  &:hover {
    border-color: #d7e2f2;
    background: #f2f7ff;
  }
}

.chip > input {
  position: absolute;
  opacity: 0;
  width: 1px;
  height: 1px;
  pointer-events: none;
}

/* Selected State */
.chip:has(> input:checked) {
  background: color-mix(in srgb, var(--green) 10%, white);
  border: 1px solid var(--green);
  box-shadow: 0 0 0 1px var(--green) inset;
  color: var(--font-color-dark);
}

.chip:has(> input:focus-visible) {
  box-shadow: 0 0 0 3px var(--ring);
}

/* Fallback for browsers without :has support */
@supports not (selector(:has(*))) {
  .chip > input:checked + span {
    background: #e8f1ff;
    border-radius: 999px;
    padding: 6px 10px;
    margin: -6px -10px;
    box-shadow: 0 0 0 2px #bfd0ff inset;
  }
  .chip > input:focus-visible + span {
    box-shadow: 0 0 0 3px var(--ring);
  }
}

/* Error State */
.field.has-error .chips {
  outline: 2px solid #ef4444;
  border-color: #ef4444;
  border-radius: 12px;
  padding: 8px;
}
</style>
