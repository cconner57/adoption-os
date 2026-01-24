<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { InputTextArea } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const additionalInfoString = computed({
  get: () => {
    return formData.value.descriptions?.additionalInformation?.join('\n') || ''
  },
  set: (val: string) => {
    if (!formData.value.descriptions) return
    formData.value.descriptions.additionalInformation = val
      .split('\n')
      .map((s) => s.trim())
      .filter(Boolean)
  },
})
</script>

<style scoped>
@import url('./form.css');

.form-section {
  gap: 12px;
}
</style>

<template>
  <div class="form-section">
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Primary Bio"
        :model-value="formData.descriptions.primary || null"
        @update:model-value="val => formData.descriptions && (formData.descriptions.primary = val)"
        placeholder="Tell their story..."
        :maxChars="500"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Origin / History"
        :model-value="formData.descriptions.origin || null"
        @update:model-value="val => formData.descriptions && (formData.descriptions.origin = val)"
        placeholder="Where did they come from?"
        :maxChars="300"
      />
    </div>

    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Fun Facts"
        :model-value="formData.descriptions.fun || null"
        @update:model-value="val => formData.descriptions && (formData.descriptions.fun = val)"
        placeholder="Quirks and cute habits..."
        :maxChars="200"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Special Needs Description"
        :model-value="formData.descriptions.specialNeeds || null"
        @update:model-value="val => formData.descriptions && (formData.descriptions.specialNeeds = val)"
        placeholder="Details about medical/behavioral needs..."
        :maxChars="300"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Spotlight Blurb (Short)"
        placeholder="Featured on homepage..."
        :model-value="formData.descriptions.spotlight || null"
        @update:model-value="val => formData.descriptions && (formData.descriptions.spotlight = val)"
        :maxChars="100"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Additional Information"
        v-model="additionalInfoString"
        placeholder="Any extra details (one per line)..."
        :rows="4"
        :maxChars="1000"
      />
    </div>
  </div>
</template>
