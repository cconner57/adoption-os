import { defineStore } from 'pinia'
import { computed,reactive, ref } from 'vue'

import { useDemoMode } from '../composables/useDemoMode'
import { useMetrics } from '../composables/useMetrics'
import type { FormState } from '../models/adopt-form'
import { usePetStore } from './pets'

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
    catAccess: [],
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

      // Fields required for everyone
      if (!formState.breedRestrictionsWeightLimit) errors.push('Breed Restrictions')
      if (!formState.monthlyFee) errors.push('Pet Deposit/Fee')
      if (!formState.allergies) errors.push('Allergies')
      if (!formState.primaryOwner) errors.push('Primary Owner')
      if (!formState.yearsAtAddress) errors.push('Years at Address')
      if (!formState.previousAddress) errors.push('Previous Address')
      if (!formState.expectToMove) errors.push('Expect to Move')
      if (!formState.travelPlan) errors.push('Travel Plan')
      if (formState.catAccess.length === 0) errors.push('Cat Access')
      if (!formState.catIndoorOutdoor) errors.push('Indoor/Outdoor')

      // Additional fields for Renters/Leasers
      if (formState.homeOwnership === 'Rent' || formState.homeOwnership === 'Lease') {
        if (!formState.landlordName) errors.push('Landlord Name')
        if (!formState.landlordPhoneNumber) errors.push('Landlord Phone')
        if (!formState.allowPets) errors.push('Pets Allowed')
      }
    }

    if (step.value === 2) {
      if (!formState.catPreferenceBreed) errors.push('Preferred Breed')
      if (!formState.catPreferencePhysical) errors.push('Physical Preferences')
      if (!formState.catPreferencePersonality) errors.push('Personality Preferences')
      if (!formState.catPreferenceNotWant) errors.push('Traits Not Wanted')
      if (!formState.whyInterested) errors.push('Why Interested')
      if (!formState.adoptionReason) errors.push('Adoption Reason')
      if (!formState.ownCatBefore) errors.push('Owned Cat Before')
      if (!formState.ownKittenBefore) errors.push('Owned Kitten Before')
      if (!formState.alreadyHaveVeterinarian) errors.push('Veterinarian')
      if (!formState.catAllowedHomeArea) errors.push('Allowed Areas')
      if (!formState.catHomeAloneHours) errors.push('Home Alone Hours')
      if (!formState.catDisciplineType) errors.push('Discipline Type')
      if (!formState.catEscapeSteps) errors.push('Escape Plan')
    }

    if (step.value === 3) {
      if (!formState.currentlyHavePets) errors.push('Do you have pets?')
      if (formState.currentlyHavePets === 'Yes') {
        formState.currentPets.forEach((pet, index) => {
          if (!pet.name) errors.push(`Pet ${index + 1} Name`)
          if (!pet.speciesBreedSize) errors.push(`Pet ${index + 1} Breed/Size`)
          if (!pet.age) errors.push(`Pet ${index + 1} Age`)
          if (!pet.likesDogs) errors.push(`Pet ${index + 1} Likes Dogs`)
          if (!pet.source) errors.push(`Pet ${index + 1} Source`)
          if (!pet.spayedNeutered) errors.push(`Pet ${index + 1} Spayed/Neutered`)
        })
      }
    }

    if (step.value === 4) {
      if (!formState.ownPetsBefore) errors.push('Have you owned pets?')
      if (formState.ownPetsBefore === 'Yes') {
        formState.pastPets.forEach((pet, index) => {
          if (!pet.name) errors.push(`Past Pet ${index + 1} Name`)
          if (!pet.speciesBreedSize) errors.push(`Past Pet ${index + 1} Breed/Size`)
          if (!pet.age) errors.push(`Past Pet ${index + 1} Age`)
          if (!pet.source) errors.push(`Past Pet ${index + 1} Source`)
          if (!pet.spayedNeutered) errors.push(`Past Pet ${index + 1} Spayed/Neutered`)
          if (!pet.passedAwayReason) errors.push(`Past Pet ${index + 1} Outcome`)
        })
      }
    }

    if (step.value === 5) {
      if (!formState.bredAnimalDescription) errors.push('Bred Animals')
      if (!formState.ownedDeclawedOrDebarked) errors.push('Declawed/Debarked')
      if (!formState.movedWithPet) errors.push('Moved with Pet')
      if (!formState.ownedSpecialNeedsPet) errors.push('Special Needs Pet')
      if (!formState.mobilityDevice) errors.push('Mobility Device')
      if (!formState.foodTypeBrand) errors.push('Food Type')
      if (formState.surrenderConditions.length === 0) errors.push('Surrender Conditions')
      if (!formState.surrenderPlan) errors.push('Surrender Plan')
      if (!formState.affordVetCare) errors.push('Vet Care')
      if (!formState.affordEmergencyCost) errors.push('Emergency Cost')
    }

    if (step.value === 6) {
      if (!formState.agreementSignature1) errors.push('Commitment Signature')
      if (!formState.agreementSignature2) errors.push('Homecheck Consent')
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
      if (isDemoMode.value) {
        // Simulate API call for Demo Mode
        await new Promise((resolve) => setTimeout(resolve, 1500))
        isSubmitted.value = true
        sessionStorage.removeItem(STORAGE_KEY)
        return true
      }

      // Prepare payload matching backend AdoptionApplication struct
      const petStore = usePetStore()
      const payload = {
        petId: petStore.selectedPet?.id || null,
        petName: petStore.selectedPet?.petName || null,
        firstName: formState.firstName,
        lastName: formState.lastName,
        age: Number(formState.age), // Ensure number
        spouseFirstName: formState.spouseFirstName,
        spouseLastName: formState.spouseLastName,
        roommatesNames: formState.roommatesNames,
        childrenNamesAges: formState.childrenNamesAges,
        email: formState.email,
        address: formState.address,
        addressLine2: formState.addressLine2,
        city: formState.city,
        state: formState.state,
        zip: formState.zip,
        phoneNumber: formState.phoneNumber,
        cellPhoneNumber: formState.cellPhoneNumber,
        adultMembersAgreed: formState.adultMembersAgreed,

        homeType: formState.homeType,
        homeOwnership: formState.homeOwnership,
        landlordName: formState.landlordName,
        landlordPhoneNumber: formState.landlordPhoneNumber,
        allowPets: formState.allowPets,
        breedRestrictionsWeightLimit: formState.breedRestrictionsWeightLimit,
        monthlyFee: formState.monthlyFee,
        allergies: formState.allergies,
        primaryOwner: formState.primaryOwner,
        yearsAtAddress: formState.yearsAtAddress,
        previousAddress: formState.previousAddress,
        expectToMove: formState.expectToMove,
        travelPlan: formState.travelPlan,

        catAccess: formState.catAccess.length > 0 ? formState.catAccess.join(', ') : null, // Convert array to string as backend expects *string
        catIndoorOutdoor: formState.catIndoorOutdoor,
        catPreferenceBreed: formState.catPreferenceBreed,
        catPreferencePhysical: formState.catPreferencePhysical,
        catPreferencePersonality: formState.catPreferencePersonality,
        catPreferenceNotWant: formState.catPreferenceNotWant,
        whyInterested: formState.whyInterested,
        adoptionReason: formState.adoptionReason,
        ownCatBefore: formState.ownCatBefore,
        ownKittenBefore: formState.ownKittenBefore,
        alreadyHaveVeterinarian: formState.alreadyHaveVeterinarian,
        catAllowedHomeArea: formState.catAllowedHomeArea,
        catHomeAloneHours: formState.catHomeAloneHours,
        catDisciplineType: formState.catDisciplineType,
        catEscapeSteps: formState.catEscapeSteps,

        currentPets: formState.currentPets || [],
        currentlyHavePets: formState.currentlyHavePets,

        pastPets: formState.pastPets || [],
        ownPetsBefore: formState.ownPetsBefore,

        bredAnimalDescription: formState.bredAnimalDescription,
        ownedDeclawedOrDebarked: formState.ownedDeclawedOrDebarked,
        movedWithPet: formState.movedWithPet,
        ownedSpecialNeedsPet: formState.ownedSpecialNeedsPet,
        mobilityDevice: formState.mobilityDevice,
        surrenderConditions: formState.surrenderConditions, // Backend expects []string, usually
        surrenderPlan: formState.surrenderPlan,
        foodTypeBrand: formState.foodTypeBrand,
        affordVetCare: formState.affordVetCare,
        affordEmergencyCost: formState.affordEmergencyCost,

        agreementSignature1: formState.agreementSignature1,
        agreementSignature2: formState.agreementSignature2,
        agreementSignature3: formState.agreementSignature3,
        signatureData: formState.signatureData,

        fax_number: formState.fax_number // Honeypot
      }

      // Convert catAccess array to single string if needed, but backend struct has it as *string
      // Re-verifying struct: CatAccess *string `json:"catAccess"`
      // Frontend is string[]. Must join it. Done above.

      const response = await fetch('/applications/adoption', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}))
        throw new Error(errorData.error || errorData.message || 'Submission failed')
      }

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
