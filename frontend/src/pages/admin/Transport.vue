<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockTrips, mockDrivers, type ITrip, type IDriver } from '../../stores/mockTransport'
import { Capsules, InputField } from '../../components/common/ui'
import Button from '../../components/common/ui/Button.vue'

const selectedTripId = ref<string | null>(null)
const etaInput = ref('')
const vehicleInput = ref('')
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

const currentDriver = computed(() => {
  if (!selectedTrip.value?.driverId) return null
  return mockDrivers.value.find((d) => d.id === selectedTrip.value?.driverId)
})

const tripDirectionLabel = (direction: ITrip['direction']) => {
  return direction === 'to_vet' ? 'Shelter ➝ Vet (Drop-off)' : 'Vet ➝ Shelter (Pick-up)'
}

const getStatusColor = (status: ITrip['status']) => {
  if (status === 'incident') return '#fee2e2' // Red (Emergency)
  if (status.includes('en_route')) return '#bfdbfe' // Blue (Moving)
  if (status.includes('at_')) return '#fde68a' // Yellow (Waiting/Location)
  if (status === 'completed') return '#d1fae5' // Green
  if (status === 'delayed') return '#ffedd5' // Orange
  return '#f3f4f6' // Gray
}

const getStatusLabel = (status: ITrip['status']) => {
  switch (status) {
    case 'en_route_vet':
      return 'En Route to Vet'
    case 'at_vet':
      return 'At Vet'
    case 'en_route_shelter':
      return 'En Route to Shelter'
    case 'at_shelter':
      return 'Back at Shelter'
    case 'pending':
      return 'Needs Driver'
    case 'incident':
      return 'Incident Reported'
    default:
      return status.replace('_', ' ')
  }
}

// Actions
const reportIssue = (issueType: 'accident' | 'breakdown' | 'traffic' | 'pet_issue') => {
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
  vehicleInput.value = trip.driverNotes || currentDriver.value?.vehicle || ''
  etaInput.value = '' // Reset ETA on fresh selection
}

const sendEta = () => {
  if (!selectedTrip.value || !etaInput.value) return

  // Mock sending Notification
  selectedTrip.value.messages.push({
    id: `m-${Date.now()}`,
    sender: 'driver',
    text: `📢 ETA Update: I'll be there in ${etaInput.value}`,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
  })

  etaInput.value = ''
  alert('ETA notification sent to Director and Adopters!')
}

const updateStatus = (newStatus: ITrip['status']) => {
  if (!selectedTrip.value) return
  loadingAction.value = newStatus

  // Simulate network delay
  setTimeout(() => {
    if (selectedTrip.value) selectedTrip.value.status = newStatus
    loadingAction.value = null
  }, 600)
}

