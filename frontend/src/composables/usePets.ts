import { computed } from 'vue'
import type { IPet } from '../models/common.ts'
import { usePetStore } from '../stores/pets'
import { storeToRefs } from 'pinia'

export function usePets() {
  const store = usePetStore()
  const { currentPets, isFetching: loading } = storeToRefs(store)

  // Derive spotlight pets from the main list (first 4 or random 4)
  // Derive spotlight pets from the main list (filter by isSpotlightFeatured flag)
  const spotlightPets = computed(() => {
    // Filter pets that have the spotlight flag enabled
    const featured = currentPets.value.filter(
      (p) => p.profileSettings && p.profileSettings.isSpotlightFeatured,
    )

    // If we have featured pets, return them (up to 4)
    if (featured.length > 0) {
      return featured.slice(0, 4)
    }

    // Fallback: If no pets are featured, you might want to return empty or a random set.
    // Based on user request "There should only be Allison, Colby... and one other",
    // it implies we should strictly respect the flag.
    // However, if NO pets are featured, the UI might look empty.
    // Let's return empty if none match, as "strict" compliance seems requested.
    return []
  })

  // We expose a "fetch" method that delegates to the store
  const fetchSpotlight = async () => {
    await store.fetchPets()
  }

  return {
    spotlightPets,
    loading,
    error: null, // Store handles errors globally or throw, can simplify here
    fetchSpotlight,
  }
}
