<script setup lang="ts">
import type { IVolunteer } from '../../../../stores/mockVolunteerData'
import { formatPhoneInput, formatZipInput } from '../../../../utils/formatters'
import { InputField, InputFileUpload } from '../../../common/ui'

defineProps<{
  formData: Partial<IVolunteer>
  canEdit: boolean
}>()

defineEmits(['update:formData'])
</script>

<template>
  <div class="form-section">
    <h3 class="section-title">General Information</h3>
    <div class="form-grid">
      <div class="full-width" style="margin-bottom: 24px">
        <InputFileUpload
          label="Profile Photo"
          :model-value="formData.photoUrl"
          @update:model-value="(val) => $emit('update:formData', { ...formData, photoUrl: val })"
          accept="image/*"
          :disabled="!canEdit"
        />
      </div>
      <InputField
        label="First Name"
        placeholder="First Name"
        :model-value="formData.firstName || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, firstName: val as string })"
        :disabled="!canEdit"
      />
      <InputField
        label="Last Name"
        placeholder="Last Name"
        :model-value="formData.lastName || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, lastName: val as string })"
        :disabled="!canEdit"
      />
      <InputField
        label="Email"
        placeholder="Email"
        :model-value="formData.email || ''"
        @update:model-value="(val) => $emit('update:formData', { ...formData, email: val as string })"
        :disabled="!canEdit"
      />
      <div class="row-full">
        <InputField
          label="Phone"
          placeholder="Phone"
          :model-value="formData.phone || ''"
          @update:model-value="
            (val: any) =>
              $emit('update:formData', {
                ...formData,
                phone: formatPhoneInput(String(val || '')),
              })
          "
          maxlength="14"
          fullWidth
          :disabled="!canEdit"
        />
      </div>

      <div class="row-full">
        <InputField
          label="Address"
          placeholder="Address"
          :model-value="formData.address || ''"
          @update:model-value="(val) => $emit('update:formData', { ...formData, address: val as string })"
          class="full-width"
          :disabled="!canEdit"
        />
      </div>

      <div class="row-full">
        <InputField
          label="City"
          placeholder="City"
          :model-value="formData.city || ''"
          @update:model-value="(val) => $emit('update:formData', { ...formData, city: val as string })"
          class="full-width"
          :disabled="!canEdit"
        />
      </div>

      <div class="row-2 full-width">
        <InputField
          label="Zip"
          placeholder="Zip"
          :model-value="formData.zip || ''"
          @update:model-value="
            (val: any) =>
              $emit('update:formData', { ...formData, zip: formatZipInput(String(val || '')) })
          "
          maxlength="5"
          fullWidth
          class="full-width"
          :disabled="!canEdit"
        />
        <InputField
          label="Birthday"
          placeholder="Birthday"
          type="date"
          :model-value="formData.birthday || ''"
          @update:model-value="(val) => $emit('update:formData', { ...formData, birthday: val as string })"
          fullWidth
          class="full-width"
          :disabled="!canEdit"
        />
      </div>
    </div>

    <h3 class="section-title mt-6">Emergency Contact</h3>
    <div class="form-grid">
      <InputField
        label="Name"
        placeholder="Name"
        :model-value="formData.emergencyContactName || ''"
        @update:model-value="
          (val) => $emit('update:formData', { ...formData, emergencyContactName: val as string })
        "
        :disabled="!canEdit"
      />
      <InputField
        label="Phone"
        placeholder="Phone"
        :model-value="formData.emergencyContactPhone || ''"
        @update:model-value="
          (val: any) =>
            $emit('update:formData', {
              ...formData,
              emergencyContactPhone: formatPhoneInput(String(val || '')),
            })
        "
        maxlength="14"
        fullWidth
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

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.full-width {
  grid-column: 1 / -1;
}

.mt-6 {
  margin-top: 24px;
}

.row-full {
  width: 100%;
}

.row-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  width: 100%;
}

.row-2 :deep(input),
.row-full :deep(input) {
  width: 100% !important;
}
</style>
