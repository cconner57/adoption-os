import { ref } from 'vue'

const DEMO_STORAGE_KEY = 'adoption_os_demo_mode'
const isDemoMode = ref(sessionStorage.getItem(DEMO_STORAGE_KEY) === 'true')

export function useDemoMode() {
  const toggleDemoMode = (value: boolean) => {
    isDemoMode.value = value
    if (value) {
      sessionStorage.setItem(DEMO_STORAGE_KEY, 'true')
    } else {
      sessionStorage.removeItem(DEMO_STORAGE_KEY)
    }
  }

  return {
    isDemoMode,
    toggleDemoMode,
  }
}
