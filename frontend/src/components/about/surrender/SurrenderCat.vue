<script setup lang="ts">
// Importing the router to handle navigation
import { useRouter } from 'vue-router'

import Button from '../../common/ui/Button.vue'

const router = useRouter()

// Function to navigate to the surrender form page
const goToSurrenderForm = () => {
  router.push('/surrender')
}
</script>

<template>
  <div class="surrender-cat-section">
    <h2>Surrender Pet</h2>
    <p class="description">
      If you need to surrender your pet, we understand this can be a difficult decision. Please fill
      out our surrender form to help us find the best solution for your pet.
    </p>
    <Button
      class="surrender-button"
      title="Fill Out Surrender Form"
      @click="goToSurrenderForm"
      color="purple"
      size="medium"
    />
  </div>
</template>

<style scoped lang="css">
h2 {
  font-size: 2.5rem;
  color: var(--text-inverse);
  text-align: center;
  margin: 50px 0 10px;
}

.surrender-cat-section {
  text-align: center;
  background-color: var(--color-primary);
  padding: 80px var(--layout-padding-side);
  display: flex;
  flex-direction: column;
  align-items: center;
}

.description {
  font-size: 1.2rem;
  color: var(--text-inverse);
  margin-bottom: 50px;
  text-align: center;
  font-weight: 400;
}

.surrender-button {
  margin-bottom: 60px;
}
</style>
