<script setup lang="ts">
import { ButtonToggle, InputField, InputSelectGroup } from '../../common/ui'
import SettingsCard from './SettingsCard.vue'

defineProps<{
  settings: {
    organization: {
      name: string
      email: string
      phone: string
      timezone: string
    }
    notifications: {
      incidentAlerts: boolean
      newApplicationAlerts: boolean
      emailDigests: string
    }
  }
  accountForm: any // eslint-disable-line @typescript-eslint/no-explicit-any
}>()

defineEmits<{
  'update:settings': [value: Record<string, unknown>]
  'update:accountForm': [value: Record<string, unknown>]
}>()
</script>

<template>
  <div class="settings-grid">
    <SettingsCard title="Organization Profile" icon="🏢">
      <InputField
        label="Shelter Name"
        placeholder="e.g. Happy Tails"
        v-model="settings.organization.name"
        class="mb-4"
      />
      <InputField
        label="Contact Email"
        placeholder="email@example.com"
        type="email"
        v-model="settings.organization.email"
        class="mb-4"
      />
      <div class="row">
        <InputField label="Phone" placeholder="555-0123" v-model="settings.organization.phone" />
        <div class="select-wrapper">
          <label class="u-label">Timezone</label>
          <InputSelectGroup
            :modelValue="settings.organization.timezone"
            @update:modelValue="(val) => (settings.organization.timezone = val as string)"
            :options="['PST', 'EST', 'UTC']"
          />
        </div>
      </div>
    </SettingsCard>

    <SettingsCard title="Account Settings" icon="🔒">
      <p class="text-sm text-gray mb-4">Update your personal account details.</p>
      <div class="row">
        <InputField label="Your Name" placeholder="Name" v-model="accountForm.name" class="mb-4" />
        <InputField
          label="Your Email"
          placeholder="email@example.com"
          type="email"
          v-model="accountForm.email"
          class="mb-4"
        />
      </div>
      <div class="row">
        <InputField
          label="New Password"
          placeholder="Leave blank to keep current"
          type="password"
          v-model="accountForm.password"
          class="mb-4"
        />
        <InputField
          label="Confirm Password"
          placeholder="Confirm New Password"
          type="password"
          v-model="accountForm.confirmPassword"
          class="mb-4"
        />
      </div>
    </SettingsCard>

    <SettingsCard title="Notifications" icon="🔔">
      <div class="setting-row">
        <div class="setting-info">
          <label>Incident Alerts</label>
          <p class="setting-desc">Get notified immediately when an incident is logged.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.notifications.incidentAlerts"
            true-label="On"
            false-label="Off"
            :true-value="true"
            :false-value="false"
          />
        </div>
      </div>
      <div class="setting-row">
        <div class="setting-info">
          <label>New Application Alerts</label>
          <p class="setting-desc">In-app notification when a new volunteer applies.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.notifications.newApplicationAlerts"
            true-label="On"
            false-label="Off"
            :true-value="true"
            :false-value="false"
          />
        </div>
      </div>
      <div class="mt-4">
        <label class="u-label">Report Digest Frequency</label>
        <InputSelectGroup
          :modelValue="settings.notifications.emailDigests"
          @update:modelValue="(val) => (settings.notifications.emailDigests = val as string)"
          :options="['daily', 'weekly', 'monthly', 'off']"
        />
      </div>
    </SettingsCard>

    <SettingsCard title="System" icon="⚙️">
      <div class="setting-row">
        <div class="setting-info">
          <label>Maintenance Mode</label>
          <p class="setting-desc">Disable public access.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle label="" :modelValue="false" true-label="Active" false-label="Inactive" />
        </div>
      </div>
      <div class="system-info mt-4">
        <p><strong>Version:</strong> v1.2.0 (Beta)</p>
        <p><strong>Environment:</strong> Production</p>
      </div>
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

.setting-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px 0;
  border-bottom: 1px solid #f3f4f6;
}

.setting-row:last-child {
  border-bottom: none;
}

.setting-info label {
  display: block;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.setting-desc {
  font-size: 0.85rem;
  color: #6b7280;
  margin: 0;
  line-height: 1.4;
}

.text-sm {
  font-size: 0.875rem;
}

.text-gray {
  color: #6b7280;
}

.mb-4 {
  margin-bottom: 1rem;
}

.mt-4 {
  margin-top: 1rem;
}

.u-label {
  font-size: 0.85rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  display: block;
  color: var(--text-secondary);
}
</style>
