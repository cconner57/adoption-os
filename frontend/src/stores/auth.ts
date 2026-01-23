import { defineStore } from 'pinia'
import { computed,ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<{
    ID: number
    Name: string
    Email: string
    Role: string
  } | null>(null)

  const isAuthenticated = computed(() => !!user.value)

  const checkAuth = async () => {
    try {
      const response = await fetch('/api/users/me', {
        credentials: 'include',
      })
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

  const login = async (email: string, password: string): Promise<boolean> => {
    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({ email, password }),
      })

      const data = await response.json()

      if (response.ok) {
        
        if (data.token) {
          localStorage.setItem('token', data.token)
        }

        if (data.user) {
          user.value = data.user
        } else {
          await checkAuth()
        }
        return true
      }
      return false
    } catch (error) {
      console.error('Login error:', error)
      return false
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
      localStorage.removeItem('token')
      user.value = null
      window.location.reload() 
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
