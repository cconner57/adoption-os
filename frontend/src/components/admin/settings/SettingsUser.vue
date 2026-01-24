<script setup lang="ts">
import { onMounted,ref } from 'vue'

import { usePushNotifications } from '../../../composables/usePushNotifications'
import { useAuthStore } from '../../../stores/auth'
import { formatPhoneInput, formatZipInput } from '../../../utils/formatters'
import {
  InputField,
  InputFileUpload,
  InputTextArea,
  Toggle,
} from '../../common/ui'
import Availability from '../../volunteer/availability/Availability.vue'
import PositionPreferences from '../../volunteer/preferences/PositionPreferences.vue'
import SettingsCard from './SettingsCard.vue'

const authStore = useAuthStore()
const { isSupported, isSubscribed, subscribe, unsubscribe, sendTestNotification, checkSubscription, loading: pushLoading, error: pushError } = usePushNotifications()

onMounted(async () => {
  await checkSubscription()
})

const handleToggle = async (val: boolean) => {
    console.log('AdoptionOS: Toggled to', val)
    if (val) {
        await subscribe()
    } else {
        await unsubscribe()
    }
}


const formData = ref({
  // General
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  address: '',
  city: '',
  zip: '',
  birthday: '',
  photoUrl: '',

  // Emergency
  emergencyContactName: '',
  emergencyContactPhone: '',

  // Bio & Skills
  bio: '',
  skills: [] as string[],
  allergies: false,

  // Preferences
  interestReason: '',
  volunteerExperience: '',
  positionPreferences: [] as string[],
  availability: [] as string[],

  role: 'Volunteer'
})

onMounted(() => {
  if (authStore.user) {
    const splitName = (authStore.user.Name || '').split(' ')
    formData.value.firstName = splitName[0] || ''
    formData.value.lastName = splitName.slice(1).join(' ') || ''
    formData.value.email = authStore.user.Email || ''

    // Default dummy data
    formData.value.bio = "I love animals!"
    formData.value.skills = ['Walking', 'Cleaning']
    formData.value.role = authStore.user.Role || 'Volunteer'
  }
})

defineExpose({
  save: async () => {
    console.log("Saving user settings", formData.value)
    await new Promise(r => setTimeout(r, 500))
    return true
  }
})

</script>

