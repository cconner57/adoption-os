<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed, nextTick, onMounted,ref } from 'vue'
import { useRoute } from 'vue-router'

import AdoptDetail from '../components/adopt/adopt-view/AdoptDetail.vue'
import AdoptSummary from '../components/adopt/adopt-view/AdoptSummary.vue'
import FilterPanel from '../components/adopt/FilterPanel.vue'
import { usePetStore } from '../stores/pets'

const props = defineProps<{ id?: string }>()
const route = useRoute()
const store = usePetStore()
const { currentPets } = storeToRefs(store)

const id = computed(() => props.id ?? (route.params.id as string | undefined))
const isFilterPanelOpen = ref(false)

const activeFilter = ref('All')
const advancedFilters = ref({
  age: [] as string[],
  size: [] as string[],
  sex: '',
  goodWith: [] as string[],
})

const applyAdvancedFilters = (newFilters: typeof advancedFilters.value) => {
  advancedFilters.value = newFilters
  isFilterPanelOpen.value = false
}

const clearFilters = () => {
  advancedFilters.value = {
    age: [],
    size: [],
    sex: '',
    goodWith: [],
  }
}

const resetAllFilters = () => {
  activeFilter.value = 'All'
  clearFilters()
  isFilterPanelOpen.value = false
}

onMounted(() => {
  store.fetchPets()
})

