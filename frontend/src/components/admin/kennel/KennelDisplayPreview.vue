<script setup lang="ts">
import type { IDisplayDevice } from '../../../stores/mockKennelDisplays'

defineProps<{
  selectedDevice: IDisplayDevice | null
  previewPet: any // eslint-disable-line @typescript-eslint/no-explicit-any
}>()
</script>

<template>
  <div class="preview-area" v-if="selectedDevice">
    <h3>Screen Preview</h3>
    <div class="e-ink-frame">
      <div class="e-ink-screen" :class="selectedDevice.template">
        <template v-if="previewPet">
          
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
</template>

<style scoped>
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
  background: #f4f4f4; 
  border: 1px solid #d4d4d4;
  color: black;
  font-family: 'Georgia', serif; 
  display: flex;
  flex-direction: column;
  padding: 16px;
  position: relative;
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
</style>
