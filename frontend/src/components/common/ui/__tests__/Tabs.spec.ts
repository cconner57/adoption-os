import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import Tabs from '../Tabs.vue'

describe('Tabs.vue', () => {
  const items = [
    { id: 'tab1', label: 'Overview' },
    { id: 'tab2', label: 'Details' },
  ]

  it('renders tabs', () => {
    const wrapper = mount(Tabs, {
      props: { items, modelValue: 'tab1' },
    })
    expect(wrapper.text()).toContain('Overview')
    expect(wrapper.text()).toContain('Details')
  })

  it('applies active class', () => {
    const wrapper = mount(Tabs, {
      props: { items, modelValue: 'tab2' },
    })
    const buttons = wrapper.findAll('button')
    // tab2 is second item
    expect(buttons[1].classes()).toContain('active')
    expect(buttons[0].classes()).not.toContain('active')
  })

  it('emits update on click', async () => {
    const wrapper = mount(Tabs, {
      props: { items, modelValue: 'tab1' },
    })

    await wrapper.findAll('button')[1].trigger('click')
    expect(wrapper.emitted('update:modelValue')![0]).toEqual(['tab2'])
  })
})
