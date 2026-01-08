<script setup lang="ts">
import { ref, computed } from 'vue'
import { InputField, ButtonToggle, InputSelectGroup, Capsules } from '../../components/common/ui'
import Button from '../../components/common/ui/Button.vue'

const saving = ref(false)
const showToast = ref(false)
const activeCategory = ref<string | null>(null)

const categories = [
  {
    id: 'general',
    label: 'General',
    icon: '⚙️',
    description: 'Organization profile, notifications, system status.',
  },
  {
    id: 'website',
    label: 'Website',
    icon: '🌐',
    description: 'Manage public facing application forms and email routing.',
  },
  {
    id: 'overview',
    label: 'Overview',
    icon: '📊',
    description: 'Dashboard widgets and visibility settings.',
  },
  {
    id: 'volunteers',
    label: 'Volunteers',
    icon: '🤝',
    description: 'Gamification, shift rules, and application settings.',
  },
  { id: 'calendar', label: 'Calendar', icon: '📅', description: 'Scheduling preferences.' },
  {
    id: 'pets',
    label: 'Pet Records',
    icon: '🐾',
    description: 'Default intake fields and species.',
  },
  {
    id: 'health',
    label: 'Medical',
    icon: '🩺',
    description: 'Vaccine protocols and vet contacts.',
  },
  {
    id: 'transport',
    label: 'Transport',
    icon: '🚙',
    description: 'Vehicle fleet and driver requirements.',
  },
  {
    id: 'timesheets',
    label: 'Timesheets',
    icon: '⏱️',
    description: 'Clock-in/out rules and roundings.',
  },
  {
    id: 'messages',
    label: 'Messages',
    icon: '✉️',
    description: 'Auto-replies and signature settings.',
  },
  {
    id: 'donations',
    label: 'Donations',
    icon: '💰',
    description: 'Payment gateways and receipt templates.',
  },
  {
    id: 'inventory',
    label: 'Inventory',
    icon: '📦',
    description: 'Low stock alerts and categories.',
  },
  {
    id: 'marketing',
    label: 'Marketing',
    icon: '📣',
    description: 'Social media integrations and branding.',
  },
  {
    id: 'kiosk',
    label: 'Intake Kiosk',
    icon: '🖥️',
    description: 'Kiosk display and timeout settings.',
  },
  {
    id: 'kennel',
    label: 'Smart Kennel Cards',
    icon: '🏷️',
    description: 'Digital cage card layouts.',
  },
  {
    id: 'events',
    label: 'Event Signage',
    icon: '🎪',
    description: 'Digital signage playlist management.',
  },
]

const settings = ref({
  organization: {
    name: 'Happy Tails Shelter',
    email: 'contact@happytails.org',
    phone: '555-0123',
    timezone: 'PST',
  },
  volunteers: {
    enableGamification: true,
    autoApproveShifts: false,
    minHoursForTier1: 20,
    shiftReminderHours: '24',
    allowTeenVolunteers: true,
  },
  notifications: {
    emailDigests: 'daily',
    incidentAlerts: true,
    newApplicationAlerts: true,
  },
  forms: {
    volunteer: {
      enabled: true,
      emails: ['director@happytails.org', 'volunteers@happytails.org'],
      newEmail: '',
    },
    surrender: {
      enabled: true,
      emails: ['director@happytails.org', 'intake@happytails.org'],
      newEmail: '',
    },
    adoption: {
      enabled: true,
      emails: ['adoptions@happytails.org'],
      newEmail: '',
    },
  },
  overview: {
    showRecentActivity: true,
    showPendingTasks: true,
    showStatsGraph: true,
  },
})

function addEmail(type: 'volunteer' | 'surrender' | 'adoption') {
  const form = settings.value.forms[type]
  if (form.newEmail && form.newEmail.includes('@') && !form.emails.includes(form.newEmail)) {
    form.emails.push(form.newEmail)
    form.newEmail = ''
  }
}

