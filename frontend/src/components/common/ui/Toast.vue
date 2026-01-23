<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  show: boolean
  message: string
  type?: 'success' | 'error' | 'info'
  duration?: number
}>()

const emit = defineEmits<{
  close: []
}>()

let timer: number

function startTimer() {
  clearTimeout(timer)
  if (props.duration !== 0) {
    timer = window.setTimeout(() => {
      emit('close')
    }, props.duration || 6000)
  }
}

onMounted(() => {
  if (props.show) {
    startTimer()
  }
})

// Correct way to watch props in script setup if needed, but simplistic approach:
// The parent controls 'show'. When 'show' becomes true, we should start timer.
// But usually, Toast is conditionally rendered v-if="show" from parent.
// If v-if is used, onMounted works.
// If v-show is used, we need to watch `props.show`.
// Let's assume parent uses `v-if` for simplicity or we watch.

import { watch } from 'vue'
watch(
  () => props.show,
  (val) => {
    if (val) startTimer()
  },
)



onUnmounted(() => {
  clearTimeout(timer)
})
</script>

<template>
  <Transition name="toast-fade">
    <div v-if="show" class="toast-container" :class="type || 'success'">
      <div class="icon">
        <span v-if="type === 'error'">✕</span>
        <span v-else>✓</span>
      </div>
      <span class="message">{{ message }}</span>
      <button class="close-btn" @click="emit('close')">✕</button>
    </div>
  </Transition>
</template>

<style scoped>
.toast-container {
  position: fixed;
  bottom: 24px;
  right: 24px;
  padding: 16px 24px;
  border-radius: 12px;
  background: white;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 12px;
  z-index: 2000;
  min-width: 300px;
  border-left: 6px solid;
}

.toast-container.success {
  border-left-color: var(--color-primary);
}

.toast-container.error {
  border-left-color: var(--color-danger);
}

.icon {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 14px;
}

.success .icon {
  background: hsla(from var(--color-primary) h s l / 0.1);
  color: var(--color-primary);
}

.error .icon {
  background: hsla(from var(--color-danger) h s l / 0.1);
  color: var(--color-danger);
}

.message {
  flex: 1;
  font-weight: 500;
  color: var(--text-primary);
  font-size: 0.95rem;
}

.close-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-secondary);
  font-size: 1.1rem;
  padding: 4px;
  transition: color 0.2s;
}

.close-btn:hover {
  color: var(--text-primary);
}

.toast-fade-enter-active,
.toast-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.toast-fade-enter-from,
.toast-fade-leave-to {
  opacity: 0;
  transform: translateY(20px) scale(0.9);
}
</style>
