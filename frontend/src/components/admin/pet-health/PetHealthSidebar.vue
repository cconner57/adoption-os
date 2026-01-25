<script setup lang="ts">
export interface IPetHealthItem {
  id: string
  name: string
  species: string
  photos?: { url: string; isPrimary?: boolean }[]
  latestLog?: { date: string }
}

defineProps<{
  pets: IPetHealthItem[]
  modelValue: string | null
  searchQuery: string
}>()

const emit = defineEmits<{
  'update:modelValue': [id: string]
  'update:searchQuery': [query: string]
  'add-log': []
}>()

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}
</script>

<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <h2>Pet Health</h2>
      <button class="add-log-btn" title="Add Vitals record" @click="emit('add-log')">+</button>
    </div>
    <div class="search-box">
      <input
        :value="searchQuery"
        @input="emit('update:searchQuery', ($event.target as HTMLInputElement).value)"
        type="text"
        placeholder="Search pets..."
      />
    </div>

    <div class="pet-list">
      <div
        v-for="pet in pets"
        :key="pet.id"
        class="pet-card"
        :class="{ active: modelValue === pet.id }"
        @click="emit('update:modelValue', pet.id)"
      >
        <div class="pet-avatar">
          <img
            v-if="pet.photos && pet.photos.length > 0"
            :src="`/images/${pet.photos?.find((p) => p.isPrimary)?.url ?? ''}`"
            alt="Pet Avatar"
          />
          <span v-else>{{ pet.species === 'cat' ? '🐱' : '🐶' }}</span>
        </div>
        <div class="pet-info">
          <div class="pet-name">{{ pet.name }}</div>
          <div class="pet-meta">
            <span v-if="pet.latestLog">Last update: {{ formatDate(pet.latestLog.date) }}</span>
            <span v-else>No records</span>
          </div>
        </div>
        <div class="chevron">›</div>
      </div>
    </div>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 320px;
  background: #fff;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 4px 12px rgb(0 0 0 / 3%);
  border: 1px solid #f3f4f6;
  flex-shrink: 0;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid #f3f4f6;
  display: flex;
  justify-content: space-between;
  align-items: center;

  h2 {
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--text-primary);
    margin: 0;
  }
}

.add-log-btn {
  background: var(--color-primary);
  color: #fff;
  border: none;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: opacity 0.2s;

  &:hover {
    opacity: 0.9;
  }
}

.search-box {
  padding: 16px;
  border-bottom: 1px solid #f3f4f6;

  input {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    background: #f9fafb;
    outline: none;

    &:focus {
      border-color: var(--color-secondary);
      background: #fff;
    }
  }
}

.pet-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.pet-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;

  &:hover {
    background: #f9fafb;
  }

  &.active {
    background: #eff6ff;
    border-color: hsl(from var(--color-secondary) h s 80%);
  }
}

.pet-avatar {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  background: #e5e7eb;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.pet-info {
  flex: 1;
}

.pet-name {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 0.95rem;
}

.pet-meta {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  margin-top: 2px;
}

.chevron {
  color: #cbd5e1;
  font-size: 1.2rem;
}
</style>
