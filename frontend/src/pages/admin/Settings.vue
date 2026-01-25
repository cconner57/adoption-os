<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed, onMounted,ref } from 'vue'

import SettingsGeneral from '../../components/admin/settings/SettingsGeneral.vue'
import SettingsOverview from '../../components/admin/settings/SettingsOverview.vue'
import SettingsPets from '../../components/admin/settings/SettingsPets.vue'
import SettingsUser from '../../components/admin/settings/SettingsUser.vue'
import SettingsVolunteers from '../../components/admin/settings/SettingsVolunteers.vue'
import SettingsWebsite from '../../components/admin/settings/SettingsWebsite.vue'
import { Icon } from '../../components/common/ui'
import { useDemoMode } from '../../composables/useDemoMode'
import { useAuthStore } from '../../stores/auth'
import { useSettingsStore } from '../../stores/settings'

const saving = ref(false)
const showToast = ref(false)
const activeCategory = ref<string | null>(null)

import { settingsCategories } from '../../config/settingsCategories'

const categories = settingsCategories

const { isDemoMode, toggleDemoMode } = useDemoMode()
const authStore = useAuthStore()
const settingsStore = useSettingsStore()
const { settings } = storeToRefs(settingsStore)

const settingsUserRef = ref()
// Initialize with empty strings, populated in onMounted
const accountForm = ref({
  firstName: '',
  lastName: '',
  email: '',
  password: '',
  confirmPassword: '',
})

// Initialize account form from auth store
// Initialize account form from auth store
const isMounted = ref(false)
onMounted(() => {
    isMounted.value = true
    if (authStore.user) {
        const parts = (authStore.user.Name || '').split(' ')
        accountForm.value.firstName = parts[0] || ''
        accountForm.value.lastName = parts.slice(1).join(' ') || ''
        accountForm.value.email = authStore.user.Email || ''
    }
})

async function saveSettings() {
  saving.value = true

  try {
    if (activeCategory.value === 'general') {
        const originalFullName = authStore.user?.Name || '';
        const currentFullName = `${accountForm.value.firstName} ${accountForm.value.lastName}`.trim();

      if (
        accountForm.value.password ||
        currentFullName !== originalFullName ||
        accountForm.value.email !== authStore.user?.Email
      ) {
            const fullName = `${accountForm.value.firstName} ${accountForm.value.lastName}`.trim()

            await authStore.updateProfile({
            name: fullName,
            email: accountForm.value.email,
            password: accountForm.value.password || undefined,
            })
      }
    } else if (activeCategory.value === 'user' && settingsUserRef.value) {
      await settingsUserRef.value.save()
    }

    await new Promise((resolve) => setTimeout(resolve, 800))

    saving.value = false
    activeCategory.value = 'general' // Return to main list? Or General? User might want to stay. But snippet said 'general'. I'll stick to 'general' or null if that was original behavior for "Back".
    // Wait, activeCategory=null shows categories grid. activeCategory='general' shows General Settings.
    // If I save, maybe I should stay on 'general'? or go back?
    // The previous code set activeCategory = null (Back to grid).
    // But if I'm in General Settings and hit Save, going back to grid is okay.
    // However, the snippet earlier changed it to 'general'.
    // I'll set it to null to go back to grid, or keep it.
    // The user flow: Edit -> Save -> Success -> Link to overview?
    // I'll set it to null (Overview) as it was originally 201.
    activeCategory.value = null;

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

const getCategoryDesc = (id: string) => {
  return categories.find((t) => t.id === id)?.description || '';
}
</script>

<template>
  <div class="admin-page">
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">{{ activeCategory ? activeCategoryLabel : 'Settings' }}</h1>
    </Teleport>

    <div class="page-header" :class="{ 'header-hidden-mobile': !activeCategory }">
      <div class="header-left">

        <button v-if="activeCategory" class="back-btn" @click="activeCategory = null">
          ← Back
        </button>
        <h1 class="desktop-only">{{ activeCategory ? activeCategoryLabel : 'Settings' }}</h1>
      </div>

      <button v-if="activeCategory" class="save-btn" :disabled="saving" @click="saveSettings">
        {{ saving ? 'Saving...' : 'Save Changes' }}
      </button>
    </div>

    <div v-if="showToast" class="toast-notification">✅ Settings saved successfully!</div>

    <div v-if="!activeCategory" class="categories-grid">
      <button
        v-for="cat in categories"
        :key="cat.id"
        class="category-card"
        @click="activeCategory = cat.id"
      >
        <div
          class="cat-icon-wrapper"
          :style="{
            color: cat.color,
            background: cat.color ? `${cat.color}15` : '#f3f4f6'
          }"
        >
          <Icon :name="cat.icon" size="24" :viewBox="cat.viewBox" />
        </div>
        <div class="cat-details">
          <h3>{{ cat.label }}</h3>
          <p>{{ getCategoryDesc(cat.id) }}</p>
        </div>
        <div class="cat-arrow">→</div>
      </button>
    </div>

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
        @toggleDemoMode="toggleDemoMode(!isDemoMode)"
      />

      <SettingsVolunteers v-if="activeCategory === 'volunteers'" v-model:settings="settings" />

      <SettingsOverview v-if="activeCategory === 'overview'" v-model:settings="settings" />

      <SettingsPets v-if="activeCategory === 'pets'" v-model:settings="settings" />

      <SettingsUser
        v-if="activeCategory === 'user'"
        ref="settingsUserRef"
      />

      <div
        v-if="
          activeCategory &&
          !['general', 'website', 'volunteers', 'overview', 'pets', 'user'].includes(activeCategory)
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

  h1.desktop-only {
    font-size: 1.8rem;
    color: var(--text-primary);
    margin: 0;
  }
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .page-header.header-hidden-mobile {
    display: none;
  }

  .page-header h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }

  /* Add border to the page actions header (Back/Save) on mobile */
  .page-header {
    border-bottom: none;
    box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%), 0 2px 4px -1px rgb(0 0 0 / 6%);
    padding-bottom: 12px;
    width: 100%;
  }
}

