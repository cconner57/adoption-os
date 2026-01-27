<script setup lang="ts">
import { computed, onMounted,ref } from 'vue'

import { Button, InputField, InputTextArea, Select, Toggle } from '../../common/ui'

const props = defineProps<{
  initialData?: any // eslint-disable-line @typescript-eslint/no-explicit-any
  volunteerOptions?: { label: string; value: string }[]
  currentVolunteerName?: string
}>()

const emit = defineEmits(['close', 'save'])

const formData = ref({
  id: props.initialData?.id || undefined,
  volunteerId: props.initialData?.volunteerId || undefined,
  date: props.initialData?.date || '',
  startTime: props.initialData?.startTime || '',
  endTime: props.initialData?.endTime || '',
  role: props.initialData?.role || 'Feeding/Cleaning',
  status: props.initialData?.status || 'scheduled',
  isRecurring: false, 
  frequency: 'weekly',
  endDate: '',
  isCovering: false,
  coveringName: '',
  coveredBy: [] as (string | number)[],
  coverageNotice: 'more_24h',
  notes: props.initialData?.notes || '',
})

const showCoveredBy = computed(() => {
  return formData.value.status && formData.value.status.startsWith('covered')
})

const roles = ref([
  { label: 'Feeding/Cleaning', value: 'Feeding/Cleaning' },
  { label: 'Cat Socializing', value: 'Cat Socializing' },
  { label: 'Dog Walking', value: 'Dog Walking' },
  { label: 'Customer Support', value: 'Customer Support' },
  { label: 'Adoption Event', value: 'Adoption Event' },
  { label: 'Adoption Event Setup', value: 'Adoption Event Setup' },
  { label: 'AM Vet Pickup', value: 'AM Vet Pickup' },
  { label: 'PM Vet Dropoff', value: 'PM Vet Dropoff' },
])

const coveringOptions = computed(() => {
  if (!props.volunteerOptions) return []
  if (!props.currentVolunteerName) return props.volunteerOptions
  return props.volunteerOptions.filter((opt) => opt.value !== props.currentVolunteerName)
})

onMounted(async () => {
  try {
    const res = await fetch('/v1/shifts/meta/roles')
    if (res.ok) {
      const data = await res.json()
      const counts = data.data.roleCounts || {}

      roles.value.sort((a, b) => {
        const countA = counts[a.value] || 0
        const countB = counts[b.value] || 0
        
        if (countB !== countA) {
          return countB - countA
        }
        return 0
      })
    }
  } catch (e) {
    console.error('Failed to load role stats', e)
  }
})

const statuses = [
  { label: 'Scheduled', value: 'scheduled' },
  { label: 'Completed (+20)', value: 'completed' },
  { label: 'Late (+10)', value: 'late' },
  { label: 'Missed (-50)', value: 'missed' },
  { label: 'Covered >24h notice (-5)', value: 'covered_24h' },
  { label: 'Covered <24h notice (-10)', value: 'covered_less_24h' },
  { label: 'Covered <1h notice (-20)', value: 'covered_less_1h' },
]

const frequencies = [
  { label: 'Weekly', value: 'weekly' },
  { label: 'Bi-Weekly', value: 'biweekly' },
  { label: 'Monthly', value: 'monthly' },
]

function handleSave() {
  const payload = { ...formData.value }

  if (!payload.date || !payload.startTime || !payload.endTime) {
    alert('Please fill in Date, Start Time, and End Time')
    return
  }

  if (payload.isCovering && payload.coveringName) {
    const noticeText = payload.coverageNotice === 'less_24h' ? '<24h notice' : '>24h notice'
    payload.notes = `Covering for ${payload.coveringName} (${noticeText})`
  }

  if (
    showCoveredBy.value &&
    payload.coveredBy &&
    (Array.isArray(payload.coveredBy) ? payload.coveredBy.length > 0 : payload.coveredBy)
  ) {
    const names = Array.isArray(payload.coveredBy)
      ? payload.coveredBy.join(', ')
      : payload.coveredBy
    const prefix = `Covered by ${names}. `
    
    const currentNotes = payload.notes || ''
    if (!currentNotes.startsWith('Covered by')) {
      payload.notes = prefix + currentNotes
    }
  }

  emit('save', payload)

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
    notes: '',
    id: undefined,
    volunteerId: undefined,
    status: 'scheduled',
    coveredBy: [],
    coverageNotice: 'more_24h',
  }
}

