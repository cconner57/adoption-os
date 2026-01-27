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

  const updatePet = async (pet: IPet) => {
    // Optimistic update
    const index = currentPets.value.findIndex((p) => p.id === pet.id)
    if (index !== -1) {
      currentPets.value[index] = pet
    }
    if (selectedPet.value?.id === pet.id) {
       // Only update if it's the currently selected pet in detailed view
       // But wait, selectedPet is a partial object in store definition:
       // selectedPet: { id?: string, name?: string ... } | null
       // So we might not need to update it fully here, or just basic info.
       // The store typings for selectedPet are weirdly minimal compared to IPet.
       // Let's just update currentPets for now to satisfy the list view.
    }

    try {
      const response = await fetch(`${API_ENDPOINTS.PETS}/${pet.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(pet),
      })

      if (!response.ok) throw new Error('Failed to update pet')

      // Update successful, maybe fetch fresh data or just leave optimistic?
      // Leaving optimistic is fine.
    } catch (error) {
      console.error('Error updating pet:', error)
      // Revert or fetch?
      await fetchPets(true) // Revert by fetching fresh
      throw error
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
    updatePet,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
