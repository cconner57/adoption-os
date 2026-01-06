<script setup lang="ts">
const props = defineProps<{
  color: 'green' | 'blue' | 'purple' | 'green-weak' | 'orange' | 'white'
  onClick?: () => void
  size?: 'small' | 'medium' | 'large'
  title: string
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
      'w-full': props.fullWidth,
      'button-disabled': props.disabled || props.loading,
    }"
    @click="props.onClick && props.onClick()"
    :disabled="props.disabled || props.loading"
  >
    <span v-if="props.loading" class="spinner"></span>
    <span v-else>{{ props.title }}</span>
  </button>
</template>

<style scoped lang="css">
button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  border-radius: 6px;
  color: var(--white);
  transition: background-color 0.2s;
  white-space: nowrap; /* Force single line */
  &:hover {
    cursor: pointer;
  }
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
  background-color: var(--blue);

  &:hover {
    background-color: var(--blue-hover);
  }
}

.button-color-purple {
  background-color: var(--purple);
  &:hover {
    background-color: var(--purple-hover);
  }
}

.button-color-green {
  background-color: var(--green);
  &:hover {
    background-color: var(--green-hover);
  }
}

.button-color-orange {
  background-color: #ed8936;
  &:hover {
    background-color: #dd6b20;
  }
}

.button-green-weak {
  background-color: var(--green-weak);
  &:hover {
    background-color: var(--green-weak-hover);
  }
}

.button-color-white {
  background-color: var(--white);
  color: var(--font-color-dark);
  border: 1px solid #e2e8f0;
  &:hover {
    background-color: #f8fafc;
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
