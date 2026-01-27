<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

import PetCard from '../../components/admin/pets/PetCard.vue'
import PetCardSkeleton from '../../components/admin/pets/PetCardSkeleton.vue'
import PetEditor from '../../components/admin/pets/PetEditor.vue'
import PetTable from '../../components/admin/pets/PetTable.vue'
import TableSettings from '../../components/admin/pets/TableSettings.vue'
import { Button, InputField, Select, TableSkeleton,Toast } from '../../components/common/ui'
import type { IPet } from '../../models/common'

const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')

function showNotification(message: string, type: 'success' | 'error' = 'success') {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
}

const pets = ref<IPet[]>([])
const isLoading = ref(false)
const isEditorOpen = ref(false)
const selectedPet = ref<IPet | null>(null)
const searchQuery = ref('')
const statusFilter = ref('available')
const speciesFilter = ref('all')
const expandedPetId = ref<string | null>(null)
const isMounted = ref(false)

const visibleColumns = ref({
  photo: true,
  name: true,
  breed: true,
  sex: true,
  sn: true,
  microchip: true,
  age: true,
  dob: true,
  intake: true,
  status: true,
  actions: true,
})

async function fetchPets() {
  isLoading.value = true
  try {
    const params = new URLSearchParams()
    if (statusFilter.value !== 'all') {
      params.append('status', statusFilter.value)
    }
    if (searchQuery.value) {
      params.append('search', searchQuery.value)
    }

    const response = await fetch(`${import.meta.env.VITE_API_URL}/pets?${params.toString()}`, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    })

    if (!response.ok) throw new Error('Failed to fetch pets')
    const data = await response.json()

    pets.value = Array.isArray(data) ? data : data.data || []
  } catch (error) {
    console.error('Error fetching pets:', error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  const saved = localStorage.getItem('petTableColumns')
  if (saved) {
    try {
      const parsed = JSON.parse(saved)
      visibleColumns.value = { ...visibleColumns.value, ...parsed }
    } catch (e) {
      console.error('Failed to parse saved columns', e)
    }
  }
  isMounted.value = true
  fetchPets()
})

watch([statusFilter, searchQuery], () => {

  fetchPets()
})

watch(
  visibleColumns,
  (newVal) => {
    localStorage.setItem('petTableColumns', JSON.stringify(newVal))
  },
  { deep: true },
)

const filteredPets = computed(() => {

  let result = pets.value

  if (speciesFilter.value !== 'all') {
    result = result.filter((p) => p.species === speciesFilter.value)
  }
  return result
})

function handleAddPet() {
  selectedPet.value = null
  isEditorOpen.value = true
}

function handleEditPet(pet: IPet) {
  selectedPet.value = pet
  isEditorOpen.value = true
}

async function handleSavePet(petData: Partial<IPet>) {
  try {
    // Exclude read-only fields that backend might reject or that we don't want to send
    // eslint-disable-next-line no-unused-vars, @typescript-eslint/no-unused-vars
    const { id, createdAt, updatedAt, ...rest } = petData
    const payload = { ...rest }

    let url = `${import.meta.env.VITE_API_URL}/pets`
    let method = 'POST'
    let successMessage = 'Pet added successfully!'

    if (petData.id) {
      url = `${import.meta.env.VITE_API_URL}/pets/${petData.id}`
      method = 'PUT'
      successMessage = 'Pet updated successfully!'
    }

    const response = await fetch(url, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
      credentials: 'include',
      body: JSON.stringify(payload),
    })

    if (!response.ok) {
      const text = await response.text()
      console.error(`${method} failed:`, text)
      throw new Error(text || `Failed to ${petData.id ? 'update' : 'create'} pet`)
    }

    const savedPet = await response.json()
    const finalPet = savedPet.data || savedPet

    if (petData.id) {

      const idx = pets.value.findIndex((p) => p.id === petData.id)
      if (idx !== -1) {
        pets.value[idx] = finalPet
      }
    } else {

      pets.value.unshift(finalPet)
    }

    isEditorOpen.value = false

    showNotification(successMessage, 'success')
  } catch (error) {
    console.error('Error saving pet:', error)
    showNotification(`Failed to save pet: ${error}`, 'error')
  }
}

async function handleQuickAdopt(pet: IPet) {
  if (!confirm(`Mark ${pet.name} as Adopted?`)) return

  // Use structuredClone to ensure we have a clean object without reactivity/proxies
  const updatedPet = structuredClone(pet)

  // Update status
  if (!updatedPet.details) updatedPet.details = { status: 'adopted' }
  updatedPet.details.status = 'adopted'

  // Set default adoption date if missing
  if (!updatedPet.adoption) updatedPet.adoption = {}
  if (!updatedPet.adoption.date) {
    updatedPet.adoption.date = new Date().toISOString().split('T')[0]
  }

  await handleSavePet(updatedPet)
}

