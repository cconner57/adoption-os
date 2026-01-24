<script setup lang="ts">
import { ref, watch } from 'vue'

import type { IRaffleParticipant } from '../../../stores/mockMarketing'
import { Button, InputField } from '../../common/ui'

const props = defineProps<{
  participant: IRaffleParticipant | null
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'save'])

const formData = ref({
  name: '',
  contact: '',
  date: '',
  paymentMethod: '',
  amount: '',
  comments: '',
})

watch(
  () => props.participant,
  (newVal) => {
    if (newVal) {
      formData.value = {
        name: newVal.name || '',
        contact: newVal.contact || '',
        date: newVal.date || '',
        paymentMethod: newVal.paymentMethod || '',
        amount: newVal.amount || '',
        comments: newVal.comments || '',
      }
    } else {
      // Reset for new entry if needed (though we currently only use this for edit)
      formData.value = {
        name: '',
        contact: '',
        date: new Date().toISOString().split('T')[0],
        paymentMethod: '',
        amount: '$10',
        comments: '',
      }
    }
  },
  { immediate: true }
)

const handleClose = () => {
  emit('close')
}

const handleSave = () => {
  emit('save', { ...props.participant, ...formData.value })
  handleClose()
}
</script>

<template>
  <div class="editor-overlay" :class="{ open: isOpen }" @click.self="handleClose">
    <div class="editor-drawer">
      <div class="drawer-header">
        <h2>{{ participant ? `Edit Ticket #${participant.ticketNumber}` : 'Add Participant' }}</h2>
        <div class="header-actions">
          <Button color="white" title="Cancel" @click="handleClose" />
          <Button color="green" title="Save Changes" @click="handleSave" />
        </div>
      </div>

      <div class="drawer-body">
        <form @submit.prevent>
          <div class="form-section">
            <h3>Participant Details</h3>
            <div class="form-row">
              <InputField label="Name" v-model="formData.name" placeholder="Full Name" required />
              <InputField label="Contact Info" v-model="formData.contact" placeholder="Email or Phone" />
            </div>
            <div class="form-row">
               <InputField label="Date" type="date" v-model="formData.date" placeholder="YYYY-MM-DD" />
               <InputField label="Payment Method" v-model="formData.paymentMethod" placeholder="e.g. Cash, Venmo" />
            </div>
            <div class="form-row">
               <InputField label="Amount" v-model="formData.amount" placeholder="$10" />
            </div>
          </div>

          <div class="form-section">
            <h3>Notes</h3>
            <InputField
              label="Comments"
              v-model="formData.comments"
              placeholder="Any additional notes..."
              fullWidth
            />
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
