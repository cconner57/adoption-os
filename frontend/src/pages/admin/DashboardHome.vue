<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Capsules, Button } from '../../components/common/ui'
import { useAuthStore } from '../../stores/auth'
import { usePetStore } from '../../stores/pets'
import { useActiveVolunteersStore } from '../../stores/activeVolunteers'

const authStore = useAuthStore()
const petStore = usePetStore()
const volunteerStore = useActiveVolunteersStore()
const userName = computed(() => authStore.user?.Name || 'Admin')
const pendingCount = ref(0)
const API_URL = import.meta.env.VITE_API_URL

// Helper format time
const formatTime = (timeStr: string) => {
  if (!timeStr) return ''
  // Handle 14:30:00 or 14:30
  const [h, m] = timeStr.split(':')
  const hour = parseInt(h)
  const ampm = hour >= 12 ? 'PM' : 'AM'
  const hour12 = hour % 12 || 12
  return `${hour12}:${m} ${ampm}`
}

// Calculate week start (Monday)
const getWeekRange = () => {
  const today = new Date()
  const day = today.getDay() // 0 = Sun
  const diff = today.getDate() - day + (day === 0 ? -6 : 1) // Adjust when day is Sunday
  const monday = new Date(today.setDate(diff))
  const sunday = new Date(today.setDate(diff + 6))

  return {
    start: monday.toISOString().split('T')[0],
    end: sunday.toISOString().split('T')[0],
    mondayObject: monday,
  }
}

onMounted(() => {
  petStore.fetchPets()
  volunteerStore.fetchActiveCount()
  const { start, end } = getWeekRange()
  volunteerStore.fetchWeeklyShifts(start, end)
  fetchPendingCount()
})

