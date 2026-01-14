<script setup lang="ts">
import { ref, computed } from 'vue'
import type { IVolunteer, IShift, IIncident } from '../../../stores/mockVolunteerData'
import { availableBadges } from '../../../stores/mockVolunteerData'
import { Tag, Toast, Button } from '../../common/ui'

import VolunteerEditor from './VolunteerEditor.vue'
import ShiftForm from './ShiftForm.vue'
import { useAuthStore } from '../../../stores/auth'
import { formatPhoneNumber } from '../../../utils/formatters'
import { formatDate } from '../../../utils/date'

const props = defineProps<{
  volunteer: IVolunteer
  shifts: IShift[]
  incidents: IIncident[]
}>()

const authStore = useAuthStore()
const canEdit = computed(() => {
  if (!authStore.user?.Role) return true
  return ['admin', 'super_admin'].includes(authStore.user.Role.toLowerCase())
})

const isEditorOpen = ref(false)
const isAddingShift = ref(false)
const editingShiftData = ref<any>(null)
const showToast = ref(false)

const emit = defineEmits(['add-shift', 'update', 'update-shift', 'delete-shift'])

function openEditor() {
  isEditorOpen.value = true
}

function toggleAddShift() {
  if (isAddingShift.value) {
    isAddingShift.value = false
    editingShiftData.value = null
  } else {
    editingShiftData.value = null
    isAddingShift.value = true
  }
}

function editShift(shift: any) {
  editingShiftData.value = shift
  isAddingShift.value = true
}

function deleteShift(shift: any) {
  emit('delete-shift', shift.id)
}

function handleShiftSave(shiftData: any) {
  if (shiftData.id) {
    emit('update-shift', shiftData)
  } else {
    emit('add-shift', shiftData)
  }
  isAddingShift.value = false
  editingShiftData.value = null
  showToast.value = true
}

function handleSave(updatedData: Partial<IVolunteer>) {
  emit('update', updatedData)
  isEditorOpen.value = false
}

function verifyShift(shift: IShift, status: string) {
  const updatedShift = { ...shift, status }
  emit('update-shift', updatedShift)
}

function markLate(shift: IShift) {
  // Pre-fill with late status so user sees it, then open form to edit time
  editingShiftData.value = { ...shift, status: 'late' }
  isAddingShift.value = true
}

function handleArchive() {
  if (confirm('Are you sure you want to archive this volunteer?')) {
    // In real app, call API
    console.log('Archiving volunteer', props.volunteer.id)
    isEditorOpen.value = false
    // Show a different toast or redirect
  }
}

const activeTab = ref<'overview' | 'schedule' | 'incidents' | 'performance' | 'suggestions'>(
  'overview',
)
const tabs = ['overview', 'schedule', 'incidents', 'performance', 'suggestions'] as const

const showAllUpcoming = ref(false)
const upcomingShifts = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  const all = props.shifts
    .filter((s) => s.date >= today)
    .sort((a, b) => a.date.localeCompare(b.date))
  return showAllUpcoming.value ? all : all.slice(0, 2)
})

const unverifiedShifts = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return props.shifts
    .filter((s) => s.date < today && s.status === 'scheduled')
    .sort((a, b) => b.date.localeCompare(a.date)) // Newest to oldest
})

const verifiedPastShifts = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return props.shifts
    .filter((s) => s.date < today && s.status !== 'scheduled')
    .sort((a, b) => b.date.localeCompare(a.date))
})

