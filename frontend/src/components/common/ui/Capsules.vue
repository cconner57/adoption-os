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
  background-color: var(--color-primary-weak);
  color: var(--color-primary-strong);
}

.capsule--blue {
  background-color: var(--color-secondary-weak);
  color: var(--color-secondary-strong);
}

.capsule--orange {
  background-color: var(--color-warning-weak);
  color: var(--color-warning-strong);
}

.capsule--red {
  background-color: var(--color-danger-weak);
  color: var(--color-danger-strong);
}

.capsule--purple {
  background-color: var(--color-tertiary-weak);
  color: var(--color-tertiary-strong);
}

.capsule--gray {
  background-color: var(--color-neutral-weak);
  color: var(--color-neutral-strong);
}

.capsule--white {
  background-color: #fff;
  color: var(--color-neutral-text-soft);
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
