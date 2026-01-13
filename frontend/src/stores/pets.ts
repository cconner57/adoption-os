import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { IPet } from '../models/common'

export const usePetStore = defineStore('pets', () => {
  // State
  const currentPets = ref<IPet[]>([])
  const lastFetched = ref<number>(0)
  const isFetching = ref(false)

  const selectedPet = ref<{
    id?: string
    name?: string
    petName?: string
    species: 'cat' | 'dog'
  } | null>(null)

  // Constants
  const STORAGE_KEY = 'adoption_pet'
  const CACHE_DURATION = 5 * 60 * 1000 // 5 minutes

  // Actions
  const fetchPets = async (forceRefresh = false) => {
    // If we have data and it's fresh, and not forcing refresh, return early
    const isFresh = Date.now() - lastFetched.value < CACHE_DURATION
    if (currentPets.value.length > 0 && isFresh && !forceRefresh) {
      return
    }

    isFetching.value = true
    try {
      // NOTE: We always fetch *all* available pets for the cache.
      // Filtering is done client-side or via specific OTHER actions if we wanted server-side filtering.
      // For this simple cache, getting all available pets is efficient enough (~1-2KB compressed).
      const response = await fetch(`${import.meta.env.VITE_API_URL}/pets?status=available&sort=age`)
      if (!response.ok) throw new Error('Failed to fetch pets')

      const data = await response.json()
      currentPets.value = Array.isArray(data) ? data : data.data || []
      lastFetched.value = Date.now()
    } catch (error) {
      console.error('Error fetching pets:', error)
      throw error
    } finally {
      isFetching.value = false
    }
  }

  // Session Storage Logic (Existing)
  const initFromSession = () => {
    const stored = sessionStorage.getItem(STORAGE_KEY)
    if (stored) {
      try {
        selectedPet.value = JSON.parse(stored)
      } catch (e) {
        console.error('Failed to parse selected pet from session storage', e)
        sessionStorage.removeItem(STORAGE_KEY)
      }
    }
  }

  const selectPet = (pet: {
    id?: string
    name?: string
    petName?: string
    species: 'cat' | 'dog'
  }) => {
    selectedPet.value = pet
    sessionStorage.setItem(STORAGE_KEY, JSON.stringify(selectedPet.value))
  }

  const clearSelectedPet = () => {
    selectedPet.value = null
    sessionStorage.removeItem(STORAGE_KEY)
  }

  initFromSession()

  return {
    // State
    currentPets,
    selectedPet,
    isFetching,

    // Actions
    fetchPets,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
