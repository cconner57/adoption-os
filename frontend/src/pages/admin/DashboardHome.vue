<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '../../stores/auth'
import { mockPetsData } from '../../stores/mockPetData'

const authStore = useAuthStore()
const userName = computed(() => authStore.user?.name || 'Admin')

const adoptablePetsCount = computed(() => {
  return mockPetsData.filter((pet) => pet.details.status === 'available').length
})

// Recurring Schedule Definition
const recurringShifts: Record<number, Array<{ time: string; title: string; type: string }>> = {
  0: [
    // Sunday
    { time: '10:00 AM', title: 'Allison (10-12)', type: 'volunteer' },
    { time: '4:00 PM', title: 'Brandon (4-6)', type: 'volunteer' },
  ],
  1: [
    // Monday
    { time: '10:00 AM', title: 'Leanne (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Arianna (6-8)', type: 'volunteer' },
  ],
  2: [
    // Tuesday
    { time: '10:00 AM', title: 'Sonia (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Lynn (6-8)', type: 'volunteer' },
  ],
  3: [
    // Wednesday
    { time: '10:00 AM', title: 'Bella (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katelyn (6-8)', type: 'volunteer' },
  ],
  4: [
    // Thursday
    { time: '10:00 AM', title: 'Alejandra (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Nathan (6-8)', type: 'volunteer' },
  ],
  5: [
    // Friday
    { time: '10:00 AM', title: 'Linda (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katie (6-8)', type: 'volunteer' },
  ],
  6: [
    // Saturday
    { time: '10:00 AM', title: 'Lindsey & Celina (10-12)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Chris (6-8)', type: 'volunteer' },
  ],
}

// Mock Calendar Data
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const today = new Date()

  // Find Monday of current week
  const day = today.getDay()
  const diff = today.getDate() - day + (day === 0 ? -6 : 1)
  const monday = new Date(today.setDate(diff))

  return days.map((name, index) => {
    const date = new Date(monday)
    date.setDate(monday.getDate() + index)

    // Mock events based on day index
    const dayOfWeek = date.getDay()
    const events = []

    // Add recurring shifts
    const shifts = recurringShifts[dayOfWeek] || []
    shifts.forEach((shift, idx) => {
      events.push({ id: `v-${index}-${idx}`, ...shift })
    })

    // Add mock vet appointments
    if (dayOfWeek === 2) {
      // Tuesday
      events.push({
        id: `vet-${index}`,
        type: 'vet',
        time: '2:00 PM',
        title: 'Amari - Vaccination',
      })
    }
    if (dayOfWeek === 4) {
      // Thursday
      events.push({ id: `vet-${index}`, type: 'vet', time: '11:30 AM', title: 'Review: Apollo' })
    }

    return {
      name,
      date: date.getDate(),
      isToday: date.toDateString() === new Date().toDateString(),
      events,
    }
  })
})

