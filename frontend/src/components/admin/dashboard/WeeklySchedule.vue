<script setup lang="ts">
import { computed } from 'vue'

import { useActiveVolunteersStore } from '../../../stores/activeVolunteers'

const volunteerStore = useActiveVolunteersStore()

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
</script>

<template>
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
</template>

<style scoped>
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

.widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 12px;
}

.weekly-schedule {
  min-height: 340px;
  grid-column: span 3;
}

@media (max-width: 900px) {
  .weekly-schedule {
    grid-column: span 1;
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
}

.calendar-day.today {
  background-color: var(--text-inverse);
  border-color: hsl(from var(--color-secondary) h s 70%);
  box-shadow: 0 0 0 2px hsla(from var(--color-secondary) h s l / 0.1);
}

.calendar-day.today .day-header {
  color: var(--color-secondary);
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
}

.day-header .day-date {
  font-size: 1rem;
  color: var(--text-primary);
  margin-top: 2px;
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
}

.event-pill.feeding {
  background-color: hsl(from var(--color-secondary) h s 96%);
  color: var(--color-secondary);
  border-left-color: var(--color-secondary);
}

.event-pill.adoption {
  background-color: hsl(from var(--color-tertiary) h s 96%);
  color: var(--color-tertiary);
  border-left-color: var(--color-tertiary);
}

.event-pill.transport {
  background-color: hsl(from var(--color-warning) h s 96%);
  color: hsl(from var(--color-warning) h s 40%);
  border-left-color: var(--color-warning);
}

.event-pill.vet {
  background-color: hsl(from var(--color-danger) h s 96%);
  color: var(--color-danger);
  border-left-color: var(--color-danger);
}

.event-pill.default {
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
</style>
