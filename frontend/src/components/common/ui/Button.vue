<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
  defineProps<{
    variant?: 'primary' | 'secondary' | 'tertiary' | 'text'
    theme?: 'primary' | 'secondary' | 'tertiary' | 'warning' | 'danger' | 'neutral'
    color?: 'green' | 'blue' | 'purple' | 'green-weak' | 'orange' | 'white'
    onClick?: () => void
    size?: 'small' | 'medium' | 'large'
    align?: 'center' | 'start' | 'between'
    title?: string
    fullWidth?: boolean
    disabled?: boolean
    loading?: boolean
  }>(),
  {
    variant: 'primary',
    size: 'medium',
    align: 'center',
  },
)

const resolvedTheme = computed(() => {
  if (props.theme) return props.theme

  switch (props.color) {
    case 'green':
    case 'green-weak':
      return 'primary'
    case 'blue':
      return 'secondary'
    case 'purple':
      return 'tertiary'
    case 'orange':
      return 'warning'
    case 'white':
      return 'neutral'
    default:
      return 'primary'
  }
})

const resolvedVariant = computed(() => {
  if (['primary', 'secondary', 'tertiary', 'text'].includes(props.variant)) {
    // Legacy support: if color is white/neutral and variant is primary, it should look like tertiary/neutral
    if (props.color === 'white' && props.variant === 'primary') {
      return 'tertiary'
    }
    return props.variant
  }
  return 'primary'
})

const classes = computed(() => {
  const base = [
    `variant-${resolvedVariant.value}`,
    `theme-${resolvedTheme.value}`,
    props.size,
    `justify-${props.align}`,
  ]

  if (props.color === 'green-weak') base.push('legacy-green-weak')
  if (props.fullWidth) base.push('w-full')
  if (props.disabled || props.loading) base.push('button-disabled')

  return base.join(' ')
})
</script>

<template>
  <button
    :class="classes"
    @click="props.onClick && props.onClick()"
    :disabled="props.disabled || props.loading"
  >
    <span v-if="props.loading" class="spinner"></span>
    <slot v-else>
      <span>{{ props.title }}</span>
    </slot>
  </button>
</template>

<style scoped lang="css">
button {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-weight: 600;
  border-radius: 6px;
  transition: all 0.2s ease;
  white-space: nowrap;
  border: 1px solid transparent;
  font-family: inherit;

  &:hover:not(:disabled) {
    cursor: pointer;
    transform: translateY(-1px);
    box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  }

  &:active:not(:disabled) {
    transform: translateY(0);
  }
}

.small {
  height: 32px;
  min-width: 80px;
  padding: 0 12px;
  font-size: 0.85rem;
}

.medium {
  height: 40px;
  min-width: 120px;
  padding: 0 20px;
  font-size: 0.95rem;
}

.large {
  height: 48px;
  min-width: 180px;
  padding: 0 32px;
  font-size: 1.1rem;
}

.justify-center {
  justify-content: center;
}

.justify-start {
  justify-content: flex-start;
}

.justify-between {
  justify-content: space-between;
}

.w-full {
  width: 100%;
  display: flex;
}

.variant-primary {
  color: #fff;
  border: none;
}

.variant-primary.theme-primary {
  background-color: var(--color-primary);
}

.variant-primary.theme-secondary {
  background-color: var(--color-secondary);
}

.variant-primary.theme-tertiary {
  background-color: var(--color-tertiary);
}

.variant-primary.theme-warning {
  background-color: var(--color-warning);
}

.variant-primary.theme-danger {
  background-color: var(--color-danger);
}

.variant-primary.theme-neutral {
  background-color: var(--color-neutral);
}

.variant-primary:hover:not(:disabled) {
  opacity: 0.9;
}

.variant-secondary {
  background-color: #fff;
  border: 1px solid currentcolor;
}

.variant-secondary.theme-primary {
  color: var(--color-primary);
}

.variant-secondary.theme-secondary {
  color: var(--color-secondary);
}

.variant-secondary.theme-tertiary {
  color: var(--color-tertiary);
}

.variant-secondary.theme-warning {
  color: var(--color-warning);
}

.variant-secondary.theme-danger {
  color: var(--color-danger);
}

.variant-secondary.theme-neutral {
  color: var(--text-primary);
  border-color: var(--border-color);
}

.variant-secondary:hover:not(:disabled) {
  background-color: #f8fafc;
}

.variant-tertiary {
  background-color: #fff;
  border: 1px solid #cbd5e1;
  color: var(--text-primary);
}

.variant-tertiary.theme-primary {
  color: var(--text-primary);
}

.variant-tertiary:hover:not(:disabled) {
  border-color: #94a3b8;
  background-color: #f1f5f9;
}

.variant-text {
  background: none;
  border: none;
  box-shadow: none !important;
  padding: 0;
  min-width: 0;
  height: auto;
  text-decoration: underline;
  text-underline-offset: 4px;
}

.variant-text:hover:not(:disabled) {
  transform: none;
  opacity: 0.8;
}

.variant-text.theme-primary {
  color: var(--color-primary);
}

.variant-text.theme-secondary {
  color: var(--color-secondary);
}

.variant-text.theme-tertiary {
  color: var(--color-tertiary);
}

.variant-text.theme-neutral {
  color: var(--text-secondary);
}

.legacy-green-weak {
  background-color: var(--color-primary-weak);
  color: var(--color-primary);
  border: none;
}

.legacy-green-weak:hover:not(:disabled) {
  background-color: var(--color-primary-border);
}

.button-disabled {
  filter: grayscale(100%);
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
  box-shadow: none !important;
}

.spinner {
  width: 1em;
  height: 1em;
  border: 2px solid currentcolor;
  border-right-color: transparent;
  border-radius: 50%;
  animation: spin 0.75s linear infinite;
  display: inline-block;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}
</style>
