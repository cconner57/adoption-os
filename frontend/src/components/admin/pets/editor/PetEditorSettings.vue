<script setup lang="ts">
import { computed } from 'vue'
import { ButtonToggle } from '../../../common/ui'
import type { IPet } from '../../../../../models/common'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})
</script>

<style scoped>
@import './form.css';
</style>

<template>
  <div class="form-section">
    <div class="form-group checkbox-list" v-if="formData.profileSettings">
      <ButtonToggle
        label="Feature in Spotlight"
        v-model="formData.profileSettings.isSpotlightFeatured"
        :true-value="true"
        :false-value="false"
      />
      <ButtonToggle
        label="Show Medical History Publicly"
        v-model="formData.profileSettings.showMedicalHistory"
        :true-value="true"
        :false-value="false"
      />
      <ButtonToggle
        label="Show Additional Information"
        v-model="formData.profileSettings.showAdditionalInformation"
        :true-value="true"
        :false-value="false"
      />
    </div>
  </div>
</template>