function handleArchivePet(pet: IPet) {
  if (confirm(`Are you sure you want to archive ${pet.name}?`)) {

    console.log('Archive logic pending API implementation')
  }
}

function handleToggleExpand(pet: IPet) {
  if (expandedPetId.value === pet.id) {
    expandedPetId.value = null
  } else {
    expandedPetId.value = pet.id
  }
}

const speciesOptions = [
  { label: 'All Species', value: 'all' },
  { label: 'Cats', value: 'cat' },
  { label: 'Dogs', value: 'dog' },
]

const statusOptions = [
  { label: 'All Statuses', value: 'all' },
  { label: 'Available', value: 'available' },
  { label: 'Adoption Pending', value: 'adoption-pending' },
  { label: 'Adopted', value: 'adopted' },
  { label: 'In Foster', value: 'foster' },
  { label: 'Medical/Behavioral Hold', value: 'hold' },
  { label: 'Intake Processing', value: 'intake' },
  { label: 'Archived', value: 'archived' },
]

</script>

<template>
  <div class="admin-page">
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Pet Records</h1>
    </Teleport>

    <div class="page-header">
      <div class="header-left">
        <h1 class="desktop-only">Pet Records</h1>
        <span class="count-badge">{{ filteredPets.length }} Pets</span>
      </div>
      <div class="header-actions">
        <div class="filter-group">
          <Select v-model="speciesFilter" :options="speciesOptions" />
          <Select v-model="statusFilter" :options="statusOptions" />
        </div>

        <div class="search-row">
          <div class="search-wrapper">
            <InputField v-model="searchQuery" placeholder="Search pets..." />
          </div>

          <TableSettings v-model="visibleColumns" />
        </div>

        <Button title="New Pet +" color="green" :onClick="handleAddPet" size="large" />
      </div>
    </div>

    <div class="desktop-skeleton-wrapper" v-if="isLoading">
      <TableSkeleton :rows="10" :columns="11" />
    </div>

    <PetTable
      v-else
      :pets="filteredPets"
      :visible-columns="visibleColumns"
      :expanded-pet-id="expandedPetId"
      :status-filter="statusFilter"
      @toggle-expand="handleToggleExpand"
      @edit="handleEditPet"
      @archive="handleArchivePet"
      @mark-adopted="handleQuickAdopt"
    />

    <div class="mobile-pet-list">
      <template v-if="isLoading">
        <PetCardSkeleton v-for="n in 5" :key="n" />
      </template>
      <template v-else>
        <PetCard
          v-for="pet in filteredPets"
          :key="pet.id"
          :pet="pet"
          :status-filter="statusFilter"
          @edit="handleEditPet"
          @archive="handleArchivePet"
          @mark-adopted="handleQuickAdopt"
        />
      </template>
    </div>

    <PetEditor
      :is-open="isEditorOpen"
      :pet="selectedPet"
      :available-pets="pets"
      @close="isEditorOpen = false"
      @save="handleSavePet"
      @archive="handleArchivePet"
    />

    <Toast :show="showToast" :message="toastMessage" :type="toastType" @close="showToast = false" />
  </div>
</template>

<style scoped>
.admin-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: calc(100vh - 100px);
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;

  h1.desktop-only {
    font-size: 1.8rem;
    color: var(--text-primary);
    margin: 0;
  }
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .header-left h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.count-badge {
  background: var(--color-neutral-weak);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
  color: var(--color-neutral-text-soft);
  font-weight: 600;
}

.header-actions {
  display: flex;
  gap: 16px;
  position: relative;
  align-items: center;
  z-index: 20;
}

.filter-group {
  display: flex;
  gap: 12px;
  align-items: center;
}

.search-wrapper {
  width: 300px;
}

.search-wrapper :deep(.field) {
  margin-bottom: 0;
}

.search-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

@media (width <= 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .header-actions {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
    isolation: isolate;
  }

  .filter-group {
    width: 100%;
    padding-bottom: 4px;
    flex-wrap: wrap;
    overflow: visible;
  }

  .filter-group :deep(.select-container) {
    width: 100% !important;
  }

  .search-row {
    display: flex;
    gap: 8px;
    width: 100%;
  }

  .search-wrapper {
    flex: 1;
  }

  :deep(.table-container),
  .desktop-skeleton-wrapper {
    display: none;
  }

  .mobile-pet-list {
    display: block;
    padding-bottom: 24px;
  }
}

@media (width >= 769px) {
  .mobile-pet-list {
    display: none;
  }
}
</style>
