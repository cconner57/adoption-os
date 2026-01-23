<script setup lang="ts">
import { computed, onMounted, onUnmounted,ref, watch } from 'vue'

const props = withDefaults(
  defineProps<{
    label?: string
    modelValue: string | string[] | null
    options: { label: string; value: string }[]
    multiple?: boolean
    placeholder?: string
    hasError?: boolean
    allowCreate?: boolean
    noResultsText?: string
  }>(),
  {
    multiple: false,
    placeholder: 'Select...',
    options: () => [],
    noResultsText: 'No options found',
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string | string[] | null]
}>()

const isOpen = ref(false)
const searchQuery = ref('')
const containerRef = ref<HTMLElement | null>(null)
const inputRef = ref<HTMLInputElement | null>(null)



const filteredOptions = computed(() => {
  if (!searchQuery.value) return props.options
  const q = searchQuery.value.toLowerCase().trim()

  // Check for exact match ignoring case and extra whitespace
  const exactMatch = props.options.some((o) => o.label.toLowerCase().trim() === q)

  const filtered = props.options.filter((o) => o.label.toLowerCase().includes(q))

  if (props.allowCreate && !exactMatch && q.length > 0) {
    return [
      ...filtered,
      { label: `Create "${searchQuery.value}"`, value: searchQuery.value, isCustom: true },
    ]
  }

  return filtered
})

watch(
  () => props.modelValue,
  (newVal) => {
    if (!props.multiple) {
      if (newVal) {
        const option = props.options.find((o) => o.value === newVal)
        searchQuery.value = option ? option.label : (newVal as string)
      } else {
        searchQuery.value = ''
      }
    }
  },
  { immediate: true },
)

function isSelected(value: string) {
  if (props.multiple) {
    return ((props.modelValue as string[]) || []).includes(value)
  }
  return props.modelValue === value
}

function selectOption(option: { label: string; value: string; isCustom?: boolean }) {
  const valToEmit = option.isCustom ? option.value : option.value

  if (props.multiple) {
    const current = (props.modelValue as string[]) || []
    const newVal = current.includes(valToEmit)
      ? current.filter((v) => v !== valToEmit)
      : [...current, valToEmit]
    emit('update:modelValue', newVal)
    searchQuery.value = ''
    inputRef.value?.focus()
  } else {
    emit('update:modelValue', valToEmit)
    // Keep the selected label/value in the input for single select
    searchQuery.value = option.isCustom ? valToEmit : option.label
    isOpen.value = false
  }
}

function removeTag(value: string) {
  if (props.multiple) {
    const current = (props.modelValue as string[]) || []
    emit(
      'update:modelValue',
      current.filter((v) => v !== value),
    )
  }
}

function handleInput() {
  isOpen.value = true
}

function handleFocus() {
  isOpen.value = true
}

// Close on click outside
function handleClickOutside(e: MouseEvent) {
  if (containerRef.value && !containerRef.value.contains(e.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <div class="combobox-field" :class="{ 'has-error': hasError }" ref="containerRef">
    <label v-if="label" class="label">{{ label }}</label>

    <div class="combobox-wrapper">
      <!-- Selected Tags (Multiple Mode) -->
      <div v-if="multiple && (modelValue as string[])?.length" class="selected-tags">
        <span v-for="val in modelValue as string[]" :key="val" class="tag">
          {{ options.find((o) => o.value === val)?.label || val }}
          <button type="button" class="remove-tag" @click="removeTag(val)">×</button>
        </span>
      </div>

      <!-- Input Area -->
      <div class="input-container">
        <input
          ref="inputRef"
          type="text"
          class="combobox-input"
          :placeholder="multiple && (modelValue as string[])?.length ? '' : placeholder"
          v-model="searchQuery"
          @focus="handleFocus"
          @input="handleInput"
          @keydown.down.prevent
          @keydown.enter.prevent
        />
        <span class="chevron">▼</span>
      </div>

      <!-- Dropdown -->
      <div v-show="isOpen" class="dropdown-menu">
        <template v-if="filteredOptions.length > 0">
          <button
            v-for="option in filteredOptions"
            :key="option.value"
            type="button"
            class="dropdown-item"
            :class="{
              selected: isSelected(option.value),
              'create-option': (option as any).isCustom,
            }"
            @mousedown.prevent="selectOption(option)"
          >
            {{ option.label }}
            <span v-if="isSelected(option.value)" class="check">✓</span>
          </button>
        </template>
        <div v-else class="no-results">
          <span v-if="allowCreate && !searchQuery">Type to create...</span>
          <span v-else>{{ noResultsText }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.combobox-field {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.label {
  font-weight: 600;
  font-size: 0.9rem;
  color: var(--text-primary);
}

.combobox-wrapper {
  position: relative;
  border: 1px solid var(--border-color);
  background: white;
  border-radius: 8px;
  min-height: 48px;
  padding: 4px;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.2s;
}

.combobox-wrapper:focus-within {
  border-color: var(--color-secondary);
  box-shadow: 0 0 0 2px hsla(from var(--color-secondary) h s l / 0.1);
}

.selected-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  padding-left: 6px;
}

.tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background: hsl(from var(--color-primary) h s 95%);
  color: var(--text-primary);
  border: 1px solid hsl(from var(--color-primary) h s 80%);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.85rem;
  font-weight: 500;
}

.remove-tag {
  background: none;
  border: none;
  cursor: pointer;
  color: hsl(from var(--color-danger) h s 60%);
  font-size: 1.1rem;
  line-height: 0.8;
  padding: 0 2px;
  margin-right: -2px;

  &:hover {
    color: var(--color-danger);
  }
}

.input-container {
  flex: 1;
  min-width: 120px;
  display: flex;
  align-items: center;
  position: relative;
}

.combobox-input {
  width: 100%;
  border: none;
  outline: none;
  padding: 8px 12px;
  font-size: 1rem;
  color: var(--text-primary);
  background: transparent;

  &::placeholder {
    opacity: 0.5;
  }
}

.chevron {
  pointer-events: none;
  font-size: 0.7rem;
  color: var(--text-secondary);
  margin-right: 12px;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: white;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  margin-top: 4px;
  max-height: 200px;
  overflow-y: auto;
  z-index: 1000;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
}

.dropdown-item {
  width: 100%;
  text-align: left;
  padding: 10px 16px;
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: var(--text-primary);
  font-size: 0.95rem;

  &:hover {
    background: hsl(from var(--color-neutral) h s 96%);
  }

  &.selected {
    background: hsl(from var(--color-primary) h s 96%);
    color: var(--color-primary);
    font-weight: 500;
  }

  &.create-option {
    color: var(--color-secondary);
    font-weight: 600;
    border-top: 1px dashed var(--border-color);
  }
}

.check {
  font-weight: bold;
}

.no-results {
  padding: 12px;
  text-align: center;
  color: var(--text-secondary);
  font-size: 0.9rem;
}
</style>
