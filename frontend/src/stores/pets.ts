import { defineStore } from 'pinia'
import { ref } from 'vue'

export const usePetStore = defineStore('pets', () => {
  const selectedPet = ref<{
    id?: string
    name?: string
    petName?: string
    species: 'cat' | 'dog'
  } | null>(null)

  const STORAGE_KEY = 'adoption_pet'

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
    selectedPet,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
