<script setup lang="ts">
const props = defineProps<{
  color: 'green' | 'blue' | 'purple' | 'green-weak'
  onClick?: () => void
  size?: 'small' | 'medium' | 'large'
  title: string
  fullWidth?: boolean
  disabled?: boolean
}>()
</script>

<template>
  <button
    :class="{
      'button-color-blue': props.color === 'blue',
      'button-color-green': props.color === 'green',
      'button-color-purple': props.color === 'purple',
      'button-green-weak': props.color === 'green-weak',
      small: props.size === 'small',
      medium: props.size === 'medium' || !props.size,
      large: props.size === 'large',
      'w-full': props.fullWidth,
      'button-disabled': props.disabled,
    }"
    @click="props.onClick && props.onClick()"
    :disabled="props.disabled"
  >
    {{ props.title }}
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
  &:hover {
    cursor: pointer;
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

.button-green-weak {
  background-color: var(--green-weak);
  &:hover {
    background-color: var(--green-weak-hover);
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