function removeEmail(type: 'volunteer' | 'surrender' | 'adoption', email: string) {
  const form = settings.value.forms[type]
  form.emails = form.emails.filter((e) => e !== email)
}

function saveSettings() {
  saving.value = true
  setTimeout(() => {
    saving.value = false
    activeCategory.value = null // Optional: go back to home on save? Maybe better to stay.
    showToast.value = true
    setTimeout(() => {
      showToast.value = false
    }, 3000)
  }, 800)
}

const activeCategoryLabel = computed(
  () => categories.find((t) => t.id === activeCategory.value)?.label,
)
</script>

<template>
  <div class="admin-page">
    <div class="page-header">
      <div class="header-left">
        <!-- Back Button (only shown when category active) -->
        <button v-if="activeCategory" class="back-btn" @click="activeCategory = null">
          ← Back
        </button>
        <h1>{{ activeCategory ? activeCategoryLabel : 'Settings' }}</h1>
      </div>

      <button v-if="activeCategory" class="save-btn" :disabled="saving" @click="saveSettings">
        {{ saving ? 'Saving...' : 'Save Changes' }}
      </button>
    </div>

    <!-- Toast Notification -->
    <div v-if="showToast" class="toast-notification">✅ Settings saved successfully!</div>

    <!-- HOME GRID VIEW -->
    <div v-if="!activeCategory" class="categories-grid">
      <button
        v-for="cat in categories"
        :key="cat.id"
        class="category-card"
        @click="activeCategory = cat.id"
      >
        <div class="cat-icon-wrapper">
          <span class="cat-icon">{{ cat.icon }}</span>
        </div>
        <div class="cat-details">
          <h3>{{ cat.label }}</h3>
          <p>{{ cat.description }}</p>
        </div>
        <div class="cat-arrow">→</div>
      </button>
    </div>

    <!-- DETAIL VIEW -->
    <div v-else class="settings-content">
      <!-- GENERAL TAB -->
      <div v-if="activeCategory === 'general'" class="settings-grid">
        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">🏢</div>
            <h3>Organization Profile</h3>
          </div>
          <div class="card-content">
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
              <InputField
                label="Phone"
                placeholder="555-0123"
                v-model="settings.organization.phone"
              />
              <div class="select-wrapper">
                <label class="u-label">Timezone</label>
                <InputSelectGroup
                  :modelValue="settings.organization.timezone"
                  @update:modelValue="(val) => (settings.organization.timezone = val as string)"
                  :options="['PST', 'EST', 'UTC']"
                />
              </div>
            </div>
          </div>
        </section>

        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">🔔</div>
            <h3>Notifications</h3>
          </div>
          <div class="card-content">
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
          </div>
        </section>

        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">⚙️</div>
            <h3>System</h3>
          </div>
          <div class="card-content">
            <div class="setting-row">
              <div class="setting-info">
                <label>Maintenance Mode</label>
                <p class="setting-desc">Disable public access.</p>
              </div>
              <div class="toggle-wrapper">
                <ButtonToggle
                  label=""
                  :modelValue="false"
                  true-label="Active"
                  false-label="Inactive"
                />
              </div>
            </div>
            <div class="system-info mt-4">
              <p><strong>Version:</strong> v1.2.0 (Beta)</p>
              <p><strong>Environment:</strong> Production</p>
            </div>
          </div>
        </section>
      </div>

      <!-- WEBSITE TAB -->
      <div v-if="activeCategory === 'website'" class="settings-grid single-col">
        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">📝</div>
            <h3>Applications & Forms</h3>
          </div>
          <div class="card-content">
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
          </div>
        </section>
      </div>

      <!-- VOLUNTEERS TAB -->
      <div v-if="activeCategory === 'volunteers'" class="settings-grid">
        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">🤝</div>
            <h3>Volunteer Management</h3>
          </div>
          <div class="card-content">
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
                @update:modelValue="
                  (val) => (settings.volunteers.shiftReminderHours = val as string)
                "
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
          </div>
        </section>

        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">🏷️</div>
            <h3>Volunteer Form</h3>
          </div>
          <div class="card-content">
            <div class="setting-row">
              <div class="setting-info">
                <label>Require References</label>
                <p class="setting-desc">Ask for 2 personal references on application.</p>
              </div>
              <div class="toggle-wrapper">
                <ButtonToggle label="" :modelValue="true" true-label="Yes" false-label="No" />
              </div>
            </div>
          </div>
        </section>
      </div>

      <!-- OVERVIEW TAB -->
      <div v-if="activeCategory === 'overview'" class="settings-grid">
        <section class="card settings-card">
          <div class="card-header">
            <div class="icon-box">📊</div>
            <h3>Dashboard Configuration</h3>
          </div>
          <div class="card-content">
            <div class="setting-row">
              <div class="setting-info">
                <label>Show Recent Activity</label>
                <p class="setting-desc">Display recent system actions on the dashboard.</p>
              </div>
              <div class="toggle-wrapper">
                <ButtonToggle
                  label=""
                  v-model="settings.overview.showRecentActivity"
                  true-label="Show"
                  false-label="Hide"
                />
              </div>
            </div>
            <div class="setting-row">
              <div class="setting-info">
                <label>Show Pending Tasks</label>
                <p class="setting-desc">Show your to-do list widget.</p>
              </div>
              <div class="toggle-wrapper">
                <ButtonToggle
                  label=""
                  v-model="settings.overview.showPendingTasks"
                  true-label="Show"
                  false-label="Hide"
                />
              </div>
            </div>
            <div class="setting-row">
              <div class="setting-info">
                <label>Show Stats Graph</label>
                <p class="setting-desc">Display the monthly adoption trends graph.</p>
              </div>
              <div class="toggle-wrapper">
                <ButtonToggle
                  label=""
                  v-model="settings.overview.showStatsGraph"
                  true-label="Show"
                  false-label="Hide"
                />
              </div>
            </div>
          </div>
        </section>
      </div>

      <!-- PLACEHOLDER TABS -->
      <div
        v-if="
          activeCategory &&
          !['general', 'website', 'volunteers', 'overview'].includes(activeCategory)
        "
        class="placeholder-content"
      >
        <div class="placeholder-icon">🛠️</div>
        <h3>{{ categories.find((t) => t.id === activeCategory)?.label }} Settings</h3>
        <p>This section is coming soon.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
  position: relative;
  height: 100%;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-shrink: 0;
  min-height: 48px;

  h1 {
    font-size: 1.8rem;
    color: var(--font-color-dark);
    margin: 0;
  }
}

