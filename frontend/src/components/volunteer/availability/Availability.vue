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
  <div class="availability-section">
    <h3 class="section-label">Preference on time of day (select all that apply):</h3>

    <div class="times-grid" :class="{ 'has-error': hasError }">
      <label class="time-card">
        <input type="checkbox" value="10AM – 12PM" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>10AM – 12PM</strong>
            <div class="check-icon">
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
          </div>
          <small>Clean, feed, & socialize with cats</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="12PM – 2PM" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>12PM – 2PM</strong>
            <div class="check-icon">
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
          </div>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="2PM – 4PM" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>2PM – 4PM</strong>
            <div class="check-icon">
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
          </div>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="4PM – 6PM" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>4PM – 6PM</strong>
            <div class="check-icon">
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
          </div>
          <small>Focus on cat socializing & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="6PM – 8PM" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>6PM – 8PM</strong>
            <div class="check-icon">
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
          </div>
          <small>Clean, feed, & socialize with cats</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="Adoption Event: 12pm - 2pm" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>Adoption Event: 12pm - 2pm</strong>
            <div class="check-icon">
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
          </div>
          <small>Focus on cat adoptions & customer support</small>
        </div>
      </label>

      <label class="time-card">
        <input type="checkbox" value="Adoption Event: 2pm - 4pm" v-model="selected" />
        <div class="time-card__content">
          <div class="time-header">
            <strong>Adoption Event: 2pm - 4pm</strong>
            <div class="check-icon">
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
          </div>
          <small>Focus on cat adoptions & customer support</small>
        </div>
      </label>
    </div>
  </div>
</template>

<style scoped lang="css">
.availability-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 100%;
}

.section-label {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.times-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 12px;
}

.times-grid.has-error {
  border: 1px solid var(--color-danger);
  border-radius: 12px;
  padding: 8px;
}

.time-card {
  position: relative;
  display: flex;
  flex-direction: column;
  padding: 16px;
  border: 1px solid var(--border-color);
  border-radius: 12px;
  background: #fff;
  cursor: pointer;
  user-select: none;
  transition: all 0.2s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.time-card:hover {
  background-color: hsl(from var(--color-neutral) h s 98%);
  border-color: hsl(from var(--color-neutral) h s 80%);
}

.time-card > input {
  position: absolute;
  opacity: 0;
  width: 1px;
  height: 1px;
  pointer-events: none;
}

.time-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.time-header strong {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
}

.check-icon {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 2px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  transition: all 0.2s;
  background: white;
}

.check-icon svg {
  opacity: 0;
  transform: scale(0.5);
  transition: all 0.2s;
}

/* Checked State */
.time-card:has(> input:checked) {
  background: color-mix(in srgb, var(--color-primary) 5%, white);
  border-color: var(--color-primary);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.time-card:has(> input:checked) .check-icon {
  background: var(--color-primary);
  border-color: var(--color-primary);
}

.time-card:has(> input:checked) .check-icon svg {
  opacity: 1;
  transform: scale(1);
}

.time-card__content small {
  font-size: 0.85rem;
  color: var(--text-secondary);
  line-height: 1.4;
  display: block;
}
</style>
