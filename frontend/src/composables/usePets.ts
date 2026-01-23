import { storeToRefs } from 'pinia'
import { computed } from 'vue'

import { usePetStore } from '../stores/pets'

export function usePets() {
  const store = usePetStore()
  const { currentPets, isFetching: loading } = storeToRefs(store)

  const spotlightPets = computed(() => {
    
    const featured = currentPets.value.filter(
      (p) => p.profileSettings && p.profileSettings.isSpotlightFeatured,
    )

    if (featured.length > 0) {
      return featured.slice(0, 4)
    }

    return []
  })

  const fetchSpotlight = async () => {
    await store.fetchPets()
  }

  return {
    spotlightPets,
    loading,
    error: null, 
    fetchSpotlight,
  }
}
