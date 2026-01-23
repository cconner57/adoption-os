import { flushPromises,mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { defineComponent, h } from 'vue'

import { useMediaQuery } from '../useMediaQuery.js'

describe('useMediaQuery', () => {
  let matchMediaMock: ReturnType<typeof vi.fn>

  beforeEach(() => {
    
    matchMediaMock = vi.fn().mockImplementation((query: string) => ({
      matches: false,
      media: query,
      onchange: null,
      addListener: vi.fn(),
      removeListener: vi.fn(),
      addEventListener: vi.fn(),
      removeEventListener: vi.fn(),
      dispatchEvent: vi.fn(),
    }))

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ;(window as any).matchMedia = matchMediaMock
  })

  it('should return false initially when media query does not match', async () => {
    const TestComponent = defineComponent({
      setup() {
        const isMobile = useMediaQuery('(max-width: 440px)')
        return () => h('div', isMobile.value ? 'mobile' : 'desktop')
      },
    })

    const wrapper = mount(TestComponent)
    await flushPromises()
    expect(wrapper.text()).toBe('desktop')
  })

  it('should return true when media query matches', async () => {
    
    const addEventListener = vi.fn()
    const removeEventListener = vi.fn()
    
    matchMediaMock.mockImplementation((query: string) => ({
      matches: true,
      media: query,
      onchange: null,
      addListener: vi.fn(),
      removeListener: vi.fn(),
      addEventListener,
      removeEventListener,
      dispatchEvent: vi.fn(),
    }))

    const TestComponent = defineComponent({
      setup() {
        const isMobile = useMediaQuery('(max-width: 440px)')
        return () => h('div', isMobile.value ? 'mobile' : 'desktop')
      },
    })

    const wrapper = mount(TestComponent)
    await flushPromises()
    expect(wrapper.text()).toBe('mobile')
  })

  it('should add event listener on mount', () => {
    const addEventListener = vi.fn()
    matchMediaMock.mockImplementation((query: string) => ({
      matches: false,
      media: query,
      onchange: null,
      addListener: vi.fn(),
      removeListener: vi.fn(),
      addEventListener,
      removeEventListener: vi.fn(),
      dispatchEvent: vi.fn(),
    }))

    const TestComponent = defineComponent({
      setup() {
        useMediaQuery('(max-width: 440px)')
        return () => h('div', 'test')
      },
    })

    mount(TestComponent)
    expect(addEventListener).toHaveBeenCalledWith('change', expect.any(Function))
  })

  it('should remove event listener on unmount', () => {
    const removeEventListener = vi.fn()
    matchMediaMock.mockImplementation((query: string) => ({
      matches: false,
      media: query,
      onchange: null,
      addListener: vi.fn(),
      removeListener: vi.fn(),
      addEventListener: vi.fn(),
      removeEventListener,
      dispatchEvent: vi.fn(),
    }))

    const TestComponent = defineComponent({
      setup() {
        useMediaQuery('(max-width: 440px)')
        return () => h('div', 'test')
      },
    })

    const wrapper = mount(TestComponent)
    wrapper.unmount()
    expect(removeEventListener).toHaveBeenCalledWith('change', expect.any(Function))
  })

  it('should handle SSR environment gracefully', () => {
    const originalMatchMedia = window.matchMedia
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    delete (window as any).matchMedia

    const TestComponent = defineComponent({
      setup() {
        const isMobile = useMediaQuery('(max-width: 440px)')
        return () => h('div', isMobile.value ? 'mobile' : 'desktop')
      },
    })

    expect(() => mount(TestComponent)).not.toThrow()

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ;(window as any).matchMedia = originalMatchMedia
  })
})