const yearlyStats = computed(() => {
  const stats: Record<
    string,
    { year: string; hours: number; totalShifts: number; completedShifts: number; missed: number }
  > = {}

  // Initialize with join year if no shifts? No, just use shift data for accuracy.
  props.shifts.forEach((shift) => {
    const year = shift.date.split('-')[0]
    if (!stats[year]) {
      stats[year] = { year, hours: 0, totalShifts: 0, completedShifts: 0, missed: 0 }
    }

    // Calculate hours
    const start = parseInt(shift.startTime.split(':')[0])
    const end = parseInt(shift.endTime.split(':')[0])
    const duration = end - start

    // Logic: Only count hours if "all_good" or "completed" or "late"
    if (['completed', 'all_good', 'late'].includes(shift.status)) {
      stats[year].hours += duration
      stats[year].completedShifts += 1
    }

    stats[year].totalShifts += 1

    if (['no_show', 'missed'].includes(shift.status)) {
      stats[year].missed += 1
    }
  })

  // Calculate reliability % per year
  return Object.values(stats)
    .map((s) => {
      // Reliability is based on Total Scheduled/Attempted shifts vs Missed
      const reliability =
        s.totalShifts > 0 ? Math.round(((s.totalShifts - s.missed) / s.totalShifts) * 100) : 0

      return {
        year: s.year,
        hours: s.hours,
        shifts: s.completedShifts, // Display only completed shifts in the 'Total Shifts' column
        reliability,
      }
    })
    .sort((a, b) => b.year.localeCompare(a.year))
})

// Helper to determine display status
function getDisplayStatus(shift: IShift) {
  return shift.status.replace('_', ' ')
}

const suggestions = computed(() => {
  const list: { title: string; desc: string; icon: string }[] = []

  // Reliability Suggestions
  if (props.volunteer.reliabilityScore < 100) {
    list.push({
      title: 'Boost Reliability',
      desc: 'Ensure you check in for every shift on time. Consistent attendance over 3 months will restore your score.',
      icon: '📈',
    })
  }

  // Tier Advancement Suggestions
  if (props.volunteer.role === 'Teen') {
    list.push({
      title: 'Become Tier 1',
      desc: 'Complete 20 hours of service and the "Basic Handling" workshop to move up to Tier 1.',
      icon: '⭐',
    })
  } else if (props.volunteer.role === 'Tier 1') {
    list.push({
      title: 'Unlock Tier 2',
      desc: 'Accumulate 100 volunteer hours and earn the "Reliable" badge to qualify for Tier 2 training.',
      icon: '🔓',
    })
  }

  if (list.length === 0) {
    list.push({
      title: 'All Systems Go!',
      desc: 'You are doing great! Keep maintaining your excellent reliability and helping out.',
      icon: '🚀',
    })
  }

  return list
})

function getShiftStatusColor(shift: IShift) {
  const displayStatus = getDisplayStatus(shift).toLowerCase()
  switch (displayStatus) {
    case 'all_good':
    case 'completed':
      return 'green'
    case 'late':
      return 'orange'
    case 'no_show':
    case 'missed':
      return 'red'
    case 'cancelled':
      return 'gray'
    case 'covered':
      return 'purple'
    case 'covered_late':
      return 'pink'
    default:
      return 'blue'
  }
}

function formatTime(timeStr: string) {
  if (!timeStr) return ''
  const [hours, minutes] = timeStr.split(':')
  const h = parseInt(hours, 10)
  const suffix = h >= 12 ? 'PM' : 'AM'
  const hour12 = h % 12 || 12
  return `${hour12}:${minutes} ${suffix}`
}
</script>

