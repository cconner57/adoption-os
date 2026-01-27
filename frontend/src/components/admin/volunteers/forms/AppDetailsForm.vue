<script setup lang="ts">
import type { IVolunteer } from '../../../../stores/mockVolunteerData'
import { InputTextArea } from '../../../common/ui'
import Availability from '../../../volunteer/availability/Availability.vue'
import PositionPreferences from '../../../volunteer/preferences/PositionPreferences.vue'

defineProps<{
  formData: Partial<IVolunteer>
  canEdit: boolean
}>()

defineEmits(['update:formData'])
</script>

<template>
  <div class="form-section">
    <h3 class="section-title">Application Details</h3>
    <div class="stack">
      <InputTextArea
        label="Why they joined"
        placeholder="Reason for joining..."
        :model-value="formData.interestReason || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, interestReason: val as string })"
        :disabled="!canEdit"
      />
      <InputTextArea
        label="Volunteer Experience"
        placeholder="Previous experience..."
        :model-value="formData.volunteerExperience || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, volunteerExperience: val as string })"
        :disabled="!canEdit"
      />

      <h3 class="section-subtitle mt-4">Position Preferences</h3>
      <PositionPreferences
        :model-value="formData.positionPreferences || []"
        @update:model-value="(val) => $emit('update:formData', { ...formData, positionPreferences: val })"
        :disabled="!canEdit"
      />

      <h3 class="section-subtitle mt-4">Availability</h3>
      <Availability
        :model-value="formData.availability || []"
        @update:model-value="(val) => $emit('update:formData', { ...formData, availability: val })"
        :disabled="!canEdit"
      />
    </div>
  </div>
</template>

<style scoped>
.section-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-color);
}

.section-subtitle {
  font-size: 1rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.stack {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.mt-4 {
  margin-top: 16px;
}
</style>