const filteredPets = computed(() => {
  let result = currentPets.value

  if (activeFilter.value !== 'All') {
    result = result.filter((p) => p.species.toLowerCase() === activeFilter.value.toLowerCase())
  }

  const { age, size, sex, goodWith } = advancedFilters.value

  if (age.length > 0) {
    result = result.filter((p) => p.physical.ageGroup && age.includes(p.physical.ageGroup))
  }

  if (size.length > 0) {
    result = result.filter((p) => p.physical.size && size.includes(p.physical.size))
  }

  if (sex) {
    result = result.filter((p) => p.sex && p.sex.toLowerCase() === sex)
  }

  if (goodWith.length > 0) {
    result = result.filter((p) => {
      
      return goodWith.every((trait) => {
        if (trait === 'kids') return p.behavior.isGoodWithKids
        if (trait === 'dogs') return p.behavior.isGoodWithDogs
        if (trait === 'cats') return p.behavior.isGoodWithCats
        return false
      })
    })
  }

  return result
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

const pet = computed(() => {
  const param = id.value
  if (!param) return undefined
  return currentPets.value.find((p) => p.id === param || p.slug === param)
})
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
        <div class="species-group">
          <button
            class="reset-btn"
            :class="{ active: activeFilter === 'All' }"
            @click="resetAllFilters"
          >
            View All Pets
          </button>
          <button :class="{ active: activeFilter === 'Cat' }" @click="setFilter('Cat')">
            Cats
          </button>
          <button :class="{ active: activeFilter === 'Dog' }" @click="setFilter('Dog')">
            Dogs
          </button>
        </div>
        <div class="mobile-break"></div>
        <div class="divider"></div>
        <button
          class="filter-btn"
          :class="{ active: isFilterPanelOpen }"
          @click="isFilterPanelOpen = !isFilterPanelOpen"
        >
          Filters
          <span v-if="Object.values(advancedFilters).flat().filter(Boolean).length" class="badge">
            {{ Object.values(advancedFilters).flat().filter(Boolean).length }}
          </span>
        </button>
      </div>

      <FilterPanel
        :isOpen="isFilterPanelOpen"
        :currentFilters="advancedFilters"
        @close="isFilterPanelOpen = false"
        @apply="applyAdvancedFilters"
        @clear="clearFilters"
      />

      <main>
        <AdoptDetail v-if="pet" :pet="pet!" />
        <AdoptSummary v-else-if="filteredPets.length > 0" :pets="filteredPets" />
        <div v-else class="empty-state">
          <h2>No pets found</h2>
          <p>We couldn't find any friends matching that filter.</p>
          <button class="reset-btn" @click="resetAllFilters">View All Pets</button>
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
  align-items: center;
  flex-wrap: wrap;
  justify-content: center;

  .species-group {
    display: flex;
    gap: 12px;
    align-items: center;
    flex-wrap: wrap; 
    justify-content: center;
  }

  .divider {
    width: 1px;
    height: 24px;
    background: rgb(255 255 255 / 30%);
    margin: 0 4px;
  }

  .mobile-break {
    display: none;
    width: 100%;
  }

  button {
    all: unset;
    padding: 8px 20px;
    border-radius: 20px;
    background-color: rgb(255 255 255 / 20%);
    color: var(--text-inverse);
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    border: 2px solid transparent;
    display: flex;
    align-items: center;
    gap: 8px;
    white-space: nowrap;

    &:hover {
      background-color: rgb(255 255 255 / 30%);
    }

    &.active {
      background-color: var(--text-inverse);
      color: var(--color-primary);
    }
  }

  .filter-btn {
    position: relative;

    .badge {
      background: var(--text-inverse);
      color: var(--color-primary);
      border-radius: 50%;
      width: 20px;
      height: 20px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.75rem;
      font-weight: 800;
    }
  }
}

.adopt {
  width: 100%;
  min-height: 100vh;
  background-color: var(--color-primary);
  display: flex;
  justify-content: center;

  .content-wrapper {
    width: 100%;
    margin: 0 auto;
    padding: 8rem var(--layout-padding-side) 3rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
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
      color: var(--text-inverse);
      min-width: 360px;
    }

    & p {
      font-size: 1.25rem;
      color: var(--text-inverse);
      min-width: 340px;
      max-width: 100%;
      font-weight: 400;
      margin-bottom: 2rem;
    }
  }

  @media (width >= 0) and (width <= 320px) {
    
  }

  @media (width >= 321px) and (width <= 430px) {
    .content-wrapper {
      padding-top: 5.5rem;
      text-align: center;
    }

    & h1 {
      font-size: 2rem;
      color: var(--text-inverse);
      line-height: 1.2;
      margin-bottom: 12px;
    }

    & p {
      font-size: 1rem;
      min-width: auto;
      max-width: 280px; 
      margin: 0 auto; 
      padding: 0;
      line-height: 1.5;
    }

    main {
      width: 100%;
      max-width: 1600px;
    }

    .filters {
      flex-wrap: wrap;
      justify-content: center;
    }
  }
  
  @media (width <= 600px) {
    .filters {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 16px;
      overflow-x: visible;

      .species-group {
        display: flex;
        flex-flow: row nowrap;
        gap: 8px;
        width: 100%;
        justify-content: center; 
      }

      .mobile-break {
        display: none; 
      }

      .divider {
        width: 100%;
        max-width: 200px; 
        height: 1px;
        background: rgb(255 255 255 / 20%);
        margin: 0;
      }

      button {
        padding: 8px 16px;
        font-size: 0.95rem;
        white-space: nowrap;
      }
    }
  }

  @media (width >= 431px) and (width <= 768px) {
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

    .filters {
      flex-wrap: nowrap;
      justify-content: flex-start;
      overflow-x: auto;
      padding-bottom: 8px; 
      width: 100%;
      -webkit-overflow-scrolling: touch;
      padding-left: 16px; 
      padding-right: 16px;
      box-sizing: border-box;

      &::-webkit-scrollbar {
        display: none;
      }
      
      -ms-overflow-style: none; 
      scrollbar-width: none; 
    }

    .filters button {
      white-space: nowrap;
      flex-shrink: 0;
    }
  }
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  color: var(--text-inverse);
  margin-top: 2rem;
  padding: 3rem;
  background: rgb(255 255 255 / 10%);
  border-radius: 24px;
  backdrop-filter: blur(8px);
  border: 1px solid rgb(255 255 255 / 20%);
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
    background-color: var(--text-inverse);
    color: var(--color-primary);
    padding: 12px 32px;
    border-radius: 30px;
    font-weight: 600;
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
    box-shadow: 0 4px 12px rgb(0 0 0 / 10%);

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 6px 16px rgb(0 0 0 / 15%);
    }

    &:active {
      transform: translateY(0);
    }
  }
}
</style>
