<script setup lang="ts">
import { ref, computed } from 'vue'

type ViewMode = 'month' | 'week'
const viewMode = ref<ViewMode>('month')

const today = new Date()
const currentMonth = ref(today.getMonth())
const currentYear = ref(today.getFullYear())

// Mock Calendar Data Logic
const daysInMonth = computed(() => {
  return new Date(currentYear.value, currentMonth.value + 1, 0).getDate()
})

const firstDayOfMonth = computed(() => {
  return new Date(currentYear.value, currentMonth.value, 1).getDay()
})

// Month View Data
const monthDays = computed(() => {
  const days = []

  // Padding for start of month
  // Monday is 1, Sunday is 0 in JS. We want Monday to be first column if we follow week view?
  // Let's stick to standard Sun-Sat for Month view for now as it's easier, or alignment with week view (Mon-Sun)
  // Let's do Mon-Sun alignment to match our Weekly view

  let padding = firstDayOfMonth.value - 1
  if (padding < 0) padding = 6 // Sunday is 0, so if Mon(1) is start, Sun is 6

  for (let i = 0; i < padding; i++) {
    days.push({ id: `pad-${i}`, isEmpty: true })
  }

  for (let i = 1; i <= daysInMonth.value; i++) {
    // Generate mock events
    const events = []

    // Pattern similar to dashboard
    const dayOfWeek = (padding + i - 1) % 7 // 0=Mon, ... 6=Sun

    if (i === 12 || i === 25) {
      // Specific dates
      events.push({ id: `evt-${i}-1`, type: 'volunteer', time: '10:00 AM', title: 'Orientation' })
    }

    if (dayOfWeek % 2 === 0) {
      // Mon, Wed, Fri
      events.push({ id: `v-${i}`, type: 'volunteer', time: '9:00 AM', title: 'Morning Shift' })
    }

    if (dayOfWeek === 1) {
      // Tue
      events.push({ id: `vet-${i}`, type: 'vet', time: '2:00 PM', title: 'Vet Visits' })
    }

    days.push({
      id: `day-${i}`,
      date: i,
      events,
      isToday: i === today.getDate() && currentMonth.value === today.getMonth(),
    })
  }

  return days
})

// Week View Data (Reused logic effectively)
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

  // Find Monday of current week (relative to today representing the "view" week)
  // For simplicity this mock always shows "Current Week"
  const now = new Date()
  const monday = new Date(now)
  monday.setDate(now.getDate() - now.getDay() + 1)

  return days.map((name, index) => {
    const date = new Date(monday)
    date.setDate(monday.getDate() + index)

    // Mock events
    const events = []
    if (index % 2 === 0)
      events.push({
        id: `wk-v-${index}`,
        type: 'volunteer',
        time: '9:00 AM',
        title: 'Shift: Group A',
      })
    if (index === 1)
      events.push({ id: `wk-vet-${index}`, type: 'vet', time: '2:00 PM', title: 'Dr. Smith Visit' })

    return {
      name,
      date: date.getDate(),
      fullDate: date,
      isToday: date.toDateString() === now.toDateString(),
      events,
    }
  })
})

const weekDaysHeader = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

const monthName = computed(() => {
  return new Date(currentYear.value, currentMonth.value).toLocaleString('default', {
    month: 'long',
    year: 'numeric',
  })
})
</script>

