import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<{
    ID: number
    Name: string
    Email: string
    Role: string
  } | null>(null)

  const isAuthenticated = computed(() => !!user.value)

  const login = async (email: string, password: string): Promise<boolean> => {
    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (response.ok) {
        // Login successful, now fetch user details
        await checkAuth()
        return true
      }
      return false
    } catch (error) {
      console.error('Login error:', error)
      return false
    }
  }

  const checkAuth = async () => {
    try {
      const response = await fetch('/api/users/me')
      if (response.ok) {
        const data = await response.json()
        if (data && data.data) {
          user.value = data.data
        }
      } else {
        user.value = null
      }
    } catch (error) {
      console.error('Check auth error:', error)
      user.value = null
    }
  }

  const logout = async () => {
    try {
      await fetch('/api/users/logout', {
        method: 'POST',
        credentials: 'include',
      })
    } catch (error) {
      console.error('Logout error:', error)
    } finally {
      user.value = null
      window.location.reload() // Force reload to clear any memory/state
    }
  }

  const initialized = ref(false)
  const initialize = async () => {
    if (initialized.value) return
    await checkAuth()
    initialized.value = true
  }

  return {
    user,
    isAuthenticated,
    login,
    checkAuth,
    logout,
    initialize,
    initialized,
  }
})
