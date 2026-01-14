<script setup lang="ts">
import { ref } from 'vue'
import { Button, InputField, Select, Toggle } from '../../common/ui'

const props = defineProps<{
  initialData?: any
}>()

const emit = defineEmits(['close', 'save'])

const formData = ref({
  id: props.initialData?.id || undefined,
  date: props.initialData?.date || '',
  startTime: props.initialData?.startTime || '',
  endTime: props.initialData?.endTime || '',
  role: props.initialData?.role || 'Feeding/Cleaning',
  status: props.initialData?.status || 'scheduled',
  isRecurring: false, // Editing recurrence is tricky, maybe disable for edit?
  frequency: 'weekly',
  endDate: '',
  isCovering: false,
  coveringName: '',
})

const roles = [
  { label: 'Feeding/Cleaning', value: 'Feeding/Cleaning' },
  { label: 'Cat Socializing', value: 'Cat Socializing' },
  { label: 'Dog Walking', value: 'Dog Walking' },
  { label: 'Customer Support', value: 'Customer Support' },
  { label: 'Adoption Event', value: 'Adoption Event' },
]

const statuses = [
  { label: 'Scheduled', value: 'scheduled' },
  { label: 'Completed', value: 'completed' }, // mapped to all_good or logic?
  // Let's stick to the user's requested terms but map to DB values if needed or just use these keys
  // DB has 'scheduled', 'all_good', 'late', 'no_show', 'cancelled' in current schema?
  // I should probably map "Completed" -> "all_good", "Missed" -> "no_show".
  // Or I can add new strings to DB since it is text.
  // Let's use clean keys.
  { label: 'Completed', value: 'completed' },
  { label: 'Late', value: 'late' },
  { label: 'Missed', value: 'missed' },
  { label: 'Covered', value: 'covered' },
]

const frequencies = [
  { label: 'Weekly', value: 'weekly' },
  { label: 'Bi-Weekly', value: 'biweekly' },
  { label: 'Monthly', value: 'monthly' },
]

function handleSave() {
  const payload = { ...formData.value }

  // If covering, save to notes
  if (payload.isCovering && payload.coveringName) {
    payload.notes = `Covering for ${payload.coveringName}`
  }

  emit('save', payload)

  // Reset form
  formData.value = {
    date: '',
    startTime: '',
    endTime: '',
    role: 'Feeding/Cleaning',
    isRecurring: false,
    frequency: 'weekly',
    endDate: '',
    isCovering: false,
    coveringName: '',
  }
}
</script>

<template>
  <div class="shift-form-card">
    <div class="form-grid">
      <div class="row-2">
        <InputField label="Date" type="date" v-model="formData.date" class="full-width" />
        <div class="field-group">
          <label class="field-label">Role</label>
          <Select v-model="formData.role" :options="roles" />
        </div>
      </div>

      <div class="row-2">
        <InputField label="Start Time" type="time" v-model="formData.startTime" />
        <InputField label="End Time" type="time" v-model="formData.endTime" />
      </div>

      <div v-if="formData.id" class="field-group">
        <label class="field-label">Status</label>
        <Select v-model="formData.status" :options="statuses" />
      </div>

      <div v-if="!formData.id" class="recurring-section">
        <div class="field-group">
          <Toggle v-model="formData.isRecurring" label="Recurring Shift?" labelPosition="left" />
        </div>

        <div v-if="formData.isRecurring" class="recurring-options">
          <div class="row-2">
            <div class="field-group">
              <label class="field-label">Frequency</label>
              <Select v-model="formData.frequency" :options="frequencies" />
            </div>
            <InputField label="End Date" type="date" v-model="formData.endDate" />
          </div>
        </div>
      </div>

      <div class="field-group mt-4">
        <Toggle
          v-model="formData.isCovering"
          label="Is this covering another volunteer?"
          labelPosition="left"
        />
      </div>

      <div v-if="formData.isCovering" class="field-group mt-2">
        <InputField
          label="Who are you covering for?"
          placeholder="Volunteer Name"
          v-model="formData.coveringName"
        />
      </div>

      <div class="form-actions">
        <Button color="white" @click="emit('close')">Cancel</Button>
        <Button color="green" @click="handleSave">{{
          formData.id ? 'Update Shift' : 'Save Shift'
        }}</Button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.shift-form-card {
  background: white;
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  animation: slideDown 0.2s ease-out;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.02);
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.form-grid {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.row-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.full-width {
  width: 100%;
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

.recurring-section {
  background: hsl(from var(--color-neutral) h s 98%);
  padding: 16px;
  border-radius: 8px;
  border: 1px solid var(--border-color);
}

.recurring-options {
  margin-top: 16px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}
</style>
