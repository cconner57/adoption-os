<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Medical History</h5>
    <InputField
      label="Please list all veterniarians that have seen the cat, including address and number"
      placeholder="Answer"
      :modelValue="formState.catVeterinarianList"
      @update:modelValue="(val) => (formState.catVeterinarianList = val)"
    />
    <InputField
      label="Does the cat see a veterinarian at least once a year?"
      placeholder="Answer"
      :modelValue="formState.catVeterinarianYearlyVisits"
      @update:modelValue="(val) => (formState.catVeterinarianYearlyVisits = val)"
    />
    <InputField
      label="Is the cat spayed/neutered?"
      placeholder="Answer"
      :modelValue="formState.catSpayedNeutered"
      @update:modelValue="(val) => (formState.catSpayedNeutered = val)"
    />
    <InputField
      label="What vaccinations has the cat received?"
      placeholder="Answer"
      :modelValue="formState.catVaccinationHistory"
      @update:modelValue="(val) => (formState.catVaccinationHistory = val)"
    />
    <InputField
      label="Are vaccinations current?"
      placeholder="Answer"
      :modelValue="formState.catVaccinationsCurrent"
      @update:modelValue="(val) => (formState.catVaccinationsCurrent = val)"
    />
    <InputField
      label="Has the cat been tested for heartworm?"
      placeholder="Answer"
      :modelValue="formState.catTestedHeartworm"
      @update:modelValue="(val) => (formState.catTestedHeartworm = val)"
    />
    <InputField
      label="If yes, what were the test results"
      placeholder="Answer"
      :modelValue="formState.catTestedHeartwormExplanation"
      @update:modelValue="(val) => (formState.catTestedHeartwormExplanation = val)"
    />
    <InputField
      label="Is the pet taking heartworm preventative?"
      placeholder="Answer"
      :modelValue="formState.catHeartwormPrevention"
      @update:modelValue="(val) => (formState.catHeartwormPrevention = val)"
    />
    <InputField
      label="If yes, when is the next dosage due?"
      placeholder="Answer"
      :modelValue="formState.catHeartwormPreventionExplanation"
      @update:modelValue="(val) => (formState.catHeartwormPreventionExplanation = val)"
    />
    <InputField
      label="Is the pet microchipped?"
      placeholder="Answer"
      :modelValue="formState.catMicrochipped"
      @update:modelValue="(val) => (formState.catMicrochipped = val)"
    />
    <InputField
      label="If yes, please provide microchip number"
      placeholder="Answer"
      :modelValue="formState.catMicrochippedExplanation"
      @update:modelValue="(val) => (formState.catMicrochippedExplanation = val)"
    />
    <InputField
      label="Check if the cat has ever shown any of the following behaviors when handled by a vet or groomer"
      placeholder="Answer"
      :modelValue="formState.catVetOrGroomerBehavior"
      @update:modelValue="(val) => (formState.catVetOrGroomerBehavior = val)"
    />
    <InputField
      label="Does the cat have to be muzzled at the veterinarians?"
      placeholder="Answer"
      :modelValue="formState.catVetMuzzled"
      @update:modelValue="(val) => (formState.catVetMuzzled = val)"
    />
    <InputField
      label="Does the pet have any past or present health problems"
      placeholder="Answer"
      :modelValue="formState.catPastOrPresentHealthProblems"
      @update:modelValue="(val) => (formState.catPastOrPresentHealthProblems = val)"
    />
    <InputField
      label="If yes, please describe"
      placeholder="Answer"
      :modelValue="formState.catPastOrPresentHealthProblemsExplanation"
      @update:modelValue="(val) => (formState.catPastOrPresentHealthProblemsExplanation = val)"
    />
    <InputField
      label="Is the cat currently taking any medications?"
      placeholder="Answer"
      :modelValue="formState.catCurrentMedications"
      @update:modelValue="(val) => (formState.catCurrentMedications = val)"
    />
    <InputField
      label="If yes, please list medications"
      placeholder="Answer"
      :modelValue="formState.catCurrentMedicationsExplanation"
      @update:modelValue="(val) => (formState.catCurrentMedicationsExplanation = val)"
    />
  </div>
</template>

<style scoped lang="css">
h5 {
  margin-bottom: 24px;
}
div > div {
    margin-bottom: 16px;
}
</style>
