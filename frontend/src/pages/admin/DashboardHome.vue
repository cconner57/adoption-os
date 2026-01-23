<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

import DashboardListWidget from '../../components/admin/dashboard/DashboardListWidget.vue'
import StatisticsGrid from '../../components/admin/dashboard/StatisticsGrid.vue'
import WeeklySchedule from '../../components/admin/dashboard/WeeklySchedule.vue'
import { Button } from '../../components/common/ui'
import { useActiveVolunteersStore } from '../../stores/activeVolunteers'
import { useAuthStore } from '../../stores/auth'
import { usePetStore } from '../../stores/pets'

const authStore = useAuthStore()
const petStore = usePetStore()
const volunteerStore = useActiveVolunteersStore()
const userName = computed(() => authStore.user?.Name || 'Admin')
const pendingCount = ref(0)
const API_URL = import.meta.env.VITE_API_URL

// Calculate week start (Monday) for initial fetch
const getWeekRange = () => {
  const today = new Date()
  const day = today.getDay() // 0 = Sun
  const diff = today.getDate() - day + (day === 0 ? -6 : 1) // Adjust when day is Sunday
  const monday = new Date(today.setDate(diff))
  const sunday = new Date(today.setDate(diff + 6))

  return {
    start: monday.toISOString().split('T')[0],
    end: sunday.toISOString().split('T')[0],
  }
}

