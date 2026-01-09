<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { mockVolunteers, mockShifts, mockIncidents } from '../../stores/mockVolunteerData'
import VolunteerList from '../../components/admin/volunteers/VolunteerList.vue'
import VolunteerDetail from '../../components/admin/volunteers/VolunteerDetail.vue'

const selectedVolunteerId = ref<string | null>(null)

const selectedVolunteer = computed(() => {
  return mockVolunteers.find((v) => v.id === selectedVolunteerId.value) || null
})

onMounted(() => {
  if (!selectedVolunteerId.value) {
    const defaults = mockVolunteers
      .filter((v) => v.status === 'active' || v.status === 'pending')
      .sort((a, b) => {
        if (b.reliabilityScore !== a.reliabilityScore)
          return b.reliabilityScore - a.reliabilityScore
        return a.firstName.localeCompare(b.firstName)
      })

    if (defaults.length > 0) {
      selectedVolunteerId.value = defaults[0].id
    }
  }
})

const selectedShifts = computed(() => {
  if (!selectedVolunteerId.value) return []
  return mockShifts.filter((s) => s.volunteerId === selectedVolunteerId.value)
})

const selectedIncidents = computed(() => {
  if (!selectedVolunteerId.value) return []
  return mockIncidents
    .filter((i) => i.volunteerId === selectedVolunteerId.value)
    .sort((a, b) => b.date.localeCompare(a.date))
})
</script>

<template>
  <div class="volunteers-page">
    <!-- Sidebar -->
    <aside class="sidebar">
      <VolunteerList
        :selectedId="selectedVolunteerId"
        @select="(vol) => (selectedVolunteerId = vol.id)"
      />
    </aside>

    <!-- Main Content -->
    <main class="main-content">
      <VolunteerDetail
        v-if="selectedVolunteer"
        :volunteer="selectedVolunteer"
        :shifts="selectedShifts"
        :incidents="selectedIncidents"
      />

      <div v-else class="empty-selection">
        <div class="icon">🤝</div>
        <h2>Select a volunteer</h2>
        <p>View their profile, schedule, and incident logs.</p>
      </div>
    </main>
  </div>
</template>

<style scoped>
.volunteers-page {
  display: flex;
  height: calc(100vh - 80px);
  gap: 24px;
  max-width: 1400px;
  margin: 0 auto;
}

.sidebar {
  width: 320px;
  background: var(--text-inverse);
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
  border: 1px solid var(--border-color);
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;

  h2 {
    font-size: 1.2rem;
    font-weight: 700;
    color: var(--text-primary);
    margin: 0;
  }
}

.invite-btn {
  background: var(--color-secondary);
  color: var(--text-inverse);
  border: none;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  &:hover {
    opacity: 0.9;
  }
}

.search-box {
  padding: 16px;
  border-bottom: 1px solid var(--border-color);
  input {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    background: hsl(from var(--color-neutral) h s 98%);
    outline: none;
    &:focus {
      border-color: var(--color-secondary);
      background: var(--text-inverse);
    }
  }
}

.main-content {
  flex: 1;
  background: var(--text-inverse);
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.03);
  border: 1px solid var(--border-color);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  scrollbar-gutter: stable;
}

.empty-selection {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: hsl(from var(--color-neutral) h s 50%);

  .icon {
    font-size: 4rem;
    margin-bottom: 24px;
    opacity: 0.5;
  }
  h2 {
    color: var(--text-primary);
    margin-bottom: 8px;
  }
}
</style>
<style scoped>
.admin-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  h1 {
    font-size: 1.8rem;
    color: var(--text-primary);
    margin: 0;
  }
}

.add-btn {
  background-color: var(--color-secondary); /* Assuming purple mapped to secondary */
  color: var(--text-inverse);
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: hsl(from var(--color-secondary) h s 40%);
  }
}

.placeholder-content {
  background: var(--text-inverse);
  padding: 48px;
  border-radius: 12px;
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
  border: 1px dashed var(--border-color);
}
</style>
