<script setup lang="ts">
import { computed,ref } from 'vue'

import type { IVolunteer } from '../../../stores/mockVolunteerData'
import { ButtonToggle, Capsules,InputField } from '../../common/ui'
const props = defineProps<{
  selectedId?: string
  volunteers: IVolunteer[]
  loading?: boolean
}>()

const emit = defineEmits<{
  select: [volunteer: IVolunteer]
  add: []
}>()

const searchQuery = ref('')
const filterType = ref<'active' | 'archived'>('active')
const sortType = ref<'alphabetical' | 'level'>('level')

const filteredVolunteers = computed(() => {
  const filtered = props.volunteers.filter((v) => {
    const matchesSearch = `${v.firstName} ${v.lastName}`
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

    const nameA = a.firstName.toLowerCase()
    const nameB = b.firstName.toLowerCase()
    return nameA.localeCompare(nameB)
  })
})

import { watch } from 'vue'

watch(
  () => filteredVolunteers.value,
  (newVal) => {
    const currentSelected = newVal.find((v) => String(v.id) === String(props.selectedId))
    if (!currentSelected && newVal.length > 0) {

      emit('select', newVal[0])
    }
  },
  { immediate: true },
)

function selectVolunteer(vol: IVolunteer) {
  emit('select', vol)
}
</script>

<template>
  <div class="vol-list-container">
    <div class="list-header">
      <div class="title-row">
        <h2>Volunteers</h2>
        <button class="add-btn" @click="emit('add')">+</button>
      </div>

      <div class="search-box">
        <InputField
          v-model="searchQuery"
          placeholder="Search volunteers..."
          style="margin-bottom: 0"
        />
      </div>

      <div class="list-tabs">
        <ButtonToggle
          label=""
          v-model="filterType"
          true-value="active"
          false-value="archived"
          true-label="Active"
          false-label="Archived"
        />
      </div>

      <div class="sort-row">
        <span class="sort-label">Sort by:</span>
        <div class="sort-options">
          <Capsules
            label="A-Z"
            :color="sortType === 'alphabetical' ? 'blue' : 'white'"
            style="cursor: pointer"
            @click="sortType = 'alphabetical'"
          />
          <Capsules
            label="Level"
            :color="sortType === 'level' ? 'blue' : 'white'"
            style="cursor: pointer"
            @click="sortType = 'level'"
          />
        </div>
      </div>
    </div>

    <div class="vol-list">

      <div v-if="loading" class="skeleton-list">
        <div v-for="n in 6" :key="n" class="vol-item skeleton-item">
          <div class="skeleton-avatar"></div>
          <div class="vol-info">
            <div class="skeleton-text w-60"></div>
            <div class="skeleton-text w-40"></div>
          </div>
        </div>
      </div>

      <div v-else-if="filteredVolunteers.length === 0" class="empty-list">
        No {{ filterType }} volunteers found.
      </div>

      <div
        v-else
        v-for="vol in filteredVolunteers"
        :key="vol.id"
        class="vol-item"
        :class="{ selected: String(vol.id) === selectedId }"
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
          :class="
            vol.reliabilityScore >= 80
              ? 'high'
              : vol.reliabilityScore > 0
                ? 'mid'
                : vol.reliabilityScore === 0
                  ? 'neutral'
                  : 'low'
          "
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
  color: #fff;
  border: none;
  font-size: 1.2rem;
  line-height: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: var(--color-secondary-strong);
  }
}

.list-tabs {
  margin-top: 8px;
}

.sort-row {
  display: flex;
  align-items: center;
  margin-top: 12px;
  justify-content: space-between;
}

.sort-label {
  font-size: 0.8rem;
  color: var(--color-neutral-text-soft);
}

.sort-options {
  display: flex;
  gap: 8px;
}

.vol-list {
  flex: 1;
  overflow-y: auto;
  padding: 12px;
}

.empty-list {
  text-align: center;
  color: var(--color-neutral-text-soft);
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
    background: var(--color-neutral-surface);
  }

  &.selected {
    background: #fff;
    box-shadow: 0 2px 8px rgb(0 0 0 / 5%);
    transform: scale(1.02);
    margin-bottom: 8px;
    margin-top: 4px;
    z-index: 1;

    .name {
      color: var(
        --text-primary
      );
    }
  }
}

.vol-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: var(--color-secondary-light);
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
    box-shadow: 0 0 4px oklch(from var(--color-primary) l c h / 0.40);
  }

  &.inactive {
    background: var(--color-neutral-border-strong);
  }

  &.pending {
    background: var(--color-warning);
  }
}

.role {
  font-size: 0.8rem;
  color: var(--color-neutral-text-soft);
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

  &.neutral {
    color: var(--color-neutral-text-soft);
  }

  &.low {
    color: var(--color-danger);
  }
}
</style>

<style scoped>

.skeleton-item {
  cursor: default;

  &:hover {
    background: transparent;
  }
}

.skeleton-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: var(--color-neutral-light);
  margin-right: 12px;
  animation: pulse 1.5s infinite;
}

.skeleton-text {
  height: 12px;
  background: var(--color-neutral-light);
  border-radius: 4px;
  margin-bottom: 6px;
  animation: pulse 1.5s infinite;

  &.w-60 {
    width: 60%;
  }

  &.w-40 {
    width: 40%;
  }
}

@keyframes pulse {
  0% {
    opacity: 1;
  }

  50% {
    opacity: 0.5;
  }

  100% {
    opacity: 1;
  }
}
</style>
