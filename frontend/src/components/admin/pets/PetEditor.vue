<script setup lang="ts">
import { ref, watch } from 'vue'
import type { IPet } from '../../../models/common'
import PetEditorBasic from './editor/PetEditorBasic.vue'
import PetEditorPhysical from './editor/PetEditorPhysical.vue'
import PetEditorBehavior from './editor/PetEditorBehavior.vue'
import PetEditorMedical from './editor/PetEditorMedical.vue'
import PetEditorStatus from './editor/PetEditorStatus.vue'
import PetEditorStory from './editor/PetEditorStory.vue'
import PetEditorPhotos from './editor/PetEditorPhotos.vue'
import PetEditorSettings from './editor/PetEditorSettings.vue'

const props = defineProps<{
  pet: IPet | null
  isOpen: boolean
  availablePets?: IPet[]
}>()

const emit = defineEmits(['close', 'save'])

const activeTab = ref('basic')
const formData = ref<Partial<IPet>>({})

// Initialize form data when pet changes or modal opens
watch(
  () => props.pet,
  (newPet) => {
    if (newPet) {
      // Merge with defaults to ensure all fields exist
      const defaults = {
        profileSettings: {
          isSpotlightFeatured: false,
          showMedicalHistory: false,
          showAdditionalInformation: false,
        },
        medical: {
          microchip: { microchipped: false },
        },
        behavior: {
          energyLevel: 'medium',
        },
      }
      formData.value = {
        ...JSON.parse(JSON.stringify(newPet)),
      }
      // Shallow merge specific nested objects if missing
      if (!formData.value.profileSettings) formData.value.profileSettings = defaults.profileSettings
      else {
        // Ensure properties exist
        if (formData.value.profileSettings.isSpotlightFeatured === undefined)
          formData.value.profileSettings.isSpotlightFeatured = false
        if (formData.value.profileSettings.showMedicalHistory === undefined)
          formData.value.profileSettings.showMedicalHistory = false
        if (formData.value.profileSettings.showAdditionalInformation === undefined)
          formData.value.profileSettings.showAdditionalInformation = false
      }

      // Ensure Medical object structure
      if (!formData.value.medical) {
        formData.value.medical = {
          spayedOrNeutered: false,
          vaccinationsUpToDate: false,
          vaccinations: {
            other: [],
            rabies: { dateAdministered: '' },
          },
          surgeries: [],
          microchip: { microchipped: false },
        }
      } else {
        // Deep ensure microchip
        if (!formData.value.medical.microchip) {
          formData.value.medical.microchip = { microchipped: false }
        }
        // Ensure vaccinations
        if (!formData.value.medical.vaccinations) {
          formData.value.medical.vaccinations = { other: [], rabies: { dateAdministered: '' } }
        } else {
          if (!formData.value.medical.vaccinations.other)
            formData.value.medical.vaccinations.other = []
          if (!formData.value.medical.vaccinations.rabies)
            formData.value.medical.vaccinations.rabies = { dateAdministered: '' }
        }
      }

      // Ensure Behavior object exists
      if (!formData.value.behavior) {
        formData.value.behavior = {
          energyLevel: 'medium',
          isGoodWithCats: null,
          isGoodWithDogs: null,
          isGoodWithKids: null,
          isHouseTrained: null,
          personalityTags: [],
          bonded: { isBonded: false, bondedWith: [] },
        }
      } else {
        if (!formData.value.behavior.bonded) {
          formData.value.behavior.bonded = { isBonded: false, bondedWith: [] }
        }
        if (!formData.value.behavior.personalityTags) {
          formData.value.behavior.personalityTags = []
        }
      }

      // Ensure Adoption object exists
      if (!formData.value.adoption) {
        formData.value.adoption = {
          adopterContactInfo: { name: '', email: '', phone: '' },
        }
      } else if (!formData.value.adoption.adopterContactInfo) {
        formData.value.adoption.adopterContactInfo = { name: '', email: '', phone: '' }
      }

      // Ensure Foster object exists
      if (!formData.value.foster) formData.value.foster = {}

      // Ensure details
      if (!formData.value.details) formData.value.details = { status: 'available' }

      // Ensure descriptions
      if (!formData.value.descriptions) formData.value.descriptions = { primary: '' }

      // Ensure photos
      if (!formData.value.photos) formData.value.photos = []

      // Ensure returned
      if (!formData.value.returned) {
        formData.value.returned = { isReturned: false, history: [] }
      }
      // Migration for old returned data if needed (though backend handles JSONB, frontend needs array)
      if (!Array.isArray(formData.value.returned.history)) {
        formData.value.returned.history = []
      }
    } else {
      // Reset to Defaults for New Pet
      activeTab.value = 'basic'
      formData.value = {
        name: '',
        species: 'cat',
        sex: 'unknown',
        physical: {
          breed: 'Unknown',
          size: 'medium',
          ageGroup: 'baby',
          color: '',
          coatLength: 'short',
        },
        medical: {
          spayedOrNeutered: false,
          vaccinationsUpToDate: false,
          vaccinations: {
            other: [],
            rabies: { dateAdministered: '' },
          },
          surgeries: [],
          microchip: { microchipped: false },
        },
        behavior: {
          energyLevel: 'medium',
          personalityTags: [],
          bonded: { isBonded: false, bondedWith: [] },
        },
        descriptions: {
          primary: '',
          spotlight: '',
          behavioral: '',
          fun: '',
          origin: '',
          additionalInformation: [],
          specialNeeds: '',
        },
        details: {
          status: 'available',
          intakeDate: '',
          shelterLocation: '',
          preferredPetLitterType: '',
          environmentType: null,
        },
        adoption: {
          adoptedBy: '',
          date: '',
          newAdoptedName: '',
          fee: null,
          surveyCompleted: false,
          adopterContactInfo: {
            name: '',
            email: '',
            phone: '',
          },
        },
        foster: {
          startDate: '',
          endDate: '',
          parentName: '',
        },
        returned: {
          isReturned: false,
          history: [],
        },
        sponsored: {
          isSponsored: false,
          sponsoredBy: '',
          date: '',
          amount: null,
        },
        photos: [],
        profileSettings: {
          isSpotlightFeatured: false,
          showMedicalHistory: false,
          showAdditionalInformation: false,
        },
      }
    }
  },
  { immediate: true },
)

