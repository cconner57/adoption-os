<script setup lang="ts">
import { ref, useAttrs, watch } from 'vue'

defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
  id?: string
  label?: string
  placeholder?: string
  modelValue: string | null
  fullWidth?: boolean
  required?: boolean
  hasError?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | null]
  blur: [event: Event]
}>()

const attrs = useAttrs()
const inputId = props.id ?? `date-input-${Math.random().toString(36).slice(2, 9)}`

const inputRef = ref<HTMLInputElement | null>(null)

function onInput(e: Event) {
  const target = e.target as HTMLInputElement | null
  const val: string | null = target?.value ?? null
  emit('update:modelValue', val)
}

function onBlur(e: Event) {
  emit('blur', e)
}

watch(
  () => props.modelValue,
  (newVal) => {
    if (inputRef.value && inputRef.value.value !== String(newVal ?? '')) {
      inputRef.value.value = String(newVal ?? '')
    }
  },
)
</script>

<template>
  <div
    class="control field"
    :class="{ 'is-fullwidth': props.fullWidth, 'has-error': props.hasError }"
  >
    <label v-if="props.label" class="label" :for="inputId">{{ props.label }}</label>
    <div class="input-wrapper">
        <input
        ref="inputRef"
        v-bind="attrs"
        :id="inputId"
        type="date"
        :class="{ 'is-empty': !props.modelValue }"
        :value="props.modelValue"
        :aria-invalid="props.hasError"
        @input="onInput"
        @blur="onBlur"
        :required="props.required"
        />
        <!-- Custom icon or styling hooks could go here -->
    </div>
  </div>
</template>

<style scoped>
.control {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.input-wrapper {
    position: relative;
    width: 100%;
}

input {
  width: 100%;
  padding: 12px 16px;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  font-size: 1rem;
  transition: all 0.2s;
  background-color: #fff;
  caret-color: var(--text-primary);
  box-shadow: 0 2px 4px rgb(0 0 0 / 10%);
  color: var(--text-primary);

  /* Ensure consistent height with text inputs */
  min-height: 48px;
  font-family: inherit;
  appearance: none;
  appearance: none;
}

input:focus {
    outline: none;
    border-color: var(--color-primary);
    box-shadow: 0 0 0 3px oklch(from var(--color-primary) l c h / 0.20);
}

/* Placeholder styling for date inputs is tricky, usually relies on 'is-empty' class or pseudo elements */
input.is-empty {
    color: var(--text-secondary);
}

</style>
