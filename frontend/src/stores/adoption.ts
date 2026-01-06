import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { FormState } from '../models/adopt-form'

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

  // Validation Logic
  const validationErrors = computed(() => {
    const errors: string[] = []

    if (step.value === 0) {
      // General Section / Applicant Info (Mapped to Step 0 in PetAdoption.vue logic)
      // Note: Previous logic in PetAdoption.vue step 0 checked these:
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

    // Add validation for other steps as needed, matching the sections in PetAdoption.vue
    // Step 1: Home Section
    // Step 2: New Cat Section
    // Step 3: Current Pets
    // Step 4: Past Pets
    // Step 5: Other Section
    // Step 6: Summary

    if (step.value === 1) {
      // Home Section
      if (!formState.homeType) errors.push('Home Type')
      if (!formState.homeOwnership) errors.push('Own or Rent')
      if (formState.homeOwnership === 'Rent') {
        if (!formState.landlordName) errors.push('Landlord Name')
        if (!formState.landlordPhoneNumber) errors.push('Landlord Phone')
      }
    }

    // ... (Continue adding validation as previously intended but with correct keys)

    if (step.value === 6) {
      // Summary
      if (!formState.agreementSignature1) errors.push('Signature 1')
      if (!formState.agreementSignature2) errors.push('Signature 2')
      // if (!formState.agreementSignature3) errors.push('Signature 3')
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

  // Initialize immediately
  initFromStorage()

  const nextStep = () => {
    // Validation bypassed for testing
    // hasAttemptedSubmit.value = true
    // if (!isStepValid.value) return false

    console.log(`[AdoptionStore] Next Step Clicked. Moving from ${step.value} to ${step.value + 1}`)
    console.log('[AdoptionStore] Current Form State:', JSON.parse(JSON.stringify(formState)))

    step.value++
    persistState() // Save state on navigation
    hasAttemptedSubmit.value = false
    return true
  }

  const prevStep = () => {
    if (step.value > 0) {
      step.value--
      persistState() // Save state on navigation
    }
  }

  const resetForm = () => {
    step.value = 0
    isSubmitted.value = false
    hasAttemptedSubmit.value = false
    sessionStorage.removeItem(STORAGE_KEY)

    // Reset formState fields (simplified reset for now, assuming page reload or explicit re-navigation handles full reset often)
    // To truly reset reactive object in place without page reload:
    // Object.keys(formState).forEach(key => formState[key] = null) // simplest approximation
    // But specific arrays need []
    // For now, removing from storage is the key request.
  }

  const submitApplication = async () => {
    try {
      // Create a sanitized payload to ensure types match backend expectations
      const payload = {
        ...formState,
        age: formState.age ? String(formState.age) : null,
        primaryOwner:
          formState.primaryOwner === true ? 'Yes' : formState.primaryOwner === false ? 'No' : null,
      }

      const response = await fetch('http://localhost:8080/applications/adoption', {
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
    } catch (error: any) {
      console.error('Error submitting application:', error)
      alert(`There was an error submitting your application: ${error.message}`)
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
