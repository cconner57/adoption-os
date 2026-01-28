import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { useAuthStore } from '../auth'

const MOCK_USER = {
    ID: 1,
    Name: 'Test User',
    Email: 'test@example.com',
    Role: 'admin'
}

describe('useAuthStore', () => {
    beforeEach(() => {
        setActivePinia(createPinia())
        vi.stubGlobal('fetch', vi.fn())
        vi.stubGlobal('localStorage', {
            getItem: vi.fn(),
            setItem: vi.fn(),
            removeItem: vi.fn()
        })

        // Mock window.location.reload
        Object.defineProperty(window, 'location', {
            value: { reload: vi.fn() },
            writable: true
        })
    })

    afterEach(() => {
        vi.restoreAllMocks()
    })

    it('initializes with null user', () => {
        const store = useAuthStore()
        expect(store.user).toBeNull()
        expect(store.isAuthenticated).toBe(false)
        expect(store.initialized).toBe(false)
    })

    it('checkAuth updates user on success', async () => {
        const store = useAuthStore()
        vi.mocked(fetch).mockResolvedValueOnce({
            ok: true,
            json: async () => ({ data: MOCK_USER })
        } as Response)

        await store.checkAuth()

        expect(store.user).toEqual(MOCK_USER)
        expect(store.isAuthenticated).toBe(true)
    })

    it('checkAuth handles failure gracefully', async () => {
        const store = useAuthStore()
        vi.mocked(fetch).mockResolvedValueOnce({
            ok: false
        } as Response)

        await store.checkAuth()

        expect(store.user).toBeNull()
    })

    it('login success sets token and user', async () => {
        const store = useAuthStore()
        vi.mocked(fetch).mockResolvedValueOnce({
            ok: true,
            json: async () => ({ token: 'abc-123', user: MOCK_USER })
        } as Response)

        const success = await store.login('test@example.com', 'password')

        expect(success).toBe(true)
        expect(localStorage.setItem).toHaveBeenCalledWith('token', 'abc-123')
        expect(store.user).toEqual(MOCK_USER)
    })

    it('login failure returns false', async () => {
        const store = useAuthStore()
        vi.mocked(fetch).mockResolvedValueOnce({
            ok: false,
            json: async () => ({ error: 'Invalid creds' })
        } as Response)

        const success = await store.login('test@example.com', 'badpass')

        expect(success).toBe(false)
        expect(store.user).toBeNull()
    })

    it('logout clears state and reloads', async () => {
        const store = useAuthStore()
        store.user = MOCK_USER

        vi.mocked(fetch).mockResolvedValueOnce({ ok: true } as Response)

        await store.logout()

        expect(fetch).toHaveBeenCalledWith('/api/users/logout', expect.any(Object))
        expect(localStorage.removeItem).toHaveBeenCalledWith('token')
        expect(store.user).toBeNull()
        expect(window.location.reload).toHaveBeenCalled()
    })

    it('initialize only runs once', async () => {
        const store = useAuthStore()
        vi.mocked(fetch).mockResolvedValue({
            ok: true,
            json: async () => ({ data: MOCK_USER })
        } as Response)

        await store.initialize()
        expect(store.initialized).toBe(true)
        expect(fetch).toHaveBeenCalledTimes(1)

        await store.initialize()
        expect(fetch).toHaveBeenCalledTimes(1) // Should not call again
    })
})