<template>
  <div class="vol-detail">
    <VolunteerEditor
      :volunteer="volunteer"
      :isOpen="isEditorOpen"
      @close="isEditorOpen = false"
      @save="handleSave"
      @archive="handleArchive"
    />
    <Toast :show="showToast" message="Updated successfully" @close="showToast = false" />
    <!-- Header -->
    <header class="detail-header">
      <div class="profile-main">
        <div class="big-avatar">{{ volunteer.firstName[0] }}{{ volunteer.lastName[0] }}</div>
        <div class="header-text">
          <h1>{{ volunteer.firstName }} {{ volunteer.lastName }}</h1>
          <div class="badges">
            <span class="role-badge" :class="volunteer.role.toLowerCase().replace(' ', '-')">{{
              volunteer.role
            }}</span>
            <span class="status-badge" :class="volunteer.status">{{ volunteer.status }}</span>
          </div>
        </div>
      </div>
      <div class="actions">
        <Button v-if="canEdit" color="blue" @click="openEditor">Edit Profile</Button>
        <Button color="white">Message</Button>
      </div>
    </header>

    <!-- Metrics Bar (Gamification) -->
    <div class="metrics-bar">
      <div class="metric-item">
        <span class="metric-label">Total Hours</span>
        <span class="metric-value">{{ volunteer.totalHours || 0 }}</span>
      </div>
      <div class="metric-item">
        <span class="metric-label">Streak</span>
        <span class="metric-value">{{ volunteer.streak }} 🔥</span>
      </div>
      <div class="metric-item">
        <span class="metric-label">Reliability</span>
        <span
          class="metric-val-colored"
          :style="{
            color:
              volunteer.reliabilityScore >= 80
                ? 'var(--color-primary)'
                : volunteer.reliabilityScore >= 60
                  ? 'var(--color-warning)'
                  : 'var(--color-danger)',
          }"
          >{{ volunteer.reliabilityScore }}%</span
        >
      </div>
      <div class="metric-item badges-metric">
        <span class="metric-label">Badges</span>
        <div class="mini-badges">
          <span v-for="badge in volunteer.badges" :key="badge" class="mini-badge" :title="badge"
            >🏆</span
          >
          <span v-if="!volunteer.badges.length" class="no-badge">-</span>
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <div class="tabs">
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'overview' }"
        @click="activeTab = 'overview'"
      >
        Overview
      </button>
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'schedule' }"
        @click="activeTab = 'schedule'"
      >
        Schedule
      </button>
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'incidents' }"
        @click="activeTab = 'incidents'"
      >
        Incidents
      </button>
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'performance' }"
        @click="activeTab = 'performance'"
      >
        Performance
      </button>
      <button
        class="tab-btn"
        :class="{ active: activeTab === 'suggestions' }"
        @click="activeTab = 'suggestions'"
      >
        Suggestions
      </button>
    </div>

    <!-- Content -->
    <div class="tab-content">
      <!-- Overview Tab -->
      <div v-if="activeTab === 'overview'" class="overview-grid">
        <!-- Column 1: Personal Info -->
        <div class="left-col">
          <div class="card info-card">
            <h3>Contact Information</h3>
            <div class="info-row">
              <label>Email</label>
              <span>{{ volunteer.email }}</span>
            </div>
            <div class="info-row">
              <label>Phone</label>
              <span>{{ formatPhoneNumber(volunteer.phone) }}</span>
            </div>
            <div class="info-row">
              <label>Address</label>
              <span>{{ volunteer.address }}, {{ volunteer.city }} {{ volunteer.zip }}</span>
            </div>
            <div class="info-row">
              <label>Birthday</label>
              <span
                >{{ formatDate(volunteer.birthday) }}
                <span v-if="volunteer.age">({{ volunteer.age }} yo)</span></span
              >
            </div>
            <div class="info-row">
              <label>Emergency Contact</label>
              <span
                >{{ volunteer.emergencyContactName }}
                <span v-if="volunteer.emergencyContactPhone">
                  ({{ formatPhoneNumber(volunteer.emergencyContactPhone) }})
                </span>
              </span>
            </div>
          </div>

          <div class="card bio-card">
            <h3>Bio & Skills</h3>

            <div class="app-section">
              <label>Bio:</label>
              <p class="bio">{{ volunteer.bio || 'No bio provided.' }}</p>
            </div>

            <div class="app-section">
              <label>Skills:</label>
              <div class="skills-list">
                <Tag
                  v-for="skill in volunteer.skills"
                  :key="skill"
                  color="neutral"
                  style="margin-right: 8px; margin-bottom: 8px"
                >
                  {{ skill }}
                </Tag>
                <span v-if="!volunteer.skills.length" class="text-muted">None listed</span>
              </div>
            </div>

            <div class="info-row">
              <label>Allergies?</label>
              <span :class="{ 'text-red': volunteer.allergies }">{{
                volunteer.allergies ? 'Yes' : 'No'
              }}</span>
            </div>
          </div>
        </div>

        <!-- Column 2: Application Details -->
        <div class="right-col">
          <div class="card app-card">
            <h3>Application Details</h3>

            <div class="app-section">
              <label>Why they joined:</label>
              <p>{{ volunteer.interestReason || 'N/A' }}</p>
            </div>
            <div class="app-section">
              <label>Volunteer Experience:</label>
              <p>{{ volunteer.volunteerExperience || 'None listed.' }}</p>
            </div>
            <div class="app-section">
              <label>Position Preferences:</label>
              <div class="pref-list">
                <Tag
                  v-for="pref in volunteer.positionPreferences"
                  :key="pref"
                  color="neutral"
                  style="margin-right: 8px; margin-bottom: 8px"
                >
                  {{ pref }}
                </Tag>
              </div>
            </div>
            <div class="app-section">
              <label>Availability:</label>
              <div class="pref-list">
                <Tag
                  v-for="avail in volunteer.availability"
                  :key="avail"
                  variant="outline"
                  color="neutral"
                  style="margin-right: 8px; margin-bottom: 8px"
                >
                  {{ avail }}
                </Tag>
              </div>
            </div>
            <div class="app-info-row">
              <label>Joined Date:</label>
              <span>{{ formatDate(volunteer.joinDate) }}</span>
            </div>

            <div class="app-actions mt-4">
              <button class="view-app-btn">📄 View Original Application</button>
            </div>
          </div>
        </div>
      </div>

      <!-- Schedule Tab -->
      <div v-if="activeTab === 'schedule'" class="schedule-section">
        <div class="section-block">
          <div
            class="block-header"
            style="
              display: flex;
              justify-content: space-between;
              align-items: center;
              margin-bottom: 20px;
            "
          >
            <h3>Upcoming Shifts</h3>
            <div class="header-actions" style="display: flex; gap: 8px">
              <button
                v-if="
                  !showAllUpcoming &&
                  shifts.filter((s) => s.date >= new Date().toISOString().split('T')[0]).length > 2
                "
                class="text-btn"
                @click="showAllUpcoming = true"
              >
                View All
              </button>
              <button v-if="showAllUpcoming" class="text-btn" @click="showAllUpcoming = false">
                Show Less
              </button>
              <Button :color="isAddingShift ? 'white' : 'green'" @click="toggleAddShift">
                {{ isAddingShift ? 'Cancel' : '+ Add Shift' }}
              </Button>
            </div>
          </div>

          <ShiftForm
            v-if="isAddingShift"
            :initialData="editingShiftData"
            @close="toggleAddShift"
            @save="handleShiftSave"
          />

          <div v-if="upcomingShifts.length === 0" class="empty-msg">
            No upcoming shifts scheduled.
          </div>
          <div v-else class="shift-list">
            <div v-for="shift in upcomingShifts" :key="shift.id" class="shift-card">
              <div class="shift-date">
                <span class="day">{{ new Date(shift.date).getDate() }}</span>
                <span class="month">{{
                  new Date(shift.date).toLocaleDateString('en-US', { month: 'short' })
                }}</span>
              </div>
              <div class="shift-info">
                <div class="shift-role">{{ shift.role }}</div>
                <div class="shift-time">
                  {{ formatTime(shift.startTime) }} - {{ formatTime(shift.endTime) }}
                </div>
              </div>
              <div class="shift-status">
                <span class="status-pill" :class="getShiftStatusColor(shift)">{{
                  getDisplayStatus(shift)
                }}</span>
              </div>
              <div class="shift-actions">
                <button class="icon-btn edit-btn" @click="editShift(shift)" title="Edit">✏️</button>
                <button class="icon-btn delete-btn" @click="deleteShift(shift)" title="Delete">
                  🗑️
                </button>
              </div>
            </div>
          </div>
        </div>

        <div v-if="unverifiedShifts.length > 0" class="section-block mt-8">
          <h3>Shifts Needing Verification</h3>
          <div class="table-container">
            <table class="simple-table">
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Role</th>
                  <th>Time</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="shift in unverifiedShifts" :key="shift.id">
                  <td>{{ formatDate(shift.date) }}</td>
                  <td>{{ shift.role }}</td>
                  <td>{{ formatTime(shift.startTime) }} - {{ formatTime(shift.endTime) }}</td>
                  <td>
                    <div class="verify-actions">
                      <button
                        class="action-btn small green"
                        @click="verifyShift(shift, 'completed')"
                        title="Mark Completed"
                      >
                        ✅ Complete
                      </button>
                      <button
                        class="action-btn small red"
                        @click="verifyShift(shift, 'missed')"
                        title="Mark Missed"
                      >
                        ❌ Missed
                      </button>
                      <div class="verify-split-btn">
                        <button
                          class="action-btn small purple"
                          @click="verifyShift(shift, 'covered')"
                          title="Covered (>24h Notice)"
                        >
                          🔄 >24h
                        </button>
                        <button
                          class="action-btn small pink"
                          @click="verifyShift(shift, 'covered_late')"
                          title="Covered (<24h Notice)"
                        >
                          ⏰ &lt;24h
                        </button>
                      </div>
                      <button
                        class="action-btn small orange"
                        @click="markLate(shift)"
                        title="Mark Late & Edit Time"
                      >
                        ⚠️ Late
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="section-block mt-8">
          <h3>Past Shifts</h3>
          <div class="table-container">
            <table class="simple-table">
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Role</th>
                  <th>Time</th>
                  <th>Status</th>
                  <th>Notes</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="shift in verifiedPastShifts" :key="shift.id">
                  <td>{{ formatDate(shift.date) }}</td>
                  <td>{{ shift.role }}</td>
                  <td>{{ formatTime(shift.startTime) }} - {{ formatTime(shift.endTime) }}</td>
                  <td>
                    <span class="status-pill small" :class="getShiftStatusColor(shift)">{{
                      getDisplayStatus(shift)
                    }}</span>
                  </td>
                  <td class="notes">{{ shift.notes || '-' }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Incidents Tab -->
      <div v-if="activeTab === 'incidents'" class="incidents-section">
        <div class="incidents-header">
          <h3>Incident Log</h3>
          <button class="add-incident-btn">+ Log Incident</button>
        </div>

        <div v-if="incidents.length === 0" class="empty-state">
          <p>No incidents recorded for this volunteer. 🎉</p>
        </div>
        <div v-else class="incident-list">
          <div v-for="inc in incidents" :key="inc.id" class="incident-card">
            <div class="inc-meta">
              <span class="inc-date">{{ formatDate(inc.date) }}</span>
              <span class="inc-severity" :class="inc.severity">{{
                inc.severity.toUpperCase()
              }}</span>
            </div>
            <p class="inc-desc">{{ inc.description }}</p>
            <div class="inc-footer">Recorded by: {{ inc.recordedBy }}</div>
          </div>
        </div>
      </div>

      <!-- Performance Tab -->
      <div v-if="activeTab === 'performance'" class="performance-section">
        <div class="perf-row">
          <!-- Yearly Stats -->
          <div class="card perf-card">
            <div>
              <h3>Yearly Performance</h3>
            </div>
            <table class="simple-table perf-table">
              <thead>
                <tr>
                  <th>Year</th>
                  <th>Total Shifts</th>
                  <th>Hours</th>
                  <th>Reliability</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="stat in yearlyStats" :key="stat.year">
                  <td class="year-cell">{{ stat.year }}</td>
                  <td>{{ stat.shifts }}</td>
                  <td>{{ stat.hours }} hrs</td>
                  <td>
                    <span
                      class="status-pill small"
                      :class="
                        stat.reliability >= 80 ? 'green' : stat.reliability >= 60 ? 'orange' : 'red'
                      "
                    >
                      {{ stat.reliability }}%
                    </span>
                  </td>
                </tr>
                <tr v-if="yearlyStats.length === 0">
                  <td colspan="4" class="text-center">No data available yet.</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="badges-row mt-8">
          <h3>Badges & Achievements</h3>
          <div class="badges-grid">
            <div
              v-for="badge in availableBadges"
              :key="badge.id"
              class="badge-card"
              :class="{ earned: volunteer.badges.includes(badge.name) }"
            >
              <div class="badge-icon">{{ badge.icon }}</div>
              <div class="badge-info">
                <div class="badge-name">{{ badge.name }}</div>
                <div class="badge-desc">{{ badge.description }}</div>
              </div>
              <div class="badge-status">
                {{ volunteer.badges.includes(badge.name) ? 'EARNED' : 'LOCKED' }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Suggestions Tab -->
      <div v-if="activeTab === 'suggestions'" class="suggestions-section">
        <div class="card suggestion-card">
          <h3>Suggestions for Improvement</h3>
          <div class="suggestion-list">
            <div v-for="(item, idx) in suggestions" :key="idx" class="suggestion-item">
              <div class="sug-icon">{{ item.icon }}</div>
              <div class="sug-content">
                <div class="sug-title">{{ item.title }}</div>
                <div class="sug-desc">{{ item.desc }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.vol-detail {
  padding: 32px;
  height: 100%;
  padding-bottom: 60px; /* Space for scroll */
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.profile-main {
  display: flex;
  gap: 20px;
  align-items: center;
}

.big-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: hsl(from var(--color-secondary) h s 90%);
  color: var(--color-secondary);
  font-size: 2rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-text h1 {
  margin: 0 0 8px 0;
  font-size: 1.8rem;
  color: var(--text-primary);
}

.badges {
  display: flex;
  gap: 8px;
}

.role-badge {
  background: hsl(from var(--color-neutral) h s 95%);
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: capitalize;
  color: var(--text-primary);

  &.teen {
    background: hsl(from var(--color-secondary) h s 90%);
    color: var(--color-secondary);
  }
  &.tier-2 {
    background: hsl(from var(--color-warning) h s 90%);
    color: var(--color-warning);
  }
  &.tier-1 {
    background: hsl(from var(--color-neutral) h s 90%);
    color: hsl(from var(--color-neutral) h s 30%);
  }
}

.status-badge {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: capitalize;

  &.active {
    background: hsl(from var(--color-primary) h s 90%);
    color: var(--color-primary);
  }
  &.inactive {
    background: hsl(from var(--color-danger) h s 90%);
    color: var(--color-danger);
  }
  &.pending {
    background: hsl(from var(--color-warning) h s 90%);
    color: var(--color-warning);
  }
}

.actions {
  display: flex;
  gap: 12px;
}

.action-btn {
  padding: 8px 16px;
  border-radius: 8px;
  border: 1px solid transparent;
  background: var(--color-secondary);
  color: var(--text-inverse);
  font-weight: 600;
  cursor: pointer;

  &.outline {
    background: var(--text-inverse);
    border: 1px solid var(--border-color);
    color: var(--text-primary);
    &:hover {
      background: hsl(from var(--color-neutral) h s 98%);
    }
  }
}

/* Metrics Bar */
.metrics-bar {
  display: flex;
  gap: 24px;
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 16px 24px;
  margin-bottom: 24px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
}

.tab-content {
  min-height: 400px; /* Prevent layout shift */
}

.metric-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding-right: 24px;
  border-right: 1px solid var(--border-color);

  &:last-child {
    border-right: none;
  }
}

.metric-label {
  font-size: 0.75rem;
  text-transform: uppercase;
  color: hsl(from var(--color-neutral) h s 50%);
  font-weight: 600;
}

.metric-value {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--text-primary);
}

.metric-val-colored {
  font-size: 1.2rem;
  font-weight: 700;
}

.score-badge {
  &.high {
    color: var(--color-primary);
  }
  &.mid {
    color: var(--color-warning);
  }
  &.low {
    color: var(--color-danger);
  }
}

.mini-badge {
  font-size: 1.2rem;
  cursor: help;
}

.tabs {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid var(--border-color);
  margin-bottom: 24px;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 0;
  font-size: 1rem;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  text-transform: capitalize;

  &.active {
    color: var(--color-secondary);
    border-bottom-color: var(--color-secondary);
    font-weight: 600;
  }
}

/* Overview Grid */
.overview-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  align-items: start;
}

.left-col,
.right-col {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.card {
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
  height: fit-content;

  h3 {
    margin: 0 0 20px 0;
    font-size: 1.1rem;
    color: var(--text-primary);
    border-bottom: 1px solid var(--border-color);
    padding-bottom: 12px;
  }
}

.info-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
  font-size: 0.95rem;

  label {
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 500;
  }
  span {
    font-weight: 500;
    color: var(--text-primary);
    text-align: right;
    max-width: 60%;
  }
}

.skill-tag {
  display: inline-block;
  background: hsl(from var(--color-neutral) h s 95%);
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-right: 8px;
  margin-bottom: 8px;
  color: var(--text-primary);
}

.bio {
  color: hsl(from var(--color-neutral) h s 50%);
  line-height: 1.5;
  /* margin-bottom handled by app-section */
}

.app-section {
  margin-bottom: 20px;
  label {
    display: block;
    font-size: 0.85rem;
    color: hsl(from var(--color-neutral) h s 50%);
    margin-bottom: 8px;
    font-weight: 600;
  }
  p {
    margin: 0;
    color: var(--text-primary);
    line-height: 1.5;
    background: hsl(from var(--color-neutral) h s 98%);
    padding: 12px;
    border-radius: 8px;
  }
}

.text-red {
  color: var(--color-danger);
  font-weight: 700;
}

.app-info-row {
  display: flex;
  justify-content: space-between;
  margin-top: 24px;
  padding-top: 16px;
  border-top: 1px solid var(--border-color);
  font-size: 0.9rem;
  label {
    color: hsl(from var(--color-neutral) h s 50%);
  }
  span {
    font-weight: 600;
  }
}

.view-app-btn {
  width: 100%;
  margin-top: 16px;
  padding: 10px;
  background: var(--text-inverse);
  border: 1px dashed var(--border-color);
  color: hsl(from var(--color-neutral) h s 50%);
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s;

  &:hover {
    border-color: var(--color-secondary);
    color: var(--color-secondary);
    background: hsl(from var(--color-secondary) h s 98%);
  }
}

/* Schedule */
.shift-card {
  display: flex;
  align-items: center;
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  padding: 16px;
  border-radius: 12px;
  margin-bottom: 12px;
  gap: 20px;
}

.shift-date {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: hsl(from var(--color-neutral) h s 98%);
  padding: 8px 16px;
  border-radius: 8px;

  .day {
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--text-primary);
  }
  .month {
    font-size: 0.8rem;
    text-transform: uppercase;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.shift-info {
  flex: 1;
  .shift-role {
    font-weight: 600;
    font-size: 1.1rem;
    margin-bottom: 4px;
  }
  .shift-time {
    color: hsl(from var(--color-neutral) h s 50%);
    font-size: 0.9rem;
  }
}

.status-pill {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: capitalize;

  &.green {
    background: hsl(from var(--color-primary) h s 95%);
    color: var(--color-primary);
  }
  &.orange {
    background: hsl(from var(--color-warning) h s 95%);
    color: var(--color-warning);
  }
  &.red {
    background: hsl(from var(--color-danger) h s 95%);
    color: var(--color-danger);
  }
  &.gray {
    background: hsl(from var(--color-neutral) h s 95%);
    color: hsl(from var(--color-neutral) h s 50%);
  }
  &.blue {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
  }

  &.small {
    font-size: 0.75rem;
    padding: 2px 8px;
  }
}

.mt-8 {
  margin-top: 32px;
}

.table-container {
  overflow-x: auto; /* Fix layout break */
}

.simple-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 600px; /* Force spread */

  th {
    text-align: left;
    padding: 12px;
    color: hsl(from var(--color-neutral) h s 50%);
    font-size: 0.9rem;
    border-bottom: 1px solid var(--border-color);
  }
  td {
    padding: 12px;
    border-bottom: 1px solid var(--border-color);
    font-size: 0.95rem;
  }
}

/* Incidents */
.incidents-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.add-incident-btn {
  background: var(--text-inverse);
  border: 1px solid var(--color-danger);
  color: var(--color-danger);
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  &:hover {
    background: hsl(from var(--color-danger) h s 98%);
  }
}

.incident-card {
  border: 1px solid hsl(from var(--color-danger) h s 80%);
  background: hsl(from var(--color-danger) h s 98%);
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;
}

.inc-meta {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.inc-severity {
  font-size: 0.75rem;
  background: white;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;

  &.low {
    color: var(--color-warning);
    border: 1px solid hsl(from var(--color-warning) h s 80%);
  }
  &.medium {
    color: hsl(from var(--color-warning) h s 40%);
    border: 1px solid hsl(from var(--color-warning) h s 60%);
  }
  &.high {
    color: var(--color-danger);
    border: 1px solid hsl(from var(--color-danger) h s 80%);
  }
}

.inc-desc {
  margin: 0 0 12px 0;
  color: hsl(from var(--color-danger) h s 30%);
  line-height: 1.5;
}

.inc-footer {
  font-size: 0.85rem;
  color: hsl(from var(--color-danger) h s 40%);
  font-style: italic;
}

.empty-state {
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
  padding: 40px;
  background: hsl(from var(--color-neutral) h s 98%);
  border-radius: 12px;
}

/* Performance */
.perf-card {
  width: 100%;
}

.perf-table td {
  padding: 16px 12px;
}

.year-cell {
  font-weight: 700;
  color: var(--color-secondary);
}

.badges-row h3 {
  margin-bottom: 16px;
  font-size: 1.1rem;
  color: var(--text-primary);
}

.badges-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.badge-card {
  display: flex;
  align-items: center;
  padding: 16px;
  background: hsl(from var(--color-neutral) h s 98%);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  gap: 16px;
  opacity: 0.7;
  filter: grayscale(1);
  transition: all 0.2s;

  &.earned {
    opacity: 1;
    filter: none;
    background: var(--text-inverse);
    border-color: var(--color-warning);
    box-shadow: 0 4px 6px hsla(from var(--color-warning) h s l / 0.1);
  }
}

.badge-icon {
  font-size: 2rem;
}

.badge-info {
  flex: 1;
}

.badge-name {
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.badge-desc {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  line-height: 1.3;
}

.badge-status {
  font-size: 0.7rem;
  font-weight: 800;
  color: hsl(from var(--color-neutral) h s 60%);
  letter-spacing: 0.5px;

  .earned & {
    color: var(--color-warning);
  }
}

.text-center {
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
}

.suggestion-card {
  width: 100%;
  border-left: 4px solid var(--color-secondary);
}

.suggestion-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.suggestion-item {
  display: flex;
  gap: 16px;
  align-items: flex-start;
}

.sug-icon {
  font-size: 1.5rem;
  background: hsl(from var(--color-neutral) h s 95%);
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.sug-title {
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.sug-desc {
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
  line-height: 1.4;
}
</style>

<style scoped>
.shift-card {
  display: flex;
  align-items: center;
}

.shift-actions {
  margin-left: auto;
  display: flex;
  gap: 8px;
  /* Always show for now to ensure visibility on touch/mobile if needed, or stick to opacity */
  opacity: 0;
  transition: opacity 0.2s;
}

.shift-card:hover .shift-actions {
  opacity: 1;
}

.icon-btn {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 1.1rem;
  padding: 4px;
  border-radius: 4px;
  transition: background 0.2s;
}

.edit-btn:hover {
  background: hsl(from var(--color-warning) h s 90%);
}

.delete-btn:hover {
  background: hsl(from var(--color-danger) h s 90%);
}

.verify-actions {
  display: flex;
  gap: 8px;
}

.action-btn.small {
  padding: 4px 8px;
  font-size: 0.75rem;
  border-radius: 4px;
}

.action-btn.green {
  background: hsl(from var(--color-primary) h s 95%);
  color: var(--color-primary);
  border: 1px solid var(--color-primary);
}
.action-btn.green:hover {
  background: var(--color-primary);
  color: white;
}

.action-btn.red {
  background: hsl(from var(--color-danger) h s 95%);
  color: var(--color-danger);
  border: 1px solid var(--color-danger);
}
.action-btn.red:hover {
  background: var(--color-danger);
  color: white;
}

.action-btn.purple {
  background: hsl(270 95% 95%); /* Approximating purple */
  color: hsl(270 60% 50%);
  border: 1px solid hsl(270 60% 50%);
}
.action-btn.purple:hover {
  background: hsl(270 60% 50%);
  color: white;
}

.action-btn.orange {
  background: hsl(from var(--color-warning) h s 95%);
  color: var(--color-warning);
  border: 1px solid var(--color-warning);
}
.action-btn.orange:hover {
  background: var(--color-warning);
  color: white;
}

.action-btn.pink {
  background: hsl(330 60% 90%);
  color: hsl(330 60% 40%);
  border: 1px solid hsl(330 60% 40%);
}
.action-btn.pink:hover {
  background: hsl(330 60% 50%);
  color: white;
}

.text-btn {
  background: none;
  border: none;
  color: var(--color-primary);
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  padding: 4px 8px;
}
.text-btn:hover {
  text-decoration: underline;
}
</style>
