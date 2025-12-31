<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/surrender-form.ts'
import InputTextArea from '../../common/ui/InputTextArea.vue'
import InputSelectGroup from '../../common/ui/InputSelectGroup.vue'
import ButtonToggle from '../../common/ui/ButtonToggle.vue'

import {
  formatPhoneNumber,
  sanitizeName,
  sanitizeCity,
  sanitizeZip,
  sanitizeAddress,
} from '../../../utils/validators.ts'

const { formState, touched, handleBlur, hasAttemptedSubmit, selectedAnimal } = defineProps<{
  formState: SurrenderFormState
  touched: Record<string, boolean>
  handleBlur: (_field: string) => void // eslint-disable-line no-unused-vars
  hasAttemptedSubmit: boolean
  selectedAnimal: string
}>()
</script>

<template>
  <div class="household-section">
    <h5>{{ selectedAnimal }} & Household Information</h5>
    <fieldset class="household-grid">
      <InputField
        label="First Name"
        placeholder="First Name"
        :modelValue="formState.firstName"
        @update:modelValue="(val) => (formState.firstName = sanitizeName(val))"
        :hasError="(touched.firstName && !formState.firstName) || (hasAttemptedSubmit && !formState.firstName)"
        @blur="handleBlur('firstName')"
      />
      <InputField
        label="Last Name"
        placeholder="Last Name"
        :modelValue="formState.lastName"
        @update:modelValue="(val) => (formState.lastName = sanitizeName(val))"
        :hasError="(touched.lastName && !formState.lastName) || (hasAttemptedSubmit && !formState.lastName)"
        @blur="handleBlur('lastName')"
      />
      <InputField
        label="Phone Number"
        placeholder="Phone Number"
        :modelValue="formState.phoneNumber"
        @update:modelValue="(val) => (formState.phoneNumber = formatPhoneNumber(val))"
        maxlength="13"
        :hasError="(touched.phoneNumber && !formState.phoneNumber) || (hasAttemptedSubmit && !formState.phoneNumber)"
        @blur="handleBlur('phoneNumber')"
      />
      <InputField
        label="Email"
        placeholder="Email"
        v-model="formState.email"
        :hasError="(touched.email && !formState.email) || (hasAttemptedSubmit && !formState.email)"
        @blur="handleBlur('email')"
      />
      <InputField
        label="Street Address"
        placeholder="Street Address"
        :modelValue="formState.streetAddress"
        @update:modelValue="(val) => (formState.streetAddress = sanitizeAddress(val))"
        :hasError="(touched.streetAddress && !formState.streetAddress) || (hasAttemptedSubmit && !formState.streetAddress)"
        @blur="handleBlur('streetAddress')"
      />
      <InputField
        label="City"
        placeholder="City"
        :modelValue="formState.city"
        @update:modelValue="(val) => (formState.city = sanitizeCity(val))"
        :hasError="(touched.city && !formState.city) || (hasAttemptedSubmit && !formState.city)"
        @blur="handleBlur('city')"
      />
      <InputField
        label="State"
        placeholder="State"
        :modelValue="formState.state"
        @update:modelValue="(val) => (formState.state = sanitizeCity(val))"
        :hasError="(touched.state && !formState.state) || (hasAttemptedSubmit && !formState.state)"
        @blur="handleBlur('state')"
      />
      <InputField
        label="Zip Code"
        placeholder="Zip Code"
        :modelValue="formState.zipCode"
        @update:modelValue="(val) => (formState.zipCode = sanitizeZip(val))"
        maxlength="5"
        :hasError="(touched.zipCode && !formState.zipCode) || (hasAttemptedSubmit && !formState.zipCode)"
        @blur="handleBlur('zipCode')"
      />

      <InputTextArea
        class="full-width"
        :label="`When do you need to surrender your ${selectedAnimal.toLowerCase()}`"
        :placeholder="`When do you need to surrender your ${selectedAnimal.toLowerCase()}`"
        :modelValue="formState.whenToSurrenderAnimal"
        @update:modelValue="(val) => (formState.whenToSurrenderAnimal = val)"
        :hasError="(touched.whenToSurrenderAnimal && !formState.whenToSurrenderAnimal) || (hasAttemptedSubmit && !formState.whenToSurrenderAnimal)"
        @blur="handleBlur('whenToSurrenderAnimal')"
      />
      <InputField
        :label="`${selectedAnimal}'s Name`"
        :placeholder="`${selectedAnimal}'s Name`"
        :modelValue="formState.animalName"
        @update:modelValue="(val) => (formState.animalName = val as string)"
        :hasError="(touched.animalName && !formState.animalName) || (hasAttemptedSubmit && !formState.animalName)"
        @blur="handleBlur('animalName')"
      />
      <InputField
        label="Age"
        placeholder="Age"
        :modelValue="formState.animalAge"
        @update:modelValue="(val) => (formState.animalAge = val as string)"
        :hasError="(touched.animalAge && !formState.animalAge) || (hasAttemptedSubmit && !formState.animalAge)"
        @blur="handleBlur('animalAge')"
      />
      <InputSelectGroup
        label="Sex"
        :options="['Male', 'Female', 'Unknown']"
        :modelValue="formState.animalSex"
        @update:modelValue="(val) => (formState.animalSex = val as string)"
        :hasError="(touched.animalSex && !formState.animalSex) || (hasAttemptedSubmit && !formState.animalSex)"
        @blur="handleBlur('animalSex')"
      />

      <InputTextArea
        class="full-width"
        :label="`How long have you had your ${selectedAnimal.toLowerCase()}?`"
        :placeholder="`How long have you had your ${selectedAnimal.toLowerCase()}?`"
        :modelValue="formState.animalOwnershipDuration"
        @update:modelValue="(val) => (formState.animalOwnershipDuration = val)"
        :hasError="(touched.animalOwnershipDuration && !formState.animalOwnershipDuration) || (hasAttemptedSubmit && !formState.animalOwnershipDuration)"
        @blur="handleBlur('animalOwnershipDuration')"
      />
      <InputTextArea
        class="full-width"
        :label="`Where did you get your ${selectedAnimal.toLowerCase()}?`"
        :placeholder="`Where did you get your ${selectedAnimal.toLowerCase()}?`"
        :modelValue="formState.animalLocationFound"
        @update:modelValue="(val) => (formState.animalLocationFound = val)"
        :hasError="(touched.animalLocationFound && !formState.animalLocationFound) || (hasAttemptedSubmit && !formState.animalLocationFound)"
        @blur="handleBlur('animalLocationFound')"
      />
      <InputTextArea
        class="full-width"
        :label="`Why are you surrendering your ${selectedAnimal.toLowerCase()}?`"
        :placeholder="`Why are you surrendering your ${selectedAnimal.toLowerCase()}?`"
        :modelValue="formState.animalWhySurrendered"
        @update:modelValue="(val) => (formState.animalWhySurrendered = val)"
        :hasError="(touched.animalWhySurrendered && !formState.animalWhySurrendered) || (hasAttemptedSubmit && !formState.animalWhySurrendered)"
        @blur="handleBlur('animalWhySurrendered')"
      />
    </fieldset>
    <div class="household-members-section">
      <h5>Including yourself, how many people live in your home?</h5>
      <p class="subtitle">Please list the age and gender of each person.</p>

      <div v-for="(member, index) in formState.householdMembers" :key="index" class="member-row">
        <div class="field-group gender-group">
          <ButtonToggle
            label="Gender"
            :modelValue="member.gender"
            @update:modelValue="(val) => (member.gender = val as 'Male' | 'Female')"
            true-value="Male"
            false-value="Female"
          />
        </div>

        <div class="field-group age-group">
          <InputField
            label="Age"
            placeholder="Age"
            :modelValue="member.age"
            @update:modelValue="(val) => (member.age = String(val).replace(/\D/g, '').substring(0, 3))"
            maxlength="3"
            class="clean-input"
            :hasError="(touched[`householdMembers[${index}].age`] || hasAttemptedSubmit) && !member.age"
            @blur="handleBlur(`householdMembers[${index}].age`)"
          />
        </div>

        <div class="field-group quantity-group">
          <InputField
            label="Quantity"
            placeholder="Qty"
            type="number"
            :modelValue="member.count"
            @update:modelValue="(val) => (member.count = Number(val))"
            min="1"
            class="clean-input"
            :hasError="(touched[`householdMembers[${index}].count`] || hasAttemptedSubmit) && (!member.count || member.count < 1)"
            @blur="handleBlur(`householdMembers[${index}].count`)"
          />
        </div>

        <button
          v-if="formState.householdMembers.length > 1"
          type="button"
          class="remove-btn"
          @click="formState.householdMembers.splice(index, 1)"
          aria-label="Remove member"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2-2h4a2 2 0 0 1 2-2h4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
        </button>
      </div>

      <button type="button" class="add-btn" @click="formState.householdMembers.push({ age: '', gender: 'Female', count: 1 })">
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        Add Another Person / Group
      </button>
    </div>
    <section class="full-width">
      <InputSelectGroup
        label="What other animals did the cat live with?"
        :options="['Dogs', 'Cats', 'Other', 'No other animals']"
        :modelValue="formState.otherPetsInHousehold"
        @update:modelValue="(val) => (formState.otherPetsInHousehold = val as string)"
        :hasError="(touched.otherPetsInHousehold && !formState.otherPetsInHousehold) || (hasAttemptedSubmit && !formState.otherPetsInHousehold)"
        @blur="handleBlur('otherPetsInHousehold')"
      />
    </section>
  </div>
