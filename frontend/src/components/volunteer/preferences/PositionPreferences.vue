<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  modelValue: string[]
  hasError?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string[]]
}>()

const selected = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})
</script>

<template>
  <fieldset class="field col-span-2" aria-labelledby="position-preferences-legend">
    <legend id="position-preferences-legend" class="label">Volunteer Position Preferences:</legend>
    <fieldset
      class="chips"
      :class="{ 'has-error': hasError }"
      aria-labelledby="position-preferences-legend"
    >
      <label class="chip"
        ><input type="checkbox" value="Feeding/Cleaning" v-model="selected" />
        <span>Feeding/Cleaning</span></label
      >
      <label class="chip"
        ><input type="checkbox" value="Customer Support" v-model="selected" />
        <span>Customer Support</span></label
      >
      <label class="chip"
        ><input type="checkbox" value="Cat Socializing" v-model="selected" />
        <span>Cat Socializing</span></label
      >
    </fieldset>
  </fieldset>
</template>

<style scoped lang="css">
fieldset.field {
  border: 0;
  padding: 0;
  margin: 0;
  @media (max-width: 440px) {
    .col-span-2 {
      grid-column: span 1;
    }
  }
}

.field {
  @media (max-width: 440px) {
    margin-bottom: 8px;
  }
}
.col-span-2 {
  grid-column: span 2;
  @media (max-width: 440px) {
    grid-column: span 1;
  }
}

.label {
  margin-bottom: 8px;
  @media (max-width: 440px) {
    font-size: 13px;
    line-height: 1.5;
  }
}
.chips {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  border: 1px solid transparent; /* Reserve space for border */
  border-radius: 12px;
  padding: 4px; /* formatting spacing */
  @media (max-width: 440px) {
    margin-top: 8px;
    flex-direction: column;
    text-align: center;
    gap: 8px;
  }
}

.chips.has-error {
  border-color: #ef4444;
  background-color: #fef2f2;
}

.chip {
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border-radius: 999px;
  border: 1px solid #e7ebf0;
  background-color: #fff;
  cursor: pointer;
  user-select: none;
  transition:
    background 0.2s,
    border-color 0.2s,
    box-shadow 0.2s;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

  span {
    font-weight: 600;
    color: var(--text-primary);
    line-height: 1.5;
  }

  &:hover {
    border-color: #d7e2f2;
    background: #f2f7ff;
  }

  @media (max-width: 440px) {
    padding: 6px 10px;
    gap: 6px;
    display: flex;
    justify-content: center;
    span {
      font-size: 13px;
      @media (max-width: 440px) {
        font-size: 13px;
        line-height: 1.5;
      }
    }
  }
}
.chip > input {
  position: absolute;
  opacity: 0;
  width: 1px;
  height: 1px;
  pointer-events: none;
  @media (max-width: 440px) {
    width: 0;
    height: 0;
    padding: 0;
    margin: 0;
  }
}
.chip:has(> input:checked) {
  background: color-mix(in srgb, var(--color-primary) 10%, white);
  border-color: var(--color-primary);
  box-shadow: 0 0 0 1px var(--color-primary) inset;
  color: var(--text-primary);
}
.chip:has(> input:focus-visible) {
  box-shadow: 0 0 0 3px var(--ring);
}

@supports not (selector(:has(*))) {
  .chip > input:checked + span {
    background: #e8f1ff;
    border-radius: 999px;
    padding: 6px 10px;
    margin: -6px -10px;
    box-shadow: 0 0 0 2px #bfd0ff inset;
  }
  .chip > input:focus-visible + span {
    box-shadow: 0 0 0 3px var(--ring);
  }
}
</style>
