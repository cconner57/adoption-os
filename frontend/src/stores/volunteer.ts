import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { IVolunteerFormState } from '../models/volunteer-form'

export const useVolunteerStore = defineStore('volunteer', () => {
  const isSubmitted = ref(false)
  const hasAttemptedSubmit = ref(false)
  const apiError = ref<string | null>(null)

  const formState = reactive<IVolunteerFormState>({
    firstName: '',
    lastName: '',
    address: '',
    city: '',
    zip: '',
    phoneNumber: '',
    birthday: '',
    age: null,
    allergies: false,
    emergencyContactName: '',
    emergencyContactPhone: '',
    volunteerExperience: '',
    interestReason: '',
    positionPreferences: [],
    availability: [],
    nameFull: '',
    signatureData: null,
    signatureDate: '',
    parentName: '',
    parentSignatureData: null,
    parentSignatureDate: '',
  })

  // Sanitization helpers could technically be here or utility functions

  // Validation Logic
  const validationErrors = computed(() => {
    const errors: string[] = []
    // ... migration of validation logic from Volunteer.vue
    if (!formState.firstName) errors.push('First Name')
    if (!formState.lastName) errors.push('Last Name')
    if (!formState.address) errors.push('Address')
    if (!formState.city) errors.push('City')
    if (!formState.zip) errors.push('Zip Code')
    if (!formState.phoneNumber) errors.push('Phone Number')
    if (!formState.birthday) errors.push('Birthday')
    if (formState.age === null) errors.push('Age')
    if (!formState.emergencyContactName) errors.push('Emergency Contact Name')
    if (!formState.emergencyContactPhone) errors.push('Emergency Contact Phone')
    if (!formState.interestReason) errors.push('Interest Reason')
    if (formState.positionPreferences.length === 0) errors.push('Position Preferences')
    if (formState.availability.length === 0) errors.push('Availability')
    if (!formState.nameFull) errors.push('Agreement Name')
    if (!formState.signatureDate) errors.push('Agreement Date')
    if (!formState.signatureData) errors.push('Signature')

    if (formState.age !== null && formState.age < 21) {
      if (!formState.parentName) errors.push('Parent Name')
      if (!formState.parentSignatureDate) errors.push('Parent Date')
      if (!formState.parentSignatureData) errors.push('Parent Signature')
    }

    return errors
  })

  const isFormValid = computed(() => validationErrors.value.length === 0)

  const submit = async () => {
    hasAttemptedSubmit.value = true
    apiError.value = null

    if (!isFormValid.value) return false

    // API logic usually stays in the store for clean components
    const payload: Partial<IVolunteerFormState> = { ...formState }
    if (formState.age !== null && formState.age >= 21) {
      delete payload.parentName
      delete payload.parentSignatureData
      delete payload.parentSignatureDate
    }

    try {
      const response = await fetch('http://localhost:8080/applications/volunteer', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}))
        apiError.value = errorData.error || `Server Error (${response.status})`
        return false
      }

      const result = await response.json()
      console.log('Form submitted successfully:', result)
      isSubmitted.value = true
      return true
    } catch (error) {
      console.error('Network error:', error)
      apiError.value = 'Network error. Please try again later.'
      return false
    }
  }

  const resetForm = () => {
    isSubmitted.value = false
    hasAttemptedSubmit.value = false
    apiError.value = null
    // Reset fields...
  }

  return {
    formState,
    isSubmitted,
    hasAttemptedSubmit,
    apiError,
    validationErrors,
    isFormValid,
    submit,
    resetForm,
  }
})