.header-left {
  display: flex;
  gap: 16px;
  align-items: center;
}

.back-btn {
  background: white;
  border: 1px solid #e5e7eb;
  padding: 8px 16px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  color: var(--font-color-medium);
  transition: all 0.2s;

  &:hover {
    color: var(--font-color-dark);
    background: #f9fafb;
    border-color: #d1d5db;
  }
}

/* CATEGORIES GRID (HOME) */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.category-card {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 24px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  user-select: none;
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease-in-out;
  position: relative;
  overflow: hidden;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.05);
    border-color: #d1d5db;

    .cat-icon-wrapper {
      background: #e8f6f4;
      transform: scale(1.1);
    }

    .cat-arrow {
      opacity: 1;
      transform: translateX(0);
    }
  }
}

.cat-icon-wrapper {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 16px;
  transition: all 0.3s;
}

.cat-icon {
  font-size: 1.5rem;
}

.cat-details {
  h3 {
    margin: 0 0 8px 0;
    font-size: 1.1rem;
    color: var(--font-color-dark);
  }
  p {
    margin: 0;
    font-size: 0.9rem;
    color: var(--font-color-medium);
    line-height: 1.4;
  }
}

.cat-arrow {
  position: absolute;
  top: 24px;
  right: 24px;
  font-size: 1.2rem;
  color: var(--font-color-medium);
  opacity: 0;
  transform: translateX(-10px);
  transition: all 0.2s;
}

