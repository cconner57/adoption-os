<script setup lang="ts">
import { computed } from 'vue'



const props = defineProps<{
  modelValue?: boolean
  label?: string
  disabled?: boolean
  id?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

const inputId = computed(() => props.id || `checkbox-${Math.random().toString(36).slice(2, 11)}`)

const toggle = () => {
  if (props.disabled) return
  emit('update:modelValue', !props.modelValue)
}
</script>

<template>
  <div
    class="checkbox-wrapper"
    :class="{ disabled: props.disabled, checked: props.modelValue }"
    @click="toggle"
  >
    <div class="checkbox-input">
      <input
        type="checkbox"
        :id="inputId"
        :checked="props.modelValue"
        :disabled="props.disabled"
        readonly
        tabindex="-1"
      />
      <div class="visual-checkbox">
        <div class="inner-circle"></div>
      </div>
    </div>
    <label v-if="label" :for="inputId" class="checkbox-label" @click.stop>
      {{ label }}
    </label>
  </div>
</template>

<style scoped>
.checkbox-wrapper {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  user-select: none;
  position: relative;

  &.disabled {
    cursor: not-allowed;
    opacity: 0.6;
  }
}

.checkbox-input {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
}

/* Hide native input visually but keep for accessibility/form submission if needed,
   though we mainly use div click for custom UI.
   Actually, standard practice for custom checkbox:
   Hidden input, visual sibling.
*/
input[type="checkbox"] {
  position: absolute;
  opacity: 0;
  width: 100%;
  height: 100%;
  margin: 0;
  cursor: inherit;
  z-index: 1;
}

.visual-checkbox {
  width: 20px;
  height: 20px;
  border-radius: 50%; /* Circle container */
  border: 2px solid var(--text-secondary); /* Default border (greyish) */
  background: var(--text-inverse);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.inner-circle {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--color-primary);
  opacity: 0;
  transform: scale(0.5);
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Checked State */
.checkbox-wrapper.checked .visual-checkbox {
  border-color: var(--color-primary);

  /* Optional: Background fill? User said "when checked, we can show a smaller colored circle".
     Usually this implies transparent bg + colored inner circle, OR colored bg + white check.
     "smaller colored circle" implies the inner dot style (Radio button style, but for checkbox).
     Let's stick to the Radio-like appearance as requested.
  */
}

.checkbox-wrapper.checked .inner-circle {
  opacity: 1;
  transform: scale(1);
}

/* Hover State */
.checkbox-wrapper:not(.disabled):hover .visual-checkbox {
  border-color: var(--color-primary);
  background: var(--color-neutral-weak);
}

/* Focus State (on wrapper or input) */
.checkbox-input:focus-within .visual-checkbox {
  box-shadow: 0 0 0 3px oklch(from var(--color-primary) l c h / 20%);
  border-color: var(--color-primary);
}

.checkbox-label {
  font-size: 1rem; /* Increased from 0.95rem to match inputs */
  color: var(--text-primary);
  cursor: inherit;
}
</style>
