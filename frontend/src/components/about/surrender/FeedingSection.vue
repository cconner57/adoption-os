<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Feeding</h5>
    <InputField
      label="What type of food does the cat eat?"
      placeholder="Answer"
      :modelValue="formState.catTypeOfFood"
      @update:modelValue="(val) => (formState.catTypeOfFood = val)"
    />
    <InputField
      label="How many times a day is pet fed?"
      placeholder="Answer"
      :modelValue="formState.catEatingFrequency"
      @update:modelValue="(val) => (formState.catEatingFrequency = val)"
    />
    <InputField
      label="How much is fed per feeding?"
      placeholder="Answer"
      :modelValue="formState.catAmountOfFood"
      @update:modelValue="(val) => (formState.catAmountOfFood = val)"
    />
    <InputField
      label="Does the pet get treats?"
      placeholder="Answer"
      :modelValue="formState.catFoodTreats"
      @update:modelValue="(val) => (formState.catFoodTreats = val)"
    />
    <InputField
      label="If yes, what treats does the pet like?"
      placeholder="Answer"
      :modelValue="formState.catFoodTreatsExplanation"
      @update:modelValue="(val) => (formState.catFoodTreatsExplanation = val)"
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
