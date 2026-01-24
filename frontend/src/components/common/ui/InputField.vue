<script setup lang="ts">
import { nextTick,ref, useAttrs, watch } from 'vue'

defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
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

const inputRef = ref<HTMLInputElement | null>(null)

function onInput(e: Event) {
  const target = e.target as HTMLInputElement | null
  let val: string | number | null = target?.value ?? null

  console.log(`DEBUG: InputField[${props.label}] onInput:`, val)

  if (props.type === 'number' && val !== null && val !== '') {
    const num = Number(val)
    if (!isNaN(num)) {
      val = num
    }
  }
  emit('update:modelValue', val)

  nextTick(() => {
    console.log(
      `DEBUG: InputField[${props.label}] nextTick check. Prop: '${props.modelValue}', Input: '${inputRef.value?.value}'`,
    )
    if (inputRef.value && String(props.modelValue ?? '') !== inputRef.value.value) {
      console.warn(`DEBUG: InputField[${props.label}] RESETTING value to prop`)
      inputRef.value.value = String(props.modelValue ?? '')
    }
  })
}

watch(
  () => props.modelValue,
  (newVal) => {
    if (inputRef.value && inputRef.value.value !== String(newVal ?? '')) {
      inputRef.value.value = String(newVal ?? '')
    }
  },
)

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
    <input
      ref="inputRef"
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
</template>

<style scoped>
.control {
  display: flex;
  flex-direction: column;
  gap: 8px;
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
  flex: 1; 
}

</style>
