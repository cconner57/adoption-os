<template>
  <div class="skeleton-table-container">
    <table class="skeleton-table">
      <thead>
        <tr>
          <th v-for="i in columns" :key="i">
            <div class="skeleton-header"></div>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="r in rows" :key="r">
          <td v-for="c in columns" :key="c">
            <div class="skeleton-cell"></div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
withDefaults(
  defineProps<{
    rows?: number
    columns?: number
  }>(),
  {
    rows: 10,
    columns: 5,
  },
)
</script>

<style scoped>
.skeleton-table-container {
  width: 100%;
  background: white;
  border-radius: 12px;
  overflow: hidden;
}

.skeleton-table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 16px 24px;
  border-bottom: 1px solid var(--border-color, #e2e8f0);
}

.skeleton-header {
  height: 24px;
  width: 60%;
  background: #f1f5f9;
  border-radius: 4px;
}

.skeleton-cell {
  height: 20px;
  width: 80%;
  background: #f1f5f9;
  border-radius: 4px;
  animation: pulse 1.5s infinite ease-in-out;
}

@keyframes pulse {
  0% {
    opacity: 0.6;
  }
  50% {
    opacity: 0.3;
  }
  100% {
    opacity: 0.6;
  }
}
</style>
