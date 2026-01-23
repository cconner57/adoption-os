<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

import PetCard from '../../components/admin/pets/PetCard.vue'
import PetEditor from '../../components/admin/pets/PetEditor.vue'
import PetTable from '../../components/admin/pets/PetTable.vue'
import TableSettings from '../../components/admin/pets/TableSettings.vue'
import { Button, InputField, Select, TableSkeleton,Toast } from '../../components/common/ui'
import type { IPet } from '../../models/common'

// Toast State
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')

function showNotification(message: string, type: 'success' | 'error' = 'success') {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
}
// State
const pets = ref<IPet[]>([])
const isLoading = ref(false)
const isEditorOpen = ref(false)
const selectedPet = ref<IPet | null>(null)
const searchQuery = ref('')
const statusFilter = ref('available')
const speciesFilter = ref('all')
const expandedPetId = ref<string | null>(null)

// Configurable Columns
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

// Fetch Data
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
        Authorization: `Bearer ${localStorage.getItem('token')}`, // Ensure auth if needed
      },
    })

    if (!response.ok) throw new Error('Failed to fetch pets')
    const data = await response.json()
    // Backend returns wrapped { data: [...] } or { message: ... }
    // If it's an array, use it directly (legacy fallback), otherwise checks data.data
    pets.value = Array.isArray(data) ? data : data.data || []
  } catch (error) {
    console.error('Error fetching pets:', error)
  } finally {
    isLoading.value = false
  }
}

// Load from LocalStorage & Initial Fetch
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
  fetchPets()
})

// Watch Filters to Refetch
watch([statusFilter, searchQuery], () => {
  // Debounce search if needed, but for now direct call
  fetchPets()
})

// Save to LocalStorage
watch(
  visibleColumns,
  (newVal) => {
    localStorage.setItem('petTableColumns', JSON.stringify(newVal))
  },
  { deep: true },
)

// Computed
const filteredPets = computed(() => {
  // Client-side Species filtering (since API doesn't support it yet)
  let result = pets.value

  if (speciesFilter.value !== 'all') {
    result = result.filter((p) => p.species === speciesFilter.value)
  }
  return result
})

// Actions
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
    const payload = { ...petData }
    let url = `${import.meta.env.VITE_API_URL}/pets`
    let method = 'POST'
    let successMessage = 'Pet added successfully!'

    if (petData.id) {
      // Update Mode
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

    // Parse the updated/created pet from response
    const savedPet = await response.json()
    const finalPet = savedPet.data || savedPet // Handle { data: ... } wrapper if present

    if (petData.id) {
      // Update existing in place
      const idx = pets.value.findIndex((p) => p.id === petData.id)
      if (idx !== -1) {
        pets.value[idx] = finalPet
      }
    } else {
      // Add new pet to the top
      pets.value.unshift(finalPet)
    }

    // No refetch needed
    isEditorOpen.value = false

    showNotification(successMessage, 'success')
  } catch (error) {
    console.error('Error saving pet:', error)
    showNotification(`Failed to save pet: ${error}`, 'error')
  }
}

async function handleQuickAdopt(pet: IPet) {
  if (!confirm(`Mark ${pet.name} as Adopted?`)) return

  // Create a copy with updated status
  // Backend will handle the date if missing
  const updatedPet = JSON.parse(JSON.stringify(pet))
  updatedPet.details.status = 'adopted'

  await handleSavePet(updatedPet)
}

function handleArchivePet(pet: IPet) {
  if (confirm(`Are you sure you want to archive ${pet.name}?`)) {
    // API call needed here

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
    <div class="page-header">
      <div class="header-left">
        <h1>Pet Records</h1>
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

        <Button title="New Pet +" color="green" :onClick="handleAddPet" />
      </div>
    </div>

    <!-- Pet List Table -->
    <TableSkeleton v-if="isLoading" :rows="10" :columns="11" />
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

    <!-- Mobile Card View -->
    <div v-if="!isLoading" class="mobile-pet-list">
      <template v-for="pet in filteredPets" :key="pet.id">
        <PetCard
          :pet="pet"
          :status-filter="statusFilter"
          @edit="handleEditPet"
          @archive="handleArchivePet"
          @mark-adopted="handleQuickAdopt"
        />
      </template>
    </div>

    <!-- Editor Drawer -->
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

  h1 {
    font-size: 1.8rem;
    color: var(--text-primary);
    margin: 0;
  }
}

.count-badge {
  background: hsl(from var(--color-neutral) h s 95%);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
  font-weight: 600;
}

.header-actions {
  display: flex;
  gap: 16px;
  position: relative;
  align-items: center;
  z-index: 20; /* Ensure this is higher than table */
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

/* Responsive Design */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .header-actions {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
    isolation: isolate; /* Create new stacking context */
  }

  .filter-group {
    width: 100%;
    /* overflow-x: auto; */
    padding-bottom: 4px;
    flex-wrap: wrap; /* Allow wrapping so dropdowns don't get cut off */
    /* Ensure visible overflow */
    overflow: visible;
  }

  .search-row {
    display: flex;
    gap: 8px;
    width: 100%;
  }

  .search-wrapper {
    flex: 1;
  }

  /* Hide table container on mobile since we show cards */
  /* This targets the component root of PetTable? No, PetTable has .table-container */
  :deep(.table-container) {
    display: none;
  }

  .mobile-pet-list {
    display: block;
    padding-bottom: 24px;
  }
}

@media (min-width: 769px) {
  .mobile-pet-list {
    display: none;
  }
}
</style>
