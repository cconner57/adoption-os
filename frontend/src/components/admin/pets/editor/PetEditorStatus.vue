<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../../models/common'
import { ButtonToggle,InputField } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

function handleAdoptionPhotoUpload() {
  const mockPhoto = {
    url: 'https://placekitten.com/400/400',
    thumbnailUrl: 'https://placekitten.com/100/100',
    isPrimary: false,
    uploadedAt: new Date().toISOString(),
  }
  if (formData.value.adoption) {
    formData.value.adoption.photo = mockPhoto
  }
}

function removeAdoptionPhoto() {
  if (formData.value.adoption) {
    formData.value.adoption.photo = null
  }
}
</script>

<style scoped>
@import './form.css';
</style>

<template>
  <div class="form-section">
    <template v-if="formData.details?.status === 'adopted' && formData.adoption">
      <h4 class="section-title">Adoption Details</h4>
      <div class="form-row">
        <!-- Debugging raw input -->
        <div class="field">
          <label class="label">Adoption Date (Debug)</label>
          <div class="control">
            <input
              class="input"
              type="date"
              :value="formData.adoption.date"
              @input="
                (e) => {
                  const val = (e.target as HTMLInputElement).value
                  console.error('DEBUG: RAW INPUT UPDATE:', val)
                  formData.adoption.date = val
                }
              "
            />
          </div>
        </div>
        <InputField label="Adoption Fee ($)" v-model="formData.adoption.fee" type="number" />
      </div>
      <div class="form-row">
        <InputField
          label="Adopted By"
          v-model="formData.adoption.adoptedBy"
          placeholder="Adopter Name"
        />
        <InputField
          label="New Adopted Name"
          v-model="formData.adoption.newAdoptedName"
          placeholder="New Name"
        />
      </div>
      <div class="form-row" v-if="formData.adoption.adopterContactInfo">
        <InputField
          label="Adopter Email"
          v-model="formData.adoption.adopterContactInfo.email"
          placeholder="email@example.com"
        />
        <InputField
          label="Adopter Phone"
          v-model="formData.adoption.adopterContactInfo.phone"
          placeholder="(555) 123-4567"
        />
      </div>
      <div class="form-group checkbox-list">
        <ButtonToggle
          label="Post-Adoption Survey Completed?"
          v-model="formData.adoption.surveyCompleted"
          :true-value="true"
          :false-value="false"
        />
      </div>

      <!-- Adoption Photo -->
      <div class="form-group">
        <label class="input-label">Adoption Photo (Family)</label>
        <div style="display: flex; gap: 12px; align-items: flex-start; margin-top: 8px">
          <div
            v-if="formData.adoption.photo"
            style="
              position: relative;
              width: 120px;
              height: 120px;
              border-radius: 8px;
              overflow: hidden;
              border: 1px solid var(--border-color);
            "
          >
            <img
              :src="formData.adoption.photo.url"
              style="width: 100%; height: 100%; object-fit: cover"
            />
            <button
              @click="removeAdoptionPhoto"
              style="
                position: absolute;
                top: 4px;
                right: 4px;
                background: rgba(0, 0, 0, 0.6);
                color: white;
                border: none;
                border-radius: 50%;
                width: 20px;
                height: 20px;
                display: flex;
                align-items: center;
                justify-content: center;
                cursor: pointer;
                font-size: 0.7rem;
              "
            >
              ✕
            </button>
          </div>

          <div
            @click="handleAdoptionPhotoUpload"
            style="
              width: 120px;
              height: 120px;
              border: 1px dashed var(--border-color);
              border-radius: 8px;
              display: flex;
              flex-direction: column;
              align-items: center;
              justify-content: center;
              cursor: pointer;
              background: var(--bg-secondary);
              color: var(--text-secondary);
              transition: all 0.2s;
            "
          >
            <span style="font-size: 1.5rem">📸</span>
            <span style="font-size: 0.75rem; margin-top: 4px">{{
              formData.adoption.photo ? 'Change' : 'Upload'
            }}</span>
          </div>
        </div>
      </div>

      <hr class="divider" />
    </template>

    <template v-if="formData.details?.status === 'foster' && formData.foster">
      <h4 class="section-title">Foster Program</h4>
      <div class="form-row">
        <InputField
          label="Foster Parent"
          v-model="formData.foster.parentName"
          placeholder="Foster Name"
        />
      </div>
      <div class="form-row">
        <InputField label="Start Date" v-model="formData.foster.startDate" type="date" />
        <InputField label="End Date" v-model="formData.foster.endDate" type="date" />
      </div>
      <hr class="divider" />
    </template>

    <h4 class="section-title">Sponsorship</h4>
    <div class="form-group checkbox-list" v-if="formData.sponsored">
      <ButtonToggle
        label="Is Sponsored?"
        v-model="formData.sponsored.isSponsored"
        :true-value="true"
        :false-value="false"
      />
    </div>
    <div class="form-row" v-if="formData.sponsored?.isSponsored">
      <InputField
        label="Sponsored By"
        v-model="formData.sponsored.sponsoredBy"
        placeholder="Sponsor Name"
      />
      <InputField label="Amount ($)" v-model="formData.sponsored.amount" type="number" />
    </div>

    <hr class="divider" />
    <h4 class="section-title">Return History</h4>
    <div class="form-group checkbox-list" v-if="formData.returned">
      <ButtonToggle
        label="Has been Returned?"
        v-model="formData.returned.isReturned"
        :true-value="true"
        :false-value="false"
      />
    </div>
    <div class="form-row" v-if="formData.returned?.isReturned">
      <!-- Helper to add new return -->
      <div style="width: 100%; display: flex; flex-direction: column; gap: 12px">
        <h5 style="margin: 0; font-size: 0.9rem; color: var(--text-secondary)">History</h5>

        <div
          v-for="(record, idx) in formData.returned.history"
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
          <InputField label="Return Date" v-model="record.date" type="date" />
          <InputField label="Reason" v-model="record.reason" placeholder="Reason for return" />
          <button
            @click="formData.returned.history.splice(idx, 1)"
            style="color: #ef4444; background: none; border: none; padding: 8px; cursor: pointer"
          >
            ✕
          </button>
        </div>

        <button
          @click="formData.returned.history.push({ date: '', reason: '' })"
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
        >
          + Add Return Record
        </button>
      </div>
    </div>
  </div>
</template>
