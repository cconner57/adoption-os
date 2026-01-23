<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { ButtonToggle } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue', 'archive'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})
</script>

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

    <div class="form-section danger-zone">
      <h3>Danger Zone</h3>
      <p class="description">Archiving a pet will hide it from the public list.</p>
      <button class="btn-archive" @click="emit('archive')">Archive Pet</button>
    </div>
  </div>
</template>

<style scoped>
@import './form.css';

.danger-zone {
  margin-top: 40px;
  border-top: 1px solid var(--border-color);
  padding-top: 24px;
}

.danger-zone h3 {
  font-size: 1rem;
  color: var(--color-danger);
  margin-bottom: 8px;
}

.description {
  font-size: 0.9rem;
  color: var(--text-secondary);
  margin-bottom: 16px;
}

.btn-archive {
  background: white;
  border: 1px solid var(--color-danger);
  color: var(--color-danger);
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s;
}

.btn-archive:hover {
  background: var(--color-danger);
  color: white;
}
</style>
