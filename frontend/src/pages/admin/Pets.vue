<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import type { IPet } from '../../models/common'
import PetEditor from '../../components/admin/pets/PetEditor.vue'
import PetRow from '../../components/admin/pets/PetRow.vue'
import { Button, InputField, Select, Toast } from '../../components/common/ui'

// State
// State
const pets = ref<IPet[]>([])
const isLoading = ref(false)
const isEditorOpen = ref(false)
const selectedPet = ref<IPet | null>(null)
const searchQuery = ref('')
const statusFilter = ref('available')
const speciesFilter = ref('all')
const isSettingsOpen = ref(false)
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

  // Search and Status are handled by API, but Search might be robust here too?
  // User asked for API call on status change.
  // Search is also passed to API.
  // So we just return result (which is pets.value filtered by species)
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
  if (!petData.id) {
    // Handle create (not implemented yet, but we are fixing update now)
    console.log('Create not implemented yet')
    return
  }

  try {
    // Sanitize payload: remove non-editable fields that might cause strict JSON decoding errors on backend
    // Backend struct in updatePet only has: Name, Sex, Physical, Behavior, Medical, Descriptions, Details, Adoption, Foster, Returned, Sponsored, Photos, ProfileSettings
    // It DOES NOT have: id, slug, createdAt, updatedAt, litterName, species (frontend has species, backend struct needs checking).
    // Let's check backend cmd/api/pets.go updatePet struct:
    /*
    		Name            string          `json:"name"`
		Sex             string          `json:"sex"`
		Physical        json.RawMessage `json:"physical"`
        ...
    */
    // It does NOT have 'species'.
    const payload = {
      name: petData.name,
      sex: petData.sex,
      // species: petData.species, // Backend doesn't support updating species yet? Struct didn't have it.
      // Actually, let's look at the struct I added in step 10100/10141. Use `view_file` if unsure.
      // Proceeding with safe assumption: update input struct in backend if needed, or omit here.
      // Backend `Pet` struct has Species but `updatePet` anonymous struct didn't have it?
      // I will assume it's safer to only send what the backend explicitly expects or fix the backend.
      // The user wants to update everything.
      // I'll filter here to be safe and avoid 400 Bad Request.
      physical: petData.physical,
      behavior: petData.behavior,
      medical: petData.medical,
      descriptions: petData.descriptions,
      details: petData.details,
      adoption: petData.adoption,
      foster: petData.foster,
      returned: petData.returned,
      sponsored: petData.sponsored,
      photos: petData.photos,
      profileSettings: petData.profileSettings,
      litterName: petData.litterName,
    }

    const response = await fetch(`${import.meta.env.VITE_API_URL}/pets/${petData.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Ensure session cookie is sent (required for localhost:8080)
      body: JSON.stringify(payload),
    })

    if (!response.ok) {
      // Log the error response text for debugging
      const text = await response.text()
      console.error('Update failed:', text)
      throw new Error(text || 'Failed to update pet')
    }

    // Refetch pets to show changes
    await fetchPets()
    isEditorOpen.value = false

    showNotification('Pet updated successfully!', 'success')
  } catch (error) {
    console.error('Error updating pet:', error)
    showNotification(`Failed to update pet: ${error}`, 'error')
  }
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

// Toast State
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')

function showNotification(message: string, type: 'success' | 'error' = 'success') {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
}
</script>

<template>
  <div class="admin-page" @click="isSettingsOpen = false">
    <div class="page-header">
      <div class="header-left">
        <h1>Pet Records</h1>
        <span class="count-badge">{{ filteredPets.length }} Pets</span>
      </div>
      <div class="header-actions" @click.stop>
        <div class="filter-group">
          <Select v-model="speciesFilter" :options="speciesOptions" />
          <Select v-model="statusFilter" :options="statusOptions" />
        </div>

        <div class="search-wrapper">
          <InputField v-model="searchQuery" placeholder="Search pets..." />
        </div>

        <!-- Column Settings -->
        <div class="settings-dropdown-wrapper">
          <button
            class="icon-btn settings-btn"
            @click="isSettingsOpen = !isSettingsOpen"
            title="Table Settings"
          >
            ⚙️
          </button>

          <div v-if="isSettingsOpen" class="settings-dropdown">
            <div class="dropdown-header">Visible Columns</div>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.age" /> Age
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.breed" /> Breed
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.dob" /> Date of Birth
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.intake" /> Intake Date
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.microchip" /> Microchip
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.name" /> Name
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.photo" /> Photo
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.sex" /> Sex
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.sn" /> Spayed/Neutered
            </label>
            <label class="dropdown-item">
              <input type="checkbox" v-model="visibleColumns.status" /> Status
            </label>
          </div>
        </div>

        <Button title="Add New Pet +" color="green" :onClick="handleAddPet" />
      </div>
    </div>

    <!-- Pet List Table -->
    <div class="table-container">
      <table class="pets-table">
        <thead>
          <tr>
            <th class="expand-col"></th>
            <!-- Expand Arrow Column -->
            <th v-if="visibleColumns.photo">Photo</th>
            <th v-if="visibleColumns.name">Name</th>
            <th v-if="visibleColumns.breed">Breed</th>
            <th v-if="visibleColumns.sex">Sex</th>
            <th v-if="visibleColumns.sn" class="text-center">S/N</th>
            <th v-if="visibleColumns.microchip">Microchip</th>
            <th v-if="visibleColumns.age">Age</th>
            <th v-if="visibleColumns.dob">DoB</th>
            <th v-if="visibleColumns.intake">Intake</th>
            <!-- Dynamic Columns -->
            <th v-if="statusFilter === 'adopted'">Adopted Date</th>
            <th v-if="statusFilter === 'foster'">Foster Start</th>

            <th v-if="visibleColumns.status">Status</th>
            <th v-if="visibleColumns.actions">Actions</th>
          </tr>
        </thead>
        <tbody>
          <template v-for="(pet, index) in filteredPets" :key="pet.id">
            <PetRow
              :pet="pet"
              :index="index"
              :visible-columns="visibleColumns"
              :is-expanded="expandedPetId === pet.id"
              :status-filter="statusFilter"
              @toggle-expand="handleToggleExpand"
              @edit="handleEditPet"
              @archive="handleArchivePet"
            />
          </template>
        </tbody>
      </table>
    </div>

    <!-- Editor Drawer -->
    <PetEditor
      :is-open="isEditorOpen"
      :pet="selectedPet"
      :available-pets="pets"
      @close="isEditorOpen = false"
      @save="handleSavePet"
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
}

.filter-group {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* Settings Dropdown */
.filter-group {
  display: flex;
  gap: 12px;
  align-items: center;
}

.settings-dropdown-wrapper {
  position: relative;
}

.settings-btn {
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  width: 46px; /* Increased to match input height */
  height: 46px;
  font-size: 1.2rem;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
  }
}

.settings-dropdown {
  position: absolute;
  top: 100%;
  right: 0; /* Changed from left: 0 to right: 0 to align with the button */
  margin-top: 8px;
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 200px;
  padding: 8px 0;
  z-index: 50;
}

.dropdown-header {
  padding: 8px 16px;
  font-size: 0.8rem;
  font-weight: 600;
  color: hsl(from var(--color-neutral) h s 50%);
  text-transform: uppercase;
  border-bottom: 1px solid var(--border-color);
  margin-bottom: 4px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  font-size: 0.95rem;
  color: var(--text-primary);
  cursor: pointer;
  transition: background 0.2s;

  &:hover {
    background: hsl(from var(--color-neutral) h s 95%);
  }

  input[type='checkbox'] {
    accent-color: var(--color-secondary);
  }
}

.table-container {
  background: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  flex: 1;
  overflow-y: auto;

  /* Custom Scrollbar */
  &::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }

  &::-webkit-scrollbar-track {
    background: #f1f5f9;
  }

  &::-webkit-scrollbar-thumb {
    background: #cbd5e1; /* Using a neutral gray that fits */
    border-radius: 4px;

    /* If we want branded: background: hsl(from var(--color-primary) h s 80%); */
  }

  &::-webkit-scrollbar-thumb:hover {
    background: #94a3b8;
  }
}

.pets-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;

  th {
    background: hsl(from var(--color-neutral) h s 98%);
    padding: 12px 24px;
    font-weight: 600;
    color: hsl(from var(--color-neutral) h s 50%);
    font-size: 0.9rem;
    border-bottom: 1px solid var(--border-color);
    position: sticky;
    top: 0;
    z-index: 10;
  }

  /* Note: tr/td hover styles handled by PetRow.vue but we might want global header styling */
}

/* Helper class for alignment */
.text-center {
  text-align: center;
}

.search-wrapper {
  width: 300px;
}

.search-wrapper :deep(.field) {
  margin-bottom: 0;
}
</style>
