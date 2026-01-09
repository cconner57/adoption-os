import { defineStore } from 'pinia'
import { ref, reactive, computed } from 'vue'
import type { IVolunteerFormState } from '../models/volunteer-form'
import { API_ENDPOINTS } from '../constants/api'
import { useMetrics } from '../composables/useMetrics'

import { useDemoMode } from '../composables/useDemoMode'

export const useVolunteerStore = defineStore('volunteer', () => {
  const { isDemoMode } = useDemoMode()
  const isSubmitted = ref(false)
  const hasAttemptedSubmit = ref(false)
  const apiError = ref<string | null>(null)

  const formState = reactive<IVolunteerFormState>({
    fax_number: '',
    firstName: '',
    lastName: '',
    address: '',
    city: '',
    zip: '',
    phoneNumber: '',
    birthday: '',
    age: null,
    allergies: null,
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

  const validationErrors = computed(() => {
    const errors: string[] = []

    if (!formState.firstName) errors.push('First Name')
    if (!formState.lastName) errors.push('Last Name')
    if (!formState.address) errors.push('Address')
    if (!formState.city) errors.push('City')
    if (!formState.zip) errors.push('Zip Code')
    if (!formState.phoneNumber) errors.push('Phone Number')
    if (!formState.birthday) errors.push('Birthday')

    let isUnder21 = false
    if (formState.birthday) {
      let birthDate: Date
      if (formState.birthday.includes('-')) {
        birthDate = new Date(`${formState.birthday}T00:00:00`)
      } else {
        birthDate = new Date(formState.birthday)
      }

      if (!isNaN(birthDate.getTime())) {
        const today = new Date()
        let calculatedAge = today.getFullYear() - birthDate.getFullYear()
        const m = today.getMonth() - birthDate.getMonth()
        if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
          calculatedAge--
        }
        if (calculatedAge < 21) {
          isUnder21 = true
        }
      }
    }

    if (isUnder21 && formState.age === null) errors.push('Age')

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

  const isFormValid = computed(() => {
    if (isDemoMode.value) return true
    return validationErrors.value.length === 0
  })

  const isSubmitting = ref(false)

  const clearFormData = () => {
    formState.firstName = ''
    formState.lastName = ''
    formState.address = ''
    formState.city = ''
    formState.zip = ''
    formState.phoneNumber = ''
    formState.birthday = ''
    formState.age = null
    formState.allergies = null
    formState.emergencyContactName = ''
    formState.emergencyContactPhone = ''
    formState.volunteerExperience = ''
    formState.interestReason = ''
    formState.positionPreferences = []
    formState.availability = []
    formState.nameFull = ''
    formState.signatureData = null
    formState.signatureDate = ''
    formState.parentName = ''
    formState.parentSignatureData = null
    formState.parentSignatureDate = ''
  }

  const submit = async () => {
    hasAttemptedSubmit.value = true
    apiError.value = null

    if (!isFormValid.value && !isDemoMode.value) return false

    isSubmitting.value = true

    if (isDemoMode.value) {
      console.log('Demo Mode: Simulating submission success')
      await new Promise((resolve) => setTimeout(resolve, 1000))
      isSubmitted.value = true
      clearFormData()
      isSubmitting.value = false
      return true
    }

    const payload: Partial<IVolunteerFormState> = { ...formState }

    if (formState.age !== null && formState.age >= 21) {
      delete payload.parentName
      delete payload.parentSignatureData
      delete payload.parentSignatureDate
    }

    try {
      const response = await fetch(API_ENDPOINTS.VOLUNTEER_APPLICATION, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        if (response.status === 400 || response.status >= 500) {
          apiError.value = 'There was an error on the server. Please try again later.'
        } else {
          const errorData = await response.json().catch(() => ({}))
          apiError.value = errorData.error || `Server Error (${response.status})`
        }
        return false
      }

      const result = await response.json()
      if (result.error) {
        apiError.value = result.error
        return false
      }

      const { submitMetric } = useMetrics()
      submitMetric('form_submit', { form: 'volunteer' })

      isSubmitted.value = true
      clearFormData()
      return true
    } catch (error) {
      console.error('Network error:', error)
      apiError.value = 'Network error. Please try again later.'
      return false
    } finally {
      isSubmitting.value = false
    }
  }

  const resetForm = () => {
    isSubmitted.value = false
    hasAttemptedSubmit.value = false
    apiError.value = null
    clearFormData()
  }

  return {
    formState,
    isSubmitted,
    isSubmitting,
    hasAttemptedSubmit,
    apiError,
    validationErrors,
    isFormValid,
    submit,
    resetForm,
  }
})
