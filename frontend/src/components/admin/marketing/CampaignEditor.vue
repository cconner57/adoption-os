<script setup lang="ts">
import { ref, watch } from 'vue'

import type { ICampaign } from '../../../stores/mockMarketing'
import { Button, InputDate, InputField, Select } from '../../common/ui'

const statusOptions = [
  { label: 'Draft', value: 'draft' },
  { label: 'Active', value: 'active' },
  { label: 'Completed', value: 'completed' }
]

const props = defineProps<{
  campaign: ICampaign | null
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'save'])

const formData = ref<{
  prize: string
  ticketPrice: string | number
  goal: string | number
  startDate: string
  endDate: string
  status: string
  metric: string
}>({
  prize: '',
  ticketPrice: '',
  goal: '',
  startDate: '',
  endDate: '',
  status: 'draft',
  metric: 'entries',
})

watch(
  () => props.campaign,
  (newVal) => {
    if (newVal) {
      formData.value = {
        prize: newVal.prize || '',
        ticketPrice: newVal.ticketPrice || '',
        goal: newVal.goal || '',
        startDate: newVal.startDate || '',
        endDate: newVal.endDate || '',
        status: newVal.status || 'draft',
        metric: newVal.metric || 'entries',
      }
    } else {
      formData.value = {
        prize: '',
        ticketPrice: '',
        goal: '',
        startDate: new Date().toISOString().split('T')[0],
        endDate: '',
        status: 'draft',
        metric: 'entries',
      }
    }
  },
  { immediate: true }
)

const handleClose = () => {
  emit('close')
}

const handleSave = () => {
  const base = props.campaign || {}
  emit('save', {
    ...base,
    prize: formData.value.prize,
    ticketPrice: Number(formData.value.ticketPrice),
    goal: String(formData.value.goal),
    startDate: formData.value.startDate,
    endDate: formData.value.endDate,
    metric: formData.value.metric,
    status: formData.value.status as 'draft' | 'active' | 'completed'
  })
  handleClose()
}
</script>

<template>
  <div class="editor-overlay" :class="{ open: isOpen }" @click.self="handleClose">
    <div class="editor-drawer">
      <div class="drawer-header">
        <h2>{{ campaign ? 'Edit Campaign Details' : 'New Campaign' }}</h2>
        <div class="header-actions">
          <Button color="white" title="Cancel" @click="handleClose" />
          <Button color="green" title="Save Changes" @click="handleSave" />
        </div>
      </div>

      <div class="drawer-body">
        <form @submit.prevent>
          <div class="form-section">
            <h3>Campaign Settings</h3>
            <div class="form-row">
               <InputField label="Prize Name" v-model="formData.prize" placeholder="e.g. Gift Basket" fullWidth />
            </div>
            <div class="form-row">
               <InputField label="Ticket Price ($)" type="number" v-model="formData.ticketPrice" placeholder="10" fullWidth />
               <InputField label="Goal" type="number" v-model="formData.goal" placeholder="150" fullWidth />
            </div>
            <div class="form-row">
               <Select
                  label="Metric"
                  v-model="formData.metric"
                  :options="[
                      { label: 'Entries', value: 'entries' },
                      { label: 'Dollars Raised', value: 'dollars' }
                  ]"
                  fullWidth
               />
               <Select
                  label="Status"
                  v-model="formData.status"
                  :options="statusOptions"
                  fullWidth
               />
            </div>
            <div class="form-row">
               <InputDate label="Start Date" v-model="formData.startDate" placeholder="Select Start Date" fullWidth />
               <InputDate label="End Date" v-model="formData.endDate" placeholder="Select End Date" fullWidth />
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.editor-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgb(0 0 0 / 50%);
  display: flex;
  justify-content: flex-end;
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.3s;
  z-index: 1000;
}

.editor-overlay.open {
  opacity: 1;
  pointer-events: auto;
}

.editor-drawer {
  width: 500px;
  max-width: 90vw;
  background: #fff;
  height: 100%;
  box-shadow: -4px 0 20px rgb(0 0 0 / 10%);
  display: flex;
  flex-direction: column;
  transform: translateX(100%);
  transition: transform 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.editor-overlay.open .editor-drawer {
  transform: translateX(0);
}

.drawer-header {
  padding: 24px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
  z-index: 10;
}

.drawer-header h2 {
  font-size: 1.25rem;
  font-weight: 700;
  margin: 0;
  color: var(--text-primary);
}

.header-actions {
  display: flex;
  gap: 12px;
}

.drawer-body {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
  background: var(--text-inverse);
}

.form-section {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  border: 1px solid var(--border-color);
  margin-bottom: 24px;
}

.form-section h3 {
  margin: 0 0 16px;
  font-size: 1rem;
  color: var(--text-secondary);
  font-weight: 600;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 8px;
}

.form-row {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
}

.form-row > * {
  flex: 1;
}
</style>
