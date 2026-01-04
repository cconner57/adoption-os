<script setup lang="ts">
import { reactive, ref, computed, onMounted } from 'vue'
import ApplicationHeader from '../components/volunteer/application-header/ApplicationHeader.vue'
import Button from '../components/common/ui/Button.vue'
import AdoptionSteps from '../components/pet-adoption/adoption-steps/AdoptionSteps.vue'
import GeneralSection from '../components/pet-adoption/cat-adoption/GeneralSection.vue'
import HomeSection from '../components/pet-adoption/cat-adoption/HomeSection.vue'
import NewCatSection from '../components/pet-adoption/cat-adoption/NewCatSection.vue'
import CurrentPetsSection from '../components/pet-adoption/cat-adoption/CurrentPetsSection.vue'
import PastPetsSection from '../components/pet-adoption/cat-adoption/PastPetsSection.vue'
import OtherSection from '../components/pet-adoption/cat-adoption/OtherSection.vue'
import SummarySection from '../components/pet-adoption/cat-adoption/SummarySection.vue'
import type { FormState } from '../models/adopt-form.ts'

const formState = reactive<FormState>({
  firstName: null,
  lastName: null,
  age: null,
  spouseFirstName: null,
  spouseLastName: null,
  roommatesNames: [''],
  childrenNamesAges: [{ name: '', age: '' }],
  currentPets: [
    { name: '', speciesBreedSize: '', age: '', source: '', spayedNeutered: '', likesDogs: '' },
  ],
  currentlyHavePets: null,
  email: null,
  address: null,
  addressLine2: null,
  city: null,
  state: null,
  zip: null,
  phoneNumber: null,
  cellPhoneNumber: null,
  adultMembersAgreed: 'No',
  homeType: null,
  homeOwnership: null,
  landlordName: null,
  landlordPhoneNumber: null,
  allowPets: null,
  breedRestrictionsWeightLimit: null,
  monthlyFee: null,
  allergies: null,
  primaryOwner: null,
  yearsAtAddress: null,
  previousAddress: null,
  expectToMove: null,
  travelPlan: null,
  catAccess: null,
  catIndoorOutdoor: null,
  catPreferenceBreed: null,
  catPreferencePhysical: null,
  catPreferencePersonality: null,
  catPreferenceNotWant: null,
  whyInterested: null,
  adoptionReason: null,
  ownCatBefore: null,
  ownKittenBefore: null,
  alreadyHaveVeterinarian: null,
  catAllowedHomeArea: null,
  catHomeAloneHours: null,
  catDisciplineType: null,
  catEscapeSteps: null,
})

const formStep = ref(0)
const hasAttemptedSubmit = ref(false)
const touched = reactive<Record<string, boolean>>({})

const handleBlur = (field: string) => {
  touched[field] = true
}

const selectedCat = ref()

onMounted(() => {
  const storedPet = sessionStorage.getItem('adoption_pet_id')
  if (storedPet) {
    const pet = JSON.parse(storedPet)
    selectedCat.value = pet.petName
  }
})

const validationErrors = computed(() => {
  const errors: string[] = []

  if (formStep.value === 0) {
    const requiredFields: { key: keyof FormState; label: string }[] = [
      { key: 'firstName', label: 'First Name' },
      { key: 'lastName', label: 'Last Name' },
      { key: 'age', label: 'Age' },
      { key: 'email', label: 'Email' },
      { key: 'address', label: 'Address' },
      { key: 'addressLine2', label: 'Address Line 2' },
      { key: 'city', label: 'City' },
      { key: 'state', label: 'State' },
      { key: 'zip', label: 'Zip Code' },
      { key: 'phoneNumber', label: 'Phone Number' },
      { key: 'adultMembersAgreed', label: 'Household Agreement' },
    ]

    requiredFields.forEach(({ key, label }) => {
      const value = formState[key]
      if (key === 'adultMembersAgreed') {
        if (value === null) errors.push(label)
      } else if (value === null || value === '') {
        errors.push(label)
      }
    })
  }

  return errors
})

// const isFormValid = computed(() => validationErrors.value.length === 0)

const handleSubmit = () => {
  hasAttemptedSubmit.value = true

  // if (!isFormValid.value) {
  //   return
  // }

  if (formStep.value < 6) {
    formStep.value++
    hasAttemptedSubmit.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
    return
  }

  console.log('Form submitted with state:', formState)
  // Proceed to API call
}
</script>

<template>
  <section class="page-shell">
    <section class="form-card" aria-labelledby="form-title">
      <ApplicationHeader
        header-title="Pet"
        header-text="This application is intended as a means to match the right cat with the right home. The more detail you provide, the better.  All of our adoptable pets are spayed/neutered, vaccinated, and microchipped. Typical adoption fees are $300 for kittens and $250 for adults. Adoption fees are tax-deductible donations, not purchase prices. Thank you for considering adoption!"
      />
      <AdoptionSteps :formStep="formStep" selectedAnimal="cat" />
      <div class="cat-name-display">
        <h2>Adopting Pet:</h2>
        <p>{{ selectedCat }}</p>
      </div>
      <GeneralSection
        v-show="formStep === 0"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <HomeSection
        v-show="formStep === 1"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <NewCatSection
        v-show="formStep === 2"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <CurrentPetsSection
        v-show="formStep === 3"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <PastPetsSection
        v-show="formStep === 4"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <OtherSection
        v-show="formStep === 5"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <SummarySection
        v-show="formStep === 6"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />

      <div v-if="hasAttemptedSubmit && validationErrors.length > 0" class="validation-summary">
        <p class="summary-title">Please complete the following required fields:</p>
        <div class="tags">
          <span v-for="err in validationErrors" :key="err" class="tag is-danger">{{ err }}</span>
        </div>
      </div>

      <div class="actions">
        <Button
          @click="handleSubmit"
          type="submit"
          :title="formStep < 6 ? 'Next' : 'Submit Application'"
          color="green"
          size="large"
        />
      </div>
    </section>
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 9rem var(--layout-padding-side) 64px;
  .form-card {
    max-width: 1600px;
    margin: 0 auto;
    background: var(--white);
    color: var(--font-color-dark);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;
    fieldset {
      border: 0;
      margin: 24px 0;
      padding: 0;
      .section-title {
        font-weight: 700;
        font-size: 18px;
        margin: 18px 0 12px;
      }
    }
    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;
    }
    .grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 16px;
    }
  }
  @media (min-width: 321px) and (max-width: 430px) {
    padding: 6rem 16px 32px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
}
.cat-name-display {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 2rem;
}
.cat-name-display h2 {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--font-color-dark);
}
.cat-name-display p {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--green);
}
.validation-summary {
  background-color: #fef2f2;
  border: 1px solid #ef4444;
  border-radius: 12px;
  padding: 16px;
  margin: 24px 0;
  text-align: center;

  .summary-title {
    color: #b91c1c;
    font-weight: 600;
    margin-bottom: 12px;
  }

  .tags {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    justify-content: center;
  }

  .tag.is-danger {
    background-color: #fee2e2;
    color: #b91c1c;
    padding: 4px 12px;
    border-radius: 16px;
    font-size: 0.875rem;
    font-weight: 500;
  }
}
</style>
