<script setup lang="ts">
import { defineProps } from 'vue'

import type { IPet } from '../../../models/common'
import { formatDate } from '../../../utils/common'
import { useIsMobile } from '../../../utils/useIsMobile'
import PetItem from '../../common/pet-item/PetItem.vue'

defineProps<{
  pet: IPet
}>()

const isMobile = useIsMobile()
</script>

<template>
  <div class="adopt-detail__more-friends">
    <h2>More Friends You Might Like</h2>
    <div class="adopt-detail__more-friends__list">
      <PetItem
        name="Crystal"
        :id="pet.id"
        :capsules="[
          pet.species || '',
          (pet.sex as string) || '',
          formatDate((pet.physical?.ageGroup || '') as string),
        ]"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || pet.photos?.[0]?.url"
      />
      <PetItem
        name="Crystal"
        :id="pet.id"
        :capsules="[
          pet.species || '',
          (pet.sex as string) || '',
          formatDate((pet.physical?.ageGroup || '') as string),
        ]"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || pet.photos?.[0]?.url"
      />
      <PetItem
        v-if="!isMobile"
        :id="pet.id"
        name="Crystal"
        :capsules="[
          pet.species || '',
          (pet.sex as string) || '',
          formatDate((pet.physical?.ageGroup || '') as string),
        ]"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || pet.photos?.[0]?.url"
      />
      <PetItem
        v-if="!isMobile"
        :id="pet.id"
        name="Crystal"
        :capsules="[
          pet.species || '',
          (pet.sex as string) || '',
          formatDate((pet.physical?.ageGroup || '') as string),
        ]"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || pet.photos?.[0]?.url"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.adopt-detail__more-friends {
  margin-top: 30px;
  width: 100%;
  background-color: var(--text-inverse);
  color: var(--text-primary);
  padding: 20px 20px 30px;
  border-radius: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  h2 {
    font-size: 1.5rem;
    margin-bottom: 16px;
  }
  .adopt-detail__more-friends__list {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    padding: 0 30px;
  }
  @media (min-width: 321px) and (max-width: 430px) {
    margin: 30px auto 0;
    width: 90%;
    padding: 20px;
    h2 {
      font-size: 1.25rem;
      margin-bottom: 12px;
    }
    .adopt-detail__more-friends__list {
      flex-direction: column;
      gap: 15px;
      padding: 0;
      align-items: center;
    }
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    width: 100%;
    max-width: 1100px;
    .adopt-detail__more-friends__list {
      padding: 0px;
    }
  }
  @media (min-width: 1441px) {
    width: 100%;
    max-width: 1500px;
    .adopt-detail__more-friends__list {
      padding: 0px;
    }
  }
}
</style>
