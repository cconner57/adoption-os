<script setup lang="ts"></script>

<template>
  <div class="page-container">
    <h1>Transport & Dispatch</h1>
    <div class="placeholder-content">
      <p>
        Coordinate vet appointments, track volunteer drivers, and manage safe transport updates.
      </p>
      <ul>
        <li>Anonymous communication system</li>
        <li>Real-time trip status (En route, Arrived, Issue)</li>
        <li>ETA tracking for adopters</li>
      </ul>
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

ul {
  margin-top: 16px;
  padding-left: 20px;
}

li {
  margin-bottom: 8px;
}
</style>
