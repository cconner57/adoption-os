<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    label: string
    placeholder: string
    modelValue: string | null
    hasError?: boolean
    spanFull?: boolean
    spanFull?: boolean
    maxChars?: number
  }>(),
  {
    spanFull: true,
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string | null]
}>()
const inputId = `textarea-${Math.random().toString(36).slice(2, 9)}`
</script>

<template>
  <div class="field" :class="{ 'col-span-2': props.spanFull, 'has-error': props.hasError }">
    <label :for="inputId" class="label">{{ props.label }}</label>
    <div class="control">
      <textarea
        :id="inputId"
        :aria-invalid="props.hasError"
        class="textarea"
        rows="3"
        :placeholder="props.placeholder"
        :value="props.modelValue"
        :maxlength="props.maxChars"
        @input="(e) => emit('update:modelValue', (e.target as HTMLTextAreaElement).value)"
        style="box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); resize: none"
      />
      <div v-if="props.maxChars" class="char-count">
        {{ (props.modelValue || '').length }} / {{ props.maxChars }}
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.control {
  display: flex;
  flex-direction: column;
}

.label {
  color: var(--text-primary);
  font-weight: 600;
  margin-bottom: 0.5em;
  display: block;
}
.col-span-2 {
  grid-column: span 2;
}
.full-width {
  grid-column: 1 / -1;
}
.has-error .textarea {
  border-color: var(--color-danger) !important;
  outline: 2px solid var(--color-danger) !important;
}
.char-count {
  text-align: right;
  font-size: 0.75rem;
  color: var(--text-secondary);
  margin-top: 4px;
}
</style>