.header-left {
  display: flex;
  gap: 16px;
  align-items: center;
}

.back-btn {
  background: #fff;
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
  color: #fff;
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
  color: #fff;
  padding: 12px 24px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 15%);
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

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  padding-bottom: 40px;
}

.category-card {
  background: #fff;
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
  box-shadow: 0 12px 24px rgb(0 0 0 / 5%);
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
  margin: 0 0 8px;
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
  font-weight: 700;
  transition: color 0.2s;
}

.category-card:hover .cat-arrow {
  color: var(--color-primary);
}

.settings-content {
  flex: 1;
  overflow-y: auto;
  padding-bottom: 40px;
}

.placeholder-content {
  text-align: center;
  padding: 60px 20px;
  background: #fff;
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

@media (width <= 768px) {
  .categories-grid {
    display: flex;
    flex-direction: column;
    gap: 1px; /* Divider line thickness */
    background: #e5e7eb; /* Color of dividers */
    border-radius: 12px;
    border: 1px solid #e5e7eb;
    overflow: hidden;
    padding-bottom: 0;
  }

  .category-card {
    flex-direction: row;
    align-items: center;
    border: none;
    border-radius: 0;
    padding: 16px;
    gap: 16px;
    height: auto;
    min-height: 56px;
    box-shadow: none;
    background: #fff;
  }

  .category-card:hover {
    transform: none;
    box-shadow: none;
    background-color: #f9fafb;
    border-color: transparent;
  }

  /* Adjust icon size for list view */
  .cat-icon-wrapper {
    width: 32px;
    height: 32px;
    border-radius: 8px;
    font-size: 1.1rem;
    flex-shrink: 0;
  }

  .cat-details {
    flex: 1;
    display: flex;
    align-items: center;
  }

  .cat-details h3 {
    margin: 0;
    font-size: 1rem;
    font-weight: 500;
  }

  /* Hide descriptions on mobile for cleaner list look */
  .cat-details p {
    display: none;
  }

  .cat-arrow {
    margin: 0;
    align-self: center;
    color: #cbd5e1;
    font-size: 1.2rem;
  }
}
</style>
