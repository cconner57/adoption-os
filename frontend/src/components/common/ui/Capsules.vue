<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  label?: string | null
  color?: 'green' | 'blue' | 'red' | 'orange' | 'purple' | 'gray' | 'neutral' | string
  size?: 'sm' | 'md' | 'lg'
}>()

const sizeClass = computed(() => {
  if (props.size === 'sm') return 'capsule--sm'
  if (props.size === 'lg') return 'capsule--lg'
  return 'capsule--md'
})

const colorClass = computed(() => {
  const c = props.color?.toLowerCase()
  if (c === 'green' || c === 'active' || c === 'published' || c === 'sent') return 'capsule--green'
  if (c === 'blue' || c === 'completed' || c === 'adopted') return 'capsule--blue'
  if (c === 'orange' || c === 'scheduled' || c === 'warning') return 'capsule--orange'
  if (c === 'red' || c === 'danger' || c === 'rejected') return 'capsule--red'
  if (c === 'purple' || c === 'foster') return 'capsule--purple'
  if (c === 'gray' || c === 'neutral' || c === 'draft' || c === 'pending' || c === 'intake') return 'capsule--gray'

  // Fallback for custom colors if key not found (though CSS vars expect specific classes)
  return 'capsule--gray'
})
</script>

<template>
  <span
    class="capsule"
    :class="[sizeClass, colorClass]"
  >
    <slot>{{ props.label }}</slot>
  </span>
</template>

<style scoped lang="css">
.capsule {
  border-radius: 12px;
  font-weight: 700;
  line-height: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  white-space: nowrap;
  text-overflow: ellipsis;
  max-width: 12rem;
  overflow: hidden;
  text-transform: uppercase;
  letter-spacing: 0.02em;
}

/* Base Themes (Pastel Background + Strong Text) */
.capsule--green {
  background-color: hsl(from var(--color-primary) h s 92%);
  color: var(--color-primary);
}

.capsule--blue {
  background-color: hsl(from var(--color-secondary) h s 92%);
  color: var(--color-secondary);
}

.capsule--orange {
  background-color: hsl(from var(--color-warning) h s 92%);
  color: var(--color-warning);
}

.capsule--red {
  background-color: hsl(from var(--color-danger) h s 92%);
  color: var(--color-danger);
}

.capsule--purple {
  background-color: hsl(from var(--color-tertiary) h s 92%);
  color: var(--color-tertiary);
}

.capsule--gray {
  background-color: hsl(from var(--color-neutral) h s 94%);
  color: hsl(from var(--color-neutral) h s 40%);
}

.capsule--white {
  background-color: #fff;
  color: hsl(from var(--color-neutral) h s 50%);
  border: 1px solid var(--border-color);
}

/* Sizes */
.capsule--sm {
  padding: 4px 10px;
  font-size: 0.7rem;
}

.capsule--md {
  padding: 6px 14px;
  font-size: 0.8rem;
}

.capsule--lg {
  padding: 8px 18px;
  font-size: 0.9rem;
}
</style>
