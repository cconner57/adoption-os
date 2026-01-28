import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import InputDate from '../InputDate.vue'

describe('InputDate.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(InputDate, {
      props: {
        label: 'Date of Birth',
        modelValue: null,
      },
    })
    expect(wrapper.find('label').text()).toBe('Date of Birth')
    expect(wrapper.find('input[type="date"]').exists()).toBe(true)
  })

  it('emits update:modelValue on change', async () => {
    const wrapper = mount(InputDate, {
      props: { modelValue: '' },
    })

    await wrapper.find('input').setValue('2023-01-01')
    expect(wrapper.emitted('update:modelValue')![0]).toEqual(['2023-01-01'])
  })

  it('applies error state', () => {
    const wrapper = mount(InputDate, {
      props: {
        hasError: true,
        modelValue: '',
      },
    })
    expect(wrapper.find('.control').classes()).toContain('has-error')
  })
})
