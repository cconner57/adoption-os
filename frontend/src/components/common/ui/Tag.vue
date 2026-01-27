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
  background-color: var(--color-neutral-light);
  color: var(--color-neutral-dark); 
  border-color: var(--color-neutral-border);
}

.tag.default.primary {
  background-color: var(--color-primary-weak);
  color: var(--color-primary);
  border-color: var(--color-primary-border);
}

.tag.default.secondary {
  background-color: var(--color-secondary-weak);
  color: var(--color-secondary);
  border-color: var(--color-secondary-border);
}

.tag.outline {
  background-color: var(--text-inverse);
  border-width: 1px;
}

.tag.outline.neutral {
  border-color: var(--color-neutral-border-strong);
  color: var(--color-neutral-strong);
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
  color: var(--color-neutral-text-soft);
}
</style>
