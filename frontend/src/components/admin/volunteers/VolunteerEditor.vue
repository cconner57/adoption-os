<script setup lang="ts">
import { ref, watch } from 'vue'

import type { IVolunteer } from '../../../stores/mockVolunteerData'
import { formatPhoneInput, formatZipInput } from '../../../utils/formatters'
import {
  Button,
  InputField,
  InputFileUpload,
  InputTextArea,
  Select,
  SidebarNav,
  Toggle,
} from '../../common/ui'
import Availability from '../../volunteer/availability/Availability.vue'
import PositionPreferences from '../../volunteer/preferences/PositionPreferences.vue'

const props = defineProps<{
  volunteer: IVolunteer | null
  isOpen: boolean
  isSaving?: boolean
}>()

const emit = defineEmits(['close', 'save', 'archive'])

const activeTab = ref('general')
const formData = ref<Partial<IVolunteer>>({})

watch(
  () => props.volunteer,
  (newVal) => {
    if (newVal) {
      formData.value = structuredClone(newVal)
    } else {

      formData.value = {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        status: 'active',
        role: 'Tier 1',
        bio: '',
        skills: [],
        positionPreferences: [],
        availability: [],
        allergies: false,
        joinDate: new Date().toISOString().split('T')[0],
      }
    }
  },
  { immediate: true },
)

const navItems = [
  { id: 'general', label: 'General Info', icon: 'user' },
  { id: 'bio', label: 'Bio & Skills', icon: 'fileText' },
  { id: 'preferences', label: 'Preferences', icon: 'star' },
  { id: 'settings', label: 'Settings', icon: 'cog' },
]

function handleSave() {
  emit('save', formData.value)
}
</script>

