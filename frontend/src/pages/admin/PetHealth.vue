<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import PetHealthSidebar from '../../components/admin/pet-health/PetHealthSidebar.vue'
import RecentLogsTable from '../../components/admin/pet-health/RecentLogsTable.vue'
import VitalsCard from '../../components/admin/pet-health/VitalsCard.vue'
import WeightChart from '../../components/common/WeightChart.vue'
import { mockHealthLogs } from '../../stores/mockHealthLogs'
import { mockPetsData } from '../../stores/mockPetData'

const searchQuery = ref('')
const selectedPetId = ref<string | null>(null)
const chartTimeframe = ref<'1M' | '3M' | '6M' | '1Y' | 'ALL'>('ALL')

const petsWithLogs = computed(() => {
  return mockPetsData
    .map((pet) => {
      const logs = mockHealthLogs.filter((l) => l.petId === pet.id)
      const latestLog = logs.sort(
        (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime(),
      )[0]
      return {
        ...pet,
        latestLog,
        hasLogs: logs.length > 0,
      }
    })
    .filter((p) => p.hasLogs)
    .sort((a, b) => {
      const dateA = a.latestLog ? new Date(a.latestLog.date).getTime() : 0
      const dateB = b.latestLog ? new Date(b.latestLog.date).getTime() : 0
      return dateB - dateA
    })
})

const filteredPets = computed(() => {
  if (!searchQuery.value) return petsWithLogs.value
  const q = searchQuery.value.toLowerCase()
  return petsWithLogs.value.filter((p) => p.name.toLowerCase().includes(q))
})

const selectedPet = computed(() => {
  if (!selectedPetId.value) return null
  return petsWithLogs.value.find((p) => p.id === selectedPetId.value)
})

const selectedPetLogs = computed(() => {
  if (!selectedPetId.value) return []
  return mockHealthLogs
    .filter((l) => l.petId === selectedPetId.value)
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
})

const filteredChartLogs = computed(() => {
  if (chartTimeframe.value === 'ALL') return selectedPetLogs.value

  const now = new Date()
  const cutoff = new Date()

  if (chartTimeframe.value === '1M') cutoff.setMonth(now.getMonth() - 1)
  if (chartTimeframe.value === '3M') cutoff.setMonth(now.getMonth() - 3)
  if (chartTimeframe.value === '6M') cutoff.setMonth(now.getMonth() - 6)
  if (chartTimeframe.value === '1Y') cutoff.setFullYear(now.getFullYear() - 1)

  return selectedPetLogs.value.filter((l) => new Date(l.date) >= cutoff)
})

const weightHistory = computed(() => {
  return filteredChartLogs.value
    .map((l) => ({
      date: l.date,
      value: (l.weight || 0) / 1000,
    }))
    .reverse()
})

const latestVitals = computed(() => {
  return selectedPetLogs.value[0] || null
})

function formatWeight(grams?: number | null) {
  if (!grams) return '-'
  const kg = (grams / 1000).toFixed(2)
  const lbs = (grams * 0.00220462).toFixed(2)
  return `${kg}kg (${lbs}lbs)`
}

watch(
  filteredPets,
  (newPets) => {
    if (!selectedPetId.value && newPets.length > 0) {
      selectedPetId.value = newPets[0].id
    }
  },
  { immediate: true },
)
</script>

<template>
  <div class="health-page">
    <Teleport to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Pet Health</h1>
    </Teleport>
    <PetHealthSidebar
      :pets="filteredPets"
      v-model="selectedPetId"
      v-model:searchQuery="searchQuery"
    />

    <main class="dashboard">
      <div v-if="selectedPet" class="dashboard-content">
        <header class="dash-header">
          <div class="dash-title">
            <h1>{{ selectedPet.name }}</h1>
            <span class="breed">{{ selectedPet.physical.breed }} • {{ selectedPet.sex }}</span>
          </div>
          <div class="actions">
            <button class="action-btn">History</button>
            <button class="action-btn primary">Record Vitals +</button>
          </div>
        </header>

        <div class="vitals-grid">
          <div class="stat-card weight-card">
            <div class="card-header">
              <h3>Weight Tracking</h3>
              <span class="current-value">{{ formatWeight(latestVitals?.weight) }}</span>
            </div>
            <div class="chart-wrapper">
              <WeightChart :data="weightHistory" />
            </div>
            <div class="footer-filters">
              <div class="timeframe-filters">
                <button
                  v-for="time in ['1M', '3M', '6M', '1Y', 'ALL']"
                  :key="time"
                  class="filter-btn"
                  :class="{ active: chartTimeframe === time }"
                  @click="chartTimeframe = time as any"
                >
                  {{ time === 'ALL' ? 'Lifetime' : time }}
                </button>
              </div>
            </div>
          </div>

          <VitalsCard :latestVitals="latestVitals" />
        </div>

        <RecentLogsTable :logs="selectedPetLogs" />
      </div>
      <div v-else class="empty-state">
        <div class="empty-icon">🐾</div>
        <h2>Select a pet to view health records</h2>
        <p>Choose a pet from the list on the left to see their weight history and vitals.</p>
      </div>
    </main>
  </div>
</template>

<style scoped>
.health-page {
  display: flex;
  height: calc(100vh - 80px);
  gap: 24px;
  max-width: 1400px;
  margin: 0 auto;
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.dashboard {
  flex: 1;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 3%);
  border: 1px solid #f3f4f6;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.dashboard-content {
  padding: 32px;
  max-width: 1200px;
  margin: 0 auto;
  height: 100%;
  overflow-y: auto;
}

.dash-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.dash-title h1 {
  font-size: 2rem;
  margin: 0 0 4px;
  color: var(--text-primary);
}

.breed {
  color: hsl(from var(--color-neutral) h s 50%);
  font-size: 0.95rem;
}

.actions {
  display: flex;
  gap: 12px;
}

.action-btn {
  padding: 8px 16px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  background: #fff;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;

  &:hover {
    background: #f9fafb;
  }

  &.primary {
    background: var(--color-secondary);
    color: #fff;
    border: none;

    &:hover {
      background: hsl(from var(--color-secondary) h s 40%);
    }
  }
}

.vitals-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 24px;
  margin-bottom: 32px;
}

.stat-card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  padding: 24px;
  border: 1px solid #f3f4f6;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--text-primary);
  }
}

.current-value {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--text-primary);
}

.footer-filters {
  display: flex;
  justify-content: center;
  margin-top: 16px;
}

.timeframe-filters {
  display: flex;
  gap: 4px;
  background: #f3f4f6;
  padding: 4px;
  border-radius: 8px;
}

.filter-btn {
  border: none;
  background: transparent;
  padding: 4px 12px;
  border-radius: 6px;
  font-size: 0.8rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;

  &:hover {
    color: var(--text-primary);
    background: rgb(0 0 0 / 5%);
  }

  &.active {
    background: #fff;
    color: var(--color-secondary);
    box-shadow: 0 2px 4px rgb(0 0 0 / 5%);
  }
}

.chart-wrapper {
  height: 300px;
  width: 100%;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);

  h2 {
    color: var(--text-primary);
    margin-bottom: 8px;
  }
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 24px;
  opacity: 0.5;
}
</style>
