<script setup lang="ts">
import { useAttrs } from 'vue'

defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
  id?: string
  id?: string
  label?: string
  placeholder: string
  type?: string
  modelValue: string | number | null
  fullWidth?: boolean
  required?: boolean
  hasError?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | number | null]
  blur: [event: Event]
}>()

const attrs = useAttrs()
const inputId = props.id ?? `input-${Math.random().toString(36).slice(2, 9)}`

function onInput(e: Event) {
  const target = e.target as HTMLInputElement | null
  let val: string | number | null = target?.value ?? null

  if (props.type === 'number' && val !== null && val !== '') {
    const num = Number(val)
    if (!isNaN(num)) {
      val = num
    }
  }
  emit('update:modelValue', val)
}

function onBlur(e: Event) {
  emit('blur', e)
}
</script>

<template>
  <div
    class="control field"
    :class="{ 'is-fullwidth': props.fullWidth, 'has-error': props.hasError }"
  >
    <label v-if="props.label" class="label" :for="inputId">{{ props.label }}</label>
    <div class="field">
      <input
        v-bind="attrs"
        :id="inputId"
        :class="{ 'is-empty': !props.modelValue }"
        :placeholder="props.placeholder"
        :type="props.type"
        :value="props.modelValue"
        :aria-invalid="props.hasError"
        @input="onInput"
        @blur="onBlur"
        :required="props.required"
      />
    </div>
  </div>
</template>

<style scoped>
.control {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label {
  font-weight: 600;
  font-size: 0.9rem;
}

.has-error input {
  border-color: var(--color-danger); /* Red-500 */
  outline: 2px solid var(--color-danger);
}
/* Focus state for error to keep it red */
.has-error input:focus {
  border-color: var(--color-danger);
  outline: 2px solid var(--color-danger);
}

input {
  width: 100%;
  padding: 12px 16px;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  font-size: 1rem;
  transition: all 0.2s;
  background-color: #ffffff; /* Explicit white background for contrast */
  caret-color: var(--text-primary); /* Ensure caret is always visible */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  color: var(--text-primary);

  &::placeholder {
    color: var(--text-primary); /* Darker gray for better visibility */
    opacity: 0.6;
  }
}

.is-empty {
  color: var(--text-primary);
}

/* Date input specific handling */
input[type='date'] {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  min-height: 2.5rem; /* Ensure height matches text inputs */
}

input[type='date'].is-empty,
input[type='time'].is-empty {
  color: var(--text-primary);
  opacity: 0.6;
}

input[type='date']::placeholder,
input[type='time']::placeholder {
  color: var(--text-primary);
  opacity: 1;
}

.field.is-fullwidth {
  width: 100%;
}
</style>
