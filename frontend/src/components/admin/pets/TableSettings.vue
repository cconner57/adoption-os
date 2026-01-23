<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  modelValue: Record<string, boolean>
}>()

const emit = defineEmits<{
  'update:modelValue': [val: Record<string, boolean>]
}>()

const isOpen = ref(false)

const updateCol = (key: string, val: boolean) => {
  const newVal = { ...props.modelValue, [key]: val }
  emit('update:modelValue', newVal)
}
</script>

<template>
  <div class="settings-dropdown-wrapper" @click.stop>
    <button
      class="icon-btn settings-btn"
      @click="isOpen = !isOpen"
      title="Table Settings"
      :class="{ active: isOpen }"
    >
      ⚙️
    </button>

    <div v-if="isOpen" class="settings-dropdown">
      <div class="dropdown-header">Visible Columns</div>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.age"
          @change="(e) => updateCol('age', (e.target as HTMLInputElement).checked)"
        />
        Age
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.breed"
          @change="(e) => updateCol('breed', (e.target as HTMLInputElement).checked)"
        />
        Species
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.dob"
          @change="(e) => updateCol('dob', (e.target as HTMLInputElement).checked)"
        />
        Date of Birth
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.intake"
          @change="(e) => updateCol('intake', (e.target as HTMLInputElement).checked)"
        />
        Intake Date
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.microchip"
          @change="(e) => updateCol('microchip', (e.target as HTMLInputElement).checked)"
        />
        Microchip
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.name"
          @change="(e) => updateCol('name', (e.target as HTMLInputElement).checked)"
        />
        Name
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.photo"
          @change="(e) => updateCol('photo', (e.target as HTMLInputElement).checked)"
        />
        Photo
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.sex"
          @change="(e) => updateCol('sex', (e.target as HTMLInputElement).checked)"
        />
        Sex
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.sn"
          @change="(e) => updateCol('sn', (e.target as HTMLInputElement).checked)"
        />
        Spayed/Neutered
      </label>
      <label class="dropdown-item">
        <input
          type="checkbox"
          :checked="modelValue.status"
          @change="(e) => updateCol('status', (e.target as HTMLInputElement).checked)"
        />
        Status
      </label>
    </div>

    <!-- Click outside handler could be added but simpler relies on parent click -->
    <div v-if="isOpen" class="overlay" @click="isOpen = false"></div>
  </div>
</template>

<style scoped>
.settings-dropdown-wrapper {
  position: relative;
}

.overlay {
  position: fixed;
  inset: 0;
  z-index: 40;
}

.settings-btn {
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  width: 46px; /* Increased to match input height */
  height: 46px;
  font-size: 1.2rem;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
  }

  &.active {
    background: hsl(from var(--color-neutral) h s 95%);
    border-color: var(--color-secondary);
  }
}

.settings-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 8px;
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 200px;
  padding: 8px 0;
  z-index: 50;
}

.dropdown-header {
  padding: 8px 16px;
  font-size: 0.8rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  text-transform: uppercase;
  border-bottom: 1px solid var(--border-color);
  margin-bottom: 4px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  font-size: 0.95rem;
  color: var(--text-primary);
  cursor: pointer;
  transition: background 0.2s;

  &:hover {
    background: hsl(from var(--color-neutral) h s 95%);
  }

  input[type='checkbox'] {
    accent-color: var(--color-secondary);
  }
}
</style>
