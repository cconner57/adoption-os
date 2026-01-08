<script setup lang="ts">
import { ref, computed } from 'vue'
import type { IVolunteer, IShift, IIncident } from '../../../stores/mockVolunteerData'
import { availableBadges } from '../../../stores/mockVolunteerData'

const props = defineProps<{
  volunteer: IVolunteer
  shifts: IShift[]
  incidents: IIncident[]
}>()

const activeTab = ref<'overview' | 'schedule' | 'incidents' | 'performance' | 'suggestions'>(
  'overview',
)
const tabs = ['overview', 'schedule', 'incidents', 'performance', 'suggestions'] as const

const upcomingShifts = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return props.shifts.filter((s) => s.date >= today).sort((a, b) => a.date.localeCompare(b.date))
})

const pastShifts = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return props.shifts.filter((s) => s.date < today).sort((a, b) => b.date.localeCompare(a.date))
})

const yearlyStats = computed(() => {
  const stats: Record<string, { year: string; hours: number; shifts: number; missed: number }> = {}

  // Initialize with join year if no shifts? No, just use shift data for accuracy.
  props.shifts.forEach((shift) => {
    const year = shift.date.split('-')[0]
    if (!stats[year]) {
      stats[year] = { year, hours: 0, shifts: 0, missed: 0 }
    }

    // Calculate hours
    const start = parseInt(shift.startTime.split(':')[0])
    const end = parseInt(shift.endTime.split(':')[0])
    const duration = end - start

    // Logic: Only count hours if "all_good" or "late" (assuming late still worked some).
    // Let's assume completed shifts count.
    if (['all_good', 'late'].includes(shift.status)) {
      stats[year].hours += duration
    }

    stats[year].shifts += 1

    if (['no_show'].includes(shift.status)) {
      stats[year].missed += 1
    }
  })

  // Calculate reliability % per year
  return Object.values(stats)
    .map((s) => {
      const reliability = s.shifts > 0 ? Math.round(((s.shifts - s.missed) / s.shifts) * 100) : 0
      return { ...s, reliability }
    })
    .sort((a, b) => b.year.localeCompare(a.year))
})

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

