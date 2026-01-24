<script setup lang="ts">


const props = defineProps({
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
  <span :class="{ 'gold-border': props.type === 'Gold', 'silver-border': props.type === 'Silver' }">
    <p>{{ props.type }} Transparency</p>
    <p>{{ props.year }}</p>
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

  @media (width <= 440px) {
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
    rgb(255 215 0) 0 0 0 5px,
    var(--text-inverse) 0 0 0 8px,
    rgb(255 215 0) 0 0 0 14px;

  @media (width <= 440px) {
    padding: 8px;
  }
}

.silver-border {
  box-shadow:
    0 0 0 5px #c0c0c0,
    0 0 0 8px var(--text-inverse),
    0 0 0 14px #c0c0c0;

  @media (width <= 440px) {
    padding: 8px;
  }
}
</style>
