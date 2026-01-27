<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import { usePermissions } from '../../../composables/usePermissions'
import type { IVolunteer } from '../../../stores/mockVolunteerData'
import {
  Button,
  SidebarNav,
} from '../../common/ui'
import AppDetailsForm from './forms/AppDetailsForm.vue'
import BioSkillsForm from './forms/BioSkillsForm.vue'
import GeneralInfoForm from './forms/GeneralInfoForm.vue'
import SettingsForm from './forms/SettingsForm.vue'

const { canEditUser } = usePermissions()

const props = defineProps<{
  volunteer: IVolunteer | null
  isOpen: boolean
  isSaving?: boolean
}>()

const canEdit = computed(() => {
  if (!props.volunteer) return true // Creating new user
  return canEditUser(props.volunteer.role)
})

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

function updateFormData(updates: Partial<IVolunteer>) {
  formData.value = updates
}
</script>

<template>
  <div v-if="isOpen" class="editor-overlay" @click.self="emit('close')">
    <div class="editor-drawer">

      <header class="editor-header">
        <h2>{{ volunteer ? `Edit ${volunteer.firstName}` : 'Add New Volunteer' }}</h2>
        <div class="header-actions">
          <Button color="white" title="Cancel" @click="emit('close')" />
          <Button color="green" title="Save Changes" @click="handleSave" :loading="isSaving" :disabled="!canEdit" />
        </div>
      </header>

      <div class="permission-warning" v-if="!canEdit">
         ⚠️ You do not have permission to edit this user.
      </div>

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
          <GeneralInfoForm
            v-if="activeTab === 'general'"
            :form-data="formData"
            :can-edit="canEdit"
            @update:form-data="updateFormData"
          />

          <BioSkillsForm
            v-if="activeTab === 'bio'"
            :form-data="formData"
            :can-edit="canEdit"
            @update:form-data="updateFormData"
          />

          <AppDetailsForm
            v-if="activeTab === 'preferences'"
            :form-data="formData"
            :can-edit="canEdit"
            @update:form-data="updateFormData"
          />

          <SettingsForm
            v-if="activeTab === 'settings'"
            :form-data="formData"
            :can-edit="canEdit"
            @update:form-data="updateFormData"
            @archive="emit('archive')"
          />
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

.permission-warning {
  background: #fefce8;
  color: #854d0e;
  padding: 12px 32px;
  font-weight: 600;
  border-bottom: 1px solid #fde047;
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>
