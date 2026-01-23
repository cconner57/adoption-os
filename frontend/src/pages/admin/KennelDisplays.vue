<script setup lang="ts">
import { computed,ref } from 'vue'

import { Button,Capsules } from '../../components/common/ui'
import { mockDevices } from '../../stores/mockKennelDisplays'
import { mockPetsData } from '../../stores/mockPetData'

const mockPets = ref(mockPetsData)

// Helpers
const getPetName = (id: string | null) => {
  if (!id) return 'Unassigned'
  const pet = mockPets.value.find((p) => p.id === id)
  return pet ? pet.name : 'Unknown Pet'
}

// Battery Icon Logic
const getBatteryIcon = (level: number) => {
  if (level > 90) return '🔋'
  if (level > 50) return '🔋' // Simplification
  if (level > 20) return '🪫'
  return '🔌'
}

const getBatteryColor = (level: number) => {
  if (level <= 20) return '#ef4444' // Red
  return '#10b981' // Green
}

// Device Selection
const selectedDeviceId = ref<string | null>(null)
const selectedDevice = computed(() =>
  mockDevices.value.find((d) => d.id === selectedDeviceId.value),
)

// Linked pet for preview
const previewPet = computed(() => {
  if (!selectedDevice.value?.assignedPetId) return null
  return mockPets.value.find((p) => p.id === selectedDevice.value?.assignedPetId)
})

const selectDevice = (id: string) => {
  selectedDeviceId.value = id
}

// Stats
const onlineCount = computed(() => mockDevices.value.filter((d) => d.status === 'online').length)
const lowBatteryCount = computed(() => mockDevices.value.filter((d) => d.batteryLevel < 20).length)

// Actions
const refreshDisplay = () => {
  if (!selectedDevice.value) return
  selectedDevice.value.status = 'syncing'
  setTimeout(() => {
    if (selectedDevice.value) {
      selectedDevice.value.status = 'online'
      selectedDevice.value.lastSync = 'Just now'
    }
  }, 1500)
}

const unpairDevice = () => {
  if (!selectedDevice.value) return
  if (confirm(`Unpair device ${selectedDevice.value.macAddress}?`)) {
    selectedDevice.value.assignedPetId = null
  }
}

import KennelDisplayPreview from '../../components/admin/kennel/KennelDisplayPreview.vue'
import PetAssignmentModal from '../../components/admin/kennel/PetAssignmentModal.vue'

// ... (keep existing refs/computed)

// Modal for Assigning Pet (Mock)
const showAssignModal = ref(false)

const openAssignModal = () => {
  showAssignModal.value = true
}

const handleConfirmAssignment = (petId: string) => {
  if (selectedDevice.value) {
    selectedDevice.value.assignedPetId = petId
    showAssignModal.value = false
  }
}
</script>

<template>
  <div class="kennel-displays-page">
    <!-- HEADER -->
    <div class="page-header">
      <h1>Kennel Displays</h1>
      <div class="header-stats">
        <div class="stat-pill">
          <span class="dot green"></span>
          {{ onlineCount }} Online
        </div>
        <div class="stat-pill" :class="{ warning: lowBatteryCount > 0 }">
          <span class="dot red"></span>
          {{ lowBatteryCount }} Low Battery
        </div>
      </div>
    </div>

    <div class="split-layout">
      <!-- LEFT: DEVICE LIST -->
      <div class="device-list-card">
        <div class="list-header">
          <h2>Connected Tags</h2>
          <Button title="+ Add Device" size="small" color="white" />
        </div>

        <div class="list-content">
          <div
            v-for="device in mockDevices"
            :key="device.id"
            class="device-item"
            :class="{ active: selectedDeviceId === device.id }"
            @click="selectDevice(device.id)"
          >
            <div class="device-icon">
              <span class="e-ink-icon">📑</span>
            </div>

            <div class="device-info">
              <div class="top-line">
                <span class="pet-name-lbl">{{ getPetName(device.assignedPetId) }}</span>
                <Capsules
                  :label="device.status"
                  :color="device.status === 'online' ? '#d1fae5' : '#f3f4f6'"
                  size="sm"
                />
              </div>
              <div class="sub-line">
                <span class="mac">{{ device.macAddress }}</span>
              </div>
            </div>

            <div class="device-meta">
              <span :style="{ color: getBatteryColor(device.batteryLevel) }">
                {{ getBatteryIcon(device.batteryLevel) }} {{ device.batteryLevel }}%
              </span>
              <span class="signal">📶 {{ device.signalStrength }}%</span>
            </div>
          </div>
        </div>
      </div>

      <!-- RIGHT: DETAILS & PREVIEW -->
      <div class="details-panel" v-if="selectedDevice">
        <div class="panel-header">
          <h2>Device Details</h2>
          <div class="actions">
            <Button title="Refresh" size="small" color="white" :onClick="refreshDisplay" />
            <Button title="Unpair" size="small" color="white" :onClick="unpairDevice" />
          </div>
        </div>

        <div class="meta-grid">
          <div class="meta-item">
            <label>MAC Address</label>
            <span>{{ selectedDevice.macAddress }}</span>
          </div>
          <div class="meta-item">
            <label>Last Sync</label>
            <span>{{ selectedDevice.lastSync }}</span>
          </div>
          <div class="meta-item">
            <label>Assigned To</label>
            <div class="assign-row">
              <span class="assign-val">{{ getPetName(selectedDevice.assignedPetId) }}</span>
              <button class="change-link" @click="openAssignModal">Change</button>
            </div>
          </div>
          <div class="meta-item">
            <label>Template</label>
            <select v-model="selectedDevice.template" class="native-select">
              <option value="standard">Standard Profile</option>
              <option value="medical">Medical Alert</option>
              <option value="urgent">Urgent / Adopt Me</option>
            </select>
          </div>
        </div>

        <!-- E-INK PREVIEW -->
        <KennelDisplayPreview :selectedDevice="selectedDevice" :previewPet="previewPet" />
      </div>

      <div v-else class="empty-selection">Select a device to view details and manage content.</div>
    </div>

    <!-- MOCK ASSIGN MODAL -->
    <!-- MOCK ASSIGN MODAL -->
    <PetAssignmentModal
      :isOpen="showAssignModal"
      :mockPets="mockPets"
      @close="showAssignModal = false"
      @confirm="handleConfirmAssignment"
    />
  </div>