// Updated to include full year
function formatDate(dateStr: string) {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleDateString('en-US', {
    weekday: 'short',
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

function getShiftStatusColor(status: string) {
  switch (status) {
    case 'all_good':
      return 'green'
    case 'late':
      return 'orange'
    case 'no_show':
      return 'red'
    case 'cancelled':
      return 'gray'
    default:
      return 'blue'
  }
}
</script>

<template>
  <div class="vol-detail">
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
        <button class="action-btn">Edit Profile</button>
        <button class="action-btn outline">Message</button>
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
        <span class="metric-value">{{ volunteer.streak || 0 }} 🔥</span>
      </div>
      <div class="metric-item">
        <span class="metric-label">Reliability</span>
        <span
          class="metric-val-colored"
          :style="{ color: volunteer.reliabilityScore >= 90 ? 'var(--green)' : 'var(--orange)' }"
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
              <span>{{ volunteer.phone }}</span>
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
                >{{ volunteer.emergencyContactName }} ({{ volunteer.emergencyContactPhone }})</span
              >
            </div>
          </div>

          <div class="card bio-card">
            <h3>Bio & Skills</h3>
            <p class="bio">{{ volunteer.bio || 'No bio provided.' }}</p>
            <div class="info-row">
              <label>Allergies?</label>
              <span :class="{ 'text-red': volunteer.allergies }">{{
                volunteer.allergies ? 'Yes' : 'No'
              }}</span>
            </div>
            <div class="skills-list">
              <span v-for="skill in volunteer.skills" :key="skill" class="skill-tag">
                {{ skill }}
              </span>
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
                <span v-for="pref in volunteer.positionPreferences" :key="pref" class="pref-tag">{{
                  pref
                }}</span>
              </div>
            </div>
            <div class="app-section">
              <label>Availability:</label>
              <div class="pref-list">
                <span
                  v-for="avail in volunteer.availability"
                  :key="avail"
                  class="pref-tag outline"
                  >{{ avail }}</span
                >
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
          <h3>Upcoming Shifts</h3>
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
                <div class="shift-time">{{ shift.startTime }} - {{ shift.endTime }}</div>
              </div>
              <div class="shift-status">
                <span class="status-pill" :class="getShiftStatusColor(shift.status)">{{
                  shift.status.replace('_', ' ')
                }}</span>
              </div>
            </div>
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
                <tr v-for="shift in pastShifts" :key="shift.id">
                  <td>{{ formatDate(shift.date) }}</td>
                  <td>{{ shift.role }}</td>
                  <td>{{ shift.startTime }} - {{ shift.endTime }}</td>
                  <td>
                    <span class="status-pill small" :class="getShiftStatusColor(shift.status)">{{
                      shift.status.replace('_', ' ')
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
            <h3>Yearly Performance</h3>
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
                        stat.reliability >= 90 ? 'green' : stat.reliability >= 70 ? 'orange' : 'red'
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
  background: var(--purple-light);
  color: var(--purple);
  font-size: 2rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-text h1 {
  margin: 0 0 8px 0;
  font-size: 1.8rem;
  color: var(--font-color-dark);
}

.badges {
  display: flex;
  gap: 8px;
}

.role-badge {
  background: #f3f4f6;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: capitalize;
  color: var(--font-color-dark);

  &.teen {
    background: #e0f2fe;
    color: #0369a1;
  }
  &.tier-2 {
    background: #fef3c7;
    color: #b45309;
  }
  &.tier-1 {
    background: #f3f4f6;
    color: #374151;
  }
}

.status-badge {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 0.85rem;
  font-weight: 600;
  text-transform: capitalize;

  &.active {
    background: #dcfce7;
    color: #166534;
  }
  &.inactive {
    background: #fee2e2;
    color: #991b1b;
  }
  &.pending {
    background: #ffedd5;
    color: #9a3412;
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
  background: var(--purple);
  color: white;
  font-weight: 600;
  cursor: pointer;

  &.outline {
    background: white;
    border: 1px solid #e5e7eb;
    color: var(--font-color-dark);
    &:hover {
      background: #f9fafb;
    }
  }
}

/* Metrics Bar */
.metrics-bar {
  display: flex;
  gap: 24px;
  background: white;
  border: 1px solid #e5e7eb;
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
  border-right: 1px solid #f3f4f6;

  &:last-child {
    border-right: none;
  }
}

.metric-label {
  font-size: 0.75rem;
  text-transform: uppercase;
  color: var(--font-color-medium);
  font-weight: 600;
}

.metric-value {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--font-color-dark);
}

.metric-val-colored {
  font-size: 1.2rem;
  font-weight: 700;
}

.mini-badge {
  font-size: 1.2rem;
  cursor: help;
}

.tabs {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 24px;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 0;
  font-size: 1rem;
  color: var(--font-color-medium);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  text-transform: capitalize;

  &.active {
    color: var(--purple);
    border-bottom-color: var(--purple);
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
  background: white;
  border: 1px solid #f3f4f6;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
  height: fit-content;

  h3 {
    margin: 0 0 20px 0;
    font-size: 1.1rem;
    color: var(--font-color-dark);
    border-bottom: 1px solid #f3f4f6;
    padding-bottom: 12px;
  }
}

.info-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
  font-size: 0.95rem;

  label {
    color: var(--font-color-medium);
    font-weight: 500;
  }
  span {
    font-weight: 500;
    color: var(--font-color-dark);
    text-align: right;
    max-width: 60%;
  }
}

.skill-tag,
.pref-tag {
  display: inline-block;
  background: #f3f4f6;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-right: 8px;
  margin-bottom: 8px;
  color: var(--font-color-dark);

  &.outline {
    background: white;
    border: 1px solid #e5e7eb;
  }
}

.bio {
  color: var(--font-color-medium);
  line-height: 1.5;
  margin-bottom: 16px;
}

.app-section {
  margin-bottom: 20px;
  label {
    display: block;
    font-size: 0.85rem;
    color: var(--font-color-medium);
    margin-bottom: 8px;
    font-weight: 600;
  }
  p {
    margin: 0;
    color: var(--font-color-dark);
    line-height: 1.5;
    background: #f9fafb;
    padding: 12px;
    border-radius: 8px;
  }
}

.text-red {
  color: var(--red);
  font-weight: 700;
}

.app-info-row {
  display: flex;
  justify-content: space-between;
  margin-top: 24px;
  padding-top: 16px;
  border-top: 1px solid #f3f4f6;
  font-size: 0.9rem;
  label {
    color: var(--font-color-medium);
  }
  span {
    font-weight: 600;
  }
}

.view-app-btn {
  width: 100%;
  margin-top: 16px;
  padding: 10px;
  background: white;
  border: 1px dashed #cbd5e1;
  color: var(--font-color-medium);
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s;

  &:hover {
    border-color: var(--purple);
    color: var(--purple);
    background: #fdf4ff;
  }
}

/* Schedule */
.shift-card {
  display: flex;
  align-items: center;
  background: white;
  border: 1px solid #e5e7eb;
  padding: 16px;
  border-radius: 12px;
  margin-bottom: 12px;
  gap: 20px;
}

.shift-date {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #f9fafb;
  padding: 8px 16px;
  border-radius: 8px;

  .day {
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--font-color-dark);
  }
  .month {
    font-size: 0.8rem;
    text-transform: uppercase;
    color: var(--font-color-medium);
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
    color: var(--font-color-medium);
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
    background: #dcfce7;
    color: #166534;
  }
  &.orange {
    background: #ffedd5;
    color: #9a3412;
  }
  &.red {
    background: #fee2e2;
    color: #991b1b;
  }
  &.gray {
    background: #f3f4f6;
    color: #6b7280;
  }
  &.blue {
    background: #dbeafe;
    color: #1e40af;
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
    color: var(--font-color-medium);
    font-size: 0.9rem;
    border-bottom: 1px solid #e5e7eb;
  }
  td {
    padding: 12px;
    border-bottom: 1px solid #f3f4f6;
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
  background: white;
  border: 1px solid #ef4444;
  color: #ef4444;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  &:hover {
    background: #fef2f2;
  }
}

.incident-card {
  border: 1px solid #fca5a5;
  background: #fef2f2;
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
    color: #d97706;
    border: 1px solid #fbbf24;
  }
  &.medium {
    color: #ea580c;
    border: 1px solid #fdba74;
  }
  &.high {
    color: #dc2626;
    border: 1px solid #fca5a5;
  }
}

.inc-desc {
  margin: 0 0 12px 0;
  color: #7f1d1d;
  line-height: 1.5;
}

.inc-footer {
  font-size: 0.85rem;
  color: #991b1b;
  font-style: italic;
}

.empty-state {
  text-align: center;
  color: var(--font-color-medium);
  padding: 40px;
  background: #f9fafb;
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
  color: var(--purple);
}

.badges-row h3 {
  margin-bottom: 16px;
  font-size: 1.1rem;
  color: var(--font-color-dark);
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
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  gap: 16px;
  opacity: 0.7;
  filter: grayscale(1);
  transition: all 0.2s;

  &.earned {
    opacity: 1;
    filter: none;
    background: white;
    border-color: #fbbf24;
    box-shadow: 0 4px 6px rgba(251, 191, 36, 0.1);
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
  color: var(--font-color-dark);
  margin-bottom: 4px;
}

.badge-desc {
  font-size: 0.8rem;
  color: var(--font-color-medium);
  line-height: 1.3;
}

.badge-status {
  font-size: 0.7rem;
  font-weight: 800;
  color: #9ca3af;
  letter-spacing: 0.5px;

  .earned & {
    color: #d97706;
  }
}

.text-center {
  text-align: center;
  color: var(--font-color-medium);
}

.suggestion-card {
  width: 100%;
  border-left: 4px solid var(--purple);
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
  background: #f3f4f6;
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.sug-title {
  font-weight: 700;
  color: var(--font-color-dark);
  margin-bottom: 4px;
}

.sug-desc {
  font-size: 0.9rem;
  color: var(--font-color-medium);
  line-height: 1.4;
}
</style>
