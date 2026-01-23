<script setup lang="ts">
import Button from '../components/common/ui/Button.vue'

const handleGoHome = () => {
  window.location.href = '/'
}
</script>

<template>
  <div class="not-found-container">
    <div class="content">
      <img src="/images/404.jpeg" alt="404 Not Found" class="not-found-image" />
      <h1>Oops! Page Not Found</h1>
      <p>The page you are looking for might have been removed or is temporarily unavailable.</p>
      <Button @click="handleGoHome" title="Return Home" color="orange" size="large" />
    </div>
  </div>
</template>

<style scoped>
.not-found-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
  padding: 2rem;
  padding-top: 8rem; 
  text-align: center;
}

@media (max-width: 768px) {
  .not-found-container {
    padding-top: 6rem; 
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
  max-width: 600px;
  position: relative;
  z-index: 1; 
}

.not-found-image {
  max-width: 100%;
  height: auto;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
  margin-top: 2rem;
}

h1 {
  color: var(--text-inverse);
  font-size: 2.5rem;
  font-weight: 800; 
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

p {
  color: var(--text-inverse); 
  font-size: 1.25rem; 
  font-weight: 600; 
  margin: 0 0 1rem 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3); 
  line-height: 1.5;
}
</style>
