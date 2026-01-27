<script setup lang="ts">
import { ref } from 'vue'

import type { ITrip } from '../../../stores/mockTransport'
import { Button, InputField } from '../../common/ui'

defineProps<{
  selectedTrip: ITrip | undefined
  loadingAction: string | null
}>()

const emit = defineEmits<{
  updateStatus: [status: ITrip['status']]
  reportIssue: [type: string]
  sendEta: [eta: string]
  updateVehicle: [info: string]
}>()

const vehicleInput = ref('')
const etaInput = ref('')

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

const sendEta = () => {
  if (etaInput.value) {
    emit('sendEta', etaInput.value)
    etaInput.value = ''
  }
}

const updateVehicle = () => {
  if (vehicleInput.value) {
    emit('updateVehicle', vehicleInput.value)
  }
}
</script>

<template>
  <div class="action-center" v-if="selectedTrip">
    <div class="panel-header">
      <h2>Action Center</h2>
      <span class="trip-id">Trip #{{ selectedTrip.id }}</span>
    </div>

    <div class="panel-scroll">
      
      <section class="action-section">
        <div class="section-label">🚗 My Vehicle (Visible to Adopters)</div>
        <div class="vehicle-input-row">
          <InputField v-model="vehicleInput" placeholder="e.g. Silver Toyota Camry" />
          <Button title="Save" size="small" color="white" :onClick="updateVehicle" />
        </div>
      </section>

      <section class="action-section">
        <div class="section-label">⚡ Quick Actions</div>

        <div v-if="selectedTrip.status === 'en_route_vet'" class="button-grid">
          <div class="eta-sender">
            <InputField v-model="etaInput" placeholder="ETA (e.g. 15 mins)" />
            <Button title="Send ETA" color="white" :onClick="sendEta" :disabled="!etaInput" />
          </div>
          <Button
            title="✅ Arrived at Vet Safely"
            color="green"
            :loading="loadingAction === 'at_vet'"
            :onClick="() => $emit('updateStatus', 'at_vet')"
          />
        </div>

        <div v-else-if="selectedTrip.status === 'at_vet'" class="button-grid">
          <p class="status-helper">Wait for vet to finish intake/procedures.</p>
          <Button
            title="👋 Leaving Vet (En Route)"
            color="white"
            :loading="loadingAction === 'en_route_shelter'"
            :onClick="() => $emit('updateStatus', 'en_route_shelter')"
          />
          <Button
            title="⏳ Pets Not Ready / Delayed"
            color="orange"
            :loading="loadingAction === 'delayed'"
            :onClick="() => $emit('updateStatus', 'delayed')"
          />
        </div>

        <div v-else-if="selectedTrip.status === 'en_route_shelter'" class="button-grid">
          <div class="eta-sender">
            <InputField v-model="etaInput" placeholder="ETA to Shelter" />
            <Button title="Send ETA" color="white" :onClick="sendEta" :disabled="!etaInput" />
          </div>
          <Button
            title="🏠 Arrived at Shelter"
            color="green"
            :loading="loadingAction === 'at_shelter'"
            :onClick="() => $emit('updateStatus', 'at_shelter')"
          />
        </div>

        <div v-else class="button-grid">
          <Button
            v-if="selectedTrip.status === 'pending'"
            title="✋ Sign Up for Shift"
            color="purple"
            :onClick="() => $emit('updateStatus', 'assigned')"
          />
          <Button
            v-else-if="selectedTrip.status === 'assigned'"
            title="🚀 Start Trip"
            color="white"
            :onClick="
              () =>
                $emit(
                  'updateStatus',
                  selectedTrip?.direction === 'to_vet' ? 'en_route_vet' : 'en_route_shelter',
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
              :onClick="() => $emit('reportIssue', 'accident')"
            />
            <Button
              title="🔧 Breakdown"
              color="white"
              style="color: #f97316; border-color: #ffedd5; flex: 1"
              :onClick="() => $emit('reportIssue', 'breakdown')"
            />
          </div>
          <div class="issue-row">
            <Button
              title="🐢 Traffic Delay"
              color="white"
              style="color: #eab308; border-color: #fef08a; flex: 1"
              :onClick="() => $emit('reportIssue', 'traffic')"
            />
            <Button
              title="🤢 Pet Sick/Anxious"
              color="white"
              style="color: #a855f7; border-color: #f3e8ff; flex: 1"
              :onClick="() => $emit('reportIssue', 'pet_issue')"
            />
          </div>
        </div>
      </section>

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
</template>

<style scoped>
.action-center {
  background: #fff;
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
    color: var(--color-neutral-text-soft);
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
    color: var(--color-neutral-text-soft);
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
  color: var(--color-neutral-text-soft);
  font-style: italic;
  text-align: center;
}

.issue-row {
  display: flex;
  gap: 12px;
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
  color: var(--text-primary);
  font-size: 1rem;
}

.pet-status-badge {
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;

  &.adopted {
    background: var(--color-primary-weak);
    color: var(--color-primary);
  }

  &.shelter {
    background: var(--color-secondary-weak);
    color: var(--color-secondary);
  }
}

.pet-reason {
  color: var(--color-neutral-text-soft);
  font-size: 0.85rem;
  font-weight: 500;
}

.pet-description {
  font-size: 0.9rem;
  color: var(--color-neutral-text-soft);
  line-height: 1.4;
  background: #fff;
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
  color: var(--text-primary);
}

.no-logs {
  color: var(--color-neutral-text-soft);
  font-style: italic;
}
</style>
