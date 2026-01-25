<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

import VolunteerDetail from '../../components/admin/volunteers/VolunteerDetail.vue'
import VolunteerEditor from '../../components/admin/volunteers/VolunteerEditor.vue'
import VolunteerList from '../../components/admin/volunteers/VolunteerList.vue'
import { type IShift, type IVolunteer, mockIncidents } from '../../stores/mockVolunteerData'
import {
  calculateReliabilityScore,
  calculateStreak,
  calculateTotalHours,
} from '../../utils/reliability'

const allVolunteers = ref<IVolunteer[]>([])
const isLoading = ref(true)
const isSaving = ref(false)

async function fetchVolunteers() {
  isLoading.value = true
  try {
    const res = await fetch('/v1/volunteers?page_size=100')
    if (res.ok) {
      const data = await res.json()

      allVolunteers.value = data.data.volunteers || []
    }
  } catch (err) {
    console.error('Failed to fetch volunteers', err)
  } finally {
    isLoading.value = false
  }
}

const selectedVolunteerId = ref<string | null>(null)

const selectedVolunteer = computed(() => {
  return (
    allVolunteers.value.find(
      (v) => v.id == selectedVolunteerId.value ,
    ) || null
  )
})

onMounted(async () => {
  await fetchVolunteers()

  if (!selectedVolunteerId.value && allVolunteers.value.length > 0) {

    selectedVolunteerId.value = allVolunteers.value[0].id
  }
})

const shifts = ref<IShift[]>([])

async function fetchShifts(volunteerId: string) {
  try {
    const res = await fetch(`/v1/volunteers/${volunteerId}/shifts?_t=${Date.now()}`)
    if (res.ok) {
      const data = await res.json()
      shifts.value = data.data.shifts || []

      const currentYear = new Date().getFullYear().toString()
      const currentYearShifts = shifts.value.filter((s) => s.date.startsWith(currentYear))

      const volIndex = allVolunteers.value.findIndex((v) => v.id == volunteerId)
      if (volIndex !== -1) {
        allVolunteers.value[volIndex].reliabilityScore =
          calculateReliabilityScore(currentYearShifts)
        allVolunteers.value[volIndex].totalHours = calculateTotalHours(currentYearShifts)
        allVolunteers.value[volIndex].streak = calculateStreak(currentYearShifts)
      }
    }
  } catch (err) {
    console.error('Failed to fetch shifts', err)
  }
}


watch(selectedVolunteerId, (newId) => {
  if (newId) {
    fetchShifts(newId)
  } else {
    shifts.value = []
  }
})

const selectedShifts = computed(() => shifts.value)

const selectedIncidents = computed(() => {
  if (!selectedVolunteerId.value) return []
  return mockIncidents
    .filter((i) => i.volunteerId == selectedVolunteerId.value)
    .sort((a, b) => b.date.localeCompare(a.date))
})

const isCreating = ref(false)

function sanitizeVolunteerData(data: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
  const payload = { ...data }

  if (!payload.birthday) payload.birthday = null
  if (!payload.joinDate) payload.joinDate = null
  console.log('Sanitized payload:', payload)
  return payload
}

function handleOpenCreate() {
  isCreating.value = true
}

async function handleCreateSave(newVolunteer: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
  isSaving.value = true
  try {
    const res = await fetch('/v1/volunteers', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(sanitizeVolunteerData(newVolunteer)),
    })

    if (res.ok) {
      await fetchVolunteers()
      const data = await res.json()
      if (data.data && data.data.volunteer) {
        selectedVolunteerId.value = data.data.volunteer.id
      }
      isCreating.value = false
    } else {
      const errText = await res.text()
      console.error('Create failed:', res.status, errText)
      alert(`Failed to create volunteer: ${errText}`)
    }
  } catch (e) {
    console.error(e)
    alert('Error creating volunteer')
  } finally {
    isSaving.value = false
  }
}

