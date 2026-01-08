<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { mockDailyLogs, type IDailyLog } from '../../stores/mockDailyCare'
import { InputField, BooleanCapsule } from '../../components/common/ui'

const router = useRouter()

// Get today's log or fallback
const todayLog = computed(() => mockDailyLogs.value[0]) // Simplification for mock

const shift = ref<'AM' | 'PM'>('AM')

// Computed Tasks based on Shift
const currentTasks = computed(() => {
  return shift.value === 'AM' ? todayLog.value.amTasks : todayLog.value.pmTasks
})

const currentVolunteerName = computed({
  get: () => (shift.value === 'AM' ? todayLog.value.amVolunteer : todayLog.value.pmVolunteer),
  set: (val) => {
    if (shift.value === 'AM') todayLog.value.amVolunteer = val
    else todayLog.value.pmVolunteer = val
  },
})

// Columns config
const columns = [
  { key: 'food', label: 'Food', icon: '🥣' },
  { key: 'pee', label: 'Pee', icon: '💧' },
  { key: 'poo', label: 'Poo', icon: '💩' },
  { key: 'clean', label: 'Clean', icon: '🧹' },
  { key: 'play', label: 'Play', icon: '🧸' },
]

const generalChecklist = [
  { key: 'sweep', label: 'Sweep Floor' },
  { key: 'litterScoop', label: 'Wash Litter Scoopers' },
  { key: 'bowls', label: 'Wash Dirty Bowls & Pails' },
  { key: 'water', label: 'Refill Water & Brita' },
  { key: 'trash', label: 'Empty Trash / Clean Bag' },
  { key: 'windows', label: 'Clean Windows' },
  { key: 'mop', label: 'Mop Floor' },
]

const goHome = () => router.push('/kiosk')
</script>

<template>
  <div class="care-sheet-page">
    <!-- HEADER -->
    <div class="sheet-header">
      <button class="back-link" @click="goHome">← Back</button>
      <div class="title-group">
        <h1>Daily Care Sheet</h1>
        <div class="date-badge">{{ todayLog.date }}</div>
      </div>

      <div class="shift-selector">
        <button class="shift-btn" :class="{ active: shift === 'AM' }" @click="shift = 'AM'">
          AM Shift
        </button>
        <button class="shift-btn" :class="{ active: shift === 'PM' }" @click="shift = 'PM'">
          PM Shift
        </button>
      </div>
    </div>

    <!-- VOLUNTEER INFO -->
    <div class="volunteer-row">
      <span>Volunteer on Duty:</span>
      <input
        v-model="currentVolunteerName"
        type="text"
        placeholder="Enter your name..."
        class="volunteer-input"
      />
    </div>

    <!-- PET GRID -->
    <div class="grid-container">
      <table class="care-table">
        <thead>
          <tr>
            <th class="narrow-col">Cage</th>
            <th class="pet-col">Cat / Pet</th>
            <th v-for="col in columns" :key="col.key" class="check-col">
              <div class="col-header">
                <span class="col-icon">{{ col.icon }}</span>
                <span class="col-lbl">{{ col.label }}</span>
              </div>
            </th>
            <th>Mood / Health / Notes</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="entry in todayLog.entries" :key="entry.id">
            <td class="narrow-col center-text">{{ entry.cageNumber }}</td>
            <td class="pet-col fw-bold">{{ entry.petName }}</td>

            <!-- DYNAMIC CHECKBOXES -->
            <td v-for="col in columns" :key="col.key" class="check-cell">
              <input
                type="checkbox"
                v-model="entry[shift.toLowerCase()][col.key]"
                class="big-checkbox"
              />
            </td>

            <!-- NOTES INPUT -->
            <td class="notes-cell">
              <input
                type="text"
                v-model="entry[shift.toLowerCase()].notes"
                class="table-input"
                placeholder="Enter notes..."
              />
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- BOTTOM SECTION: CHECKLIST & EXTRAS -->
    <div class="bottom-section">
      <!-- GENERAL TASKS -->
      <div class="card tasks-card">
        <h3>🧹 General Tasks</h3>
        <div class="checklist-grid">
          <label
            v-for="task in generalChecklist"
            :key="task.key"
            class="task-row"
            :class="{ checked: currentTasks[task.key] }"
          >
            <input type="checkbox" v-model="currentTasks[task.key]" class="big-checkbox" />
            <span>{{ task.label }}</span>
          </label>
        </div>
      </div>

      <!-- SUMMARY INPUTS -->
      <div class="card summary-card">
        <h3>📝 Shift Summary</h3>

        <div class="input-group">
          <label>Donations Collected ($)</label>
          <input
            type="text"
            v-model="currentTasks.donationsCollected"
            placeholder="$0.00"
            class="summary-input"
          />
        </div>

        <div class="input-group">
          <label>Applications Mentioned</label>
          <input
            type="text"
            v-model="currentTasks.applicationsInfo"
            placeholder="Which cats?"
            class="summary-input"
          />
        </div>

        <div class="input-group">
          <label>Additional Notes</label>
          <textarea
            v-model="currentTasks.additionalNotes"
            rows="3"
            class="summary-textarea"
          ></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.care-sheet-page {
  max-width: 1200px;
  margin: 0 auto;
  padding-bottom: 50px;
}

