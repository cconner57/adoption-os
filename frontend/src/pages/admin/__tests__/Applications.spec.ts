import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick } from 'vue'

import Applications from '../Applications.vue'

// Mock useApplications composable
const { updateStatusSpy, fetchApplicationsSpy } = vi.hoisted(() => ({
  updateStatusSpy: vi.fn(),
  fetchApplicationsSpy: vi.fn()
}))

vi.mock('../../../composables/useApplications', async () => {
  const { ref } = await import('vue')
  return {
    useApplications: () => ({
      activeTab: ref('adoption'),
      filterStatus: ref('all'),
      selectedYear: ref(2025),
      currentYear: 2025,
      applications: ref([]),
      loading: ref(false),
      filteredApplications: ref([]),
      pendingGroup: ref([]),
      approvedGroup: ref([]),
      deletedGroup: ref([]),
      resendingAppId: ref(null),
      resendSuccessAppId: ref(null),
      fetchApplications: fetchApplicationsSpy,
      updateStatus: updateStatusSpy,
      resendEmail: vi.fn(),
      viewOriginal: vi.fn()
    })
  }
})

describe('Applications Page Integration', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    let wrapper: any

    interface AppVM {
        applications: unknown[]
        filteredApplications: unknown[]
        pendingGroup: unknown[]
    }

    beforeEach(() => {
        setActivePinia(createPinia())
        vi.clearAllMocks()
    })

    const mountApp = () => {
        return mount(Applications, {
            global: {
                stubs: {
                    Teleport: true
                }
            }
        })
    }

    it('fetches applications on mount', () => {
        wrapper = mountApp()
        expect(fetchApplicationsSpy).toHaveBeenCalledWith(true)
    })

    it('renders pending applications', async () => {
        wrapper = mountApp()
        const vm = wrapper.vm as unknown as AppVM


        const testApp = {
            id: 'app-123',
            type: 'adoption',
            status: 'pending',
            date: '2025-01-01',
            applicantName: 'John Doe',
            email: 'john@example.com',
            details: { petName: 'Buddy' },
            fullApplication: {}
        }

        vm.applications = [testApp]
        vm.filteredApplications = [testApp] // Required to bypass empty state
        vm.pendingGroup = [testApp]

        await nextTick()

        expect(wrapper.text()).toContain('John Doe')
        expect(wrapper.text()).toContain('Buddy')
    })

    it('handles approve action', async () => {
        wrapper = mountApp()
        const vm = wrapper.vm as unknown as AppVM

        const testApp = {
            id: 'app-123',
            type: 'adoption',
            status: 'pending',
            date: '2025-01-01',
            applicantName: 'John Doe',
            email: 'john@example.com',
            details: { petName: 'Rx' },
            fullApplication: {}
        }

        vm.applications = [testApp]
        vm.filteredApplications = [testApp]
        vm.pendingGroup = [testApp]
        await nextTick()

        // Find the ApplicationCard component
        const cardComponent = wrapper.findComponent({ name: 'ApplicationCard' })
        expect(cardComponent.exists()).toBe(true)

        // Emit update-status event directly from the child component
        cardComponent.vm.$emit('update-status', testApp, 'approved')

        expect(updateStatusSpy).toHaveBeenCalledWith(testApp, 'approved')
    })

    it('handles deny action', async () => {
        wrapper = mountApp()
        const vm = wrapper.vm as unknown as AppVM

        const testApp = {
            id: 'app-456',
            type: 'adoption',
            status: 'pending',
            date: '2025-01-01',
            applicantName: 'Jane Doe',
            email: 'jane@example.com',
            details: { petName: 'Max' },
            fullApplication: {}
        }

        vm.applications = [testApp]
        vm.filteredApplications = [testApp]
        vm.pendingGroup = [testApp]
        await nextTick()

        const cardComponent = wrapper.findComponent({ name: 'ApplicationCard' })
        expect(cardComponent.exists()).toBe(true)

        // Emit update-status event
        cardComponent.vm.$emit('update-status', testApp, 'denied')

        expect(updateStatusSpy).toHaveBeenCalledWith(testApp, 'denied')
    })
})