<template>
  <div class="calendar-page">
    <div class="page-header">
      <h1>Calendar</h1>

      <div class="controls">
        <h2 class="current-period">{{ monthName }}</h2>

        <div class="view-toggle">
          <button
            class="toggle-btn"
            :class="{ active: viewMode === 'week' }"
            @click="viewMode = 'week'"
          >
            Week
          </button>
          <button
            class="toggle-btn"
            :class="{ active: viewMode === 'month' }"
            @click="viewMode = 'month'"
          >
            Month
          </button>
        </div>
      </div>
    </div>

    <!-- Legend -->
    <div class="legend-bar">
      <span class="legend-item"><span class="dot volunteer"></span> Volunteer Shift</span>
      <span class="legend-item"><span class="dot vet"></span> Vet Appointment</span>
    </div>

    <!-- Month View -->
    <div v-if="viewMode === 'month'" class="month-view">
      <div class="week-header">
        <div v-for="day in weekDaysHeader" :key="day" class="header-cell">{{ day }}</div>
      </div>
      <div class="month-grid">
        <div
          v-for="day in monthDays"
          :key="day.id"
          class="day-cell"
          :class="{ empty: day.isEmpty, today: day.isToday }"
        >
          <div v-if="!day.isEmpty" class="day-content">
            <span class="date-number">{{ day.date }}</span>
            <div class="events-list">
              <div
                v-for="event in day.events"
                :key="event.id"
                class="event-dot"
                :class="event.type"
                :title="`${event.time} - ${event.title}`"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Week View -->
    <div v-else class="week-view">
      <div class="week-grid">
        <div
          v-for="day in weekDays"
          :key="day.name"
          class="week-day-column"
          :class="{ today: day.isToday }"
        >
          <div class="column-header">
            <span class="day-name">{{ day.name }}</span>
            <span class="day-numero">{{ day.date }}</span>
          </div>
          <div class="column-content">
            <div v-for="event in day.events" :key="event.id" class="event-card" :class="event.type">
              <span class="time">{{ event.time }}</span>
              <span class="title">{{ event.title }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.calendar-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: 100%;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  h1 {
    font-size: 1.8rem;
    color: var(--font-color-dark);
    margin: 0;
  }
}

.controls {
  display: flex;
  align-items: center;
  gap: 24px;
}

.current-period {
  font-size: 1.2rem;
  font-weight: 600;
  color: var(--font-color-medium);
  margin: 0;
}

.view-toggle {
  background: #e5e7eb;
  padding: 4px;
  border-radius: 8px;
  display: flex;
  gap: 4px;
}

.toggle-btn {
  background: none;
  border: none;
  padding: 6px 16px;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--font-color-medium);
  cursor: pointer;

  &.active {
    background: white;
    color: var(--font-color-dark);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  }
}

.legend-bar {
  display: flex;
  gap: 16px;
  margin-bottom: 8px;
}
.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
  color: var(--font-color-medium);
}
.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}
.dot.volunteer {
  background-color: var(--purple);
}
.dot.vet {
  background-color: var(--green);
}

/* Month View */
.month-view {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  padding: 24px;
  display: flex;
  flex-direction: column;
  flex: 1;
}

.week-header {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  text-align: center;
  margin-bottom: 12px;
  font-weight: 600;
  color: var(--font-color-medium);
}

.month-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  grid-auto-rows: 1fr;
  gap: 1px;
  background-color: #f3f4f6; /* Grid lines */
  border: 1px solid #f3f4f6;
  flex: 1;
}

.day-cell {
  background-color: white;
  min-height: 100px;
  padding: 8px;

  &.empty {
    background-color: #fcfcfc;
  }

  &.today {
    background-color: #f0f9ff;
    .date-number {
      color: var(--blue);
      font-weight: 700;
    }
  }
}

.date-number {
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--font-color-dark);
  margin-bottom: 8px;
  display: block;
}

.events-list {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.event-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;

  &.volunteer {
    background-color: var(--purple);
  }
  &.vet {
    background-color: var(--green);
  }
}

/* Week View */
.week-view {
  flex: 1;
  overflow-x: auto;
}

.week-grid {
  display: grid;
  grid-template-columns: repeat(7, minmax(140px, 1fr));
  gap: 16px;
  height: 100%;
}

.week-day-column {
  background: white;
  border-radius: 12px;
  padding: 16px;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  gap: 16px;

  &.today {
    border: 2px solid var(--blue-weak);
  }
}

.column-header {
  text-align: center;
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 12px;

  .day-name {
    display: block;
    font-size: 0.9rem;
    color: var(--font-color-medium);
    text-transform: uppercase;
    font-weight: 600;
  }

  .day-numero {
    display: block;
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--font-color-dark);
  }
}

.column-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.event-card {
  padding: 12px;
  border-radius: 8px;
  border-left: 4px solid transparent;

  &.volunteer {
    background: #faf5ff;
    border-left-color: var(--purple);
    color: var(--purple-hover);
  }

  &.vet {
    background: #f0fdf4;
    border-left-color: var(--green);
    color: var(--green-hover);
  }

  .time {
    display: block;
    font-size: 0.75rem;
    font-weight: 700;
    margin-bottom: 4px;
    opacity: 0.8;
  }

  .title {
    display: block;
    font-size: 0.9rem;
    font-weight: 500;
    line-height: 1.3;
  }
}
</style>
