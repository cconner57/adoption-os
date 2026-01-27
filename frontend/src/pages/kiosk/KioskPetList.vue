<script setup lang="ts">
import { computed,ref } from 'vue'
import { useRouter } from 'vue-router'

import { Capsules } from '../../components/common/ui'
import { mockPetsData } from '../../stores/mockPetData' 

const mockPets = ref(mockPetsData)

const router = useRouter()
const searchQuery = ref('')

const filteredPets = computed(() => {
  return mockPets.value.filter(
    (p) =>
      p.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (p.physical.breed &&
        p.physical.breed.toLowerCase().includes(searchQuery.value.toLowerCase())),
  )
})

const calculateAge = (dob: string | null | undefined): string => {
  if (!dob) return 'Unknown Age'
  if (dob.length < 5) return dob 
  const birthDate = new Date(dob)
  if (isNaN(birthDate.getTime())) return dob

  const today = new Date()
  let age = today.getFullYear() - birthDate.getFullYear()
  const m = today.getMonth() - birthDate.getMonth()
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  return `${age} yrs`
}

const formattedPets = computed(() => {
  return filteredPets.value.map((p) => ({
    id: p.id,
    name: p.name,
    species: p.species,
    status: p.details.status,
    breed: p.physical.breed,
    age: calculateAge(p.physical.dateOfBirth),
    kennelNo: p.details.shelterLocation,
    weight: p.physical.currentWeight,
    medicalNotes: p.medical.healthConcerns?.join(', ') || null,
  }))
})

const goHome = () => router.push('/kiosk')
</script>

<template>
  <div class="kiosk-page">
    <div class="page-header">
      <button class="back-link" @click="goHome">← Back</button>
      <h1>Resident Pets</h1>
      <div class="search-box">
        <span class="icon">🔍</span>
        <input v-model="searchQuery" placeholder="Search by name or breed..." />
      </div>
    </div>

    <div class="pets-grid">
      <div v-for="pet in formattedPets" :key="pet.id" class="pet-card">
        <div class="img-area">
          
          <span class="emoji-avatar">{{ pet.species === 'dog' ? '🐶' : '🐱' }}</span>
        </div>
        <div class="pet-info">
          <div class="top-row">
            <h2>{{ pet.name }}</h2>
            <Capsules
              :label="pet.status"
              color="var(--color-primary-weak)"
              size="sm"
            />
          </div>
          <p class="breed">{{ pet.breed || 'Unknown Breed' }} • {{ pet.age }}</p>
          <div class="details">
            <span>📍 {{ pet.kennelNo || 'N/A' }}</span>
            <span>⚖️ {{ pet.weight ? pet.weight + ' lbs' : '-- lbs' }}</span>
          </div>

          <div v-if="pet.medicalNotes" class="med-note">⚠️ {{ pet.medicalNotes }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.kiosk-page {
  max-width: 1100px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;

  h1 {
    margin: 0;
    font-size: 2rem;
    color: var(--text-primary);
  }
}

.back-link {
  background: none;
  border: none;
  font-size: 1.2rem;
  color: var(--color-neutral-text-soft);
  cursor: pointer;
  font-weight: 600;
}

.search-box {
  background: #fff;
  padding: 12px 20px;
  border-radius: 50px;
  border: 1px solid #cbd5e1;
  display: flex;
  align-items: center;
  gap: 12px;
  width: 350px;
  box-shadow: 0 4px 6px -2px rgb(0 0 0 / 5%);

  .icon {
    font-size: 1.2rem;
  }

  input {
    border: none;
    outline: none;
    font-size: 1.1rem;
    width: 100%;
    color: var(--text-primary);
  }
}

.pets-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
}

.pet-card {
  background: #fff;
  border-radius: 20px;
  overflow: hidden;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 5%);
  display: flex;
  flex-direction: column;
}

.img-area {
  height: 160px;
  background: #f1f5f9;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
}

.pet-info {
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.top-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;

  h2 {
    margin: 0;
    font-size: 1.4rem;
    color: var(--text-primary);
  }
}

.breed {
  color: var(--color-neutral-text-soft);
  margin: 0 0 16px;
  font-size: 1rem;
}

.details {
  display: flex;
  gap: 16px;
  font-size: 0.95rem;
  font-weight: 500;
  color: #475569;
  margin-bottom: 12px;
}

.med-note {
  margin-top: auto;
  background: #fef2f2;
  color: #b91c1c;
  padding: 8px 12px;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
}
</style>
