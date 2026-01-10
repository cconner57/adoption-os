<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { IPet } from '../../../models/common'
import {
  InputField,
  Select,
  InputTextArea,
  InputSelectGroup,
  ButtonToggle,
  Combobox,
} from '../../common/ui'

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
      }

      // Ensure Medical object structure
      if (!formData.value.medical) {
        formData.value.medical = {
          spayedOrNeutered: false,
          vaccinationsUpToDate: false,
          vaccinations: {},
          surgeries: [],
          microchip: { microchipped: false },
        }
      } else {
        // Deep ensure microchip
        if (!formData.value.medical.microchip) {
          formData.value.medical.microchip = { microchipped: false }
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
    } else {
      resetForm()
    }
  },
  { immediate: true },
)

function resetForm() {
  formData.value = {
    name: '',
    species: 'cat',
    sex: 'unknown',
    details: {
      status: 'available',
    },
    physical: {
      breed: '',
      color: '',
      dateOfBirth: '',
      ageGroup: 'baby',
      size: 'medium',
      coatLength: 'short',
    },
    behavior: {
      isGoodWithCats: null,
      isGoodWithDogs: null,
      isGoodWithKids: null,
      isHouseTrained: null,
      energyLevel: 'medium',
      personalityTags: [],
    },
    medical: {
      spayedOrNeutered: false,
      vaccinationsUpToDate: false,
      vaccinations: {},
      surgeries: [],
      microchip: { microchipped: false },
    },
    descriptions: {
      primary: '',
      spotlight: '',
    },
    photos: [],
    profileSettings: {
      isSpotlightFeatured: false,
      showAdditionalInformation: true,
      showMedicalHistory: false,
    },
  }
}

function handleSave() {
  emit('save', formData.value)
}

function handleClose() {
  emit('close')
}

// Mock photo upload
function handlePhotoUpload() {
  const mockPhoto = {
    url: 'https://placekitten.com/300/300',
    thumbnailUrl: 'https://placekitten.com/100/100',
    isPrimary: formData.value.photos?.length === 0,
    uploadedAt: new Date().toISOString(),
  }
  if (!formData.value.photos) formData.value.photos = []
  formData.value.photos.push(mockPhoto)
}

function removePhoto(index: number) {
  formData.value.photos?.splice(index, 1)
}

// Computed Properties for UI mapping
const behaviorOptions = [
  { label: 'Good with Kids', value: 'kids' },
  { label: 'Good with Dogs', value: 'dogs' },
  { label: 'Good with Cats', value: 'cats' },
  { label: 'House Trained', value: 'trained' },
]

const behaviorFlags = computed({
  get: () => {
    const list: string[] = []
    if (formData.value.behavior?.isGoodWithKids) list.push('kids')
    if (formData.value.behavior?.isGoodWithDogs) list.push('dogs')
    if (formData.value.behavior?.isGoodWithCats) list.push('cats')
    if (formData.value.behavior?.isHouseTrained) list.push('trained')
    return list
  },
  set: (vals: string[] | null) => {
    if (!formData.value.behavior) return
    const v = vals || []
    formData.value.behavior.isGoodWithKids = v.includes('kids')
    formData.value.behavior.isGoodWithDogs = v.includes('dogs')
    formData.value.behavior.isGoodWithCats = v.includes('cats')
    formData.value.behavior.isHouseTrained = v.includes('trained')
  },
})

// Tabs configuration
const tabs = [
  { id: 'basic', label: 'Basic Info', icon: '📝' },
  { id: 'physical', label: 'Physical', icon: '📏' },
  { id: 'behavior', label: 'Behavior', icon: '🧠' },
  { id: 'medical', label: 'Medical', icon: '⚕️' },
  { id: 'descriptions', label: 'Story', icon: '📖' },
  { id: 'photos', label: 'Photos', icon: '📸' },
  { id: 'settings', label: 'Settings', icon: '⚙️' },
]

// Dropdown Options
const speciesOptions = [
  { label: 'Cat', value: 'cat' },
  { label: 'Dog', value: 'dog' },
]

const sexOptions = [
  { label: 'Male', value: 'male' },
  { label: 'Female', value: 'female' },
  { label: 'Unknown', value: 'unknown' },
]

// Temperament Options
const temperamentOptions = [
  { label: 'Affectionate', value: 'affectionate' },
  { label: 'Anxious', value: 'anxious' },
  { label: 'Bossy', value: 'bossy' },
  { label: 'Curious', value: 'curious' },
  { label: 'Hunter', value: 'hunter' },
  { label: 'Independent', value: 'independent' },
  { label: 'Laid-back', value: 'laid-back' },
  { label: 'Playful', value: 'playful' },
  { label: 'Shy', value: 'shy' },
  { label: 'Vocal', value: 'vocal' },
]

