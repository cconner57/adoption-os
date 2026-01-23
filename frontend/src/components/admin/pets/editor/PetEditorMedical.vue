<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { ButtonToggle, Combobox, InputField, InputTextArea } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const medicalConcernOptions = [
  { label: 'Allergies - Flea', value: 'allergies - flea' },
  { label: 'Allergies - Food', value: 'allergies - food' },
  { label: 'Allergies - Skin', value: 'allergies - skin' },
  { label: 'Anemia', value: 'anemia' },
  { label: 'Asthma', value: 'asthma' },
  { label: 'Bladder Infection', value: 'bladder infection' },
  { label: 'Cancer', value: 'cancer' },
  { label: 'Cystitis', value: 'cystitis' },
  { label: 'Dental Problems', value: 'dental problems' },
  { label: 'Diabetes', value: 'diabetes' },
  { label: 'Ear Infections', value: 'ear infections' },
  { label: 'FIV', value: 'feline immunodeficiency (fiv)' },
  { label: 'FIP', value: 'feline infectious peritonitis (fip)' },
  { label: 'FeLV', value: 'feline leukemia virus (felv)' },
  { label: 'Gastrointestinal Issues', value: 'gastrointestinal issues' },
  { label: 'Heartworm Disease', value: 'heartworm disease' },
  { label: 'Hyperthyroidism', value: 'hyperthyroidism' },
  { label: 'Kidney Disease', value: 'kidney disease' },
  { label: 'Obesity', value: 'obesity' },
  { label: 'Upper Respiratory Infections', value: 'upper respiratory infections' },
]
</script>

<style scoped>
@import './form.css';
</style>

<template>
  <div class="form-section" v-if="formData.medical">
    <div class="form-group checkbox-list">
      <ButtonToggle
        label="Spayed / Neutered"
        v-model="formData.medical.spayedOrNeutered"
        :true-value="true"
        :false-value="false"
      />

      <div v-if="formData.medical.spayedOrNeutered" style="margin-left: 12px; margin-top: 8px">
        <label class="input-label" style="font-size: 0.85rem; margin-bottom: 4px"
          >Date Spayed/Neutered</label
        >
        <input type="date" v-model="formData.medical.spayedOrNeuteredDate" class="text-input" />
      </div>
    </div>

    <div class="form-group checkbox-list">
      <ButtonToggle
        label="Vaccinations Up to Date"
        v-model="formData.medical.vaccinationsUpToDate"
        :true-value="true"
        :false-value="false"
      />
    </div>

    <div class="form-group checkbox-list">
      <ButtonToggle
        label="Microchipped"
        v-model="formData.medical.microchip.microchipped"
        :true-value="true"
        :false-value="false"
      />
      <div
        style="margin-left: 12px; margin-top: 8px; display: flex; flex-direction: column; gap: 12px"
        v-if="formData.medical.microchip.microchipped"
      >
        <InputField
          label="Microchip ID"
          :model-value="formData.medical.microchip.microchipID || null"
          @update:model-value="val => formData.medical?.microchip && (formData.medical.microchip.microchipID = (val as string))"
          placeholder="e.g. 98102000..."
        />
        <InputField
          label="Company / Brand"
          :model-value="formData.medical.microchip.microchipCompany || null"
          @update:model-value="val => formData.medical?.microchip && (formData.medical.microchip.microchipCompany = (val as string))"
          placeholder="e.g. HomeAgain"
        />
      </div>
    </div>

    <div class="form-group" v-if="formData.behavior">
      <InputTextArea
        label="Health Summary"
        :model-value="formData.behavior.healthSummary || null"
        @update:model-value="val => formData.behavior && (formData.behavior.healthSummary = val)"
        placeholder="General summary of health status..."
        :maxChars="500"
      />
    </div>

    <hr class="divider" />

    <div class="form-group">
      <Combobox
        label="Health Concerns / Conditions"
        :model-value="(formData.medical.healthConcerns as unknown as string[]) || null"
        @update:model-value="val => formData.medical && (formData.medical.healthConcerns = (Array.isArray(val) ? val : val ? [val] : []) as any)"
        :options="medicalConcernOptions"
        multiple
        placeholder="Select conditions..."
      />
    </div>
    <div class="form-group">
      <Combobox
        label="Current Medications"
        :model-value="formData.medical.currentMedications || null"
        @update:model-value="val => formData.medical && (formData.medical.currentMedications = (Array.isArray(val) ? val : val ? [val] : []))"
        :options="[]"
        allow-create
        multiple
        placeholder="Add medications..."
        no-results-text="Type to add medication..."
      />
    </div>

    <hr class="divider" />
    <h4 class="section-title">Vaccinations</h4>
    <div v-if="formData.medical.vaccinations">
      
      <div class="form-row">
        <InputField
          label="Rabies Date"
          :model-value="formData.medical.vaccinations.rabies?.dateAdministered || null"
          @update:model-value="val => formData.medical?.vaccinations?.rabies && (formData.medical.vaccinations.rabies.dateAdministered = (val as string))"
          placeholder="YYYY-MM-DD"
          type="date"
        />
      </div>

      <div style="display: flex; flex-direction: column; gap: 10px; margin-top: 10px">
        <label class="input-label">Other Vaccinations</label>
        <div
          v-for="(vax, idx) in formData.medical.vaccinations.other"
          :key="idx"
          style="
            display: flex;
            gap: 8px;
            align-items: flex-end;
            padding: 8px;
            border: 1px solid var(--border-color);
            border-radius: 6px;
          "
        >
          <InputField label="Name" :model-value="vax.name || null" @update:model-value="val => vax.name = (val as string)" placeholder="Vaccine Name" />
          <InputField label="Date" placeholder="YYYY-MM-DD" v-model="vax.dateAdministered" type="date" />
          <button
            @click="formData.medical?.vaccinations?.other?.splice(idx, 1)"
            style="color: #ef4444; background: none; border: none; padding: 8px; cursor: pointer"
            type="button"
          >
            ✕
          </button>
        </div>
        <button
          @click="formData.medical?.vaccinations?.other?.push({ name: '', dateAdministered: '' })"
          style="
            align-self: flex-start;
            background: none;
            border: 1px dashed var(--color-primary);
            color: var(--color-primary);
            padding: 8px 12px;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.85rem;
          "
          type="button"
        >
          + Add Vaccination
        </button>
      </div>
    </div>
  </div>
</template>
