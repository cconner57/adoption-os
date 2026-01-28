import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import SuccessModal from '../SuccessModal.vue'

describe('SuccessModal.vue', () => {
  it('renders when visible', () => {
    const wrapper = mount(SuccessModal, {
      props: {
        isVisible: true,
        title: 'Great Job',
        message: 'You did it'
      },
    })
    expect(wrapper.text()).toContain('Great Job')
    expect(wrapper.text()).toContain('You did it')
  })

  it('does not render when hidden', () => {
    const wrapper = mount(SuccessModal, {
      props: { isVisible: false },
    })
    expect(wrapper.find('.modal-overlay').exists()).toBe(false)
  })

  it('emits close on background click', async () => {
    const wrapper = mount(SuccessModal, {
      props: { isVisible: true },
    })

    await wrapper.find('.modal-overlay').trigger('click')
    expect(wrapper.emitted('close')).toHaveLength(1)
  })

  it('emits close on button click', async () => {
    const wrapper = mount(SuccessModal, {
      props: { isVisible: true },
    })

    await wrapper.find('button').trigger('click')
    expect(wrapper.emitted('close')).toHaveLength(1)
  })
})