.sheet-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.back-link {
  background: none;
  border: none;
  font-size: 1.1rem;
  color: var(--font-color-medium);
  font-weight: 500;
  cursor: pointer;
}

.title-group {
  text-align: center;
  h1 {
    margin: 0;
    font-size: 2rem;
  }
  .date-badge {
    color: var(--font-color-medium);
    font-size: 1.1rem;
  }
}

.shift-selector {
  background: #e2e8f0;
  padding: 4px;
  border-radius: 12px;
  display: flex;
  gap: 4px;
}

.shift-btn {
  border: none;
  background: transparent;
  padding: 10px 24px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  color: #64748b;
  cursor: pointer;

  &.active {
    background: white;
    color: var(--purple);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  }
}

.volunteer-row {
  background: white;
  padding: 16px 24px;
  border-radius: 12px;
  border: 1px solid #cbd5e1;
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
  font-size: 1.2rem;
  font-weight: 600;
}

.volunteer-input {
  flex: 1;
  font-size: 1.2rem;
  padding: 8px;
  border: none;
  border-bottom: 2px solid #e2e8f0;
  outline: none;
  &:focus {
    border-color: var(--purple);
  }
}

/* TABLE */
.grid-container {
  background: white;
  border-radius: 16px;
  border: 2px solid #000; /* Mimic bold lines of paper form */
  overflow: hidden;
  margin-bottom: 32px;
}

.care-table {
  width: 100%;
  border-collapse: collapse;

  th,
  td {
    border: 1px solid #cbd5e1;
    padding: 12px;
    vertical-align: middle;
  }

  th {
    background: #f1f5f9;
    text-align: center;
    font-size: 0.85rem;
    text-transform: uppercase;
    color: #475569;
  }
}

.narrow-col {
  width: 60px;
  text-align: center;
  font-weight: 700;
}
.pet-col {
  width: 180px;
  font-size: 1.1rem;
}
.center-text {
  text-align: center;
}
.fw-bold {
  font-weight: 700;
}

.col-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}
.col-icon {
  font-size: 1.4rem;
}
.col-lbl {
  font-size: 0.7rem;
}

.check-cell {
  text-align: center;
  background: #fff;
}

.big-checkbox {
  width: 28px;
  height: 28px;
  cursor: pointer;
  accent-color: var(--purple);
}

.table-input {
  width: 100%;
  padding: 8px;
  font-size: 1rem;
  border: none;
  background: transparent;
  outline: none;
  font-family: 'Inter', sans-serif; /* For handwriting feel? */
}

/* BOTTOM SECTION */
.bottom-section {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.card {
  background: white;
  padding: 24px;
  border-radius: 16px;
  border: 1px solid #cbd5e1;
  h3 {
    margin-top: 0;
    margin-bottom: 20px;
  }
}

.checklist-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 12px;
}

.task-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border-radius: 8px;
  cursor: pointer;
  border: 1px solid transparent;

  &.checked {
    background: #f0fdf4;
    border-color: #bbf7d0;
  }

  span {
    font-size: 1.1rem;
  }
}

.input-group {
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  label {
    font-weight: 600;
    font-size: 0.95rem;
    color: #475569;
  }
}

.summary-input {
  padding: 12px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 1.1rem;
}

.summary-textarea {
  padding: 12px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 1rem;
  resize: vertical;
}
</style>
