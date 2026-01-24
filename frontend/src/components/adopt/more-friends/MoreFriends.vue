
<script setup lang="ts">


import type { IPet } from '../../../models/common'
import { formatDate } from '../../../utils/common'
import { useIsMobile } from '../../../utils/useIsMobile'
import PetItem from '../../common/pet-item/PetItem.vue'

const props = defineProps<{
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
        :id="props.pet.id"
        :capsules="[
          props.pet.species || '',
          (props.pet.sex as string) || '',
          formatDate((props.pet.physical?.ageGroup || '') as string),
        ]"
        :photo="props.pet.photos?.find((p) => p.isPrimary)?.url || props.pet.photos?.[0]?.url"
      />
      <PetItem
        name="Crystal"
        :id="props.pet.id"
        :capsules="[
          props.pet.species || '',
          (props.pet.sex as string) || '',
          formatDate((props.pet.physical?.ageGroup || '') as string),
        ]"
        :photo="props.pet.photos?.find((p) => p.isPrimary)?.url || props.pet.photos?.[0]?.url"
      />
      <PetItem
        v-if="!isMobile"
        :id="props.pet.id"
        name="Crystal"
        :capsules="[
          props.pet.species || '',
          (props.pet.sex as string) || '',
          formatDate((props.pet.physical?.ageGroup || '') as string),
        ]"
        :photo="props.pet.photos?.find((p) => p.isPrimary)?.url || props.pet.photos?.[0]?.url"
      />
      <PetItem
        v-if="!isMobile"
        :id="props.pet.id"
        name="Crystal"
        :capsules="[
          props.pet.species || '',
          (props.pet.sex as string) || '',
          formatDate((props.pet.physical?.ageGroup || '') as string),
        ]"
        :photo="props.pet.photos?.find((p) => p.isPrimary)?.url || props.pet.photos?.[0]?.url"
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
  box-shadow: 0 4px 6px rgb(0 0 0 / 25%);

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

  @media (width >= 321px) and (width <= 430px) {
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

  @media (width >= 1025px) and (width <= 1440px) {
    width: 100%;
    max-width: 1100px;

    .adopt-detail__more-friends__list {
      padding: 0;
    }
  }

  @media (width >= 1441px) {
    width: 100%;
    max-width: 1500px;

    .adopt-detail__more-friends__list {
      padding: 0;
    }
  }
}
</style>
