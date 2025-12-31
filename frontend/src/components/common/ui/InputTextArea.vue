<script setup lang="ts">
const props = withDefaults(defineProps<{
  label: string
  placeholder: string
  modelValue: string
  hasError?: boolean
  spanFull?: boolean
}>(), {
  spanFull: true
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()
</script>

<template>
  <div class="field" :class="{ 'col-span-2': props.spanFull, 'has-error': props.hasError }">
    <label class="label">{{ props.label }}</label>
    <div class="control">
      <textarea
        class="textarea"
        rows="3"
        :placeholder="props.placeholder"
        :value="props.modelValue"
        @input="(e) => emit('update:modelValue', (e.target as HTMLTextAreaElement).value)"
        style="box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); resize: none"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.label {
  color: var(--font-color-dark);
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
  border-color: #ef4444 !important;
  outline: 2px solid #ef4444 !important;
}
</style>
