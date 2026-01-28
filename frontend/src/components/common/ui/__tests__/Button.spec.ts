import { mount } from '@vue/test-utils'
import { describe, expect, it, vi } from 'vitest'

import Button from '../Button.vue'

describe('Button.vue', () => {
  it('renders default button correctly', () => {
    const wrapper = mount(Button, {
      slots: { default: 'Click Me' },
    })
    expect(wrapper.text()).toBe('Click Me')
    expect(wrapper.classes()).toContain('variant-primary')
    expect(wrapper.classes()).toContain('theme-primary') // default theme logic
  })

  it('applies variant and theme classes', () => {
    const wrapper = mount(Button, {
      props: {
        variant: 'secondary',
        theme: 'danger',
      },
      slots: { default: 'Delete' },
    })
    expect(wrapper.classes()).toContain('variant-secondary')
    expect(wrapper.classes()).toContain('theme-danger')
  })

  it('renders legacy colors correctly', () => {
    // Test the computed logic for legacy colors
    const wrapper = mount(Button, {
      props: {
        color: 'blue'
      }
    })
    // blue -> secondary theme
    expect(wrapper.classes()).toContain('theme-secondary')
  })

  it('shows loading spinner', () => {
    const wrapper = mount(Button, {
      props: { loading: true },
    })
    expect(wrapper.find('.spinner').exists()).toBe(true)
    expect(wrapper.attributes('disabled')).toBeDefined()
  })

  it('handles disabled state', async () => {
    const onClick = vi.fn()
    const wrapper = mount(Button, {
      props: {
        disabled: true,
        onClick,
      },
    })

    await wrapper.trigger('click')
    expect(onClick).not.toHaveBeenCalled()
    expect(wrapper.attributes('disabled')).toBeDefined()
  })

  it('handles click events', async () => {
    const onClick = vi.fn()
    const wrapper = mount(Button, {
      props: { onClick },
    })

    await wrapper.trigger('click')
    expect(onClick).toHaveBeenCalled()
  })
})