async function handleUpdateSave(updatedData: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
  if (!selectedVolunteerId.value) return

  isSaving.value = true
  try {
    const res = await fetch(`/v1/volunteers/${selectedVolunteerId.value}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(sanitizeVolunteerData(updatedData)),
    })

    if (res.ok) {
      await fetchVolunteers()

    } else {
      const errText = await res.text()
      console.error('Update failed:', res.status, errText)
      alert(`Failed to update volunteer: ${errText}`)
    }
  } catch (e) {
    console.error('Error updating volunteer', e)
    alert('Error updating volunteer')
  } finally {
    isSaving.value = false
  }
}

async function handleAddShift(shiftData: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
  if (!selectedVolunteerId.value) return

  const shiftsToCreate = []
  const baseDate = new Date(shiftData.date)

  if (shiftData.isRecurring) {
    const endDate = shiftData.endDate
      ? new Date(shiftData.endDate)
      : new Date(baseDate.getTime() + 90 * 24 * 60 * 60 * 1000)

    const currentDate = new Date(baseDate)

    let count = 0
    while (currentDate <= endDate && count < 50) {
      shiftsToCreate.push({
        volunteerId: selectedVolunteerId.value,
        date: currentDate.toISOString().split('T')[0],
        startTime: shiftData.startTime,
        endTime: shiftData.endTime,
        role: shiftData.role,
      })

      if (shiftData.frequency === 'weekly') {
        currentDate.setDate(currentDate.getDate() + 7)
      } else if (shiftData.frequency === 'biweekly') {
        currentDate.setDate(currentDate.getDate() + 14)
      } else if (shiftData.frequency === 'monthly') {
        currentDate.setMonth(currentDate.getMonth() + 1)
      } else {
        break
      }
      count++
    }
  } else {
    shiftsToCreate.push({
      volunteerId: selectedVolunteerId.value,
      date: shiftData.date,
      startTime: shiftData.startTime,
      endTime: shiftData.endTime,
      role: shiftData.role,
    })
  }

  try {
    for (const s of shiftsToCreate) {
      await fetch('/v1/shifts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(s),
      })
    }

    await new Promise((resolve) => setTimeout(resolve, 300))

    await fetchShifts(selectedVolunteerId.value)

    await fetchVolunteers()
  } catch (e) {
    console.error('Error creating shifts', e)
    alert('Failed to create shifts')
  }
}

async function handleUpdateShift(shiftData: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
  if (!shiftData.id) return
  try {

    const payload = {
      id: shiftData.id,
      volunteerId: shiftData.volunteerId,
      date: shiftData.date,
      startTime: shiftData.startTime,
      endTime: shiftData.endTime,
      role: shiftData.role,
      status: shiftData.status,
      notes: shiftData.notes,
    }

    const res = await fetch(`/v1/shifts/${shiftData.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    })

    if (res.ok) {
      await fetchShifts(selectedVolunteerId.value!)
      await fetchVolunteers()
    } else {
      const err = await res.text()
      alert(`Failed to update shift: ${  err}`)
    }
  } catch (e) {
    console.error('Error updating shift', e)
    alert('Error updating shift')
  }
}

async function handleDeleteShift(shiftId: string | number) {
  if (!confirm('Are you sure you want to delete this shift?')) return

  try {
    const res = await fetch(`/v1/shifts/${shiftId}`, {
      method: 'DELETE',
    })

    if (res.ok) {
      await fetchShifts(selectedVolunteerId.value!)
      await fetchVolunteers()
    } else {
      alert('Failed to delete shift')
    }
  } catch (e) {
    console.error('Error deleting shift', e)
    alert('Error deleting shift')
  }
}
</script>

<template>
  <div class="volunteers-page">
    <Teleport to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Volunteers</h1>
    </Teleport>

    <aside class="sidebar">
      <VolunteerList
        :volunteers="allVolunteers"
        :loading="isLoading"
        :selectedId="String(selectedVolunteerId)"
        @select="(vol) => (selectedVolunteerId = vol.id)"
        @add="handleOpenCreate"
      />
    </aside>

    <VolunteerEditor
      :volunteer="null"
      :isOpen="isCreating"
      :isSaving="isSaving"
      @close="isCreating = false"
      @save="handleCreateSave"
    />

    <main class="main-content">
      <VolunteerDetail
        v-if="selectedVolunteer"
        :volunteer="selectedVolunteer"
        :volunteers="allVolunteers"
        :shifts="selectedShifts"
        :incidents="selectedIncidents"
        :isSaving="isSaving"
        @add-shift="handleAddShift"
        @update="handleUpdateSave"
        @update-shift="handleUpdateShift"
        @delete-shift="handleDeleteShift"
      />

      <div v-else class="empty-selection">
        <div class="icon">🤝</div>
        <h2>Select a volunteer</h2>
        <p>View their profile, schedule, and incident logs.</p>
      </div>
    </main>
  </div>
</template>

<style scoped>
.volunteers-page {
  display: flex;
  height: calc(100vh - 80px);
  gap: 24px;
  max-width: 1400px;
  margin: 0 auto;
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.sidebar {
  width: 320px;
  background: var(--text-inverse);
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 4px 12px rgb(0 0 0 / 3%);
  border: 1px solid var(--border-color);
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;

  h2 {
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--text-primary);
    margin: 0;
  }
}

.invite-btn {
  background: var(--color-secondary);
  color: var(--text-inverse);
  border: none;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.2rem;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    opacity: 0.9;
  }
}

.search-box {
  padding: 16px;
  border-bottom: 1px solid var(--border-color);

  input {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    background: hsl(from var(--color-neutral) h s 98%);
    outline: none;

    &:focus {
      border-color: var(--color-secondary);
      background: var(--text-inverse);
    }
  }
}

.main-content {
  flex: 1;
  background: var(--text-inverse);
  border-radius: 16px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 3%);
  border: 1px solid var(--border-color);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  scrollbar-gutter: stable;
}

.empty-selection {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: hsl(from var(--color-neutral) h s 50%);

  .icon {
    font-size: 4rem;
    margin-bottom: 24px;
    opacity: 0.5;
  }

  h2 {
    color: var(--text-primary);
    margin-bottom: 8px;
  }
}
</style>
<style scoped>
.admin-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
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

.add-btn {
  background-color: var(--color-secondary);
  color: var(--text-inverse);
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: hsl(from var(--color-secondary) h s 40%);
  }
}

.placeholder-content {
  background: var(--text-inverse);
  padding: 48px;
  border-radius: 12px;
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
  border: 1px dashed var(--border-color);
}
</style>
