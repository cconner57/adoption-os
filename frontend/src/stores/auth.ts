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
          // Temporary Role mapping until backend sends role
          // Defaulting Admin for now since we only seeded Admin
          user.value!.Role = 'admin'
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
    // Ideally call backend logout endpoint if we had one to clear cookie
    // For now just clear local state, but cookie will persist until expiry or cleared
    // Note: To do this properly, we need a POST /users/logout endpoint that clears the cookie.
    user.value = null
    window.location.reload() // Force reload to clear any memory/state
  }

  return {
    user,
    isAuthenticated,
    login,
    checkAuth,
    logout,
  }
})
