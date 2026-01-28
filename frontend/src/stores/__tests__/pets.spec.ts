import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { API_ENDPOINTS } from '../../constants/api'
import { usePetStore } from '../pets'

const MOCK_PETS = [
    { id: '1', name: 'Buddy', species: 'dog', status: 'available' },
    { id: '2', name: 'Luna', species: 'cat', status: 'available' },
]

describe('usePetStore', () => {
    let fetchMock: ReturnType<typeof vi.fn>
    let setItemMock: ReturnType<typeof vi.fn>
    let removeItemMock: ReturnType<typeof vi.fn>

    beforeEach(() => {
        setActivePinia(createPinia())
        fetchMock = vi.fn()
        vi.stubGlobal('fetch', fetchMock)

        setItemMock = vi.fn()
        removeItemMock = vi.fn()

        vi.stubGlobal('sessionStorage', {
            getItem: vi.fn(),
            setItem: setItemMock,
            removeItem: removeItemMock
        })

        vi.stubGlobal('localStorage', {
            getItem: vi.fn(() => 'mock-token')
        })
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('initializes with default state', () => {
        const store = usePetStore()
        expect(store.currentPets).toEqual([])
        expect(store.isFetching).toBe(false)
        expect(store.selectedPet).toBeNull()
    })

    it('fetching pets populates currentPets', async () => {
        const store = usePetStore()
        fetchMock.mockResolvedValueOnce({
            ok: true,
            json: async () => ({ data: MOCK_PETS })
        })

        await store.fetchPets()

        expect(fetchMock).toHaveBeenCalledWith(`${API_ENDPOINTS.PETS}?status=available&sort=age`)
        expect(store.currentPets).toEqual(MOCK_PETS)
        expect(store.isFetching).toBe(false)
    })

    it('fetch pets handles errors', async () => {
        const store = usePetStore()
        fetchMock.mockResolvedValueOnce({ ok: false })

        await expect(store.fetchPets()).rejects.toThrow('Failed to fetch pets')
        expect(store.currentPets).toEqual([])
        expect(store.isFetching).toBe(false)
    })

    it('fetch uses cache if fresh', async () => {
        const store = usePetStore()
        fetchMock.mockResolvedValue({
            ok: true,
            json: async () => ({ data: MOCK_PETS })
        })

        await store.fetchPets()
        expect(fetchMock).toHaveBeenCalledTimes(1)

        await store.fetchPets() // cached
        expect(fetchMock).toHaveBeenCalledTimes(1)

        await store.fetchPets(true) // force refresh
        expect(fetchMock).toHaveBeenCalledTimes(2)
    })

    it('fetchAdminPets calls correct endpoint with auth', async () => {
        const store = usePetStore()
        fetchMock.mockResolvedValue({
            ok: true,
            json: async () => ({ data: MOCK_PETS })
        })

        const params = new URLSearchParams({ page: '1', sort: 'name' })
        await store.fetchAdminPets(params)

        expect(fetchMock).toHaveBeenCalledWith(
            expect.stringContaining(`${API_ENDPOINTS.PETS}?page=1&sort=name`),
            expect.objectContaining({
                headers: expect.objectContaining({
                    Authorization: 'Bearer mock-token'
                })
            })
        )
        expect(store.adminPets).toEqual(MOCK_PETS)
    })

    it('updatePet updates locally and remotely', async () => {
        const store = usePetStore()
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        store.currentPets = [...MOCK_PETS] as any

        fetchMock.mockResolvedValue({ ok: true })

        const updatedPet = { ...MOCK_PETS[0], name: 'Updated Buddy' }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        await store.updatePet(updatedPet as any)

        expect(store.currentPets[0].name).toBe('Updated Buddy')
        expect(fetchMock).toHaveBeenCalledWith(
            expect.stringContaining(`/${updatedPet.id}`),
            expect.objectContaining({ method: 'PUT' })
        )
    })

    it('reverts optimistic update on failure', async () => {
        const store = usePetStore()
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        store.currentPets = [...MOCK_PETS] as any

        fetchMock.mockResolvedValueOnce({ ok: false })

        // Mock the refetch that happens on error
        fetchMock.mockResolvedValueOnce({
            ok: true,
            json: async () => ({ data: MOCK_PETS })
        })

        const updatedPet = { ...MOCK_PETS[0], name: 'Updated Buddy' }

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        await expect(store.updatePet(updatedPet as any)).rejects.toThrow('Failed to update pet')

        // Should have attempted refetch to restore state
        expect(fetchMock).toHaveBeenCalledTimes(2)
    })

    it('manages selectedPet in sessionStorage', () => {
        const store = usePetStore()
        const pet = { id: '1', species: 'dog' }

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        store.selectPet(pet as any)
        expect(store.selectedPet).toEqual(pet)
        expect(setItemMock).toHaveBeenCalledTimes(1)

        store.clearSelectedPet()
        expect(store.selectedPet).toBeNull()
        expect(removeItemMock).toHaveBeenCalledTimes(1)
    })
})
