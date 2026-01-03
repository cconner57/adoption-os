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
  width: 600px;
  margin: 0 auto 20px;
  align-items: center;
  position: relative;
  & .line {
    flex-grow: 1;
    height: 2px;
    background-color: var(--green);
    margin: 0 5px;
    width: 550px;
    min-width: 550px;
    max-width: 550px;
    position: absolute;
    top: 14px;
    left: 20px;
    z-index: 1;
  }
  .step {
    display: flex;
    flex-direction: column;
    align-items: center;

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
    }
    .step-label {
      font-size: 0.875rem;
      text-align: center;
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
