<script setup lang="ts">
import {
  Agreement,
  Allergies,
  ApplicationHeader,
  Availability,
  PositionPreferences,
} from '../components/volunteer/index.ts'
import Button from '../components/common/ui/Button.vue'
import InputField from '../components/common/ui/InputField.vue'
import InputTextArea from '../components/common/ui/InputTextArea.vue'
import type { IVolunteerFormState } from '../models/volunteer-form.ts'
import { computed, reactive, ref } from 'vue'

type FormInput = string | number | null

const formatPhoneNumber = (value: FormInput): string => {
  if (!value) return ''
  const digits = String(value).replace(/\D/g, '').substring(0, 10)
  if (digits.length === 0) return ''
  if (digits.length <= 3) return `(${digits}`
  if (digits.length <= 6) return `(${digits.slice(0, 3)})${digits.slice(3)}`
  return `(${digits.slice(0, 3)})${digits.slice(3, 6)}-${digits.slice(6, 10)}` // Limit to 10 digits total (approx 13 chars with parens/dash)
}

const sanitizeName = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 ]/g, '')
}

const sanitizeCity = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}


const sanitizeZip = (value: FormInput): string => {
  if (!value) return ''
  // Allow only digits, max 5 chars
  return String(value).replace(/\D/g, '').substring(0, 5)
}


