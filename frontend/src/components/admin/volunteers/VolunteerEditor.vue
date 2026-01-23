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

// Init data
watch(
  () => props.volunteer,
  (newVal) => {
    if (newVal) {
      formData.value = JSON.parse(JSON.stringify(newVal))
    } else {
      // Defaults for new volunteer
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
  { id: 'general', label: 'General Info', icon: '👤' },
  { id: 'bio', label: 'Bio & Skills', icon: '📝' },
  { id: 'preferences', label: 'Preferences', icon: '⭐' },
  { id: 'settings', label: 'Settings', icon: '⚙️' },
]

function handleSave() {
  emit('save', formData.value)
}
</script>

<template>
  <div v-if="isOpen" class="editor-overlay" @click.self="emit('close')">
    <div class="editor-drawer">
      <!-- Header -->
      <header class="editor-header">
        <h2>{{ volunteer ? `Edit ${volunteer.firstName}` : 'Add New Volunteer' }}</h2>
        <div class="header-actions">
          <Button color="white" title="Cancel" @click="emit('close')" />
          <Button color="green" title="Save Changes" @click="handleSave" :loading="isSaving" />
        </div>
      </header>

      <!-- Body -->
      <div class="editor-body">
        <SidebarNav
          variant="editor"
          :items="navItems"
          :modelValue="activeTab"
          @update:modelValue="(id) => (activeTab = id as string)"
          style="
            width: 200px;
            background: hsl(from var(--color-neutral) h s 98%);
            border-right: 1px solid var(--border-color);
            padding-top: 16px;
          "
        />

        <div class="editor-content">
          <!-- General Tab -->
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
              <InputField label="First Name" v-model="formData.firstName" />
              <InputField label="Last Name" v-model="formData.lastName" />
              <InputField label="Email" v-model="formData.email" />
              <div class="row-full">
                <InputField
                  label="Phone"
                  :model-value="formData.phone"
                  @update:model-value="
                    (val: any) => (formData.phone = formatPhoneInput(String(val || '')))
                  "
                  maxlength="14"
                  fullWidth
                />
              </div>

              <div class="row-full">
                <InputField label="Address" v-model="formData.address" class="full-width" />
              </div>

              <!-- City has its own row for more space -->
              <div class="row-full">
                <InputField label="City" v-model="formData.city" class="full-width" />
              </div>

              <!-- Zip moved next to Birthday -->
              <div class="row-2 full-width">
                <InputField
                  label="Zip"
                  :model-value="formData.zip"
                  @update:model-value="
                    (val: any) => (formData.zip = formatZipInput(String(val || '')))
                  "
                  maxlength="5"
                  fullWidth
                  class="full-width"
                />
                <InputField
                  label="Birthday"
                  type="date"
                  v-model="formData.birthday"
                  fullWidth
                  class="full-width"
                />
              </div>
            </div>

            <h3 class="section-title mt-6">Emergency Contact</h3>
            <div class="form-grid">
              <InputField label="Name" v-model="formData.emergencyContactName" />
              <InputField
                label="Phone"
                :model-value="formData.emergencyContactPhone"
                @update:model-value="
                  (val: any) =>
                    (formData.emergencyContactPhone = formatPhoneInput(String(val || '')))
                "
                maxlength="14"
                fullWidth
              />
            </div>
          </div>

          <!-- Bio & Skills Tab -->
          <div v-else-if="activeTab === 'bio'" class="form-section">
            <h3 class="section-title">Bio & Skills</h3>
            <div class="stack">
              <InputTextArea label="Bio" v-model="formData.bio" rows="4" />

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
                <Toggle v-model="formData.allergies" label="Has Allergies?" labelPosition="left" />
              </div>
            </div>
          </div>

          <!-- Preferences Tab -->
          <div v-else-if="activeTab === 'preferences'" class="form-section">
            <h3 class="section-title">Application Details</h3>
            <div class="stack">
              <InputTextArea label="Why they joined" v-model="formData.interestReason" />
              <InputTextArea label="Volunteer Experience" v-model="formData.volunteerExperience" />

              <h3 class="section-subtitle mt-4">Position Preferences</h3>
              <PositionPreferences v-model="formData.positionPreferences" />

              <h3 class="section-subtitle mt-4">Availability</h3>
              <Availability v-model="formData.availability" />
            </div>
          </div>

          <!-- Settings Tab -->
          <div v-else-if="activeTab === 'settings'" class="form-section">
            <h3 class="section-title">Settings</h3>

            <div class="field-group" style="margin-bottom: 32px">
              <label class="field-label" style="margin-bottom: 8px; display: block"
                >Permission Level</label
              >
              <Select
                v-model="formData.role"
                :options="[
                  { label: 'Tier 1 Volunteer', value: 'Tier 1' },
                  { label: 'Tier 2 Volunteer', value: 'Tier 2' },
                  { label: 'Teen Volunteer', value: 'Teen' },
                  { label: 'Admin', value: 'Admin' },
                ]"
              />
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
  background: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(2px);
  z-index: 1000;
  display: flex;
  justify-content: flex-end;
}

.editor-drawer {
  width: 800px; /* Match PetEditor width logic roughly, or slightly smaller */
  background: #fff;
  height: 100%;
  box-shadow: -4px 0 24px rgba(0, 0, 0, 0.15);
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
  background: white;
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
  background: white;
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

/* Reused input styles (consider making global shared CSS) */
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
  background: white;
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

/* Force deep width override for Inputs in grid cells */
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
  border-color: #fca5a5; /* red-300 */
}

.danger-zone .panel-header {
  background: #fef2f2; /* red-50 */
  padding: 16px;
  border-bottom: 1px solid #fca5a5;
}

.danger-zone h4 {
  color: #991b1b; /* red-800 */
  font-weight: 600;
  margin-bottom: 4px;
}

.danger-zone p {
  color: #b91c1c; /* red-700 */
  font-size: 0.9rem;
  margin: 0;
}

.panel-actions {
  padding: 16px;
  background: white;
}
</style>
