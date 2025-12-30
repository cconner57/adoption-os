<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Other</h5>
    <InputField
      label="Please feel free to tell us any other information about the pet you feel is important"
      placeholder="Answer"
      :modelValue="formState.additionalInformation"
      @update:modelValue="(val) => (formState.additionalInformation = val)"
    />
    <InputField
      label="Please upload a full body picture of the pet"
      placeholder="Answer"
      :modelValue="formState.fullBodyPhotoOfCat"
      @update:modelValue="(val) => (formState.fullBodyPhotoOfCat = val)"
    />
    <InputField
      label="Please upload a up close head shot of the pet"
      placeholder="Answer"
      :modelValue="formState.closeUpPhotoOfCatsFace"
      @update:modelValue="(val) => (formState.closeUpPhotoOfCatsFace = val)"
    />
    <InputField
      label="Please upload any records you have for the pet"
      placeholder="Answer"
      :modelValue="formState.copiesOfRecords"
      @update:modelValue="(val) => (formState.copiesOfRecords = val)"
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
