<script setup lang="ts">
import { computed } from 'vue';

import { icons } from './icons';

const props = withDefaults(
  defineProps<{
    name: string
    size?: string | number
    viewBox?: string
  }>(),
  {
    size: 24,
    viewBox: '0 0 24 24',
  },
)

const path = computed(() => {
  return icons[props.name as keyof typeof icons] || ''
})

const style = computed(() => ({
  width: typeof props.size === 'number' ? `${props.size}px` : props.size,
  height: typeof props.size === 'number' ? `${props.size}px` : props.size,
}))
</script>

<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    :width="size"
    :height="size"
    :viewBox="viewBox"
    :style="style"
    fill="currentColor"
    role="img"
    aria-hidden="true"
  >
    <path :d="path" />
  </svg>
</template>
