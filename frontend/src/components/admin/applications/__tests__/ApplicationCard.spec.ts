import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'

import ApplicationCard, { type IApplicationItem } from '../ApplicationCard.vue'

describe('ApplicationCard.vue', () => {
    const mockApp: IApplicationItem = {
        id: '1',
        type: 'adoption',
        status: 'submitted',
        date: '2025-01-01',
        applicantName: 'Test User',
        email: 'test@example.com',
        details: {
            petName: 'Buddy',
            role: null,
            reason: null
        },
        fullApplication: {
            q1: 'Answer 1',
            firstName: 'Test',
            lastName: 'User',
            email: 'test@example.com'
        }
    }

    const mountCard = (props: Partial<IApplicationItem> & { expanded?: boolean } & Record<string, unknown> = {}) => {
        return mount(ApplicationCard, {
            props: {
                app: { ...mockApp, ...(props.app || {}) },
                expanded: false,
                isExpandedId: false,
                ...props
            },
            global: {
                stubs: {
                    Capsules: true,
                    Button: {
                        name: 'Button',
                        template: '<button class="button-stub" :title="title" @click="$emit(\'click\')"></button>',
                        props: ['title', 'onClick'],
                        emits: ['click']
                    }
                }
            }
        })
    }

    it('renders basic info correctly', () => {
        const wrapper = mountCard()
        expect(wrapper.text()).toContain('Test User')
        expect(wrapper.text()).toContain('test@example.com')
        expect(wrapper.text()).toContain('Pet: Buddy')
    })

    it('shows warning for old pending applications', () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'submitted', date: '2020-01-01' },
            expanded: true
        })
        expect(wrapper.text()).toContain('Warning')
        expect(wrapper.text()).toContain('deleted')
    })

    it('emits toggle event on click', async () => {
        const wrapper = mountCard()
        await wrapper.find('.app-card').trigger('click')
        expect(wrapper.emitted('toggle')).toBeTruthy()
    })

    // --- New Workflow Tests ---

    it('shows Start Review / Deny for submitted application', () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'submitted' },
            expanded: true
        })

        const buttons = wrapper.findAll('.button-stub')
        const titles = buttons.map(b => b.attributes('title'))
        expect(titles).toContain('Start Review')
        expect(titles).toContain('Deny')
    })

    it('emits under_review when Start Review clicked', async () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'submitted' },
            expanded: true
        })

        const btn = wrapper.findAll('.button-stub').find(b => b.attributes('title') === 'Start Review')
        await btn?.trigger('click')

        expect(wrapper.emitted('update-status')?.[0]).toEqual([expect.objectContaining({ id: '1' }), 'under_review'])
    })

    it('shows Request Video Tour for under_review application', () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'under_review' },
            expanded: true
        })

        const btn = wrapper.findAll('.button-stub').find(b => b.attributes('title') === 'Request Video Tour')
        expect(btn).toBeDefined()
    })

    it('emits video_requested when Request Video Tour clicked', async () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'under_review' },
            expanded: true
        })

        const btn = wrapper.findAll('.button-stub').find(b => b.attributes('title') === 'Request Video Tour')
        await btn?.trigger('click')

        expect(wrapper.emitted('update-status')?.[0]).toEqual([expect.objectContaining({ id: '1' }), 'video_requested'])
    })

    it('shows Approve Video / Deny Video for video_requested application', () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'video_requested' },
            expanded: true
        })

        const buttons = wrapper.findAll('.button-stub')
        const titles = buttons.map(b => b.attributes('title'))
        expect(titles).toContain('Approve Video')
        expect(titles).toContain('Deny Video')
    })

    it('emits payment_pending when Approve Video clicked', async () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'video_requested' },
            expanded: true
        })

        const btn = wrapper.findAll('.button-stub').find(b => b.attributes('title') === 'Approve Video')
        await btn?.trigger('click')

        expect(wrapper.emitted('update-status')?.[0]).toEqual([expect.objectContaining({ id: '1' }), 'payment_pending'])
    })

    it('emits rejected when Deny Video clicked', async () => {
        const wrapper = mountCard({
            app: { ...mockApp, status: 'video_requested' },
            expanded: true
        })

        const btn = wrapper.findAll('.button-stub').find(b => b.attributes('title') === 'Deny Video')
        await btn?.trigger('click')

        expect(wrapper.emitted('update-status')?.[0]).toEqual([expect.objectContaining({ id: '1' }), 'rejected'])
    })
})
