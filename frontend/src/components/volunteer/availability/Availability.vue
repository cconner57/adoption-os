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
  <fieldset aria-labelledby="avail">
    <legend id="avail" class="section-title">
      Preference on time of day (select all that apply):
    </legend>

    <div class="times" :class="{ 'has-error': hasError }">
      <label class="time-card">
        <input type="checkbox" value="10AM – 12PM" v-model="selected" />
        <div class="time-card__content">
          <strong>10AM – 12PM</strong>
          <small>Clean, feed, & socialize with cats</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="12PM – 2PM" v-model="selected" />
        <div class="time-card__content">
          <strong>12PM – 2PM</strong>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="2PM – 4PM" v-model="selected" />
        <div class="time-card__content">
          <strong>2PM – 4PM</strong>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="4PM – 6PM" v-model="selected" />
        <div class="time-card__content">
          <strong>4PM – 6PM</strong>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="6PM – 8PM" v-model="selected" />
        <div class="time-card__content">
          <strong>6PM – 8PM</strong>
          <small>Clean, feed, & socialize with cats</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="Adoption Event: 12pm - 2pm" v-model="selected" />
        <div class="time-card__content">
          <strong>Adoption Event: 12pm - 2pm</strong>
          <small>Focus on cat adoptions & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="Adoption Event: 2pm - 4pm" v-model="selected" />
        <div class="time-card__content">
          <strong>Adoption Event: 2pm - 4pm</strong>
          <small>Focus on cat adoptions & customer support</small>
        </div>
      </label>
    </div>
  </fieldset>
</template>

<style scoped lang="css">
.section-title {
  font-weight: 700;
  font-size: 18px;
  margin: 18px 0 12px;
}
.times {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;
  border: 1px solid transparent;
  padding: 8px;
  border-radius: 16px;
}

.times.has-error {
  border-color: #ef4444;
  background-color: #fef2f2;
}

.time-card {
  position: relative;
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 12px;
  border: 1px solid #e7ebf0;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  cursor: pointer;
  user-select: none;
  transition:
    background 0.2s,
    border-color 0.2s,
    box-shadow 0.2s;
}
.time-card > input {
  position: absolute;
  opacity: 0;
  width: 1px;
  height: 1px;
  pointer-events: none;
}
.time-card:has(> input:checked) {
  background: color-mix(in srgb, var(--color-primary) 10%, white);
  border-color: var(--color-primary);
  box-shadow: 0 0 0 1px var(--color-primary) inset;
}
.time-card:has(> input:focus-visible) {
  box-shadow: 0 0 0 3px rgba(30, 99, 217, 0.45);
}
.time-card__content small {
  color: #6b7b8a;
  display: block;
  margin-top: 4px;
}
@supports not (selector(:has(*))) {
  .time-card > input:checked + .time-card__content {
    background: #e8f1ff;
    border-radius: 10px;
    margin: -12px;
    padding: 12px;
    box-shadow: 0 0 0 2px #bfd0ff inset;
  }
}
</style>
