import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { useApplications } from '../useApplications'

const MOCK_APPS = [
  { id: '1', type: 'adoption', status: 'submitted', created_at: '2023-01-01' },
  { id: '2', type: 'adoption', status: 'adopted', created_at: '2023-01-02' },
  { id: '3', type: 'surrender', status: 'submitted', created_at: '2023-01-03' },
]

describe('useApplications', () => {
    const fetchMock = vi.fn()

    beforeEach(() => {
        vi.stubGlobal('fetch', fetchMock)

        const localStorageMock = {
            getItem: vi.fn(() => 'mock-token'),
            setItem: vi.fn(),
            clear: vi.fn()
        }
        vi.stubGlobal('localStorage', localStorageMock)

        fetchMock.mockReset()
        fetchMock.mockResolvedValue({
            ok: true,
            json: async () => ({ applications: MOCK_APPS }),
        })

        // Fix: Ensure setAppBadge returns a Promise to avoid "undefined reading catch" error
        Object.defineProperty(navigator, 'setAppBadge', {
            value: vi.fn().mockImplementation(() => Promise.resolve()),
            writable: true,
            configurable: true
        })
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

  it('initializes with default state', () => {
    const { activeTab, filterStatus, loading } = useApplications()
    expect(activeTab.value).toBe('adoption')
    expect(filterStatus.value).toBe('all')
    expect(loading.value).toBe(true)
  })

  it('fetches applications correctly', async () => {
    const { fetchApplications, applications } = useApplications()

    await fetchApplications()

    expect(fetchMock).toHaveBeenCalledWith(
        expect.stringContaining('/v1/applications?page_size=100'),
        expect.objectContaining({
            headers: expect.objectContaining({
                Authorization: 'Bearer mock-token'
            })
        })
    )
    expect(applications.value).toHaveLength(3)
  })

  it('filters applications by tab', async () => {
    const { fetchApplications, activeTab, filteredApplications } = useApplications()

    await fetchApplications()

    activeTab.value = 'adoption'
    expect(filteredApplications.value).toHaveLength(2)
    expect(filteredApplications.value.every(a => a.type === 'adoption')).toBe(true)

    activeTab.value = 'surrender'
    expect(filteredApplications.value).toHaveLength(1)
    expect(filteredApplications.value[0].id).toBe('3')
  })

  it('filters applications by status', async () => {
    const { fetchApplications, activeTab, filterStatus, filteredApplications } = useApplications()

    await fetchApplications()
    activeTab.value = 'adoption'

    filterStatus.value = 'adopted'
    expect(filteredApplications.value).toHaveLength(1)
    expect(filteredApplications.value[0].status).toBe('adopted')
  })

  it('computes grouped applications', async () => {
     const { fetchApplications, pendingGroup, approvedGroup } = useApplications()
     await fetchApplications()

     expect(pendingGroup.value).toHaveLength(1) // App 1
     expect(approvedGroup.value).toHaveLength(1) // App 2
  })

  it('updates application status', async () => {
      const { updateStatus, applications } = useApplications()
      applications.value = JSON.parse(JSON.stringify(MOCK_APPS))

      const targetApp = applications.value[0]

      fetchMock.mockResolvedValueOnce({ ok: true })

      await updateStatus(targetApp, 'adopted')

      // Sorted order is 3, 2, 1. So [0] is 3.
      expect(fetchMock).toHaveBeenCalledWith(
          expect.stringContaining(`/v1/applications/3`),
          expect.objectContaining({
              method: 'PUT',
              headers: expect.objectContaining({
                Authorization: 'Bearer mock-token'
            })
          })
      )

      await new Promise(resolve => setTimeout(resolve, 0))

      expect(targetApp.status).toBe('adopted')
  })

  it('auto-switches tabs if current is empty but others have data', async () => {
      const { fetchApplications, activeTab, applications } = useApplications()

      const surrenderApps = [
          { id: '3', type: 'surrender', status: 'submitted', created_at: '2023-01-03' }
      ]

      fetchMock.mockReset()
      fetchMock.mockResolvedValue({
        ok: true,
        json: async () => ({ applications: surrenderApps }),
      })

      await fetchApplications(true)

      expect(applications.value).toHaveLength(1)
      expect(applications.value[0].type).toBe('surrender')
      expect(activeTab.value).toBe('surrender')
  })
})
