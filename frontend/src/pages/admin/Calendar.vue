<script setup lang="ts">
import { ref, computed, watch } from 'vue'

type ViewMode = 'month' | 'week'
const savedView = localStorage.getItem('adminCalendarView') as ViewMode | null
const viewMode = ref<ViewMode>(savedView === 'week' ? 'week' : 'month')

watch(viewMode, (newMode) => {
  localStorage.setItem('adminCalendarView', newMode)
})

const today = new Date()
const activeDate = ref(new Date())

// Navigation Logic
const prev = () => {
  const d = new Date(activeDate.value)
  if (viewMode.value === 'month') {
    d.setMonth(d.getMonth() - 1)
  } else {
    d.setDate(d.getDate() - 7)
  }
  activeDate.value = d
}

const next = () => {
  const d = new Date(activeDate.value)
  if (viewMode.value === 'month') {
    d.setMonth(d.getMonth() + 1)
  } else {
    d.setDate(d.getDate() + 7)
  }
  activeDate.value = d
}

const goToday = () => {
  activeDate.value = new Date()
}

const currentMonth = computed(() => activeDate.value.getMonth())
const currentYear = computed(() => activeDate.value.getFullYear())

// Recurring Volunteer Schedule
const recurringShifts: Record<number, Array<{ time: string; title: string; type: string }>> = {
  0: [
    // Sunday
    { time: '10:00 AM', title: 'Allison (10AM-12PM)', type: 'volunteer' },
    { time: '4:00 PM', title: 'Brandon (4PM-6PM)', type: 'volunteer' },
  ],
  1: [
    // Monday
    { time: '10:00 AM', title: 'Leanne (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Arianna (6PM-8PM)', type: 'volunteer' },
  ],
  2: [
    // Tuesday
    { time: '10:00 AM', title: 'Sonia (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Lynn (6PM-8PM)', type: 'volunteer' },
  ],
  3: [
    // Wednesday
    { time: '10:00 AM', title: 'Bella (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katelyn (6PM-8PM)', type: 'volunteer' },
  ],
  4: [
    // Thursday
    { time: '10:00 AM', title: 'Alejandra (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Nathan (6PM-8PM)', type: 'volunteer' },
  ],
  5: [
    // Friday
    { time: '10:00 AM', title: 'Linda (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katie (6PM-8PM)', type: 'volunteer' },
  ],
  6: [
    // Saturday
    { time: '10:00 AM', title: 'Lindsey & Celina (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Chris (6PM-8PM)', type: 'volunteer' },
  ],
}

// Helper to parse time for sorting (e.g. "9:00 AM")
const parseTime = (timeStr: string): number => {
  if (!timeStr) return 0
  const [time, modifier] = timeStr.split(' ')
  let [hours, minutes] = time.split(':').map(Number)
  if (modifier === 'PM' && hours < 12) hours += 12
  if (modifier === 'AM' && hours === 12) hours = 0
  return hours * 60 + minutes
}

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
  let padding = firstDayOfMonth.value - 1
  if (padding < 0) padding = 6

  for (let i = 0; i < padding; i++) {
    days.push({ id: `pad-${i}`, isEmpty: true })
  }

  for (let i = 1; i <= daysInMonth.value; i++) {
    // Generate mock events
    const events = []
    const currentDayDate = new Date(currentYear.value, currentMonth.value, i)
    const dayOfWeek = currentDayDate.getDay()

    // Add recurring shifts
    const shifts = recurringShifts[dayOfWeek] || []
    shifts.forEach((shift, idx) => {
      events.push({ id: `v-${i}-${idx}`, ...shift })
    })

    if (i === 5) {
      // Vet mock logic
      events.push({
        id: `vet-${i}-special`,
        type: 'vet',
        time: '9:00 AM',
        title: 'Vet Appointments',
      })
    }

    // Sort events by time
    events.sort((a, b) => parseTime(a.time) - parseTime(b.time))

    days.push({
      id: `day-${i}`,
      date: i,
      events,
      isToday:
        i === today.getDate() &&
        currentMonth.value === today.getMonth() &&
        currentYear.value === today.getFullYear(),
    })
  }

  return days
})

