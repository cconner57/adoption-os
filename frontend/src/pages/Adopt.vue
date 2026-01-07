<script setup lang="ts">
import AdoptDetail from '../components/adopt/adopt-view/AdoptDetail.vue'
import AdoptSummary from '../components/adopt/adopt-view/AdoptSummary.vue'
import { computed, ref, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { mockPetsData } from '../stores/mockPetData.ts'

const props = defineProps<{ id?: string }>()
const route = useRoute()

const id = computed(() => props.id ?? (route.params.id as string | undefined))

const activeFilter = ref('All')

const filteredPets = computed(() => {
  if (activeFilter.value === 'All') return mockPetsData
  return mockPetsData.filter((p) => p.species.toLowerCase() === activeFilter.value.toLowerCase())
})

const setFilter = (filter: string) => {
  if (!document.startViewTransition) {
    activeFilter.value = filter
    return
  }

  document.startViewTransition(async () => {
    activeFilter.value = filter
    await nextTick()
  })
}

const pets = mockPetsData // Keep for finding single pet
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
      <div class="filters" v-if="!pet">
        <button :class="{ active: activeFilter === 'All' }" @click="setFilter('All')">All</button>
        <button :class="{ active: activeFilter === 'Cat' }" @click="setFilter('Cat')">Cats</button>
        <button :class="{ active: activeFilter === 'Dog' }" @click="setFilter('Dog')">Dogs</button>
      </div>
      <main>
        <AdoptDetail v-if="pet" :pet="pet!" />
        <AdoptSummary v-else-if="filteredPets.length > 0" :pets="filteredPets" />
        <div v-else class="empty-state">
          <h2>No pets found</h2>
          <p>We couldn't find any friends matching that filter.</p>
          <button class="reset-btn" @click="setFilter('All')">View All Pets</button>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped lang="css">
.filters {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;

  button {
    all: unset;
    padding: 8px 20px;
    border-radius: 20px;
    background-color: rgba(255, 255, 255, 0.2);
    color: var(--font-color-light);
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    border: 2px solid transparent;

    &:hover {
      background-color: rgba(255, 255, 255, 0.3);
    }

    &.active {
      background-color: var(--white);
      color: var(--green);
    }
  }
}

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
      min-width: 360px;
    }
    & p {
      font-size: 1.25rem;
      color: var(--font-color-light);
      min-width: 340px;
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
      font-size: 2rem;
      color: var(--font-color-light);
      line-height: 1.2;
      margin-bottom: 12px;
    }
    & p {
      font-size: 1rem;
      min-width: auto;
      max-width: 280px; /* Constrain width */
      margin: 0 auto; /* Center it */
      padding: 0;
      line-height: 1.5;
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

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  color: var(--font-color-light);
  margin-top: 2rem;
  padding: 3rem;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 24px;
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  max-width: 400px;

  h2 {
    font-size: 1.75rem;
    margin-bottom: 0.5rem;
    font-weight: 600;
  }

  p {
    font-size: 1.1rem;
    opacity: 0.9;
    margin-bottom: 2rem;
    line-height: 1.5;
  }

  .reset-btn {
    all: unset;
    background-color: var(--white);
    color: var(--green);
    padding: 12px 32px;
    border-radius: 30px;
    font-weight: 600;
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 6px 16px rgba(0, 0, 0, 0.15);
    }

    &:active {
      transform: translateY(0);
    }
  }
}
</style>
