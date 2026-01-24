import { defineStore } from 'pinia'
import { ref } from 'vue'

import { API_ENDPOINTS } from '../constants/api'
import type { IPet } from '../models/common'

export const usePetStore = defineStore('pets', () => {

  const currentPets = ref<IPet[]>([])
  const lastFetched = ref<number>(0)
  const isFetching = ref(false)

  const selectedPet = ref<{
    id?: string
    name?: string
    petName?: string
    species: 'cat' | 'dog'
  } | null>(null)

  const STORAGE_KEY = 'adoption_pet'
  const CACHE_DURATION = 5 * 60 * 1000

  const fetchPets = async (forceRefresh = false) => {

    const isFresh = Date.now() - lastFetched.value < CACHE_DURATION
    if (currentPets.value.length > 0 && isFresh && !forceRefresh) {
      return
    }

    isFetching.value = true
    try {


    // ... (in fetchPets)
      const response = await fetch(`${API_ENDPOINTS.PETS}?status=available&sort=age`)
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

    currentPets,
    selectedPet,
    isFetching,

    fetchPets,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
