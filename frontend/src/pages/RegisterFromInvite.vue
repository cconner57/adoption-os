<template>
  <div class="register-page">
    <div class="register-card">
      <div class="logo-section">
        <h1>Complete Registration</h1>
        <p class="subtitle">Welcome to Adoption OS</p>
      </div>

      <div v-if="loading" class="loading-state">
        <p>Validating invite...</p>
      </div>

      <div v-else-if="error" class="error-state">
        <div class="error-icon">⚠️</div>
        <h3>Invalid Invite</h3>
        <p>{{ error }}</p>
        <router-link to="/login" class="back-link">Return to Login</router-link>
      </div>

      <div v-else-if="success" class="success-state">
        <div class="success-icon">✅</div>
        <h3>Account Created!</h3>
        <p>Your account has been successfully created.</p>
        <router-link to="/login" class="btn-primary" style="display: block; text-align: center; margin-top: 1rem; text-decoration: none;">Proceed to Login</router-link>
      </div>

      <form v-else @submit.prevent="handleSubmit" class="register-form">
        <div class="form-group">
          <label for="email-input">Email</label>
          <input id="email-input" type="email" :value="email" disabled class="input-disabled" />
          <span class="help-text">Email matches your invitation</span>
        </div>

        <div class="form-group">
          <label for="name">Full Name</label>
          <input
            id="name"
            v-model="name"
            type="text"
            required
            placeholder="Jane Doe"
            autofocus
          />
        </div>

        <div class="form-group">
          <label for="password">Set Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            required
            placeholder="••••••••"
            minlength="8"
          />
        </div>

        <div v-if="formError" class="error-message">
            {{ formError }}
        </div>

        <div class="form-actions">
          <button type="submit" :disabled="submitting" class="btn-primary">
            {{ submitting ? 'Creating Account...' : 'Create Account' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

import { useUIStore } from '../stores/ui'

const route = useRoute()
const uiStore = useUIStore()

const token = route.query.token as string

const loading = ref(true)
const submitting = ref(false)
const error = ref<string | null>(null)
const formError = ref<string | null>(null)
const success = ref(false)

const email = ref('')
const name = ref('')
const password = ref('')
const role = ref('')

onMounted(async () => {
  if (!token) {
    error.value = 'Missing invite token.'
    loading.value = false
    return
  }

  try {
    const res = await fetch(`/api/validate-invite/${token}`)

    if (!res.ok) {
      if (res.status === 404) throw new Error('Invite not found.')
      if (res.status === 400) throw new Error('Invite has expired or is invalid.')
      throw new Error('Failed to validate invite.')
    }

    const data = await res.json()
    email.value = data.invite.email
    role.value = data.invite.role
  } catch (e) {
    if (e instanceof Error) {
        error.value = e.message
    } else {
        error.value = 'An error occurred during validation.'
    }
  } finally {
    loading.value = false
  }
})

const handleSubmit = async () => {
  if (!name.value || !password.value) return

  submitting.value = true
  formError.value = null
  uiStore.startLoading()

  try {
    const response = await fetch('/api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: name.value,
        email: email.value,
        password: password.value,
        token: token
      }),
    })

    if (!response.ok) {
      const errData = await response.json()
      throw new Error(errData.error || 'Registration failed')
    }

    success.value = true

  } catch (e) {
    if (e instanceof Error) {
        formError.value = e.message
    } else {
        formError.value = 'An error occurred during registration.'
    }
  } finally {
    submitting.value = false
    uiStore.stopLoading()
  }
}
</script>

<style scoped>
.register-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-background-soft);
  padding: 1rem;
}

.register-card {
  background: #fff;
  padding: 2.5rem;
  border-radius: 1rem;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%), 0 2px 4px -2px rgb(0 0 0 / 10%);
  width: 100%;
  max-width: 420px;
}

.logo-section {
  text-align: center;
  margin-bottom: 2rem;
}

h1 {
  color: var(--color-primary);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.subtitle {
  color: var(--text-secondary);
  font-size: 0.95rem;
}

.form-group {
  margin-bottom: 1.25rem;
}

label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 0.5rem;
}

input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-border);
  border-radius: 0.5rem;
  font-size: 1rem;
  transition: all 0.2s;
}

input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px rgb(0 165 173 / 10%);
}

.input-disabled {
  background-color: #f3f4f6;
  color: var(--text-secondary);
  cursor: not-allowed;
}

.help-text {
  display: block;
  font-size: 0.75rem;
  color: var(--text-secondary);
  margin-top: 0.25rem;
}

.btn-primary {
  width: 100%;
  padding: 0.75rem;
  background-color: var(--color-primary);
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-primary:hover:not(:disabled) {
  filter: brightness(0.95);
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.error-state {
  text-align: center;
  padding: 2rem 0;
}

.error-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.back-link {
  display: inline-block;
  margin-top: 1rem;
  color: var(--color-primary);
  text-decoration: none;
  font-weight: 500;
}

.back-link:hover {
  text-decoration: underline;
}
</style>
