<script setup lang="ts">
import { ref, watch } from 'vue'

import { Button,InputField } from '../../components/common/ui'
import { mockKioskSettings } from '../../stores/mockKiosk'

// Mock saving state
const isSaving = ref(false)
const hasChanges = ref(false)

const saveSettings = () => {
  isSaving.value = true
  setTimeout(() => {
    isSaving.value = false
    hasChanges.value = false
    alert('Kiosk settings published!')
  }, 800)
}

watch(
  mockKioskSettings,
  () => {
    hasChanges.value = true
  },
  { deep: true },
)
</script>

<template>
  <div class="kiosk-page">
    <!-- HEADER -->
    <div class="page-header">
      <h1>Kiosk Management</h1>
      <div class="header-actions">
        <span v-if="hasChanges" class="unsaved-badge">Unsaved Changes</span>
        <Button title="Publish Changes" color="black" :loading="isSaving" :onClick="saveSettings" />
      </div>
    </div>

    <div class="split-view">
      <!-- LEFT: SETTINGS -->
      <div class="settings-panel">
        <!-- SECTION: GENERAL -->
        <div class="settings-group">
          <h2>General</h2>
          <div class="form-stack">
            <label>Kiosk Name</label>
            <InputField v-model="mockKioskSettings.general.kioskName" placeholder="Enter Kiosk Name" />

            <label>Welcome Message</label>
            <InputField v-model="mockKioskSettings.general.welcomeMessage" placeholder="Welcome Message" />

            <label>Screen Timeout (Seconds)</label>
            <input
              type="range"
              v-model.number="mockKioskSettings.general.timeoutSeconds"
              min="30"
              max="300"
              step="30"
            />
            <div class="range-val">{{ mockKioskSettings.general.timeoutSeconds }}s</div>
          </div>
        </div>

        <!-- SECTION: BRANDING -->
        <div class="settings-group">
          <h2>Branding & Theme</h2>
          <div class="form-stack">
            <label>Primary Color</label>
            <div class="color-picker-row">
              <input
                type="color"
                v-model="mockKioskSettings.theme.primaryColor"
                class="native-color"
              />
              <span>{{ mockKioskSettings.theme.primaryColor }}</span>
            </div>

            <label>Background Image URL</label>
            <InputField v-model="mockKioskSettings.theme.backgroundImage" placeholder="https://..." />
          </div>
        </div>

        <!-- SECTION: FEATURES -->
        <div class="settings-group">
          <h2>Enabled Features</h2>
          <div class="toggle-list">
            <div class="toggle-row">
              <span>Volunteer Check-In</span>
              <input
                type="checkbox"
                v-model="mockKioskSettings.features.allowCheckIn"
                class="toggle-switch"
              />
            </div>
            <div class="toggle-row">
              <span>Browse Available Pets</span>
              <input
                type="checkbox"
                v-model="mockKioskSettings.features.showAvailablePets"
                class="toggle-switch"
              />
            </div>
            <div class="toggle-row">
              <span>Upcoming Events</span>
              <input
                type="checkbox"
                v-model="mockKioskSettings.features.showEvents"
                class="toggle-switch"
              />
            </div>
            <div class="toggle-row">
              <span>Accept Donations</span>
              <input
                type="checkbox"
                v-model="mockKioskSettings.features.donationsEnabled"
                class="toggle-switch"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- RIGHT: PREVIEW -->
      <div class="preview-panel">
        <div class="device-frame">
          <div class="camera-hole"></div>
          <div
            class="screen-content"
            :style="{
              backgroundImage: `linear-gradient(rgba(0,0,0,0.5), rgba(0,0,0,0.5)), url(${mockKioskSettings.theme.backgroundImage})`,
            }"
          >
            <div class="kiosk-ui">
              <h1 class="welcome-text">{{ mockKioskSettings.general.welcomeMessage }}</h1>

              <div class="action-grid">
                <button
                  v-if="mockKioskSettings.features.allowCheckIn"
                  class="kiosk-btn"
                  :style="{ backgroundColor: mockKioskSettings.theme.primaryColor }"
                >
                  👋 Check In
                </button>

                <button
                  v-if="mockKioskSettings.features.showAvailablePets"
                  class="kiosk-btn"
                  :style="{ backgroundColor: mockKioskSettings.theme.primaryColor }"
                >
                  🐾 View Pets
                </button>

                <button
                  v-if="mockKioskSettings.features.showEvents"
                  class="kiosk-btn"
                  :style="{ backgroundColor: mockKioskSettings.theme.primaryColor }"
                >
                  📅 Events
                </button>

                <button
                  v-if="mockKioskSettings.features.donationsEnabled"
                  class="kiosk-btn"
                  :style="{ backgroundColor: mockKioskSettings.theme.primaryColor }"
                >
                  ❤️ Donate
                </button>
              </div>
            </div>

            <div class="kiosk-footer">Tap to start • {{ mockKioskSettings.general.kioskName }}</div>
          </div>
        </div>
        <p class="preview-label">Live Preview (iPad Pro 12.9")</p>
        <button class="launch-link" @click="$router.push('/kiosk')">↗ Launch Kiosk Mode</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.kiosk-page {
  display: flex;
  flex-direction: column;
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

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.unsaved-badge {
  background: #fef08a; /* yellow-200 */
  color: #854d0e; /* yellow-800 */
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
}

.split-view {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
}

/* LEFT: SETTINGS */
.settings-panel {
  display: flex;
  flex-direction: column;
  gap: 24px;
  max-width: 500px;
}

.settings-group {
  background: white;
  padding: 24px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;

  h2 {
    margin-top: 0;
    font-size: 1.1rem;
    border-bottom: 1px solid #f3f4f6;
    padding-bottom: 12px;
    margin-bottom: 20px;
  }
}

.form-stack {
  display: flex;
  flex-direction: column;
  gap: 12px;
  label {
    font-weight: 500;
    font-size: 0.9rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.range-val {
  text-align: right;
  font-weight: 600;
  font-size: 0.9rem;
}

.color-picker-row {
  display: flex;
  align-items: center;
  gap: 12px;
  font-family: monospace;
}

.native-color {
  border: none;
  width: 40px;
  height: 40px;
  cursor: pointer;
  background: none;
}

.toggle-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.toggle-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  span {
    font-weight: 500;
  }
}

/* Custom checkbox as toggle */
.toggle-switch {
  appearance: none;
  width: 44px;
  height: 24px;
  background: #e5e7eb;
  border-radius: 12px;
  position: relative;
  cursor: pointer;
  transition: background 0.3s;

  &:checked {
    background: var(--color-secondary);
  }

  &::after {
    content: '';
    position: absolute;
    top: 2px;
    left: 2px;
    width: 20px;
    height: 20px;
    background: white;
    border-radius: 50%;
    transition: transform 0.3s;
  }

  &:checked::after {
    transform: translateX(20px);
  }
}

/* RIGHT: PREVIEW */
.preview-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding-top: 20px;
}

.device-frame {
  width: 100%;
  max-width: 600px;
  aspect-ratio: 4/3;
  background: #1e293b;
  border-radius: 24px;
  padding: 16px;
  position: relative;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}

.camera-hole {
  width: 6px;
  height: 6px;
  background: #334155;
  border-radius: 50%;
  position: absolute;
  left: 50%;
  top: 8px;
}

.screen-content {
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: white;
  padding: 40px;
  text-align: center;
}

.kiosk-ui {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 40px;
}

.welcome-text {
  font-size: 2rem;
  font-weight: 800;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
  margin: 0;
}

.action-grid {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  justify-content: center;
}

.kiosk-btn {
  padding: 16px 24px;
  border: none;
  border-radius: 12px;
  color: white;
  font-weight: 700;
  font-size: 1.1rem;
  cursor: pointer;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  gap: 8px;
  transition: transform 0.2s;

  &:hover {
    transform: translateY(-2px);
  }
}

.kiosk-footer {
  font-size: 0.8rem;
  opacity: 0.8;
}

.preview-label {
  margin-top: 16px;
  color: hsl(from var(--color-neutral) h s 50%);
  font-size: 0.9rem;
}

/* Responsive */
@media (max-width: 1024px) {
  .split-view {
    grid-template-columns: 1fr;
  }
  .settings-panel {
    max-width: 100%;
  }
}
</style>
