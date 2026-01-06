import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { FormState } from '../models/adopt-form'
import { API_ENDPOINTS } from '../constants/api'

export const useAdoptionStore = defineStore('adoption', () => {
  const step = ref(0)
  const isSubmitted = ref(false)
  const hasAttemptedSubmit = ref(false)

  const formState = reactive<FormState>({
    firstName: '',
    lastName: '',
    age: null,
    spouseFirstName: null,
    spouseLastName: null,
    roommatesNames: [''],
    childrenNamesAges: [{ name: '', age: '' }],

    currentPets: [
      { name: '', speciesBreedSize: '', age: '', source: '', spayedNeutered: '', likesDogs: '' },
    ],
    currentlyHavePets: null,
    pastPets: [
      {
        name: '',
        speciesBreedSize: '',
        age: '',
        source: '',
        spayedNeutered: '',
        passedAwayReason: '',
      },
    ],
    ownPetsBefore: null,
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
    bredAnimalDescription: null,
    ownedDeclawedOrDebarked: null,
    movedWithPet: null,
    ownedSpecialNeedsPet: null,
    mobilityDevice: null,
    surrenderConditions: [],
    surrenderPlan: null,
    foodTypeBrand: null,
    affordVetCare: null,
    affordEmergencyCost: null,
    agreementSignature1: null,
    agreementSignature2: null,
    agreementSignature3: null,
    signatureData: null,
  })

  const validationErrors = computed(() => {
    const errors: string[] = []

    if (step.value === 0) {
      if (!formState.firstName) errors.push('First Name')
      if (!formState.lastName) errors.push('Last Name')
      if (!formState.age) errors.push('Age')
      if (!formState.email) errors.push('Email')
      if (!formState.address) errors.push('Address')
      if (!formState.city) errors.push('City')
      if (!formState.state) errors.push('State')
      if (!formState.zip) errors.push('Zip Code')
      if (!formState.phoneNumber) errors.push('Phone Number')
      if (!formState.adultMembersAgreed) errors.push('Household Agreement')
    }

    if (step.value === 1) {
      if (!formState.homeType) errors.push('Home Type')
      if (!formState.homeOwnership) errors.push('Own or Rent')
      if (formState.homeOwnership === 'Rent') {
        if (!formState.landlordName) errors.push('Landlord Name')
        if (!formState.landlordPhoneNumber) errors.push('Landlord Phone')
      }
    }

    if (step.value === 6) {
      if (!formState.agreementSignature1) errors.push('Signature 1')
      if (!formState.agreementSignature2) errors.push('Signature 2')
      if (!formState.signatureData) errors.push('Final Signature')
    }

    return errors
  })

  const isStepValid = computed(() => validationErrors.value.length === 0)

  const STORAGE_KEY = 'adoption_form_state'

  const persistState = () => {
    sessionStorage.setItem(
      STORAGE_KEY,
      JSON.stringify({
        step: step.value,
        formState: formState,
      }),
    )
  }

  const initFromStorage = () => {
    const stored = sessionStorage.getItem(STORAGE_KEY)
    if (stored) {
      try {
        const parsed = JSON.parse(stored)
        step.value = parsed.step || 0
        Object.assign(formState, parsed.formState)
      } catch (e) {
        console.error('Failed to restore adoption form state', e)
      }
    }
  }

  initFromStorage()

  const nextStep = () => {
    step.value++
    persistState()
    hasAttemptedSubmit.value = false
    return true
  }

  const prevStep = () => {
    if (step.value > 0) {
      step.value--
      persistState()
    }
  }

  const resetForm = () => {
    step.value = 0
    isSubmitted.value = false
    hasAttemptedSubmit.value = false
    sessionStorage.removeItem(STORAGE_KEY)
  }

  const submitApplication = async () => {
    try {
      let primaryOwner = null
      if (formState.primaryOwner !== null) {
        primaryOwner = formState.primaryOwner ? 'Yes' : 'No'
      }

      const payload = {
        ...formState,
        age: formState.age ? String(formState.age) : null,
        primaryOwner,
      }

      const response = await fetch(API_ENDPOINTS.ADOPTION_APPLICATION, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        const errorText = await response.text()

        console.error('Backend Submission Error:', response.status, errorText)
        throw new Error(`Server responded with ${response.status}: ${errorText}`)
      }

      const data = await response.json()
      console.log('Submission success:', data)
      isSubmitted.value = true
      sessionStorage.removeItem(STORAGE_KEY)
      return true
    } catch (error: unknown) {
      console.error('Error submitting application:', error)
      const message = error instanceof Error ? error.message : String(error)
      alert(`There was an error submitting your application: ${message}`)
      return false
    }
  }

  return {
    formState,
    step,
    isSubmitted,
    hasAttemptedSubmit,
    validationErrors,
    isStepValid,
    nextStep,
    prevStep,
    resetForm,
    persistState,
    submitApplication,
  }
})
