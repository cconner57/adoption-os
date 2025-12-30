<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/surrender-form.ts'
import InputTextArea from '../../common/ui/InputTextArea.vue'

import {
  formatPhoneNumber,
  sanitizeName,
  sanitizeCity,
  sanitizeZip,
  sanitizeAddress,
} from '../../../utils/validators.ts'

const { formState, touched, handleBlur, hasAttemptedSubmit } = defineProps<{
  formState: SurrenderFormState
  touched: Record<string, boolean>
  handleBlur: (field: string) => void
  hasAttemptedSubmit: boolean
}>()
</script>

<template>
  <div class="household-section">
    <h5>Cat & Household Information</h5>
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
        label="When do you need to surrender your cat"
        placeholder="When do you need to surrender your cat"
        :modelValue="formState.whenToSurrenderCat"
        @update:modelValue="(val) => (formState.whenToSurrenderCat = val)"
        :hasError="(touched.whenToSurrenderCat && !formState.whenToSurrenderCat) || (hasAttemptedSubmit && !formState.whenToSurrenderCat)"
        @blur="handleBlur('whenToSurrenderCat')"
      />
      <InputField
        label="Cat's Name"
        placeholder="Cat's Name"
        :modelValue="formState.catName"
        @update:modelValue="(val) => (formState.catName = val)"
        :hasError="(touched.catName && !formState.catName) || (hasAttemptedSubmit && !formState.catName)"
        @blur="handleBlur('catName')"
      />
      <InputField
        label="Age"
        placeholder="Age"
        :modelValue="formState.catAge"
        @update:modelValue="(val) => (formState.catAge = val)"
        :hasError="(touched.catAge && !formState.catAge) || (hasAttemptedSubmit && !formState.catAge)"
        @blur="handleBlur('catAge')"
      />
      <fieldset
        class="field full-width"
        aria-labelledby="cat-sex-legend"
        :class="{ 'has-error': (touched.catSex && !formState.catSex) || (hasAttemptedSubmit && !formState.catSex) }"
      >
        <legend id="cat-sex-legend" class="label centralized-legend">Sex</legend>
        <fieldset class="chips" aria-labelledby="cat-sex-legend">
          <label class="chip">
            <input type="radio" value="Male" v-model="formState.catSex" @blur="handleBlur('catSex')" />
            <span>Male</span>
          </label>
          <label class="chip">
            <input type="radio" value="Female" v-model="formState.catSex" @blur="handleBlur('catSex')" />
            <span>Female</span>
          </label>
          <label class="chip">
            <input type="radio" value="Unknown" v-model="formState.catSex" @blur="handleBlur('catSex')" />
            <span>Unknown</span>
          </label>
        </fieldset>
      </fieldset>

      <InputTextArea
        class="full-width"
        label="How long have you had your cat?"
        placeholder="How long have you had your cat?"
        :modelValue="formState.catOwnershipDuration"
        @update:modelValue="(val) => (formState.catOwnershipDuration = val)"
        :hasError="(touched.catOwnershipDuration && !formState.catOwnershipDuration) || (hasAttemptedSubmit && !formState.catOwnershipDuration)"
        @blur="handleBlur('catOwnershipDuration')"
      />
      <InputTextArea
        class="full-width"
        label="Where did you get your cat?"
        placeholder="Where did you get your cat?"
        :modelValue="formState.catLocationFound"
        @update:modelValue="(val) => (formState.catLocationFound = val)"
        :hasError="(touched.catLocationFound && !formState.catLocationFound) || (hasAttemptedSubmit && !formState.catLocationFound)"
        @blur="handleBlur('catLocationFound')"
      />
      <InputTextArea
        class="full-width"
        label="Why are you surrendering your cat?"
        placeholder="Why are you surrendering your cat?"
        :modelValue="formState.catWhySurrendered"
        @update:modelValue="(val) => (formState.catWhySurrendered = val)"
        :hasError="(touched.catWhySurrendered && !formState.catWhySurrendered) || (hasAttemptedSubmit && !formState.catWhySurrendered)"
        @blur="handleBlur('catWhySurrendered')"
      />
    </fieldset>
    <div class="household-members-section">
      <h5>Including yourself, how many people live in your home?</h5>
      <p class="subtitle">Please list the age and gender of each person.</p>

      <div v-for="(member, index) in formState.householdMembers" :key="index" class="member-row">
        <div class="field-group gender-group">
          <label class="field-label">Gender</label>
          <div class="gender-toggle">
            <button
              type="button"
              class="toggle-btn"
              :class="{ active: member.gender === 'Female' }"
              @click="member.gender = 'Female'"
            >
              Female
            </button>
            <button
              type="button"
              class="toggle-btn"
              :class="{ active: member.gender === 'Male' }"
              @click="member.gender = 'Male'"
            >
              Male
            </button>
          </div>
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
    <section>
      <fieldset
        class="field col-span-2"
        aria-labelledby="other-animals-legend"
        :class="{ 'has-error': (touched.otherPetsInHousehold && !formState.otherPetsInHousehold) || (hasAttemptedSubmit && !formState.otherPetsInHousehold) }"
      >
        <legend id="other-animals-legend" class="label centralized-legend">
          What other animals did the cat live with?
        </legend>
        <fieldset class="chips" aria-labelledby="other-animals-legend">
          <label class="chip">
            <input type="radio" value="Dogs" v-model="formState.otherPetsInHousehold" @blur="handleBlur('otherPetsInHousehold')" />
            <span>Dogs</span>
          </label>
          <label class="chip">
            <input type="radio" value="Cats" v-model="formState.otherPetsInHousehold" @blur="handleBlur('otherPetsInHousehold')" />
            <span>Cats</span>
          </label>
          <label class="chip">
            <input type="radio" value="Other" v-model="formState.otherPetsInHousehold" @blur="handleBlur('otherPetsInHousehold')" />
            <span>Other</span>
          </label>
          <label class="chip">
            <input type="radio" value="No other animals" v-model="formState.otherPetsInHousehold" @blur="handleBlur('otherPetsInHousehold')" />
            <span>No other animals</span>
          </label>
        </fieldset>
      </fieldset>
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

  .centralized-legend {
    @media (max-width: 640px) {
      text-align: center;
      display: block;
      width: 100%;
    }
  }

  .chips {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    border: none;
    padding: 0;
    margin: 0;
    @media (max-width: 440px) {
      margin-top: 8px;
      flex-direction: column;
      text-align: center;
      gap: 8px;
      /* Center chips content */
      align-items: center;
    }
  }
  .chip {
    position: relative;
    display: inline-flex;
    align-items: center;
    gap: 8px;
    padding: 8px 12px;
    border-radius: 999px;
    border: 1px solid #e7ebf0;
    background-color: #fff;
    cursor: pointer;
    user-select: none;
    font-size: 1rem;
    transition:
      background 0.2s,
      border-color 0.2s,
      box-shadow 0.2s;

    @media (max-width: 440px) {
       width: 100%;
       justify-content: center;
    }

    span {
      font-weight: 600;
      color: var(--text-900);
      line-height: 1.5;
    }

    &:hover {
      border-color: #d7e2f2;
      background: #f2f7ff;
    }
  }
  .chip > input {
    position: absolute;
    opacity: 0;
    width: 1px;
    height: 1px;
    pointer-events: none;
  }
  .chip:has(> input:checked) {
    background: color-mix(in srgb, var(--green) 10%, white);
    border: 1px solid var(--green);
    box-shadow: 0 0 0 1px var(--green) inset;
    color: var(--font-color-dark);
  }
  .chip:has(> input:focus-visible) {
    box-shadow: 0 0 0 3px var(--ring);
  }

  @supports not (selector(:has(*))) {
    .chip > input:checked + span {
      background: #e8f1ff;
      border-radius: 999px;
      padding: 6px 10px;
      margin: -6px -10px;
      box-shadow: 0 0 0 2px #bfd0ff inset;
    }
    .chip > input:focus-visible + span {
      box-shadow: 0 0 0 3px var(--ring);
    }
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
