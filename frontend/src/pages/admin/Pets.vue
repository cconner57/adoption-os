<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { mockPetsData } from '../../stores/mockPetData'
import type { IPet } from '../../models/common'
import PetEditor from '../../components/admin/pets/PetEditor.vue'
import PetRow from '../../components/admin/pets/PetRow.vue'

// State
const pets = ref<IPet[]>([...mockPetsData])
const isEditorOpen = ref(false)
const selectedPet = ref<IPet | null>(null)
const searchQuery = ref('')
const statusFilter = ref('all')
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
  actions: true, // Generally always visible, but configurable if desired
})

// Load from LocalStorage
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
  let result = pets.value

  // 1. Filter by Status
  if (statusFilter.value !== 'all') {
    result = result.filter((p) => p.details.status === statusFilter.value)
  }

  // 2. Filter by Species
  if (speciesFilter.value !== 'all') {
    result = result.filter((p) => p.species === speciesFilter.value)
  }

  // 3. Filter by Search Query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(
      (p) =>
        p.name.toLowerCase().includes(query) ||
        p.physical.breed?.toLowerCase().includes(query) ||
        p.details.status.toLowerCase().includes(query),
    )
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

function handleSavePet(petData: Partial<IPet>) {
  if (selectedPet.value) {
    // Edit Mode
    const index = pets.value.findIndex((p) => p.id === selectedPet.value?.id)
    if (index !== -1) {
      pets.value[index] = { ...selectedPet.value, ...petData } as IPet
    }
  } else {
    // Create Mode
    const newPet = {
      ...petData,
      id: Math.random().toString(36).substr(2, 9),
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    } as IPet
    pets.value.unshift(newPet)
  }
  isEditorOpen.value = false
}

function handleArchivePet(pet: IPet) {
  if (confirm(`Are you sure you want to archive ${pet.name}?`)) {
    const index = pets.value.findIndex((p) => p.id === pet.id)
    if (index !== -1) {
      pets.value[index].details.status = 'archived'
    }
  }
}

function handleToggleExpand(pet: IPet) {
  if (expandedPetId.value === pet.id) {
    expandedPetId.value = null
  } else {
    expandedPetId.value = pet.id
  }
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
        <select v-model="speciesFilter" class="filter-select">
          <option value="all">All Species</option>
          <option value="cat">Cats</option>
          <option value="dog">Dogs</option>
        </select>

        <select v-model="statusFilter" class="filter-select">
          <option value="all">All Statuses</option>
          <option value="available">Available</option>
          <option value="adoption-pending">Adoption Pending</option>
          <option value="adopted">Adopted</option>
          <option value="foster">In Foster</option>
          <option value="hold">Medical/Behavioral Hold</option>
          <option value="intake">Intake Processing</option>
          <option value="archived">Archived</option>
        </select>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search pets..."
          class="search-input"
        />

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

        <button class="add-btn" @click="handleAddPet">Add New Pet +</button>
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
      @close="isEditorOpen = false"
      @save="handleSavePet"
    />
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
    color: var(--font-color-dark);
    margin: 0;
  }
}

.count-badge {
  background: #f3f4f6;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
  color: var(--font-color-medium);
  font-weight: 600;
}

.header-actions {
  display: flex;
  gap: 16px;
  position: relative;
}

/* Settings Dropdown */
.settings-dropdown-wrapper {
  position: relative;
}

.settings-btn {
  background: white;
  border: 1px solid #e5e7eb;
  width: 42px; /* Match height of inputs/selects approx */
  height: 42px;
  font-size: 1.2rem;
  border-radius: 6px;
  cursor: pointer;

  &:hover {
    background: #f9fafb;
  }
}

.settings-dropdown {
  position: absolute;
  top: 100%;
  right: 0; /* Changed from left: 0 to right: 0 to align with the button */
  margin-top: 8px;
  background: white;
  border: 1px solid #e5e7eb;
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
  color: var(--font-color-medium);
  text-transform: uppercase;
  border-bottom: 1px solid #f3f4f6;
  margin-bottom: 4px;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  font-size: 0.95rem;
  color: var(--font-color-dark);
  cursor: pointer;
  transition: background 0.2s;

  &:hover {
    background: #f3f4f6;
  }

  input[type='checkbox'] {
    accent-color: var(--blue);
  }
}

.search-input {
  padding: 10px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  width: 250px;
  outline: none;
  font-size: 0.95rem;
  background-color: white;
  color: var(--font-color-dark);

  &:focus {
    border-color: var(--blue);
  }

  &::placeholder {
    color: #9ca3af;
  }
}

.filter-select {
  padding: 10px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  outline: none;
  font-size: 0.95rem;
  background-color: white;
  color: var(--font-color-dark);
  cursor: pointer;

  &:focus {
    border-color: var(--blue);
  }
}

.add-btn {
  background-color: var(--green);
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: var(--green-hover);
  }
}

.table-container {
  background: white;
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

    /* If we want branded: background: var(--green-weak); */
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
    background: #f9fafb;
    padding: 12px 24px;
    font-weight: 600;
    color: var(--font-color-medium);
    font-size: 0.9rem;
    border-bottom: 1px solid #e5e7eb;
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
</style>
