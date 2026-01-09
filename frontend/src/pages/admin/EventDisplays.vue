<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockEventDisplays } from '../../stores/mockEventDisplays'
import { mockPetsData } from '../../stores/mockPetData'
import { Button, InputField, InputSelectGroup } from '../../components/common/ui'

const selectedId = ref<string | null>(null)
const selectedDevice = computed(() =>
  mockEventDisplays.value.find((d) => d.id === selectedId.value),
)

const allPets = ref(mockPetsData)

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
        <div class="settings-panel">
          <div class="panel-header">
            <h3>Configuration</h3>
            <Button
              title="Push Update"
              size="small"
              color="black"
              :loading="selectedDevice.status === 'updating'"
              @click="forceUpdate"
            />
          </div>

          <div class="form-group">
            <InputSelectGroup
              label="Template"
              v-model="selectedDevice.template"
              :options="templateOptions"
            />
          </div>

          <div class="form-group">
            <label>Main Title</label>
            <InputField v-model="selectedDevice.config.title" />
          </div>

          <div class="form-group">
            <label>Subtitle / Message</label>
            <InputField v-model="selectedDevice.config.subtitle" />
          </div>

          <div v-if="selectedDevice.template === 'featured-grid'" class="form-group">
            <label>Select Pets (Max 4)</label>
            <div class="pet-picker">
              <div
                v-for="pet in allPets"
                :key="pet.id"
                class="pet-chip"
                :class="{ selected: selectedDevice.config.featuredPetIds.includes(pet.id) }"
                @click="togglePetSelection(pet.id)"
              >
                {{ pet.name }}
              </div>
            </div>
          </div>

          <div v-if="selectedDevice.template === 'wayfinding'" class="form-group">
            <InputSelectGroup
              label="Direction"
              v-model="selectedDevice.config.direction"
              :options="directionOptions"
            />
          </div>
        </div>

        <div class="preview-panel">
          <p class="preview-lbl">Live Preview (13.3" E-Ink)</p>
          <div class="e-ink-canvas">
            <div v-if="selectedDevice.template === 'welcome-sign'" class="layout-welcome">
              <div class="logo-area">🐾 Paws & Claws</div>
              <h1 class="big-title">{{ selectedDevice.config.title }}</h1>
              <p class="sub-text">{{ selectedDevice.config.subtitle }}</p>
              <div class="bottom-deco">
                <span>Scan for Available Pets</span>
                <div class="qr-mock">QR</div>
              </div>
            </div>

            <div v-if="selectedDevice.template === 'featured-grid'" class="layout-grid">
              <h2 class="grid-title">{{ selectedDevice.config.title }}</h2>
              <div class="pets-display-grid">
                <div
                  v-for="pid in selectedDevice.config.featuredPetIds"
                  :key="pid"
                  class="pet-tile"
                >
                  <div class="pt-photo">📷</div>
                  <div class="pt-name">{{ getPetById(pid)?.name }}</div>
                  <div class="pt-breed">{{ getPetById(pid)?.breed }}</div>
                </div>
                <div v-if="selectedDevice.config.featuredPetIds.length === 0" class="empty-msg">
                  Select pets to display
                </div>
              </div>
            </div>

            <div v-if="selectedDevice.template === 'wayfinding'" class="layout-wayfinding">
              <div class="arrow-container">
                <span v-if="selectedDevice.config.direction === 'left'" class="giant-arrow"
                  >⬅</span
                >
                <span v-if="selectedDevice.config.direction === 'straight'" class="giant-arrow"
                  >⬆</span
                >
                <span v-if="selectedDevice.config.direction === 'right'" class="giant-arrow"
                  >➡</span
                >
              </div>
              <h1 class="way-title">{{ selectedDevice.config.title }}</h1>
              <p class="way-sub">{{ selectedDevice.config.subtitle }}</p>
            </div>
          </div>
        </div>
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

  .settings-panel {
    background: white;
    border-radius: 12px;
    border: 1px solid #e2e8f0;
    padding: 24px;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: 20px;

    .panel-header {
      display: flex;
      justify-content: space-between;
      align-items: center;

      h3 {
        margin: 0;
      }
    }

    .form-group {
      display: flex;
      flex-direction: column;
      gap: 8px;

      label {
        font-size: 0.85rem;
        font-weight: 600;
      label {
        font-size: 0.85rem;
        font-weight: 600;
        color: var(--text-secondary);
      }
      }
    }

    .pet-picker {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      max-height: 200px;
      overflow-y: auto;
      padding: 4px;

      .pet-chip {
        padding: 6px 12px;
        background: #f1f5f9;
        border-radius: 20px;
        font-size: 0.9rem;
        cursor: pointer;
        border: 1px solid transparent;

        &.selected {
          background: hsl(from var(--color-secondary) h s 95%);
          color: var(--color-secondary);
          border-color: var(--color-secondary);
          font-weight: 600;
        }
      }
    }
  }

  .preview-panel {
    background: #e2e8f0;
    border-radius: 12px;
    padding: 24px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    .preview-lbl {
      margin-bottom: 12px;
      color: #64748b;
      font-size: 0.9rem;
    }

    .e-ink-canvas {
      width: 600px;
      height: 400px;
      background: #f4f4f4;
      border: 8px solid #334155;
      border-radius: 4px;
      box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.2);
      position: relative;
      overflow: hidden;
      color: black;

      .layout-welcome {
        height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;

        p {
          margin: 0;
        }

        .logo-area {
          font-size: 1.2rem;
          font-weight: 900;
          margin-bottom: 20px;
          border-bottom: 2px solid black;
          padding-bottom: 8px;
        }

        .big-title {
          font-size: 3rem;
          margin: 0 0 16px 0;
          font-family: serif;
        }

        .sub-text {
          font-size: 1.4rem;
          font-style: italic;
        }

        .bottom-deco {
          margin-top: auto;
          padding-bottom: 20px;
          display: flex;
          align-items: center;
          gap: 12px;

          .qr-mock {
            width: 50px;
            height: 50px;
            background: black;
            color: white;
            display: flex;
            align-items: center;
            justify-content: center;
          }
        }
      }

      .layout-grid {
        height: 100%;
        padding: 20px;
        display: flex;
        flex-direction: column;

        .grid-title {
          text-align: center;
          margin: 0 0 20px 0;
          border-bottom: 2px solid black;
          padding-bottom: 10px;
        }

        .pets-display-grid {
          display: grid;
          grid-template-columns: 1fr 1fr;
          gap: 16px;
          flex: 1;

          .pet-tile {
            border: 1px solid #ccc;
            padding: 10px;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            text-align: center;

            .pt-photo {
              font-size: 2rem;
              margin-bottom: 8px;
            }

            .pt-name {
              font-weight: 800;
              font-size: 1.2rem;
            }
          }
        }

        .empty-msg {
          grid-column: span 2;
          display: flex;
          align-items: center;
          justify-content: center;
          color: #666;
          font-style: italic;
        }
      }

      .layout-wayfinding {
        height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;

        .giant-arrow {
          font-size: 8rem;
          line-height: 1;
          margin-bottom: 20px;
        }

        .way-title {
          font-size: 3.5rem;
          margin: 0 0 10px 0;
          font-weight: 900;
        }

        .way-sub {
          font-size: 1.5rem;
          color: #444;
        }
      }
    }
  }
}

.empty-state {
  flex: 1;
  background: #f8fafc;
  border-radius: 12px;
  border: 2px dashed #cbd5e1;
  display: flex;
  align-items: center;
  justify-content: center;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  font-size: 1.1rem;
}
  font-size: 1.1rem;
}
</style>
