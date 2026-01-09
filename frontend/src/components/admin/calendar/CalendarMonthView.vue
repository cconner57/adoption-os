<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  currentYear: number
  currentMonth: number
  activeDate: Date
  today: Date
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

const daysInMonth = computed(() => {
  return new Date(props.currentYear, props.currentMonth + 1, 0).getDate()
})

const firstDayOfMonth = computed(() => {
  return new Date(props.currentYear, props.currentMonth, 1).getDay()
})

const weekDaysHeader = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

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
    const currentDayDate = new Date(props.currentYear, props.currentMonth, i)
    const dayOfWeek = currentDayDate.getDay()

    // Add recurring shifts
    const shifts = props.recurringShifts[dayOfWeek] || []
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
        i === props.today.getDate() &&
        props.currentMonth === props.today.getMonth() &&
        props.currentYear === props.today.getFullYear(),
    })
  }

  return days
})
</script>

<template>
  <div class="month-view">
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
</template>

<style scoped>
/* Month View */
.month-view {
  background: var(--text-inverse);
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
  color: hsl(from var(--color-neutral) h s 50%);
}

.month-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  grid-auto-rows: 1fr;
  gap: 1px;
  background-color: var(--border-color); /* Grid lines */
  border: 1px solid var(--border-color);
  flex: 1;
}

.day-cell {
  background-color: var(--text-inverse);
  min-height: 100px;
  padding: 8px;

  &.empty {
    background-color: hsl(from var(--color-neutral) h s 98%);
  }

  &.today {
    background-color: hsl(from var(--color-secondary) h s 95%);
    .date-number {
      color: var(--color-secondary);
      font-weight: 700;
    }
  }
}

.date-number {
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--text-primary);
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
  background-color: hsl(from var(--color-neutral) h s 95%);
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

  /* Specific event type styling needs to be here or passed down/global */
  &.volunteer {
    background-color: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
    .dot {
      background-color: var(--color-secondary);
    }
  }

  &.vet {
    background-color: hsl(from var(--color-primary) h s 95%);
    color: var(--color-primary);
    .dot {
      background-color: var(--color-primary);
    }
  }
}
</style>
