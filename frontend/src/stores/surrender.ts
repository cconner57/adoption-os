import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { SurrenderFormState } from '../models/surrender-form'
import { API_ENDPOINTS } from '../constants/api'

export const useSurrenderStore = defineStore('surrender', () => {
  const step = ref(0)

  const isSubmitted = ref(false)
  const hasAttemptedSubmit = ref(false)
  const selectedAnimal = ref<'dog' | 'cat' | null>(null)

  const formState = reactive<SurrenderFormState>({
    firstName: '',
    lastName: '',
    phoneNumber: '',
    email: '',
    streetAddress: '',
    city: '',
    state: '',
    zipCode: '',
    whenToSurrenderAnimal: '',
    animalName: '',
    animalSex: '',
    animalAge: '',
    animalOwnershipDuration: '',
    animalLocationFound: '',
    animalWhySurrendered: '',
    householdMembers: [{ age: '', gender: 'Female', count: 1 }],
    otherPetsInHousehold: '',

    animalsBehaviorTowardsKnownPeople: '',

    animalsBehaviorTowardsStrangers: '',
    animalsBehaviorTowardsKnownAnimals: '',
    commentsOnBehavior: '',
    animalsReactionToNewPeople: '',
    animalHouseTrained: '',
    animalSpendMajorityOfTime: '',
    animalLeftAloneDuration: '',
    animalWhenLeftAlone: '',
    animalLeftAloneBehaviors: '',
    animalHowItPlays: '',
    animalToysItLikes: '',
    animalGamesItLikes: '',
    animalScaredOfAnything: '',
    animalScaredOfAnythingExplanation: '',
    animalBadHabits: '',
    animalAllowedOnFurniture: '',
    animalSleepAtNight: '',
    animalBehaviorFoodOthers: '',
    animalBehaviorToysOthers: '',
    animalProblemsRidingInCar: '',
    animalProblemsRidingInCarExplanation: '',
    animalEscapedBefore: '',
    animalEscapedBeforeExplanation: '',

    animalEverAttackedPeople: '',

    animalEverAttackedPeopleExplanation: '',
    animalEverAttackedOtherCats: '',
    animalEverAttackedOtherCatsExplanation: '',
    animalEverAttackedOtherDogs: '',
    animalEverAttackedOtherDogsExplanation: '',

    animalVeterinarianList: '',

    animalVeterinarianYearlyVisits: '',
    animalSpayedNeutered: '',
    animalVaccinationHistory: '',
    animalVaccinationsCurrent: '',
    animalTestedHeartworm: '',
    animalTestedHeartwormExplanation: '',
    animalHeartwormPrevention: '',
    animalHeartwormPreventionExplanation: '',
    animalMicrochipped: '',
    animalMicrochippedExplanation: '',
    animalVetOrGroomerBehavior: '',
    animalVetMuzzled: '',
    animalPastOrPresentHealthProblems: '',
    animalPastOrPresentHealthProblemsExplanation: '',
    animalCurrentMedications: '',
    animalCurrentMedicationsExplanation: '',

    animalTypeOfFood: '',

    animalEatingFrequency: '',
    animalAmountOfFood: '',
    animalFoodTreats: '',
    animalFoodTreatsExplanation: '',

    additionalInformation: '',

    fullBodyPhotoOfAnimal: '',
    closeUpPhotoOfAnimalFace: '',
    copiesOfRecords: '',
  })

  const validationErrors = computed(() => {
    const errors: string[] = []

    if (step.value === 1) {
      if (!formState.firstName) errors.push('First Name')
      if (!formState.lastName) errors.push('Last Name')
      if (!formState.phoneNumber) errors.push('Phone Number')
      if (!formState.email) errors.push('Email')
      if (!formState.streetAddress) errors.push('Street Address')
      if (!formState.city) errors.push('City')
      if (!formState.state) errors.push('State')
      if (!formState.zipCode) errors.push('Zip Code')
      if (!formState.whenToSurrenderAnimal) errors.push('When do you need to surrender your animal')
      if (!formState.animalName) errors.push("Animal's Name")
      if (!formState.animalAge) errors.push('Age')
      if (!formState.animalSex) errors.push('Sex')
      if (!formState.animalOwnershipDuration) errors.push('How long have you had your animal?')
      if (!formState.animalLocationFound) errors.push('Where did you get your animal?')
      if (!formState.animalWhySurrendered) errors.push('Why are you surrendering your animal?')
      if (!formState.otherPetsInHousehold) errors.push('Other pets in household')

      let hasAgeError = false
      let hasQtyError = false
      formState.householdMembers.forEach((member) => {
        if (!member.age) hasAgeError = true
        if (!member.count || member.count < 1) hasQtyError = true
      })
      if (hasAgeError) errors.push('Household - Age')
      if (hasQtyError) errors.push('Household - Quantity')
    }

    return errors
  })

  const isStepValid = computed(() => {
    if (step.value === 0) return !!selectedAnimal.value
    if (step.value === 1) return validationErrors.value.length === 0
    return true
  })

  const nextStep = () => {
    hasAttemptedSubmit.value = true
    if (!isStepValid.value) return false

    if (step.value === 3 && selectedAnimal.value === 'dog') {
      step.value += 2
    } else {
      step.value++
    }

    hasAttemptedSubmit.value = false
    return true
  }

  const prevStep = () => {
    if (step.value > 0) {
      if (step.value === 5 && selectedAnimal.value === 'dog') {
        step.value -= 2
      } else {
        step.value--
      }
    }
  }

  const resetForm = () => {
    step.value = 0
    isSubmitted.value = false
    hasAttemptedSubmit.value = false
    selectedAnimal.value = null
  }

  const submitApplication = async () => {
    try {
      isSubmitted.value = false
      // Import API_ENDPOINTS if not already available in scope (adding import at top of file)
      const response = await fetch(API_ENDPOINTS.SURRENDER_APPLICATION, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formState),
      })

      if (!response.ok) {
        throw new Error('Failed to submit application')
      }

      isSubmitted.value = true
      resetForm()
    } catch (error) {
      console.error('Error submitting form:', error)
      alert('There was an error submitting your application. Please try again.')
    }
  }

  return {
    formState,
    step,
    isSubmitted,
    hasAttemptedSubmit,
    selectedAnimal,
    validationErrors,
    isStepValid,
    nextStep,
    prevStep,
    submitApplication,
    resetForm,
  }
})
