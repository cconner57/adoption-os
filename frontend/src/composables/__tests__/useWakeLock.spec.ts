import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick } from 'vue'

import { useWakeLock } from '../useWakeLock'

describe('useWakeLock', () => {
    const requestMock = vi.fn()
    const releaseMock = vi.fn()


    // Mock WakeLockSentinel
    const sentinelMock = {
        release: releaseMock,
        addEventListener: vi.fn()
    }

    beforeEach(() => {
        releaseMock.mockReset()
        requestMock.mockReset().mockResolvedValue(sentinelMock)

        // Mock navigator.wakeLock
        const wakeLockMock = {
            request: requestMock
        }

        Object.defineProperty(navigator, 'wakeLock', {
            value: wakeLockMock,
            writable: true,
            configurable: true
        })
    })

    afterEach(() => {
        // @ts-expect-error - restoring navigator
        delete navigator.wakeLock
    })

// ...

    it('requests lock successfully', async () => {
        const { requestLock, isActive } = useWakeLock()

        // Mock navigator.wakeLock
        const mockSentinel = {
            release: vi.fn(),
        }
        const mockWakeLock = {
            request: vi.fn().mockResolvedValue(mockSentinel)
        }
        Object.defineProperty(navigator, 'wakeLock', {
            value: mockWakeLock,
            writable: true,
            configurable: true
        })

        await requestLock()

        expect(mockWakeLock.request).toHaveBeenCalledWith('screen')
        expect(isActive.value).toBe(true)

        // Clean up
        // @ts-expect-error - restoring navigator
        delete navigator.wakeLock
    })

    it('handles release manually', async () => {
        const { isActive, releaseLock } = useWakeLock()

        // Mock navigator.wakeLock and sentinel
        const mockRelease = vi.fn()
        const mockSentinel = { release: mockRelease, onrelease: null }
        const mockWakeLock = {
            request: vi.fn().mockResolvedValue(mockSentinel)
        }
        Object.defineProperty(navigator, 'wakeLock', {
            value: mockWakeLock,
            writable: true,
            configurable: true
        })

        // Simulate a lock being active
        isActive.value = true
        // @ts-expect-error - Mocking internal property
        useWakeLock()._wakeLockSentinel.value = mockSentinel

        await releaseLock()
        expect(mockRelease).toHaveBeenCalled()
        expect(isActive.value).toBe(false)

        // Clean up
        // @ts-expect-error - restoring navigator
        delete navigator.wakeLock
    })

    it('handles system release event', async () => {
        const { isActive, requestLock } = useWakeLock()

        // Mock navigator.wakeLock
        // eslint-disable-next-line no-unused-vars, @typescript-eslint/no-explicit-any
        let _onReleaseCallback: ((this: WakeLockSentinel, ev: Event) => any) | null = null
        const mockSentinel = {
            release: vi.fn(),
            // eslint-disable-next-line no-unused-vars
            set onrelease(cb: ((this: WakeLockSentinel, ev: Event) => void) | null) {
                _onReleaseCallback = cb
            },
            get onrelease() {
                return _onReleaseCallback
            }
        }
        const mockWakeLock = {
            request: vi.fn().mockResolvedValue(mockSentinel)
        }
        Object.defineProperty(navigator, 'wakeLock', {
            value: mockWakeLock,
            writable: true,
            configurable: true
        })

        await requestLock()
        expect(isActive.value).toBe(true)

        // Simulate release event
        if (_onReleaseCallback) {
            // @ts-expect-error - mocking event
            _onReleaseCallback.call(mockSentinel, new Event('release'))
        }
        await nextTick() // Allow Vue to react to state changes

        expect(isActive.value).toBe(false)

        // Clean up
        // @ts-expect-error - restoring navigator
        delete navigator.wakeLock
    })
})
