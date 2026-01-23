<script setup lang="ts">
import { computed,ref } from 'vue'

import ActionCenter from '../../components/admin/transport/ActionCenter.vue'
import ShiftList from '../../components/admin/transport/ShiftList.vue'
import { type ITrip,mockTrips } from '../../stores/mockTransport'

const selectedTripId = ref<string | null>(null)
const loadingAction = ref<string | null>(null)

// Computed
const filteredTrips = computed(() => {
  // Sort by time
  return [...mockTrips.value].sort((a, b) => {
    // Put active statuses first
    const activeStatuses = ['en_route_vet', 'en_route_shelter']
    const aActive = activeStatuses.includes(a.status)
    const bActive = activeStatuses.includes(b.status)
    if (aActive && !bActive) return -1
    if (!aActive && bActive) return 1

    // Then sort by Date + Time
    const dateA = new Date(`${a.date} ${a.pickupTime}`).getTime()
    const dateB = new Date(`${b.date} ${b.pickupTime}`).getTime()
    return dateA - dateB
  })
})

const selectedTrip = computed(() => {
  return mockTrips.value.find((t) => t.id === selectedTripId.value)
})





// Actions
// Actions
const updateStatus = (newStatus: ITrip['status']) => {
  if (!selectedTrip.value) return
  loadingAction.value = newStatus

  // Simulate network delay
  setTimeout(() => {
    if (selectedTrip.value) selectedTrip.value.status = newStatus
    loadingAction.value = null
  }, 600)
}

const reportIssue = (issueType: string) => {
  if (!selectedTrip.value) return

  let message = ''
  let newStatus: ITrip['status'] = 'delayed'

  switch (issueType) {
    case 'accident':
      message = '🚨 URGENT: Was in a car accident. Please send assistance.'
      newStatus = 'incident'
      break
    case 'breakdown':
      message = '⚠️ Vehicle breakdown. Stopped safely.'
      newStatus = 'incident'
      break
    case 'traffic':
      message = '🐢 Stuck in heavy traffic. Will be delayed.'
      newStatus = 'delayed'
      break
    case 'pet_issue':
      message = '🤢 Pet is having an issue (anxiety/sickness) in transit.'
      newStatus = 'incident' // or keep current status but log warning? Let's say incident for visibility
      break
  }

  // Push Urgent Message
  selectedTrip.value.messages.push({
    id: `m-${Date.now()}`,
    sender: 'driver',
    text: message,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
  })

  // Update Status
  updateStatus(newStatus)
  alert('Incident reported! Dispatch has been notified.')
}

const selectTrip = (trip: ITrip) => {
  selectedTripId.value = trip.id
  // We'll let the ActionCenter component handle initializing its own local state for vehicle/eta
  // via props/emits if needed, or it can manage it locally since it's transient form data.
  // Actually, for vehicle info, if it persists, we might want to pass it down.
  // The original code initialized vehicleInput from trip.driverNotes or currentDriver.value?.vehicle
  // The ActionCenter component can read this from the `selectedTrip` prop.
}

const sendEta = (eta: string) => {
  if (!selectedTrip.value) return

  // Mock sending Notification
  selectedTrip.value.messages.push({
    id: `m-${Date.now()}`,
    sender: 'driver',
    text: `📢 ETA Update: I'll be there in ${eta}`,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
  })

  alert('ETA notification sent to Director and Adopters!')
}





const updateVehicleInfo = (info: string) => {
  if (!selectedTrip.value) return
  selectedTrip.value.driverNotes = info
  alert('Vehicle info updated for adopters!')
}
</script>

<template>
  <div class="transport-page">
    <div class="page-header">
      <h1>Transport & Dispatch</h1>
      <div class="header-actions">
        <span class="user-role-badge">Volunteer View</span>
      </div>
    </div>

    <div class="content-grid">
      <!-- Left: My Shifts / Available Shifts -->
      <ShiftList
        :trips="filteredTrips"
        :selected-trip-id="selectedTripId"
        @select="selectTrip"
      />

      <!-- Right: Action Center -->
      <ActionCenter
        :selected-trip="selectedTrip"
        :loading-action="loadingAction"
        @update-status="updateStatus"
        @report-issue="reportIssue"
        @send-eta="sendEta"
        @update-vehicle="updateVehicleInfo"
      />
    </div>
  </div>
</template>

<style scoped>
.transport-page {
  display: flex;
  flex-direction: column;
  height: 100%;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  h1 {
    margin: 0;
    font-size: 1.8rem;
    color: var(--text-primary);
  }
}

.user-role-badge {
  background: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
  padding: 6px 12px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 0.85rem;
}

.content-grid {
  display: grid;
  grid-template-columns: 350px 1fr;
  gap: 24px;
  flex: 1;
  min-height: 0;
}
</style>
