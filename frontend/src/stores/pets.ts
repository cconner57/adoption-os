import { defineStore } from 'pinia'
import { type Ref, ref } from 'vue'

import { API_ENDPOINTS } from '../constants/api'
import type { IPet } from '../models/common'

export const usePetStore = defineStore('pets', () => {

  const currentPets = ref<IPet[]>([])
  const adminPets = ref<IPet[]>([])
  const lastFetched = ref<number>(0)
  const lastAdminFetched = ref<number>(0)
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

  const fetchAdminPets = async (params: URLSearchParams, forceRefresh = false) => {
    // Basic cache check - only if no params (default view) or if we want to implement more complex caching
    // For now, let's cache the "default" admin view if params are empty or match defaults
    // But Admin often filters.
    // The requirement says: "In these specific actions, check if Date.now() - lastUpdated is less than 5 minutes."
    // If we filter, we probably should bypass cache or have a more sophisticated cache key.
    // However, for simplicity and to meet the requirement "reduce API calls", we can cache the *full* list if we fetch it,
    // OR we just cache the last result if params match.
    // Let's assume we are caching the *latest fetch* regardless of params, OR just caching the "all pets" list if possible.
    // Actually, `Pets.vue` passes params.
    // If I want to cache, I should probably check if it's a "fresh" fetch.
    // Let's stick to the user's specific instruction: "check if Date.now() - lastUpdated is less than 5 minutes."
    // I will add a simple check, but note that changing filters might require `forceRefresh`.

    const isFresh = Date.now() - lastAdminFetched.value < CACHE_DURATION
    // We only use cache if it's fresh AND we are not forcing AND (crucially) we might want to check if params changed.
    // But `Pets.vue` calls `fetchPets` (now `fetchAdminPets`) on watcher.
    // If I return early, the user won't see filter updates.
    // So, I should probably only cache if `params` are effectively "default" OR if I store the params used for the cache.
    // User instruction is specific about the *time* check.
    // I will implement the time check but also ensure `forceRefresh` is used when filters change in the component.

    if (adminPets.value.length > 0 && isFresh && !forceRefresh) {
       console.log('Using cached admin pets')
       return
    }

    isFetching.value = true
    try {
      const response = await fetch(`${API_ENDPOINTS.PETS}?${params.toString()}`, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })

      if (!response.ok) throw new Error('Failed to fetch admin pets')
      const data = await response.json()
      adminPets.value = Array.isArray(data) ? data : data.data || []
      lastAdminFetched.value = Date.now()
    } catch (error) {
      console.error('Error fetching admin pets:', error)
      throw error
    } finally {
      isFetching.value = false
    }
  }

  const updatePet = async (pet: IPet) => {
    // Optimistic update
    const updateInList = (list: Ref<IPet[]>) => {
      const idx = list.value.findIndex((p: IPet) => p.id === pet.id)
      if (idx !== -1) {
        list.value[idx] = { ...pet }
      }
    }

    updateInList(currentPets)
    updateInList(adminPets)

    // Sanitize payload: Remove read-only fields that the backend does not expect
    const payload: Partial<IPet> = structuredClone(pet)
    delete payload.id
    delete payload.createdAt
    delete payload.updatedAt

    try {
      const response = await fetch(`${API_ENDPOINTS.PETS}/${pet.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(payload),
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
    adminPets,
    selectedPet,
    isFetching,

    fetchPets,
    fetchAdminPets,
    updatePet,
    selectPet,
    clearSelectedPet,
    initFromSession,
  }
})