<template>
  <div v-if="isOpen" class="editor-overlay" @click.self="emit('close')">
    <div class="editor-drawer">

      <header class="editor-header">
        <h2>{{ volunteer ? `Edit ${volunteer.firstName}` : 'Add New Volunteer' }}</h2>
        <div class="header-actions">
          <Button color="white" title="Cancel" @click="emit('close')" />
          <Button color="green" title="Save Changes" @click="handleSave" :loading="isSaving" />
        </div>
      </header>

      <div class="editor-body">
        <SidebarNav
          variant="editor"
          :items="navItems"
          :modelValue="activeTab"
          @update:modelValue="(id) => (activeTab = id as string)"
          style="
            width: 200px;
            background: var(--color-neutral-surface);
            border-right: 1px solid var(--border-color);
            padding-top: 16px;
          "
        />

        <div class="editor-content">

          <div v-if="activeTab === 'general'" class="form-section">
            <h3 class="section-title">General Information</h3>
            <div class="form-grid">
              <div class="full-width" style="margin-bottom: 24px">
                <InputFileUpload
                  label="Profile Photo"
                  v-model="formData.photoUrl"
                  accept="image/*"
                />
              </div>
              <InputField label="First Name" placeholder="First Name" :model-value="formData.firstName || ''" @update:model-value="val => formData.firstName = (val as string)" />
              <InputField label="Last Name" placeholder="Last Name" :model-value="formData.lastName || ''" @update:model-value="val => formData.lastName = (val as string)" />
              <InputField label="Email" placeholder="Email" :model-value="formData.email || ''" @update:model-value="val => formData.email = (val as string)" />
              <div class="row-full">
                <InputField
                  label="Phone"
                  placeholder="Phone"
                  :model-value="formData.phone || ''"
                  @update:model-value="
                    (val: any) => (formData.phone = formatPhoneInput(String(val || '')))
                  "
                  maxlength="14"
                  fullWidth
                />
              </div>

              <div class="row-full">
                <InputField label="Address" placeholder="Address" :model-value="formData.address || ''" @update:model-value="val => formData.address = (val as string)" class="full-width" />
              </div>

              <div class="row-full">
                <InputField label="City" placeholder="City" :model-value="formData.city || ''" @update:model-value="val => formData.city = (val as string)" class="full-width" />
              </div>

              <div class="row-2 full-width">
                <InputField
                  label="Zip"
                  placeholder="Zip"
                  :model-value="formData.zip || ''"
                  @update:model-value="
                    (val: any) => (formData.zip = formatZipInput(String(val || '')))
                  "
                  maxlength="5"
                  fullWidth
                  class="full-width"
                />
                <InputField
                  label="Birthday"
                  placeholder="Birthday"
                  type="date"
                  :model-value="formData.birthday || ''"
                  @update:model-value="val => formData.birthday = (val as string)"
                  fullWidth
                  class="full-width"
                />
              </div>
            </div>

            <h3 class="section-title mt-6">Emergency Contact</h3>
            <div class="form-grid">
              <InputField label="Name" placeholder="Name" :model-value="formData.emergencyContactName || ''" @update:model-value="val => formData.emergencyContactName = (val as string)" />
              <InputField
                label="Phone"
                placeholder="Phone"
                :model-value="formData.emergencyContactPhone || ''"
                @update:model-value="
                  (val: any) =>
                    (formData.emergencyContactPhone = formatPhoneInput(String(val || '')))
                "
                maxlength="14"
                fullWidth
              />
            </div>
          </div>

          <div v-else-if="activeTab === 'bio'" class="form-section">
            <h3 class="section-title">Bio & Skills</h3>
            <div class="stack">
              <InputTextArea label="Bio" placeholder="Bio description..." :model-value="formData.bio || ''" @update:model-value="val => formData.bio = (val as string)" rows="4" />

              <div class="field-group">
                <label class="field-label">Skills (comma separated)</label>
                <input
                  class="simple-input"
                  :value="formData.skills?.join(', ')"
                  @input="
                    (e) =>
                      (formData.skills = (e.target as HTMLInputElement).value
                        .split(',')
                        .map((s) => s.trim())
                        .filter(Boolean))
                  "
                />
              </div>

              <div class="field-group">
                <Toggle :model-value="formData.allergies || false" @update:model-value="val => formData.allergies = val" label="Has Allergies?" labelPosition="left" />
              </div>
            </div>
          </div>

          <div v-else-if="activeTab === 'preferences'" class="form-section">
            <h3 class="section-title">Application Details</h3>
            <div class="stack">
              <InputTextArea label="Why they joined" placeholder="Reason for joining..." :model-value="formData.interestReason || ''" @update:model-value="val => formData.interestReason = (val as string | undefined)" />
              <InputTextArea label="Volunteer Experience" placeholder="Previous experience..." :model-value="formData.volunteerExperience || ''" @update:model-value="val => formData.volunteerExperience = (val as string | undefined)" />

              <h3 class="section-subtitle mt-4">Position Preferences</h3>
              <PositionPreferences :model-value="formData.positionPreferences || []" @update:model-value="val => formData.positionPreferences = val" />

              <h3 class="section-subtitle mt-4">Availability</h3>
              <Availability :model-value="formData.availability || []" @update:model-value="val => formData.availability = val" />
            </div>
          </div>

          <div v-else-if="activeTab === 'settings'" class="form-section">
            <h3 class="section-title">Settings</h3>

            <div class="field-group" style="margin-bottom: 32px">
              <label class="field-label" style="margin-bottom: 8px; display: block"
                >Permission Level</label
              >
              <Select
                :model-value="formData.role || 'Tier 1'"
                @update:model-value="val => formData.role = (val as any)"
                :options="[
                  { label: 'Tier 1 Volunteer', value: 'Tier 1' },
                  { label: 'Tier 2 Volunteer', value: 'Tier 2' },
                  { label: 'Teen Volunteer', value: 'Teen' },
                  { label: 'Admin', value: 'Admin' },
                ]"
              />
            </div>

            <div class="panel" style="margin-bottom: 24px;">
                <div class="panel-header" style="padding: 16px; border-bottom: 1px solid var(--border-color); background: #f9fafb;">
                    <h4 style="margin: 0 0 4px; font-weight: 600; font-size: 1rem;">User Account</h4>
                    <p style="margin: 0; color: var(--text-secondary); font-size: 0.9rem;">
                        Invite this volunteer to create a login account to access the dashboard.
                    </p>
                </div>
                <div class="panel-actions">
                    <Button
                        title="Invite User"
                        color="purple"
                        style="width: 100%; justify-content: center;"
                        :onClick="() => {}"
                    />
                </div>
            </div>

            <div class="panel danger-zone">
              <div class="panel-header">
                <h4>Danger Zone</h4>
                <p>
                  Archive this volunteer to hide them from the active list. You can unarchive them
                  later.
                </p>
              </div>
              <div class="panel-actions">
                <Button
                  color="white"
                  title="Archive Member"
                  @click="emit('archive')"
                  style="
                    color: var(--color-danger);
                    border-color: var(--color-danger);
                    width: 100%;
                    justify-content: center;
                  "
                />
              </div>
            </div>
          </div>
        </div>
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
  background: rgb(0 0 0 / 40%);
  backdrop-filter: blur(2px);
  z-index: 1000;
  display: flex;
  justify-content: flex-end;
}

.editor-drawer {
  width: 800px;
  background: #fff;
  height: 100%;
  box-shadow: -4px 0 24px rgb(0 0 0 / 15%);
  display: flex;
  flex-direction: column;
  animation: slideIn 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }

  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.editor-header {
  padding: 24px 32px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
}

.header-left h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.subtitle {
  color: var(--text-secondary);
  font-size: 0.95rem;
}

.header-actions {
  display: flex;
  gap: 12px;
}

.editor-body {
  flex: 1;
  display: flex;
  overflow: hidden;
  background: var(--bg-secondary);
}

.editor-content {
  flex: 1;
  padding: 32px;
  overflow-y: auto;
  background: #fff;
  border-left: 1px solid var(--border-color);
}

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

.mt-4 {
  margin-top: 16px;
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

.checkbox-field {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: var(--text-primary);

  label {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
  }

  input[type='checkbox'] {
    width: 18px;
    height: 18px;
    accent-color: var(--color-primary);
  }
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
