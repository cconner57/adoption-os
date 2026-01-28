import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import Checkbox from '../Checkbox.vue'

describe('Checkbox.vue', () => {
  it('renders label', () => {
    const wrapper = mount(Checkbox, {
      props: { label: 'Agree to terms' },
    })
    expect(wrapper.text()).toContain('Agree to terms')
  })

  it('toggles value on click', async () => {
    const wrapper = mount(Checkbox, {
      props: { modelValue: false },
    })

    await wrapper.find('.checkbox-wrapper').trigger('click')
    expect(wrapper.emitted('update:modelValue')![0]).toEqual([true])
  })

  it('respects disabled state', async () => {
    const wrapper = mount(Checkbox, {
      props: { modelValue: false, disabled: true },
    })

    await wrapper.find('.checkbox-wrapper').trigger('click')
    expect(wrapper.emitted('update:modelValue')).toBeUndefined()
    expect(wrapper.find('.checkbox-wrapper').classes()).toContain('disabled')
  })

  it('renders checked state', () => {
    const wrapper = mount(Checkbox, {
      props: { modelValue: true },
    })
    expect(wrapper.find('.checkbox-wrapper').classes()).toContain('checked')
  })
})
