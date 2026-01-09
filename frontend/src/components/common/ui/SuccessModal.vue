<script setup lang="ts">
import Button from './Button.vue'

defineProps<{
  isVisible: boolean
  title?: string
  message?: string
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()
</script>

<template>
  <div v-if="isVisible" class="modal-overlay" @click.self="emit('close')">
    <div class="modal-card">
      <div class="icon-wrapper">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="48"
          height="48"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          class="success-icon"
        >
          <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
          <polyline points="22 4 12 14.01 9 11.01"></polyline>
        </svg>
      </div>
      <h3 class="modal-title">{{ title || 'Success!' }}</h3>
      <p class="modal-message">{{ message || 'Operation completed successfully.' }}</p>
      <div class="modal-actions">
        <Button title="Close" color="green" @click="emit('close')" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background-color: rgba(
    0,
    0,
    0,
    0.5
  ); /* Neutral overlay can stay or use neutral var with opacity */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
  animation: fadeIn 0.3s ease-out;
}

.modal-card {
  background: var(--text-inverse);
  padding: 40px 32px; /* Slight increase in vertical padding */
  border-radius: 24px;
  width: 90%;
  max-width: 400px;
  display: flex;
  flex-direction: column;
  align-items: center; /* Ensures horizontal centering of all children */
  justify-content: center;
  box-shadow:
    0 20px 25px -5px rgba(0, 0, 0, 0.1),
    0 10px 10px -5px rgba(0, 0, 0, 0.04);
  animation: scaleIn 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.icon-wrapper {
  color: var(--color-primary);
  background-color: color-mix(in srgb, var(--color-primary) 10%, white);
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 24px; /* Reduced auto margin reliance */
  animation: popIn 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) 0.1s both;
}

.modal-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 12px;
  color: var(--text-primary);
  text-align: center;
}

.modal-message {
  color: hsl(from var(--color-neutral) h s 30%); /* Darker gray for better contrast */
  margin-bottom: 32px;
  line-height: 1.5;
  text-align: center;
}

.modal-actions {
  display: flex;
  justify-content: center;
  width: 100%;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes scaleIn {
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

@keyframes popIn {
  from {
    opacity: 0;
    transform: scale(0.5);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
