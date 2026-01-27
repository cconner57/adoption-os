<template>
  <div class="pet-card-skeleton">
    <div class="skeleton-header">
      <div class="skeleton-avatar"></div>
      <div class="skeleton-info">
        <div class="skeleton-row top">
          <div class="skeleton-text title"></div>
          <div class="skeleton-text badge"></div>
        </div>
        <div class="skeleton-row bottom">
          <div class="skeleton-pill"></div>
          <div class="skeleton-pill"></div>
          <div class="skeleton-pill"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.pet-card-skeleton {
  background: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 2px 8px rgb(0 0 0 / 5%);
  overflow: hidden;
  border: 1px solid var(--color-neutral-border);
  margin-bottom: 16px;
  padding: 16px;
}

.skeleton-header {
  display: flex;
  gap: 12px;
}

.skeleton-avatar {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  background: var(--color-neutral-weak);
  flex-shrink: 0;
  animation: pulse 1.5s infinite ease-in-out;
}

.skeleton-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  justify-content: center;
}

.skeleton-row {
  display: flex;
  gap: 8px;
  align-items: center;
}

.skeleton-text {
  height: 20px;
  background: var(--color-neutral-weak);
  border-radius: 4px;
  animation: pulse 1.5s infinite ease-in-out;
}

.skeleton-text.title {
  width: 120px;
  height: 24px;
}

.skeleton-text.badge {
  width: 80px;
}

.skeleton-pill {
  width: 60px;
  height: 24px;
  background: var(--color-neutral-weak);
  border-radius: 6px;
  animation: pulse 1.5s infinite ease-in-out;
}

@keyframes pulse {
  0% { opacity: 0.6; }
  50% { opacity: 0.3; }
  100% { opacity: 0.6; }
}
</style>