const fetchPendingCount = async () => {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_URL}/v1/applications?status=pending&page_size=1`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    if (res.ok) {
      const data = await res.json()
      pendingCount.value = data.metadata.total_records
    }
  } catch (e) {
    console.error('Failed to fetch pending count', e)
  }
}

onMounted(() => {
  petStore.fetchPets()
  volunteerStore.fetchActiveCount()
  const { start, end } = getWeekRange()
  volunteerStore.fetchWeeklyShifts(start, end)
  fetchPendingCount()
})

const adoptablePetsCount = computed(() => {
  return petStore.currentPets.length
})

// --- List Widget Data Adapters ---

const medicalNeedsItems = computed(() => [
  {
    id: 1,
    title: 'Apollo',
    subtitle: 'Neuter Surgery',
    statusLabel: 'Tomorrow',
    statusColor: 'hsl(from var(--color-warning) h s 90%)',
    statusTextColor: 'hsl(from var(--color-warning) h s 40%)',
  },
  {
    id: 2,
    title: 'Luna',
    subtitle: 'Rabies Booster',
    statusLabel: 'Overdue',
    statusColor: 'hsl(from var(--color-danger) h s 95%)',
    statusTextColor: 'var(--color-danger)',
  },
  {
    id: 3,
    title: 'Max',
    subtitle: 'Eye Infection Check',
    statusLabel: 'Today',
    statusColor: 'hsl(from var(--color-warning) h s 95%)',
    statusTextColor: 'var(--color-warning)',
  },
])

const inventoryAlertsItems = computed(() => [
  {
    id: 1,
    title: 'Kitten Formula',
    subtitle: '2 cans remaining',
    statusLabel: 'Critical',
    statusColor: 'hsl(from var(--color-danger) h s 95%)',
    statusTextColor: 'var(--color-danger)',
  },
  {
    id: 2,
    title: 'Paper Towels',
    subtitle: '4 rolls remaining',
    statusLabel: 'Low',
    statusColor: 'hsl(from var(--color-warning) h s 95%)',
    statusTextColor: 'var(--color-warning)',
  },
  {
    id: 3,
    title: 'Rabies Support',
    subtitle: '5 doses remaining',
    statusLabel: 'Low',
    statusColor: 'hsl(from var(--color-warning) h s 95%)',
    statusTextColor: 'var(--color-warning)',
  },
])

const activeVolunteersItems = computed(() => [
  {
    id: 1,
    title: 'Sarah J.',
    subtitle: 'Kennels',
    statusLabel: '2h 15m',
  },
  {
    id: 2,
    title: 'Mike T.',
    subtitle: 'Front Desk',
    statusLabel: '45m',
  },
  {
    id: 3,
    title: 'Emily R.',
    subtitle: 'Cats',
    statusLabel: '1h 10m',
  },
])

const recentApplicationsItems = computed(() => [
  {
    id: 1,
    title: 'John Doe',
    subtitle: 'for Buddy',
    statusLabel: 'new',
    statusColor: 'hsl(from var(--color-secondary) h s 95%)',
    statusTextColor: 'var(--color-secondary)',
  },
  {
    id: 2,
    title: 'Jane Smith',
    subtitle: 'for Luna',
    statusLabel: 'review',
    statusColor: 'hsl(from var(--color-tertiary) h s 95%)',
    statusTextColor: 'var(--color-tertiary)',
  },
  {
    id: 3,
    title: 'Robert B.',
    subtitle: 'for Max',
    statusLabel: 'pending',
    statusColor: 'hsl(from var(--color-warning) h s 95%)',
    statusTextColor: 'var(--color-warning)',
  },
])

const recentDonationsItems = computed(() => [
  {
    id: 1,
    title: '$50.00',
    subtitle: 'Alice W.',
    statusLabel: '2h ago',
  },
  {
    id: 2,
    title: '$100.00',
    subtitle: 'Tom H.',
    statusLabel: '5h ago',
  },
  {
    id: 3,
    title: '$25.00',
    subtitle: 'Anonymous',
    statusLabel: '1d ago',
  },
])

const newIntakesItems = computed(() => [
  {
    id: 1,
    title: 'Simba',
    subtitle: 'Tabby Cat',
    statusLabel: 'Stray',
    statusColor: 'hsl(from var(--color-warning) h s 90%)',
    statusTextColor: 'hsl(from var(--color-warning) h s 40%)',
  },
  {
    id: 2,
    title: 'Rocky',
    subtitle: 'Boxer Mix',
    statusLabel: 'Surrender',
    statusColor: 'hsl(from var(--color-warning) h s 90%)',
    statusTextColor: 'hsl(from var(--color-warning) h s 40%)',
  },
  {
    id: 3,
    title: 'Daisy',
    subtitle: 'Golden Ret.',
    statusLabel: 'Transfer',
    statusColor: 'hsl(from var(--color-warning) h s 90%)',
    statusTextColor: 'hsl(from var(--color-warning) h s 40%)',
  },
])
</script>

<template>
  <div class="dashboard-home">
    <div class="welcome-section">
      <h1>Hello, {{ userName }} 👋</h1>
      <p>Here's what's happening at the rescue today.</p>
    </div>

    <StatisticsGrid :pendingCount="pendingCount" :adoptablePetsCount="adoptablePetsCount" />

    <div class="dashboard-widgets">
      <WeeklySchedule />

      <div class="widget quick-actions">
        <h3>Quick Actions</h3>
        <div class="action-buttons">
          <Button variant="secondary" fullWidth align="between">
            New Pet Profile <span>→</span>
          </Button>
          <Button variant="secondary" fullWidth align="between">
            Review Applications <span>→</span>
          </Button>
          <Button variant="secondary" fullWidth align="between">
            Send Newsletter <span>→</span>
          </Button>
          <Button variant="secondary" fullWidth align="between">
            Log Incident <span>→</span>
          </Button>
        </div>
      </div>
    </div>

    <!-- Secondary Widgets Row -->
    <div class="widgets-row-three">
      <DashboardListWidget title="🏥 Urgent Medical" :items="medicalNeedsItems" />
      <DashboardListWidget title="⚠️ Low Inventory" :items="inventoryAlertsItems" />
      <DashboardListWidget title="📍 Active Now" :items="activeVolunteersItems" />
    </div>

    <!-- Third Widgets Row -->
    <div class="widgets-row-three">
      <DashboardListWidget title="📝 Recent Applications" :items="recentApplicationsItems" />
      <DashboardListWidget title="💰 Recent Donations" :items="recentDonationsItems" />
      <DashboardListWidget title="🐾 New Intakes" :items="newIntakesItems" />
    </div>
  </div>
</template>

<style scoped>
.dashboard-home {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.welcome-section h1 {
  font-size: 2rem;
  color: var(--text-primary);
  margin-bottom: 8px;
}

.welcome-section p {
  color: var(--text-secondary);
  font-size: 1.1rem;
}

.dashboard-widgets {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;

  @media (max-width: 900px) {
    grid-template-columns: 1fr;
  }
}

.widgets-row-three {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;

  @media (max-width: 900px) {
    grid-template-columns: 1fr;
  }
}

/* Quick Actions Widget Styles (Specific override for this special widget) */
.widget {
  background: var(--text-inverse);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
}

.widget h3 {
  font-size: 1.1rem;
  margin: 0;
  font-weight: 700;
  color: var(--text-primary);
}

.quick-actions {
  grid-column: span 1;
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
  margin-top: 16px;
}
</style>
