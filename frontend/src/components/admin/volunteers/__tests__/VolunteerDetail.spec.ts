import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import type { IVolunteer } from '@/stores/mockVolunteerData'

import VolunteerDetail from '../VolunteerDetail.vue'

// Mock sub-components
vi.mock('../../common/ui/Button.vue', () => ({ default: { template: '<button><slot /></button>' } }))
vi.mock('../../common/ui/InputField.vue', () => ({ default: { template: '<input />' } }))
vi.mock('../../common/ui/Capsules.vue', () => ({ default: { template: '<div></div>' } }))
vi.mock('../../common/ui/Tabs.vue', () => ({ default: { template: '<div></div>' } }))
vi.mock('../../common/ui/Tag.vue', () => ({ default: { template: '<div><slot /></div>' } }))
vi.mock('../../common/ui/Toast.vue', () => ({ default: { template: '<div></div>' } }))
vi.mock('../ShiftForm.vue', () => ({ default: { template: '<div></div>' } }))
vi.mock('../VolunteerEditor.vue', () => ({ default: { template: '<div></div>' } }))

const mockVolunteer = {
    id: 'vol-1',
    firstName: 'John',
    lastName: 'Doe',
    role: 'volunteer',
    status: 'active',
    email: 'john@example.com',
    phone: '555-0123',
    badges: [],
    reliabilityScore: 100,
    streak: 5,
    totalHours: 20,
    joinDate: '2024-01-01',
    skills: [],
    availability: [],
    positionPreferences: []
} as unknown as IVolunteer

describe('VolunteerDetail.vue', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('renders volunteer details correctly', () => {
    const wrapper = mount(VolunteerDetail, {
      global: {
        stubs: {
           RouterLink: true
        }
      },
      props: {
        volunteer: mockVolunteer,
        selectedTags: [],
        shifts: [],
        incidents: []
      }
    })

    expect(wrapper.text()).toContain('John Doe')
    expect(wrapper.text()).toContain('active')
  })

})
