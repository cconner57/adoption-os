import { mount } from '@vue/test-utils'
import { afterEach, describe, expect, it } from 'vitest'

import Drawer from '../Drawer.vue'

describe('Drawer.vue', () => {
  // Reset body overflow after each test
  afterEach(() => {
    document.body.style.overflow = ''
  })

  it('renders content when open', () => {
    const wrapper = mount(Drawer, {
      props: { isOpen: true, title: 'Menu' },
      global: {
        stubs: {
          Teleport: true, // Stub teleport to render in-place for easier testing
          Transition: true
        }
      },
      slots: {
        default: '<div class="content">Hello World</div>'
      }
    })

    expect(wrapper.text()).toContain('Menu')
    expect(wrapper.find('.content').exists()).toBe(true)
  })

  it('locks body scroll when open', async () => {
    const wrapper = mount(Drawer, {
      props: { isOpen: false, title: 'Menu' },
      global: { stubs: { Teleport: true } }
    })

    expect(document.body.style.overflow).toBe('')

    await wrapper.setProps({ isOpen: true })
    expect(document.body.style.overflow).toBe('hidden')

    await wrapper.setProps({ isOpen: false })
    expect(document.body.style.overflow).toBe('')
  })

  it('emits close event', async () => {
    const wrapper = mount(Drawer, {
      props: { isOpen: true, title: 'Menu' },
      global: { stubs: { Teleport: true } }
    })

    await wrapper.find('.close-btn').trigger('click')
    expect(wrapper.emitted('close')).toHaveLength(1)

    await wrapper.find('.drawer-backdrop').trigger('click')
    expect(wrapper.emitted('close')).toHaveLength(2)
  })
})
