<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuthStore } from '../../stores/auth'
import { mockPetsData } from '../../stores/mockPetData'

const authStore = useAuthStore()
const userName = computed(() => authStore.user?.name || 'Admin')

const adoptablePetsCount = computed(() => {
  return mockPetsData.filter((pet) => pet.details.status === 'available').length
})

// Mock Calendar Data
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const today = new Date()

  // Find Monday of current week
  const monday = new Date(today)
  monday.setDate(today.getDate() - today.getDay() + 1)

  return days.map((name, index) => {
    const date = new Date(monday)
    date.setDate(monday.getDate() + index)

    // Mock events based on day index
    const events = []

    // Add mock volunteer shifts
    if (index % 2 === 0) {
      // Mon, Wed, Fri
      events.push({
        id: `v-${index}`,
        type: 'volunteer',
        time: '9:00 AM',
        title: 'Start Shift: Sarah & Mike',
      })
    }
    if (index === 5 || index === 6) {
      // Weekends
      events.push({
        id: `v-${index}-2`,
        type: 'volunteer',
        time: '10:00 AM',
        title: 'Adoption Event Team',
      })
    }

    // Add mock vet appointments
    if (index === 1) {
      // Tuesday
      events.push({
        id: `vet-${index}`,
        type: 'vet',
        time: '2:00 PM',
        title: 'Amari - Vaccination',
      })
    }
    if (index === 3) {
      // Thursday
      events.push({ id: `vet-${index}`, type: 'vet', time: '11:30 AM', title: 'Review: Apollo' })
    }

    return {
      name,
      date: date.getDate(),
      isToday: date.toDateString() === today.toDateString(),
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
        </div>
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

/* Calendar Styles */
.weekly-schedule {
  grid-column: span 1;
  display: flex;
  flex-direction: column;
}

.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 12px;

  h3 {
    margin: 0;
    border: none;
    padding: 0;
  }
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
}

.calendar-day {
  background-color: #fcfcfc;
  border-radius: 8px;
  padding: 8px;
  min-height: 140px;
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
  font-size: 0.9rem;
  font-weight: 600;
  padding-bottom: 6px;
  border-bottom: 1px solid #f3f4f6;

  .day-date {
    font-size: 1.1rem;
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
  padding: 6px;
  border-radius: 4px;
  font-size: 0.75rem;
  display: flex;
  flex-direction: column;
  gap: 2px;
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
    font-weight: 600;
    opacity: 0.8;
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

.dashboard-widgets {
  display: grid;
  grid-template-columns: 2fr 1fr;
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
  min-height: 300px;
}

.widget h3 {
  font-size: 1.25rem;
  margin-bottom: 20px;
  color: var(--font-color-dark);
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 12px;
}

.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
  color: var(--font-color-medium);
  font-style: italic;
  background-color: #f9fafb;
  border-radius: 8px;
  border: 1px dashed #e5e7eb;
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-btn {
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  background-color: var(--white);
  color: var(--font-color-dark);
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;

  &:hover {
    background-color: var(--green-weak);
    color: var(--green);
    border-color: var(--green-weak);
  }
}
</style>