function handleSave() {
  emit('save', formData.value)
}

function handleClose() {
  emit('close')
}

// Tabs configuration
const tabs = [
  { id: 'basic', label: 'Basic Info', icon: '📝' },
  { id: 'physical', label: 'Physical', icon: '📏' },
  { id: 'behavior', label: 'Behavior', icon: '🧠' },
  { id: 'medical', label: 'Medical', icon: '⚕️' },
  { id: 'status', label: 'Status', icon: '🏠' },
  { id: 'descriptions', label: 'Story', icon: '📖' },
  { id: 'photos', label: 'Photos', icon: '📸' },
  { id: 'settings', label: 'Settings', icon: '⚙️' },
]
</script>

<template>
  <div class="pet-editor-overlay" :class="{ open: isOpen }" @click.self="handleClose">
    <div class="pet-editor-drawer">
      <div class="drawer-header">
        <h2>{{ pet ? `Edit ${pet.name}` : 'Add New Pet' }}</h2>
        <button class="close-btn" @click="handleClose">✕</button>
      </div>

      <div class="drawer-body">
        <div class="sidebar-nav">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            class="tab-btn"
            :class="{ active: activeTab === tab.id }"
            @click="activeTab = tab.id"
          >
            <span class="tab-icon">{{ tab.icon }}</span>
            {{ tab.label }}
          </button>
        </div>

        <div class="tab-content" v-if="formData">
          <form @submit.prevent>
            <PetEditorBasic
              v-if="activeTab === 'basic'"
              v-model="formData"
              :available-pets="availablePets"
            />
            <PetEditorPhysical v-if="activeTab === 'physical'" v-model="formData" />
            <PetEditorBehavior
              v-if="activeTab === 'behavior'"
              v-model="formData"
              :available-pets="availablePets"
            />
            <PetEditorMedical v-if="activeTab === 'medical'" v-model="formData" />
            <PetEditorStatus v-if="activeTab === 'status'" v-model="formData" />
            <PetEditorStory v-if="activeTab === 'descriptions'" v-model="formData" />
            <PetEditorPhotos v-if="activeTab === 'photos'" v-model="formData" />
            <PetEditorSettings v-if="activeTab === 'settings'" v-model="formData" />
          </form>
        </div>
      </div>

      <div class="drawer-footer">
        <button class="btn-cancel" @click="handleClose">Cancel</button>
        <button class="btn-save" @click="handleSave">Save Pet</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.pet-editor-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: flex-end;
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.3s;
  z-index: 1000;
}

