import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<{ email: string; name: string; role: 'admin' | 'volunteer' } | null>(null)

  // Mock function to check if user is logged in (persisted state could be added later)
  const isAuthenticated = computed(() => !!user.value)

  const login = async (email: string, password: string): Promise<boolean> => {
    // Simulate API delay
    await new Promise((resolve) => setTimeout(resolve, 800))

    // Mock verification
    if (email === 'admin@idohr.org' && password === 'admin123') {
      user.value = {
        email,
        name: 'Test-User',
        role: 'admin',
      }
      return true
    }

    if (password === 'volunteer123') {
      user.value = {
        email,
        name: 'Volunteer User',
        role: 'volunteer',
      }
      return true
    }

    return false
  }

  const devLogin = () => {
    user.value = {
      email: 'admin@idohr.org',
      name: 'John',
      role: 'admin',
    }
  }

  const logout = () => {
    user.value = null
  }

  return {
    user,
    isAuthenticated,
    login,
    devLogin,
    logout,
  }
})
