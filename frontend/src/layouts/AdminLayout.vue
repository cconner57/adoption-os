<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const route = useRoute()
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
  { name: 'Overview', path: '/admin', icon: '📊' },
  { name: 'Calendar', path: '/admin/calendar', icon: '📅' },
  { name: 'Pets', path: '/admin/pets', icon: '🐾' },
  { name: 'Pet Health', path: '/admin/pet-health', icon: '🩺' },
  { name: 'Volunteers', path: '/admin/volunteers', icon: '🤝' },
  { name: 'Messages', path: '/admin/messages', icon: '✉️' },
  { name: 'Settings', path: '/admin/settings', icon: '⚙️' },
]
</script>

<template>
  <div class="admin-layout">
    <!-- Sidebar -->
    <aside class="sidebar" :class="{ 'mobile-open': isMobileMenuOpen }">
      <div class="sidebar-header">
        <img src="/images/idohr-logo.jpg" alt="IDOHR Logo" class="logo" />
        <h2>Admin</h2>
      </div>

      <nav class="sidebar-nav">
        <router-link
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          class="nav-item"
          :class="{ active: route.path === item.path }"
          @click="isMobileMenuOpen = false"
        >
          <span class="icon">{{ item.icon }}</span>
          <span class="label">{{ item.name }}</span>
        </router-link>
      </nav>

      <div class="sidebar-footer">
        <div class="user-info">
          <div class="avatar">AD</div>
          <div class="details">
            <span class="name">Admin User</span>
            <span class="role">Administrator</span>
          </div>
        </div>
        <button @click="handleLogout" class="logout-btn" title="Logout">🚪</button>
      </div>
    </aside>

    <!-- Overlay for mobile -->
    <div v-if="isMobileMenuOpen" class="sidebar-overlay" @click="isMobileMenuOpen = false"></div>

    <!-- Main Content -->
    <div class="main-content">
      <header class="top-bar">
        <button class="menu-toggle" @click="toggleMobileMenu">☰</button>
        <h3>Dashboard</h3>
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

/* Sidebar Styles */
.sidebar {
  width: 280px;
  background-color: var(--white);
  border-right: 1px solid rgba(0, 0, 0, 0.08);
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  transition: transform 0.3s ease;
  z-index: 50;
  box-shadow: 4px 0 24px rgba(0, 0, 0, 0.02);
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
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  }

  h2 {
    font-size: 1.5rem;
    font-weight: 800;
    color: var(--font-color-dark);
    margin: 0;
    letter-spacing: -0.02em;
  }
}

.sidebar-nav {
  flex: 1;
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  overflow-y: auto;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 20px;
  border-radius: 12px;
  text-decoration: none;
  color: var(--font-color-medium);
  transition: all 0.2s cubic-bezier(0.2, 0.8, 0.2, 1);
  font-weight: 600;
  font-size: 1rem;

  &:hover {
    background-color: #f8fafc;
    color: var(--font-color-dark);
    transform: translateX(4px);
  }

  &.active {
    background-color: var(--green);
    color: var(--white);
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.2);

    &:hover {
      transform: none;
    }
  }

  .icon {
    font-size: 1.25rem;
  }
}

.sidebar-footer {
  padding: 24px;
  border-top: 1px solid rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #fcfcfc;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;

  .avatar {
    width: 42px;
    height: 42px;
    background: linear-gradient(135deg, var(--purple), #8b5cf6);
    color: white;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 700;
    font-size: 0.9rem;
    box-shadow: 0 2px 8px rgba(124, 58, 237, 0.2);
  }

  .details {
    display: flex;
    flex-direction: column;
    gap: 2px;

    .name {
      font-size: 0.95rem;
      font-weight: 700;
      color: var(--font-color-dark);
    }
    .role {
      font-size: 0.75rem;
      color: var(--font-color-medium);
      font-weight: 500;
      text-transform: uppercase;
      letter-spacing: 0.05em;
    }
  }
}

.logout-btn {
  background: var(--white);
  border: 1px solid #fee2e2;
  cursor: pointer;
  font-size: 1.1rem;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  transition: all 0.2s;
  color: #ef4444;

  &:hover {
    background-color: #fee2e2;
    transform: translateY(-2px);
  }
}

/* Main Content Styles */
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: #f8fafc; /* Cooler, lighter gray */
}

.top-bar {
  height: 80px;
  background-color: transparent; /* Cleaner look */
  display: flex;
  align-items: center;
  padding: 0 40px;
  gap: 16px;
  flex-shrink: 0;

  h3 {
    font-size: 1.75rem;
    font-weight: 800;
    color: var(--font-color-dark);
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
  color: var(--font-color-dark);
  padding: 4px;
}

.page-content {
  flex: 1;
  overflow-y: auto;
  padding: 0 40px 40px 40px;
}

/* Mobile Responsive */
.sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(4px);
  z-index: 40;
  display: none;
}

@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    transform: translateX(-100%);
    box-shadow: 4px 0 24px rgba(0, 0, 0, 0.1);
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
    padding: 0 24px;
    height: 64px;
    background-color: var(--white);
    border-bottom: 1px solid rgba(0, 0, 0, 0.05);

    h3 {
      font-size: 1.25rem;
    }
  }

  .page-content {
    padding: 24px;
  }
}
</style>
