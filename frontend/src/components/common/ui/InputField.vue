<script setup lang="ts">
import { useAttrs } from 'vue'

defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
  id?: string
  label: string
  placeholder: string
  type?: string
  modelValue: string | number | null
  fullWidth?: boolean
  required?: boolean
  hasError?: boolean
}>()

// eslint-disable-next-line no-unused-vars
const emit = defineEmits<{
  (e: 'update:modelValue', value: string | number | null): void
  (e: 'blur', event: Event): void
}>()

const attrs = useAttrs()
const inputId = props.id ?? `input-${Math.random().toString(36).slice(2, 9)}`

function onInput(e: Event) {
  const target = e.target as HTMLInputElement | null
  emit('update:modelValue', target?.value ?? null)
}

function onBlur(e: Event) {
  emit('blur', e)
}
</script>

<template>
  <div class="control field" :class="{ 'is-fullwidth': props.fullWidth, 'has-error': props.hasError }">
    <label class="label" :for="inputId">{{ props.label }}</label>
    <div class="field">
      <input
        v-bind="attrs"
        :id="inputId"
        :class="{ 'is-empty': !props.modelValue }"
        :placeholder="props.placeholder"
        :type="props.type"
        :value="props.modelValue"
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
  border-color: #ef4444; /* Red-500 */
  outline: 2px solid #ef4444;
}
/* Focus state for error to keep it red */
.has-error input:focus {
  border-color: #ef4444;
  outline: 2px solid #ef4444;
}

input {
  width: 100%;
  padding: 12px 16px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  font-size: 1rem;
  transition: all 0.2s;
  background-color: #ffffff; /* Explicit white background */

  &::placeholder {
    color: var(--font-color-light);
    opacity: 0.6;
  }
}

.is-empty {
  color: var(--font-color-light) !important;
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
  color: var(--font-color-dark);
  opacity: 0.6;
}

input[type='date']::placeholder,
input[type='time']::placeholder {
  color: var(--font-color-dark);
  opacity: 0.6;
}

.field.is-fullwidth {
  width: 100%;
}
</style>
