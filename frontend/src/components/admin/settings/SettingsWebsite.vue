<script setup lang="ts">
import { defineEmits,defineProps } from 'vue'

import { ButtonToggle, Capsules,InputField } from '../../common/ui'
import Button from '../../common/ui/Button.vue'
import SettingsCard from './SettingsCard.vue'

const props = defineProps<{
  settings: any // eslint-disable-line @typescript-eslint/no-explicit-any
  isDemoMode: boolean
}>()

const emit = defineEmits<{
  'update:settings': [value: any] // eslint-disable-line @typescript-eslint/no-explicit-any
  'toggleDemoMode': []
}>()

function addEmail(type: 'volunteer' | 'surrender' | 'adoption') {
  const form = props.settings.forms[type]
  if (form.newEmail && form.newEmail.includes('@') && !form.emails.includes(form.newEmail)) {
    form.emails.push(form.newEmail)
    form.newEmail = ''
  }
}

function removeEmail(type: 'volunteer' | 'surrender' | 'adoption', email: string) {
  const form = props.settings.forms[type]
  form.emails = form.emails.filter((e: string) => e !== email)
}
</script>

<template>
  <div class="settings-grid single-col">
    <SettingsCard title="Development Tools" icon="🚧">
      <div class="setting-row">
        <div class="setting-info">
          <label>Demo Mode</label>
          <p class="setting-desc">Bypass validation and simulate submissions for all forms.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            :modelValue="isDemoMode"
            @update:modelValue="emit('toggleDemoMode')"
            true-label="On"
            false-label="Off"
            :true-value="true"
            :false-value="false"
          />
        </div>
      </div>
    </SettingsCard>

    <SettingsCard title="Applications & Forms" icon="📝">
      <div class="forms-grid">
        <!-- Volunteer Form -->
        <div class="form-config-col">
          <div class="config-header">
            <h4>Volunteer Application</h4>
            <div class="toggle-wrapper small">
              <ButtonToggle
                label=""
                v-model="settings.forms.volunteer.enabled"
                true-label="Active"
                false-label="Disabled"
                :true-value="true"
                :false-value="false"
              />
            </div>
          </div>
          <div class="email-routing">
            <label class="sub-label">Send Notifications To:</label>
            <div class="capsule-list">
              <Capsules
                v-for="email in settings.forms.volunteer.emails"
                :key="email"
                size="sm"
                color="#f3e8ff"
              >
                {{ email }}
                <span class="remove-x" @click="removeEmail('volunteer', email)">×</span>
              </Capsules>
            </div>
            <div class="add-email-row">
              <InputField
                label=""
                placeholder="Add..."
                v-model="settings.forms.volunteer.newEmail"
                @keydown.enter.prevent="addEmail('volunteer')"
              />
              <Button
                title="Add"
                color="white"
                size="small"
                :onClick="() => addEmail('volunteer')"
              />
            </div>
          </div>
        </div>

        <!-- Surrender Form -->
        <div class="form-config-col">
          <div class="config-header">
            <h4>Surrender Request</h4>
            <div class="toggle-wrapper small">
              <ButtonToggle
                label=""
                v-model="settings.forms.surrender.enabled"
                true-label="Active"
                false-label="Disabled"
                :true-value="true"
                :false-value="false"
              />
            </div>
          </div>
          <div class="email-routing">
            <label class="sub-label">Send Notifications To:</label>
            <div class="capsule-list">
              <Capsules
                v-for="email in settings.forms.surrender.emails"
                :key="email"
                size="sm"
                color="#fee2e2"
              >
                {{ email }}
                <span class="remove-x" @click="removeEmail('surrender', email)">×</span>
              </Capsules>
            </div>
            <div class="add-email-row">
              <InputField
                label=""
                placeholder="Add..."
                v-model="settings.forms.surrender.newEmail"
                @keydown.enter.prevent="addEmail('surrender')"
              />
              <Button
                title="Add"
                color="white"
                size="small"
                :onClick="() => addEmail('surrender')"
              />
            </div>
          </div>
        </div>

        <!-- Adoption Form -->
        <div class="form-config-col">
          <div class="config-header">
            <h4>Adoption Application</h4>
            <div class="toggle-wrapper small">
              <ButtonToggle
                label=""
                v-model="settings.forms.adoption.enabled"
                true-label="Active"
                false-label="Disabled"
                :true-value="true"
                :false-value="false"
              />
            </div>
          </div>
          <div class="email-routing">
            <label class="sub-label">Send Notifications To:</label>
            <div class="capsule-list">
              <Capsules
                v-for="email in settings.forms.adoption.emails"
                :key="email"
                size="sm"
                color="#dbeafe"
              >
                {{ email }}
                <span class="remove-x" @click="removeEmail('adoption', email)">×</span>
              </Capsules>
            </div>
            <div class="add-email-row">
              <InputField
                label=""
                placeholder="Add..."
                v-model="settings.forms.adoption.newEmail"
                @keydown.enter.prevent="addEmail('adoption')"
              />
              <Button
                title="Add"
                color="white"
                size="small"
                :onClick="() => addEmail('adoption')"
              />
            </div>
          </div>
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
.settings-grid.single-col {
  grid-template-columns: 1fr;
  max-width: 800px; /* Optional, keep readable width */
}

/* Row & Setting Row (Shared) */
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

/* Forms Grid */
.forms-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
}

.form-config-col {
  background: #f9fafb;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #f3f4f6;
}

.config-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;

  h4 {
    margin: 0;
    font-size: 0.95rem;
    font-weight: 600;
  }
}

.sub-label {
  display: block;
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: #6b7280;
  margin-bottom: 8px;
  font-weight: 600;
}

.capsule-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 12px;
}

.remove-x {
  margin-left: 6px;
  cursor: pointer;
  opacity: 0.5;
}
.remove-x:hover {
  opacity: 1;
}

.add-email-row {
  display: flex;
  gap: 8px;
  align-items: center;
}
</style>
