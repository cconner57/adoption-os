<script setup lang="ts">
import { computed } from 'vue'

import { ButtonToggle, InputField, InputTextArea,Select } from '../../common/ui'

const props = defineProps<{
  settings: Record<string, any> // eslint-disable-line @typescript-eslint/no-explicit-any
}>()

const emit = defineEmits(['update:settings'])

const localSettings = computed({
  get: () => props.settings,
  set: (val) => emit('update:settings', val),
})

const speciesOptions = [
  { label: 'Cat', value: 'cat' },
  { label: 'Dog', value: 'dog' },
]

const statusOptions = [
  { label: 'Available', value: 'available' },
  { label: 'Intake Processing', value: 'intake' },
  { label: 'Medical Hold', value: 'hold' },
  { label: 'Foster', value: 'foster' },
]

const environmentOptions = [
  { label: 'Indoor', value: 'indoor' },
  { label: 'Outdoor', value: 'outdoor' },
  { label: 'Indoor/Outdoor', value: 'indoor/outdoor' },
]

const unitOptions = [
  { label: 'Imperial (lbs)', value: 'imperial' },
  { label: 'Metric (kg)', value: 'metric' },
]

const catBreeds = [
  'Domestic Short Hair',
  'Domestic Medium Hair',
  'Domestic Long Hair',
  'Siamese',
  'Maine Coon',
  'Persian',
  'Ragdoll',
  'Bengal',
  'Sphynx',
  'Other',
]

const dogBreeds = [
  'Mixed Breed',
  'Labrador Retriever',
  'German Shepherd',
  'Golden Retriever',
  'Bulldog',
  'Poodle',
  'Beagle',
  'Rottweiler',
  'Chihuahua',
  'Pit Bull Terrier',
  'Other',
]

const breedOptions = computed(() => {
  const species = props.settings.pets.defaultSpecies?.toLowerCase() || 'cat'
  const list = species === 'dog' ? dogBreeds : catBreeds
  return list.map((b) => ({ label: b, value: b }))
})

const compatibilityOptions = [
  { label: 'Unknown', value: 'unknown' },
  { label: 'Yes', value: 'yes' },
  { label: 'No', value: 'no' },
]
</script>

<template>
  <div class="settings-section">
    <div class="settings-group">
      <h3>Defaults</h3>
      <p class="description">Set default values for new pet records to speed up intake.</p>

      <div class="form-row">
        <Select
          label="Default Species"
          v-model="localSettings.pets.defaultSpecies"
          :options="speciesOptions"
          fullWidth
        />
        <Select
          label="Default Breed"
          v-model="localSettings.pets.defaultBreed"
          :options="breedOptions"
          fullWidth
        />
      </div>

      <div class="form-row">
        <Select
          label="Default Intake Status"
          v-model="localSettings.pets.defaultIntakeStatus"
          :options="statusOptions"
          fullWidth
        />
        <Select
          label="Default Environment"
          v-model="localSettings.pets.defaultEnvironment"
          :options="environmentOptions"
          fullWidth
        />
      </div>
      <div class="form-row">
        <InputField
          label="Default Shelter Location"
          v-model="localSettings.pets.defaultShelterLocation"
          placeholder="e.g. Main Location"
        />
        <InputField
          label="Default Adoption Fee ($)"
          v-model.number="localSettings.pets.defaultAdoptionFee"
          type="number"
          placeholder="0.00"
        />
      </div>
    </div>

    <div class="settings-group">
      <h3>Units & Formatting</h3>
      <div class="form-row">
        <Select
          label="Weight Units"
          v-model="localSettings.pets.measurementSystem"
          :options="unitOptions"
          fullWidth
        />
      </div>
    </div>

    <div class="settings-group">
      <h3>Defaults: Story & Compatibility</h3>
      <div class="form-row">
        <InputTextArea
          label="Default Bio Template"
          v-model="localSettings.pets.defaultBioTemplate"
          placeholder="e.g. Meet [Name]! This [Breed] is..."
          rows="4"
          style="grid-column: span 2"
        />
      </div>
      <div class="form-row">
        <Select
          label="Good with Cats?"
          v-model="localSettings.pets.defaultGoodWithCats"
          :options="compatibilityOptions"
          fullWidth
        />
        <Select
          label="Good with Dogs?"
          v-model="localSettings.pets.defaultGoodWithDogs"
          :options="compatibilityOptions"
          fullWidth
        />
        <Select
          label="Good with Kids?"
          v-model="localSettings.pets.defaultGoodWithKids"
          :options="compatibilityOptions"
          fullWidth
        />
      </div>
    </div>

    <div class="settings-group">
      <h3>Automation & Visibility</h3>
      <div class="form-row">
        <InputField
          label="Default Auto-Hold Period (Days)"
          v-model.number="localSettings.pets.defaultAutoHoldPeriod"
          type="number"
          placeholder="0"
        />
      </div>
      <div class="checkbox-list">
        <ButtonToggle
          label="Show Medical History by Default"
          v-model="localSettings.pets.defaultShowMedical"
          :true-value="true"
          :false-value="false"
        />
        <ButtonToggle
          label="Require Microchip on Intake"
          v-model="localSettings.pets.requireMicrochip"
          :true-value="true"
          :false-value="false"
        />
        <ButtonToggle
          label="Require Photo to Publish"
          v-model="localSettings.pets.requirePhotoForPublic"
          :true-value="true"
          :false-value="false"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.settings-section {
  max-width: 800px;
  background: white;
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 32px;
}

.settings-group {
  margin-bottom: 32px;
}

.settings-group:last-child {
  margin-bottom: 0;
}

h3 {
  font-size: 1.1rem;
  margin: 0 0 8px 0;
  color: var(--text-primary);
}

.description {
  color: var(--text-secondary);
  font-size: 0.9rem;
  margin-bottom: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 20px;
}

.checkbox-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

@media (max-width: 640px) {
  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