const fetchPendingCount = async () => {
  try {
    const token = localStorage.getItem('token')
    const res = await fetch(`${API_URL}/v1/applications?status=pending&page_size=1`, {
      // Page size 1 is enough to get metadata total
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

const adoptablePetsCount = computed(() => {
  return petStore.currentPets.length
})

// Helper to determine event style based on role
const getEventType = (role: string = '') => {
  const r = role.toLowerCase()
  if (r.includes('feeding')) return 'feeding'
  if (r.includes('adoption')) return 'adoption'
  if (r.includes('pickup') || r.includes('dropoff') || r.includes('transport')) return 'transport'
  if (r.includes('vet') && !r.includes('pickup') && !r.includes('dropoff')) return 'vet'
  return 'default'
}

// Real Calendar Data
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const { mondayObject } = getWeekRange()
  const shifts = volunteerStore.weeklyShifts || []

  return days.map((name, index) => {
    // Clone monday to avoid mutation issues
    const currentDay = new Date(mondayObject)
    currentDay.setDate(mondayObject.getDate() + index)
    const dateStr = currentDay.toISOString().split('T')[0]

    // Filter shifts for this day
    const dayShifts = shifts
      .filter((s) => s.date === dateStr)
      .map((s) => ({
        id: s.id,
        name: s.volunteerName,
        time: `${formatTime(s.startTime)} - ${formatTime(s.endTime)}`,
        type: getEventType(s.role),
        role: s.role,
      }))

    return {
      name,
      date: currentDay.getDate(),
      isToday: currentDay.toDateString() === new Date().toDateString(),
      events: dayShifts,
    }
  })
})

const stats = computed(() => [
  { label: 'Pending Applications', value: pendingCount.value, color: 'orange', icon: '📝' },
  { label: 'Adoptable Pets', value: adoptablePetsCount.value, color: 'green', icon: '🐾' },
  { label: 'Volunteers', value: volunteerStore.activeCount, color: 'purple', icon: '🤝' },
  { label: 'Donations (Month)', value: '$3,250', color: 'blue', icon: '❤️' },
])

const medicalNeeds = [
  { id: 1, pet: 'Apollo', issue: 'Neuter Surgery', date: 'Tomorrow', status: 'urgent' },
  { id: 2, pet: 'Luna', issue: 'Rabies Booster', date: 'Overdue', status: 'critical' },
  { id: 3, pet: 'Max', issue: 'Eye Infection Check', date: 'Today', status: 'warning' },
]

const inventoryAlerts = [
  { id: 1, item: 'Kitten Formula', level: 'Critical', quantity: '2 cans' },
  { id: 2, item: 'Paper Towels', level: 'Low', quantity: '4 rolls' },
  { id: 3, item: 'Rabies Support', level: 'Low', quantity: '5 doses' },
]

const activeVolunteers = [
  { id: 1, name: 'Sarah J.', role: 'Kennels', time: '2h 15m' },
  { id: 2, name: 'Mike T.', role: 'Front Desk', time: '45m' },
  { id: 3, name: 'Emily R.', role: 'Cats', time: '1h 10m' },
]

const recentApplications = [
  { id: 1, applicant: 'John Doe', pet: 'Buddy', status: 'new' },
  { id: 2, applicant: 'Jane Smith', pet: 'Luna', status: 'review' },
  { id: 3, applicant: 'Robert B.', pet: 'Max', status: 'pending' },
]

const recentDonations = [
  { id: 1, donor: 'Alice W.', amount: '$50.00', time: '2h ago' },
  { id: 2, donor: 'Tom H.', amount: '$100.00', time: '5h ago' },
  { id: 3, donor: 'Anonymous', amount: '$25.00', time: '1d ago' },
]

const newIntakes = [
  { id: 1, pet: 'Simba', breed: 'Tabby Cat', status: 'Stray' },
  { id: 2, pet: 'Rocky', breed: 'Boxer Mix', status: 'Surrender' },
  { id: 3, pet: 'Daisy', breed: 'Golden Ret.', status: 'Transfer' },
]
</script>

<template>
  <div class="dashboard-home">
    <div class="welcome-section">
      <h1>Hello, {{ userName }} 👋</h1>
      <p>Here's what's happening at the rescue today.</p>
    </div>

    <div class="stats-grid">
      <div
        v-if="volunteerStore.error"
        style="
          grid-column: 1 / -1;
          background: #fee2e2;
          color: #b91c1c;
          padding: 12px;
          border-radius: 8px;
          margin-bottom: 16px;
        "
      >
        Error loading data: {{ volunteerStore.error }}
      </div>
      <div v-for="stat in stats" :key="stat.label" class="stat-card" :class="`color-${stat.color}`">
        <div class="stat-icon">{{ stat.icon }}</div>
        <div class="stat-info">
          <span class="stat-value">{{ stat.value }}</span>
          <span class="stat-label">{{ stat.label }}</span>
        </div>
      </div>
    </div>

    <div class="dashboard-widgets">
      <div class="widget weekly-schedule">
        <div class="widget-header">
          <h3>Weekly Schedule</h3>
          <div class="legend">
            <span class="legend-item">Feeding <span class="dot feeding"></span></span>
            <span class="legend-item">Adoption <span class="dot adoption"></span></span>
            <span class="legend-item">Transport <span class="dot transport"></span></span>
            <span class="legend-item">Vet <span class="dot vet"></span></span>
            <span class="legend-item">Other <span class="dot default"></span></span>
          </div>
        </div>

        <div class="calendar-grid">
          <div
            v-for="day in weekDays"
            :key="day.date"
            class="calendar-day"
            :class="{ today: day.isToday }"
          >
            <div class="day-header">
              <span class="day-name">{{ day.name }}</span>
              <span class="day-date">{{ day.date }}</span>
            </div>
            <div class="day-events">
              <div
                v-for="event in day.events"
                :key="event.id"
                class="event-pill"
                :class="event.type"
                :title="`${event.name} (${event.time})`"
              >
                <span class="event-name">{{ event.name }}</span>
                <span class="event-time">{{ event.time }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

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
      <!-- Urgent Medical -->
      <div class="widget">
        <div class="widget-header">
          <h3>🏥 Urgent Medical</h3>
        </div>
        <ul class="list-widget">
          <li v-for="item in medicalNeeds" :key="item.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ item.pet }}</span>
              <span class="item-subtitle">{{ item.issue }}</span>
            </div>
            <Capsules
              size="sm"
              :label="item.date"
              :color="
                item.status === 'critical'
                  ? 'hsl(from var(--color-danger) h s 95%)'
                  : item.status === 'urgent'
                    ? 'hsl(from var(--color-warning) h s 90%)'
                    : 'hsl(from var(--color-warning) h s 95%)'
              "
              :textColor="
                item.status === 'critical'
                  ? 'var(--color-danger)'
                  : item.status === 'urgent'
                    ? 'hsl(from var(--color-warning) h s 40%)'
                    : 'var(--color-warning)'
              "
            />
          </li>
        </ul>
      </div>

      <!-- Inventory Alerts -->
      <div class="widget">
        <div class="widget-header">
          <h3>⚠️ Low Inventory</h3>
        </div>
        <ul class="list-widget">
          <li v-for="item in inventoryAlerts" :key="item.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ item.item }}</span>
              <span class="item-subtitle">{{ item.quantity }} remaining</span>
            </div>
            <Capsules
              v-if="item.level === 'Critical'"
              size="sm"
              label="Critical"
              color="hsl(from var(--color-danger) h s 95%)"
              textColor="var(--color-danger)"
            />
            <Capsules
              v-else
              size="sm"
              label="Low"
              color="hsl(from var(--color-warning) h s 95%)"
              textColor="var(--color-warning)"
            />
          </li>
        </ul>
      </div>

      <!-- Active Volunteers -->
      <div class="widget">
        <div class="widget-header">
          <h3>📍 Active Now</h3>
        </div>
        <ul class="list-widget">
          <li v-for="vol in activeVolunteers" :key="vol.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ vol.name }}</span>
              <span class="item-subtitle">{{ vol.role }}</span>
            </div>
            <Capsules size="sm" :label="vol.time" />
          </li>
        </ul>
      </div>
    </div>

    <!-- Third Widgets Row -->
    <div class="widgets-row-three">
      <!-- Recent Applications -->
      <div class="widget">
        <div class="widget-header">
          <h3>📝 Recent Applications</h3>
        </div>
        <ul class="list-widget">
          <li v-for="app in recentApplications" :key="app.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ app.applicant }}</span>
              <span class="item-subtitle">for {{ app.pet }}</span>
            </div>
            <Capsules
              size="sm"
              :label="app.status"
              :color="
                app.status === 'new'
                  ? 'hsl(from var(--color-secondary) h s 95%)'
                  : app.status === 'review'
                    ? 'hsl(from var(--color-tertiary) h s 95%)'
                    : 'hsl(from var(--color-warning) h s 95%)'
              "
              :textColor="
                app.status === 'new'
                  ? 'var(--color-secondary)'
                  : app.status === 'review'
                    ? 'var(--color-tertiary)'
                    : 'var(--color-warning)'
              "
            />
          </li>
        </ul>
      </div>

      <!-- Recent Donations -->
      <div class="widget">
        <div class="widget-header">
          <h3>💰 Recent Donations</h3>
        </div>
        <ul class="list-widget">
          <li v-for="donation in recentDonations" :key="donation.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ donation.amount }}</span>
              <span class="item-subtitle">{{ donation.donor }}</span>
            </div>
            <Capsules size="sm" :label="donation.time" />
          </li>
        </ul>
      </div>

      <!-- New Intakes -->
      <div class="widget">
        <div class="widget-header">
          <h3>🐾 New Intakes</h3>
        </div>
        <ul class="list-widget">
          <li v-for="intake in newIntakes" :key="intake.id" class="list-item">
            <div class="item-main">
              <span class="item-title">{{ intake.pet }}</span>
              <span class="item-subtitle">{{ intake.breed }}</span>
            </div>
            <Capsules
              size="sm"
              :label="intake.status"
              color="hsl(from var(--color-warning) h s 90%)"
              textColor="hsl(from var(--color-warning) h s 40%)"
            />
          </li>
        </ul>
      </div>
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

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
}

.stat-card {
  background: var(--text-inverse);
  padding: 24px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s;

  &:hover {
    transform: translateY(-2px);
    border-color: #cbd5e1;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  }
}

.stat-icon {
  font-size: 2rem;
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: hsl(from var(--color-neutral) h s 95%);
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
}

.stat-label {
  color: var(--text-secondary);
  font-size: 0.9rem;
}

/* Color variants */
/* Color variants - Darkened backgrounds (88% lightness) for better contrast */
.stat-card.color-orange .stat-icon {
  background-color: hsl(from var(--color-warning) h s 88%);
  color: var(--color-warning);
}
.stat-card.color-green .stat-icon {
  background-color: hsl(from var(--color-primary) h s 88%);
  color: var(--color-primary);
}
.stat-card.color-purple .stat-icon {
  background-color: hsl(from var(--color-secondary) h s 88%);
  color: var(--color-secondary);
}
.stat-card.color-blue .stat-icon {
  background-color: hsl(from var(--color-secondary) h s 85%);
  color: var(--color-secondary);
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

/* Updated Widget Header */
.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 12px;
}

/* Calendar Styles */
.weekly-schedule {
  min-height: 340px;
  grid-column: span 3;

  @media (max-width: 900px) {
    grid-column: span 1;
  }
}

.quick-actions {
  grid-column: span 1;
}

.legend {
  display: flex;
  gap: 12px;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.85rem;
  color: var(--text-secondary);
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot.feeding {
  background-color: var(--color-secondary);
}
.dot.adoption {
  background-color: var(--color-tertiary);
}
.dot.transport {
  background-color: var(--color-warning);
}
.dot.vet {
  background-color: var(--color-danger);
}
.dot.default {
  background-color: var(--text-secondary);
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 12px;
  overflow-x: auto;
  padding-bottom: 8px;
  flex: 1;
}

.calendar-day {
  background-color: hsl(from var(--color-neutral) h s 98%);
  border-radius: 8px;
  padding: 8px;
  min-height: 120px;
  border: 1px solid transparent;
  display: flex;
  flex-direction: column;
  gap: 8px;

  &.today {
    background-color: var(--text-inverse);
    border-color: hsl(from var(--color-secondary) h s 70%);
    box-shadow: 0 0 0 2px hsla(from var(--color-secondary) h s l / 0.1);

    .day-header {
      color: var(--color-secondary);
    }
  }
}

.day-header {
  text-align: center;
  display: flex;
  flex-direction: column;
  color: var(--text-secondary);
  font-size: 0.8rem;
  font-weight: 600;
  padding-bottom: 6px;
  border-bottom: 1px solid #f3f4f6;

  .day-date {
    font-size: 1rem;
    color: var(--text-primary);
    margin-top: 2px;
  }
}

.day-events {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.event-pill {
  padding: 4px 6px;
  border-radius: 4px;
  font-size: 0.7rem;
  display: flex;
  flex-direction: column;
  gap: 1px;
  border-left: 2px solid transparent;

  &.feeding {
    background-color: hsl(from var(--color-secondary) h s 96%);
    color: var(--color-secondary);
    border-left-color: var(--color-secondary);
  }

  &.adoption {
    background-color: hsl(from var(--color-tertiary) h s 96%);
    color: var(--color-tertiary);
    border-left-color: var(--color-tertiary);
  }

  &.transport {
    background-color: hsl(from var(--color-warning) h s 96%);
    color: hsl(from var(--color-warning) h s 40%);
    border-left-color: var(--color-warning);
  }

  &.vet {
    background-color: hsl(from var(--color-danger) h s 96%);
    color: var(--color-danger);
    border-left-color: var(--color-danger);
  }

  &.default {
    background-color: #f3f4f6;
    color: var(--text-secondary);
    border-left-color: var(--text-secondary);
  }

  .event-name {
    font-weight: 600;
    font-size: 0.75rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    line-height: 1.2;
  }

  .event-time {
    font-size: 0.65rem;
    opacity: 0.85;
    font-weight: 400;
  }
}

/* Quick Actions */
.quick-actions {
  display: flex;
  flex-direction: column;
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
  margin-top: 16px;
}

/* List Widgets (Medical, Inventory, Volunteers) */
.list-widget {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-color);

  &:last-child {
    border-bottom: none;
    padding-bottom: 0;
  }
}

.item-main {
  display: flex;
  flex-direction: column;
}

.item-title {
  font-family: 600;
  font-size: 0.95rem;
  color: var(--text-primary);
}

.item-subtitle {
  font-size: 0.85rem;
  color: var(--text-secondary);
}
</style>