const stats = computed(() => [
  { label: 'Pending Applications', value: 12, color: 'orange', icon: '📝' },
  { label: 'Adoptable Pets', value: adoptablePetsCount.value, color: 'green', icon: '🐾' },
  { label: 'Volunteers', value: 28, color: 'purple', icon: '🤝' },
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
            <span class="legend-item"><span class="dot volunteer"></span> Volunteer</span>
            <span class="legend-item"><span class="dot vet"></span> Vet</span>
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
                :title="event.title"
              >
                <span class="event-time">{{ event.time }}</span>
                <span class="event-title">{{ event.title }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="widget quick-actions">
        <h3>Quick Actions</h3>
        <div class="action-buttons">
          <button class="action-btn">New Pet Profile</button>
          <button class="action-btn">Review Applications</button>
          <button class="action-btn">Send Newsletter</button>
          <button class="action-btn">Log Incident</button>
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
            <span class="status-tag" :class="item.status">{{ item.date }}</span>
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
            <span class="status-tag critical" v-if="item.level === 'Critical'">Critical</span>
            <span class="status-tag warning" v-else>Low</span>
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
            <span class="time-tag">{{ vol.time }}</span>
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
            <span class="status-tag" :class="app.status">{{ app.status }}</span>
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
            <span class="time-tag">{{ donation.time }}</span>
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
            <span class="status-tag urgent">{{ intake.status }}</span>
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
  max-width: 1200px;
  margin: 0 auto;
}

.welcome-section h1 {
  font-size: 2rem;
  color: var(--font-color-dark);
  margin-bottom: 8px;
}

.welcome-section p {
  color: var(--font-color-medium);
  font-size: 1.1rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
}

.stat-card {
  background: var(--white);
  padding: 24px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s;

  &:hover {
    transform: translateY(-2px);
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
  background-color: #f3f4f6;
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--font-color-dark);
}

.stat-label {
  color: var(--font-color-medium);
  font-size: 0.9rem;
}

/* Color variants */
.stat-card.color-orange .stat-icon {
  background-color: #fff7ed;
  color: #ed8936;
}
.stat-card.color-green .stat-icon {
  background-color: #f0fdf4;
  color: var(--green);
}
.stat-card.color-purple .stat-icon {
  background-color: #faf5ff;
  color: var(--purple);
}
.stat-card.color-blue .stat-icon {
  background-color: #eff6ff;
  color: var(--blue);
}

.dashboard-widgets {
  display: grid;
  grid-template-columns: 2fr 1fr;
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
  background: var(--white);
  padding: 24px;
  border-radius: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
}

.widget h3 {
  font-size: 1.1rem;
  margin: 0;
  font-weight: 700;
  color: var(--font-color-dark);
}

/* Updated Widget Header */
.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 12px;
}

/* Calendar Styles */
.weekly-schedule {
  min-height: 340px;
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
  color: var(--font-color-medium);
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.dot.volunteer {
  background-color: var(--purple);
}
.dot.vet {
  background-color: var(--green);
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
  background-color: #fcfcfc;
  border-radius: 8px;
  padding: 8px;
  min-height: 120px;
  border: 1px solid transparent;
  display: flex;
  flex-direction: column;
  gap: 8px;

  &.today {
    background-color: #fff;
    border-color: var(--blue-weak);
    box-shadow: 0 0 0 2px rgba(25, 118, 210, 0.1);

    .day-header {
      color: var(--blue);
    }
  }
}

.day-header {
  text-align: center;
  display: flex;
  flex-direction: column;
  color: var(--font-color-medium);
  font-size: 0.8rem;
  font-weight: 600;
  padding-bottom: 6px;
  border-bottom: 1px solid #f3f4f6;

  .day-date {
    font-size: 1rem;
    color: var(--font-color-dark);
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

  &.volunteer {
    background-color: #faf5ff;
    color: var(--purple);
    border-left-color: var(--purple);
  }

  &.vet {
    background-color: #f0fdf4;
    color: var(--green);
    border-left-color: var(--green);
  }

  .event-time {
    font-weight: 700;
    opacity: 0.8;
    font-size: 0.65rem;
  }

  .event-title {
    line-height: 1.2;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
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
}

.action-btn {
  padding: 14px;
  border-radius: 10px;
  border: 1px solid #e5e7eb;
  background-color: var(--white);
  color: var(--font-color-dark);
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
  display: flex;
  align-items: center;
  justify-content: space-between;

  &:after {
    content: '→';
    opacity: 0;
    transform: translateX(-5px);
    transition: all 0.2s;
  }

  &:hover {
    background-color: #f9fafb;
    border-color: var(--green);
    color: var(--green-hover);

    &:after {
      opacity: 1;
      transform: translateX(0);
    }
  }
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
  border-bottom: 1px solid #f3f4f6;

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
  font-weight: 600;
  font-size: 0.95rem;
  color: var(--font-color-dark);
}

.item-subtitle {
  font-size: 0.85rem;
  color: var(--font-color-medium);
}

.status-tag {
  font-size: 0.75rem;
  font-weight: 700;
  padding: 4px 8px;
  border-radius: 12px;
  text-transform: uppercase;
}

.status-tag.critical {
  background-color: #fef2f2;
  color: #ef4444;
}

.status-tag.urgent {
  background-color: #fff7ed;
  color: #f97316;
}

.status-tag.warning {
  background-color: #fffbeb;
  color: #d97706;
}

.status-tag.new {
  background-color: #eff6ff;
  color: #3b82f6;
}

.status-tag.review {
  background-color: #faf5ff;
  color: #8b5cf6;
}

.status-tag.pending {
  background-color: #fff7ed;
  color: #f97316;
}

.time-tag {
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--green);
  background-color: #f0fdf4;
  padding: 4px 8px;
  border-radius: 6px;
}
</style>
