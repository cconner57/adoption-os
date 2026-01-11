<script setup lang="ts">
import { ref, computed } from 'vue'
import { Button, InputField, Select, Toggle } from '../../common/ui'

const props = defineProps<{}>()

const emit = defineEmits(['close', 'save'])

const formData = ref({
  date: '',
  startTime: '',
  endTime: '',
  role: 'Feeding/Cleaning',
  isRecurring: false,
  frequency: 'weekly',
  endDate: '',
})

const roles = [
  { label: 'Feeding/Cleaning', value: 'Feeding/Cleaning' },
  { label: 'Cat Socializing', value: 'Cat Socializing' },
  { label: 'Dog Walking', value: 'Dog Walking' },
  { label: 'Customer Support', value: 'Customer Support' },
  { label: 'Adoption Event', value: 'Adoption Event' },
]

const frequencies = [
  { label: 'Weekly', value: 'weekly' },
  { label: 'Bi-Weekly', value: 'biweekly' },
  { label: 'Monthly', value: 'monthly' },
]

function handleSave() {
  emit('save', { ...formData.value })
  // Reset form or keep? Resetting is better.
  formData.value = {
    date: '',
    startTime: '',
    endTime: '',
    role: 'Feeding/Cleaning',
    isRecurring: false,
    frequency: 'weekly',
    endDate: '',
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

      <div class="recurring-section">
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

      <div class="form-actions">
        <Button color="white" title="Cancel" @click="emit('close')" />
        <Button color="green" title="Save Shift" @click="handleSave" />
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
