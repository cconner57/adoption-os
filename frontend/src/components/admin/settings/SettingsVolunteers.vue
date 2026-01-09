<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import { InputField, ButtonToggle, InputSelectGroup } from '../../common/ui'
import SettingsCard from './SettingsCard.vue'

defineProps<{
  settings: any
}>()

const emit = defineEmits<{
  (e: 'update:settings', value: any): void
}>()
</script>

<template>
  <div class="settings-grid">
    <SettingsCard title="Volunteer Management" icon="🤝">
      <div class="setting-row">
        <div class="setting-info">
          <label>Enable Gamification</label>
          <p class="setting-desc">Show badges, streaks, and leaderboards to volunteers.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.volunteers.enableGamification"
            true-label="On"
            false-label="Off"
            :true-value="true"
            :false-value="false"
          />
        </div>
      </div>
      <div class="setting-row">
        <div class="setting-info">
          <label>Teen Volunteers</label>
          <p class="setting-desc">Allow applications from users under 18.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.volunteers.allowTeenVolunteers"
            true-label="Allowed"
            false-label="Blocked"
            :true-value="true"
            :false-value="false"
          />
        </div>
      </div>
      <hr class="my-4" />
      <div class="mb-4">
        <label class="u-label">Shift Reminder</label>
        <InputSelectGroup
          :modelValue="settings.volunteers.shiftReminderHours"
          @update:modelValue="(val) => (settings.volunteers.shiftReminderHours = val as string)"
          :options="[
            { label: '12 Hours Before', value: '12' },
            { label: '24 Hours Before', value: '24' },
            { label: '48 Hours Before', value: '48' },
          ]"
        />
      </div>
      <InputField
        label="Hours for Tier 1 Promotion"
        type="number"
        placeholder="20"
        v-model="settings.volunteers.minHoursForTier1"
      />
    </SettingsCard>

    <SettingsCard title="Volunteer Form" icon="🏷️">
      <div class="setting-row">
        <div class="setting-info">
          <label>Require References</label>
          <p class="setting-desc">Ask for 2 personal references on application.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle label="" :modelValue="true" true-label="Yes" false-label="No" />
        </div>
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

.my-4 {
  margin-top: 1rem;
  margin-bottom: 1rem;
  border: 0;
  border-top: 1px solid #e5e7eb;
}
.mb-4 {
  margin-bottom: 1rem;
}
.u-label {
  font-size: 0.85rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  display: block;
  color: var(--text-secondary);
}
</style>
