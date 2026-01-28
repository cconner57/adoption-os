<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import { usePermissions } from '../../../composables/usePermissions'
import type { IVolunteer } from '../../../stores/mockVolunteerData'
import {
  EditorDrawer,
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
  <EditorDrawer
    :is-open="isOpen"
    :title="volunteer ? `Edit ${volunteer.firstName}` : 'Add New Volunteer'"
    save-label="Save Changes"
    :is-saving="isSaving"
    :can-save="canEdit"
    width="800px"
    @close="emit('close')"
    @save="handleSave"
  >
    <template #pre-body>
      <div class="permission-warning" v-if="!canEdit">
        ⚠️ You do not have permission to edit this user.
      </div>
    </template>
    <template #sidebar>
      <SidebarNav
        variant="editor"
        :items="navItems"
        :modelValue="activeTab"
        @update:modelValue="(id) => (activeTab = id as string)"
      />
    </template>
    <template #content>
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
    </template>
  </EditorDrawer>
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
}

/* Scoped styles removed as they are now handled by EditorDrawer */

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
