<script setup lang="ts">
import type { IEventDisplay } from '../../../stores/mockEventDisplays'
import { calculateAge,formatDate } from '../../../utils/date'

defineProps<{
  device: IEventDisplay
  getPetById: (_id: string) => unknown // eslint-disable-line no-unused-vars
}>()
</script>

<template>
  <div class="preview-panel">
    <p class="preview-lbl">Live Preview ({{ device.model }})</p>
    <div class="e-ink-canvas">
      <div v-if="device.template === 'welcome-sign'" class="layout-welcome">
        <div class="logo-area">🐾 Paws & Claws</div>
        <h1 class="big-title">{{ device.config.title }}</h1>
        <p class="sub-text">{{ device.config.subtitle }}</p>
        <div class="bottom-deco">
          <span>Scan for Available Pets</span>
          <div class="qr-mock">QR</div>
        </div>
      </div>

      <div v-if="device.template === 'featured-grid'" class="layout-grid">
        <h2 class="grid-title">{{ device.config.title }}</h2>
        <div class="pets-display-grid">
          <div v-for="pid in device.config.featuredPetIds" :key="pid" class="pet-tile">
            <div class="pt-photo">📷</div>
            <div class="pt-name">{{ getPetById(pid)?.name }}</div>
            <div class="pt-breed">{{ getPetById(pid)?.breed }}</div>
          </div>
          <div v-if="device.config.featuredPetIds.length === 0" class="empty-msg">
            Select pets to display
          </div>
        </div>
      </div>

      <div v-if="device.template === 'wayfinding'" class="layout-wayfinding">
        <div class="arrow-container">
          <span v-if="device.config.direction === 'left'" class="giant-arrow">⬅</span>
          <span v-if="device.config.direction === 'straight'" class="giant-arrow">⬆</span>
          <span v-if="device.config.direction === 'right'" class="giant-arrow">➡</span>
        </div>
        <h1 class="way-title">{{ device.config.title }}</h1>
        <p class="way-sub">{{ device.config.subtitle }}</p>
      </div>

      <div v-if="device.template === 'kennel-card'" class="layout-kennel-card">
        <div v-if="device.config.featuredPetIds.length" class="kc-container">
          <!-- We assume single pet for kennel card -->
          <div
            v-for="pid in device.config.featuredPetIds.slice(0, 1)"
            :key="pid"
            class="kc-content"
          >
            <div class="kc-left">
              <h1 class="kc-name">{{ getPetById(pid)?.name }}</h1>
              <div class="kc-stats">
                <div>Age: {{ calculateAge(getPetById(pid)?.physical.dateOfBirth) }}</div>
                <div>DOB: {{ formatDate(getPetById(pid)?.physical.dateOfBirth) }}</div>
                <div>Sex: {{ getPetById(pid)?.sex }}</div>
              </div>
              <div class="kc-breed">
                {{ getPetById(pid)?.physical.color || getPetById(pid)?.species }}
              </div>

              <div class="kc-personality">
                <div class="kc-lbl">Personality:</div>
                <div class="kc-tags-text">
                  {{
                    getPetById(pid)?.behavior?.personalityTags?.join(', ') || 'Playful, Loving'
                  }}
                </div>
              </div>
            </div>
            <div class="kc-right">
              <div class="kc-qr">
                <img
                  :src="`https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=https://adoption-os.com/adopt/${pid}`"
                  alt="QR"
                  style="width: 100%; height: 100%; image-rendering: pixelated"
                />
              </div>
              <div class="kc-cta">Adopt Me!</div>
            </div>
          </div>
        </div>
        <div v-else class="empty-msg" style="color: black; padding-top: 40px">
          Select a pet to view card
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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

    .layout-kennel-card {
      height: 100%;
      background: white;
      color: black;
      font-family: 'Courier New', Courier, monospace; /* Monospace for that printed label look */
      padding: 16px;
      box-sizing: border-box;

      .kc-container,
      .kc-content {
        height: 100%;
      }

      .kc-content {
        display: flex;
        justify-content: space-between;
      }

      .kc-left {
        flex: 1;
        display: flex;
        flex-direction: column;
        gap: 12px;
      }

      .kc-name {
        font-size: 3rem;
        font-weight: 800;
        margin: 0;
        line-height: 1;
        letter-spacing: -1px;
      }

      .kc-stats {
        font-size: 1.2rem;
        font-weight: 600;
        display: flex;
        flex-direction: column;
        gap: 4px;
      }

      .kc-breed {
        font-size: 1.1rem;
        margin-top: 8px;
      }

      .kc-personality {
        margin-top: auto;
        margin-bottom: 8px;

        .kc-lbl {
          font-weight: 800;
          font-size: 1.1rem;
          margin-bottom: 4px;
        }
        .kc-tags-text {
          font-size: 1rem;
          line-height: 1.3;
        }
      }

      .kc-right {
        width: 140px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        gap: 8px;
      }

      .kc-qr {
        width: 130px;
        height: 130px;
        background: white;
        border: 2px solid black;
        padding: 4px;
      }

      .kc-cta {
        font-size: 1.2rem;
        font-weight: 800;
        text-align: center;
        font-family: 'Courier New', Courier, monospace;
      }
    }
  }
}
</style>