.pet-editor-overlay.open {
  opacity: 1;
  pointer-events: auto;
}

.pet-editor-drawer {
  width: 900px;
  max-width: 95vw;
  background: white;
  height: 100%;
  box-shadow: -4px 0 20px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  transform: translateX(100%);
  transition: transform 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.pet-editor-overlay.open .pet-editor-drawer {
  transform: translateX(0);
}

.drawer-header {
  padding: 20px 24px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  z-index: 10;
}

.drawer-header h2 {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  color: var(--text-primary);
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: hsl(from var(--color-neutral) h s 50%);
}

.drawer-body {
  flex: 1;
  overflow: hidden; /* Hide outer overflow, inner sections scroll */
  display: flex;
  flex-direction: row; /* Horizontal Layout */
}

/* SIDEBAR NAV */
.sidebar-nav {
  width: 200px;
  background: hsl(from var(--color-neutral) h s 98%);
  display: flex;
  flex-direction: column;
  padding: 16px 0;
  border-right: 1px solid var(--border-color);
  overflow-y: auto;
  flex-shrink: 0;
}

.tab-btn {
  padding: 12px 24px;
  background: none;
  border: none;
  border-left: 3px solid transparent;
  cursor: pointer;
  font-weight: 500;
  color: hsl(from var(--color-neutral) h s 50%);
  text-align: left;
  display: flex;
  align-items: center;
  gap: 10px;
  transition: all 0.2s;
  font-size: 0.95rem;

  &.active {
    color: var(--color-primary);
    background: hsla(from var(--color-primary) h s l / 0.05);
    border-left-color: var(--color-primary);
    font-weight: 600;
  }

  &:hover:not(.active) {
    background: hsl(from var(--color-neutral) h s 95%);
    color: var(--text-primary);
  }

  .tab-icon {
    font-size: 1.1rem;
    width: 24px;
    text-align: center;
  }
}

.tab-content {
  flex: 1;
  padding: 32px;
  overflow-y: auto;
  background: var(--text-inverse);
}

/* Form Sections are now inside components but layout styling remains generic */

.drawer-footer {
  padding: 16px 24px;
  border-top: 1px solid var(--border-color);
  background: white;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  z-index: 10;
}

.btn-cancel {
  background: none;
  border: 1px solid var(--border-color);
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  color: var(--text-primary);
}

.btn-save {
  background: var(--color-primary);
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  box-shadow: 0 2px 4px hsla(from var(--color-primary) h s l / 0.3);
}

/* Scrollbar Styling */
.sidebar-nav::-webkit-scrollbar,
.tab-content::-webkit-scrollbar {
  width: 6px;
}
.sidebar-nav::-webkit-scrollbar-thumb,
.tab-content::-webkit-scrollbar-thumb {
  background-color: var(--border-color);
  border-radius: 3px;
}
</style>
