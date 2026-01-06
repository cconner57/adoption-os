import { defineStore } from 'pinia'
import { ref } from 'vue'

export const usePetStore = defineStore('pets', () => {
  // Matches the structure used in the rest of the app (e.g. PetAdoption.vue template uses petName)
  const selectedPet = ref<{ name?: string; petName?: string; species: 'cat' | 'dog' } | null>(null)

  const STORAGE_KEY = 'adoption_pet'

  // Initialize from session storage if available
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

  const selectPet = (pet: { name?: string; petName?: string; species: 'cat' | 'dog' }) => {
    selectedPet.value = pet
    sessionStorage.setItem(STORAGE_KEY, JSON.stringify(selectedPet.value))
  }

  const clearSelectedPet = () => {
    selectedPet.value = null
    sessionStorage.removeItem(STORAGE_KEY)
  }

  // Auto-init
  initFromSession()

  return {
    selectedPet,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
