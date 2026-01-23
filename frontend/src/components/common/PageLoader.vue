<script setup lang="ts">
import { storeToRefs } from 'pinia'

import { useUIStore } from '@/stores/ui'

const uiStore = useUIStore()
const { isLoading } = storeToRefs(uiStore)
</script>

<template>
  <div v-if="isLoading" class="page-loader">
    <div class="bar"></div>
  </div>
</template>

<style scoped>
.page-loader {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background: transparent;
  z-index: 9999;
  pointer-events: none;
}

.bar {
  width: 100%;
  height: 100%;
  background-color: var(--color-primary);
  transform-origin: left;
  animation: loading 2s infinite;
}

@keyframes loading {
  0% {
    transform: scaleX(0);
  }
  50% {
    transform: scaleX(0.5);
  }
  100% {
    transform: scaleX(1);
    opacity: 0;
  }
}
</style>
