<script setup lang="ts">
defineProps<{
  title: string
  icon: string
}>()
</script>

<template>
  <section class="card settings-card">
    <div class="card-header">
      <div class="icon-box">{{ icon }}</div>
      <h3>{{ title }}</h3>
    </div>
    <div class="card-content">
      <slot></slot>
    </div>
  </section>
</template>

<style scoped>

.settings-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden;
}

.card-header {
  padding: 16px 24px;
  border-bottom: 1px solid #f3f4f6;
  display: flex;
  align-items: center;
  gap: 12px;
  background: #f9fafb;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    font-weight: 600;
    color: var(--text-primary);
  }
}

.icon-box {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 1.2rem;
}

.card-content {
  padding: 24px;
}
</style>
