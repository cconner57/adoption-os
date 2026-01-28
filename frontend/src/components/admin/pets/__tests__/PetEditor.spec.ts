import { mount, VueWrapper } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick } from 'vue'

import PetEditor from '../PetEditor.vue'

// Mock Stores
const mockSettingsStore = {
  settings: {
    pets: {
      defaultSpecies: 'cat',
      requirePhotoForPublic: true,
      defaultIntakeStatus: 'available'
    }
  }
}

vi.mock('../../../../stores/settings', () => ({
  useSettingsStore: () => mockSettingsStore
}))

// Mock Sub-Components
const MockEditorComponent = {
  template: '<div class="mock-editor"></div>',
  props: ['modelValue'],
  emits: ['update:modelValue']
}

// Robust Sidebar Stub
const MockSidebar = {
    name: 'SidebarNav',
    template: '<div class="editor-sidebar"></div>',
    props: ['modelValue', 'items'],
    emits: ['update:modelValue']
}

interface EditorVM {
    formData: {
        name: string
        species: string
        details: {
            status: string
        }
        photos: unknown[]
    }
    activeTab: string
}

describe('PetEditor Interaction', () => {
    let wrapper: VueWrapper

    beforeEach(() => {
        setActivePinia(createPinia())
    })

    const mountEditor = (props: Record<string, unknown> = {}) => {
        return mount(PetEditor, {
            props: {
                isOpen: true,
                pet: null,
                availablePets: [],
                ...props
            },
            global: {
                stubs: {
                    PetEditorBasic: MockEditorComponent,
                    PetEditorPhysical: MockEditorComponent,
                    PetEditorBehavior: MockEditorComponent,
                    PetEditorMedical: MockEditorComponent,
                    PetEditorStatus: MockEditorComponent,
                    PetEditorStory: MockEditorComponent,
                    PetEditorPhotos: MockEditorComponent,
                    PetEditorSettings: MockEditorComponent,
                    SidebarNav: MockSidebar, // Use robust stub
                    Button: {
                        name: 'Button',
                        props: ['title'],
                        template: '<button class="btn" :title="title" @click="$emit(\'click\')">{{ title }}</button>'
                    }
                }
            }
        })
    }

    it('initializes with default data for new pet', () => {
        wrapper = mountEditor()
        // Check if PetEditorBasic is rendered (it has .mock-editor class)
        const basicEditorHTML = wrapper.find('.mock-editor')
        expect(basicEditorHTML.exists()).toBe(true)

        // Logic check: Default species is 'cat' from settings
        expect((wrapper.vm as unknown as EditorVM).formData.species).toBe('cat')
    })

    it('populates data when editing existing pet', () => {
        const mockPet = {
            id: '123',
            name: 'Fluffy',
            species: 'dog',
            physical: { breed: 'Poodle' }
        }

        wrapper = mountEditor({ pet: mockPet })
        expect((wrapper.vm as unknown as EditorVM).formData.name).toBe('Fluffy')
        expect((wrapper.vm as unknown as EditorVM).formData.species).toBe('dog')
    })

    it('updates internal state when child component emits update', async () => {
        wrapper = mountEditor()
        // Find the component instance corresponding to the mock
        const child = wrapper.findComponent(MockEditorComponent)

        // Simulate child updating model
        // We must clone the data to simulate a new object ref if needed, or just modify
        const currentData = (wrapper.vm as unknown as EditorVM).formData
        const newFormData = { ...currentData, name: 'Rex' }

        child.vm.$emit('update:modelValue', newFormData)

        await nextTick()
        expect((wrapper.vm as unknown as EditorVM).formData.name).toBe('Rex')
    })

    it('validates photo requirement for available pets', async () => {
        wrapper = mountEditor()
        const vm = wrapper.vm as unknown as EditorVM

        // Ensure data reactivity
        vm.formData.details.status = 'available'
        vm.formData.photos = [] // explicit empty

        await nextTick()

        // Use vitest stubGlobal for alert if spy fails (but spy should work)
        const alertMock = vi.fn()
        vi.stubGlobal('alert', alertMock)

        // Find Save button by component or class
        // Our Button stub has class 'btn' and title attribute
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const saveButton = wrapper.findAll('.btn').find((b: any) => b.attributes('title') === 'Save Pet')
        expect(saveButton).toBeDefined()

        await saveButton!.trigger('click')

        // If alert logic depends on settings store, ensure mock is correct.
        // The mock store is static: requirePhotoForPublic: true.

        expect(alertMock).toHaveBeenCalled()
        expect(alertMock).toHaveBeenCalledWith(expect.stringContaining('Photo required'))
        expect(wrapper.emitted('save')).toBeFalsy()
    })

    it('emits save event with valid data', async () => {
        wrapper = mountEditor()
        const vm = wrapper.vm as unknown as EditorVM

        vm.formData.details.status = 'intake'
        vm.formData.name = 'Valid Pet'

        await nextTick()

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const saveButton = wrapper.findAll('.btn').find((b: any) => b.attributes('title') === 'Save Pet')
        expect(saveButton).toBeDefined()
        await saveButton!.trigger('click')

        expect(wrapper.emitted('save')).toBeTruthy()
        expect(wrapper.emitted('save')).toBeTruthy()
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        expect((wrapper.emitted('save')![0][0] as any).name).toBe('Valid Pet')
    })

    it('switches tabs', async () => {
        wrapper = mountEditor()

        expect((wrapper.vm as unknown as EditorVM).activeTab).toBe('basic')
        expect(wrapper.findComponent(MockEditorComponent).exists()).toBe(true)

        // Find stub by component now
        const sidebar = wrapper.findComponent(MockSidebar)
        expect(sidebar.exists()).toBe(true)

        // Emit update event
        sidebar.vm.$emit('update:modelValue', 'medical')

        await nextTick()

        expect((wrapper.vm as unknown as EditorVM).activeTab).toBe('medical')
    })
})