// Placement Options
const placementOptions = [
  { label: 'Must go with another cat', value: 'must_cat' },
  { label: 'Must go with another dog', value: 'must_dog' },
  { label: 'Prefers to be alone', value: 'alone' },
]

const placementRequirements = computed({
  get: () => {
    const list: string[] = []
    if (formData.value.behavior?.mustGoWithAnotherCat) list.push('must_cat')
    if (formData.value.behavior?.mustGoWithAnotherDog) list.push('must_dog')
    if (formData.value.behavior?.prefersToBeAlone) list.push('alone')
    return list
  },
  set: (vals: string[] | null) => {
    if (!formData.value.behavior) return
    const v = vals || []
    formData.value.behavior.mustGoWithAnotherCat = v.includes('must_cat')
    formData.value.behavior.mustGoWithAnotherDog = v.includes('must_dog')
    formData.value.behavior.prefersToBeAlone = v.includes('alone')
  },
})

// Bonded Computed Helpers
const isBondedComputed = computed({
  get: () => formData.value.behavior?.bonded?.isBonded ?? false,
  set: (val: boolean) => {
    if (!formData.value.behavior) return
    if (!formData.value.behavior.bonded) {
      formData.value.behavior.bonded = { isBonded: false, bondedWith: [] }
    }
    formData.value.behavior.bonded.isBonded = val
  },
})

// Ensure bondedWith is initialized array
const bondedWithOptions = computed(() => {
  const all = props.availablePets || []
  // Filter out self
  return all.filter((p) => p.id !== props.pet?.id).map((p) => ({ label: p.name, value: p.name }))
})

const statusOptions = [
  { label: 'Available', value: 'available' },
  { label: 'Adoption Pending', value: 'adoption-pending' },
  { label: 'Adopted', value: 'adopted' },
  { label: 'In Foster', value: 'foster' },
  { label: 'Medical/Behavioral Hold', value: 'hold' },
  { label: 'Intake Processing', value: 'intake' },
  { label: 'Archived', value: 'archived' },
]

const ageGroupOptions = [
  { label: 'Baby (Kitten/Puppy)', value: 'baby' },
  { label: 'Young', value: 'young' },
  { label: 'Adult', value: 'adult' },
  { label: 'Senior', value: 'senior' },
]

const sizeOptions = [
  { label: 'Extra Small', value: 'extra-small' },
  { label: 'Small', value: 'small' },
  { label: 'Medium', value: 'medium' },
  { label: 'Large', value: 'large' },
  { label: 'Extra Large', value: 'extra-large' },
]

const coatLengthOptions = [
  { label: 'Short', value: 'short' },
  { label: 'Medium', value: 'medium' },
  { label: 'Long', value: 'long' },
  { label: 'Wire', value: 'wire' },
  { label: 'Hairless', value: 'hairless' },
]

// Breeds
const catBreeds = [
  'Domestic Short Hair',
  'Domestic Medium Hair',
  'Domestic Long Hair',
  'Siamese',
  'Maine Coon',
  'Persian',
  'Ragdoll',
  'Bengal',
  'Sphynx',
  'Other',
]

const dogBreeds = [
  'Mixed Breed',
  'Labrador Retriever',
  'German Shepherd',
  'Golden Retriever',
  'Bulldog',
  'Poodle',
  'Beagle',
  'Rottweiler',
  'Chihuahua',
  'Pit Bull Terrier',
  'Other',
]

const breedOptions = computed(() => {
  const species = formData.value.species?.toLowerCase() || 'cat'
  const list = species === 'dog' ? dogBreeds : catBreeds
  return list.map((b) => ({ label: b, value: b }))
})

const energyLevelOptions = [
  { label: 'Low', value: 'low' },
  { label: 'Medium', value: 'medium' },
  { label: 'High', value: 'high' },
]
</script>

