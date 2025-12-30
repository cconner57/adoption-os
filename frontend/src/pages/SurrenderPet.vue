<script setup lang="ts">
import type { SurrenderFormState } from '../models/surrender-form.ts'
import { reactive, ref, computed } from 'vue'
import {
  AggressiveSection,
  FeedingSection,
  HouseholdSection,
  MedicalSection,
  BehaviorSection,
  OtherSection,
} from '../components/about/surrender/index.ts'
import Button from '../components/common/ui/Button.vue'
import SurrenderSteps from '../components/about/surrender/SurrenderSteps.vue'
import PetSelectSection from '../components/about/surrender/PetSelectSection.vue'
import FormSubmitted from '../components/common/form-submitted/FormSubmitted.vue'

const formState = reactive<SurrenderFormState>({
  // Cat & Household Information
  firstName: '',
  lastName: '',
  phoneNumber: '',
  email: '',
  streetAddress: '',
  city: '',
  state: '',
  zipCode: '',
  whenToSurrenderCat: '',
  catName: '',
  catSex: '',
  catAge: '',
  catOwnershipDuration: '',
  catLocationFound: '',
  catWhySurrendered: '',
  householdMembers: [{ age: '', gender: 'Female', count: 1 }],
  otherPetsInHousehold: '',

  // Behavior
  catsBehaviorTowardsKnownPeople: '',
  catsBehaviorTowardsStrangers: '',
  catsBehaviorTowardsKnownAnimals: '',
  commentsOnBehavior: '',
  catsReactionToNewPeople: '',
  catHouseTrained: '',
  catSpendMajorityOfTime: '',
  catLeftAloneDuration: '',
  catWhenLeftAlone: '',
  catLeftAloneBehaviors: '',
  catHowItPlays: '',
  catToysItLikes: '',
  catGamesItLikes: '',
  catScaredOfAnything: '',
  catScaredOfAnythingExplanation: '',
  catBadHabits: '',
  catAllowedOnFurniture: '',
  catSleepAtNight: '',
  catBehaviorFoodOthers: '',
  catBehaviorToysOthers: '',
  catProblemsRidingInCar: '',
  catProblemsRidingInCarExplanation: '',
  catEscapedBefore: '',
  catEscapedBeforeExplanation: '',

  // Aggressive Behavior
  catEverAttackedPeople: '',
  catEverAttackedPeopleExplanation: '',
  catEverAttackedOtherCats: '',
  catEverAttackedOtherCatsExplanation: '',

  // Medical History
  catVeterinarianList: '',
  catVeterinarianYearlyVisits: '',
  catSpayedNeutered: '',
  catVaccinationHistory: '',
  catVaccinationsCurrent: '',
  catTestedHeartworm: '',
  catTestedHeartwormExplanation: '',
  catHeartwormPrevention: '',
  catHeartwormPreventionExplanation: '',
  catMicrochipped: '',
  catMicrochippedExplanation: '',
  catVetOrGroomerBehavior: '',
  catVetMuzzled: '',
  catPastOrPresentHealthProblems: '',
  catPastOrPresentHealthProblemsExplanation: '',
  catCurrentMedications: '',
  catCurrentMedicationsExplanation: '',

  // Feeding
  catTypeOfFood: '',
  catEatingFrequency: '',
  catAmountOfFood: '',
  catFoodTreats: '',
  catFoodTreatsExplanation: '',

  // Other
  additionalInformation: '',
  fullBodyPhotoOfCat: '',
  closeUpPhotoOfCatsFace: '',
  copiesOfRecords: '',
})

const formStep = ref(0)
const selectedAnimal = ref<'dog' | 'cat' | null>(null)
const formError = ref<boolean>(false)
const isSubmitted = ref(false)

const touched = reactive<Record<string, boolean>>({})
const hasAttemptedSubmit = ref(false)

const handleBlur = (field: string) => {
  touched[field] = true
}

const validationErrors = computed(() => {
  const errors: string[] = []

  if (formStep.value === 1) {
    const {
      firstName, lastName, phoneNumber, email, streetAddress, city, state, zipCode,
      catName, catAge,
      whenToSurrenderCat, catSex, catOwnershipDuration, catLocationFound, catWhySurrendered, otherPetsInHousehold
    } = formState

    if (!firstName) errors.push('First Name')
    if (!lastName) errors.push('Last Name')
    if (!phoneNumber) errors.push('Phone Number')
    if (!email) errors.push('Email')
    if (!streetAddress) errors.push('Street Address')
    if (!city) errors.push('City')
    if (!state) errors.push('State')
    if (!zipCode) errors.push('Zip Code')
    if (!whenToSurrenderCat) errors.push('When do you need to surrender your cat')
    if (!catName) errors.push("Cat's Name")
    if (!catAge) errors.push("Age")
    if (!catSex) errors.push("Sex")
    if (!catOwnershipDuration) errors.push('How long have you had your cat?')
    if (!catLocationFound) errors.push('Where did you get your cat?')
    if (!catWhySurrendered) errors.push('Why are you surrendering your cat?')

    // Check dynamic household members
    let hasAgeError = false
    let hasQtyError = false
    formState.householdMembers.forEach((member) => {
       if (!member.age) hasAgeError = true
       if (!member.count || member.count < 1) hasQtyError = true
    })

    if (hasAgeError) errors.push('Household - Age')
    if (hasQtyError) errors.push('Household - Quantity')

    if (!otherPetsInHousehold) errors.push('What other animals did the cat live with?')
  }

  return errors
})

