<script setup lang="ts">
import { ref, computed } from 'vue'
import { useDemoMode } from '../../composables/useDemoMode'
import { useAuthStore } from '../../stores/auth'

// Components
import SettingsGeneral from '../../components/admin/settings/SettingsGeneral.vue'
import SettingsWebsite from '../../components/admin/settings/SettingsWebsite.vue'
import SettingsVolunteers from '../../components/admin/settings/SettingsVolunteers.vue'
import SettingsOverview from '../../components/admin/settings/SettingsOverview.vue'
import SettingsPets from '../../components/admin/settings/SettingsPets.vue'

import { useSettingsStore } from '../../stores/settings'
import { storeToRefs } from 'pinia'

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

// ... existing imports ...

const { isDemoMode, toggleDemoMode } = useDemoMode()
const authStore = useAuthStore()
const settingsStore = useSettingsStore()
const { settings } = storeToRefs(settingsStore)

// Account Settings Form
const accountForm = ref({
  name: authStore.user?.Name || '',
  email: authStore.user?.Email || '',
  password: '',
  confirmPassword: '',
})

const savingAccount = ref(false)

async function updateAccount(showToastOnSuccess = true) {
  if (
    accountForm.value.password &&
    accountForm.value.password !== accountForm.value.confirmPassword
  ) {
    alert('Passwords do not match')
    throw new Error('Passwords do not match')
  }

  savingAccount.value = true
  try {
    const res = await fetch('/api/users', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: accountForm.value.name,
        email: accountForm.value.email,
        password: accountForm.value.password || undefined,
      }),
    })

    if (!res.ok) throw new Error('Failed to update')

    const data = await res.json()
    // Update local store user
    if (data.user) {
      authStore.user = data.user
    }

    accountForm.value.password = ''
    accountForm.value.confirmPassword = ''

    if (showToastOnSuccess) {
      showToast.value = true
      setTimeout(() => {
        showToast.value = false
      }, 3000)
    }
  } catch (e) {
    console.error(e)
    alert('Failed to update profile')
    throw e // Re-throw to inform parent
  } finally {
    savingAccount.value = false
  }
}

async function saveSettings() {
  saving.value = true

  try {
    // If on General tab, also try to save account settings if dirty
    if (activeCategory.value === 'general') {
      if (
        accountForm.value.password ||
        accountForm.value.name !== authStore.user?.Name ||
        accountForm.value.email !== authStore.user?.Email
      ) {
        await updateAccount(false) // false = don't show separate toast
      }
    }

    // Simulate other settings save (or real API calls if they existed)
    await new Promise((resolve) => setTimeout(resolve, 800))

    saving.value = false
    activeCategory.value = null
    showToast.value = true
    setTimeout(() => {
      showToast.value = false
    }, 3000)
  } catch (e) {
    saving.value = false
    console.error(e)
  }
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
      <SettingsGeneral
        v-if="activeCategory === 'general'"
        v-model:settings="settings"
        v-model:accountForm="accountForm"
      />

      <SettingsWebsite
        v-if="activeCategory === 'website'"
        v-model:settings="settings"
        :isDemoMode="isDemoMode"
        @toggleDemoMode="toggleDemoMode"
      />

      <SettingsVolunteers v-if="activeCategory === 'volunteers'" v-model:settings="settings" />

      <SettingsOverview v-if="activeCategory === 'overview'" v-model:settings="settings" />

      <SettingsPets v-if="activeCategory === 'pets'" v-model:settings="settings" />

      <!-- PLACEHOLDER TABS -->
      <div
        v-if="
          activeCategory &&
          !['general', 'website', 'volunteers', 'overview', 'pets'].includes(activeCategory)
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
    color: var(--text-primary);
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
  font-weight: 500;
  color: var(--text-secondary);
  transition: all 0.2s;
}

.back-btn:hover {
  background: #f9fafb;
  border-color: #d1d5db;
}

.save-btn {
  background: var(--color-primary);
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

.save-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.toast-notification {
  position: fixed;
  bottom: 24px;
  right: 24px;
  background: #10b981;
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  animation: slideIn 0.3s ease-out;
  z-index: 100;
  font-weight: 500;
}

@keyframes slideIn {
  from {
    transform: translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* HOME GRID */
.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  padding-bottom: 40px;
}

.category-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 24px;
  text-align: left;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: relative;
  height: 100%;
}

.category-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.05);
  border-color: var(--primary-color-light, #bfdbfe);
}

.cat-icon-wrapper {
  width: 48px;
  height: 48px;
  background: #f3f4f6;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
}

.cat-details h3 {
  margin: 0 0 8px 0;
  font-size: 1.1rem;
  color: var(--text-primary);
}

.cat-details p {
  margin: 0;
  font-size: 0.9rem;
  color: var(--text-secondary);
  line-height: 1.5;
}

.cat-arrow {
  margin-top: auto;
  align-self: flex-end;
  color: #d1d5db;
  font-weight: bold;
  transition: color 0.2s;
}

.category-card:hover .cat-arrow {
  color: var(--color-primary);
}

/* DETAIL CONTENT */
.settings-content {
  flex: 1;
  overflow-y: auto;
  padding-bottom: 40px;
}

/* PLACEHOLDER */
.placeholder-content {
  text-align: center;
  padding: 60px 20px;
  background: white;
  border-radius: 12px;
  border: 2px dashed #e5e7eb;
  color: var(--text-secondary);
}

.placeholder-icon {
  font-size: 3rem;
  margin-bottom: 20px;
  display: inline-block;
  opacity: 0.5;
}
</style>
