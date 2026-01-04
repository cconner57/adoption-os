<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import InputSelectGroup from '../../common/ui/InputSelectGroup.vue'
import type { FormState, PastPet } from '../../../models/adopt-form.ts'

const { modelValue } = defineProps<{
  modelValue: FormState
  touched?: Record<string, boolean>
  handleBlur?: (field: string) => void
}>()

const addPet = () => {
  modelValue.pastPets.push({
    name: '',
    speciesBreedSize: '',
    age: '',
    source: '',
    spayedNeutered: '',
    passedAwayReason: '',
  })
}

const removePet = (index: number) => {
  modelValue.pastPets.splice(index, 1)
}
</script>

<template>
  <div class="past-pets-section">
    <InputSelectGroup
      label="Did you have pets in the past?"
      :options="['Yes', 'No']"
      :modelValue="modelValue.ownPetsBefore"
      @update:modelValue="(val) => (modelValue.ownPetsBefore = val as string)"
    />
    <div class="desktop-spacer"></div>
    <div class="children" v-if="modelValue.ownPetsBefore === 'Yes'">
      <div v-for="(pet, index) in modelValue.pastPets" :key="index" class="pet-entry">
        <div class="pet-header">
          <h4>PAST PET {{ index + 1 }}</h4>
          <button
            v-if="index === modelValue.pastPets.length - 1"
            class="add-btn"
            @click.prevent="addPet"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              width="24"
              height="24"
            >
              <path
                d="M12 4.5c.414 0 .75.336.75.75v6h6c.414 0 .75.336.75.75s-.336.75-.75.75h-6v6c0 .414-.336.75-.75.75s-.75-.336-.75-.75v-6h-6c-.414 0-.75-.336-.75-.75s.336-.75.75-.75h6v-6c0-.414.336-.75.75-.75z"
              />
            </svg>
          </button>
          <button v-if="index > 0" class="remove-btn" @click.prevent="removePet(index)">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              width="24"
              height="24"
            >
              <rect x="5" y="11" width="14" height="2" />
            </svg>
          </button>
        </div>

        <div class="pet-fields">
          <InputField
            v-model="pet.name"
            label="Name:"
            :name="`pet-name-${index}`"
            placeholder="Ex: Fluffy"
          />
          <InputField
            v-model="pet.speciesBreedSize"
            label="Species, Breed and Size:"
            :name="`pet-breed-${index}`"
            placeholder="Ex: Dog, Boxer, 50lbs"
          />
          <InputField
            v-model="pet.age"
            label="How long did you have this pet (and when did you get this pet)?:"
            :name="`pet-age-${index}`"
            placeholder="Ex: I've had him 10 years - from 2010 to Current"
          />
          <InputSelectGroup
            label="Where did you get this pet?"
            :options="['Friend/Family', 'Rescue', 'Shelter', 'Breeder', 'Online Ad', 'Found it']"
            :modelValue="pet.source"
            @update:modelValue="(val) => (pet.source = val as string)"
          />
          <InputSelectGroup
            label="Is this pet spayed/neutered?"
            :options="['Yes', 'No', 'Not sure']"
            :modelValue="pet.spayedNeutered"
            @update:modelValue="(val) => (pet.spayedNeutered = val as string)"
          />
          <InputSelectGroup
            label="What happened to this pet?"
            :options="[
              'Passed of old age',
              'Hit by car',
              'Died of disease',
              'Gave away',
              'Gave to shelter',
              'Put to sleep',
              'N/A',
            ]"
            :modelValue="pet.passedAwayReason"
            @update:modelValue="(val) => (pet.passedAwayReason = val as string)"
          />
          <hr class="pet-divider" v-if="index < modelValue.pastPets.length - 1" />
        </div>
      </div>

      <div v-if="modelValue.pastPets.length === 0" class="no-pets-placeholder">
        <button class="add-btn-large" @click.prevent="addPet">Add a Pet</button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.past-pets-section {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

@media (max-width: 768px) {
  .past-pets-section {
    grid-template-columns: 1fr;
  }
  .desktop-spacer {
    display: none;
  }
}

.children {
  grid-column: 1 / -1;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.pet-entry {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.pet-header {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.pet-header h4 {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--font-color-dark);
}

.pet-fields {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

@media (min-width: 768px) {
  .pet-fields {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1.5rem 2rem;
    align-items: start;
  }
}

.pet-divider {
  border: 0;
  border-top: 2px solid #cbd5e1; /* Darker and slightly thicker for contrast */
  margin: 2rem 0; /* More spacing to separate distinct pets */
  grid-column: 1 / -1;
}

/* Button styles */
.add-btn,
.remove-btn {
  width: 40px;
  height: 40px;
  background: none;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--green);
  transition: all 0.2s;
}

.add-btn:hover {
  color: var(--blue-hover);
  border-color: var(--blue-hover);
  background: #f0f9ff;
}

.remove-btn:hover {
  color: var(--red);
  border-color: var(--red);
  background: #fff5f5;
}

.add-btn-large {
  padding: 12px 24px;
  background-color: var(--green);
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
}
</style>
