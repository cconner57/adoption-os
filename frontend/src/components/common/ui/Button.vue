<script setup lang="ts">
const props = defineProps<{
  color: 'green' | 'blue' | 'purple' | 'green-weak' | 'orange' | 'white'
  onClick?: () => void
  size?: 'small' | 'medium' | 'large'
  align?: 'center' | 'start' | 'between'
  title?: string
  fullWidth?: boolean
  disabled?: boolean
  loading?: boolean
}>()
</script>

<template>
  <button
    :class="{
      'button-color-blue': props.color === 'blue',
      'button-color-green': props.color === 'green',
      'button-color-purple': props.color === 'purple',
      'button-color-orange': props.color === 'orange',
      'button-green-weak': props.color === 'green-weak',
      'button-color-white': props.color === 'white',
      small: props.size === 'small',
      medium: props.size === 'medium' || !props.size,
      large: props.size === 'large',
      'justify-center': !props.align || props.align === 'center',
      'justify-start': props.align === 'start',
      'justify-between': props.align === 'between',
      'w-full': props.fullWidth,
      'button-disabled': props.disabled || props.loading,
    }"
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
  font-weight: 600;
  border-radius: 6px;
  color: var(--text-inverse);
  transition: background-color 0.2s;
  white-space: nowrap; /* Force single line */
  &:hover {
    cursor: pointer;
  }
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

.spinner {
  width: 1em; /* Scales with font size */
  height: 1em;
  border: 2px solid currentColor;
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

.small {
  height: 30px;
  min-width: 120px;
  padding: 0 16px;
  font-size: 0.875rem;
}

.medium {
  height: 40px;
  min-width: 160px;
  padding: 0 24px;
  font-size: 1rem;

  @media (max-width: 480px) {
    padding: 0 16px; /* Slightly reduced padding but comfortable */
    font-size: 0.95rem; /* Readable size, close to desktop */
    min-width: 140px;
  }
}

.large {
  height: 60px;
  min-width: 220px;
  padding: 0 32px;
  font-size: 1.5rem;

  @media (max-width: 768px) {
    height: 50px;
    font-size: 1.25rem;
    padding: 0 24px;
  }
}

.button-color-blue {
  background-color: var(--color-secondary);

  &:hover {
    background-color: hsl(from var(--color-secondary) h s 40%);
  }
}

.button-color-purple {
  background-color: var(--color-tertiary);
  &:hover {
    background-color: hsl(from var(--color-tertiary) h s 40%);
  }
}

.button-color-green {
  background-color: var(--color-primary);
  &:hover {
    background-color: hsl(from var(--color-primary) h s 40%);
  }
}

.button-color-orange {
  background-color: var(--color-warning);
  &:hover {
    background-color: hsl(from var(--color-warning) h s 40%);
  }
}

.button-green-weak {
  background-color: hsl(from var(--color-primary) h s 95%);
  color: var(--color-primary);
  &:hover {
    background-color: hsl(from var(--color-primary) h s 80%);
  }
}

.button-color-white {
  background-color: var(--text-inverse);
  color: var(--text-primary);
  border: 1px solid var(--border-color);
  &:hover {
    background-color: hsl(from var(--text-inverse) h s 95%);
  }
}

.w-full {
  width: 100%;
}

.button-disabled {
  filter: grayscale(60%);
  cursor: not-allowed;
  &:hover {
    filter: grayscale(70%);
  }
}
</style>
