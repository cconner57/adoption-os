<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const menuItems = [
  {
    title: 'Clock In / Out',
    icon: '🕒',
    route: '/admin/time-logs',
    desc: 'Start or end your volunteer shift.',
  }, 
  {
    title: 'Daily Care Sheet',
    icon: '📋',
    route: '/kiosk/daily-care',
    desc: 'Fill out feeding, cleaning, & health logs.',
  },
  {
    title: 'Pet Profiles',
    icon: '🐾',
    route: '/kiosk/pets',
    desc: 'Browse info for all current residents.',
  },
  {
    title: 'Vet Schedule',
    icon: '🩺',
    route: '/kiosk/vet',
    desc: 'Upcoming medical appointments.',
  },
  {
    title: 'Report Issue',
    icon: '⚠️',
    route: '/admin/inventory',
    desc: 'Log low supplies or maintenance issues.',
  }, 
]
</script>

<template>
  <div class="kiosk-home">
    <div class="welcome-banner">
      <h1>Welcome, Volunteer! 👋</h1>
      <p>Thank you for helping us care for these amazing animals today.</p>
    </div>

    <div class="menu-grid">
      <button
        v-for="item in menuItems"
        :key="item.title"
        class="menu-card"
        @click="router.push(item.route)"
      >
        <div class="icon-circle">{{ item.icon }}</div>
        <div class="card-text">
          <h3>{{ item.title }}</h3>
          <p>{{ item.desc }}</p>
        </div>
      </button>
    </div>
  </div>
</template>

<style scoped>
.kiosk-home {
  max-width: 1000px;
  margin: 0 auto;
  padding-top: 40px;
}

.welcome-banner {
  text-align: center;
  margin-bottom: 60px;

  h1 {
    font-size: 2.5rem;
    margin-bottom: 12px;
    color: var(--text-primary);
  }
  p {
    font-size: 1.2rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.menu-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
}

.menu-card {
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 20px;
  padding: 32px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);

  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    border-color: var(--color-secondary);
  }

  &:active {
    transform: scale(0.98);
  }
}

.icon-circle {
  width: 80px;
  height: 80px;
  background: #f3f4f6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.5rem;
}

.card-text {
  h3 {
    margin: 0 0 8px 0;
    font-size: 1.4rem;
    color: var(--text-primary);
  }
  p {
    margin: 0;
    color: hsl(from var(--color-neutral) h s 50%);
    font-size: 1rem;
    line-height: 1.4;
  }
}
</style>
