<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

import ActionCenter from '../../components/admin/transport/ActionCenter.vue'
import ShiftList from '../../components/admin/transport/ShiftList.vue'
import { type ITrip,mockTrips } from '../../stores/mockTransport'

const selectedTripId = ref<string | null>(null)
const loadingAction = ref<string | null>(null)
const isMounted = ref(false)

onMounted(() => {
  isMounted.value = true
})

const filteredTrips = computed(() => {

  return [...mockTrips.value].sort((a, b) => {

    const activeStatuses = new Set(['en_route_vet', 'en_route_shelter'])
    const aActive = activeStatuses.has(a.status)
    const bActive = activeStatuses.has(b.status)
    if (aActive && !bActive) return -1
    if (!aActive && bActive) return 1

    const dateA = new Date(`${a.date} ${a.pickupTime}`).getTime()
    const dateB = new Date(`${b.date} ${b.pickupTime}`).getTime()
    return dateA - dateB
  })
})

const selectedTrip = computed(() => {
  return mockTrips.value.find((t) => t.id === selectedTripId.value)
})

const updateStatus = (newStatus: ITrip['status']) => {
  if (!selectedTrip.value) return
  loadingAction.value = newStatus

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
      break
    case 'pet_issue':
      message = '🤢 Pet is having an issue (anxiety/sickness) in transit.'
      newStatus = 'incident'
      break
  }

  selectedTrip.value.messages.push({
    id: `m-${Date.now()}`,
    sender: 'driver',
    text: message,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
  })

  updateStatus(newStatus)
  alert('Incident reported! Dispatch has been notified.')
}

const selectTrip = (trip: ITrip) => {
  selectedTripId.value = trip.id

}

const sendEta = (eta: string) => {
  if (!selectedTrip.value) return

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
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Transport & Dispatch</h1>
    </Teleport>

    <div class="page-header">
      <h1 class="desktop-only">Transport & Dispatch</h1>
      <div class="header-actions">
        <span class="user-role-badge">Volunteer View</span>
      </div>
    </div>

    <div class="content-grid">

      <ShiftList
        :trips="filteredTrips"
        :selected-trip-id="selectedTripId"
        @select="selectTrip"
      />

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

  h1.desktop-only {
    margin: 0;
    font-size: 1.8rem;
    color: var(--text-primary);
  }
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .page-header h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.user-role-badge {
  background: var(--color-secondary-weak);
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
