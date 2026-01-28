<script setup lang="ts">
/* eslint-disable */
import { ref, watch } from 'vue'
import type { IPet, TEnvironment, TPetStatus, TSpecies } from '../../../models/common'
import type { TPetBreed } from '../../../constants/breeds'
import { EditorDrawer, SidebarNav } from '../../common/ui'
import { useSettingsStore } from '../../../stores/settings'
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

const emit = defineEmits(['close', 'save', 'archive'])

const settingsStore = useSettingsStore()
const activeTab = ref('basic')
const formData = ref<Partial<IPet>>({})

function initFormData(newPet: IPet) {

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

  if (formData.value.litterName === undefined) formData.value.litterName = null

  if (!formData.value.profileSettings) formData.value.profileSettings = defaults.profileSettings
  else {

    if (formData.value.profileSettings.isSpotlightFeatured === undefined)
      formData.value.profileSettings.isSpotlightFeatured = false
    if (formData.value.profileSettings.showMedicalHistory === undefined)
      formData.value.profileSettings.showMedicalHistory = false
    if (formData.value.profileSettings.showAdditionalInformation === undefined)
      formData.value.profileSettings.showAdditionalInformation = false
  }

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

    if (!formData.value.medical.microchip) {
      formData.value.medical.microchip = { microchipped: false }
    }

    if (!formData.value.medical.vaccinations) {
      formData.value.medical.vaccinations = { other: [], rabies: { dateAdministered: '' } }
    } else {
      if (!formData.value.medical.vaccinations.other) formData.value.medical.vaccinations.other = []
      if (!formData.value.medical.vaccinations.rabies)
        formData.value.medical.vaccinations.rabies = { dateAdministered: '' }
    }
  }

  if (!formData.value.behavior) {
    formData.value.behavior = {
      energyLevel: 'medium',
      isGoodWithCats: null,
      isGoodWithDogs: null,
      isGoodWithKids: null,
      isHouseTrained: null,
      personalityTags: [],
      bonded: { isBonded: false, bondedWith: [] },
      prefersToBeAlone: false,
    }
  } else {
    if (!formData.value.behavior.bonded) {
      formData.value.behavior.bonded = { isBonded: false, bondedWith: [] }
    }
    if (!formData.value.behavior.personalityTags) {
      formData.value.behavior.personalityTags = []
    }
  }

  if (!formData.value.adoption) {
    formData.value.adoption = {
      adoptedBy: '',
      date: '',
      newAdoptedName: '',
      fee: null,
      surveyCompleted: false,
      adopterContactInfo: { name: '', email: '', phone: '' },
    }
  } else {

    if (!formData.value.adoption.adopterContactInfo) {
      formData.value.adoption.adopterContactInfo = { name: '', email: '', phone: '' }
    }
    if (formData.value.adoption.date === undefined) formData.value.adoption.date = ''
    if (formData.value.adoption.adoptedBy === undefined) formData.value.adoption.adoptedBy = ''
    if (formData.value.adoption.newAdoptedName === undefined)
      formData.value.adoption.newAdoptedName = ''
    if (formData.value.adoption.surveyCompleted === undefined)
      formData.value.adoption.surveyCompleted = false
  }

  if (!formData.value.foster) formData.value.foster = {}

  if (!formData.value.details) formData.value.details = { status: 'available' }

  if (!formData.value.descriptions) formData.value.descriptions = { primary: '' }

  if (!formData.value.photos) formData.value.photos = []

  if (!formData.value.returned) {
    formData.value.returned = { isReturned: false, history: [] }
  }

  if (!Array.isArray(formData.value.returned.history)) {
    formData.value.returned.history = []
  }
}

watch(
  () => props.pet,
  (newPet) => {
    if (newPet) {
      initFormData(newPet)
    } else {

      const defaults = settingsStore.settings.pets || {}

      activeTab.value = 'basic'
      formData.value = {
        name: '',
        species: (defaults.defaultSpecies as unknown as TSpecies) || 'cat',
        sex: 'unknown',
        physical: {
          breed: (defaults.defaultBreed as unknown as TPetBreed) || 'Unknown',
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
          microchip: { microchipped: defaults.requireMicrochip || false },
        },
        behavior: {
          energyLevel: 'medium',
          personalityTags: [],
          bonded: { isBonded: false, bondedWith: [] },
          isGoodWithCats: null,
          isGoodWithDogs: null,
          isGoodWithKids: null,
          isHouseTrained: null,
          prefersToBeAlone: false,
        },
        descriptions: {
          primary: defaults.defaultBioTemplate || '',
          spotlight: '',
          behavioral: '',
          fun: '',
          origin: '',
          additionalInformation: [],
          specialNeeds: '',
        },
        details: {
          status: (defaults.defaultIntakeStatus as unknown as TPetStatus) || 'available',
          intakeDate: new Date().toISOString().split('T')[0],
          shelterLocation: defaults.defaultShelterLocation || '',
          preferredPetLitterType: '',
          environmentType: (defaults.defaultEnvironment as unknown as TEnvironment) || null,
        },
        adoption: {
          adoptedBy: '',
          date: '',
          newAdoptedName: '',
          fee: defaults.defaultAdoptionFee || null,
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
        litterName: null,
        profileSettings: {
          isSpotlightFeatured: false,
          showMedicalHistory: defaults.defaultShowMedical || false,
          showAdditionalInformation: false,
        },
      }
    }
  },
  { immediate: true },
)

function handleSave() {

  const defaults = settingsStore.settings.pets || {}

  if (defaults.requirePhotoForPublic && formData.value.details?.status === 'available') {
    if (!formData.value.photos || formData.value.photos.length === 0) {
      alert('Photo required for "Available" pets due to shelter policy.')
      return
    }
  }

  emit('save', formData.value)
}

function handleClose() {
  emit('close')
}

const tabs = [
  { id: 'basic', label: 'Basic Info', icon: 'fileText' },
  { id: 'physical', label: 'Physical', icon: 'paw' },
  { id: 'behavior', label: 'Behavior', icon: 'smile' },
  { id: 'medical', label: 'Medical', icon: 'activity' },
  { id: 'status', label: 'Status', icon: 'home' },
  { id: 'descriptions', label: 'Story', icon: 'book' },
  { id: 'photos', label: 'Photos', icon: 'camera' },
  { id: 'settings', label: 'Settings', icon: 'cog' },
]
</script>

<template>
  <EditorDrawer
    :is-open="isOpen"
    :title="pet ? `Edit ${pet.name}` : 'Add New Pet'"
    save-label="Save Pet"
    :is-saving="false"
    width="900px"
    @close="handleClose"
    @save="handleSave"
  >
    <template #sidebar>
      <SidebarNav :items="tabs" v-model="activeTab" variant="editor" />
    </template>
    <template #content>
      <div v-if="formData">
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
          <PetEditorSettings
            v-if="activeTab === 'settings'"
            v-model="formData"
            @archive="emit('archive', pet)"
          />
        </form>
      </div>
    </template>
  </EditorDrawer>
</template>

<style scoped>
/* Scoped styles removed as they are now handled by EditorDrawer */
</style>
