import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import Combobox from '../Combobox.vue'

describe('Combobox.vue', () => {
  const options = [
    { label: 'Apple', value: 'apple' },
    { label: 'Banana', value: 'banana' },
  ]

  it('renders correctly', () => {
    const wrapper = mount(Combobox, {
      props: {
        modelValue: null,
        options,
        placeholder: 'Search fruit',
      },
    })
    expect(wrapper.find('input').attributes('placeholder')).toBe('Search fruit')
  })

  it('filters options based on search', async () => {
    const wrapper = mount(Combobox, {
      props: { modelValue: null, options },
    })

    await wrapper.find('input').setValue('ban')
    await wrapper.find('input').trigger('input') // Open menu

    const items = wrapper.findAll('.dropdown-item')
    expect(items.length).toBe(1)
    expect(items[0].text()).toContain('Banana')
  })

  it('selects an option', async () => {
    const wrapper = mount(Combobox, {
      props: { modelValue: null, options },
    })

    await wrapper.find('input').trigger('focus') // Open menu
    await wrapper.findAll('.dropdown-item')[0].trigger('mousedown')

    expect(wrapper.emitted('update:modelValue')![0]).toEqual(['apple'])
  })

  it('supports custom creation', async () => {
    const wrapper = mount(Combobox, {
      props: {
        modelValue: null,
        options,
        allowCreate: true
      },
    })

    await wrapper.find('input').setValue('Cherry')
    await wrapper.find('input').trigger('input')

    const createOption = wrapper.find('.create-option')
    expect(createOption.exists()).toBe(true)
    expect(createOption.text()).toContain('Create "Cherry"')

    await createOption.trigger('mousedown')
    expect(wrapper.emitted('update:modelValue')![0]).toEqual(['Cherry'])
  })

  it('handles multiple selection with tags', async () => {
    const wrapper = mount(Combobox, {
      props: {
        modelValue: ['apple'],
        options,
        multiple: true
      },
    })

    expect(wrapper.find('.tag').text()).toContain('Apple')

    // Add another
    await wrapper.find('input').trigger('focus')
    await wrapper.findAll('.dropdown-item')[1].trigger('mousedown') // Select Banana

    expect(wrapper.emitted('update:modelValue')![0][0]).toEqual(['apple', 'banana'])
  })

  it('removes tag', async () => {
    const wrapper = mount(Combobox, {
      props: {
        modelValue: ['apple'],
        options,
        multiple: true
      },
    })

    await wrapper.find('.remove-tag').trigger('click')
    expect(wrapper.emitted('update:modelValue')![0][0]).toEqual([])
  })
})
