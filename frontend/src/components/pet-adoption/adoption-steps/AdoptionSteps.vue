<script setup lang="ts">
const { formStep, selectedAnimal } = defineProps<{
  formStep: number
  selectedAnimal: 'dog' | 'cat' | null
}>()
</script>

<template>
  <div class="steps-container">
    <div
    <div class="line"
      :class="{ activeLine: formStep >= 1 }"
      :style="{ width: selectedAnimal === 'dog' ? '100px' : '80px' }"
    ></div>
    <div class="step" :class="{ active: formStep >= 0, completed: formStep > 0 }">
      <div class="step-number">
        <svg v-if="formStep > 0" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <span v-else>1</span>
      </div>
      <div class="step-label">General</div>
    </div>
    <div class="step" :class="{ active: formStep >= 1, completed: formStep > 1 }">
      <div class="step-number">
        <svg v-if="formStep > 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <span v-else>2</span>
      </div>
      <div class="step-label">Home</div>
    </div>
    <div class="step" :class="{ active: formStep >= 2, completed: formStep > 2 }">
      <div class="step-number">
        <svg v-if="formStep > 2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <span v-else>3</span>
      </div>
      <div class="step-label">New Cat</div>
    </div>
    <div class="step" v-if="selectedAnimal === 'cat'" :class="{ active: formStep >= 3, completed: formStep > 3 }">
      <div class="step-number">
        <svg v-if="formStep > 3" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <span v-else>4</span>
      </div>
      <div class="step-label">Current Pets</div>
    </div>
    <div class="step" :class="{ active: formStep >= 4, completed: formStep > 4 }">
      <div class="step-number">
        <svg v-if="formStep > 4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <div v-else>{{ selectedAnimal === 'dog' ? 4 : 5 }}</div>
      </div>
      <div class="step-label">Past Pets</div>
    </div>
    <div class="step" :class="{ active: formStep >= 5, completed: formStep > 5 }">
      <div class="step-number">
        <svg v-if="formStep > 5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <div v-else>{{ selectedAnimal === 'dog' ? 5 : 6 }}</div>
      </div>
      <div class="step-label">Other</div>
    </div>
    <div class="step" :class="{ active: formStep >= 6, completed: formStep > 6 }">
      <div class="step-number">
        <svg v-if="formStep > 6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="16" height="16"><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/></svg>
        <div v-else>{{ selectedAnimal === 'dog' ? 6 : 7 }}</div>
      </div>
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
    background-color: var(--color-primary);
    z-index: 1;
    width: auto !important;
    max-width: none;
    min-width: 0;
  }

  @media (width <= 600px) {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
    padding-left: 16px;

    & .line {
      display: none;
    }

    &::before {
      content: '';
      position: absolute;
      top: 0;
      bottom: 0;
      left: 31px; /* 16px padding-left + 15px (half of 30px width) */
      width: 2px;
      background-color: var(--color-primary);
      z-index: 1;
    }
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    z-index: 2;

    @media (width <= 600px) {
      flex-direction: row;
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
      font-size: 0.875rem;
      text-align: center;

      @media (width <= 600px) {
        font-size: 1rem;
        font-weight: 500;
        text-align: left;
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
