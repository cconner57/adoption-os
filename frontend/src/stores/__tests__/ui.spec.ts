import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick } from 'vue'

import { useUIStore } from '../ui'

describe('useUIStore', () => {
    let getItemMock: ReturnType<typeof vi.fn>
    let setItemMock: ReturnType<typeof vi.fn>

    beforeEach(() => {
        setActivePinia(createPinia())

        getItemMock = vi.fn()
        setItemMock = vi.fn()

        vi.stubGlobal('localStorage', {
            getItem: getItemMock,
            setItem: setItemMock,
            removeItem: vi.fn()
        })
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('initializes with default state', () => {
        const store = useUIStore()
        expect(store.isLoading).toBe(false)
        expect(store.adminState.isSidebarOpen).toBe(false)
    })

    it('loads state from localStorage if version matches', () => {
        const storedState = JSON.stringify({ isSidebarOpen: true, version: 1 })
        getItemMock.mockReturnValue(storedState)

        const store = useUIStore()

        expect(store.adminState.isSidebarOpen).toBe(true)
    })

    it('resets state if version mismatch', () => {
        const storedState = JSON.stringify({ isSidebarOpen: true, version: 999 })
        getItemMock.mockReturnValue(storedState)

        const store = useUIStore()

        // Should rely on default because version mismatch cleared it (effectively, though implementation detail might vary slightly on how it resets,
        // the store implementation usually keeps default if parse fails or version mismatch)
        // Actually looking at code: it clears storage but doesn't explicitly reset state variable to default in the else block,
        // but since `adminState` is initialized with defaults, it just stays default.
        expect(store.adminState.isSidebarOpen).toBe(false)
        expect(localStorage.removeItem).toHaveBeenCalledWith('admin_ui_state')
    })

    it('updates loading state', () => {
        const store = useUIStore()
        store.startLoading()
        expect(store.isLoading).toBe(true)
        store.stopLoading()
        expect(store.isLoading).toBe(false)
    })

    it('persists state changes to localStorage', async () => {
        const store = useUIStore()

        store.adminState.isSidebarOpen = true
        await nextTick() // Wait for watcher

        expect(setItemMock).toHaveBeenCalledWith(
            'admin_ui_state',
            expect.stringContaining('"isSidebarOpen":true')
        )
        expect(setItemMock).toHaveBeenCalledWith(
            'admin_ui_state',
            expect.stringContaining('"version":1')
        )
    })
})
