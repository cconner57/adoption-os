import { ref } from 'vue'
import type { IPet } from '../models/common.ts'
import { API_ENDPOINTS } from '../constants/api'

export function usePets() {
  const spotlightPets = ref<IPet[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  const fetchSpotlight = async () => {
    loading.value = true
    try {
      const response = await fetch(API_ENDPOINTS.PET_SPOTLIGHT)
      if (!response.ok) throw new Error('Failed to fetch spotlight pets')

      spotlightPets.value = await response.json()
    } catch (err: unknown) {
      if (err instanceof Error) {
        error.value = err.message
      } else {
        error.value = 'An unknown error occurred'
      }
    } finally {
      loading.value = false
    }
  }

  return {
    spotlightPets,
    loading,
    error,
    fetchSpotlight,
  }
}
