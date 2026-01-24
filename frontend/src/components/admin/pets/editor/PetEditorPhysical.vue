<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { InputField, Select } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const ageGroupOptions = [
  { label: 'Baby (Kitten/Puppy)', value: 'baby' },
  { label: 'Young', value: 'young' },
  { label: 'Adult', value: 'adult' },
  { label: 'Senior', value: 'senior' },
]

const sizeOptions = [
  { label: 'Extra Small', value: 'extra-small' },
  { label: 'Small', value: 'small' },
  { label: 'Medium', value: 'medium' },
  { label: 'Large', value: 'large' },
  { label: 'Extra Large', value: 'extra-large' },
]

const coatLengthOptions = [
  { label: 'Short', value: 'short' },
  { label: 'Medium', value: 'medium' },
  { label: 'Long', value: 'long' },
  { label: 'Wire', value: 'wire' },
  { label: 'Hairless', value: 'hairless' },
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
  const species = formData.value.species?.toLowerCase() || 'cat'
  const list = species === 'dog' ? dogBreeds : catBreeds
  return list.map((b) => ({ label: b, value: b }))
})
</script>

<template>
  <div class="form-section" v-if="formData.physical">
    <Select label="Breed" v-model="formData.physical.breed" :options="breedOptions" fullWidth />

    <div class="form-row">
      <InputField
        label="Color"
        v-model="formData.physical.color"
        placeholder="e.g. Black & White"
      />
      <InputField
        label="Date of Birth"
        :model-value="formData.physical.dateOfBirth || null"
        @update:model-value="val => formData.physical && (formData.physical.dateOfBirth = (val as string))"
        type="date"
        placeholder="mm/dd/yyyy"
      />
    </div>

    <div class="form-row">
      <Select
        label="Age Group"
        v-model="formData.physical.ageGroup"
        :options="ageGroupOptions"
        fullWidth
      />
      <Select label="Size" v-model="formData.physical.size" :options="sizeOptions" fullWidth />
    </div>

    <div class="form-row">
      <Select
        label="Coat Length"
        v-model="formData.physical.coatLength"
        :options="coatLengthOptions"
        fullWidth
      />
      <InputField
        label="Weight (lbs/kg)"
        :model-value="formData.physical.currentWeight || null"
        @update:model-value="val => formData.physical && (formData.physical.currentWeight = (val as number))"
        type="number"
        placeholder="e.g. 10.5"
      />
    </div>
  </div>
</template>

<style scoped>
@import url('./form.css');

.dob-group {
  flex: 1;
}
</style>
