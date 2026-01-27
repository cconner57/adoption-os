<script setup lang="ts">
import type { IVolunteer } from '../../../../stores/mockVolunteerData'
import { InputTextArea, Toggle } from '../../../common/ui'

defineProps<{
  formData: Partial<IVolunteer>
  canEdit: boolean
}>()

defineEmits(['update:formData'])
</script>

<template>
  <div class="form-section">
    <h3 class="section-title">Bio & Skills</h3>
    <div class="stack">
      <InputTextArea
        label="Bio"
        placeholder="Bio description..."
        :model-value="formData.bio || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, bio: val as string })"
        rows="4"
        :disabled="!canEdit"
      />

      <div class="field-group">
        <label class="field-label">Skills (comma separated)</label>
        <input
          class="simple-input"
          :value="formData.skills?.join(', ')"
          @input="
            (e) =>
              $emit('update:formData', {
                ...formData,
                skills: (e.target as HTMLInputElement).value
                  .split(',')
                  .map((s) => s.trim())
                  .filter(Boolean),
              })
          "
          :disabled="!canEdit"
        />
      </div>

      <div class="field-group">
        <Toggle
          :model-value="formData.allergies || false"
          @update:model-value="(val) => $emit('update:formData', { ...formData, allergies: val })"
          label="Has Allergies?"
          labelPosition="left"
          :disabled="!canEdit"
        />
      </div>
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

.stack {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.field-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field-label {
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--text-primary);
}

.simple-input {
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  background: #fff;
  width: 100%;
  font-size: 0.95rem;

  &:focus {
    outline: 2px solid var(--color-primary);
    border-color: transparent;
  }
}
</style>