/* SETTINGS CONTENT (DETAIL) */
.settings-content {
  flex: 1;
  min-width: 0;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.save-btn {
  background: var(--purple);
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;

  &:disabled {
    opacity: 0.7;
    cursor: wait;
  }
}

.settings-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
  align-items: start;
}

.single-col {
  grid-template-columns: 1fr;
}

.card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden;
}

.card-header {
  padding: 16px 24px;
  background: #f9fafb;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  gap: 12px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--font-color-dark);
  }
}

.icon-box {
  width: 32px;
  height: 32px;
  background: white;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
}

.card-content {
  padding: 24px;
}

.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

/* Forms Grid - Responsive */
.forms-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
}

.form-config-col {
  border: 1px solid #f3f4f6;
  border-radius: 8px;
  padding: 16px;
  background: #fafafa;
}

.config-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;

  h4 {
    margin: 0;
    font-size: 1rem;
    color: var(--font-color-dark);
    font-weight: 600;
  }
}

.toggle-wrapper.small {
  width: 100px;
}
.toggle-wrapper.small :deep(.toggle-btn) {
  font-size: 0.8rem;
  padding: 2px;
}
.toggle-wrapper.small :deep(.toggle-container) {
  height: 36px;
}

.sub-label {
  display: block;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--font-color-medium);
  margin-bottom: 8px;
}

.capsule-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 12px;
  min-height: 30px;
}

.remove-x {
  margin-left: 6px;
  cursor: pointer;
  opacity: 0.6;
  &:hover {
    opacity: 1;
  }
}

.add-email-row {
  display: flex;
  gap: 8px;
  align-items: flex-start;

  :deep(.field-group) {
    margin-bottom: 0;
    width: 100%;
  }
}

/* General Settings Rows */
.setting-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f3f4f6;

  &:last-of-type {
    border-bottom: none;
    margin-bottom: 0px;
    padding-bottom: 0px;
  }
}

.setting-info {
  flex: 1;
  padding-right: 16px;

  label {
    display: block;
    font-weight: 600;
    color: var(--font-color-dark);
    margin-bottom: 4px;
  }
  .setting-desc {
    margin: 0;
    font-size: 0.85rem;
    color: var(--font-color-medium);
  }
}

.toggle-wrapper {
  width: 140px;
}

/* Utilities */
.mb-4 {
  margin-bottom: 16px;
}
.my-4 {
  margin: 16px 0;
  border: 0;
  border-top: 1px solid #e5e7eb;
}
.mt-4 {
  margin-top: 16px;
}

.u-label {
  display: block;
  font-weight: 600;
  font-size: 0.9rem;
  margin-bottom: 8px;
  color: var(--font-color-dark);
}

.system-info p {
  margin: 4px 0;
  font-size: 0.9rem;
  color: var(--font-color-medium);
}

.toast-notification {
  position: fixed;
  bottom: 24px;
  right: 24px;
  background: #10b981;
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  animation: slideUp 0.3s ease-out;
  z-index: 100;
}

@keyframes slideUp {
  from {
    transform: translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Placeholder */
.placeholder-content {
  text-align: center;
  padding: 60px 20px;
  background: white;
  border-radius: 12px;
  border: 1px dashed #e5e7eb;
  color: var(--font-color-medium);

  .placeholder-icon {
    font-size: 3rem;
    margin-bottom: 16px;
    display: block;
  }

  h3 {
    margin-bottom: 8px;
    color: var(--font-color-dark);
  }
}

@media (max-width: 1200px) {
  .settings-grid {
    grid-template-columns: 1fr;
  }
}
</style>
