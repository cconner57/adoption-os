<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { mockShifts, mockIncidents } from '../../stores/mockVolunteerData'
import VolunteerList from '../../components/admin/volunteers/VolunteerList.vue'
import VolunteerDetail from '../../components/admin/volunteers/VolunteerDetail.vue'
import VolunteerEditor from '../../components/admin/volunteers/VolunteerEditor.vue'

// Stores
const allVolunteers = ref<IVolunteer[]>([])
const isLoading = ref(true)

// Fetch Volunteers
async function fetchVolunteers() {
  isLoading.value = true
  try {
    const res = await fetch('/v1/volunteers?page_size=100')
    if (res.ok) {
      const data = await res.json()
      // Backend returns { volunteers: [], metadata: {} }
      // Map backend fields to frontend interface if needed
      // Currently backend uses camelCase JSON tags so it should match mostly
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
      (v) => v.id == selectedVolunteerId.value /* loose match for string/int ids */,
    ) || null
  )
})

onMounted(async () => {
  await fetchVolunteers()

  if (!selectedVolunteerId.value && allVolunteers.value.length > 0) {
    // Default select first (logic similar to before but simpler)
    selectedVolunteerId.value = allVolunteers.value[0].id
  }
})

const shifts = ref<IShift[]>([])

async function fetchShifts(volunteerId: string) {
  try {
    const res = await fetch(`/v1/volunteers/${volunteerId}/shifts`)
    if (res.ok) {
      const data = await res.json()
      shifts.value = data.data.shifts || []
    }
  } catch (err) {
    console.error('Failed to fetch shifts', err)
  }
}

// Watch selection to fetch shifts
import { watch } from 'vue'
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

function handleOpenCreate() {
  isCreating.value = true
}

async function handleCreateSave(newVolunteer: any) {
  try {
    const res = await fetch('/v1/volunteers', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newVolunteer),
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
  }
}

async function handleUpdateSave(updatedData: any) {
  if (!selectedVolunteerId.value) return

  try {
    const res = await fetch(`/v1/volunteers/${selectedVolunteerId.value}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedData),
    })

    if (res.ok) {
      await fetchVolunteers()
      // No need to reset selectedVolunteerId as it persists
    } else {
      const errText = await res.text()
      console.error('Update failed:', res.status, errText)
      alert(`Failed to update volunteer: ${errText}`)
    }
  } catch (e) {
    console.error('Error updating volunteer', e)
    alert('Error updating volunteer')
  }
}

async function handleAddShift(shiftData: any) {
  if (!selectedVolunteerId.value) return

  // Handle recurring logic or single shift
  // For MVP backed by API, let's start with single shift creation
  // If recurring is needed, we would loop and create multiple, or handle in backend.
  // The backend POST creates one shift.
  // We'll mimic the previous loop logic but call API for each.

  const shiftsToCreate = []
  const baseDate = new Date(shiftData.date)

  if (shiftData.isRecurring) {
    const endDate = shiftData.endDate
      ? new Date(shiftData.endDate)
      : new Date(baseDate.getTime() + 90 * 24 * 60 * 60 * 1000)

    let currentDate = new Date(baseDate)
    // Safety limit
    let count = 0
    while (currentDate <= endDate && count < 50) {
      shiftsToCreate.push({
        volunteerId: parseInt(selectedVolunteerId.value), // Ensure int
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
      volunteerId: parseInt(selectedVolunteerId.value),
      date: shiftData.date,
      startTime: shiftData.startTime,
      endTime: shiftData.endTime,
      role: shiftData.role,
    })
  }

  // Execute requests
  try {
    for (const s of shiftsToCreate) {
      await fetch('/v1/shifts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(s),
      })
    }
    // Refresh shifts
    await fetchShifts(selectedVolunteerId.value)
    // Refresh volunteers to get updated stats (reliability, etc)
    await fetchVolunteers()
  } catch (e) {
    console.error('Error creating shifts', e)
    alert('Failed to create shifts')
  }
}

async function handleUpdateShift(shiftData: any) {
  if (!shiftData.id) return
  try {
    const res = await fetch(`/v1/shifts/${shiftData.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(shiftData),
    })

    if (res.ok) {
      await fetchShifts(selectedVolunteerId.value!)
      await fetchVolunteers() // Stats update
    } else {
      const err = await res.text()
      alert('Failed to update shift: ' + err)
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
    <!-- Sidebar -->
    <aside class="sidebar">
      <VolunteerList
        :volunteers="allVolunteers"
        :loading="isLoading"
        :selectedId="String(selectedVolunteerId)"
        @select="(vol) => (selectedVolunteerId = vol.id)"
        @add="handleOpenCreate"
      />
    </aside>

    <!-- Create Modal/Drawer -->
    <VolunteerEditor
      :volunteer="null"
      :isOpen="isCreating"
      @close="isCreating = false"
      @save="handleCreateSave"
    />

    <!-- Main Content -->
    <main class="main-content">
      <VolunteerDetail
        v-if="selectedVolunteer"
        :volunteer="selectedVolunteer"
        :shifts="selectedShifts"
        :incidents="selectedIncidents"
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

.sidebar {
  width: 320px;
  background: var(--text-inverse);
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
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
  background-color: var(--color-secondary); /* Assuming purple mapped to secondary */
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
