<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockDevices, type IDisplayDevice } from '../../stores/mockKennelDisplays'
import { mockPetsData } from '../../stores/mockPetData'
import { Capsules, Button, InputField } from '../../components/common/ui'

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

// Modal for Assigning Pet (Mock)
const showAssignModal = ref(false)
const assignPetId = ref('')

const openAssignModal = () => {
  assignPetId.value = ''
  showAssignModal.value = true
}

const confirmAssignment = () => {
  if (selectedDevice.value && assignPetId.value) {
    selectedDevice.value.assignedPetId = assignPetId.value
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
        <div class="preview-area">
          <h3>Screen Preview</h3>
          <div class="e-ink-frame">
            <div class="e-ink-screen" :class="selectedDevice.template">
              <template v-if="previewPet">
                <!-- STANDARD LAYOUT -->
                <div class="screen-content">
                  <div class="screen-header">
                    <span class="p-name">{{ previewPet.name }}</span>
                    <span class="p-id">{{ previewPet.breed }}</span>
                  </div>

                  <div class="screen-body">
                    <div class="info-row">
                      <span>AGE: {{ previewPet.age }}y</span>
                      <span>SEX: {{ previewPet.sex }}</span>
                      <span>KENNEL: {{ previewPet.kennelNo }}</span>
                    </div>
                    <div class="desc-blk">friendly, energetic, good with kids</div>

                    <div v-if="selectedDevice.template === 'medical'" class="alert-banner">
                      💊 MEDICAL MONITORING
                    </div>
                    <div v-if="selectedDevice.template === 'urgent'" class="alert-banner urgent">
                      🚨 URGENT ADOPTION NEEDED
                    </div>
                  </div>

                  <div class="screen-footer">
                    <div class="qr-box">QR</div>
                    <span>Scan to Adopt</span>
                  </div>
                </div>
              </template>
              <template v-else>
                <div class="empty-screen">
                  <span class="logo-txt">Paws & Claws</span>
                  <span>No Pet Assigned</span>
                </div>
              </template>
            </div>
          </div>
          <p class="helper-text">4.2" E-Ink Display (400x300)</p>
        </div>
      </div>

      <div v-else class="empty-selection">Select a device to view details and manage content.</div>
    </div>

    <!-- MOCK ASSIGN MODAL -->
    <div v-if="showAssignModal" class="modal-overlay" @click.self="showAssignModal = false">
      <div class="modal-card">
        <h3>Assign Pet</h3>
        <div class="select-list">
          <div
            v-for="pet in mockPets"
            :key="pet.id"
            class="select-item"
            @click="assignPetId = pet.id"
            :class="{ selected: assignPetId === pet.id }"
          >
            {{ pet.name }} ({{ pet.breed }})
          </div>
        </div>
        <div class="modal-actions">
          <Button title="Cancel" color="white" :onClick="() => (showAssignModal = false)" />
          <Button title="Confirm Link" color="black" :onClick="confirmAssignment" />
        </div>
      </div>
    </div>
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

/* E-INK PREVIEW */
.preview-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  h3 {
    margin-top: 0;
    font-size: 1.1rem;
    color: var(--text-primary);
  }
}

.e-ink-frame {
  background: #e2e8f0;
  padding: 12px;
  border-radius: 8px;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 8px;
}

.e-ink-screen {
  width: 400px;
  height: 300px;
  background: #f4f4f4; /* Paper-ish white */
  border: 1px solid #d4d4d4;
  color: black;
  font-family: 'Georgia', serif; /* E-ink often uses serifs or simple sans */
  display: flex;
  flex-direction: column;
  padding: 16px;
  position: relative;

  /* Simulate E-ink update flash/ghosting implicitly via simple aesthetics */
}

.screen-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.screen-header {
  border-bottom: 2px solid black;
  margin-bottom: 12px;
  padding-bottom: 8px;
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}
.p-name {
  font-size: 2rem;
  font-weight: 900;
}
.p-id {
  font-size: 1.2rem;
  font-style: italic;
}

.info-row {
  font-family: sans-serif;
  font-weight: 700;
  margin-bottom: 16px;
  display: flex;
  gap: 16px;
  font-size: 0.9rem;
}

.desc-blk {
  font-size: 1.1rem;
  line-height: 1.4;
  flex: 1;
}

.screen-footer {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: auto;
}
.qr-box {
  width: 40px;
  height: 40px;
  background: black;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.7rem;
}

.alert-banner {
  background: black;
  color: white;
  text-align: center;
  font-weight: 700;
  padding: 4px;
  margin-top: 12px;
  &.urgent {
    background: black;
    border: 2px dashed white;
  }
}

.empty-screen {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
  .logo-txt {
    font-weight: 900;
    font-size: 1.5rem;
  }
}

.helper-text {
  font-size: 0.85rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.empty-selection {
  padding: 40px;
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
  background: white;
  border-radius: 12px;
  border: 1px dashed #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

/* MODAL */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}
.modal-card {
  background: white;
  padding: 24px;
  border-radius: 12px;
  width: 400px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
}
.select-list {
  overflow-y: auto;
  max-height: 300px;
  margin: 16px 0;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
}
.select-item {
  padding: 12px;
  cursor: pointer;
  &:hover {
    background: #f8fafc;
  }
  &.selected {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
    font-weight: 600;
  }
}
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