</template>

<style scoped>
.kennel-displays-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: calc(100vh - 100px); /* Fill remaining usage */
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

.header-stats {
  display: flex;
  gap: 12px;
}

.stat-pill {
  background: white;
  padding: 6px 12px;
  border-radius: 20px;
  border: 1px solid #e2e8f0;
  font-size: 0.9rem;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 8px;

  &.warning {
    background: hsl(from var(--color-danger) h s 95%);
    border-color: hsl(from var(--color-danger) h s 80%);
    color: var(--color-danger);
  }
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  &.green {
    background: var(--color-primary);
  }
  &.red {
    background: var(--color-danger);
  }
}

.split-layout {
  display: grid;
  grid-template-columns: 350px 1fr;
  gap: 24px;
  flex: 1;
  min-height: 0; /* Allow scrolling within items */
}

/* DEVICE LIST */
.device-list-card {
  background: white;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.list-header {
  padding: 16px;
  border-bottom: 1px solid #f3f4f6;
  display: flex;
  justify-content: space-between;
  align-items: center;
  h2 {
    margin: 0;
    font-size: 1.1rem;
  }
}

.list-content {
  flex: 1;
  overflow-y: auto;
}

.device-item {
  padding: 16px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  gap: 12px;
  cursor: pointer;
  transition: background 0.2s;

  &:hover {
    background: #f8fafc;
  }
  &.active {
    background: hsl(from var(--color-neutral) h s 95%);
    border-left: 4px solid var(--color-secondary);
  }
}

.device-icon {
  font-size: 1.5rem;
}

.device-info {
  flex: 1;
}
.top-line {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}
.pet-name-lbl {
  font-weight: 600;
  color: var(--text-primary);
}
.mac {
  font-family: monospace;
  color: hsl(from var(--color-neutral) h s 50%);
  font-size: 0.8rem;
}

.device-meta {
  font-size: 0.8rem;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 2px;
  color: hsl(from var(--color-neutral) h s 50%);
}

/* DETAILS PANEL */
.details-panel {
  background: white;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 24px;
  overflow-y: auto;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  h2 {
    margin: 0;
    font-size: 1.4rem;
  }
}

.actions {
  display: flex;
  gap: 8px;
}

.meta-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  background: #f8fafc;
  padding: 16px;
  border-radius: 8px;
}

.meta-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
  label {
    font-size: 0.75rem;
    text-transform: uppercase;
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 600;
  }
  span,
  .assign-val {
    font-weight: 500;
  }
}

.assign-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.change-link {
  background: none;
  border: none;
  color: var(--color-secondary);
  font-weight: 600;
  cursor: pointer;
  font-size: 0.85rem;
}

.native-select {
  padding: 6px;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  background: white;
  font-size: 0.9rem;
}

/* STYLES MOVED TO COMPONENTS */
</style>