</template>

<style scoped lang="css">
.household-section {
  display: flex;
  flex-direction: column;
  gap: 16px;

  .household-grid {
    border: 0;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 16px;

    @media (min-width: 768px) {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }

    .full-width {
      grid-column: 1 / -1;
    }
  }

  fieldset.field {
    border: 0;
    padding: 0;
    margin: 0;
  }

  .label {
    margin-bottom: 8px;
    font-weight: 600;
  }



  .household-members-section {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 16px;
    margin-bottom: 24px;

    .subtitle {
      color: var(--font-color-dark);
      margin-top: -8px;
      margin-bottom: 8px;
    }

    .member-row {
      display: flex;
      align-items: flex-start;
      gap: 16px;
      background: #f8fafc;
      padding: 16px;
      border-radius: 12px;
      border: 1px solid #e2e8f0;

      @media (max-width: 640px) {
        flex-direction: column;
        align-items: stretch;
      }
    }

    .field-group {
      display: flex;
      flex-direction: column;
      gap: 8px;
    }

    .gender-group,
    .age-group,
    .quantity-group {
      flex: 1;
      width: 0; /* Ensures flex-basis work correctly for equal widths */

      @media (max-width: 640px) {
        width: 100%;
        flex: none;
      }
    }

    .gender-group .label {
      @media (max-width: 640px) {
         text-align: center;
      }
    }

    .field-label {
      font-size: 0.875rem;
      font-weight: 600;
      color: #374151;
    }

    .gender-toggle {
      display: flex;
      background: #fff;
      border: 1px solid #e2e8f0;
      border-radius: 8px;
      padding: 4px;
      height: 48px; /* Match typical input height */

      .toggle-btn {
        flex: 1;
        border: none;
        background: transparent;
        border-radius: 6px;
        font-weight: 500;
        color: #64748b;
        cursor: pointer;
        padding: 0 4px;
        transition: all 0.2s;

        &.active {
          background: color-mix(in srgb, var(--green) 10%, white);
          border: 1px solid var(--green);
          box-shadow: 0 0 0 1px var(--green) inset;
          color: var(--font-color-dark);
          font-weight: 600;
        }
      }
    }

    .clean-input {
      margin-bottom: 0 !important;
    }

    .remove-btn {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 48px;
      height: 48px;
      border: 1px solid #fee2e2;
      background: #fff;
      color: #ef4444;
      border-radius: 8px;
      cursor: pointer;
      transition: all 0.2s;
      flex-shrink: 0;
      margin-top: 29px; /* Align with input fields */

      &:hover {
        background: #fef2f2;
      }

      @media (max-width: 640px) {
        margin-top: 0;
        width: 100%;
        height: 40px;
      }
    }

    .add-btn {
      align-self: flex-start;
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 12px 20px;
      background: #fff;
      border: 1px dashed #cbd5e1;
      border-radius: 8px;
      color: #64748b;
      font-weight: 500;
      cursor: pointer;
      transition: all 0.2s;

      &:hover {
        border: 1px solid var(--green);
        background: color-mix(in srgb, var(--green) 10%, white);
        color: var(--font-color-dark);
        box-shadow: 0 0 0 1px var(--green) inset;
        font-weight: 600;
      }
    }
  }
}
/* Add error styles for fieldsets */
fieldset.has-error .chips {
  outline: 2px solid #ef4444;
  border-color: #ef4444;
  border-radius: 12px; /* Add some radius */
  padding: 8px; /* Add padding so border doesn't hug chips too tight */
}
</style>
