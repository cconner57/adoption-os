import { mount } from '@vue/test-utils'
import { describe, expect, it, vi } from 'vitest'

import App from '../App.vue'

// Mock vue-router
vi.mock('vue-router', () => ({
  useRoute: vi.fn(() => ({
    meta: {
      hideNavbar: false,
    },
  })),
  RouterView: { template: '<div><slot /></div>' },
}))

describe('App', () => {
  it('mounts and renders properly', () => {
    const wrapper = mount(App, {
      global: {
        stubs: {
          NavBar: true,
          PageLoader: true,
        },
      },
    })

    expect(wrapper.exists()).toBe(true)
    expect(wrapper.find('main').exists()).toBe(true)
  })
})
