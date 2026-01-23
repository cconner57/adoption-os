<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'

const router = useRouter()

const exitKiosk = () => {
  if (confirm('Exit Kiosk Mode?')) {
    router.push('/')
  }
}
</script>

<template>
  <div class="kiosk-layout">
    
    <header class="kiosk-header">
      <div class="logo-area">
        <span class="logo-emoji">🐾</span>
        <span class="logo-text">Paws & Claws Kiosk</span>
      </div>

      <div class="clock-display">
        {{ new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}
      </div>

      <button class="exit-btn" @click="exitKiosk">Exit</button>
    </header>

    <main class="kiosk-main">
      <RouterView v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </RouterView>
    </main>
  </div>
</template>

<style scoped>
.kiosk-layout {
  width: 100vw;
  height: 100vh;
  background: hsl(from var(--color-neutral) h s 98%);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: 'Inter', system-ui, sans-serif;
}

.kiosk-header {
  height: 80px;
  background: var(--text-inverse);
  border-bottom: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  z-index: 10;
}

.logo-area {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--text-primary);
}

.clock-display {
  font-size: 1.25rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  font-variant-numeric: tabular-nums;
}

.exit-btn {
  background: hsl(from var(--color-neutral) h s 95%);
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 40%);
  cursor: pointer;
}

.kiosk-main {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
