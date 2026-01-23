<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../../models/common'
import { ButtonToggle, Combobox, InputSelectGroup, InputTextArea } from '../../../common/ui'

const props = defineProps<{
  modelValue: Partial<IPet>
  availablePets?: IPet[]
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const behaviorOptions = [
  { label: 'Good with Kids', value: 'kids' },
  { label: 'Good with Dogs', value: 'dogs' },
  { label: 'Good with Cats', value: 'cats' },
  { label: 'House Trained', value: 'trained' },
]

const behaviorFlags = computed({
  get: () => {
    const list: string[] = []
    if (formData.value.behavior?.isGoodWithKids) list.push('kids')
    if (formData.value.behavior?.isGoodWithDogs) list.push('dogs')
    if (formData.value.behavior?.isGoodWithCats) list.push('cats')
    if (formData.value.behavior?.isHouseTrained) list.push('trained')
    return list
  },
  set: (vals: string[] | null) => {
    if (!formData.value.behavior) return
    const v = vals || []
    formData.value.behavior.isGoodWithKids = v.includes('kids')
    formData.value.behavior.isGoodWithDogs = v.includes('dogs')
    formData.value.behavior.isGoodWithCats = v.includes('cats')
    formData.value.behavior.isHouseTrained = v.includes('trained')
  },
})

const temperamentOptions = [
  { label: 'Affectionate', value: 'affectionate' },
  { label: 'Anxious', value: 'anxious' },
  { label: 'Bossy', value: 'bossy' },
  { label: 'Curious', value: 'curious' },
  { label: 'Hunter', value: 'hunter' },
  { label: 'Independent', value: 'independent' },
  { label: 'Laid-back', value: 'laid-back' },
  { label: 'Playful', value: 'playful' },
  { label: 'Shy', value: 'shy' },
  { label: 'Vocal', value: 'vocal' },
]

const placementOptions = [
  { label: 'Must go with another cat', value: 'must_cat' },
  { label: 'Must go with another dog', value: 'must_dog' },
  { label: 'Prefers to be alone', value: 'alone' },
]

const placementRequirements = computed({
  get: () => {
    const list: string[] = []
    if (formData.value.behavior?.mustGoWithAnotherCat) list.push('must_cat')
    if (formData.value.behavior?.mustGoWithAnotherDog) list.push('must_dog')
    if (formData.value.behavior?.prefersToBeAlone) list.push('alone')
    return list
  },
  set: (vals: string[] | null) => {
    if (!formData.value.behavior) return
    const v = vals || []
    formData.value.behavior.mustGoWithAnotherCat = v.includes('must_cat')
    formData.value.behavior.mustGoWithAnotherDog = v.includes('must_dog')
    formData.value.behavior.prefersToBeAlone = v.includes('alone')
  },
})

const isBondedComputed = computed({
  get: () => formData.value.behavior?.bonded?.isBonded ?? false,
  set: (val: boolean) => {
    if (!formData.value.behavior) return
    if (!formData.value.behavior.bonded) {
      formData.value.behavior.bonded = { isBonded: false, bondedWith: [] }
    }
    formData.value.behavior.bonded.isBonded = val
  },
})

const bondedWithOptions = computed(() => {
  const all = props.availablePets || []
  
  const currentId = props.modelValue.id
  return all
    .filter((p) => p.id !== currentId && p.details?.status !== 'archived')
    .map((p) => ({ label: p.name, value: p.name }))
})
</script>

<style scoped>
@import './form.css';
</style>

<template>
  <div class="form-section" v-if="formData.behavior">
    <div class="form-group">
      <InputSelectGroup
        label="Traits"
        v-model="behaviorFlags"
        :options="behaviorOptions"
        multiple
      />
    </div>

    <div class="form-group">
      <InputSelectGroup
        label="Personality"
        v-model="formData.behavior.personalityTags"
        :options="temperamentOptions"
        multiple
      />
    </div>

    <div class="form-group">
      <InputSelectGroup
        label="Placement Requirements"
        v-model="placementRequirements"
        :options="placementOptions"
        multiple
      />
    </div>

    <div class="form-group checkbox-list">
      <ButtonToggle
        label="Part of a Bonded Pair?"
        v-model="isBondedComputed"
        :true-value="true"
        :false-value="false"
      />

      <div
        v-if="isBondedComputed && formData.behavior.bonded"
        style="margin-left: 12px; margin-top: 8px"
      >
        <Combobox
          label="Bonded With"
          :model-value="formData.behavior.bonded.bondedWith?.length ? formData.behavior.bonded.bondedWith : null"
          @update:model-value="val => formData.behavior?.bonded && (formData.behavior.bonded.bondedWith = (Array.isArray(val) ? val : val ? [val] : []))"
          :options="bondedWithOptions"
          multiple
          placeholder="Search pets..."
        />
      </div>
    </div>

    <div class="form-group">
      <InputTextArea
        label="Special Needs / Behavior Notes"
        :model-value="formData.behavior.specialNeeds || null"
        @update:model-value="val => formData.behavior && (formData.behavior.specialNeeds = val)"
        placeholder="Describe any special behavioral needs..."
        :maxChars="300"
      />
    </div>
  </div>
</template>
