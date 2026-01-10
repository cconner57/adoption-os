<script setup lang="ts">
import { computed } from 'vue'
import { InputField, Select, Combobox } from '../../../common/ui'
import type { IPet } from '../../../../../models/common'

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
  // Extract unique litter names
  const litters = new Set(props.availablePets.map((p) => p.litterName).filter(Boolean))
  return Array.from(litters).map((l) => ({ label: l!, value: l! }))
})
</script>

<style scoped>
@import './form.css';
</style>

<template>
  <div class="form-section">
    <InputField label="Name" v-model="formData.name" placeholder="Pet's Name" />

    <div class="form-row">
      <Select label="Species" v-model="formData.species" :options="speciesOptions" fullWidth />
      <Select label="Sex" v-model="formData.sex" :options="sexOptions" fullWidth />
    </div>

    <div class="form-row">
      <Combobox
        label="Litter Name"
        v-model="formData.litterName"
        :options="litterOptions"
        placeholder="Select or create litter..."
        allow-create
      />
    </div>

    <div class="form-row" style="margin-top: 20px" v-if="formData.details">
      <InputField label="Intake Date" v-model="formData.details.intakeDate" type="date" />
      <Select
        label="Environment"
        v-model="formData.details.environmentType"
        :options="environmentOptions"
        fullWidth
      />
    </div>

    <div class="form-row" style="margin-top: 12px" v-if="formData.details">
      <InputField
        label="Shelter Location"
        v-model="formData.details.shelterLocation"
        placeholder="e.g. Main Location, Foster"
      />
    </div>

    <div
      class="form-row"
      style="margin-top: 12px"
      v-if="formData.species === 'cat' && formData.details"
    >
      <Combobox
        label="Preferred Litter Type"
        v-model="formData.details.preferredPetLitterType"
        :options="litterTypeOptions"
        placeholder="Select litter type..."
        allow-create
      />
    </div>

    <Select
      v-if="formData.details"
      label="Status"
      v-model="formData.details.status"
      :options="statusOptions"
      fullWidth
    />
  </div>
</template>
