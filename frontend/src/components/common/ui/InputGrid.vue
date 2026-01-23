<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
  defineProps<{
    label?: string
    modelValue: string
    rows: string[]
    columns: string[]
    exclusiveOptions?: string[]
    singleSelect?: boolean
  }>(),
  {
    exclusiveOptions: () => ['none of these', 'never encounter'],
    singleSelect: false,
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()

const gridState = computed(() => {
  try {
    if (!props.modelValue || !props.modelValue.startsWith('{')) {
      return {}
    }
    return JSON.parse(props.modelValue) as Record<string, string[]>
  } catch {
    return {}
  }
})

const isSelected = (row: string, col: string) => {
  return gridState.value[row]?.includes(col) || false
}

const toggleSelection = (row: string, col: string) => {
  const current = { ...gridState.value }
  const rowSelections = current[row] || []

  if (props.singleSelect) {
    if (rowSelections.includes(col)) {
      
      current[row] = []
    } else {
      current[row] = [col]
    }
  } else {
    
    const isExclusive = props.exclusiveOptions.includes(col)

    if (isExclusive) {
      if (rowSelections.includes(col)) {
        current[row] = []
      } else {
        current[row] = [col]
      }
    } else {
      let newRowSelections = rowSelections.filter((b) => !props.exclusiveOptions.includes(b))

      if (newRowSelections.includes(col)) {
        newRowSelections = newRowSelections.filter((c) => c !== col)
      } else {
        newRowSelections.push(col)
      }
      current[row] = newRowSelections
    }
  }

  if (current[row] && current[row].length === 0) {
    delete current[row]
  }

  emit('update:modelValue', JSON.stringify(current))
}
</script>

<template>
  <div class="field-group span-full">
    <label v-if="label" class="field-label">
      {{ label }}
    </label>

    <div class="grid-container">
      <div
        class="grid-header"
        :style="{ 'grid-template-columns': `150px repeat(${columns.length}, 1fr)` }"
      >
        <div class="row-label-header"></div>
        <div class="col-header" v-for="col in columns" :key="col">{{ col }}</div>
      </div>

      <div
        class="grid-row"
        v-for="row in rows"
        :key="row"
        :style="{ 'grid-template-columns': `150px repeat(${columns.length}, 1fr)` }"
      >
        <div class="row-label">{{ row }}</div>
        <label
          class="grid-cell"
          v-for="col in columns"
          :key="col"
          :class="{ selected: isSelected(row, col) }"
        >
          <input
            :type="singleSelect ? 'radio' : 'checkbox'"
            :name="`${row}-group`"
            :checked="isSelected(row, col)"
            @change="toggleSelection(row, col)"
            class="hidden-input"
            :aria-label="`${row} - ${col}`"
          />
          <div class="check-indicator" v-if="isSelected(row, col)">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="3"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
          </div>
        </label>
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.field-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.field-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.span-full {
  grid-column: 1 / -1;
}

.grid-container {
  border: 1px solid var(--border-color);
  border-radius: 8px;
  overflow: hidden;
  margin-top: 8px;
  overflow-x: auto;
  background-color: var(--text-inverse); 
}

.grid-header {
  display: grid;
  background-color: hsl(from var(--color-neutral) h s 95%);
  border-bottom: 1px solid var(--border-color);
  font-weight: 600;
  color: var(--text-primary);
  font-size: 0.875rem;
}

.grid-row {
  display: grid;
  border-bottom: 1px solid var(--border-color);
  background-color: var(--text-inverse); 

  &:last-child {
    border-bottom: none;
  }
}

.row-label-header {
  padding: 12px;
}

.col-header {
  padding: 12px;
  text-align: center;
  border-left: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1.2;
}

.row-label {
  padding: 12px;
  background-color: hsl(from var(--color-neutral) h s 95%);
  font-weight: 500;
  color: var(--text-primary);
  display: flex;
  align-items: center;
  border-right: 1px solid var(--border-color);
}

.grid-cell {
  padding: 12px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-left: 1px solid var(--border-color);
  cursor: pointer;
  background-color: var(--text-inverse); 

  &:hover {
    background-color: hsl(from var(--color-neutral) h s 98%); 
    box-shadow: 0 0 0 1px var(--border-color) inset;
  }

  &.selected {
    background: color-mix(in srgb, var(--color-primary) 10%, white);
    box-shadow: 0 0 0 1px var(--color-primary) inset;

    .check-indicator {
      color: var(--color-primary);
    }

    &:hover {
      background: color-mix(in srgb, var(--color-primary) 15%, white);
    }
  }
}

.hidden-input {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.check-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-primary);
}

.grid-header .row-label-header {
  border-top-left-radius: 7px;
}

.grid-header .col-header:last-child {
  border-top-right-radius: 7px;
}

.grid-row:last-child .row-label {
  border-bottom-left-radius: 7px;
}

.grid-row:last-child .grid-cell:last-child {
  border-bottom-right-radius: 7px;
}
</style>