const updateVehicleInfo = () => {
  if (!selectedTrip.value) return
  selectedTrip.value.driverNotes = vehicleInput.value
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
      <div class="shifts-list">
        <h3>Upcoming Shifts</h3>
        <div
          v-for="trip in filteredTrips"
          :key="trip.id"
          class="trip-card"
          :class="{
            selected: selectedTripId === trip.id,
            'active-trip': trip.status.includes('en_route'),
            'urgent-trip': trip.status === 'incident',
          }"
          @click="selectTrip(trip)"
        >
          <div class="card-top">
            <span class="direction-badge" :class="trip.direction">
              {{ tripDirectionLabel(trip.direction) }}
            </span>
            <Capsules
              :label="getStatusLabel(trip.status)"
              :color="getStatusColor(trip.status)"
              size="sm"
            />
          </div>

          <div class="trip-time">⏰ {{ trip.pickupTime }}</div>

          <div class="pet-avatars">
            <div v-for="pet in trip.pets.slice(0, 4)" :key="pet.id" class="pet-capsule">
              🐾 {{ pet.name }}
            </div>
            <div v-if="trip.pets.length > 4" class="more-pets">+{{ trip.pets.length - 4 }}</div>
          </div>
        </div>
      </div>

      <!-- Right: Action Center -->
      <div class="action-center" v-if="selectedTrip">
        <div class="panel-header">
          <h2>Action Center</h2>
          <span class="trip-id">Trip #{{ selectedTrip.id }}</span>
        </div>

        <div class="panel-scroll">
          <!-- 1. Vehicle Info for Adopters -->
          <section class="action-section">
            <div class="section-label">🚗 My Vehicle (Visible to Adopters)</div>
            <div class="vehicle-input-row">
              <InputField v-model="vehicleInput" placeholder="e.g. Silver Toyota Camry" />
              <Button title="Save" size="small" color="white" :onClick="updateVehicleInfo" />
            </div>
          </section>

          <!-- 2. Workflow Actions -->
          <section class="action-section">
            <div class="section-label">⚡ Quick Actions</div>

            <!-- Context: En Route to Vet -->
            <div v-if="selectedTrip.status === 'en_route_vet'" class="button-grid">
              <div class="eta-sender">
                <InputField v-model="etaInput" placeholder="ETA (e.g. 15 mins)" />
                <Button title="Send ETA" color="black" :onClick="sendEta" :disabled="!etaInput" />
              </div>
              <Button
                title="✅ Arrived at Vet Safely"
                color="green"
                :loading="loadingAction === 'at_vet'"
                :onClick="() => updateStatus('at_vet')"
              />
            </div>

            <!-- Context: At Vet -->
            <div v-else-if="selectedTrip.status === 'at_vet'" class="button-grid">
              <p class="status-helper">Wait for vet to finish intake/procedures.</p>
              <Button
                title="👋 Leaving Vet (En Route)"
                color="black"
                :loading="loadingAction === 'en_route_shelter'"
                :onClick="() => updateStatus('en_route_shelter')"
              />
              <Button
                title="⏳ Pets Not Ready / Delayed"
                color="orange"
                :loading="loadingAction === 'delayed'"
                :onClick="() => updateStatus('delayed')"
              />
            </div>

            <!-- Context: En Route to Shelter -->
            <div v-else-if="selectedTrip.status === 'en_route_shelter'" class="button-grid">
              <div class="eta-sender">
                <InputField v-model="etaInput" placeholder="ETA to Shelter" />
                <Button title="Send ETA" color="black" :onClick="sendEta" :disabled="!etaInput" />
              </div>
              <Button
                title="🏠 Arrived at Shelter"
                color="green"
                :loading="loadingAction === 'at_shelter'"
                :onClick="() => updateStatus('at_shelter')"
              />
            </div>

            <!-- Default / Pending -->
            <div v-else class="button-grid">
              <Button
                v-if="selectedTrip.status === 'pending'"
                title="✋ Sign Up for Shift"
                color="purple"
                :onClick="() => updateStatus('assigned')"
              />
              <Button
                v-else-if="selectedTrip.status === 'assigned'"
                title="🚀 Start Trip"
                color="black"
                :onClick="
                  () =>
                    updateStatus(
                      selectedTrip.direction === 'to_vet' ? 'en_route_vet' : 'en_route_shelter',
                    )
                "
              />
              <p v-else class="status-helper">Trip is {{ getStatusLabel(selectedTrip.status) }}</p>
            </div>
          </section>

          <section class="action-section">
            <div class="section-label">⚠️ Report Issue</div>
            <div class="button-grid">
              <div class="issue-row">
                <Button
                  title="💥 Accident"
                  color="white"
                  style="color: #ef4444; border-color: #fee2e2; flex: 1"
                  :onClick="() => reportIssue('accident')"
                />
                <Button
                  title="🔧 Breakdown"
                  color="white"
                  style="color: #f97316; border-color: #ffedd5; flex: 1"
                  :onClick="() => reportIssue('breakdown')"
                />
              </div>
              <div class="issue-row">
                <Button
                  title="🐢 Traffic Delay"
                  color="white"
                  style="color: #eab308; border-color: #fef08a; flex: 1"
                  :onClick="() => reportIssue('traffic')"
                />
                <Button
                  title="🤢 Pet Sick/Anxious"
                  color="white"
                  style="color: #a855f7; border-color: #f3e8ff; flex: 1"
                  :onClick="() => reportIssue('pet_issue')"
                />
              </div>
            </div>
          </section>

          <!-- 3. Manifest -->
          <section class="action-section">
            <div class="section-label">📋 Passenger Manifest</div>
            <div class="passenger-list">
              <div v-for="pet in selectedTrip.pets" :key="pet.id" class="passenger-row">
                <div class="pet-header">
                  <div class="pet-name-wrap">
                    <span class="pet-name">🐾 {{ pet.name }}</span>
                    <span class="pet-status-badge" :class="pet.status">{{
                      pet.status === 'adopted' ? '🏠 Adopted' : '🏢 Shelter'
                    }}</span>
                  </div>
                  <span class="pet-reason">{{ pet.reason }}</span>
                </div>
                <div class="pet-description">
                  {{ pet.description }}
                </div>
              </div>
            </div>
          </section>

          <!-- 4. Communication -->
          <section class="action-section">
            <div class="section-label">💬 Updates Log</div>
            <div class="log-entries">
              <div v-for="msg in selectedTrip.messages" :key="msg.id" class="log-entry">
                <strong>{{ msg.timestamp }}:</strong> {{ msg.text }}
              </div>
              <div v-if="selectedTrip.messages.length === 0" class="no-logs">No updates yet.</div>
            </div>
          </section>
        </div>
      </div>

      <div class="action-center empty" v-else>
        <span class="icon">👈</span>
        <h3>Select a Shift</h3>
        <p>Choose a trip from the list to view actions.</p>
      </div>
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
    color: var(--font-color-dark);
  }
}

