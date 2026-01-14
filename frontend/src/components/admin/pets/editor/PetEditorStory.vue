<script setup lang="ts">
import { computed } from 'vue'
import { InputTextArea } from '../../../common/ui'
import type { IPet } from '../../../../../models/common'

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
@import './form.css';

.form-section {
  gap: 12px;
}
</style>

<template>
  <div class="form-section">
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Primary Bio"
        v-model="formData.descriptions.primary"
        placeholder="Tell their story..."
        :maxChars="500"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Origin / History"
        v-model="formData.descriptions.origin"
        placeholder="Where did they come from?"
        :maxChars="300"
      />
    </div>

    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Fun Facts"
        v-model="formData.descriptions.fun"
        placeholder="Quirks and cute habits..."
        :maxChars="200"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Special Needs Description"
        v-model="formData.descriptions.specialNeeds"
        placeholder="Details about medical/behavioral needs..."
        :maxChars="300"
      />
    </div>
    <div class="form-group" v-if="formData.descriptions">
      <InputTextArea
        label="Spotlight Blurb (Short)"
        v-model="formData.descriptions.spotlight"
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
