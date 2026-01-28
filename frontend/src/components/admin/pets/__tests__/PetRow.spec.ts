import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import PetRow from '../PetRow.vue'

// Mock sub-components
vi.mock('@/components/common/ui/Button.vue', () => ({ default: { template: '<button><slot /></button>' } }))
vi.mock('@/components/common/ui/Capsules.vue', () => ({ default: { template: '<div></div>' } }))
vi.mock('@/components/common/ui/Tag.vue', () => ({ default: { template: '<div><slot /></div>' } }))
vi.mock('@/components/common/ui/Icon.vue', () => ({ default: { template: '<i></i>' } }))

const MOCK_PET = {
    id: '123',
    name: 'Buddy',
    species: 'dog',
    status: 'available',
    gender: 'male',
    age: 2,
    breed: 'Golden Retriever',
    intakeDate: '2023-01-01',
    sex: 'Male',
    physical: {
        dateOfBirth: '2021-01-01',
        size: 'Medium',
        coatLength: 'Short',
        color: 'Golden'
    },
    medical: {
        spayedOrNeutered: true,
        vaccinations: [],
        microchip: { microchipID: '123456789' }
    },
    details: {
        status: 'available',
        intakeDate: '2023-01-01',
        shelterLocation: 'Kennel 1'
    },
    photos: [],
    behavior: {},
    descriptions: {},
    sponsored: {},
    returned: {},
    adoption: {},
    foster: {}
}

describe('PetRow.vue', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('renders correctly', () => {
    const wrapper = mount(PetRow, {
      global: {
        plugins: [createPinia()],
        stubs: {
            RouterLink: true
        }
      },
      props: {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        pet: MOCK_PET as any,
        index: 0,
        visibleColumns: { name: true, photo: true, breed: true, status: true, actions: true },
        isExpanded: false,
        statusFilter: 'available'
      }
    })

    expect(wrapper.exists()).toBe(true)
    expect(wrapper.text()).toContain('Buddy')

    const link = wrapper.findComponent({ name: 'RouterLink' })
    expect(link.props().to).toEqual({ name: 'pet-details', params: { id: '123' } })
  })
})
