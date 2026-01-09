<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockVolunteers, type IVolunteer } from '../../../stores/mockVolunteerData'

const props = defineProps<{
  selectedId?: string
}>()

const emit = defineEmits<{
  (e: 'select', volunteer: IVolunteer): void
}>()

const searchQuery = ref('')
const filterType = ref<'active' | 'archived'>('active')
const sortType = ref<'alphabetical' | 'level'>('level')

const filteredVolunteers = computed(() => {
  const filtered = mockVolunteers.filter((v) => {
    const matchesSearch = (`${v.firstName  } ${  v.lastName}`)
      .toLowerCase()
      .includes(searchQuery.value.toLowerCase())

    if (filterType.value === 'active') {
      return matchesSearch && (v.status === 'active' || v.status === 'pending')
    } else {
      return matchesSearch && v.status === 'inactive'
    }
  })

  return filtered.sort((a, b) => {
    if (sortType.value === 'level') {
      const weights: Record<string, number> = { 'Tier 2': 3, 'Tier 1': 2, Teen: 1, Admin: 4 }
      const weightA = weights[a.role] || 0
      const weightB = weights[b.role] || 0
      if (weightA !== weightB) return weightB - weightA
    }

    // Default / Tie-breaker is alphabetical
    const nameA = a.firstName.toLowerCase()
    const nameB = b.firstName.toLowerCase()
    return nameA.localeCompare(nameB)
  })
})

function selectVolunteer(vol: IVolunteer) {
  emit('select', vol)
}
</script>

<template>
  <div class="vol-list-container">
    <div class="list-header">
      <div class="title-row">
        <h2>Volunteers</h2>
        <button class="add-btn">+</button>
      </div>

      <div class="search-box">
        <input v-model="searchQuery" type="text" placeholder="Search volunteers..." />
      </div>

      <!-- Filter Tabs -->
      <div class="list-tabs">
        <button
          class="list-tab"
          :class="{ active: filterType === 'active' }"
          @click="filterType = 'active'"
        >
          Active
        </button>
        <button
          class="list-tab"
          :class="{ active: filterType === 'archived' }"
          @click="filterType = 'archived'"
        >
          Archived
        </button>
      </div>

      <!-- Sort Controls -->
      <div class="sort-row">
        <span class="sort-label">Sort by:</span>
        <div class="sort-options">
          <button
            class="sort-chip"
            :class="{ active: sortType === 'alphabetical' }"
            @click="sortType = 'alphabetical'"
          >
            A-Z
          </button>
          <button
            class="sort-chip"
            :class="{ active: sortType === 'level' }"
            @click="sortType = 'level'"
          >
            Level
          </button>
        </div>
      </div>
    </div>

    <div class="vol-list">
      <div v-if="filteredVolunteers.length === 0" class="empty-list">
        No {{ filterType }} volunteers found.
      </div>

      <div
        v-for="vol in filteredVolunteers"
        :key="vol.id"
        class="vol-item"
        :class="{ selected: vol.id === selectedId }"
        @click="selectVolunteer(vol)"
      >
        <div class="vol-avatar">{{ vol.firstName[0] }}{{ vol.lastName[0] }}</div>
        <div class="vol-info">
          <div class="name-row">
            <span class="name">{{ vol.firstName }} {{ vol.lastName }}</span>
            <span class="status-dot" :class="vol.status"></span>
          </div>
          <div class="role">{{ vol.role }}</div>
        </div>
        <div
          class="vol-score"
          :class="vol.reliabilityScore >= 90 ? 'high' : vol.reliabilityScore >= 70 ? 'mid' : 'low'"
        >
          {{ vol.reliabilityScore }}%
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.vol-list-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: var(--text-inverse);
  border-right: 1px solid var(--border-color);
}

.list-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-color);
}

.title-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;

  h2 {
    margin: 0;
    font-size: 1.2rem;
    color: var(--text-primary);
  }
}

.add-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background: var(--color-secondary);
  color: white;
  border: none;
  font-size: 1.2rem;
  line-height: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: hsl(from var(--color-secondary) h s 40%);
  }
}

.search-box input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  background: hsl(from var(--color-neutral) h s 98%);
  font-size: 0.9rem;

  &:focus {
    outline: none;
    border-color: var(--color-secondary);
    background: var(--text-inverse);
  }
}

.list-tabs {
  display: flex;
  margin-top: 12px;
  background: hsl(from var(--color-neutral) h s 95%);
  padding: 4px;
  border-radius: 8px;
}

.list-tab {
  flex: 1;
  border: none;
  background: none;
  padding: 6px;
  font-size: 0.85rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;

  &.active {
    background: var(--text-inverse);
    color: var(--text-primary);
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  }
}

.sort-row {
  display: flex;
  align-items: center;
  margin-top: 12px;
  justify-content: space-between;
}

.sort-label {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.sort-options {
  display: flex;
  gap: 8px;
}

.sort-chip {
  background: none;
  border: 1px solid var(--border-color);
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.75rem;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  transition: all 0.2s;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
  }

  &.active {
    background: hsl(from var(--color-secondary) h s 90%);
    color: var(--color-secondary);
    border-color: hsl(from var(--color-secondary) h s 90%);
    font-weight: 600;
  }
}

.vol-list {
  flex: 1;
  overflow-y: auto;
  padding: 12px;
}

.empty-list {
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
  padding: 24px;
  font-size: 0.9rem;
}

.vol-item {
  display: flex;
  align-items: center;
  padding: 12px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s;
  margin-bottom: 4px;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
  }

  &.selected {
    background: hsl(from var(--color-secondary) h s 95%);
    .name {
      color: var(--color-secondary);
    }
  }
}

.vol-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: hsl(from var(--color-secondary) h s 90%);
  color: var(--color-secondary);
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12px;
  font-size: 0.9rem;
}

.vol-info {
  flex: 1;
}

.name-row {
  display: flex;
  align-items: center;
  margin-bottom: 4px;
}

.name {
  font-weight: 600;
  font-size: 0.95rem;
  color: var(--text-primary);
  margin-right: 6px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;

  &.active {
    background: var(--color-primary);
    box-shadow: 0 0 4px hsla(from var(--color-primary) h s l / 0.4);
  }
  &.inactive {
    background: hsl(from var(--color-neutral) h s 80%);
  }
  &.pending {
    background: var(--color-warning);
  }
}

.role {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.vol-score {
  font-weight: 700;
  font-size: 0.9rem;

  &.high {
    color: var(--color-primary);
  }
  &.mid {
    color: var(--color-warning);
  }
  &.low {
    color: var(--color-danger);
  }
}
</style>
