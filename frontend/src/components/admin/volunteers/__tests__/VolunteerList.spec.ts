import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import { nextTick } from 'vue'

import type { IVolunteer } from '@/stores/mockVolunteerData'

import VolunteerList from '../VolunteerList.vue'

const mockVolunteers = [
  {
    id: '1',
    firstName: 'Alice',
    lastName: 'Smith',
    role: 'Tier 1',
    status: 'active',
    email: 'alice@example.com',
    phone: '123-456-7890',
    badges: [],
    reliabilityScore: 100,
    streak: 10,
    totalHours: 50,
    joinDate: '2024-01-01',
    skills: [],
    availability: [],
    positionPreferences: [],
    address: '123 Test St',
    city: 'Test City',
    zip: '12345',
    birthday: '1990-01-01',
    allergies: '',
    emergencyContactName: 'Emergency',
    emergencyContactPhone: '111-222-3333'
  },
  {
    id: '2',
    firstName: 'Bob',
    lastName: 'Jones',
    role: 'Tier 1',
    status: 'active',
    email: 'bob@example.com',
    phone: '098-765-4321',
    badges: [],
    reliabilityScore: 90,
    streak: 5,
    totalHours: 25,
    joinDate: '2024-02-01',
    skills: [],
    availability: [],
    positionPreferences: [],
    address: '456 Test Ave',
    city: 'Test City',
    zip: '12345',
    birthday: '1992-02-02',
    allergies: '',
    emergencyContactName: 'Emergency',
    emergencyContactPhone: '111-222-3333'
  },
  {
    id: '3',
    firstName: 'Charlie',
    lastName: 'Inactive',
    role: 'Teen',
    status: 'inactive',
    email: 'charlie@test.com',
    phone: '123-456-7890',
    badges: [],
    reliabilityScore: 0,
    streak: 0,
    totalHours: 5,
    joinDate: '2023-03-01',
    skills: [],
    availability: [],
    positionPreferences: [],
    address: '789 Inactive Rd',
    city: 'Quiet Town',
    zip: '54321',
    birthday: '2005-05-05',
    allergies: '',
    emergencyContactName: 'Emergency',
    emergencyContactPhone: '111-222-3333'
  }
]

describe('VolunteerList.vue', () => {

    const mountList = (volunteers = mockVolunteers) => {
        return mount(VolunteerList, {
            props: {
                volunteers: volunteers as unknown as IVolunteer[],
                selectedId: '1'
            },
            global: {
                stubs: {
                    // We'll trust real internal logic of these simple UI components or stub them if needed.
                    // Let's use real for deeper integration, but stub if they cause issues.
                    // InputField creates an input, easy to target.
                }
            }
        })
    }

    it('renders list of active volunteers by default', () => {
        const wrapper = mountList()
        const text = wrapper.text()

        expect(text).toContain('Alice')
        expect(text).toContain('Bob')
        expect(text).not.toContain('Charlie') // Inactive
    })

    it('filters by search query', async () => {
        const wrapper = mountList()

        // Find input
        const input = wrapper.find('input')
        await input.setValue('bob')

        expect(wrapper.text()).toContain('Bob')
        expect(wrapper.text()).not.toContain('Alice')
    })

    it('filters by status (Archived)', async () => {
        const wrapper = mountList()

        // Find toggle. ButtonToggle usually emits update:modelValue
        // We can find the button with label "Archived" or value "archived"?
        // ButtonToggle renders buttons.
        // Let's try to find the button component instance OR the button element.

        // Assumption: ButtonToggle renders buttons for true/false values.
        // We can interact with component ref if we can find it,
        // OR just click the button that says "Archived".

        // Let's rely on finding by text usually.
        const archivedBtn = wrapper.findAll('button').find(b => b.text().includes('Archived'))
        if (archivedBtn) {
            await archivedBtn.trigger('click')
        } else {
             // Fallback if ButtonToggle logic is specific
             // It might be a component event we need to trigger?
             // ButtonToggle v-model="filterType".
             // We can set data directly if UI interaction is tricky.
             // We can set data directly if UI interaction is tricky.
             // eslint-disable-next-line @typescript-eslint/no-explicit-any
             ;(wrapper.vm as any).filterType = 'archived'
        }

        await nextTick()

        expect(wrapper.text()).toContain('Charlie')
        expect(wrapper.text()).not.toContain('Alice')
    })

    it('emits select event when item clicked', async () => {
        const wrapper = mountList()

        const bobItem = wrapper.findAll('.vol-item').find(el => el.text().includes('Bob'))
        await bobItem?.trigger('click')

        // Initial selectedId='1' (Alice) matches, so no initial emit.
        // Click Bob -> emit select Bob.
        expect(wrapper.emitted('select')).toBeDefined()
        expect(wrapper.emitted('select')?.[0]).toEqual([mockVolunteers[1]]) // Index 1 is Bob
    })

    it('automatically selects first item on filtered change if no selection', async () => {
         const wrapper = mountList()
         // Initially no emit because Alice (id 1) is selected and present.
         expect(wrapper.emitted('select')).toBeUndefined()

         // Search for Bob -> Alice disappears -> Bob becomes first -> Auto select Bob?
         const input = wrapper.find('input')
         await input.setValue('Bob')

         // wrapper.vm selection logic runs in watcher

         const emitted = wrapper.emitted('select')
         expect(emitted).toBeDefined()
         expect((emitted![0][0] as IVolunteer).firstName).toBe('Bob')
    })
})
