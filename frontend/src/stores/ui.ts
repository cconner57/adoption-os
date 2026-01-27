import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useUIStore = defineStore('ui', () => {
  const isLoading = ref(false)

  // Admin UI State
  const adminState = ref({
    isSidebarOpen: false, // Default closed on mobile
    version: 1, // Schema version
  })

  // Persistence Logic
  const STORAGE_KEY = 'admin_ui_state'
  const CURRENT_VERSION = 1

  const saved = localStorage.getItem(STORAGE_KEY)
  if (saved) {
    try {
      const parsed = JSON.parse(saved)
      if (parsed.version === CURRENT_VERSION) {
        adminState.value = parsed
      } else {
        // Version mismatch: Clear specific key only
        console.warn(`Admin UI state version mismatch (found ${parsed.version}, expected ${CURRENT_VERSION}). Resetting.`)
        localStorage.removeItem(STORAGE_KEY)
      }
    } catch (e) {
      console.error('Failed to parse admin UI state', e)
      localStorage.removeItem(STORAGE_KEY)
    }
  }

  // Watch for changes (auto-save)
  watch(
    adminState,
    (state) => {
      // Ensure current version is always saved
      const payload = { ...state, version: CURRENT_VERSION }
      localStorage.setItem(STORAGE_KEY, JSON.stringify(payload))
    },
    { deep: true },
  )

  function startLoading() {
    isLoading.value = true
  }

  function stopLoading() {
    isLoading.value = false
  }

  return {
    isLoading,
    adminState,
    startLoading,
    stopLoading,
  }
})
