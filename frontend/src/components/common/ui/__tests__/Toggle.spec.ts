import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import Toggle from '../Toggle.vue'

describe('Toggle.vue', () => {
  it('renders label', () => {
    const wrapper = mount(Toggle, {
      props: { label: 'Enable notifications', modelValue: false },
    })
    expect(wrapper.text()).toContain('Enable notifications')
  })

  it('toggles value on click', async () => {
    const wrapper = mount(Toggle, {
      props: { modelValue: false },
    })

    await wrapper.find('.toggle-wrapper').trigger('click')
    expect(wrapper.emitted('update:modelValue')![0]).toEqual([true])
  })

  it('respects disabled state', async () => {
    const wrapper = mount(Toggle, {
      props: { modelValue: false, disabled: true },
    })

    await wrapper.find('.toggle-wrapper').trigger('click')
    expect(wrapper.emitted('update:modelValue')).toBeUndefined()
    expect(wrapper.find('.toggle-wrapper').classes()).toContain('disabled')
  })

  it('renders active state', () => {
    const wrapper = mount(Toggle, {
      props: { modelValue: true },
    })
    expect(wrapper.find('.toggle-switch').classes()).toContain('active')
  })
})