const isStepValid = computed(() => {
  if (formStep.value === 0) return !!selectedAnimal.value
  if (formStep.value === 1) return validationErrors.value.length === 0
  return true
})

const validateStep = (step: number): boolean => {
  if (step === 0) {
     if (!selectedAnimal.value) {
      formError.value = true
      return false
    }
    formError.value = false
    return true
  }

  if (!isStepValid.value) {
    // Mark regular fields as touched
    const fields = [
      'firstName', 'lastName', 'phoneNumber', 'email', 'streetAddress', 'city', 'state', 'zipCode',
      'catName', 'catAge', 'whenToSurrenderCat', 'catSex', 'catOwnershipDuration', 'catLocationFound', 'catWhySurrendered', 'otherPetsInHousehold'
    ]
    fields.forEach(f => touched[f] = true)

    // Mark household members as touched on submit
    formState.householdMembers.forEach((_, index) => {
      // Use consistent key format with HouseholdSection.vue
      touched[`householdMembers[${index}].age`] = true
      touched[`householdMembers[${index}].count`] = true
    })

    return false
  }

  return true
}

const handleSubmit = () => {
  hasAttemptedSubmit.value = true
  if (!validateStep(formStep.value)) {
    window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' })
    return
  }

  if (formStep.value < 6) {
    formStep.value += 1
    hasAttemptedSubmit.value = false // Reset for next step
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } else {
    // Final submission logic
    isSubmitted.value = true
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const handleReset = () => {
  isSubmitted.value = false
  window.location.reload()
}

const headerText = computed(() => {
  if (!selectedAnimal.value || formStep.value === 0) {
    return 'Surrender Pet'
  }
  return selectedAnimal.value === 'cat' ? 'Feline Surrender' : 'Canine Surrender'
})
</script>

<template>
  <section class="page-shell">
    <section v-if="!isSubmitted" class="form-card" aria-labelledby="form-title">
      <div class="form-header">
        <img
          v-if="selectedAnimal === 'cat' && formStep > 0"
          src="../../public/images/cat.png"
          alt="cat"
          height="50"
          width="100"
        />
        <img
          v-if="selectedAnimal === 'dog' && formStep > 0"
          src="../../public/images/dog.png"
          alt="cat"
          height="50"
          width="100"
        />
        <h1>{{ headerText }}</h1>
      </div>
      <SurrenderSteps
        v-if="selectedAnimal && formStep > 0"
        :formStep="formStep"
        :selectedAnimal="selectedAnimal"
      />
      <content>
        <PetSelectSection
          v-if="formStep === 0"
          :formError="formError"
          :selectedAnimal="selectedAnimal"
          @update:selectedAnimal="
            (value) => {
              selectedAnimal = value
            }
          "
        />
        <HouseholdSection
          v-if="formStep === 1"
          :formState="formState"
          :touched="touched"
          :handleBlur="handleBlur"
          :hasAttemptedSubmit="hasAttemptedSubmit"
        />
        <BehaviorSection v-if="formStep === 2" :formState="formState" />
        <AggressiveSection v-if="formStep === 3" :formState="formState" />
        <MedicalSection v-if="formStep === 4" :formState="formState" />
        <FeedingSection v-if="formStep === 5" :formState="formState" />
        <OtherSection v-if="formStep === 6" :formState="formState" />
      </content>

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
          :title="formStep === 6 ? 'Submit' : 'Next'"
          color="green"
          size="large"
          :disabled="formStep === 0 && !selectedAnimal"
        />
      </div>
    </section>

    <FormSubmitted v-else @reset="handleReset" formType="surrender" />
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
    .form-header {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 16px;
      margin-bottom: 4px;
      color: var(--green);
      & h1 {
        font-size: 4.25rem;
        line-height: 1.2;
        letter-spacing: 0.2px;
      }
      img {
        width: 100px;
      }
    }
    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;
    }
  }
  @media (max-width: 440px) {
    .form-header {
      flex-direction: column;
      align-items: center;
      gap: 0px;
      margin-bottom: 1rem;
      & h1 {
        font-size: 1.75rem;
        text-align: center;
      }
    }
  }
}

.validation-summary {
  background-color: #fef2f2;
  border: 1px solid #ef4444;
  border-radius: 12px;
  padding: 16px;
  margin: 24px 0;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;

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
