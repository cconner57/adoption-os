<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { IPet, TSpecies, TSex, TPetStatus, TAgeGroup } from '../../../models/common'

const props = defineProps<{
  pet: IPet | null
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'save'])

const activeTab = ref('basic')
const formData = ref<Partial<IPet>>({})

// Initialize form data when pet changes or modal opens
watch(
  () => props.pet,
  (newPet) => {
    if (newPet) {
      formData.value = JSON.parse(JSON.stringify(newPet))
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
      breed: 'Unknown',
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
      showMedicalHistory: true,
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
</script>

<template>
  <div class="pet-editor-overlay" :class="{ open: isOpen }" @click.self="handleClose">
    <div class="pet-editor-drawer">
      <div class="drawer-header">
        <h2>{{ pet ? 'Edit Pet' : 'Add New Pet' }}</h2>
        <button class="close-btn" @click="handleClose">✕</button>
      </div>

      <div class="drawer-body">
        <div class="tabs-nav">
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
            <div class="form-group">
              <label>Name</label>
              <input v-model="formData.name" type="text" placeholder="Pet's Name" />
            </div>

            <div class="form-row">
              <div class="form-group">
                <label>Species</label>
                <select v-model="formData.species">
                  <option value="cat">Cat</option>
                  <option value="dog">Dog</option>
                </select>
              </div>
              <div class="form-group">
                <label>Sex</label>
                <select v-model="formData.sex">
                  <option value="male">Male</option>
                  <option value="female">Female</option>
                  <option value="unknown">Unknown</option>
                </select>
              </div>
            </div>

            <div class="form-group">
              <label>Status</label>
              <select v-if="formData.details" v-model="formData.details.status">
                <option value="available">Available</option>
                <option value="adoption-pending">Adoption Pending</option>
                <option value="adopted">Adopted</option>
                <option value="foster">In Foster</option>
                <option value="hold">Medical/Behavioral Hold</option>
                <option value="intake">Intake Processing</option>
                <option value="archived">Archived</option>
              </select>
            </div>
          </div>

          <!-- Physical Tab -->
          <div v-if="activeTab === 'physical'" class="form-section">
            <div class="form-group" v-if="formData.physical">
              <label>Breed</label>
              <input
                v-model="formData.physical.breed"
                type="text"
                placeholder="e.g. Domestic Short Hair"
              />
            </div>
            <div class="form-row" v-if="formData.physical">
              <div class="form-group">
                <label>Color</label>
                <input
                  v-model="formData.physical.color"
                  type="text"
                  placeholder="e.g. Black & White"
                />
              </div>
              <div class="form-group">
                <label>Date of Birth</label>
                <input v-model="formData.physical.dateOfBirth" type="date" />
              </div>
            </div>
            <div class="form-row" v-if="formData.physical">
              <div class="form-group">
                <label>Age Group</label>
                <select v-model="formData.physical.ageGroup">
                  <option value="baby">Baby (Kitten/Puppy)</option>
                  <option value="young">Young</option>
                  <option value="adult">Adult</option>
                  <option value="senior">Senior</option>
                </select>
              </div>
              <div class="form-group">
                <label>Size</label>
                <select v-model="formData.physical.size">
                  <option value="extra-small">Extra Small</option>
                  <option value="small">Small</option>
                  <option value="medium">Medium</option>
                  <option value="large">Large</option>
                  <option value="extra-large">Extra Large</option>
                </select>
              </div>
            </div>
            <div class="form-group" v-if="formData.physical">
              <label>Coat Length</label>
              <select v-model="formData.physical.coatLength">
                <option value="short">Short</option>
                <option value="medium">Medium</option>
                <option value="long">Long</option>
                <option value="wire">Wire</option>
                <option value="hairless">Hairless</option>
              </select>
            </div>
          </div>

          <!-- Behavior Tab -->
          <div v-if="activeTab === 'behavior'" class="form-section">
            <div class="form-group" v-if="formData.behavior">
              <label>Energy Level</label>
              <div class="radio-group">
                <label
                  ><input type="radio" v-model="formData.behavior.energyLevel" value="low" />
                  Low</label
                >
                <label
                  ><input type="radio" v-model="formData.behavior.energyLevel" value="medium" />
                  Medium</label
                >
                <label
                  ><input type="radio" v-model="formData.behavior.energyLevel" value="high" />
                  High</label
                >
              </div>
            </div>

            <div class="form-group checkbox-list" v-if="formData.behavior">
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.behavior.isGoodWithKids" /> Good with Kids
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.behavior.isGoodWithDogs" /> Good with Dogs
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.behavior.isGoodWithCats" /> Good with Cats
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.behavior.isHouseTrained" /> House Trained
              </label>
            </div>
          </div>

          <!-- Medical Tab -->
          <div v-if="activeTab === 'medical'" class="form-section">
            <div class="form-group checkbox-list" v-if="formData.medical">
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.medical.spayedOrNeutered" /> Spayed /
                Neutered
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.medical.vaccinationsUpToDate" />
                Vaccinations Up to Date
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.medical.microchip.microchipped" />
                Microchipped
              </label>

              <div class="form-group" v-if="formData.medical.microchip.microchipped">
                <label>Microchip ID</label>
                <input
                  v-model="formData.medical.microchip.microchipID"
                  type="text"
                  placeholder="e.g. 98102000..."
                />
              </div>
            </div>
          </div>

          <!-- Descriptions Tab -->
          <div v-if="activeTab === 'descriptions'" class="form-section">
            <div class="form-group" v-if="formData.descriptions">
              <label>Primary Bio</label>
              <textarea
                v-model="formData.descriptions.primary"
                rows="6"
                placeholder="Tell their story..."
              ></textarea>
            </div>
            <div class="form-group" v-if="formData.descriptions">
              <label>Spotlight Blurb (Short)</label>
              <textarea
                v-model="formData.descriptions.spotlight"
                rows="3"
                placeholder="Catchy one-liner..."
              ></textarea>
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
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.profileSettings.isSpotlightFeatured" />
                Feature in Spotlight
              </label>
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.profileSettings.showMedicalHistory" /> Show
                Medical History Publicly
              </label>
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
  width: 500px;
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
  overflow-y: auto;
  display: flex;
  flex-direction: column;
}

.tabs-nav {
  display: flex;
  overflow-x: auto;
  border-bottom: 1px solid var(--border-color);
  padding: 0 12px;
  background: hsl(from var(--color-neutral) h s 98%);
}

.tab-btn {
  padding: 12px 16px;
  background: none;
  border: none;
  border-bottom: 3px solid transparent;
  cursor: pointer;
  font-weight: 500;
  color: hsl(from var(--color-neutral) h s 50%);
  white-space: nowrap;
  display: flex;
  align-items: center;
  gap: 6px;

  &.active {
    color: var(--color-secondary);
    border-bottom-color: var(--color-secondary);
  }

  &:hover {
    background: hsl(from var(--color-primary) h s 98%);
  }
}

.tab-content {
  padding: 24px;
}

.form-section {
  display: flex;
  flex-direction: column;
  gap: 20px;
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

  input[type='text'],
  select,
  textarea {
    padding: 10px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-size: 1rem;
    outline: none;
    background-color: var(--text-inverse);
    color: var(--text-primary);
    font-family: var(--font-regular);

    &:focus {
      border-color: var(--color-secondary);
      box-shadow: 0 0 0 2px hsla(from var(--color-secondary) h s l / 0.1);
    }
  }
}

.form-row {
  display: flex;
  gap: 16px;

  .form-group {
    flex: 1;
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
}

.radio-group {
  display: flex;
  gap: 20px;

  label {
    font-weight: 400;
    display: flex;
    align-items: center;
    gap: 6px;
    color: var(--text-primary);
  }
}

.photos-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
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
  top: 4px;
  right: 4px;
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
}

.primary-badge {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(46, 204, 113, 0.8); /* Keep nice green overlay */
  color: white;
  text-align: center;
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

  &:hover {
    border-color: var(--color-secondary);
    color: var(--color-secondary);
    background: hsl(from var(--color-secondary) h s 98%);
  }
}

.drawer-footer {
  padding: 16px 24px;
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  background: var(--text-inverse);
}

.btn-cancel {
  padding: 10px 20px;
  border: 1px solid var(--border-color);
  background: var(--text-inverse);
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}

.btn-save {
  padding: 10px 20px;
  background: var(--color-primary);
  color: var(--text-inverse);
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;

  &:hover {
    background: hsl(from var(--color-primary) h s 40%);
  }
}
</style>