<template>
  <div class="pet-editor-overlay" :class="{ open: isOpen }" @click.self="handleClose">
    <div class="pet-editor-drawer">
      <div class="drawer-header">
        <h2>{{ pet ? 'Edit Pet' : 'Add New Pet' }}</h2>
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
          <!-- Basic Info Tab -->
          <div v-if="activeTab === 'basic'" class="form-section">
            <InputField label="Name" v-model="formData.name" placeholder="Pet's Name" />

            <div class="form-row">
              <Select
                label="Species"
                v-model="formData.species"
                :options="speciesOptions"
                fullWidth
              />
              <Select label="Sex" v-model="formData.sex" :options="sexOptions" fullWidth />
            </div>

            <div class="form-row">
              <Combobox
                label="Litter Name"
                v-model="formData.litterName"
                :options="litterOptions"
                placeholder="Select or create litter..."
                allow-create
              />
            </div>

            <Select
              v-if="formData.details"
              label="Status"
              v-model="formData.details.status"
              :options="statusOptions"
              fullWidth
            />
          </div>

          <!-- Physical Tab -->
          <div v-if="activeTab === 'physical'" class="form-section">
            <div v-if="formData.physical">
              <!-- Replaced InputField with Select for Breed -->
              <Select
                label="Breed"
                v-model="formData.physical.breed"
                :options="breedOptions"
                fullWidth
              />

              <br />

              <div class="form-row">
                <InputField
                  label="Color"
                  v-model="formData.physical.color"
                  placeholder="e.g. Black & White"
                />
                <div class="form-group dob-group">
                  <InputField
                    label="Date of Birth"
                    v-model="formData.physical.dateOfBirth"
                    type="date"
                    placeholder="mm/dd/yyyy"
                  />
                </div>
              </div>

              <br />

              <div class="form-row">
                <Select
                  label="Age Group"
                  v-model="formData.physical.ageGroup"
                  :options="ageGroupOptions"
                  fullWidth
                />
                <Select
                  label="Coat Length"
                  v-model="formData.physical.coatLength"
                  :options="coatLengthOptions"
                  fullWidth
                />
              </div>

              <br />

              <div class="form-row">
                <Select
                  label="Size"
                  v-model="formData.physical.size"
                  :options="sizeOptions"
                  fullWidth
                />
                <InputField
                  label="Current Weight (lbs)"
                  v-model="formData.physical.currentWeight"
                  type="number"
                  placeholder="e.g. 12"
                />
              </div>
            </div>
          </div>

          <!-- Behavior Tab -->
          <div v-if="activeTab === 'behavior'" class="form-section">
            <div class="form-group" v-if="formData.behavior">
              <InputSelectGroup
                label="Energy Level"
                v-model="formData.behavior.energyLevel"
                :options="energyLevelOptions"
              />
            </div>

            <!-- Checkbox Group for Behavior Flags -->
            <div class="form-group" v-if="formData.behavior">
              <InputSelectGroup
                label="Traits"
                v-model="behaviorFlags"
                :options="behaviorOptions"
                multiple
              />
            </div>

            <!-- Personality Tags -->
            <div class="form-group" v-if="formData.behavior">
              <InputSelectGroup
                label="Personality"
                v-model="formData.behavior.personalityTags"
                :options="temperamentOptions"
                multiple
              />
            </div>

            <!-- Placement Requirements -->
            <div class="form-group" v-if="formData.behavior">
              <InputSelectGroup
                label="Placement Requirements"
                v-model="placementRequirements"
                :options="placementOptions"
                multiple
              />
            </div>

            <!-- Bonded & Special Needs -->
            <div class="form-group checkbox-list" v-if="formData.behavior">
              <ButtonToggle
                label="Part of a Bonded Pair?"
                v-model="isBondedComputed"
                :true-value="true"
                :false-value="false"
              />

              <div v-if="isBondedComputed" style="margin-left: 12px; margin-top: 8px">
                <Combobox
                  label="Bonded With"
                  v-model="formData.behavior.bonded.bondedWith"
                  :options="bondedWithOptions"
                  multiple
                  placeholder="Search pets..."
                />
              </div>
            </div>

            <div class="form-group" v-if="formData.behavior">
              <InputTextArea
                label="Special Needs / Behavior Notes"
                v-model="formData.behavior.specialNeeds"
                placeholder="Describe any special behavioral needs..."
                :maxChars="300"
              />
            </div>
          </div>

          <!-- Medical Tab -->
          <div v-if="activeTab === 'medical'" class="form-section">
            <div class="form-group checkbox-list" v-if="formData.medical">
              <ButtonToggle
                label="Spayed / Neutered"
                v-model="formData.medical.spayedOrNeutered"
                :true-value="true"
                :false-value="false"
              />
              <ButtonToggle
                label="Vaccinations Up to Date"
                v-model="formData.medical.vaccinationsUpToDate"
                :true-value="true"
                :false-value="false"
              />
              <ButtonToggle
                label="Microchipped"
                v-model="formData.medical.microchip.microchipped"
                :true-value="true"
                :false-value="false"
              />

              <div style="margin-top: 16px" v-if="formData.medical.microchip.microchipped">
                <InputField
                  label="Microchip ID"
                  v-model="formData.medical.microchip.microchipID"
                  placeholder="e.g. 98102000..."
                />
              </div>
            </div>
          </div>

          <!-- Descriptions Tab -->
          <div v-if="activeTab === 'descriptions'" class="form-section">
            <div class="form-group" v-if="formData.descriptions">
              <InputTextArea
                label="Primary Bio"
                v-model="formData.descriptions.primary"
                placeholder="Tell their story..."
                :maxChars="500"
              />
            </div>
            <div class="form-group" v-if="formData.descriptions">
              <InputTextArea
                label="Spotlight Blurb (Short)"
                v-model="formData.descriptions.spotlight"
                placeholder="Catchy one-liner..."
                :maxChars="100"
              />
            </div>
          </div>

          <!-- Photos Tab -->
          <div v-if="activeTab === 'photos'" class="form-section">
            <div class="photos-grid">
              <div class="photo-card" v-for="(photo, index) in formData.photos" :key="index">
                <img :src="photo.url" alt="Pet photo" />
                <button class="remove-photo" @click="removePhoto(index)">✕</button>
                <span v-if="photo.isPrimary" class="primary-badge">Primary</span>
              </div>
              <div class="add-photo-btn" @click="handlePhotoUpload">
                <span>+ Add Photo</span>
              </div>
            </div>
          </div>

          <!-- Settings Tab -->
          <div v-if="activeTab === 'settings'" class="form-section">
            <div class="form-group checkbox-list" v-if="formData.profileSettings">
              <ButtonToggle
                label="Feature in Spotlight"
                v-model="formData.profileSettings.isSpotlightFeatured"
                :true-value="true"
                :false-value="false"
              />
              <ButtonToggle
                label="Show Medical History Publicly"
                v-model="formData.profileSettings.showMedicalHistory"
                :true-value="true"
                :false-value="false"
              />
            </div>
          </div>
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
  width: 800px; /* Widened for Sidebar Layout */
  background: var(--text-inverse);
  height: 100%;
  display: flex;
  flex-direction: column;
  box-shadow: -4px 0 12px rgba(0, 0, 0, 0.1);
  transform: translateX(100%);
  transition: transform 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
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
  flex-shrink: 0; /* Don't shrink header */

  h2 {
    margin: 0;
    font-size: 1.5rem;
    color: var(--text-primary);
  }
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