function setEndOfYear() {
  const year = new Date().getFullYear()
  const eoy = `${year}-12-31`
  formData.value.endDate = eoy
}
</script>

<template>
  <div class="shift-form-card">
    <div class="form-grid">
      <div class="row-2">
        <InputField label="Date" placeholder="YYYY-MM-DD" type="date" v-model="formData.date" class="full-width" />
        <div class="field-group">
          <label class="field-label">Role</label>
          <Select v-model="formData.role" :options="roles" />
        </div>
      </div>

      <div class="row-2">
        <InputField label="Start Time" placeholder="09:00" type="time" v-model="formData.startTime" />
        <InputField label="End Time" placeholder="17:00" type="time" v-model="formData.endTime" />
      </div>

      <div v-if="formData.id" :class="{ 'row-2': showCoveredBy }">
        <div class="field-group">
          <label class="field-label">Status</label>
          <Select v-model="formData.status" :options="statuses" fullWidth />
        </div>

        <div v-if="showCoveredBy" class="field-group">
          <label class="field-label">Who covered this shift?</label>
          <Select
            placeholder="Select Volunteer"
            v-model="formData.coveredBy"
            :options="props.volunteerOptions || []"
            fullWidth
            multiple
          />
        </div>
      </div>

      <div :class="{ 'row-2': !formData.id }" style="align-items: start">
        <div v-if="!formData.id" class="section-box">
          <div class="field-group">
            <Toggle
              v-model="formData.isRecurring"
              label="Recurring Shift?"
              labelPosition="left"
              fullWidth
            />
          </div>

          <div v-if="formData.isRecurring" class="recurring-options">
            <div class="field-group">
              <label class="field-label">Frequency</label>
              <Select v-model="formData.frequency" :options="frequencies" />
            </div>
            <div class="field-group" style="margin-top: 16px">
              <div class="flex-between">
                <label class="field-label">End Date</label>
                <Button variant="secondary" theme="primary" size="small" @click="setEndOfYear">
                  Until End of Year
                </Button>
              </div>
              <InputField type="date" placeholder="YYYY-MM-DD" v-model="formData.endDate" />
            </div>
          </div>
        </div>

        <div class="section-box">
          <div class="field-group">
            <Toggle
              v-model="formData.isCovering"
              label="Is this covering another volunteer?"
              labelPosition="left"
              fullWidth
            />
          </div>

          <div v-if="formData.isCovering" class="field-group" style="margin-top: 16px">
            <label class="field-label">Who are you covering for?</label>
            <Select
              placeholder="Select Volunteer"
              v-model="formData.coveringName"
              :options="coveringOptions"
              fullWidth
            />
            <div class="field-group" style="margin-top: 12px">
              <label class="field-label">How much notice?</label>
              <Select
                v-model="formData.coverageNotice"
                :options="[
                  { label: 'More than 24h', value: 'more_24h' },
                  { label: 'Less than 24h', value: 'less_24h' },
                ]"
                fullWidth
              />
            </div>
          </div>
        </div>
      </div>

      <div class="field-group">
        <InputTextArea
          label="Notes"
          v-model="formData.notes"
          placeholder="Add any notes about this shift..."
          :rows="3"
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
  background: #fff;
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  animation: slideDown 0.2s ease-out;
  box-shadow: 0 4px 12px rgb(0 0 0 / 2%);
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

.section-box {
  background: var(--color-neutral-surface);
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

.flex-between {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}
</style>
