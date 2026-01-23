<script setup lang="ts">
import { computed,ref } from 'vue'

import DeviceSettings from '../../components/admin/events/DeviceSettings.vue'
import DisplayPreview from '../../components/admin/events/DisplayPreview.vue'
import { mockEventDisplays } from '../../stores/mockEventDisplays'
import { mockPetsData } from '../../stores/mockPetData'

const selectedId = ref<string | null>(null)
const selectedDevice = computed(() =>
  mockEventDisplays.value.find((d) => d.id === selectedId.value),
)

const allPets = ref(mockPetsData)
const availablePets = computed(() => allPets.value.filter((p) => p.details.status === 'available'))

const selectDevice = (id: string) => {
  selectedId.value = id
}

const togglePetSelection = (petId: string) => {
  if (!selectedDevice.value) return
  const ids = selectedDevice.value.config.featuredPetIds
  if (ids.includes(petId)) {
    selectedDevice.value.config.featuredPetIds = ids.filter((id) => id !== petId)
  } else {
    if (ids.length < 4) {
      selectedDevice.value.config.featuredPetIds.push(petId)
    } else {
      alert('Max 4 pets for this template')
    }
  }
}

const getPetById = (id: string) => allPets.value.find((p) => p.id === id)

const forceUpdate = () => {
  if (selectedDevice.value) {
    selectedDevice.value.status = 'updating'
    setTimeout(() => {
      if (selectedDevice.value) selectedDevice.value.status = 'online'
    }, 2000)
  }
}

const templateOptions = [
  { label: 'Welcome Sign', value: 'welcome-sign' },
  { label: 'Featured Pets Grid', value: 'featured-grid' },
  { label: 'Wayfinding Arrow', value: 'wayfinding' },
  { label: 'Kennel Card', value: 'kennel-card' },
]

const directionOptions = [
  { label: '⬅ Left', value: 'left' },
  { label: '⬆ Straight', value: 'straight' },
  { label: '➡ Right', value: 'right' },
]
</script>

<template>
  <div class="event-displays-page">
    <div class="page-header">
      <h1>Event Displays</h1>
    </div>

    <div class="main-layout">
      <div class="sidebar">
        <h2>Devices</h2>
        <div class="device-list">
          <div
            v-for="device in mockEventDisplays"
            :key="device.id"
            class="device-card"
            :class="{ active: selectedId === device.id }"
            @click="selectDevice(device.id)"
          >
            <div class="d-icon">🖥️</div>
            <div class="d-info">
              <span class="d-name">{{ device.name }}</span>
              <span class="d-meta" :class="device.status">
                ● {{ device.status }} • {{ device.batteryLevel }}%
              </span>
            </div>
          </div>
        </div>
      </div>

      <div v-if="selectedDevice" class="editor-area">
        <DeviceSettings
          v-model="selectedDevice"
          :availablePets="availablePets"
          :templateOptions="templateOptions"
          :directionOptions="directionOptions"
          @forceUpdate="forceUpdate"
          @togglePet="togglePetSelection"
        />

        <DisplayPreview :device="selectedDevice" :getPetById="getPetById" />
      </div>

      <div v-else class="empty-state">Select a display to manage content.</div>
    </div>
  </div>
</template>

<style scoped>
.event-displays-page {
  height: calc(100vh - 100px);
  display: flex;
  flex-direction: column;
}

.page-header {
  h1 {
    margin: 0 0 24px 0;
    color: var(--text-primary);
    font-size: 1.8rem;
  }
}

.main-layout {
  display: flex;
  gap: 24px;
  flex: 1;
  min-height: 0;
}

.sidebar {
  width: 300px;
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;

  h2 {
    padding: 16px;
    margin: 0;
    font-size: 1.1rem;
    border-bottom: 1px solid #f3f4f6;
  }

  .device-list {
    overflow-y: auto;
    flex: 1;
  }

  .device-card {
    padding: 16px;
    border-bottom: 1px solid #f3f4f6;
    display: flex;
    gap: 12px;
    cursor: pointer;

    &:hover {
      background: #f8fafc;
    }

    &.active {
      background: #f1f5f9;
      border-left: 4px solid var(--color-secondary);
    }

    .d-icon {
      font-size: 1.5rem;
    }

    .d-info {
      display: flex;
      flex-direction: column;

      .d-name {
        font-weight: 600;
        color: var(--text-primary);
      }

      .d-meta {
        font-size: 0.8rem;
        text-transform: capitalize;

        &.online {
          color: #10b981;
        }

        &.offline {
          color: #ef4444;
        }

        &.updating {
          color: #f59e0b;
        }
      }
    }
  }
}

.editor-area {
  flex: 1;
  display: grid;
  grid-template-columns: 350px 1fr;
  gap: 24px;
  min-height: 0;
}
/* STYLES MOVED TO COMPONENTS */
</style>
