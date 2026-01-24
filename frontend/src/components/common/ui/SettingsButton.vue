<script setup lang="ts">
import Icon from './Icon.vue'

withDefaults(
  defineProps<{
    active?: boolean
    title?: string
    size?: number
  }>(),
  {
    active: false,
    title: 'Settings',
    size: 20
  }
)

const emit = defineEmits<{
  click: [e: MouseEvent]
}>()
</script>

<template>
  <button
    class="settings-btn"
    :class="{ active }"
    :title="title"
    @click="(e) => emit('click', e)"
  >
    <Icon name="cog" :size="size" />
  </button>
</template>

<style scoped>
.settings-btn {
  background: #fff;
  border: 1px solid var(--border-color);
  width: 48px;
  height: 48px;
  font-size: 1.1rem;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  color: hsl(from var(--color-neutral) h s 50%);

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
    color: var(--text-primary);
    border-color: var(--color-secondary);
  }

  &.active {
    background: hsl(from var(--color-neutral) h s 95%);
    border-color: var(--color-secondary);
    color: var(--text-primary);
  }
}
</style>
