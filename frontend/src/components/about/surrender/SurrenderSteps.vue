<script setup lang="ts">
const { formStep, selectedAnimal } = defineProps<{
  formStep: number
  selectedAnimal: 'dog' | 'cat' | null
}>()
</script>

<template>
  <ol class="steps-container">
    <li
      class="step"
      :class="{ active: formStep >= 1 }"
      :aria-current="formStep === 1 ? 'step' : undefined"
    >
      <div class="step-number">1</div>
      <div class="step-label">Household</div>
    </li>
    <li
      class="step"
      :class="{ active: formStep >= 2 }"
      :aria-current="formStep === 2 ? 'step' : undefined"
    >
      <div class="step-number">2</div>
      <div class="step-label">Behavior</div>
    </li>
    <li
      class="step"
      :class="{ active: formStep >= 3 }"
      :aria-current="formStep === 3 ? 'step' : undefined"
    >
      <div class="step-number">3</div>
      <div class="step-label">Aggression</div>
    </li>
    <li
      class="step"
      v-if="selectedAnimal === 'cat'"
      :class="{ active: formStep >= 4 }"
      :aria-current="formStep === 4 ? 'step' : undefined"
    >
      <div class="step-number">4</div>
      <div class="step-label">Medical</div>
    </li>
    <li
      class="step"
      :class="{ active: formStep >= 5 }"
      :aria-current="formStep === (selectedAnimal === 'cat' ? 5 : 4) ? 'step' : undefined"
    >
      <div class="step-number">{{ selectedAnimal === 'cat' ? 5 : 4 }}</div>
      <div class="step-label">Feeding</div>
    </li>
    <li
      class="step"
      :class="{ active: formStep >= 6 }"
      :aria-current="formStep === (selectedAnimal === 'cat' ? 6 : 5) ? 'step' : undefined"
    >
      <div class="step-number">{{ selectedAnimal === 'cat' ? 6 : 5 }}</div>
      <div class="step-label">Other</div>
    </li>
  </ol>
</template>

<style scoped lang="css">
.steps-container {
  display: flex;
  justify-content: space-between;
  list-style: none;
  padding: 0;
  width: 100%;
  max-width: 600px;
  margin: 0 auto 20px;
  align-items: center; 
  position: relative; 

  &::before {
    content: '';
    position: absolute;
    top: 15px; 
    left: 27px; 
    right: 27px; 
    height: 2px;
    background-color: var(--color-primary);
    z-index: 0; 
  }

  @media (width <= 600px) {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
    padding-left: 16px;

    &::before {
      display: none;
    }
  }

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

    @media (width <= 600px) {
      flex-direction: row;
      margin: 0;
      gap: 12px;
      width: 100%;
    }

    .step-number {
      width: 30px;
      height: 30px;
      border-radius: 50%;
      background-color: var(--text-inverse);
      border: 1px solid var(--color-primary);
      color: var(--text-primary);
      display: flex;
      justify-content: center;
      align-items: center;
      margin-bottom: 8px;
      z-index: 5; 

      @media (width <= 600px) {
        margin-bottom: 0;
      }
    }

    .step-label {
      font-size: 0.75rem;
      text-align: center;
      white-space: nowrap;

      @media (width <= 600px) {
        font-size: 1rem;
        font-weight: 500;
      }
    }

    &.active {
      .step-number {
        background-color: var(--color-primary);
        color: var(--text-inverse);
      }
    }
  }
}
</style>
