<script setup lang="ts">
import { computed, onMounted, onUnmounted,ref } from 'vue'

const props = withDefaults(
  defineProps<{
    modelValue: string | number | null | (string | number)[]
    options: (string | { label: string; value: string | number })[]
    placeholder?: string
    label?: string
    hasError?: boolean
    fullWidth?: boolean
    multiple?: boolean
  }>(),
  {
    placeholder: 'Select an option',
    fullWidth: false,
    hasError: false,
    multiple: false,
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string | number | null | (string | number)[]]
}>()

const isOpen = ref(false)
const containerRef = ref<HTMLElement | null>(null)

const normalizedOptions = computed(() => {
  return props.options.map((opt) => {
    if (typeof opt === 'object' && opt !== null && 'label' in opt && 'value' in opt) {
      return opt
    }
    return { label: String(opt), value: opt }
  })
})

const selectedLabel = computed(() => {
  if (props.multiple) {
    if (!Array.isArray(props.modelValue) || props.modelValue.length === 0) {
      return props.placeholder
    }
    const selected = normalizedOptions.value.filter((opt) =>
      Array.isArray(props.modelValue) && props.modelValue.includes(opt.value)
    )
    return selected.map((s) => s.label).join(', ')
  }

  const selected = normalizedOptions.value.find((opt) => opt.value === props.modelValue)
  return selected ? selected.label : props.placeholder
})

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const selectOption = (value: string | number) => {
  if (props.multiple) {
    const current = Array.isArray(props.modelValue) ? [...props.modelValue] : []
    const index = current.indexOf(value)
    if (index > -1) {
      current.splice(index, 1)
    } else {
      current.push(value)
    }
    emit('update:modelValue', current)
    
  } else {
    emit('update:modelValue', value)
    isOpen.value = false
  }
}

const handleClickOutside = (event: MouseEvent) => {
  if (containerRef.value && !containerRef.value.contains(event.target as Node)) {
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
  <div
    class="select-container"
    :class="{ 'is-fullwidth': fullWidth, 'has-error': hasError }"
    ref="containerRef"
  >
    <label v-if="label" class="label">{{ label }}</label>

    <div
      class="select-trigger"
      :class="{ 'is-open': isOpen, 'is-placeholder': !modelValue }"
      @click="toggleDropdown"
      tabindex="0"
      @keydown.space.prevent="toggleDropdown"
      @keydown.enter.prevent="toggleDropdown"
      @keydown.esc="isOpen = false"
    >
      <span class="selected-text">{{ selectedLabel }}</span>
      <span class="chevron">▼</span>
    </div>

    <transition name="fade">
      <div v-if="isOpen" class="options-menu">
        <div
          v-for="option in normalizedOptions"
          :key="option.value"
          class="option-item"
          :class="{
            'is-selected': multiple
              ? Array.isArray(modelValue) && modelValue.includes(option.value)
              : option.value === modelValue,
          }"
          @click="selectOption(option.value)"
        >
          {{ option.label }}
          <span
            v-if="
              multiple
                ? Array.isArray(modelValue) && modelValue.includes(option.value)
                : option.value === modelValue
            "
            class="check"
            >✓</span
          >
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.select-container {
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 200px; 
}

.select-container.is-fullwidth {
  width: 100%;
}

.label {
  font-weight: 600;
  font-size: 0.9rem;
  color: var(--text-primary);
}

.select-trigger {
  width: 100%;
  padding: 12px 16px;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  font-size: 1rem;
  background-color: #ffffff;
  color: var(--text-primary);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: all 0.2s;
  user-select: none;

  &:hover {
    border-color: hsl(from var(--color-primary) h s 80%);
  }

  &.is-open {
    border-color: var(--color-primary);
    box-shadow: 0 0 0 2px hsla(from var(--color-primary) h s l / 0.2);
  }
}

.select-trigger.is-placeholder .selected-text {
  opacity: 0.8;
}

.selected-text {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 8px;
}

.chevron {
  font-size: 0.7rem;
  opacity: 0.5;
  transition: transform 0.2s;
}

.select-trigger.is-open .chevron {
  transform: rotate(180deg);
}

.options-menu {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  margin-top: 4px;
  background: white;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  max-height: 250px;
  overflow-y: auto;
  z-index: 100;
  padding: 4px;
}

.option-item {
  padding: 10px 12px;
  font-size: 0.95rem;
  color: var(--text-primary);
  cursor: pointer;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: background 0.1s;

  &:hover {
    background-color: hsl(from var(--color-neutral) h s 96%);
  }

  &.is-selected {
    background-color: hsl(from var(--color-primary) h s 95%);
    color: var(--color-primary);
    font-weight: 500;
  }
}

.check {
  font-size: 0.8rem;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.1s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
