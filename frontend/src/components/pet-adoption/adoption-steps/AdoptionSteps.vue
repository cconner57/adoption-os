<script setup lang="ts">
const { formStep, selectedAnimal } = defineProps<{
  formStep: number
  selectedAnimal: 'dog' | 'cat' | null
}>()
</script>

<template>
  <div class="steps-container">
    <div
      class="line"
      :class="{ activeLine: formStep >= 1 }"
      :style="{ width: selectedAnimal === 'dog' ? '100px' : '80px' }"
    ></div>
    <div class="step" :class="{ active: formStep >= 0 }">
      <div class="step-number">1</div>
      <div class="step-label">General</div>
    </div>
    <div class="step" :class="{ active: formStep >= 1 }">
      <div class="step-number">2</div>
      <div class="step-label">Home</div>
    </div>
    <div class="step" :class="{ active: formStep >= 2 }">
      <div class="step-number">3</div>
      <div class="step-label">New Cat</div>
    </div>
    <div class="step" v-if="selectedAnimal === 'cat'" :class="{ active: formStep >= 3 }">
      <div class="step-number">4</div>
      <div class="step-label">Current Pets</div>
    </div>
    <div class="step" :class="{ active: formStep >= 4 }">
      <div class="step-number">{{ selectedAnimal === 'dog' ? 4 : 5 }}</div>
      <div class="step-label">Past Pets</div>
    </div>
    <div class="step" :class="{ active: formStep >= 5 }">
      <div class="step-number">{{ selectedAnimal === 'dog' ? 5 : 6 }}</div>
      <div class="step-label">Other</div>
    </div>
    <div class="step" :class="{ active: formStep >= 6 }">
      <div class="step-number">{{ selectedAnimal === 'dog' ? 6 : 7 }}</div>
      <div class="step-label">Summary</div>
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
  align-items: center;
  position: relative;

  & .line {
    position: absolute;
    top: 14px;
    left: 20px;
    right: 20px;
    height: 2px;
    background-color: var(--green);
    z-index: 1;
    width: auto !important;
    max-width: none;
    min-width: 0;
  }

  @media (max-width: 600px) {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
    padding-left: 16px;

    /* Hide horizontal line on mobile */
    & .line {
      display: none;
    }
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    z-index: 2; /* Above line */

    @media (max-width: 600px) {
      flex-direction: row;
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
      z-index: 5;

      @media (max-width: 600px) {
        margin-bottom: 0;
      }
    }
    .step-label {
      font-size: 0.875rem;
      text-align: center;

      @media (max-width: 600px) {
        font-size: 1rem;
        font-weight: 500;
        text-align: left;
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
