<script setup lang="ts">
import { onUnmounted, watch } from 'vue'

import Icon from './Icon.vue'

const props = defineProps<{
  isOpen: boolean
  title: string
}>()

const emit = defineEmits<{
  close: []
}>()

// Prevent body scroll when open
watch(
  () => props.isOpen,
  (val) => {
    if (val) {
      document.body.style.overflow = 'hidden'
    } else {
      document.body.style.overflow = ''
    }
  }
)

// Ensure cleanup on unmount if component is destroyed while open
onUnmounted(() => {
  document.body.style.overflow = ''
})
</script>

<template>
  <Teleport to="body">
    <!-- Backdrop -->
    <Transition name="fade">
      <div v-if="isOpen" class="drawer-backdrop" @click="emit('close')"></div>
    </Transition>

    <!-- Drawer Content -->
    <Transition name="slide">
      <div v-if="isOpen" class="drawer-panel">
        <header class="drawer-header">
          <h3>{{ title }}</h3>
          <button class="close-btn" @click="emit('close')" aria-label="Close drawer">
            <Icon name="x" size="24" />
          </button>
        </header>

        <div class="drawer-body">
          <slot />
        </div>

        <div class="drawer-footer" v-if="$slots.footer">
          <slot name="footer" />
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.drawer-backdrop {
  position: fixed;
  inset: 0;
  background: rgb(0 0 0 / 50%);
  backdrop-filter: blur(4px);
  z-index: 9998;
}

.drawer-panel {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  width: 100%;
  max-width: 400px;
  background: var(--text-inverse);
  box-shadow: -4px 0 24px rgb(0 0 0 / 15%);
  z-index: 9999;
  display: flex;
  flex-direction: column;
  border-left: 1px solid var(--border-color);

  @media (width <= 600px) {
    max-width: 100%;
    border-left: none;
  }
}

.drawer-header {
  padding: 16px 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--border-color);

  h3 {
    margin: 0;
    font-size: 1.25rem;
    font-weight: 700;
    color: var(--text-primary);
  }
}

.close-btn {
  background: none;
  border: none;
  padding: 8px;
  margin: -8px;
  cursor: pointer;
  color: var(--text-secondary);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;

  &:hover {
    background: var(--color-neutral-weak);
    color: var(--text-primary);
  }
}

.drawer-body {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
}

.drawer-footer {
  padding: 16px 24px;
  border-top: 1px solid var(--border-color);
  background: var(--color-neutral-surface);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.slide-enter-from,
.slide-leave-to {
  transform: translateX(100%);
}
</style>
