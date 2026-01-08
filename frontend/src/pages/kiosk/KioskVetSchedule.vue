<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// Mock Kiosk Vet Data (Simpler than full calendar)
const appointments = ref([
  { id: 1, time: '09:00 AM', pet: 'Rex', type: 'Check-up', vet: 'Dr. Smith', urgent: false },
  { id: 2, time: '10:30 AM', pet: 'Luna', type: 'Vaccination', vet: 'Dr. Smith', urgent: false },
  {
    id: 3,
    time: '01:00 PM',
    pet: 'Mittens',
    type: 'Surgery Follow-up',
    vet: 'Dr. Jones',
    urgent: true,
  },
  {
    id: 4,
    time: '03:15 PM',
    pet: 'Goldie',
    type: 'Dental Cleaning',
    vet: 'Dr. Smith',
    urgent: false,
  },
])

const goHome = () => router.push('/kiosk')
</script>

<template>
  <div class="kiosk-page">
    <div class="page-header">
      <button class="back-link" @click="goHome">← Back</button>
      <h1>Vet Schedule (Today)</h1>
      <div class="date-display">
        {{
          new Date().toLocaleDateString(undefined, {
            weekday: 'long',
            month: 'long',
            day: 'numeric',
          })
        }}
      </div>
    </div>

    <div class="schedule-list">
      <div
        v-for="apt in appointments"
        :key="apt.id"
        class="apt-card"
        :class="{ urgent: apt.urgent }"
      >
        <div class="time-col">
          <span class="time">{{ apt.time }}</span>
        </div>

        <div class="info-col">
          <h2>{{ apt.pet }}</h2>
          <div class="apt-details">{{ apt.type }} • {{ apt.vet }}</div>
        </div>

        <div class="status-col">
          <span v-if="apt.urgent" class="urgent-badge">Urgent</span>
        </div>
      </div>

      <div v-if="appointments.length === 0" class="empty-state">
        No veterinary appointments scheduled for today.
      </div>
    </div>
  </div>
</template>

<style scoped>
.kiosk-page {
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;

  h1 {
    margin: 0;
    font-size: 2rem;
  }
}

.back-link {
  background: none;
  border: none;
  font-size: 1.2rem;
  color: var(--font-color-medium);
  cursor: pointer;
  font-weight: 600;
}

.date-display {
  font-size: 1.2rem;
  font-weight: 600;
  color: var(--purple);
  background: #f3e8ff;
  padding: 8px 16px;
  border-radius: 12px;
}

.schedule-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.apt-card {
  background: white;
  padding: 24px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 6px -2px rgba(0, 0, 0, 0.05);

  &.urgent {
    border-left: 6px solid #ef4444;
    background: #fef2f2;
  }
}

.time-col {
  width: 140px;
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--font-color-medium);
}

.info-col {
  flex: 1;
  h2 {
    margin: 0 0 6px 0;
    font-size: 1.6rem;
    color: var(--font-color-dark);
  }
}

.apt-details {
  font-size: 1.1rem;
  color: var(--font-color-medium);
}

.urgent-badge {
  background: #ef4444;
  color: white;
  font-weight: 700;
  padding: 6px 16px;
  border-radius: 50px;
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.empty-state {
  text-align: center;
  color: var(--font-color-medium);
  padding: 60px;
  font-size: 1.2rem;
  background: white;
  border-radius: 20px;
  border: 2px dashed #e2e8f0;
}
</style>
