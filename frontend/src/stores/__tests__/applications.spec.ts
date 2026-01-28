import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { API_ENDPOINTS } from '../../constants/api'
import { useApplicationsStore } from '../applications'

// Mock utils to avoid dependency
vi.mock('../../pages/admin/utils', () => ({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    mapApplicationToItem: (app: any) => ({
        id: app.id,
        type: app.type,
        status: app.status,
        created_at: app.created_at,
        // Add other required props if needed by store logic (sorting uses created_at)
    })
}))

const MOCK_APPS_RESPONSE = {
    applications: [
        { id: '1', type: 'adoption', status: 'pending', created_at: '2023-01-01' },
        { id: '2', type: 'volunteer', status: 'approved', created_at: '2023-01-02' },
        { id: '3', type: 'surrender', status: 'pending', created_at: '2023-01-03' },
    ]
}

describe('useApplicationsStore', () => {
    let fetchMock: ReturnType<typeof vi.fn>

    beforeEach(() => {
        setActivePinia(createPinia())
        fetchMock = vi.fn()
        vi.stubGlobal('fetch', fetchMock)

        vi.stubGlobal('localStorage', {
            getItem: vi.fn(() => 'mock-token'),
            setItem: vi.fn(),
            removeItem: vi.fn()
        })

        Object.defineProperty(navigator, 'setAppBadge', {
             value: vi.fn().mockResolvedValue(undefined),
             writable: true,
             configurable: true
        })

        // Default succesful fetch
        fetchMock.mockResolvedValue({
            ok: true,
            json: async () => MOCK_APPS_RESPONSE
        })
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('initializes with default state', () => {
        const store = useApplicationsStore()
        expect(store.activeTab).toBe('adoption')
        expect(store.filterStatus).toBe('all')
        expect(store.loading).toBe(true)
        expect(store.applications).toEqual([])
    })

    it('fetches applications and sets state', async () => {
        const store = useApplicationsStore()
        await store.fetchApplications()

        expect(fetchMock).toHaveBeenCalledWith(
            expect.stringContaining(API_ENDPOINTS.APPLICATIONS),
            expect.anything()
        )
        expect(store.applications).toHaveLength(3)
        expect(store.loading).toBe(false)
        expect(navigator.setAppBadge).toHaveBeenCalledWith(2) // 2 pending apps in mock (id 1 and 3?) Wait.
        // Mock has:
        // 1: pending adoption
        // 2: approved volunteer
        // 3: pending surrender
        // So pending count is 2. Correct.
    })

    it('caches applications within duration', async () => {
        const store = useApplicationsStore()

        await store.fetchApplications()
        expect(fetchMock).toHaveBeenCalledTimes(1)

        // Immediate second call should use cache
        await store.fetchApplications()
        expect(fetchMock).toHaveBeenCalledTimes(1)
    })

    it('force refresh ignores cache', async () => {
        const store = useApplicationsStore()
        await store.fetchApplications()
        expect(fetchMock).toHaveBeenCalledTimes(1)

        await store.fetchApplications(false, true) // forceRefresh = true
        expect(fetchMock).toHaveBeenCalledTimes(2)
    })

    it('filters applications correctly', async () => {
        const store = useApplicationsStore()
        await store.fetchApplications()

        store.activeTab = 'adoption'
        expect(store.filteredApplications).toHaveLength(1)
        expect(store.filteredApplications[0].id).toBe('1')

        store.activeTab = 'surrender'
        expect(store.filteredApplications).toHaveLength(1)
        expect(store.filteredApplications[0].id).toBe('3')

        store.activeTab = 'volunteer'
        store.filterStatus = 'adopted'
        expect(store.filteredApplications).toHaveLength(1)
        expect(store.filteredApplications[0].id).toBe('2')
    })

    it('updates application status', async () => {
         const store = useApplicationsStore()
         await store.fetchApplications()

         fetchMock.mockResolvedValueOnce({ ok: true })
         const app = store.applications.find(a => a.id === '1')!

         store.updateStatus(app, 'adopted')

         // Wait for promise resolution (helper fn doesn't await fetch return, but it returns undefined.
         // Implementation does not await, it chains .then.
         // We might need to wait a tick.)
         await new Promise(resolve => setTimeout(resolve, 0))

         expect(fetchMock).toHaveBeenCalledWith(
             expect.stringContaining(`/${app.id}`),
             expect.objectContaining({ method: 'PUT' })
         )
         expect(app.status).toBe('approved')
    })

    it('resends email handles success', async () => {
        const store = useApplicationsStore()
        fetchMock.mockResolvedValueOnce({ ok: true })

        vi.useFakeTimers()
        const promise = store.resendEmail('123')

        expect(store.resendingAppId).toBe('123')
        await promise

        expect(store.resendSuccessAppId).toBe('123')

        vi.advanceTimersByTime(2000)
        expect(store.resendSuccessAppId).toBeNull()
        vi.useRealTimers()
    })
})
