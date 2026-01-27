<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { Combobox, InputField, Select } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
  availablePets?: IPet[]
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const speciesOptions = [
  { label: 'Cat', value: 'cat' },
  { label: 'Dog', value: 'dog' },
]

const sexOptions = [
  { label: 'Male', value: 'male' },
  { label: 'Female', value: 'female' },
  { label: 'Unknown', value: 'unknown' },
]

const environmentOptions = [
  { label: 'Indoor', value: 'indoor' },
  { label: 'Outdoor', value: 'outdoor' },
  { label: 'Indoor/Outdoor', value: 'indoor/outdoor' },
]

const litterTypeOptions = [
  { label: 'Clay (Clumping)', value: 'clay_clumping' },
  { label: 'Clay (Non-Clumping)', value: 'clay_non_clumping' },
  { label: 'Crystal / Silica', value: 'crystal' },
  { label: 'Pine / Wood', value: 'pine' },
  { label: 'Paper', value: 'paper' },
  { label: 'Corn / Wheat', value: 'corn_wheat' },
  { label: 'Grass', value: 'grass' },
  { label: 'Tofu', value: 'tofu' },
]

const statusOptions = [
  { label: 'Available', value: 'available' },
  { label: 'Adoption Pending', value: 'adoption-pending' },
  { label: 'Adopted', value: 'adopted' },
  { label: 'In Foster', value: 'foster' },
  { label: 'Medical/Behavioral Hold', value: 'hold' },
  { label: 'Intake Processing', value: 'intake' },
  { label: 'Archived', value: 'archived' },
]

const litterOptions = computed(() => {
  if (!props.availablePets) return []

  const litters = new Set(props.availablePets.map((p) => p.litterName).filter(Boolean))
  return Array.from(litters).map((l) => ({ label: l!, value: l! }))
})
</script>

<style scoped>
@import url('./form.css');
</style>

<template>
  <div class="form-section">
    <InputField label="Name" :model-value="formData.name || null" @update:model-value="val => formData.name = (val as string)" placeholder="Pet's Name" />

    <div class="form-row">
      <Select
        label="Species"
        :model-value="formData.species || null"
        @update:model-value="val => formData.species = (val as any)"
        :options="speciesOptions"
        fullWidth
      />
      <Select
        label="Sex"
        :model-value="formData.sex || null"
        @update:model-value="val => formData.sex = (val as any)"
        :options="sexOptions"
        fullWidth
      />
    </div>

    <div class="form-row">
      <Combobox
        label="Litter Name"
        :model-value="formData.litterName || null"
        @update:model-value="val => formData.litterName = (val as string)"
        :options="litterOptions"
        placeholder="Select or create litter..."
        allow-create
      />
    </div>

    <div class="form-row" v-if="formData.details">
      <InputField
        label="Intake Date"
        :model-value="formData.details.intakeDate || null"
        @update:model-value="val => formData.details && (formData.details.intakeDate = (val as string))"
        placeholder="YY-MMDD"
      />
      <Select
        label="Environment"
        :model-value="formData.details.environmentType || null"
        @update:model-value="val => formData.details && (formData.details.environmentType = (val as any))"
        :options="environmentOptions"
        fullWidth
      />
    </div>

    <div class="form-row">
      <div v-if="formData.adoption" style="flex: 1">
        <InputField
          label="Adoption Fee ($)"
          :model-value="formData.adoption.fee || null"
          @update:model-value="val => formData.adoption && (formData.adoption.fee = (val as number))"
          type="number"
          placeholder="0.00"
        />
      </div>
      <div v-if="formData.details" style="flex: 1">
        <Select
          label="Status"
          v-model="formData.details.status"
          :options="statusOptions"
          fullWidth
        />
      </div>
    </div>

    <div class="form-row" v-if="formData.details">
      <InputField
        label="Shelter Location"
        :model-value="formData.details.shelterLocation || null"
        @update:model-value="val => formData.details && (formData.details.shelterLocation = (val as string))"
        placeholder="e.g. Main Location, Foster"
      />
    </div>

    <div class="form-row" v-if="formData.species === 'cat' && formData.details">
      <Combobox
        label="Preferred Litter Type"
        :model-value="formData.details.preferredPetLitterType || null"
        @update:model-value="val => formData.details && (formData.details.preferredPetLitterType = (val as string))"
        :options="litterTypeOptions"
        placeholder="Select litter type..."
        allow-create
      />
    </div>
  </div>
</template>
