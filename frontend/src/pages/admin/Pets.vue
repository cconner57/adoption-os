<script setup lang="ts"></script>

<template>
  <div class="admin-page">
    <div class="page-header">
      <h1>Pets Management</h1>
      <button class="add-btn">Add New Pet +</button>
    </div>

    <div class="placeholder-content">
      <p>Pet management features coming soon...</p>
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

.add-btn {
  background-color: var(--green);
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: var(--green-hover);
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
