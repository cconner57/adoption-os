<script setup lang="ts">
defineProps<{
  label?: string
  variant?: 'default' | 'outline' | 'ghost'
  color?: 'neutral' | 'primary' | 'secondary' | 'success' | 'warning' | 'danger'
}>()
</script>

<template>
  <span class="tag" :class="[variant || 'default', color || 'neutral']">
    <slot>{{ label }}</slot>
  </span>
</template>

<style scoped>
.tag {
  display: inline-flex;
  align-items: center;
  padding: 6px 14px;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
  line-height: 1.4;
  white-space: nowrap;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.tag.default.neutral {
  background-color: hsl(from var(--color-neutral) h s 90%);
  color: hsl(from var(--color-neutral) h s 20%); 
  border-color: hsl(from var(--color-neutral) h s 85%);
}

.tag.default.primary {
  background-color: hsl(from var(--color-primary) h s 92%);
  color: var(--color-primary);
  border-color: hsl(from var(--color-primary) h s 85%);
}

.tag.default.secondary {
  background-color: hsl(from var(--color-secondary) h s 92%);
  color: var(--color-secondary);
  border-color: hsl(from var(--color-secondary) h s 85%);
}

.tag.outline {
  background-color: var(--text-inverse);
  border-width: 1px;
}

.tag.outline.neutral {
  border-color: hsl(from var(--color-neutral) h s 80%);
  color: hsl(from var(--color-neutral) h s 30%);
}

.tag.outline.primary {
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.tag.ghost {
  background-color: transparent;
  padding: 4px 10px;
}

.tag.ghost.neutral {
  color: hsl(from var(--color-neutral) h s 50%);
}
</style>
