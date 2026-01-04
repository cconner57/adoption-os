<script setup lang="ts">
import { computed } from 'vue'
const props = withDefaults(
  defineProps<{
    imgSrc: string
    title: string
    subtitle: string
    color: string
    type?: 'button' | 'submit' | 'reset'
  }>(),
  {
    imgSrc: '',
    title: 'Adopt a Pet',
    subtitle: 'Find your perfect companion',
    color: '#2563EB',
    type: 'button',
  },
)

const imgSrcComputed = computed(() => (props.imgSrc && props.imgSrc.length ? props.imgSrc : ''))
</script>

<template>
  <button
    class="banner-button"
    :class="{
      'button-color-blue': props.color === 'blue',
      'button-color-green': props.color === 'green',
      'button-color-purple': props.color === 'purple',
    }"
    :type="props.type"
    :aria-label="props.title"
  >
    <img :src="imgSrcComputed" :alt="props.title" height="48" width="48" />

    <div class="banner-text">
      <h5>{{ props.title }}</h5>
      <p>{{ props.subtitle }}</p>
    </div>
  </button>
</template>

<style scoped lang="css">
.banner-button {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  padding: 24px;
  border-radius: 8px;
  cursor: pointer;
  min-height: 160px;
  height: auto;
  width: 100%;
  transition:
    background-color 0.2s,
    border-color 0.2s;
  color: var(--font-color-light);
  overflow: hidden;

  @media (max-width: 440px) {
    width: 100%;
  }
}
.banner-text {
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 4px;
}
img {
  border-radius: 8px;
}
h5 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
}
p {
  margin: 0;
  font-size: 1rem;
}
</style>
