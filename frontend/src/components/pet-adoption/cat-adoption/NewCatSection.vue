<script setup lang="ts">
import InputSelectGroup from '../../common/ui/InputSelectGroup.vue'
import InputField from '../../common/ui/InputField.vue'
import InputTextArea from '../../common/ui/InputTextArea.vue'

const { modelValue } = defineProps<{
  modelValue: {
    catPreferenceBreed: string | null
    catPreferencePhysical: string | null
    catPreferencePersonality: string | null
    catPreferenceNotWant: string | null
    whyInterested: string | null
    adoptionReason: string | null
    ownCatBefore: string | null
    ownKittenBefore: string | null
    alreadyHaveVeterinarian: string | null
    catAllowedHomeArea: string | null
    catHomeAloneHours: string | null
    catDisciplineType: string | null
    catEscapeSteps: string | null
  }
  touched?: Record<string, boolean>
  handleBlur?: (field: string) => void
}>()
</script>

<template>
  <div class="new-cat-section">
    <InputField
      v-model="modelValue.catPreferenceBreed"
      label="Preferred breed, age, or gender?"
      name="catPreferenceBreed"
      placeholder="e.g. Siamese, Kitten, Female"
      required
      :hasError="touched?.catPreferenceBreed && !modelValue.catPreferenceBreed"
      @blur="handleBlur?.('catPreferenceBreed')"
    />
    <InputField
      v-model="modelValue.catPreferencePhysical"
      label="Preferred size, hair length, or color?"
      name="catPreferencePhysical"
      placeholder="e.g. Short hair, Orange tabby"
      required
      :hasError="touched?.catPreferencePhysical && !modelValue.catPreferencePhysical"
      @blur="handleBlur?.('catPreferencePhysical')"
    />
    <InputField
      v-model="modelValue.catPreferencePersonality"
      label="Preferred personality or energy level?"
      name="catPreferencePersonality"
      placeholder="e.g. Cuddly, Playful, Chill"
      required
      :hasError="touched?.catPreferencePersonality && !modelValue.catPreferencePersonality"
      @blur="handleBlur?.('catPreferencePersonality')"
    />
    <InputField
      v-model="modelValue.catPreferenceNotWant"
      label="Any traits you specifically NOT want?"
      name="catPreferenceNotWant"
      placeholder="e.g. Aggressive, Super high energy"
      required
      :hasError="touched?.catPreferenceNotWant && !modelValue.catPreferenceNotWant"
      @blur="handleBlur?.('catPreferenceNotWant')"
    />
    <InputTextArea
      label="Why are you interested in adopting a new cat?"
      placeholder="Share your motivation..."
      :modelValue="modelValue.whyInterested"
      @update:modelValue="(val) => (modelValue.whyInterested = val)"
      :hasError="touched?.whyInterested && !modelValue.whyInterested"
      @blur="handleBlur?.('whyInterested')"
      :spanFull="false"
    />
    <div class="spacer desktop-only"></div>
    <InputSelectGroup
      label="Primary reason for adoption:"
      :options="[
        'Companion for me',
        'Companion for family/children',
        'Companion for another pet',
        'For protection',
        'A gift',
      ]"
      :modelValue="modelValue.adoptionReason"
      @update:modelValue="(val) => (modelValue.adoptionReason = val as string)"
      :hasError="touched?.adoptionReason && !modelValue.adoptionReason"
      @blur="handleBlur?.('adoptionReason')"
    />
    <InputSelectGroup
      label="Have you owned a cat before?"
      :options="['Yes', 'No', 'Not as an adult']"
      :modelValue="modelValue.ownCatBefore"
      @update:modelValue="(val) => (modelValue.ownCatBefore = val as string)"
      :hasError="touched?.ownCatBefore && !modelValue.ownCatBefore"
      @blur="handleBlur?.('ownCatBefore')"
    />
    <InputSelectGroup
      label="Have you owned a kitten before?"
      :options="['Yes', 'No', 'Not as an adult']"
      :modelValue="modelValue.ownKittenBefore"
      @update:modelValue="(val) => (modelValue.ownKittenBefore = val as string)"
      :hasError="touched?.ownKittenBefore && !modelValue.ownKittenBefore"
      @blur="handleBlur?.('ownKittenBefore')"
    />
    <InputSelectGroup
      label="Do you currently have a veterinarian?"
      :options="['Yes', 'No', 'I need a recommendation']"
      :modelValue="modelValue.alreadyHaveVeterinarian"
      @update:modelValue="(val) => (modelValue.alreadyHaveVeterinarian = val as string)"
      :hasError="touched?.alreadyHaveVeterinarian && !modelValue.alreadyHaveVeterinarian"
      @blur="handleBlur?.('alreadyHaveVeterinarian')"
    />
    <InputTextArea
      label="Where in the house will the cat be allowed?"
      placeholder="e.g. Everywhere, Bedrooms only..."
      :modelValue="modelValue.catAllowedHomeArea"
      @update:modelValue="(val) => (modelValue.catAllowedHomeArea = val)"
      :hasError="touched?.catAllowedHomeArea && !modelValue.catAllowedHomeArea"
      :spanFull="false"
    />
    <InputTextArea
      label="How many hours a day will the cat be alone?"
      placeholder="Please include typical work schedule..."
      :modelValue="modelValue.catHomeAloneHours"
      @update:modelValue="(val) => (modelValue.catHomeAloneHours = val)"
      :hasError="touched?.catHomeAloneHours && !modelValue.catHomeAloneHours"
      :spanFull="false"
    />
    <InputTextArea
      label="How will you correct undesirable behavior?"
      placeholder="e.g. Redirecting, Spray bottle, Positive reinforcement..."
      :modelValue="modelValue.catDisciplineType"
      @update:modelValue="(val) => (modelValue.catDisciplineType = val)"
      :hasError="touched?.catDisciplineType && !modelValue.catDisciplineType"
      :spanFull="false"
    />
    <InputTextArea
      label="If the cat escapes, what will you do?"
      placeholder="Search plan..."
      :modelValue="modelValue.catEscapeSteps"
      @update:modelValue="(val) => (modelValue.catEscapeSteps = val)"
      :hasError="touched?.catEscapeSteps && !modelValue.catEscapeSteps"
      :spanFull="false"
    />
  </div>
</template>

<style scoped lang="css">
.new-cat-section {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

@media (max-width: 768px) {
  .new-cat-section {
    grid-template-columns: 1fr;
  }
  .desktop-only {
    display: none;
  }
}
.section-group {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--green);
  border-bottom: 2px solid var(--green);
  padding-bottom: 0.5rem;
  margin-bottom: 0.5rem;
}
</style>
