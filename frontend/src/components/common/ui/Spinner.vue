<template>
  <div class="spinner"></div>
</template>

<style scoped>
.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgb(0 0 0 / 10%);
  border-left-color: var(--color-primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
