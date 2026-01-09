<script setup lang="ts">
import { defineProps } from 'vue'

defineProps({
  type: {
    type: String,
    required: true,
  },
  year: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <span :class="{ 'gold-border': type === 'Gold', 'silver-border': type === 'Silver' }">
    <p>{{ type }} Transparency</p>
    <p>{{ year }}</p>
    <p>Candid.</p>
  </span>
</template>

<style scoped lang="css">
span {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 8px;
  color: var(--text-primary);
  height: 132px;
  width: 159px;

  & p {
    font-size: 1.05rem;
  }
  & p:not(:first-child) {
    font-weight: 600;
  }

  @media (max-width: 440px) {
    height: auto;
    width: auto;
    & p {
      font-size: 1rem;
      margin-left: 0;
    }
  }
}

.gold-border {
  padding: 12px;
  box-shadow:
    rgb(255, 215, 0) 0px 0px 0px 5px,
    var(--text-inverse) 0px 0px 0px 8px,
    rgb(255, 215, 0) 0px 0px 0px 14px;

  @media (max-width: 440px) {
    padding: 8px;
  }
}

.silver-border {
  box-shadow:
    0 0 0 5px silver,
    0 0 0 8px var(--text-inverse),
    0 0 0 14px silver;

  @media (max-width: 440px) {
    padding: 8px;
  }
}
</style>
