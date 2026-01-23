<script setup lang="ts">
import { computed, ref, watch } from 'vue'

import type { IPet } from '../../../models/common.ts'
import { useIsMobile } from '../../../utils/useIsMobile.ts'
import PetItem from '../../common/pet-item/PetItem.vue'
import { Spinner } from '../../common/ui'

const props = defineProps<{
  pets: IPet[]
  loading: boolean
  error: string | null
}>()

const isMobile = useIsMobile()

const randomPet = ref<IPet | null>(null)
console.log('AdoptionSpotlight props.pets:', props.pets)
watch(
  [() => props.pets, isMobile],
  ([newPets, newIsMobile], [oldPets, oldIsMobile]) => {
    if (!newIsMobile) return

    const hasPets = newPets && newPets.length > 0
    if (!hasPets) return

    const justSwitchedToMobile = !oldIsMobile
    const justLoadedPets = !oldPets || oldPets.length === 0

    if (justSwitchedToMobile || justLoadedPets) {
      randomPet.value = newPets[Math.floor(Math.random() * newPets.length)]
    }
  },
  { immediate: true },
)

const displayedPets = computed((): IPet[] => {
  if (isMobile.value) {
    return randomPet.value ? [randomPet.value] : []
  }
  return props.pets
})
</script>

<template>
  <section class="adoption-spotlight">
    <h2>Adoption Spotlight</h2>
    <div v-if="loading" class="loader-container">
      <Spinner />
    </div>
    <div v-else class="pet-list">
      <PetItem
        v-for="pet in displayedPets"
        :key="pet.id"
        :name="pet.name"
        :id="(pet.slug || pet.id).toLowerCase()"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || null"
        :description="pet.descriptions?.spotlight || ''"
        :size="isMobile ? 'large' : 'medium'"
      />
    </div>
  </section>
</template>

<style scoped lang="css">
.loader-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 200px;
  width: 100%;
}
.adoption-spotlight {
  background-color: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin: -120px 0 0;
  width: 100%;
  padding: 24px 50px 40px 50px;

  & h2 {
    font-size: 2rem;
    color: var(--text-primary);
  }

  .pet-list {
    display: flex;
    gap: 3rem;
    flex-wrap: nowrap;
    overflow-x: auto;
    padding-bottom: 8px;
    -webkit-overflow-scrolling: touch;
    justify-content: center; /* Center by default on larger screens */

    /* Fix for cutoff cards: extend scroll area to edges */
    margin-left: -50px;
    margin-right: -50px;
    padding-left: 50px;
    padding-right: 50px;
  }

  @media (max-width: 430px) {
    margin: 2rem 0 0;
    padding: 1rem 2rem;
    gap: 0.5rem;
    & h2 {
      font-size: 1.5rem;
    }
    .pet-list {
      gap: 1rem;
      /* Adjust for mobile padding (2rem = 32px roughly, but verify padding above) */
      margin-left: -2rem;
      margin-right: -2rem;
      padding-left: 2rem;
      padding-right: 2rem;
      justify-content: flex-start; /* Keep left alignment for scrolling on mobile */
    }
  }
  @media (min-width: 431px) and (max-width: 768px) {
    margin: -20px 0 0;
    & h2 {
      font-size: 1.75rem;
    }
  }
  @media (min-width: 769px) and (max-width: 1024px) {
    & h2 {
      font-size: 1.75rem;
    }
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    width: 100%;
    padding: 24px 30px 30px;
    & h2 {
      font-size: 1.75rem;
    }
    .pet-list {
      gap: 2rem;
      /* Adjust for tablet landscape/desktop padding */
      margin-left: -30px;
      margin-right: -30px;
      padding-left: 30px;
      padding-right: 30px;
    }
  }
}
</style>
