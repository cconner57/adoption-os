import { beforeEach,describe, expect, it, vi } from 'vitest'
import { ref } from 'vue'

import { usePets } from '../usePets'

// Mock the store
const mockFetchPets = vi.fn()
const mockCurrentPets = ref<{ id: string; name: string; profileSettings: { isSpotlightFeatured: boolean } }[]>([])
const mockIsFetching = ref(false)

vi.mock('../../stores/pets', () => ({
    usePetStore: () => ({
        currentPets: mockCurrentPets,
        isFetching: mockIsFetching,
        fetchPets: mockFetchPets
    })
}))

describe('usePets', () => {
    beforeEach(() => {
        mockFetchPets.mockClear()
        mockCurrentPets.value = []
        mockIsFetching.value = false
    })

    it('initializes and computes spotlight pets', () => {
        mockCurrentPets.value = [
            { id: '1', name: 'Spot', profileSettings: { isSpotlightFeatured: true } },
            { id: '2', name: 'Rex', profileSettings: { isSpotlightFeatured: false } }, // Excluded
            { id: '3', name: 'Luna', profileSettings: { isSpotlightFeatured: true } },
            { id: '4', name: 'Max', profileSettings: { isSpotlightFeatured: true } },
            { id: '5', name: 'Bella', profileSettings: { isSpotlightFeatured: true } },
            { id: '6', name: 'Charlie', profileSettings: { isSpotlightFeatured: true } }, // Excluded (over limit)
        ]

        const { spotlightPets, loading } = usePets()

        expect(loading.value).toBe(false)
        expect(spotlightPets.value).toHaveLength(4)
        expect(spotlightPets.value.map(p => p.id)).toEqual(['1', '3', '4', '5'])
    })

    it('fetches spotlight calls store action', async () => {
        const { fetchSpotlight } = usePets()
        await fetchSpotlight()

        expect(mockFetchPets).toHaveBeenCalled()
    })
})