<template>
  <div class="settings-grid">
    <SettingsCard title="General Information" icon="user">
        <div class="full-width mb-6">
            <InputFileUpload
                label="Profile Photo"
                v-model="formData.photoUrl"
                accept="image/*"
            />
        </div>
        <div class="row">
            <InputField label="First Name" placeholder="First Name" v-model="formData.firstName" class="mb-4" />
            <InputField label="Last Name" placeholder="Last Name" v-model="formData.lastName" class="mb-4" />
        </div>
        <div class="row">
            <InputField label="Email" placeholder="Email" v-model="formData.email" class="mb-4" />
            <InputField
                label="Phone"
                placeholder="Phone"
                v-model="formData.phone"
                @update:modelValue="(val: any) => formData.phone = formatPhoneInput(String(val || ''))"
                maxlength="14"
                class="mb-4"
            />
        </div>
        <div class="row">
            <InputField label="Address" placeholder="Address" v-model="formData.address" class="mb-4" />
            <InputField label="City" placeholder="City" v-model="formData.city" class="mb-4" />
        </div>
        <div class="row">
            <InputField
                label="Zip"
                placeholder="Zip"
                v-model="formData.zip"
                @update:modelValue="(val: any) => formData.zip = formatZipInput(String(val || ''))"
                maxlength="5"
                class="mb-4"
            />
            <InputField
                label="Birthday"
                placeholder="Birthday"
                type="date"
                v-model="formData.birthday"
                class="mb-4"
            />
        </div>
    </SettingsCard>

    <SettingsCard title="Notifications" icon="bell">
      <div v-if="isSupported">
        <div class="notification-toggle-row">
            <Toggle
                :modelValue="isSubscribed"
                :disabled="pushLoading"
                :label="`Push Notifications: ${isSubscribed ? 'Enabled' : 'Disabled'}`"
                labelPosition="left"
                @update:modelValue="handleToggle"
            />
            <div v-if="pushLoading" class="loading-spinner"></div>

            <button
                class="test-btn"
                :disabled="!isSubscribed || pushLoading"
                @click="sendTestNotification"
                title="Send Test Notification"
            >
                🔔 Test
            </button>
        </div>
        <p v-if="pushError" class="error-text mb-2">{{ pushError }}</p>
        <p class="help-text">
            Receive alerts for urgent needs, new applications, and campaign milestones.
        </p>
      </div>
      <div v-else>
        <p class="error-text">Push notifications are not supported on this device/browser.</p>
      </div>
    </SettingsCard>

    <SettingsCard title="Emergency Contact" icon="phone">
        <div class="row">
            <InputField label="Name" placeholder="Name" v-model="formData.emergencyContactName" class="mb-4" />
            <InputField
                label="Phone"
                placeholder="Phone"
                v-model="formData.emergencyContactPhone"
                @update:modelValue="(val: any) => formData.emergencyContactPhone = formatPhoneInput(String(val || ''))"
                maxlength="14"
                class="mb-4"
            />
        </div>
    </SettingsCard>

    <SettingsCard title="Bio & Skills" icon="fileText">
        <InputTextArea label="Bio" placeholder="Bio description..." v-model="formData.bio" rows="4" class="mb-4" />

        <div class="mb-4">
            <label class="u-label">Skills (comma separated)</label>
            <input
                class="simple-input"
                :value="formData.skills.join(', ')"
                @input="(e) => formData.skills = (e.target as HTMLInputElement).value.split(',').map(s => s.trim()).filter(Boolean)"
            />
        </div>

        <div class="mb-4">
            <Toggle v-model="formData.allergies" label="Has Allergies?" labelPosition="left" />
        </div>
    </SettingsCard>

    <SettingsCard title="Preferences" icon="star">
        <InputTextArea label="Why you joined" placeholder="Reason for joining..." v-model="formData.interestReason" class="mb-4" />
        <InputTextArea label="Volunteer Experience" placeholder="Previous experience..." v-model="formData.volunteerExperience" class="mb-4" />

        <h3 class="section-subtitle mt-4">Position Preferences</h3>
        <PositionPreferences v-model="formData.positionPreferences" class="mb-6" />

        <h3 class="section-subtitle mt-4">Availability</h3>
        <Availability v-model="formData.availability" />
    </SettingsCard>
  </div>
</template>

<style scoped>
.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 24px;
  align-items: start;
}

.row {
  display: flex;
  gap: 16px;
}

.row > * {
  flex: 1;
}

.mb-4 {
  margin-bottom: 16px;
}

.mb-6 {
  margin-bottom: 24px;
}

.mt-4 {
    margin-top: 16px;
}

.full-width {
    width: 100%;
}

.u-label {
  font-size: 0.85rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  display: block;
  color: var(--text-secondary);
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

.section-subtitle {
  font-size: 1rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.notification-toggle-row {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 8px;
}

.help-text {
    font-size: 0.9rem;
    color: var(--text-secondary);
    margin: 0;
}

.error-text {
    font-size: 0.9rem;
    color: var(--color-danger);
    margin: 0;
}

.test-btn {
    background: #fff;
    border: 1px solid var(--color-primary);
    border-radius: 6px;
    padding: 6px 12px;
    font-size: 0.85rem;
    cursor: pointer;
    font-weight: 600;
    color: var(--color-primary);
    transition: all 0.2s;
    margin-left: auto;
    display: inline-flex;
    align-items: center;
    gap: 6px;
    box-shadow: 0 1px 2px rgb(0 0 0 / 5%);
}

.test-btn:hover:not(:disabled) {
    background: var(--color-primary);
    color: #fff;
}

.test-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
    background: #f3f4f6;
    border-color: #d1d5db;
    color: #9ca3af;
    box-shadow: none;
}
</style>
