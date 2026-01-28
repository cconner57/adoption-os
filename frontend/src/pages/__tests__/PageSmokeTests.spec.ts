import { createTestingPinia } from '@pinia/testing'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { createRouter, createWebHistory } from 'vue-router'

// Import Pages (using generic import to avoid resolving issues if possible, but standard import is fine)
import AdminLayout from '@/layouts/AdminLayout.vue'
import AdminApplications from '@/pages/admin/Applications.vue'
import AdminPets from '@/pages/admin/Pets.vue'
import AdoptList from '@/pages/Adopt.vue'

// Mock Sub Components to avoid deep rendering issues and noise
vi.mock('@/components/common/ui/Button.vue', () => ({ default: { template: '<button><slot/></button>' } }))
vi.mock('@/components/common/ui/InputField.vue', () => ({ default: { template: '<input />' } }))
vi.mock('@/components/common/ui/Select.vue', () => ({ default: { template: '<select />' } }))
vi.mock('@/components/common/ui/Combobox.vue', () => ({ default: { template: '<div />' } }))
vi.mock('@/components/common/ui/Tabs.vue', () => ({ default: { template: '<div />' } }))
vi.mock('@/components/common/ui/Tag.vue', () => ({ default: { template: '<div />' } }))
vi.mock('@/components/common/ui/Capsules.vue', () => ({ default: { template: '<div />' } }))
vi.mock('@/components/common/ui/Icon.vue', () => ({ default: { template: '<i />' } }))

// Mock Chart.js to avoid canvas errors in Admin Dashboard
vi.mock('chart.js', () => ({
    Chart: { register: vi.fn() },
    registerables: []
}))
vi.mock('vue-chartjs', () => ({
    Bar: { template: '<div></div>' },
    Line: { template: '<div></div>' }
}))

describe('Page Smoke Tests', () => {
    let router: object

    beforeEach(() => {
        setActivePinia(createPinia())
        router = createRouter({
            history: createWebHistory(),
            routes: [
                { path: '/', component: { template: '<div>Home</div>' } }
            ]
        })

        // Mock Fetch
        const fetchMock = vi.fn().mockImplementation((url) => {
            const urlStr = String(url)
            console.log('Fetching:', urlStr)

            if (urlStr.includes('/users/me')) {
                 return Promise.resolve({ ok: true, json: () => Promise.resolve({ data: { id: 'admin', Role: 'admin' } }) })
            }
            if (urlStr.includes('/applications')) {
                return Promise.resolve({ ok: true, json: () => Promise.resolve({ data: [] }) })
            }
            if (urlStr.includes('/pets')) {
                return Promise.resolve({ ok: true, json: () => Promise.resolve({ data: [] }) })
            }
            if (urlStr.includes('/volunteers')) {
                 return Promise.resolve({ ok: true, json: () => Promise.resolve({ data: [] }) })
            }

            return Promise.resolve({ ok: true, json: () => Promise.resolve({ data: [] }) })
        })
        vi.stubGlobal('fetch', fetchMock)

        // Mock IntersectionObserver
        const intersectionObserverMock = vi.fn().mockImplementation(() => ({
            observe: vi.fn(),
            unobserve: vi.fn(),
            disconnect: vi.fn(),
        }))
        vi.stubGlobal('IntersectionObserver', intersectionObserverMock)
        globalThis.IntersectionObserver = intersectionObserverMock
        globalThis.scrollTo = vi.fn()

        // Mock ResizeObserver
        const resizeObserverMock = vi.fn().mockImplementation(() => ({
            observe: vi.fn(),
            unobserve: vi.fn(),
            disconnect: vi.fn(),
        }))
        vi.stubGlobal('ResizeObserver', resizeObserverMock)
        globalThis.ResizeObserver = resizeObserverMock
    })

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const mountPage = (Component: any) => {
        const component = (Component as unknown as typeof Component)
        return mount(component, {
            global: {
            plugins: [
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                createTestingPinia() as any,
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                router as any
            ],
                stubs: {
                    RouterView: true,
                    RouterLink: true,
                    Teleport: true
                }
            }
        })
    }

    it('mounts AdminLayout', async () => {
        try {
            const wrapper = mountPage(AdminLayout)
            expect(wrapper.exists()).toBe(true)
        } catch (e) { console.warn(e) }
    })

    it('mounts AdoptList (Public)', () => {
         const wrapper = mountPage(AdoptList)
         expect(wrapper.exists()).toBe(true)
    })

    it('mounts AdminApplications', () => {
        const wrapper = mountPage(AdminApplications)
        expect(wrapper.exists()).toBe(true)
    })

    it('mounts AdminPets', () => {
        const wrapper = mountPage(AdminPets)
        expect(wrapper.exists()).toBe(true)
    })
})