.form-section {
  display: flex;
  flex-direction: column;
  gap: 24px;
  max-width: 100%; /* Ensure content stays within pane */
}

/* Reusing form-row but allowing flexible inputs */
.form-row {
  display: flex;
  gap: 16px;
  align-items: flex-start; /* vertical align top so inputs with labels align */
  width: 100%;
}
.form-row > * {
  flex: 1;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;

  label {
    font-weight: 600;
    color: var(--text-primary);
    font-size: 0.9rem;
  }
}

/* Native inputs styling (for date/textarea fallback) */
.native-date-input,
.native-textarea {
  padding: 10px 12px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  font-size: 1rem;
  outline: none;
  background-color: #ffffff; /* Changed from var(--text-inverse) to white to match InputField */
  color: var(--text-primary);
  font-family: var(--font-regular);
  width: 100%;
  box-sizing: border-box;

  &:focus {
    border-color: var(--color-secondary);
    box-shadow: 0 0 0 2px hsla(from var(--color-secondary) h s l / 0.1);
  }
}

.checkbox-list {
  gap: 12px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  font-weight: 400 !important;
  font-size: 0.95rem;
  user-select: none;
}

.radio-group {
  display: flex;
  gap: 24px;

  label {
    font-weight: 400;
    display: flex;
    align-items: center;
    gap: 8px;
    color: var(--text-primary);
    cursor: pointer;
  }
}

.photos-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr)); /* Better responsiveness */
  gap: 16px;
}

.photo-card {
  position: relative;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid var(--border-color);

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.remove-photo {
  position: absolute;
  top: 6px;
  right: 6px;
  background: rgba(0, 0, 0, 0.6);
  color: white; /* overlay text always white */
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.2s;

  &:hover {
    background: rgba(220, 38, 38, 0.9);
  }
}

.primary-badge {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(46, 204, 113, 0.9); /* Keep nice green overlay */
  color: white;
  text-align: center;
  font-weight: 600;
  font-size: 0.75rem;
  padding: 4px;
}

.add-photo-btn {
  aspect-ratio: 1;
  border: 2px dashed var(--border-color);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: hsl(from var(--color-neutral) h s 50%);
  transition: all 0.2s;
  font-weight: 500;

  &:hover {
    border-color: var(--color-secondary);
    color: var(--color-secondary);
    background: hsl(from var(--color-secondary) h s 98%);
  }
}

.drawer-footer {
  padding: 20px 32px;
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  background: var(--text-inverse);
  flex-shrink: 0;
}

.btn-cancel {
  padding: 10px 24px;
  border: 1px solid var(--border-color);
  background: white;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
  }
}

.btn-save {
  padding: 10px 24px;
  background: var(--color-primary);
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 4px hsla(from var(--color-primary) h s l / 0.2);

  &:hover {
    background: hsl(from var(--color-primary) h s 40%);
    transform: translateY(-1px);
    box-shadow: 0 4px 6px hsla(from var(--color-primary) h s l / 0.25);
  }

  &:active {
    transform: translateY(0);
  }
}
</style>
