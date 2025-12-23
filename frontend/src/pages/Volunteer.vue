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
import { computed, reactive } from 'vue'

type FormInput = string | number | null

const formatPhoneNumber = (value: FormInput): string => {
  if (!value) return ''
  const digits = String(value).replace(/\D/g, '').substring(0, 10)
  if (digits.length === 0) return ''
  if (digits.length <= 3) return `(${digits}`
  if (digits.length <= 6) return `(${digits.slice(0, 3)})${digits.slice(3)}`
  return `(${digits.slice(0, 3)})${digits.slice(3, 6)}-${digits.slice(6)}`
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
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
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
  allergies: '',
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

const handleBlur = (field: string) => {
  touched[field] = true
}

const isFormValid = computed(() => {
  // Required Personal Info
  if (!formState.firstName || !formState.lastName || !formState.address || !formState.city || !formState.zip || !formState.phoneNumber || !formState.birthday || formState.age === null || !formState.allergies || !formState.emergencyContactName || !formState.emergencyContactPhone) {
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
  if (!formState.allergies) errors.push('Allergies (enter "None" if applicable)')
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

const handleSubmit = () => {
  if (!isFormValid.value) {
    // Mark all as touched to show errors on explicit submit attempt
    Object.keys(formState).forEach(key => touched[key] = true)
    return
  }

  console.log('Form submitted successfully:', formState)
  alert('Application submitted successfully!')
  // TODO: Add API call here
}
</script>

<template>
  <section class="page-shell">
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
          maxlength="10"
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
          maxlength="15"
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
          :hasError="touched.birthday && !formState.birthday"
          @blur="handleBlur('birthday')"
        />
        <InputField
          v-model="formState.age"
          label="If under 21, Age"
          placeholder="Age"
          type="number"
          name="age"
          min="0"
          max="100"
          :hasError="touched.age && formState.age === null"
          @blur="handleBlur('age')"
        />

        <Allergies
          v-model="formState.allergies"
          :class="{ 'has-error': touched.allergies && !formState.allergies }"
          @vnode-updated="() => {} /* simple placeholder, complex components need internal error logic */"
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
          maxlength="15"
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

        <PositionPreferences v-model="formState.positionPreferences" />
      </fieldset>

      <Availability v-model="formState.availability" />

      <Agreement
        :name="formState.firstName + ' ' + formState.lastName"
        :fullName="formState.nameFull"
        @update:fullName="(val: string) => (formState.nameFull = sanitizeName(val))"
        :age="formState.age!"
        :signature="formState.signatureData"
        :signature-date="formState.signatureDate"
        :parent-name="formState.parentName"
        @update:parentName="(val: string) => (formState.parentName = sanitizeName(val))"
        :parent-signature="formState.parentSignatureData"
        :parent-date="formState.parentSignatureDate"
      />

      <div v-if="!isFormValid && validationErrors.length > 0 && Object.keys(touched).length > 0" class="validation-summary">
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
          :disabled="!isFormValid"
        />
      </div>
    </form>
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
</style>
