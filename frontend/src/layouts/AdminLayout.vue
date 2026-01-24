<script setup lang="ts">
import { computed,ref } from 'vue'
import { useRouter } from 'vue-router'

import { SidebarNav } from '../components/common/ui'
import { useAuthStore } from '../stores/auth'

const router = useRouter()

const authStore = useAuthStore()

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}

const navItems = [
  { label: 'Overview', to: '/admin', icon: 'dashboard', exact: true },
  { label: 'Calendar', to: '/admin/calendar', icon: 'calendar' },
  { label: 'Pet Records', to: '/admin/pets', icon: 'pawPrint', viewBox: '0 0 128 128' },
  { label: 'Applications', to: '/admin/applications', icon: 'fileText' },
  { label: 'Medical', to: '/admin/pet-health', icon: 'activity' },
  { label: 'Volunteers', to: '/admin/volunteers', icon: 'users' },
  { label: 'Transport', to: '/admin/transport', icon: 'truck' },
  { label: 'Timesheets', to: '/admin/time-logs', icon: 'clock' },
  { label: 'Messages', to: '/admin/messages', icon: 'mail' },
  { label: 'Donations', to: '/admin/donations', icon: 'dollar' },
  { label: 'Inventory', to: '/admin/inventory', icon: 'box' },
  { label: 'Marketing', to: '/admin/marketing', icon: 'megaphone' },
  { label: 'Intake Kiosk', to: '/admin/kiosk', icon: 'monitor' },
  { label: 'Smart Kennel Cards', to: '/admin/kennel-displays', icon: 'tag' },
  { label: 'Event Signage', to: '/admin/event-displays', icon: 'sign' },
  { label: 'Settings', to: '/admin/settings', icon: 'cog' },
]

const userName = computed(() => authStore.user?.Name || 'Admin User')
const userInitials = computed(() => {
  const name = authStore.user?.Name || 'AD'
  return name
    .split(' ')
    .map((n) => n[0])
    .join('')
    .substring(0, 2)
    .toUpperCase()
})
</script>

<template>
  <div class="admin-layout">

    <aside class="sidebar" :class="{ 'mobile-open': isMobileMenuOpen }">
      <div class="sidebar-header">
        <img src="/images/idohr-logo.jpg" alt="IDOHR Logo" class="logo" />
        <h2>Admin</h2>
      </div>

      <SidebarNav :items="navItems" variant="dashboard" @click="isMobileMenuOpen = false" />

      <div class="sidebar-footer">
        <div class="user-info">
          <div class="avatar">{{ userInitials }}</div>
          <div class="details">
            <span class="name">{{ userName }}</span>
            <span class="role">Administrator</span>
          </div>
        </div>
        <button @click="handleLogout" class="logout-btn" title="Logout">🚪</button>
      </div>
    </aside>

    <div v-if="isMobileMenuOpen" class="sidebar-overlay" @click="isMobileMenuOpen = false"></div>

    <div class="main-content">
      <header class="top-bar">
        <button class="menu-toggle" @click="toggleMobileMenu">☰</button>
      </header>

      <div class="page-content">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-layout {
  display: flex;
  height: 100vh;
  width: 100vw;
  background-color: #f3f4f6;
  overflow: hidden;
}

.sidebar {
  width: 280px;
  background-color: var(--text-inverse);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  transition: transform 0.3s ease;
  z-index: 50;
  box-shadow: 4px 0 24px rgb(0 0 0 / 2%);
}

.sidebar-header {
  padding: 32px 24px;
  display: flex;
  align-items: center;
  gap: 16px;

  .logo {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    object-fit: cover;
    box-shadow: 0 4px 12px rgb(0 0 0 / 5%);
  }

  h2 {
    font-size: 1.5rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
    letter-spacing: -0.02em;
  }
}

.sidebar-footer {
  padding: 24px;
  border-top: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: hsl(from var(--color-neutral) h s 99%);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;

  .avatar {
    width: 42px;
    height: 42px;
    background: linear-gradient(135deg, var(--color-tertiary), #8b5cf6);
    color: #fff;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 700;
    font-size: 0.9rem;
    box-shadow: 0 2px 8px hsl(from var(--color-tertiary) h s l / 20%);
  }

  .details {
    display: flex;
    flex-direction: column;
    gap: 2px;

    .name {
      font-size: 0.95rem;
      font-weight: 700;
      color: var(--text-primary);
    }

    .role {
      font-size: 0.75rem;
      color: hsl(from var(--color-neutral) h s 50%);
      font-weight: 500;
      text-transform: uppercase;
      letter-spacing: 0.05em;
    }
  }
}

.logout-btn {
  background: var(--text-inverse);
  border: 1px solid hsl(from var(--color-danger) h s 90%);
  cursor: pointer;
  font-size: 1.1rem;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  transition: all 0.2s;
  color: var(--color-danger);

  &:hover {
    background-color: hsl(from var(--color-danger) h s 95%);
    transform: translateY(-2px);
  }
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: hsl(from var(--color-neutral) h s 98%);
}

.top-bar {
  display: none;
  height: 80px;
  background-color: transparent;
  align-items: center;
  padding: 0 40px;
  gap: 16px;
  flex-shrink: 0;

  h3 {
    font-size: 1.75rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
    letter-spacing: -0.02em;
  }
}

.menu-toggle {
  display: none;
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: var(--text-primary);
  padding: 4px;
}

.page-content {
  flex: 1;
  overflow-y: auto;
  padding: 32px 40px 40px;
}

.sidebar-overlay {
  position: fixed;
  inset: 0;
  background-color: rgb(0 0 0 / 50%);
  backdrop-filter: blur(4px);
  z-index: 40;
  display: none;
}

@media (width <= 768px) {
  .sidebar {
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    transform: translateX(-100%);
    box-shadow: 4px 0 24px rgb(0 0 0 / 10%);
  }

  .sidebar.mobile-open {
    transform: translateX(0);
  }

  .sidebar-overlay {
    display: block;
  }

  .menu-toggle {
    display: block;
  }

  .top-bar {
    display: flex;
    padding: 0 24px;
    height: 64px;
    background-color: var(--white);
    border-bottom: 1px solid rgb(0 0 0 / 5%);

    h3 {
      font-size: 1.25rem;
    }
  }

  .page-content {
    padding: 24px;
  }
}
</style>
