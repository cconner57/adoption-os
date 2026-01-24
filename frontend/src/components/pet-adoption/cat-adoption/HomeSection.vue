<script setup lang="ts">
import { reactive, ref } from 'vue'

import ButtonToggle from '../../common/ui/ButtonToggle.vue'
import InputField from '../../common/ui/InputField.vue'
import InputSelectGroup from '../../common/ui/InputSelectGroup.vue'

const { modelValue } = defineProps<{
  modelValue: {
    homeType: string | null
    homeOwnership: string | null
    landlordName: string | null
    landlordPhoneNumber: string | null
    allowPets: string | null
    breedRestrictionsWeightLimit: string | null
    monthlyFee: string | null
    allergies: string | null
    primaryOwner: boolean | null
    yearsAtAddress: string | null
    previousAddress: string | null
    expectToMove: string | null
    travelPlan: string | null
    catAccess: string | null
    catIndoorOutdoor: string | null
  }
  touched?: Record<string, boolean>
  // eslint-disable-next-line no-unused-vars
  handleBlur: (_field: string) => void
}>()

const touched = reactive<Record<string, boolean>>({})
const hasAttemptedSubmit = ref(false)

const handleBlur = (field: string) => {
  touched[field] = true
}
</script>

<template>
  <div class="home-section">
    <InputSelectGroup
      label="Home Type"
      :options="['Home', 'Apartment', 'Condo', 'Townhouse', 'Other']"
      :modelValue="modelValue.homeType"
      @update:modelValue="(val) => (modelValue.homeType = val as string)"
      :hasError="
        (touched.homeType && !modelValue.homeType) || (hasAttemptedSubmit && !modelValue.homeType)
      "
      @blur="handleBlur('homeType')"
    />
    <InputSelectGroup
      label="Do you own or rent your home?"
      :options="['Own', 'Rent', 'Lease']"
      :modelValue="modelValue.homeOwnership"
      @update:modelValue="(val) => (modelValue.homeOwnership = val as string)"
      :hasError="
        (touched.homeOwnership && !modelValue.homeOwnership) ||
        (hasAttemptedSubmit && !modelValue.homeOwnership)
      "
      @blur="handleBlur('homeOwnership')"
    />
    <template v-if="modelValue.homeOwnership === 'Rent' || modelValue.homeOwnership === 'Lease'">
      <InputField
        v-model="modelValue.landlordName"
        label="Landlord's Name"
        name="landlordName"
        placeholder="Enter landlord's name"
        :hasError="touched.landlordName && !modelValue.landlordName"
        @blur="handleBlur('landlordName')"
      />
      <InputField
        v-model="modelValue.landlordPhoneNumber"
        label="Landlord's Phone Number"
        name="landlordPhoneNumber"
        placeholder="(555) 555-5555"
        :hasError="touched.landlordPhoneNumber && !modelValue.landlordPhoneNumber"
        @blur="handleBlur('landlordPhoneNumber')"
      />
      <InputField
        v-model="modelValue.allowPets"
        label="Are pets allowed by your lease?"
        name="allowPets"
        placeholder="Yes/No, explain if needed"
        :hasError="touched.allowPets && !modelValue.allowPets"
        @blur="handleBlur('allowPets')"
      />
    </template>
    <InputField
      v-model="modelValue.breedRestrictionsWeightLimit"
      label="Are there any breed restrictions or weight limits?"
      name="breedRestrictions"
      placeholder="List any restrictions"
      :hasError="touched.breedRestrictionsWeightLimit && !modelValue.breedRestrictionsWeightLimit"
      @blur="handleBlur('breedRestrictionsWeightLimit')"
    />
    <InputField
      v-model="modelValue.monthlyFee"
      label="Is there a pet deposit or monthly fee?"
      name="monthlyFee"
      placeholder="Yes/No (Amount)"
      :hasError="touched.monthlyFee && !modelValue.monthlyFee"
      @blur="handleBlur('monthlyFee')"
    />
    <InputField
      v-model="modelValue.allergies"
      label="Does anyone in the household have cat allergies?"
      name="allergies"
      placeholder="Yes/No"
      :hasError="touched.allergies && !modelValue.allergies"
      @blur="handleBlur('allergies')"
    />
    <ButtonToggle
      label="Is the person filling out this application the primary owner/leaseholder/renter?"
      :modelValue="modelValue.primaryOwner"
      @update:modelValue="(val) => (modelValue.primaryOwner = val as boolean)"
      true-value="Yes"
      false-value="No"
    />
    <InputField
      v-model="modelValue.yearsAtAddress"
      label="How long have you lived at this address?"
      name="yearsAtAddress"
      placeholder="Years at Address"
      :hasError="touched.yearsAtAddress && !modelValue.yearsAtAddress"
      @blur="handleBlur('yearsAtAddress')"
    />
    <InputField
      v-model="modelValue.previousAddress"
      label="What city/state were you in before this and how long did you live there?"
      name="previousAddress"
      placeholder="Previous Address"
      :hasError="touched.previousAddress && !modelValue.previousAddress"
      @blur="handleBlur('previousAddress')"
    />
    <InputField
      v-model="modelValue.expectToMove"
      label="Do you expect to move in the next few months? When/Where?"
      name="expectToMove"
      placeholder="Expect to Move"
      :hasError="touched.expectToMove && !modelValue.expectToMove"
      @blur="handleBlur('expectToMove')"
    />
    <InputField
      v-model="modelValue.travelPlan"
      label="Do you travel a great deal? What do you plan to do with your cat when you do travel?"
      name="travelPlan"
      placeholder="Travel Plan"
      :hasError="touched.travelPlan && !modelValue.travelPlan"
      @blur="handleBlur('travelPlan')"
    />
    <InputSelectGroup
      label="Check all that apply. Will the cat have access to:"
      :options="['Balcony', 'Patio', 'Garage', 'Yard', 'Doggie Door', 'None of these']"
      :modelValue="modelValue.catAccess"
      @update:modelValue="(val) => (modelValue.catAccess = val as string)"
      :hasError="
        (touched.catAccess && !modelValue.catAccess) ||
        (hasAttemptedSubmit && !modelValue.catAccess)
      "
      @blur="handleBlur('catAccess')"
    />
    <InputSelectGroup
      label="This cat will be:"
      :options="['Indoor Only', 'Mostly Indoor', 'Mostly Outdoor', 'Outdoor Only']"
      :modelValue="modelValue.catIndoorOutdoor"
      @update:modelValue="(val) => (modelValue.catIndoorOutdoor = val as string)"
      :hasError="
        (touched.catIndoorOutdoor && !modelValue.catIndoorOutdoor) ||
        (hasAttemptedSubmit && !modelValue.catIndoorOutdoor)
      "
      @blur="handleBlur('catIndoorOutdoor')"
    />
  </div>
</template>

<style scoped lang="css">
.home-section {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

@media (width >= 768px) {
  .home-section {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 24px; 
    align-items: start;
  }
}
</style>