const sanitizeAddress = (value: FormInput): string => {
  if (!value) return ''
  // Strict: Alphanumeric + Space + Dash only.
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

const formState = reactive<IVolunteerFormState>({
  firstName: '',
  lastName: '',
  address: '',
  city: '',
  zip: '',
  phoneNumber: '',
  birthday: '',
  age: null as number | null,
  allergies: false,
  emergencyContactName: '',
  emergencyContactPhone: '',
  volunteerExperience: '',
  interestReason: '',
  positionPreferences: [] as string[],
  availability: [] as string[],
  nameFull: '',
  signatureData: null as string | null,
  signatureDate: '',
  parentName: '',
  parentSignatureData: null as string | null,
  parentSignatureDate: '',
})

// Track which fields have been touched/blurred by the user
const touched = reactive<Record<string, boolean>>({})
const touched = reactive<Record<string, boolean>>({})
const isSubmitted = ref(false)
const hasAttemptedSubmit = ref(false)
const apiError = ref<string | null>(null)

const handleBlur = (field: string) => {
  touched[field] = true
}

const isFormValid = computed(() => {
  // Required Personal Info
  if (!formState.firstName || !formState.lastName || !formState.address || !formState.city || !formState.zip || !formState.phoneNumber || !formState.birthday || formState.age === null || !formState.emergencyContactName || !formState.emergencyContactPhone) {
    return false
  }

  // Experience & Interests
  if (!formState.interestReason || formState.positionPreferences.length === 0 || formState.availability.length === 0) {
    return false
  }

  // Agreement
  if (!formState.nameFull || !formState.signatureDate || !formState.signatureData) {
    return false
  }

  // Parental Consent (Under 21)
  if (formState.age !== null && formState.age < 21) {
    if (!formState.parentName || !formState.parentSignatureDate || !formState.parentSignatureData) {
      return false
    }
  }

  return true
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
  if (formState.age === null) errors.push('Age')
  // Allergies is optional/defaults to "None" logic handled by backend or not required
  if (!formState.emergencyContactName) errors.push('Emergency Contact Name')
  if (!formState.emergencyContactPhone) errors.push('Emergency Contact Phone')
  if (!formState.volunteerExperience && false) errors.push('Experience') // Optional
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

const handleSubmit = async () => {
  hasAttemptedSubmit.value = true
  apiError.value = null

  if (!isFormValid.value) {
    // Mark all as touched to show errors on explicit submit attempt
    Object.keys(formState).forEach(key => touched[key] = true)
    return
  }

  // Create a copy of formState to modify for payload
  const payload = { ...formState }

  // Remove parent fields if age is 21 or older
  if (payload.age !== null && payload.age >= 21) {
    delete (payload as any).parentName
    delete (payload as any).parentSignatureData
    delete (payload as any).parentSignatureDate
  }

  try {
    const response = await fetch('http://localhost:8080/applications/volunteer', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      const errorMessage = errorData.error || `Server Error (${response.status})`
      console.error('Submission failed:', errorData)
      apiError.value = errorMessage
      return
    }

    const result = await response.json()
    console.log('Form submitted successfully:', result)
    const result = await response.json()
    console.log('Form submitted successfully:', result)
    isSubmitted.value = true
    window.scrollTo({ top: 0, behavior: 'smooth' })

  } catch (error) {
    console.error('Network error:', error)
    apiError.value = 'Network error. Please try again later.'
  }
}

const handleReset = () => {
  isSubmitted.value = false
  window.location.reload()
}
</script>

<template>
  <section class="page-shell">
    <div v-if="!isSubmitted" class="form-container">
      <form class="form-card" aria-labelledby="form-title" @submit.prevent="handleSubmit">
      <ApplicationHeader
        header-title="Volunteer"
        header-text="I Dream of Home Rescue (IDOHR) is an all-volunteer, nonprofit dedicated to helping homeless cats
    and kittens find loving, permanent homes. We’re looking for responsible volunteers to help with
    feeding and cleaning, socializing cats and kittens, supporting adoptions, and light
    administrative tasks. Volunteers must be 21 or older. If under 21, a parent or guardian must
    sign the waiver below. Join us and help change a life. You’ll connect with amazing animals, work
    alongside caring volunteers, and make a meaningful impact."
      />
      <fieldset class="volunteer-grid" aria-labelledby="pi">
        <legend id="pi" class="section-title">Personal Information</legend>
      </form>
    </div>

    <!-- Success State -->
    <div v-else class="success-card">
        <div class="icon-wrapper">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="48"
            height="48"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="success-icon"
          >
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
            <polyline points="22 4 12 14.01 9 11.01"></polyline>
          </svg>
        </div>
        <h3 class="success-title">Application Received!</h3>
        <p class="success-message">Thank you for volunteering! We'll review your application and get back to you shortly.</p>
        <div class="success-actions">
          <Button title="Return to Home" color="green" @click="handleReset" />
        </div>
    </div>

        <InputField
          :modelValue="formState.firstName"
          @update:modelValue="(val) => (formState.firstName = sanitizeName(val))"
          label="First Name"
          placeholder="First name"
          autocomplete="given-name"
          name="firstName"
          maxlength="50"
          :hasError="touched.firstName && !formState.firstName"
          @blur="handleBlur('firstName')"
        />
        <InputField
          :modelValue="formState.lastName"
          @update:modelValue="(val) => (formState.lastName = sanitizeName(val))"
          label="Last Name"
          placeholder="Last name"
          autocomplete="family-name"
          name="lastName"
          maxlength="50"
          :hasError="touched.lastName && !formState.lastName"
          @blur="handleBlur('lastName')"
        />
        <InputField
          :modelValue="formState.address"
          @update:modelValue="(val) => (formState.address = sanitizeAddress(val))"
          label="Address"
          placeholder="Address"
          autocomplete="street-address"
          name="address"
          maxlength="100"
          :hasError="touched.address && !formState.address"
          @blur="handleBlur('address')"
        />
        <InputField
          :modelValue="formState.city"
          @update:modelValue="(val) => (formState.city = sanitizeCity(val))"
          label="City"
          placeholder="City"
          autocomplete="address-level2"
          name="city"
          maxlength="50"
          :hasError="touched.city && !formState.city"
          @blur="handleBlur('city')"
        />
        <InputField
          :modelValue="formState.zip"
          @update:modelValue="(val) => (formState.zip = sanitizeZip(val))"
          label="Zip"
          placeholder="Zip"
          type="text"
          autocomplete="postal-code"
          name="zip"
          maxlength="5"
          :hasError="touched.zip && !formState.zip"
          @blur="handleBlur('zip')"
        />
        <InputField
          :modelValue="formState.phoneNumber"
          @update:modelValue="(val) => (formState.phoneNumber = formatPhoneNumber(val))"
          label="Phone Number"
          placeholder="Phone"
          type="tel"
          autocomplete="tel"
          name="phoneNumber"
          maxlength="13"
          :hasError="touched.phoneNumber && !formState.phoneNumber"
          @blur="handleBlur('phoneNumber')"
        />
        <InputField
          v-model="formState.birthday"
          label="Birthday"
          placeholder="mm/dd/yyyy"
          type="date"
          autocomplete="bday"
          name="birthday"
          max="9999-12-31"
          :hasError="touched.birthday && !formState.birthday"
          @blur="handleBlur('birthday')"
        />
        <InputField
          :modelValue="formState.age"
          @update:modelValue="(val) => {
             const str = String(val).replace(/\D/g, '').substring(0, 3);
             formState.age = str ? Number(str) : null;
          }"
          label="If under 21, Age"
          placeholder="Age"
          type="text"
          name="age"
          maxlength="3"
          :hasError="touched.age && formState.age === null && false"
          @blur="handleBlur('age')"
        />

        <Allergies
          v-model="formState.allergies"
          :class="{ 'has-error': touched.allergies && !formState.allergies }"
        />

        <InputField
          :modelValue="formState.emergencyContactName"
          @update:modelValue="(val) => (formState.emergencyContactName = sanitizeName(val))"
          label="Emergency Contact Person"
          placeholder="Name"
          name="emergencyContactName"
          autocomplete="off"
          maxlength="100"
          :hasError="touched.emergencyContactName && !formState.emergencyContactName"
          @blur="handleBlur('emergencyContactName')"
        />
        <InputField
          :modelValue="formState.emergencyContactPhone"
          @update:modelValue="(val) => (formState.emergencyContactPhone = formatPhoneNumber(val))"
          label="Phone Number"
          placeholder="Phone Number"
          type="tel"
          name="emergencyContactPhone"
          autocomplete="off"
          maxlength="13"
          :hasError="touched.emergencyContactPhone && !formState.emergencyContactPhone"
          @blur="handleBlur('emergencyContactPhone')"
        />
      </fieldset>

      <fieldset aria-labelledby="exp">
        <legend id="exp" class="section-title">Experience & Interests</legend>

        <InputTextArea
          v-model="formState.volunteerExperience"
          label="Volunteer Experience (if any):"
          placeholder="Briefly describe"
          name="volunteerExperience"
          maxlength="500"
        />

        <InputTextArea
          v-model="formState.interestReason"
          label="Why are you interested in being a volunteer:"
          placeholder="Your reason"
          name="interestReason"
          maxlength="500"
          :class="{ 'has-error': touched.interestReason && !formState.interestReason }"
          @blur="handleBlur('interestReason')"
        />

        <PositionPreferences
          v-model="formState.positionPreferences"
          :hasError="touched.positionPreferences && formState.positionPreferences.length === 0"
        />
      </fieldset>

      <Availability
        v-model="formState.availability"
        :hasError="touched.availability && formState.availability.length === 0"
      />

      <Agreement
        :name="formState.firstName + ' ' + formState.lastName"
        v-model:fullName="formState.nameFull"
        :age="formState.age!"
        v-model:signature="formState.signatureData"
        v-model:signatureDate="formState.signatureDate"
        v-model:parentName="formState.parentName"
        v-model:parentSignature="formState.parentSignatureData"
        v-model:parentDate="formState.parentSignatureDate"
        :hasNameError="touched.nameFull && !formState.nameFull"
        :hasDateError="touched.signatureDate && !formState.signatureDate"
        :hasSignatureError="touched.signatureData && !formState.signatureData"
        :hasParentNameError="touched.parentName && !formState.parentName"
        :hasParentDateError="touched.parentSignatureDate && !formState.parentSignatureDate"
        :hasParentSignatureError="touched.parentSignatureData && !formState.parentSignatureData"
      />

      <div v-if="apiError" class="validation-summary error-alert">
        <p class="summary-title">Submission Error</p>
        <p>{{ apiError }}</p>
      </div>

      <div v-if="hasAttemptedSubmit && !isFormValid && validationErrors.length > 0" class="validation-summary">
        <p class="summary-title">Please complete the following required fields:</p>
        <div class="tags">
          <span v-for="err in validationErrors" :key="err" class="tag is-danger">{{ err }}</span>
        </div>
      </div>

      <div class="actions">
        <Button
          type="submit"
          title="Submit Application"
          color="green"
          size="large"
        />
      </div>
    </form>
    </div>
  </section>
</template>

<style scoped>
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
    .volunteer-grid {
      display: flex;
      flex-direction: column;
      gap: 16px;

      @media (min-width: 768px) {
        display: grid;
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
    }
  }
  @media (min-width: 321px) and (max-width: 430px) {
    padding: 6rem 16px 32px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
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

/* Helper for manual has-error class on non-InputField components if needed */
.has-error :deep(input),
.has-error :deep(textarea) {
  border-color: #ef4444 !important;
  outline: 2px solid #ef4444 !important;
}

/* Success State Styles */
.success-card {
  max-width: 600px;
  margin: 0 auto;
  background: var(--white);
  border-radius: 24px;
  padding: 60px 32px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  text-align: center;
  animation: scaleIn 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.icon-wrapper {
  color: var(--green);
  background-color: color-mix(in srgb, var(--green) 10%, white);
  width: 100px;
  height: 100px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 32px;
  animation: popIn 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) 0.1s both;
}

.success-icon {
  width: 48px;
  height: 48px;
}

.success-title {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 16px;
  color: var(--font-color-dark);
}

.success-message {
  color: #374151;
  font-size: 1.1rem;
  margin-bottom: 40px;
  line-height: 1.6;
  max-width: 400px;
}

.success-actions {
  display: flex;
  justify-content: center;
  width: 100%;
}

@keyframes scaleIn {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: scale(1); }
}

@keyframes popIn {
  from { opacity: 0; transform: scale(0.5); }
  to { opacity: 1; transform: scale(1); }
}
</style>
