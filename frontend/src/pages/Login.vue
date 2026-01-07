<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
// import { useAuthStore } from '../stores/auth'
import InputField from '../components/common/ui/InputField.vue'
import Button from '../components/common/ui/Button.vue'

const email = ref('')
const password = ref('')
const error = ref('')
const isLoading = ref(false)

const router = useRouter()
// const authStore = useAuthStore()

const handleLogin = async () => {
  // validation disabled for development
  isLoading.value = true

  // Simulate API delay for effect
  await new Promise((resolve) => setTimeout(resolve, 500))

  // Bypass auth checks for now
  router.push('/admin')
  isLoading.value = false
}
</script>

<template>
  <div class="login-page">
    <div class="login-card">
      <div class="logo-area">
        <img src="/images/idohr-logo.jpg" alt="IDOHR Logo" class="logo" />
        <h1>Welcome Back</h1>
        <p>Sign in to access the volunteer dashboard</p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <InputField
          label="Email Address"
          placeholder="name@example.com"
          type="email"
          name="email"
          v-model="email"
          :hasError="!!error"
        />

        <InputField
          label="Password"
          placeholder="Enter your password"
          type="password"
          name="password"
          v-model="password"
          :hasError="!!error"
        />

        <div v-if="error" class="error-message">
          {{ error }}
        </div>

        <Button
          type="submit"
          title="Sign In"
          color="green"
          size="medium"
          :loading="isLoading"
          fullWidth
          class="submit-btn"
        />
      </form>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--green);
  padding: 20px;
}

.login-card {
  background: var(--white);
  width: 100%;
  max-width: 440px;
  padding: 48px;
  border-radius: 24px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  gap: 32px;
  animation: slideUp 0.6s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.logo-area {
  text-align: center;
}

.logo {
  height: 80px;
  width: auto;
  margin: 0 auto 24px auto;
  display: block;
}

.logo-area h1 {
  font-size: 1.75rem;
  color: var(--font-color-dark);
  margin-bottom: 8px;
  font-weight: 700;
}

.logo-area p {
  color: var(--font-color-medium);
  font-size: 1rem;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.error-message {
  color: #ef4444;
  font-size: 0.9rem;
  text-align: center;
  background-color: #fef2f2;
  padding: 10px;
  border-radius: 8px;
}

.submit-btn {
  margin-top: 8px;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 480px) {
  .login-card {
    padding: 32px 24px;
  }
}
</style>
