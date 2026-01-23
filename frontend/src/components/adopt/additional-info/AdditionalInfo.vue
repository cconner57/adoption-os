<script setup lang="ts">
import { computed } from 'vue'

import type { IPet } from '../../../models/common.ts'

const props = defineProps<{
  pet: IPet
}>()

const isSpayedOrNeutered = (pet: IPet) => {
  return pet?.sex === 'male' ? 'Neutered' : 'Spayed'
}

const goodWithText = computed(() => {
  const traits = props.pet.behavior
  if (!traits || (!traits.isGoodWithCats && !traits.isGoodWithDogs && !traits.isGoodWithKids)) {
    return 'N/A'
  }

  const goodWith: string[] = []
  if (traits.isGoodWithCats) goodWith.push('Other Cats')
  if (traits.isGoodWithDogs) goodWith.push('Other Dogs')
  if (traits.isGoodWithKids) goodWith.push('Kids')

  return goodWith.join(', ')
})
const houseTrainedText = () => {
  if (props.pet.behavior?.isHouseTrained === undefined) {
    return 'N/A'
  }
  if (props.pet.behavior?.isHouseTrained) {
    return 'Yes';
  } else {
    return 'No'
  }
}
</script>

<template>
  <div class="adopt-detail__additional-info">
    <div class="adopt-detail__additional-info__item">
      <p>Breed</p>
      <p>{{ pet.physical?.breed ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Color</p>
      <p>{{ pet.physical?.color ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Size</p>
      <p>{{ pet.physical?.size ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>House-trained</p>
      <p>{{ houseTrainedText() }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Health</p>
      <p>
        {{ pet.medical?.vaccinationsUpToDate ? 'Vaccinated' : 'Not Vaccinated' }},
        {{
          pet.medical?.spayedOrNeutered
            ? isSpayedOrNeutered(pet)
            : `Not ${isSpayedOrNeutered(pet)}`
        }},
        {{ pet.medical?.microchip.microchipped ? 'Microchipped' : 'Not Microchipped' }}
      </p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Good in a home with</p>
      <p>{{ goodWithText }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Adoption Fee</p>
      <p>{{ pet?.adoption?.fee !== undefined ? '$' + pet?.adoption?.fee : 'N/A' }}</p>
    </div>
  </div>
</template>

<style scoped lang="css">
.adopt-detail__additional-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-height: 250px;

  @media (max-width: 440px) {
    gap: 5px;
    flex-direction: column;
    align-items: flex-start;
    p {
      font-size: 0.9rem;
      line-height: 1.5;
      text-wrap: wrap;
    }
    p:last-child {
      text-wrap: wrap;
      width: auto; /* Allow flexible width */
      flex: 1;    /* Take remaining space */
    }
  }
}

.adopt-detail__additional-info__item {
  display: flex;
  flex-direction: row;
  p {
    text-transform: capitalize;
    text-align: left;
  }
  & p:first-child {
    width: 200px;
    flex-shrink: 0;
    @media (max-width: 440px) {
      width: 110px;
      margin-right: 12px;
    }
  }
  & p:last-child {
    font-weight: bold;
    width: 300px;
    @media (max-width: 440px) {
      width: auto;
      flex: 1;
    }
  }
}
</style>