.user-role-badge {
  background: #f3e8ff;
  color: var(--purple);
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

.shifts-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  overflow-y: auto;
  padding-right: 4px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--font-color-medium);
  }
}

.trip-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
  border-left: 4px solid transparent;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    border-color: #d1d5db;
  }

  &.selected {
    border-color: var(--purple);
    background: #fdfbff;
  }

  &.active-trip {
    border-left-color: #10b981;
  }

  &.urgent-trip {
    border-left-color: #ef4444;
    background: #fef2f2;
  }
}

.card-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.direction-badge {
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;

  &.to_vet {
    color: #3b82f6;
  }
  &.from_vet {
    color: #8b5cf6;
  }
}

.trip-time {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--font-color-dark);
  margin-bottom: 12px;
}

.pet-avatars {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.pet-capsule {
  background: #f3f4f6;
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.85rem;
  color: var(--font-color-dark);
}

.more-pets {
  font-size: 0.8rem;
  color: var(--font-color-medium);
  align-self: center;
}

/* Action Center */
.action-center {
  background: white;
  border-radius: 16px;
  border: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  overflow: hidden;

  &.empty {
    align-items: center;
    justify-content: center;
    background: #f9fafb;
    border-style: dashed;
    .icon {
      font-size: 2rem;
      margin-bottom: 16px;
    }
  }
}

.panel-header {
  padding: 20px;
  border-bottom: 1px solid #f3f4f6;
  background: #fafafa;
  display: flex;
  justify-content: space-between;
  align-items: center;

  h2 {
    margin: 0;
    font-size: 1.2rem;
  }
  .trip-id {
    color: var(--font-color-medium);
    font-family: monospace;
  }
}

.panel-scroll {
  padding: 24px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.action-section {
  .section-label {
    font-size: 0.8rem;
    text-transform: uppercase;
    color: var(--font-color-medium);
    font-weight: 700;
    margin-bottom: 12px;
    letter-spacing: 0.05em;
  }
}

.vehicle-input-row {
  display: flex;
  gap: 12px;
  align-items: center;
}

.issue-row {
  display: flex;
  gap: 12px;
}

.button-grid {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.eta-sender {
  display: flex;
  gap: 12px;
  background: #f9fafb;
  padding: 12px;
  border-radius: 12px;
  border: 1px solid #f3f4f6;
}

.status-helper {
  color: var(--font-color-medium);
  font-style: italic;
  text-align: center;
}

.passenger-list {
  background: #f9fafb;
  border-radius: 12px;
  padding: 16px;
}

.passenger-row {
  display: flex;
  flex-direction: column;
  gap: 6px;
  padding: 12px 0;
  border-bottom: 1px solid #e5e7eb;

  &:last-child {
    border-bottom: none;
  }
}

.pet-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.pet-name-wrap {
  display: flex;
  align-items: center;
  gap: 8px;
}

.pet-name {
  font-weight: 600;
  color: var(--font-color-dark);
  font-size: 1rem;
}
.pet-reason {
  color: var(--font-color-medium);
  font-size: 0.85rem;
  font-weight: 500;
}

.pet-status-badge {
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;

  &.adopted {
    background: #dcfce7;
    color: #166534;
  }
  &.shelter {
    background: #e0e7ff;
    color: #3730a3;
  }
}

.pet-description {
  font-size: 0.9rem;
  color: var(--font-color-medium);
  line-height: 1.4;
  background: white;
  padding: 8px;
  border-radius: 6px;
  border: 1px dashed #e5e7eb;
}

.log-entries {
  background: #f9fafb;
  border-radius: 12px;
  padding: 16px;
  max-height: 200px;
  overflow-y: auto;
  font-size: 0.9rem;
}

.log-entry {
  margin-bottom: 8px;
  line-height: 1.4;
  color: var(--font-color-dark);
}
.no-logs {
  color: var(--font-color-medium);
  font-style: italic;
}
</style>
