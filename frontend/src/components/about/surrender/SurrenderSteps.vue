<script setup lang="ts">
const { formStep, selectedAnimal } = defineProps<{
  formStep: number
  selectedAnimal: 'dog' | 'cat' | null
}>()
</script>

<template>
  <div class="steps-container">
    <div class="step" :class="{ active: formStep >= 1 }">
      <div class="step-number">1</div>
      <div class="step-label">Household</div>
    </div>
    <div class="step" :class="{ active: formStep >= 2 }">
      <div class="step-number">2</div>
      <div class="step-label">Behavior</div>
    </div>
    <div class="step" :class="{ active: formStep >= 3 }">
      <div class="step-number">3</div>
      <div class="step-label">Aggression</div>
    </div>
    <div class="step" v-if="selectedAnimal === 'cat'" :class="{ active: formStep >= 4 }">
      <div class="step-number">4</div>
      <div class="step-label">Medical</div>
    </div>
    <div
      class="step"
      :class="{ active: formStep >= 5 }"
    >
      <div class="step-number">{{ selectedAnimal === 'cat' ? 5 : 4 }}</div>
      <div class="step-label">Feeding</div>
    </div>
    <div
      class="step"
      :class="{ active: formStep >= 6 }"
    >
      <div class="step-number">{{ selectedAnimal === 'cat' ? 6 : 5 }}</div>
      <div class="step-label">Other</div>
    </div>
  </div>
</template>

<style scoped lang="css">
.steps-container {
  display: flex;
  justify-content: space-between;
  width: 100%;
  max-width: 600px;
  margin: 0 auto 20px;
  align-items: center; /* Vertically align steps */
  position: relative; /* Establish positioning context for line */

  /* The continuous line */
  &::before {
    content: '';
    position: absolute;
    top: 15px; /* Center of 30px circle */
    left: 27px; /* 12px margin + 15px half circle */
    right: 27px; /* Match left offset */
    height: 2px;
    background-color: var(--green);
    z-index: 0; /* Behind step numbers */
  }

  @media (max-width: 600px) {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
    padding-left: 16px;

    /* Hide line on mobile */
    &::before {
      display: none;
    }
  }

  /* Hide scrollbar */
  -ms-overflow-style: none;
  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none;
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    flex-shrink: 0;
    margin: 0 12px;

    @media (max-width: 600px) {
      flex-direction: row;
      margin: 0;
      gap: 12px;
      width: 100%;
    }

    .step-number {
      width: 30px;
      height: 30px;
      border-radius: 50%;
      background-color: var(--white);
      border: 1px solid var(--green);
      color: var(--font-color-dark);
      display: flex;
      justify-content: center;
      align-items: center;
      margin-bottom: 8px;
      z-index: 5; /* Above the line */

      @media (max-width: 600px) {
        margin-bottom: 0;
      }
    }

    .step-label {
      font-size: 0.75rem;
      text-align: center;
      white-space: nowrap;

      @media (max-width: 600px) {
        font-size: 1rem;
        font-weight: 500;
      }
    }

    &.active {
      .step-number {
        background-color: var(--green);
        color: var(--white);
      }
    }
  }
}
</style>
