<script setup lang="ts">
import type { IVolunteer } from '../../../../stores/mockVolunteerData'
import { Button, Select } from '../../../common/ui'

defineProps<{
  formData: Partial<IVolunteer>
  canEdit: boolean
}>()

defineEmits(['update:formData', 'archive'])
</script>

<template>
  <div class="form-section">
    <h3 class="section-title">Settings</h3>

    <div class="field-group" style="margin-bottom: 32px">
      <label class="field-label" style="margin-bottom: 8px; display: block">Permission Level</label>
      <Select
        :model-value="formData.role || 'Tier 1'"
        @update:model-value="(val) => $emit('update:formData', { ...formData, role: val })"
        :options="[
          { label: 'Tier 1 Volunteer', value: 'Tier 1' },
          { label: 'Tier 2 Volunteer', value: 'Tier 2' },
          { label: 'Teen Volunteer', value: 'Teen' },
          { label: 'Admin', value: 'Admin' },
        ]"
        :disabled="!canEdit"
      />
    </div>

    <div class="panel" style="margin-bottom: 24px">
      <div
        class="panel-header"
        style="padding: 16px; border-bottom: 1px solid var(--border-color); background: #f9fafb"
      >
        <h4 style="margin: 0 0 4px; font-weight: 600; font-size: 1rem">User Account</h4>
        <p style="margin: 0; color: var(--text-secondary); font-size: 0.9rem">
          Invite this volunteer to create a login account to access the dashboard.
        </p>
      </div>
      <div class="panel-actions">
        <Button
          title="Invite User"
          color="purple"
          style="width: 100%; justify-content: center"
          :onClick="() => {}"
          :disabled="!canEdit"
        />
      </div>
    </div>

    <div class="panel danger-zone">
      <div class="panel-header">
        <h4>Danger Zone</h4>
        <p>
          Archive this volunteer to hide them from the active list. You can unarchive them later.
        </p>
      </div>
      <div class="panel-actions">
        <Button
          color="white"
          title="Archive Member"
          @click="$emit('archive')"
          style="
            color: var(--color-danger);
            border-color: var(--color-danger);
            width: 100%;
            justify-content: center;
          "
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

.panel {
  border: 1px solid var(--border-color);
  border-radius: 8px;
  overflow: hidden;
}

.panel.danger-zone {
  border-color: #fca5a5;
}

.danger-zone .panel-header {
  background: #fef2f2;
  padding: 16px;
  border-bottom: 1px solid #fca5a5;
}

.danger-zone h4 {
  color: #991b1b;
  font-weight: 600;
  margin-bottom: 4px;
}

.danger-zone p {
  color: #b91c1c;
  font-size: 0.9rem;
  margin: 0;
}

.panel-actions {
  padding: 16px;
  background: #fff;
}
</style>
