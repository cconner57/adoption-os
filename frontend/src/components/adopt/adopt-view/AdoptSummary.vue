<script setup lang="ts">
import type { IPet } from '../../../models/common.ts'
import { formatDate } from '../../../utils/common.ts'
import PetItem from '../../common/pet-item/PetItem.vue'

defineProps<{
  pets: IPet[]
}>()
</script>

<template>
  <div class="adopt-summary">
    <PetItem
      v-for="(pet, index) in pets"
      :capsules="[
        pet?.species ?? '',
        pet?.sex ?? '',
        pet?.physical?.dateOfBirth ? formatDate(pet?.physical?.dateOfBirth ?? '', true) : '',
      ]"
      :description="pet.descriptions?.fun ?? ''"
      :id="pet.slug || pet.id"
      :key="pet.id"
      :name="pet.name"
      :photo="pet.photos?.find((p) => p.isPrimary)?.url"
      :priority="index === 0"
    />
  </div>
</template>

<style scoped lang="css">
.adopt-summary {
  display: flex;
  flex-wrap: wrap;
  row-gap: 30px;
  column-gap: 20px;
  
  justify-content: center;

}
</style>
