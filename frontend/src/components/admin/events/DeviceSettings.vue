<script setup lang="ts">
import type { IEventDisplay } from '../../../stores/mockEventDisplays'
import { Button, InputField, InputSelectGroup } from '../../common/ui'

defineProps<{
  device: IEventDisplay
  availablePets: any[] // eslint-disable-line @typescript-eslint/no-explicit-any
  templateOptions: { label: string; value: string }[]
  directionOptions: { label: string; value: string }[]
}>()

defineEmits<{
  forceUpdate: []
  togglePet: [petId: string]
}>()
</script>

<template>
  <div class="settings-panel">
    <div class="panel-header">
      <h3>Configuration</h3>
      <Button
        title="Push Update"
        size="small"
        color="black"
        :loading="device.status === 'updating'"
        @click="$emit('forceUpdate')"
      />
    </div>

    <div class="form-group">
      <label>Device Name</label>
      <InputField v-model="device.name" />
    </div>

    <div class="form-group">
      <InputSelectGroup label="Template" v-model="device.template" :options="templateOptions" />
    </div>

    <div class="form-group">
      <label>Main Title</label>
      <InputField v-model="device.config.title" />
    </div>

    <div class="form-group">
      <label>Subtitle / Message</label>
      <InputField v-model="device.config.subtitle" />
    </div>

    <!-- Kennel Card Pet Selection (Single) -->
    <div v-if="device.template === 'kennel-card'" class="form-group">
      <label>Select Pet</label>
      <div class="pet-picker">
        <div
          v-for="pet in availablePets"
          :key="pet.id"
          class="pet-chip"
          :class="{ selected: device.config.featuredPetIds.includes(pet.id) }"
          @click="
            () => {
              // Emulating single select toggle
              const isSelected = device.config.featuredPetIds.includes(pet.id)
              device.config.featuredPetIds = isSelected ? [] : [pet.id]
            }
          "
        >
          {{ pet.name }}
        </div>
      </div>
    </div>

    <div v-if="device.template === 'featured-grid'" class="form-group">
      <label>Select Pets (Max 4)</label>
      <div class="pet-picker">
        <div
          v-for="pet in availablePets"
          :key="pet.id"
          class="pet-chip"
          :class="{ selected: device.config.featuredPetIds.includes(pet.id) }"
          @click="$emit('togglePet', pet.id)"
        >
          {{ pet.name }}
        </div>
      </div>
    </div>

    <div v-if="device.template === 'wayfinding'" class="form-group">
      <InputSelectGroup
        label="Direction"
        v-model="device.config.direction"
        :options="directionOptions"
      />
    </div>
  </div>
</template>

<style scoped>
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
      color: var(--text-secondary);
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
</style>
