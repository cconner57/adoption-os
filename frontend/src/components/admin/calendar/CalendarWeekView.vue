<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  activeDate: Date
  recurringShifts: Record<number, Array<{ time: string; title: string; type: string }>>
}>()

// Helper to parse time for sorting (e.g. "9:00 AM")
const parseTime = (timeStr: string): number => {
  if (!timeStr) return 0
  const [time, modifier] = timeStr.split(' ')
  let [hours, minutes] = time.split(':').map(Number)
  if (modifier === 'PM' && hours < 12) hours += 12
  if (modifier === 'AM' && hours === 12) hours = 0
  return hours * 60 + minutes
}

// Week View Data
const weekDays = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

  // Find Monday of the ACTIVE week
  const current = new Date(props.activeDate)
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
    const shifts = props.recurringShifts[dayOfWeek] || []
    shifts.forEach((shift, idx) => {
      events.push({ id: `wk-v-${index}-${idx}`, ...shift })
    })

    if (index === 0) {
      // Monday
      events.push({
        id: `wk-vet-special`,
        type: 'vet',
        time: '9:00 AM',
        title: 'Vet Appointments', // Fixed typo here if it was wrong before
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
</script>

<template>
  <div class="week-view">
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
</template>

<style scoped>
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
  background: var(--text-inverse);
  border-radius: 12px;
  padding: 16px;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  gap: 16px;

  &.today {
    border: 2px solid hsl(from var(--color-secondary) h s 80%);
  }
}

.column-header {
  text-align: center;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 12px;

  .day-name {
    display: block;
    font-size: 0.9rem;
    color: hsl(from var(--color-neutral) h s 50%);
    text-transform: uppercase;
    font-weight: 600;
  }

  .day-numero {
    display: block;
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--text-primary);
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
    background: hsl(from var(--color-secondary) h s 95%);
    border-left-color: var(--color-secondary);
    color: var(--color-secondary);
  }

  &.vet {
    background: hsl(from var(--color-primary) h s 95%);
    border-left-color: var(--color-primary);
    color: var(--color-primary);
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
