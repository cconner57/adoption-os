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
import { reactive, onMounted } from 'vue'
import FormSubmitted from '../components/common/form-submitted/FormSubmitted.vue'
import { storeToRefs } from 'pinia'
import { useVolunteerStore } from '../stores/volunteer.ts'
import { useRouter } from 'vue-router'
import { useMetrics } from '../composables/useMetrics'

const { submitMetric } = useMetrics()

onMounted(() => {
  submitMetric('form_start', { form: 'volunteer' })
})

type FormInput = string | number | null

const formatPhoneNumber = (value: FormInput): string => {
  if (!value) return ''
  const digits = String(value).replace(/\D/g, '').substring(0, 10)
  if (digits.length === 0) return ''
  if (digits.length <= 3) return `(${digits}`
  if (digits.length <= 6) return `(${digits.slice(0, 3)})${digits.slice(3)}`
  return `(${digits.slice(0, 3)})${digits.slice(3, 6)}-${digits.slice(6, 10)}`
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
  return String(value).replace(/\D/g, '').substring(0, 5)
}
const sanitizeAddress = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

const volunteerStore = useVolunteerStore()
const {
  formState,
  isSubmitted,
  isSubmitting,
  hasAttemptedSubmit,
  apiError,
  validationErrors,
  isFormValid,
} = storeToRefs(volunteerStore)
const { submit, resetForm } = volunteerStore

const touched = reactive<Record<string, boolean>>({})

const handleBlur = (field: string) => {
  touched[field] = true
}

const handleSubmit = async () => {
  if (!isFormValid.value) {
    Object.keys(formState.value).forEach((key) => (touched[key] = true))
  }

  const success = await submit()

  if (success) {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const router = useRouter()

const handleReset = () => {
  resetForm()
  router.push('/')
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
            @update:modelValue="
              (val) => {
                const str = String(val).replace(/\D/g, '').substring(0, 3)
                formState.age = str ? Number(str) : null
              }
            "
            label="If under 21, Age"
            placeholder="Age"
            type="text"
            name="age"
            maxlength="3"
            :hasError="touched.age && validationErrors.includes('Age')"
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

        <div
          v-if="hasAttemptedSubmit && !isFormValid && validationErrors.length > 0"
          class="validation-summary"
        >
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
            :loading="isSubmitting"
          />
        </div>
      </form>
    </div>

    <FormSubmitted v-if="isSubmitted" @reset="handleReset" formType="volunteer" />
  </section>
</template>

<style scoped>
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 9rem var(--layout-padding-side) 64px;

  /* Viewport adjustments can remain, but simplified */
  @media (max-width: 430px) {
    padding: 6rem 16px 32px;
  }

  .form-container {
    container-type: inline-size;
    container-name: form-card;
    max-width: 1600px;
    margin: 0 auto;
  }

  .form-card {
    background: var(--white);
    color: var(--font-color-dark);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;

    /* Responsive padding based on container width */
    @container form-card (max-width: 768px) {
      padding: 32px 24px;
    }

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

      /* Switch to grid when container is wide enough */
      @container form-card (min-width: 700px) {
        display: grid;
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
    }
  }

  /* Nested components that are localized */
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

  /* Utility classes nested here or could be global, but scoped so kept here */
  .has-error :deep(input),
  .has-error :deep(textarea) {
    border-color: #ef4444 !important;
    outline: 2px solid #ef4444 !important;
  }
}
</style>
