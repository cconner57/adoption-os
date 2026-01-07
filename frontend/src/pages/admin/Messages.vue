<script setup lang="ts"></script>

<template>
  <div class="admin-page">
    <div class="page-header">
      <h1>Messages</h1>
    </div>

    <div class="placeholder-content">
      <p>Inbox and messaging features coming soon...</p>
    </div>
  </div>
</template>

<style scoped>
.admin-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  h1 {
    font-size: 1.8rem;
    color: var(--font-color-dark);
    margin: 0;
  }
}

.placeholder-content {
  background: white;
  padding: 48px;
  border-radius: 12px;
  text-align: center;
  color: var(--font-color-medium);
  border: 1px dashed #e5e7eb;
}
</style>
