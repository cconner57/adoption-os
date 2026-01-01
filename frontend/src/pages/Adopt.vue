<script setup lang="ts">
import AdoptDetail from '../components/adopt/adopt-view/AdoptDetail.vue'
import AdoptSummary from '../components/adopt/adopt-view/AdoptSummary.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { mockPetsData } from '../stores/mockPetData.ts'

const props = defineProps<{ id?: string }>()
const route = useRoute()

const id = computed(() => props.id ?? (route.params.id as string | undefined))

const pets = mockPetsData

const pet = computed(() => pets.find((p) => p.id === id.value))
</script>

<template>
  <div class="adopt">
    <div class="content-wrapper">
      <div class="header" v-if="!pet">
        <h1>Find your new best friend</h1>
        <p>
          Search adoptable cats and dogs across Southern California. Every adoption helps us rescue
          another life.
        </p>
      </div>
      <main>
        <AdoptDetail v-if="pet" :pet="pet!" />
        <AdoptSummary v-else :pets="pets" />
      </main>
    </div>
  </div>
</template>

<style scoped lang="css">
.adopt {
  width: 100%;
  min-height: 100vh;
  background-color: var(--green);
  display: flex;
  justify-content: center;

  .content-wrapper {
    width: 100%;
    /* max-width: 1600px; removed to let children fill 1600px independently of padding */
    margin: 0 auto;
    padding: 8rem var(--layout-padding-side) 3rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 2rem;
    box-sizing: border-box;
  }

  .header {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    align-items: center;
    text-align: center;
    & h1 {
      font-size: 3.5rem;
      color: var(--font-color-light);
    }
    & p {
      font-size: 1.25rem;
      color: var(--font-color-light);
      min-width: 400px;
      max-width: 100%;
      font-weight: 400;
    }
  }

  @media (min-width: 0px) and (max-width: 320px) {
    /* margin-top: 1rem; removed as sticking to padding based layout */
  }
  @media (min-width: 321px) and (max-width: 430px) {
    .content-wrapper {
        padding-top: 5.5rem;
        text-align: center;
    }
    & h1 {
      font-size: 2.5rem;
      color: var(--font-color-light);
      line-height: 4rem;
    }
    & p {
      font-size: 1.05rem;
      min-width: auto;
    }
    main {
      width: 100%;
      max-width: 1600px;
    }
  }
  @media (min-width: 431px) and (max-width: 768px) {
    .content-wrapper {
        padding: 6rem var(--layout-padding-side) 2rem;
        gap: 40px;
        align-items: center;
    }
    .header {
      max-width: 580px;
      h1 {
        font-size: 2.75rem;
        line-height: 3.25rem;
      }
      p {
        font-size: 1.1rem;
      }
    }
  }
}
</style>
