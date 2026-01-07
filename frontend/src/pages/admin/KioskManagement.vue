<script setup lang="ts"></script>

<template>
  <div class="page-container">
    <h1>Kiosk Management</h1>
    <div class="placeholder-content">
      <p>Configure settings and UI for the adoption center kiosk.</p>
    </div>
  </div>
</template>

<style scoped>
.page-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

h1 {
  color: var(--font-color-dark);
  font-size: 1.8rem;
  margin: 0;
}

.placeholder-content {
  background: white;
  padding: 32px;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  color: var(--font-color-medium);
}
</style>
