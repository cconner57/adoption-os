<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import CalendarMonthView from '../../components/admin/calendar/CalendarMonthView.vue'
import CalendarWeekView from '../../components/admin/calendar/CalendarWeekView.vue'

type ViewMode = 'month' | 'week'
const savedView = localStorage.getItem('adminCalendarView') as ViewMode | null
const viewMode = ref<ViewMode>(savedView === 'week' ? 'week' : 'month')

watch(viewMode, (newMode) => {
  localStorage.setItem('adminCalendarView', newMode)
})

const today = new Date()
const activeDate = ref(new Date())

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

const recurringShifts: Record<number, Array<{ time: string; title: string; type: string }>> = {
  0: [
    
    { time: '10:00 AM', title: 'Allison (10AM-12PM)', type: 'volunteer' },
    { time: '4:00 PM', title: 'Brandon (4PM-6PM)', type: 'volunteer' },
  ],
  1: [
    
    { time: '10:00 AM', title: 'Leanne (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Arianna (6PM-8PM)', type: 'volunteer' },
  ],
  2: [
    
    { time: '10:00 AM', title: 'Sonia (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Lynn (6PM-8PM)', type: 'volunteer' },
  ],
  3: [
    
    { time: '10:00 AM', title: 'Bella (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katelyn (6PM-8PM)', type: 'volunteer' },
  ],
  4: [
    
    { time: '10:00 AM', title: 'Alejandra (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Nathan (6PM-8PM)', type: 'volunteer' },
  ],
  5: [
    
    { time: '10:00 AM', title: 'Linda (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Katie (6PM-8PM)', type: 'volunteer' },
  ],
  6: [
    
    { time: '10:00 AM', title: 'Lindsey & Celina (10AM-12PM)', type: 'volunteer' },
    { time: '6:00 PM', title: 'Chris (6PM-8PM)', type: 'volunteer' },
  ],
}

const monthName = computed(() => {
  if (viewMode.value === 'month') {
    return activeDate.value.toLocaleString('default', {
      month: 'long',
      year: 'numeric',
    })
  } else {
    
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

    <div class="legend-bar">
      <span class="legend-item"><span class="dot volunteer"></span> Volunteer Shift</span>
      <span class="legend-item"><span class="dot vet"></span> Vet Appointment</span>
    </div>

    <CalendarMonthView
      v-if="viewMode === 'month'"
      :current-year="currentYear"
      :current-month="currentMonth"
      :active-date="activeDate"
      :today="today"
      :recurring-shifts="recurringShifts"
    />

    <CalendarWeekView v-else :active-date="activeDate" :recurring-shifts="recurringShifts" />
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
    color: var(--text-primary);
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
  background: #fff;
  border: 1px solid #e5e7eb;
  padding: 6px 12px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  color: hsl(from var(--color-neutral) h s 50%);
  transition: all 0.2s;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
    border-color: var(--border-color);
    color: var(--text-primary);
  }

  &.today {
    font-size: 0.9rem;
  }
}

.current-period {
  font-size: 1.2rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  margin: 0;
  min-width: 200px;
  text-align: center;
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
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;

  &.active {
    background: var(--text-inverse);
    color: var(--text-primary);
    box-shadow: 0 2px 4px rgb(0 0 0 / 5%);
  }
}

.add-event-btn {
  background-color: var(--color-secondary);
  color: #fff;
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
    background-color: hsl(from var(--color-secondary) h s 40%);
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
  color: hsl(from var(--color-neutral) h s 50%);
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.dot.volunteer {
  background-color: var(--color-secondary);
}

.dot.vet {
  background-color: var(--color-primary);
}
</style>
