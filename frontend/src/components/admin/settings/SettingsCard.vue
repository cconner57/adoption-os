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
/* Consolidate card styles if they aren't globally available,
   but for now we assume global 'card' class or scoped style persistence.
   Since the original used scoped styles, we might need to rely on the parent or global styles.
   However, assuming 'card' is a global utility class or we should copy the scoped styles if they were local.
   Checking original Settings.vue... it has scoped styles but 'settings-card' seems specific.
*/
.settings-card {
  background: white;
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
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 1.2rem;
}

.card-content {
  padding: 24px;
}
</style>
