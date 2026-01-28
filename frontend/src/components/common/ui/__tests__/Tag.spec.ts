import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import Tag from '../Tag.vue'

describe('Tag.vue', () => {
  it('renders label', () => {
    const wrapper = mount(Tag, {
      props: { label: 'New' },
    })
    expect(wrapper.text()).toBe('New')
  })

  it('renders slot content', () => {
    const wrapper = mount(Tag, {
      slots: { default: '<strong>Urgent</strong>' },
    })
    expect(wrapper.html()).toContain('<strong>Urgent</strong>')
  })

  it('applies variant and color classes', () => {
    const wrapper = mount(Tag, {
      props: {
        variant: 'outline',
        color: 'danger'
      },
    })
    expect(wrapper.classes()).toContain('outline')
    expect(wrapper.classes()).toContain('danger')
  })
})