// Week View Data
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

  // Find Monday of the ACTIVE week
  const current = new Date(activeDate.value)
  const day = current.getDay()
  const diff = current.getDate() - day + (day === 0 ? -6 : 1) // adjust when day is sunday
  const monday = new Date(current.setDate(diff))

  return days.map((name, index) => {
    const date = new Date(monday)
    date.setDate(monday.getDate() + index)
    const dayOfWeek = date.getDay()

    // Mock events
    const events = []

    // Add recurring shifts
    const shifts = recurringShifts[dayOfWeek] || []
    shifts.forEach((shift, idx) => {
      events.push({ id: `wk-v-${index}-${idx}`, ...shift })
    })

    if (index === 0) {
      // Monday
      events.push({
        id: `wk-vet-special`,
        type: 'vet',
        time: '9:00 AM',
        title: 'Vet Appointments',
        details: ['Malachi', 'Merry', 'Ariel', 'Aragorn', 'Purina'],
      })
    }

    // Sort events by time
    events.sort((a, b) => parseTime(a.time) - parseTime(b.time))

    return {
      name,
      date: date.getDate(),
      fullDate: date,
      isToday: date.toDateString() === new Date().toDateString(),
      events,
    }
  })
})

const weekDaysHeader = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

const monthName = computed(() => {
  if (viewMode.value === 'month') {
    return activeDate.value.toLocaleString('default', {
      month: 'long',
      year: 'numeric',
    })
  } else {
    // For week view, show "Jan 2026" or "Jan - Feb 2026"
    const start = new Date(activeDate.value)
    const day = start.getDay()
    const diff = start.getDate() - day + (day === 0 ? -6 : 1)
    const monday = new Date(start.setDate(diff))
    const sunday = new Date(monday)
    sunday.setDate(monday.getDate() + 6)

    if (monday.getMonth() === sunday.getMonth()) {
      return monday.toLocaleString('default', { month: 'long', year: 'numeric' })
    } else {
      return `${monday.toLocaleString('default', { month: 'long' })} - ${sunday.toLocaleString('default', { month: 'long', year: 'numeric' })}`
    }
  }
})
</script>

<template>
  <div class="calendar-page">
    <div class="page-header">
      <h1>Calendar</h1>

      <div class="controls">
        <div class="nav-buttons">
          <button class="nav-btn" @click="prev">&lt;</button>
          <button class="nav-btn today" @click="goToday">Today</button>
          <button class="nav-btn" @click="next">&gt;</button>
        </div>

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

        <button class="add-event-btn">Add Event +</button>
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
                class="compact-event"
                :class="event.type"
                :title="`${event.time} - ${event.title}`"
              >
                <div class="dot" :class="event.type"></div>
                <div class="content">
                  <span class="time">{{ event.time }}</span>
                  <span class="title">{{ event.title }}</span>
                </div>
              </div>
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
              <ul v-if="event.details" class="event-details-list">
                <li v-for="detail in event.details" :key="detail">{{ detail }}</li>
              </ul>
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

.nav-buttons {
  display: flex;
  gap: 8px;
}

.nav-btn {
  background: white;
  border: 1px solid #e5e7eb;
  padding: 6px 12px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  color: var(--font-color-medium);
  transition: all 0.2s;

  &:hover {
    background: #f9fafb;
    border-color: #d1d5db;
    color: var(--font-color-dark);
  }

  &.today {
    font-size: 0.9rem;
  }
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

.add-event-btn {
  background-color: var(--blue);
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.2s;
  display: flex;
  align-items: center;
  gap: 6px;

  &:hover {
    background-color: var(--blue-hover);
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
  flex-direction: column;
  gap: 4px;
}

.compact-event {
  display: flex;
  align-items: flex-start;
  gap: 6px;
  padding: 4px 6px;
  border-radius: 4px;
  font-size: 0.75rem;
  background-color: #f3f4f6;
  cursor: pointer;
  white-space: normal;
  line-height: 1.3;

  .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    flex-shrink: 0;
    margin-top: 5px; /* Align with first line of text */
  }

  .content {
    display: block;
  }

  .time {
    font-weight: 700;
    opacity: 0.8;
    margin-right: 4px;
    display: inline;
  }

  .title {
    display: inline;
  }

  &.volunteer {
    background-color: #faf5ff;
    color: var(--purple); /* Fallback or use var(--purple-hover) if defined */
    .dot {
      background-color: var(--purple);
    }
  }

  &.vet {
    background-color: #f0fdf4;
    color: var(--green); /* Fallback or use var(--green-hover) if defined */
    .dot {
      background-color: var(--green);
    }
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

  .event-details-list {
    margin: 4px 0 0 0;
    padding-left: 16px;
    font-size: 0.8rem;
    opacity: 0.9;
    list-style-type: disc; /* Ensure bullets show */

    li {
      margin-bottom: 2px;
      margin-left: 4px; /* Slight indent for bullets */
    }
  }
}
</style>
