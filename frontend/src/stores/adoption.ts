import { defineStore } from 'pinia'
import { computed,reactive, ref } from 'vue'

import { useDemoMode } from '../composables/useDemoMode'
import { useMetrics } from '../composables/useMetrics'
import type { FormState } from '../models/adopt-form'
import { mockApplications } from './mockApplications'

export const useAdoptionStore = defineStore('adoption', () => {
  const { isDemoMode } = useDemoMode()
  const step = ref(0)
  const isSubmitted = ref(false)
  const hasAttemptedSubmit = ref(false)

  const formState = reactive<FormState>({
    fax_number: '',
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
    adultMembersAgreed: null,
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

  const isStepValid = computed(() => {
    if (isDemoMode.value) return true
    return validationErrors.value.length === 0
  })

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

  const { submitMetric } = useMetrics()

  const nextStep = () => {
    step.value++
    submitMetric('form_step', { form: 'adoption', step: step.value })
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
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 1500))

      const petDetails = sessionStorage.getItem('adoption_pet')
      let petInfo = { petId: 'unknown', petName: 'Unknown' }
      if (petDetails) {
        try {
          petInfo = JSON.parse(petDetails)
        } catch (e) {
          console.error(e)
        }
      }

      // Add to mock store
      const newApp = {
        id: `a-app-${Date.now()}`,
        type: 'adoption' as const,
        applicantName: `${formState.firstName} ${formState.lastName}`,
        email: formState.email || '',
        date: new Date().toISOString().split('T')[0],
        status: 'pending' as const,
        expiresAt: Date.now() + 7 * 24 * 60 * 60 * 1000, // 7 days from now
        details: {
          petName: petInfo.petName,
          petId: petInfo.petId,
          homeType: formState.homeType,
          otherPets: formState.currentPets.length > 0 || formState.currentlyHavePets === 'Yes',
        },
        fullApplication: {
          Applicant: `${formState.firstName} ${formState.lastName}`,
          Phone: formState.phoneNumber,
          Address: `${formState.address}, ${formState.city}`,
          Housing: `${formState.homeType} (${formState.homeOwnership})`,
          'Landlord Contact': formState.landlordName ? `${formState.landlordName} (${formState.landlordPhoneNumber})` : 'N/A',
          'Household Members': `Spouse: ${formState.spouseFirstName || 'N/A'}`,
          'Current Pets': formState.currentPets.map(p => `${p.speciesBreedSize} (${p.name})`).join(', ') || 'None',
          'Vet Reference': formState.alreadyHaveVeterinarian || 'N/A',
          'Time alone': formState.catHomeAloneHours || 'N/A',
        },
      }

      mockApplications.value.unshift(newApp)

      // Simulate Email
      console.log(`[Email Service] Sending adoption confirmation to ${formState.email}...`)
      console.log(`[Email Service] Subject: Application Received for ${petInfo.petName}`)
      console.log(`[Email Service] Body: Dear ${formState.firstName}, we have received your application...`)

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
