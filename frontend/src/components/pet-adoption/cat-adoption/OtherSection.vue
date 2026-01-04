<script setup lang="ts">
import InputSelectGroup from '../../common/ui/InputSelectGroup.vue'
import InputField from '../../common/ui/InputField.vue'

import type { FormState } from '../../../models/adopt-form.ts'

const { modelValue } = defineProps<{
  modelValue: FormState
  touched?: Record<string, boolean>
  handleBlur?: (field: string) => void
}>()
</script>

<template>
  <div class="other-section">
    <InputField
      v-model="modelValue.bredAnimalDescription"
      label="If you have ever bred an animal, please describe the circumstances"
      name="bredAnimalDescription"
      placeholder="Describe the circumstances"
      :hasError="touched?.bredAnimalDescription && !modelValue.bredAnimalDescription"
      @blur="handleBlur?.('bredAnimalDescription')"
    />
    <InputField
      v-model="modelValue.ownedDeclawedOrDebarked"
      label="Have you ever owned a declawed cat or a debarked dog?"
      name="ownedDeclawedOrDebarked"
      placeholder="Yes/No and details"
      :hasError="touched?.ownedDeclawedOrDebarked && !modelValue.ownedDeclawedOrDebarked"
      @blur="handleBlur?.('ownedDeclawedOrDebarked')"
    />
    <InputField
      v-model="modelValue.movedWithPet"
      label="Have you ever moved with a pet?"
      name="movedWithPet"
      placeholder="Yes/No and details"
      :hasError="touched?.movedWithPet && !modelValue.movedWithPet"
      @blur="handleBlur?.('movedWithPet')"
    />
    <InputField
      v-model="modelValue.ownedSpecialNeedsPet"
      label="Have you ever owned a special needs pet?"
      name="ownedSpecialNeedsPet"
      placeholder="Yes/No and details"
      :hasError="touched?.ownedSpecialNeedsPet && !modelValue.ownedSpecialNeedsPet"
      @blur="handleBlur?.('ownedSpecialNeedsPet')"
    />
    <InputField
      v-model="modelValue.mobilityDevice"
      label="Does anyone in your home use a mobility device?"
      name="mobilityDevice"
      placeholder="Yes/No and details"
      :hasError="touched?.mobilityDevice && !modelValue.mobilityDevice"
      @blur="handleBlur?.('mobilityDevice')"
    />
    <InputSelectGroup
      label="Check all that apply. Under what conditions would you NOT KEEP your new pet?"
      :options="[
        'Moved to a place that didn\'t allow pets',
        'Move out of state',
        'Sheds too much',
        'Litterbox problem',
        'Kids ignored pet',
        'Required prescription food or medication',
        'Cat has seizures',
        'Cat whines/cries/plays loudly at night',
        'Cat chews clothing/shoes',
        'New baby',
        'Allergy',
        'Pets didn\'t get along',
        'Scratches/chews furniture',
        'Pees on furniture',
        'Bit or snapped at another animal',
        'Bit or snapped at a person',
        'Escapes from yard/house',
        'Divorce/Separation',
        'Unemployment or financial hardship',
        'None of the above',
      ]"
      multiple
      :modelValue="modelValue.surrenderConditions"
      @update:modelValue="(val) => (modelValue.surrenderConditions = val as string[])"
      :hasError="touched?.surrenderConditions && modelValue.surrenderConditions.length === 0"
      @blur="handleBlur?.('surrenderConditions')"
    />
    <InputField
      v-model="modelValue.surrenderPlan"
      label="What would you do with your cat if you could not keep it for the above reason(s)?"
      name="surrenderPlan"
      placeholder="Details"
      :hasError="touched?.surrenderPlan && !modelValue.surrenderPlan"
      @blur="handleBlur?.('surrenderPlan')"
    />
    <InputField
      v-model="modelValue.foodTypeBrand"
      label="What type and brand of food do you plan on feeding your new cat?"
      name="foodTypeBrand"
      placeholder="Type and Brand"
      :hasError="touched?.foodTypeBrand && !modelValue.foodTypeBrand"
      @blur="handleBlur?.('foodTypeBrand')"
    />
    <InputField
      v-model="modelValue.affordVetCare"
      label="Can you afford regular veterinary care for this cat - including yearly vaccinations, annual physical exams, dental care, etc. ($300 or more per year)?"
      name="affordVetCare"
      placeholder="Yes/No"
      :hasError="touched?.affordVetCare && !modelValue.affordVetCare"
      @blur="handleBlur?.('affordVetCare')"
    />
    <InputField
      v-model="modelValue.affordEmergencyCost"
      label="Can you afford serious injury or illness costs ($1000 or more)?"
      name="affordEmergencyCost"
      placeholder="Yes/No"
      :hasError="touched?.affordEmergencyCost && !modelValue.affordEmergencyCost"
      @blur="handleBlur?.('affordEmergencyCost')"
    />
  </div>
</template>

<style scoped lang="css"></style>
